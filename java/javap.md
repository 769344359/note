相关命令
```
javap -c Test.class
```


```
class com.hello.Test {
  com.hello.Test();
    Code:
       0: aload_0
       1: invokespecial #1                  // Method java/lang/Object."<init>":()V
       4: return

  public static void main(java.lang.String[]);
    Code:
       0: new           #2                  // class java/io/File
       3: dup
       4: ldc           #3                  // String a.txt
       6: invokespecial #4                  // Method java/io/File."<init>":(Ljava/lang/String;)V
       9: astore_1
      10: return
}
```
