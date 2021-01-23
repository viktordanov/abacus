
grammar Abacus;

// Tokens


// Rules
root
    : declaration EOF
    | comparison EOF
    | expression  EOF
    ;

declaration
    : VARIABLE EQ expression
    ;

comparison
    : expression EQ EQ expression               # EqualComparison
    | expression LS expression                  # LessComparison
    | expression GR expression                  # GreaterComparison
    | expression ((LS EQ) | (EQ LS)) expression   # LessOrEqualComparison
    | expression ((GR EQ) | (EQ GR)) expression   # GreaterOrEqualComparison
    ;


expression
   : expression POW expression          # Pow
   | expression op=(MUL|DIV) expression # MulDiv
   | expression op=(ADD|SUB) expression # AddSub
   | LPAREN expression RPAREN           # Parentheses
   | sign? atom                         # AtomExpr
   ;

EQ: '=';
LS: '<';
GR: '>';

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
    : function              # FuncExpr
    | CONSTANT              # Constant
    | SCIENTIFIC_NUMBER     # Number
    | VARIABLE              # Variable
    ;

sign
    : '+' # PlusSign
    | '-' # MinusSign;


function
    : 'sqrt' LPAREN expression RPAREN                   # SqrtFunction
    | 'ln' LPAREN expression RPAREN                     # LnFunction
    | 'log' LPAREN expression RPAREN                    # LogDefFunction
    | 'log2' LPAREN expression RPAREN                   # Log2Function
    | 'log10' LPAREN expression RPAREN                  # Log10Function
    | 'floor' LPAREN expression RPAREN                  # FloorFunction
    | 'ceil' LPAREN expression RPAREN                   # CeilFunction
    | 'exp' LPAREN expression RPAREN                    # ExpFunction
    | 'sin' LPAREN expression RPAREN                    # SinFunction
    | 'cos' LPAREN expression RPAREN                    # CosFunction
    | 'tan' LPAREN expression RPAREN                    # TanFunction
    | 'round' LPAREN expression RPAREN                  # RoundDefFunction
    | 'round' LPAREN expression ',' expression RPAREN   # Round2Function
    | 'log' LPAREN expression ',' expression RPAREN     # LogFunction
    | 'min' LPAREN expression ',' expression RPAREN     # MinFunction
    | 'max' LPAREN expression ',' expression RPAREN     # MaxFunction
    ;


CONSTANT
    : 'pi' | 'e' | 'phi';

SCIENTIFIC_NUMBER
   : NUMBER (('e' | 'E') SIGN? NUMBER)?
   ;


fragment SIGN:
'+'|'-';

fragment NUMBER
   : ('0' .. '9')+ (POINT ('0' .. '9') +)?
   ;

VARIABLE
   :  VALID_ID_START VALID_ID_CHAR*
   ;


fragment VALID_ID_START
   : ('a' .. 'z') | ('A' .. 'Z') | '_'
   ;


fragment VALID_ID_CHAR
   : VALID_ID_START | ('0' .. '9')
   ;

WHITESPACE: [ \r\n\t]+ -> skip;
