#ifndef __AST_UNION__
#define __AST_UNION__
typedef union ast_union
{
    ast_node * node_p;
    char *  ast_string;
    double * number;
} ast_union;

typedef struct ast_node
{
     int ast_type;
     float val;
     int son_length;
     ast_node * son[0];
} ast_node;


#define NODE_VAL 1
#define NODE_OP 2
#define NODE_STMT 3

#define ADD 1
#define SUB 2
#define MULTIPLE 3
#define DIVIDE 4
#define BRACES 5

extern ast_node * create_ast(int ast_type, int son_num ,...);
extern ast_node * create_leaf_node(int ast_type, double val);
#endif