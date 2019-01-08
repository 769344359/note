. packet length(3字节), 包体的长度 
. packet number(1字节), 从0开始的递增的

payload 


https://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::Handshake

```
36 00 00 00 0a 35 2e 35    2e 32 2d 6d 32 00 0b 00    6....5.5.2-m2...
00 00 64 76 48 40 49 2d    43 4a 00 ff f7 08 02 00    ..dvH@I-CJ......
00 00 00 00 00 00 00 00    00 00 00 00 00 2a 34 64    .............*4d
7c 63 5a 77 6b 34 5e 5d    3a 00                      |cZwk4^]:.
```
