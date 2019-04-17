

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
