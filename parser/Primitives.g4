lexer grammar Primitives;


TEXT
    : '"' *? '"'
    | '\''.*?'\'' ;
TEXT_TYPE: 'text';

BOOLEAN : 'true' | 'false' ;
BOOLEAN_TYPE: 'bool';

NUMBER  : '-'? [0-9]+ ('.' [0-9]*)? ;
NUMBER_TYPE: 'num';

NULL : 'na' ;

DIGIT: [0-9];
OCTET
    : '0'
    | ([1-9] DIGIT?)  // 1 to 99
    | '1' DIGIT DIGIT // 100 to 199
    | '2' [0-4] DIGIT // 200 to 249
    | '25' [0-5]      // 250 to 255
    ;
IPV4: OCTET '.' OCTET '.' OCTET '.' OCTET;
IPV4_TYPE: 'ipv4';

LOOP: 'loop';
IF: 'if';
ELSE: 'else';
FUNC: 'fun';
RETURN: 'return';
RETURN_TYPE: '->';




OP_ADD: '+';
OP_SUB: '-';
OP_MUL: '*';
OP_DIV: '/';
OP_MOD: '%';
OP_POW: '^';
OP_EQ: '=';
OP_NEQ: '!=';
OP_GT: '>';
OP_LT: '<';
OP_GTE: '>=';
OP_LTE: '<=';
OP_AND: 'and';
OP_OR: 'or';
OP_NOT: 'not';
OP_XOR: 'xor';

L_BLOCK : '{' ;
R_BLOCK : '}' ;
L_PAREN : '(' ;
R_PAREN : ')' ;
L_SQUARE : '[' ;
R_SQUARE : ']' ;
COMMA : ',' ;
DOT : '.' ;

VAR_IDENTIFIER : [a-z][a-zA-Z0-9]* ;
COMMON_IDENTIFIER : [A-Z][a-zA-Z0-9]* ;

AT : '@' ;
NEWLINE : '\r'? '\n' ;
WS      : [ \t\r]+ -> skip ;