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
```


SELECT * FROM tbl LIMIT 5,10;  # Retrieve rows 6-15
To retrieve all rows from a certain offset up to the end of the result set, you can use some large number for the second parameter. This statement retrieves all rows from the 96th row to the last:

SELECT * FROM tbl LIMIT 95,18446744073709551615;
With one argument, the value specifies the number of rows to return from the beginning of the result set:

SELECT * FROM tbl LIMIT 5;     # Retrieve first 5 rows
In other words, LIMIT row_count is equivalent to LIMIT 0, row_count.

For prepared statements, you can use placeholders. The following statements will return one row from the tbl table:

SET @a=1;
PREPARE STMT FROM 'SELECT * FROM tbl LIMIT ?';
EXECUTE STMT USING @a;
The following statements will return the second to sixth row from the tbl table:

SET @skip=1; SET @numrows=5;
PREPARE STMT FROM 'SELECT * FROM tbl LIMIT ?, ?';
EXECUTE STMT USING @skip, @numrows;
For compatibility with PostgreSQL, MySQL also supports the LIMIT row_count OFFSET offset syntax.
```
