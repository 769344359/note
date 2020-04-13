```
Breakpoint 1, php_stream_filter_free (filter=0x7fffef747690) at /home/dinosaur/Downloads/php-7.2.2/main/streams/filter.c:279
279		if (filter->fops->dtor)
(gdb) bt
#0  php_stream_filter_free (filter=0x7fffef747690) at /home/dinosaur/Downloads/php-7.2.2/main/streams/filter.c:279
#1  0x00000000009b95bb in php_stream_filter_remove (filter=0x7fffef747690, call_dtor=1) at /home/dinosaur/Downloads/php-7.2.2/main/streams/filter.c:499
#2  0x00000000009af3b4 in _php_stream_free (stream=0x7fffef684c00, close_options=67) at /home/dinosaur/Downloads/php-7.2.2/main/streams/streams.c:483
#3  0x0000000000859f0f in zif_fclose (execute_data=0x7fffef61e0b0, return_value=0x7fffffffa1f0) at /home/dinosaur/Downloads/php-7.2.2/ext/standard/file.c:917
#4  0x0000000000aae986 in ZEND_DO_ICALL_SPEC_RETVAL_UNUSED_HANDLER () at /home/dinosaur/Downloads/php-7.2.2/Zend/zend_vm_execute.h:573
#5  0x0000000000b4295e in execute_ex (ex=0x7fffef61e030) at /home/dinosaur/Downloads/php-7.2.2/Zend/zend_vm_execute.h:59731
#6  0x0000000000b47d9d in zend_execute (op_array=0x7fffef684b00, return_value=0x0) at /home/dinosaur/Downloads/php-7.2.2/Zend/zend_vm_execute.h:63760
#7  0x0000000000a3afe0 in zend_execute_scripts (type=8, retval=0x0, file_count=3) at /home/dinosaur/Downloads/php-7.2.2/Zend/zend.c:1496
#8  0x000000000098c749 in php_execute_script (primary_file=0x7fffffffc8d0) at /home/dinosaur/Downloads/php-7.2.2/main/main.c:2590
#9  0x0000000000b4b2a5 in do_cli (argc=2, argv=0x1561ff0) at /home/dinosaur/Downloads/php-7.2.2/sapi/cli/php_cli.c:1011
#10 0x0000000000b4c491 in main (argc=2, argv=0x1561ff0) at /home/dinosaur/Downloads/php-7.2.2/sapi/cli/php_cli.c:1404

```

多次fclose 会coredump
