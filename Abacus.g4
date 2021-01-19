
grammar Abacus;

// Tokens


// Rules
start : expression EOF;

expression
   : expression POW expression          # Pow
   | expression op=(MUL|DIV) expression # MulDiv
   | expression op=(ADD|SUB) expression # AddSub
   | LPAREN expression RPAREN           # Parentheses
   | atom                               # AtomExpr
   ;


POW: '^' | '**';
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';

POINT
   : '.'
   ;

LPAREN: '(';
RPAREN: ')';

atom
    : SCIENTIFIC_NUMBER # Number
    | VARIABLE          # Variable
    ;


SCIENTIFIC_NUMBER
   : NUMBER
   ;


fragment NUMBER
   : ('0' .. '9')+ (POINT ('0' .. '9') +)?
   ;

VARIABLE
   : VALID_ID_START VALID_ID_CHAR*
   ;


fragment VALID_ID_START
   : ('a' .. 'z') | ('A' .. 'Z') | '_'
   ;


fragment VALID_ID_CHAR
   : VALID_ID_START | ('0' .. '9')
   ;

WHITESPACE: [ \r\n\t]+ -> skip;
