遇到的问题: 
当时想要一个守护进程去跑一下发邮件给相关的人,因为怕守护进程会挂,所以在crontab 上加上了每分钟启动一次,然后在邮件进程里加了相关的文件锁.

调试和部署的时候都好好的,但是执行久了就发现文件锁居然没有锁住,同时在跑的有5-10个进程.按道理如果一开始就没有锁住,每分钟启动一个进程一个小时至少都有60个进程,但是实际上是一天后同时跑进程只有几个,而且我收到启动了一个进程,结果是获取不到锁所以直接退出了.所以一开始肯定是获得了锁.

然后看了一下glibc 的相关文档(虽然golang不使用glibc 但是封装系统调用都是差不多)

[文档链接](http://man7.org/linux/man-pages/man2/fcntl.2.html)

>If a process **closes** any file descriptor referring to a file, then
          all of the process's locks on that file are released, regardless
          of the file descriptor(s) on which the locks were obtained.  This
          is bad: it means that a process can lose its locks on a file such
          as /etc/passwd or /etc/mtab when for some reason a library func‐
          tion decides to open, read, and close the same file.

如果关闭了关联那个锁的文件的描述符,锁就会释放

-----------
当时想了很久锁失效了只有一个可能,文件被我关闭了.
----

`runtime.SetFinalizer` 

```
(gdb) bt
#0  os.newFile (fd=7, kind=1, name=..., ~r3=0x7) at /usr/local/go/src/os/file_unix.go:102
#1  0x0000000000475605 in os.openFileNolog (flag=0, name=..., perm=0, ~r3=0x0, ~r4=...)
    at /usr/local/go/src/os/file_unix.go:216
#2  0x0000000000474517 in os.OpenFile (flag=0, name=..., perm=0, ~r3=0x0, ~r4=...)
    at /usr/local/go/src/os/file.go:284
#3  0x0000000000474405 in os.Open (name=..., ~r1=0x0, ~r2=...) at /usr/local/go/src/os/file.go:265
#4  0x00000000004af4da in main.openFile (path=..., ~r1=...) at /data/big_utils/crontab/test.go:12
#5  0x00000000004af5b0 in main.main () at /data/big_utils/crontab/test.go:17

```

```
// newFile is like NewFile, but if called from OpenFile or Pipe
// (as passed in the kind parameter) it tries to add the file to
// the runtime poller.
func newFile(fd uintptr, name string, kind newFileKind) *File {
	fdi := int(fd)
	if fdi < 0 {
		return nil
	}
	f := &File{&file{
		pfd: poll.FD{
			Sysfd:         fdi,
			IsStream:      true,
			ZeroReadIsEOF: true,
		},
		name:        name,
		stdoutOrErr: fdi == 1 || fdi == 2,
	}}

	pollable := kind == kindOpenFile || kind == kindPipe || kind == kindNonBlock

	// Don't try to use kqueue with regular files on FreeBSD.
	// It crashes the system unpredictably while running all.bash.
	// Issue 19093.
	// If the caller passed a non-blocking filedes (kindNonBlock),
	// we assume they know what they are doing so we allow it to be
	// used with kqueue.
	if runtime.GOOS == "freebsd" && kind == kindOpenFile {
		pollable = false
	}

	// On Darwin, kqueue does not work properly with fifos:
	// closing the last writer does not cause a kqueue event
	// for any readers. See issue #24164.
	if runtime.GOOS == "darwin" && kind == kindOpenFile {
		var st syscall.Stat_t
		if err := syscall.Fstat(fdi, &st); err == nil && st.Mode&syscall.S_IFMT == syscall.S_IFIFO {
			pollable = false
		}
	}

	if err := f.pfd.Init("file", pollable); err != nil {
		// An error here indicates a failure to register
		// with the netpoll system. That can happen for
		// a file descriptor that is not supported by
		// epoll/kqueue; for example, disk files on
		// GNU/Linux systems. We assume that any real error
		// will show up in later I/O.
	} else if pollable {
		// We successfully registered with netpoll, so put
		// the file into nonblocking mode.
		if err := syscall.SetNonblock(fdi, true); err == nil {
			f.nonblock = true
		}
	}

	runtime.SetFinalizer(f.file, (*file).close)
	return f
}
```
