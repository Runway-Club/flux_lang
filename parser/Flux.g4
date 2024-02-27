parser grammar Flux;
import Variables, Math;

options {
    tokenVocab = Primitives ;
}

program: statement* EOF;

statement
    :   expression
    |   loop_statement
    |   if_statement
    |   func_declaration
    |   return_statement
    | NEWLINE
    ;

expression
    :   function_call
    |   var_assignment
    |   single_var_declaration
    |   array_var_declaration
    |   get_var
    ;

block
    :   L_BLOCK statement* R_BLOCK;

loop_statement: AT LOOP comparative_expression block;
if_statement: AT IF comparative_expression block (AT ELSE block)?;

return_statement: AT RETURN expression;

data_type
    : TEXT_TYPE
    | NUMBER_TYPE
    | BOOLEAN_TYPE
    | COMMON_IDENTIFIER
    ;

func_declaration: AT FUNC VAR_IDENTIFIER L_PAREN (VAR_IDENTIFIER (COMMA VAR_IDENTIFIER)*)? R_PAREN (RETURN_TYPE data_type)?  block;