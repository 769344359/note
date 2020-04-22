```
Thread 39 "mysqld" hit Breakpoint 1, Item_func_ifnull::Item_func_ifnull (this=0x7fa2f40dd278, pos=..., a=0x7fa2f40dcc70, b=0x7fa2f40dce28) at /home/ubuntu/mysql-server/sql/item_cmpfunc.h:1301
1301	      : Item_func_coalesce(pos, a, b) {}
(gdb) bt
#0  Item_func_ifnull::Item_func_ifnull (this=0x7fa2f40dd278, pos=..., a=0x7fa2f40dcc70, b=0x7fa2f40dce28) at /home/ubuntu/mysql-server/sql/item_cmpfunc.h:1301
#1  0x000056107e7e5ddb in (anonymous namespace)::Instantiator<Item_func_ifnull, 2, 2>::instantiate (
    this=0x561082d94418 <(anonymous namespace)::Function_factory<(anonymous namespace)::Instantiator<Item_func_ifnull, 2u, 2u> >::s_singleton+8>, thd=0x7fa2f40010c0, args=0x7fa2f40dcd30)
    at /home/ubuntu/mysql-server/sql/item_create.cc:264
#2  0x000056107e7d100c in (anonymous namespace)::Function_factory<(anonymous namespace)::Instantiator<Item_func_ifnull, 2, 2> >::create_func (
    this=0x561082d94410 <(anonymous namespace)::Function_factory<(anonymous namespace)::Instantiator<Item_func_ifnull, 2u, 2u> >::s_singleton>, thd=0x7fa2f40010c0, function_name=..., 
    item_list=0x7fa2f40dcd30) at /home/ubuntu/mysql-server/sql/item_create.cc:980
#3  0x000056107e95bfd1 in PTI_function_call_generic_ident_sys::itemize (this=0x7fa2f40dcef8, pc=0x7fa31999e990, res=0x7fa2f40dd048) at /home/ubuntu/mysql-server/sql/parse_tree_items.cc:268
#4  0x000056107e95c6dc in PTI_expr_with_alias::itemize (this=0x7fa2f40dcfb0, pc=0x7fa31999e990, res=0x7fa31999e8c8) at /home/ubuntu/mysql-server/sql/parse_tree_items.cc:345
#5  0x000056107e5a5b74 in PT_item_list::contextualize (this=0x7fa2f40dd070, pc=0x7fa31999e990) at /home/ubuntu/mysql-server/sql/parse_tree_helpers.h:107
#6  0x000056107e969af9 in PT_select_item_list::contextualize (this=0x7fa2f40dd070, pc=0x7fa31999e990) at /home/ubuntu/mysql-server/sql/parse_tree_nodes.cc:2675
#7  0x000056107e962ad4 in PT_query_specification::contextualize (this=0x7fa2f40dd1b8, pc=0x7fa31999e990) at /home/ubuntu/mysql-server/sql/parse_tree_nodes.cc:1007
#8  0x000056107e96b876 in PT_query_expression::contextualize (this=0x7fa2f40dd228, pc=0x7fa31999e990) at /home/ubuntu/mysql-server/sql/parse_tree_nodes.cc:3054
#9  0x000056107e960ebf in PT_select_stmt::make_cmd (this=0x7fa2f40dd258, thd=0x7fa2f40010c0) at /home/ubuntu/mysql-server/sql/parse_tree_nodes.cc:561
#10 0x000056107e338233 in LEX::make_sql_cmd (this=0x7fa2f4007150, parse_tree=0x7fa2f40dd258) at /home/ubuntu/mysql-server/sql/sql_lex.cc:4753
#11 0x000056107e2acc48 in THD::sql_parser (this=0x7fa2f40010c0) at /home/ubuntu/mysql-server/sql/sql_class.cc:2733
#12 0x000056107e37c666 in parse_sql (thd=0x7fa2f40010c0, parser_state=0x7fa31999ecb0, creation_ctx=0x0) at /home/ubuntu/mysql-server/sql/sql_parse.cc:7104
#13 0x000056107e3774dc in mysql_parse (thd=0x7fa2f40010c0, parser_state=0x7fa31999ecb0) at /home/ubuntu/mysql-server/sql/sql_parse.cc:5192
#14 0x000056107e36c828 in dispatch_command (thd=0x7fa2f40010c0, com_data=0x7fa31999fca0, command=COM_QUERY) at /home/ubuntu/mysql-server/sql/sql_parse.cc:1777
#15 0x000056107e36acdc in do_command (thd=0x7fa2f40010c0) at /home/ubuntu/mysql-server/sql/sql_parse.cc:1275
#16 0x000056107e546b46 in handle_connection (arg=0x561084063010) at /home/ubuntu/mysql-server/sql/conn_handler/connection_handler_per_thread.cc:302
#17 0x000056108017befd in pfs_spawn_thread (arg=0x56108421e280) at /home/ubuntu/mysql-server/storage/perfschema/pfs.cc:2854
#18 0x00007fa32fc626db in start_thread (arg=0x7fa3199a0700) at pthread_create.c:463
#19 0x00007fa32e2d788f in clone () at ../sysdeps/unix/sysv/linux/x86_64/clone.S:95
```
