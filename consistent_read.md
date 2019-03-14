[consistent_read](https://dev.mysql.com/doc/refman/5.6/en/glossary.html#glos_consistent_read)
```
consistent read
A read operation that uses snapshot information to present query results based on a point in time, regardless of changes performed by other transactions running at the same time. If queried data has been changed by another transaction, the original data is reconstructed based on the contents of the undo log. This technique avoids some of the locking issues that can reduce concurrency by forcing transactions to wait for other transactions to finish.

With REPEATABLE READ isolation level, the snapshot is based on the time when the first read operation is performed. With READ COMMITTED isolation level, the snapshot is reset to the time of each consistent read operation.

Consistent read is the default mode in which InnoDB processes SELECT statements in READ COMMITTED and REPEATABLE READ isolation levels. Because a consistent read does not set any locks on the tables it accesses, other sessions are free to modify those tables while a consistent read is being performed on the table.

For technical details about the applicable isolation levels, see Section 14.7.2.3, “Consistent Nonlocking Reads”.

See Also concurrency, isolation level, locking, READ COMMITTED, REPEATABLE READ, snapshot, transaction, undo log.
```

### 一致性读
- 一种读取操作，它使用快照信息基于某个时间点显示查询结果，而不管同时运行的其他事务所执行的更改。(这句机翻)  
如果这个查询的信息已经被其他事务修改,原来的数据会通过undo log 重新构造出来,  
这种技术可以避免了一个事务等待另外一个事务去完成再执行的这种使用锁而导致的并发降低的情况.  
