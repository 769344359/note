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


## Short variable declarations  和 regular variable declaration  的区别:  
A short variable declaration uses the syntax:

ShortVarDecl = IdentifierList ":=" ExpressionList .
It is shorthand for a regular variable declaration with initializer expressions but no types:

"var" IdentifierList = ExpressionList .
i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w := os.Pipe(fd)  // os.Pipe() returns two values
_, y, _ := coord(p)  // coord() returns three values; only interested in y coordinate
Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared earlier in the same block (or the parameter lists if the block is the function body) with the same type, and at least one of the non-blank variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original.

field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // redeclares offset
a, a := 1, 2                              // illegal: double declaration of a or no new variable if a was declared elsewhere
Short variable declarations may appear only inside functions. In some contexts such as the initializers for "if", "for", or "switch" statements, they can be used to declare local temporary variables.
