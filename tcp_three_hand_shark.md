[相关链接](https://zhangbinalan.gitbooks.io/protocol/content/ipxie_yi_tou_bu.html)
```c
dinosaur@dinosaur-X550VXK:/usr/local/nginx/sbin$ sudo tcpdump -i lo port 80 -X 
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on lo, link-type EN10MB (Ethernet), capture size 262144 bytes
14:49:37.670947 IP localhost.36986 > localhost.http: Flags [S], seq 3354874316, win 43690, options [mss 65495,sackOK,TS val 1456314863 ecr 0,nop,wscale 7], length 0
	0x0000:  4500 003c 71ee 4000 4006 cacb 7f00 0001  E..<q.@.@.......
	0x0010:  7f00 0001 907a 0050 c7f7 51cc 0000 0000  .....z.P..Q.....
	0x0020:  a002 aaaa fe30 0000 0204 ffd7 0402 080a  .....0..........
	0x0030:  56cd 99ef 0000 0000 0103 0307            V...........
14:49:37.670976 IP localhost.http > localhost.36986: Flags [S.], seq 2607525835, ack 3354874317, win 43690, options [mss 65495,sackOK,TS val 1456314863 ecr 1456314863,nop,wscale 7], length 0
	0x0000:  4500 003c 0000 4000 4006 3cba 7f00 0001  E..<..@.@.<.....
	0x0010:  7f00 0001 0050 907a 9b6b afcb c7f7 51cd  .....P.z.k....Q.
	0x0020:  a012 aaaa fe30 0000 0204 ffd7 0402 080a  .....0..........
	0x0030:  56cd 99ef 56cd 99ef 0103 0307            V...V.......
14:49:37.671004 IP localhost.36986 > localhost.http: Flags [.], ack 1, win 342, options [nop,nop,TS val 1456314863 ecr 1456314863], length 0
	0x0000:  4500 0034 71ef 4000 4006 cad2 7f00 0001  E..4q.@.@.......
	0x0010:  7f00 0001 907a 0050 c7f7 51cd 9b6b afcc  .....z.P..Q..k..
	0x0020:  8010 0156 fe28 0000 0101 080a 56cd 99ef  ...V.(......V...
	0x0030:  56cd 99ef                                V...
14:49:37.673193 IP localhost.36986 > localhost.http: Flags [P.], seq 1:480, ack 1, win 342, options [nop,nop,TS val 1456314865 ecr 1456314863], length 479: HTTP: GET / HTTP/1.1
	0x0000:  4500 0213 71f0 4000 4006 c8f2 7f00 0001  E...q.@.@.......
	0x0010:  7f00 0001 907a 0050 c7f7 51cd 9b6b afcc  .....z.P..Q..k..
	0x0020:  8018 0156 0008 0000 0101 080a 56cd 99f1  ...V........V...
	0x0030:  56cd 99ef 4745 5420 2f20 4854 5450 2f31  V...GET./.HTTP/1
	0x0040:  2e31 0d0a 486f 7374 3a20 3132 372e 302e  .1..Host:.127.0.
	0x0050:  302e 310d 0a43 6f6e 6e65 6374 696f 6e3a  0.1..Connection:
	0x0060:  206b 6565 702d 616c 6976 650d 0a43 6163  .keep-alive..Cac
	0x0070:  6865 2d43 6f6e 7472 6f6c 3a20 6d61 782d  he-Control:.max-
	0x0080:  6167 653d 300d 0a55 7067 7261 6465 2d49  age=0..Upgrade-I
	0x0090:  6e73 6563 7572 652d 5265 7175 6573 7473  nsecure-Requests
	0x00a0:  3a20 310d 0a55 7365 722d 4167 656e 743a  :.1..User-Agent:
	0x00b0:  204d 6f7a 696c 6c61 2f35 2e30 2028 5831  .Mozilla/5.0.(X1
	0x00c0:  313b 204c 696e 7578 2078 3836 5f36 3429  1;.Linux.x86_64)
	0x00d0:  2041 7070 6c65 5765 624b 6974 2f35 3337  .AppleWebKit/537
	0x00e0:  2e33 3620 284b 4854 4d4c 2c20 6c69 6b65  .36.(KHTML,.like
	0x00f0:  2047 6563 6b6f 2920 4368 726f 6d65 2f36  .Gecko).Chrome/6
	0x0100:  342e 302e 3332 3832 2e31 3637 2053 6166  4.0.3282.167.Saf
	0x0110:  6172 692f 3533 372e 3336 0d0a 4163 6365  ari/537.36..Acce
	0x0120:  7074 3a20 7465 7874 2f68 746d 6c2c 6170  pt:.text/html,ap
	0x0130:  706c 6963 6174 696f 6e2f 7868 746d 6c2b  plication/xhtml+
	0x0140:  786d 6c2c 6170 706c 6963 6174 696f 6e2f  xml,application/
	0x0150:  786d 6c3b 713d 302e 392c 696d 6167 652f  xml;q=0.9,image/
	0x0160:  7765 6270 2c69 6d61 6765 2f61 706e 672c  webp,image/apng,
	0x0170:  2a2f 2a3b 713d 302e 380d 0a41 6363 6570  */*;q=0.8..Accep
	0x0180:  742d 456e 636f 6469 6e67 3a20 677a 6970  t-Encoding:.gzip
	0x0190:  2c20 6465 666c 6174 652c 2062 720d 0a41  ,.deflate,.br..A
	0x01a0:  6363 6570 742d 4c61 6e67 7561 6765 3a20  ccept-Language:.
	0x01b0:  656e 2d55 532c 656e 3b71 3d30 2e39 0d0a  en-US,en;q=0.9..
	0x01c0:  4966 2d4e 6f6e 652d 4d61 7463 683a 2022  If-None-Match:."
	0x01d0:  3562 3035 3763 6338 2d32 3634 220d 0a49  5b057cc8-264"..I
	0x01e0:  662d 4d6f 6469 6669 6564 2d53 696e 6365  f-Modified-Since
	0x01f0:  3a20 5765 642c 2032 3320 4d61 7920 3230  :.Wed,.23.May.20
	0x0200:  3138 2031 343a 3338 3a30 3020 474d 540d  18.14:38:00.GMT.
	0x0210:  0a0d 0a                                  ...
14:49:37.673227 IP localhost.http > localhost.36986: Flags [.], ack 480, win 350, options [nop,nop,TS val 1456314865 ecr 1456314865], length 0
	0x0000:  4500 0034 8fee 4000 4006 acd3 7f00 0001  E..4..@.@.......
	0x0010:  7f00 0001 0050 907a 9b6b afcc c7f7 53ac  .....P.z.k....S.
	0x0020:  8010 015e fe28 0000 0101 080a 56cd 99f1  ...^.(......V...
	0x0030:  56cd 99f1                                V...
14:49:37.673422 IP localhost.http > localhost.36986: Flags [P.], seq 1:182, ack 480, win 350, options [nop,nop,TS val 1456314865 ecr 1456314865], length 181: HTTP: HTTP/1.1 304 Not Modified
	0x0000:  4500 00e9 8fef 4000 4006 ac1d 7f00 0001  E.....@.@.......
	0x0010:  7f00 0001 0050 907a 9b6b afcc c7f7 53ac  .....P.z.k....S.
	0x0020:  8018 015e fedd 0000 0101 080a 56cd 99f1  ...^........V...
	0x0030:  56cd 99f1 4854 5450 2f31 2e31 2033 3034  V...HTTP/1.1.304
	0x0040:  204e 6f74 204d 6f64 6966 6965 640d 0a53  .Not.Modified..S
	0x0050:  6572 7665 723a 206e 6769 6e78 2f31 2e31  erver:.nginx/1.1
	0x0060:  332e 3132 0d0a 4461 7465 3a20 5375 6e2c  3.12..Date:.Sun,
	0x0070:  2030 3120 4a75 6c20 3230 3138 2030 363a  .01.Jul.2018.06:
	0x0080:  3439 3a33 3720 474d 540d 0a4c 6173 742d  49:37.GMT..Last-
	0x0090:  4d6f 6469 6669 6564 3a20 5765 642c 2032  Modified:.Wed,.2
	0x00a0:  3320 4d61 7920 3230 3138 2031 343a 3338  3.May.2018.14:38
	0x00b0:  3a30 3020 474d 540d 0a43 6f6e 6e65 6374  :00.GMT..Connect
	0x00c0:  696f 6e3a 206b 6565 702d 616c 6976 650d  ion:.keep-alive.
	0x00d0:  0a45 5461 673a 2022 3562 3035 3763 6338  .ETag:."5b057cc8
	0x00e0:  2d32 3634 220d 0a0d 0a                   -264"....
14:49:37.673444 IP localhost.36986 > localhost.http: Flags [.], ack 182, win 350, options [nop,nop,TS val 1456314865 ecr 1456314865], length 0
	0x0000:  4500 0034 71f1 4000 4006 cad0 7f00 0001  E..4q.@.@.......
	0x0010:  7f00 0001 907a 0050 c7f7 53ac 9b6b b081  .....z.P..S..k..
	0x0020:  8010 015e fe28 0000 0101 080a 56cd 99f1  ...^.(......V...
	0x0030:  56cd 99f1                                V...
14:50:24.638108 IP localhost.36986 > localhost.http: Flags [.], ack 182, win 350, options [nop,nop,TS val 1456314865 ecr 1456314865], length 0
	0x0000:  4500 0034 71f2 4000 4006 cacf 7f00 0001  E..4q.@.@.......
	0x0010:  7f00 0001 907a 0050 c7f7 53ab 9b6b b081  .....z.P..S..k..
	0x0020:  8010 015e fe28 0000 0101 080a 56cd 99f1  ...^.(......V...
	0x0030:  56cd 99f1                                V...
14:50:24.638134 IP localhost.http > localhost.36986: Flags [.], ack 480, win 350, options [nop,nop,TS val 1456361830 ecr 1456314865], length 0
	0x0000:  4500 0034 8ff0 4000 4006 acd1 7f00 0001  E..4..@.@.......
	0x0010:  7f00 0001 0050 907a 9b6b b081 c7f7 53ac  .....P.z.k....S.
	0x0020:  8010 015e fe28 0000 0101 080a 56ce 5166  ...^.(......V.Qf
	0x0030:  56cd 99f1                                V...
14:50:42.735013 IP localhost.http > localhost.36986: Flags [F.], seq 182, ack 480, win 350, options [nop,nop,TS val 1456379927 ecr 1456314865], length 0
	0x0000:  4500 0034 8ff1 4000 4006 acd0 7f00 0001  E..4..@.@.......
	0x0010:  7f00 0001 0050 907a 9b6b b081 c7f7 53ac  .....P.z.k....S.
	0x0020:  8011 015e fe28 0000 0101 080a 56ce 9817  ...^.(......V...
	0x0030:  56cd 99f1                                V...
14:50:42.778093 IP localhost.36986 > localhost.http: Flags [.], ack 183, win 350, options [nop,nop,TS val 1456379927 ecr 1456379927], length 0
	0x0000:  4500 0034 71f3 4000 4006 cace 7f00 0001  E..4q.@.@.......
	0x0010:  7f00 0001 907a 0050 c7f7 53ac 9b6b b082  .....z.P..S..k..
	0x0020:  8010 015e fe28 0000 0101 080a 56ce 9817  ...^.(......V...
	0x0030:  56ce 9817                                V...
14:51:28.126088 IP localhost.36986 > localhost.http: Flags [.], ack 183, win 350, options [nop,nop,TS val 1456379927 ecr 1456379927], length 0
	0x0000:  4500 0034 71f4 4000 4006 cacd 7f00 0001  E..4q.@.@.......
	0x0010:  7f00 0001 907a 0050 c7f7 53ab 9b6b b082  .....z.P..S..k..
	0x0020:  8010 015e fe28 0000 0101 080a 56ce 9817  ...^.(......V...
	0x0030:  56ce 9817                                V...
14:51:28.126118 IP localhost.http > localhost.36986: Flags [.], ack 480, win 350, options [nop,nop,TS val 1456425319 ecr 1456379927], length 0
	0x0000:  4500 0034 6dd1 4000 4006 cef0 7f00 0001  E..4m.@.@.......
	0x0010:  7f00 0001 0050 907a 9b6b b082 c7f7 53ac  .....P.z.k....S.
	0x0020:  8010 015e efe3 0000 0101 080a 56cf 4967  ...^........V.Ig
	0x0030:  56ce 9817                                V...
14:52:11.129030 IP localhost.36986 > localhost.http: Flags [F.], seq 480, ack 183, win 350, options [nop,nop,TS val 1456468322 ecr 1456425319], length 0
	0x0000:  4500 0034 71f5 4000 4006 cacc 7f00 0001  E..4q.@.@.......
	0x0010:  7f00 0001 907a 0050 c7f7 53ac 9b6b b082  .....z.P..S..k..
	0x0020:  8011 015e fe28 0000 0101 080a 56cf f162  ...^.(......V..b
	0x0030:  56cf 4967                                V.Ig
14:52:11.129060 IP localhost.http > localhost.36986: Flags [R], seq 2607526018, win 0, length 0
	0x0000:  4500 0028 8014 4000 4006 bcb9 7f00 0001  E..(..@.@.......
	0x0010:  7f00 0001 0050 907a 9b6b b082 0000 0000  .....P.z.k......
	0x0020:  5004 0000 d525 0000                      P....%..

```
> 第一个包  

这是我从上面的截取的一个包也是一个三次握手的第一个包　包好syn　标志位
```
14:49:37.670947 IP localhost.36986 > localhost.http: Flags [S], seq 3354874316, win 43690, options [mss 65495,sackOK,TS val 1456314863 ecr 0,nop,wscale 7], length 0
	0x0000:  4500 003c 71ee 4000 4006 cacb 7f00 0001  E..<q.@.@.......
	0x0010:  7f00 0001 907a 0050 c7f7 51cc 0000 0000  .....z.P..Q.....
	0x0020:  a002 aaaa fe30 0000 0204 ffd7 0402 080a  .....0..........
	0x0030:  56cd 99ef 0000 0000 0103 0307            V...........
```


这个包里面`0x907a` 前面的都是　ip 头，  
从　`0x907a` 后的都是tcp头   

>  然后下面我们一个一个字节分析整个报文

因为 16 = 2^4 所以两位16进制数可以代表一个字节

- 第一个字节  0x45  

ip 头的头4位代表协议，4 代表ipv4 ，如果是ipv6 则是6，由于我这是ipv4 的协议所以是4  
而0x4500 的4 就是代表ipv4 


ip 头的第二个4位代表 ip报文头的长度，倍数是4 字节， 比如当该字段是5 的时候代表ip报头的长度是20字节（因为4*5=20 字节）,本次报文第一个字节的0x4500 的5 代表的是 ip报文的报文头是20字节  

- 第二个字节 0x00

第二个字节代表 tos，0x00 代表普通

- 第三和第四个字节 的16位一共两个字节代表 ip 报文的长度（字节数， 包括ip报头和报文）

0x003c 即十进制的16*3+12 =48+12 = 60 字节 

- 第五六两个字节 0x71ee    是identify 和分片标志位一起标志是否是同一个分片

- 0x4000  第七八个字节的 16 位 ，前3位叫flag 标志是否分片 ，后面13位是用来标志分片的相对偏移，单位是 8 字节

0x4000 的头三位是 010，后面13位是0

- 0x4006 第九个字节代表ttl  这里的例子是 0x40,第10个字节是上层传输层的协议  tcp 是0x6  

- 0xcacb 十一十二字节 是ip报文头的头部校验（只对ip头做校验） 

- 7f00 0001   后面重复的两个 7f00 0001 分别是 源ip 和目的ip

以上一共20字节就是ip头的所有内容
