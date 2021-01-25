grammar Abacus;



// Rules
root
    : declaration EOF
    | comparison EOF
    | tuple  EOF
    ;

declaration
    : variablesTuple EQ tuple       # VariableDeclaration
    | LAMBDA_VARIABLE EQ lambda     # LambdaDeclaration
    ;

comparison
    : expression EQ EQ expression               # EqualComparison
    | expression LS expression                  # LessComparison
    | expression GR expression                  # GreaterComparison
    | expression ((LS EQ) | (EQ LS)) expression   # LessOrEqualComparison
    | expression ((GR EQ) | (EQ GR)) expression   # GreaterOrEqualComparison
    ;


lambda
    : variablesTuple ARROW tuple            # VariablesLambda
    | LPAREN RPAREN ARROW tuple             # NullArityLambda
    ;

expression
   : sign expression                                            # SignedExpr
   | expression PER                                             # Percent
   | expression POW expression                                  # Pow
   | expression PER PER expression                              # Mod
   | expression op=(MUL|DIV) expression                         # MulDiv
   | expression op=(ADD|SUB) expression                         # AddSub
   | LPAREN expression RPAREN                                   # Parentheses
   | LAMBDA_VARIABLE LPAREN tuple? RPAREN recursionParameters?  # LambdaExpr
   | atom                                                       # AtomExpr
   ;


recursionParameters
    : LSQPAREN expression (',' expression (',' comparison)?)? RSQPAREN;

EQ: '=';
LS: '<';
GR: '>';

ARROW: '->';

POW: '^' | '**';
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
PER: '%';

POINT
   : '.'
   ;

LPAREN: '(';
RPAREN: ')';
LSQPAREN: '[';
RSQPAREN: ']';

tuple
    : expression (',' tuple)?;

variablesTuple
    : VARIABLE (',' variablesTuple)?
    | LPAREN VARIABLE (',' variablesTuple)? RPAREN;


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
    | 'cbrt' LPAREN expression RPAREN                   # CbrtFunction
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
    | 'abs' LPAREN expression RPAREN                    # AbsFunction
    | 'round' LPAREN expression ',' expression RPAREN   # Round2Function
    | 'log' LPAREN expression ',' expression RPAREN     # LogFunction
    | 'min' LPAREN tuple RPAREN                         # MinFunction
    | 'max' LPAREN tuple RPAREN                         # MaxFunction
    | 'avg' LPAREN tuple RPAREN                         # AvgFunction
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

LAMBDA_VARIABLE
   :  ('A' .. 'Z') VALID_ID_CHAR*
   ;


fragment VALID_ID_START
   : ('a' .. 'z') | '_'
   ;


fragment VALID_ID_CHAR
   : VALID_ID_START | ('0' .. '9')
   ;

WHITESPACE: [ \r\n\t]+ -> skip;
