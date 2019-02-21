```
https://go101.org/article/channel.html
```

我现在遇到的编程语言都是保证单线程内的内容经过指令重排都能保证单线程内的顺序是一样的

```
Concurrent computations may share resources, generally memory resource. There are some circumstances may happen in a concurrent computing.
In the same period of one computation is writing data to a memory segment, another computation is reading data from the same memory segment. Then the integrity of the data read by the other computation might be not preserved.
In the same period of one computation is writing data to a memory segment, another computation is also writing data to the same memory segment. Then the integrity of the data stored at the memory segment might be not preserved.
```

`读/写`   `写/读` 以及`写/写`是不能交换顺序的,或者说只要有写的操作,就不能交换顺序.

我们首先介绍一下偏序的概念:  
偏序满足条件: 
- 反对称
- 自反
- 可传递
