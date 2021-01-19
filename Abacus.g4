
grammar Abacus;

// Tokens


// Rules
start : expression EOF;

expression
   : expression POW expression          # Pow
   | expression op=(MUL|DIV) expression # MulDiv
   | expression op=(ADD|SUB) expression # AddSub
   | LPAREN expression RPAREN           # Parentheses
   | NUMBER                             # Number
   ;


POW: '^';


MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
LPAREN: '(';
RPAREN: ')';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
