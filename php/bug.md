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

```

PHPAPI php_stream_filter *php_stream_filter_remove(php_stream_filter *filter, int call_dtor)
{
	if (filter->prev) {
		filter->prev->next = filter->next;
	} else {
		filter->chain->head = filter->next;
	}
	if (filter->next) {
		filter->next->prev = filter->prev;
	} else {
		filter->chain->tail = filter->prev;
	}

	if (filter->res) {
		zend_list_delete(filter->res);
	}

	if (call_dtor) {
		php_stream_filter_free(filter);
		return NULL;
	}
	return filter;
}

```

```
(gdb) bt
#0  strfilter_rot13_create (filtername=0x7fffe5fb1d08 "string.rot13", filterparams=0x0, persistent=0 '\000') at /home/dinosaur/Downloads/php-7.2.2/ext/standard/filters.c:72
#1  0x00000000009b8735 in php_stream_filter_create (filtername=0x7fffe5fb1d08 "string.rot13", filterparams=0x0, persistent=0 '\000') at /home/dinosaur/Downloads/php-7.2.2/main/streams/filter.c:231
#2  0x0000000000900c48 in apply_filter_to_stream (append=1, execute_data=0x7fffef61e0b0, return_value=0x7fffef61e0a0) at /home/dinosaur/Downloads/php-7.2.2/ext/standard/streamsfuncs.c:1216
#3  0x0000000000900df8 in zif_stream_filter_append (execute_data=0x7fffef61e0b0, return_value=0x7fffef61e0a0) at /home/dinosaur/Downloads/php-7.2.2/ext/standard/streamsfuncs.c:1254
#4  0x0000000000aaeb97 in ZEND_DO_ICALL_SPEC_RETVAL_USED_HANDLER () at /home/dinosaur/Downloads/php-7.2.2/Zend/zend_vm_execute.h:617
#5  0x0000000000b4296b in execute_ex (ex=0x7fffef61e030) at /home/dinosaur/Downloads/php-7.2.2/Zend/zend_vm_execute.h:59734
#6  0x0000000000b47d9d in zend_execute (op_array=0x7fffef684b00, return_value=0x0) at /home/dinosaur/Downloads/php-7.2.2/Zend/zend_vm_execute.h:63760
#7  0x0000000000a3afe0 in zend_execute_scripts (type=8, retval=0x0, file_count=3) at /home/dinosaur/Downloads/php-7.2.2/Zend/zend.c:1496
#8  0x000000000098c749 in php_execute_script (primary_file=0x7fffffffc8d0) at /home/dinosaur/Downloads/php-7.2.2/main/main.c:2590
#9  0x0000000000b4b2a5 in do_cli (argc=2, argv=0x1561ff0) at /home/dinosaur/Downloads/php-7.2.2/sapi/cli/php_cli.c:1011
#10 0x0000000000b4c491 in main (argc=2, argv=0x1561ff0) at /home/dinosaur/Downloads/php-7.2.2/sapi/cli/php_cli.c:1404

```
