mysql 的select 有limit 子句
[文档地址](https://dev.mysql.com/doc/refman/5.7/en/select.html)

>The LIMIT clause can be used to constrain the number of rows returned by the SELECT statement. LIMIT takes one or two numeric arguments, which must both be nonnegative integer constants, with these exceptions:
limit 子句可以被限制select 语句返回的行的数量.limit子句使用一个或者两个参数,除了以下这种情况之外,参数必须是非负的

>Within prepared statements, LIMIT parameters can be specified using ? placeholder markers.
第一种例外是prepare语句,limit子句的参数可以使用`?`作为占位符

>Within stored programs, LIMIT parameters can be specified using integer-valued routine parameters or local variables.
第二种是储存过程,limit子句的参数可以使用integer例程参数或者临时变量(local variables)

>With two arguments, the first argument specifies the offset of the first row to return, and the second specifies the maximum number of rows to return. The offset of the initial row is 0 (not 1):
当使用两个参数的时候,第一个参数指定第一个返回的行的偏置值,第二个参数指定返回行的最大值.偏置值(offset)的初始值是0不是1.


> For compatibility with PostgreSQL, MySQL also supports the LIMIT row_count OFFSET offset syntax.
为了兼容postgreSql 的语法,Mysql 支持`LIMIT 最大行数 OFFSET  偏置值` 的语法
