```
(gdb) bt
#0  write () at ../sysdeps/unix/syscall-template.S:84
#1  0x00007ffff556779a in handleWrite (fd=1, buf=0x7ffff7fcf2d0, len=1) at /home/dinosaur/jdk8/jdk/src/solaris/native/java/io/io_util_md.c:164
#2  0x00007ffff556710a in writeBytes (env=0x7ffff000c210, this=0x7ffff7fd1400, bytes=0x7ffff7fd13f8, off=0, len=1, append=0 '\000', fid=0x47e1043)
    at /home/dinosaur/jdk8/jdk/src/share/native/java/io/io_util.c:189
#3  0x00007ffff555a79c in Java_java_io_FileOutputStream_writeBytes (env=0x7ffff000c210, this=0x7ffff7fd1400, bytes=0x7ffff7fd13f8, off=0, len=1, append=0 '\000')
    at /home/dinosaur/jdk8/jdk/src/solaris/native/java/io/FileOutputStream_md.c:70
#4  0x00007fffe10298dc in ?? ()
#5  0x00007ffff7fd1410 in ?? ()
#6  0x00007ffff672dd43 in JVM_ArrayCopy (env=0x7ffff000c210, ignored=0x7ffff7fd1400, src=0x7ffff7fd13f8, src_pos=0, dst=0x7f00f6265bea, dst_pos=1, length=0)
    at /home/dinosaur/jdk8/hotspot/src/share/vm/prims/jvm.cpp:298
#7  0x00007fffe1007500 in ?? ()
#8  0x0000000000000000 in ?? ()

```
```
Thread 2 "java" hit Breakpoint 3, handleWrite (fd=1, buf=0x7ffff7fcf270, len=12) at /home/dinosaur/jdk8/jdk/src/solaris/native/java/io/io_util_md.c:164
164	    RESTARTABLE(write(fd, buf, len), result);
(gdb) p buf
$3 = (const void *) 0x7ffff7fcf270
(gdb) p  (char *)buf
$4 = 0x7ffff7fcf270 "Hello, World\377\177"

```
