```
Old value = 1
New value = 2
0x0000555555b9f4ba in zend_gc_addref (p=0x7fffee6593c0) at /root/php-src/Zend/zend_types.h:1035
1035		return ++(p->refcount);
(gdb) bt
#0  0x0000555555b9f4ba in zend_gc_addref (p=0x7fffee6593c0) at /root/php-src/Zend/zend_types.h:1035
#1  0x0000555555b9f5ac in zval_addref_p (pz=0x7fffee614100) at /root/php-src/Zend/zend_types.h:1070
#2  0x0000555555ba63e9 in zend_array_dup_element (source=0x7fffee659420, target=0x7fffee659480, idx=0, p=0x7fffee676640, q=0x7fffee6767c0, packed=0, static_keys=1, with_holes=0)
    at /root/php-src/Zend/zend_hash.c:1973
#3  0x0000555555ba65f3 in zend_array_dup_elements (source=0x7fffee659420, target=0x7fffee659480, static_keys=1, with_holes=0) at /root/php-src/Zend/zend_hash.c:2020
#4  0x0000555555ba6d94 in zend_array_dup (source=0x7fffee659420) at /root/php-src/Zend/zend_hash.c:2105
#5  0x0000555555b8bec6 in zend_error_va_list (type=8, error_filename=0x7fffee65ddd8 "/root/test.php", error_lineno=7, format=0x5555563f570f "Undefined variable: %s", args=0x7fffffffa7a0)
    at /root/php-src/Zend/zend.c:1344
#6  0x0000555555b8c757 in zend_error (type=8, format=0x5555563f570f "Undefined variable: %s") at /root/php-src/Zend/zend.c:1481
#7  0x0000555555bf1c07 in zval_undefined_cv (var=96) at /root/php-src/Zend/zend_execute.c:275
#8  0x0000555555bf1d53 in _get_zval_ptr_cv_BP_VAR_R (var=96) at /root/php-src/Zend/zend_execute.c:348
#9  0x0000555555c5dc78 in ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CV_HANDLER () at /root/php-src/Zend/zend_vm_execute.h:46207
#10 0x0000555555c69ebf in execute_ex (ex=0x7fffee614020) at /root/php-src/Zend/zend_vm_execute.h:57578
#11 0x0000555555c6a466 in zend_execute (op_array=0x7fffee682400, return_value=0x0) at /root/php-src/Zend/zend_vm_execute.h:57922
#12 0x0000555555b8d2b4 in zend_execute_scripts (type=8, retval=0x0, file_count=3) at /root/php-src/Zend/zend.c:1666
#13 0x0000555555aee5b4 in php_execute_script (primary_file=0x7fffffffd0e0) at /root/php-src/main/main.c:2617
#14 0x0000555555c6d070 in do_cli (argc=2, argv=0x55555697d570) at /root/php-src/sapi/cli/php_cli.c:961
#15 0x0000555555c6e232 in main (argc=2, argv=0x55555697d570) at /root/php-src/sapi/cli/php_cli.c:1356

```


