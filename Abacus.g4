
grammar Abacus;

// Tokens


// Rules
root
    : declaration EOF
    | expression  EOF
    ;

declaration
    : VARIABLE EQ expression
    ;


expression
   : expression POW expression          # Pow
   | expression op=(MUL|DIV) expression # MulDiv
   | expression op=(ADD|SUB) expression # AddSub
   | LPAREN expression RPAREN           # Parentheses
   | atom                               # AtomExpr
   ;

EQ: '=';

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
    : function          # FuncExpr
    | CONSTANT          # Constant
    | SCIENTIFIC_NUMBER # Number
    | VARIABLE          # Variable
    ;


function
    : 'sqrt' LPAREN expression RPAREN   # SqrtFunction
    | 'ln' LPAREN expression RPAREN     # LnFunction
    | 'floor' LPAREN expression RPAREN  # FloorFunction
    | 'ceil' LPAREN expression RPAREN   # CeilFunction
    | 'exp' LPAREN expression RPAREN    # ExpFunction
    | 'sin' LPAREN expression RPAREN    # SinFunction
    | 'cos' LPAREN expression RPAREN    # CosFunction
    | 'tan' LPAREN expression RPAREN    # TanFunction
    ;


CONSTANT
    : 'pi' | 'e' | 'phi';

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
