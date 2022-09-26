grammar Abacus;



// Rules
root
    : declaration EOF
    | boolExpression EOF
    | tuple  EOF
    | LAMBDA_VARIABLE EOF
    ;

declaration
    : variablesTuple EQ tuple       # VariableDeclaration
    | LAMBDA_VARIABLE EQ lambda     # LambdaDeclaration
    ;

boolExpression
    : expression EQ EQ expression                   # EqualComparison
    | expression LS expression                      # LessComparison
    | expression GR expression                      # GreaterComparison
    | expression ((LS EQ) | (EQ LS)) expression     # LessOrEqualComparison
    | expression ((GR EQ) | (EQ GR)) expression     # GreaterOrEqualComparison
    | boolExpression op=(AND|OR|XOR) boolExpression # AndOrXor
    | NOT boolExpression                            # Not
    | LPAREN boolExpression RPAREN                  # ParenthesesBoolean
    | boolAtom                                      # BooleanAtom
    ;

AND: '&&' ;
OR: '||' ;
XOR: 'xor' ;
NOT: '¬' | '~' | '!' ;

boolAtom
    : 'true' | 'false';

lambda
    : lambdaArguments ARROW tuple           # VariablesLambda
    | LPAREN RPAREN ARROW tuple             # NullArityLambda
    ;

expression
   : sign expression                                                 # SignedExpr
   | expression PER                                                  # Percent
   | expression POW expression                                       # Pow
   | expression PER PER expression                                   # Mod
   | expression op=(MUL|DIV) expression                              # MulDiv
   | expression op=(ADD|SUB) expression                              # AddSub
   | LPAREN expression RPAREN                                        # Parentheses
   | LAMBDA_VARIABLE LPAREN mixedTuple? RPAREN recursionParameters?  # LambdaExpr
   | atom                                                            # AtomExpr
   ;


parameter
    : VARIABLE (':'|EQ) (boolExpression | expression);

recursionParameters
    : LSQPAREN (parameter (',' parameter)*)? RSQPAREN;

EQ: '=';
LS: '<';
GR: '>';

ARROW: '->' | '=>';

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

mixedTuple
    : (expression|LAMBDA_VARIABLE) (',' mixedTuple)?;

tuple
    : expression (',' tuple)?;

lambdaArguments
    : (VARIABLE | LAMBDA_VARIABLE) (',' lambdaArguments)?
    | LPAREN (VARIABLE | LAMBDA_VARIABLE) (',' lambdaArguments)? RPAREN;

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
    : VARIABLE LPAREN tuple RPAREN  # FunctionInvocation
    ;


CONSTANT
    : 'pi' | 'e' | 'phi';

SCIENTIFIC_NUMBER
   : NUMBER (('e' | 'E') SIGN? NUMBER)?
   ;


fragment SIGN:
'+'|'-';

fragment NUMBER
   : DIGITS+ (POINT DIGITS +)?
   ;


VARIABLE
   :  VALID_ID_START VALID_ID_CHAR*
   ;


LAMBDA_VARIABLE
   :  UPPERCASE VALID_ID_CHAR*
   ;


fragment VALID_ID_START
   : LOWERCASE | '_'
   ;


fragment VALID_ID_CHAR
   : VALID_ID_START | UPPERCASE | DIGITS
   ;


DIGITS: ('0' .. '9');
UPPERCASE:  ('A' .. 'Z');
LOWERCASE: ('a' .. 'z');

WHITESPACE: [ \r\n\t]+ -> skip;
