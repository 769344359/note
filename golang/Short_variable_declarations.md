```
A short variable declaration uses the syntax:

ShortVarDecl = IdentifierList ":=" ExpressionList .
It is shorthand for a regular variable declaration with initializer expressions but no types:

"var" IdentifierList = ExpressionList .
```



golang 的隐式转换非常少,一般都是显式转换,现在是我找到的一个隐式转换的例子:
```
An untyped constant has a default type which is the type to which the constant is implicitly converted in contexts where a typed value is required, for instance, in a short variable declaration such as i := 0 where there is no explicit type. The default type of an untyped constant is bool, rune, int, float64, complex128 or string respectively, depending on whether it is a boolean, rune, integer, floating-point, complex, or string constant.
```
