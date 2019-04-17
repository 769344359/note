

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
