parser grammar Math;

options {
    tokenVocab = Primitives;
}

op_level1: OP_MUL | OP_DIV | OP_MOD | OP_POW;
op_level2: OP_ADD | OP_SUB;
op_level3: OP_AND | OP_OR | OP_XOR;
op_level4: OP_EQ | OP_NEQ | OP_LT | OP_GT | OP_GTE | OP_LTE;
op_level5: OP_NOT;

numeric_expression
    :   L_PAREN numeric_expression R_PAREN
    |   numeric_expression op_level1 numeric_expression
    |   numeric_expression op_level2 numeric_expression
    |   NUMBER
    |   get_var
    |   function_call
    |   OP_SUB numeric_expression
    ;

logical_expression
    :   L_PAREN logical_expression R_PAREN
    |   logical_expression op_level3 logical_expression
    |   logical_expression op_level4 logical_expression
    |   OP_NOT logical_expression
    |   BOOLEAN
    |   get_var
    |   function_call
    ;

comparative_expression
    :   numeric_expression op_level4 numeric_expression
    |   logical_expression op_level3 logical_expression
    |   comparative_expression op_level3 logical_expression
    |   op_level5 L_PAREN comparative_expression R_PAREN
    ;

math_expression
    :   get_var
    |   numeric_expression
    |   logical_expression
    ;

get_array_element
    :   VAR_IDENTIFIER L_SQUARE numeric_expression R_SQUARE
    ;

get_child
    :   VAR_IDENTIFIER DOT VAR_IDENTIFIER
    |   VAR_IDENTIFIER DOT get_array_element
    |   VAR_IDENTIFIER DOT get_child
    |   get_array_element DOT get_child
    |   get_array_element
    |   VAR_IDENTIFIER
    ;

get_var
    :  VAR_IDENTIFIER
    |  get_array_element
    |  get_child
    ;

function_call
    :   VAR_IDENTIFIER L_PAREN (math_expression (COMMA math_expression)*)? R_PAREN
    ;