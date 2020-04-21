```
(gdb) bt
#0  JOIN_CACHE_BNL::join_matching_records (this=0x7fa2ecabef50, skip_last=false) at /home/ubuntu/mysql-server/sql/sql_join_buffer.cc:1908
#1  0x000056107e327aa1 in JOIN_CACHE::join_records (this=0x7fa2ecabef50, skip_last=false) at /home/ubuntu/mysql-server/sql/sql_join_buffer.cc:1785
#2  0x000056107e329f91 in JOIN_CACHE::end_send (this=0x7fa2ecabef50) at /home/ubuntu/mysql-server/sql/sql_join_buffer.h:535
#3  0x000056107e2de693 in sub_select_op (join=0x7fa2ecabd098, qep_tab=0x7fa2ecabed80, end_of_records=true) at /home/ubuntu/mysql-server/sql/sql_executor.cc:3262
#4  0x000056107e2de821 in sub_select (join=0x7fa2ecabd098, qep_tab=0x7fa2ecabec98, end_of_records=true) at /home/ubuntu/mysql-server/sql/sql_executor.cc:3432
#5  0x000056107e2de2c4 in do_select (join=0x7fa2ecabd098) at /home/ubuntu/mysql-server/sql/sql_executor.cc:3153
#6  0x000056107e2d3538 in JOIN::exec (this=0x7fa2ecabd098) at /home/ubuntu/mysql-server/sql/sql_executor.cc:336
#7  0x000056107e48a4d0 in SELECT_LEX_UNIT::execute (this=0x7fa2ec17ebd8, thd=0x7fa2ec006fa0) at /home/ubuntu/mysql-server/sql/sql_union.cc:1631
#8  0x000056107e3ddf9a in Sql_cmd_dml::execute_inner (this=0x7fa2ecabccb0, thd=0x7fa2ec006fa0) at /home/ubuntu/mysql-server/sql/sql_select.cc:910
#9  0x000056107e3dd685 in Sql_cmd_dml::execute (this=0x7fa2ecabccb0, thd=0x7fa2ec006fa0) at /home/ubuntu/mysql-server/sql/sql_select.cc:715
#10 0x000056107e374f03 in mysql_execute_command (thd=0x7fa2ec006fa0, first_level=true) at /home/ubuntu/mysql-server/sql/sql_parse.cc:4478
#11 0x000056107e377a01 in mysql_parse (thd=0x7fa2ec006fa0, parser_state=0x7fa31999ecb0) at /home/ubuntu/mysql-server/sql/sql_parse.cc:5288
#12 0x000056107e36c828 in dispatch_command (thd=0x7fa2ec006fa0, com_data=0x7fa31999fca0, command=COM_QUERY) at /home/ubuntu/mysql-server/sql/sql_parse.cc:1777
#13 0x000056107e36acdc in do_command (thd=0x7fa2ec006fa0) at /home/ubuntu/mysql-server/sql/sql_parse.cc:1275
#14 0x000056107e546b46 in handle_connection (arg=0x56108401bd10) at /home/ubuntu/mysql-server/sql/conn_handler/connection_handler_per_thread.cc:302
#15 0x000056108017befd in pfs_spawn_thread (arg=0x56108421e280) at /home/ubuntu/mysql-server/storage/perfschema/pfs.cc:2854
#16 0x00007fa32fc626db in start_thread (arg=0x7fa3199a0700) at pthread_create.c:463
#17 0x00007fa32e2d788f in clone () at ../sysdeps/unix/sysv/linux/x86_64/clone.S:95
```