```
Hardware watchpoint 1: *(uint32_t *) 0x7fffee6593c0

Old value = 3999634464
New value = 1
zend_gc_set_refcount (p=0x7fffee6593c0, rc=1) at /root/php-src/Zend/zend_types.h:1030
1030		return p->refcount;
(gdb) bt
#0  zend_gc_set_refcount (p=0x7fffee6593c0, rc=1) at /root/php-src/Zend/zend_types.h:1030
#1  0x0000555555ba043c in _zend_hash_init_int (ht=0x7fffee6593c0, nSize=8, pDestructor=0x555555b88dc8 <zval_ptr_dtor>, persistent=0 '\000') at /root/php-src/Zend/zend_hash.c:229
#2  0x0000555555ba05ac in _zend_new_array (nSize=8) at /root/php-src/Zend/zend_hash.c:257
#3  0x0000555555c5dedb in ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CV_HANDLER () at /root/php-src/Zend/zend_vm_execute.h:46283
#4  0x0000555555c69ebf in execute_ex (ex=0x7fffee614020) at /root/php-src/Zend/zend_vm_execute.h:57578
#5  0x0000555555c6a466 in zend_execute (op_array=0x7fffee682400, return_value=0x0) at /root/php-src/Zend/zend_vm_execute.h:57922
#6  0x0000555555b8d2b4 in zend_execute_scripts (type=8, retval=0x0, file_count=3) at /root/php-src/Zend/zend.c:1666
#7  0x0000555555aee5b4 in php_execute_script (primary_file=0x7fffffffd0e0) at /root/php-src/main/main.c:2617
#8  0x0000555555c6d070 in do_cli (argc=2, argv=0x55555697d570) at /root/php-src/sapi/cli/php_cli.c:961
#9  0x0000555555c6e232 in main (argc=2, argv=0x55555697d570) at /root/php-src/sapi/cli/php_cli.c:1356
(gdb) c
Continuing.
```
```
(gdb) bt
#0  __GI_raise (sig=sig@entry=6) at ../sysdeps/unix/sysv/linux/raise.c:51
#1  0x00007ffff5d45801 in __GI_abort () at abort.c:79
#2  0x00007ffff5d3539a in __assert_fail_base (fmt=0x7ffff5ebc7d8 "%s%s%s:%u: %s%sAssertion `%s' failed.\n%n", 
    assertion=assertion@entry=0x5555563ed9f0 "(zend_gc_refcount(&(ht)->gc) == 1) || ((ht)->u.flags & (1<<6))", file=file@entry=0x5555563ed960 "/root/php-src/Zend/zend_hash.c", 
    line=line@entry=965, function=function@entry=0x5555563edee0 <__PRETTY_FUNCTION__.13579> "_zend_hash_index_add_or_update_i") at assert.c:92
#3  0x00007ffff5d35412 in __GI___assert_fail (assertion=0x5555563ed9f0 "(zend_gc_refcount(&(ht)->gc) == 1) || ((ht)->u.flags & (1<<6))", file=0x5555563ed960 "/root/php-src/Zend/zend_hash.c", 
    line=965, function=0x5555563edee0 <__PRETTY_FUNCTION__.13579> "_zend_hash_index_add_or_update_i") at assert.c:101
#4  0x0000555555ba2bdb in _zend_hash_index_add_or_update_i (ht=0x7fffee6593c0, h=0, pData=0x555556978ac0 <executor_globals>, flag=18) at /root/php-src/Zend/zend_hash.c:965
#5  0x0000555555ba3170 in zend_hash_next_index_insert (ht=0x7fffee6593c0, pData=0x555556978ac0 <executor_globals>) at /root/php-src/Zend/zend_hash.c:1078
#6  0x0000555555c5dcba in ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CV_HANDLER () at /root/php-src/Zend/zend_vm_execute.h:46211
#7  0x0000555555c69ebf in execute_ex (ex=0x7fffee614020) at /root/php-src/Zend/zend_vm_execute.h:57578
#8  0x0000555555c6a466 in zend_execute (op_array=0x7fffee682400, return_value=0x0) at /root/php-src/Zend/zend_vm_execute.h:57922
#9  0x0000555555b8d2b4 in zend_execute_scripts (type=8, retval=0x0, file_count=3) at /root/php-src/Zend/zend.c:1666
#10 0x0000555555aee5b4 in php_execute_script (primary_file=0x7fffffffd0e0) at /root/php-src/main/main.c:2617
#11 0x0000555555c6d070 in do_cli (argc=2, argv=0x55555697d570) at /root/php-src/sapi/cli/php_cli.c:961
#12 0x0000555555c6e232 in main (argc=2, argv=0x55555697d570) at /root/php-src/sapi/cli/php_cli.c:1356

```


```
(gdb) p object_ptr->value->arr->gc
$5 = {refcount = 1, u = {type_info = 23}}
```
