#include "ast_union.h"
#include <stdlib.h>

//
// 创建ast 节点
ast_node * create_ast(int ast_type, int son_num ,...){
        /* 为 num 个参数初始化 valist */
    ast_node * node = malloc(sizeof(ast_node) + son_num * sizeof(ast_node *) );
    node->son_length = son_num;
    node->ast_type = ast_type;
    node->son_length = son_num;
    node->val = NULL;
    
    va_start(valist, son_num);
 
    /* 访问所有赋给 valist 的参数 */
    for (i = 0; i < son_num; i++)
    {
       ast_node * temp_node = va_arg(valist, ast_node *);
       node->son[i] = temp_node;
    }
    /* 清理为 valist 保留的内存 */
    va_end(valist);

    return node;

}


//
// 创建ast 节点
ast_node * create_leaf_node(int ast_type, double val){
        /* 为 num 个参数初始化 valist */
    ast_node * node = malloc(sizeof(ast_node));
    node->son_length =0;
    node->ast_type = ast_type;
    node->son_length = 0;
    node->val = val;
    return node;

}
