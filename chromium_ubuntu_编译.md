我的环境 Ubuntu 18.04 

1 准备好梯子 shadowsocks    
2 比较坑的是  gclient runhook 是不支持sock5 所以在sslocal之外还要加一层http代理,不然gclient runhook 会失败   
gclient 的链路
```
glcient----xxx----sslocal----你的梯子
```
