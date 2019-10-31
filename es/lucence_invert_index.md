```
abstract class TermsHashPerField implements Comparable<TermsHashPerField> {
  ...
  final IntBlockPool intPool;
  final ByteBlockPool bytePool;
  final ByteBlockPool termBytePool;
  ...
```


倒排索引的创建首先是依赖这三个Pool来创建的
