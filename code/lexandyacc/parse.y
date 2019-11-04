/* Infix notation calculator--calc */
  
%{
#include "ast_union.h"
#define YYSTYPE ast_union
#include <math.h>
#include <stdio.h>
%}


/* BISON Declarations */
%token <node_p> NUM
%type <node_p> exp
%left '-' '+'
%left '*' '/'
%left NEG     /* negation--unary minus */
%right '^'    /* exponentiation        */

/* Grammar follows */
%%
input:    /* empty string */
             | input line
;

line:     '\n'
            | exp '\n'  { printf ("\t%.10g\n", $1); }
;

exp:      NUM                { $$ = $1; printf("the num is %lf",$1->val);          }
           | exp '+' exp        { $$ = create_ast(NODE_STMT , 3 , $1 , $3 , create_leaf_node(NODE_OP,ADD));    }
        | exp '-' exp        { $$ = create_ast(NODE_STMT , 3 , $1 , $3 , create_leaf_node(NODE_OP,SUB));     }
        | exp '*' exp        { $$ = create_ast(NODE_STMT , 3 , $1 , $3 , create_leaf_node(NODE_OP,MULTIPLE));   }
        | exp '/' exp        { $$ = create_ast(NODE_STMT , 3 , $1 , $3 , create_leaf_node(NODE_OP,DIVIDE));   }
        | '(' exp ')'        { $$  = create_ast(NODE_STMT , 2 , $2 , create_leaf_node(NODE_OP,BRACES));         }
;
%%
#include <ctype.h>
main ()
{
  yyparse ();
}
yyerror (s)  /* Called by yyparse on error */
     char *s;
{
  printf ("%s\n", s);
}

yylex ()
{
  int c;

  /* skip white space  */
  while ((c = getchar ()) == ' ' || c == '\t')
    ;
  /* process numbers   */
  if (c == '.' || isdigit (c))
    {
      ungetc (c, stdin);
      float fnum = 0;
      scanf ("%lf", &fnum);
      yylval = create_leaf_node(NODE_VAL,fnum);
      return NUM;
    }
  /* return end-of-file  */
  if (c == EOF)
    return 0;
  /* return single chars */
  return c;
}