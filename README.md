# abacus

Abacus is a simple interactive calculator CLI with support for variables, comparison checks, and math functions

```
λ ./abacus -h         
abacus - a simple interactive calculator CLI with support for variables, 
comparison checks, and math functions

v1.0.0

Usage: abacus [--no-color] [--precision PRECISION] [--eval EVAL]

Options:
  --no-color, -n         disable color in output [default: false]
  --precision PRECISION, -p PRECISION
                         precision for calculations [default: 64]
  --eval EVAL, -e EVAL   evaluate expression and quit
  --help, -h             display this help and exit
  --version              display version and exit

```

## Install

```
git clone git@github.com:viktordanov/abacus.git
git checkout feature/tuples-and-lambda-expressions
go install
```

## Exprerimental λ expression support
Lambda names must begin with a capital letter A-Z and variable names — with a lowercase a-z.
### Defining lambdas 
```
<lambda name> = <lambda variables> -> <expression>, <expression> ...
```
Note:
- Parentheses around the lambda variables are optional, except when no variables are to be provided, e.g. `F = () -> 5+5`
- Lambdas can return multiple values - a tuple.
- Both `->` and `=>` can be used between the lambda variables and the expressions tuple.

#### Examples
```
Identity = x -> x
Area = a, b => a*b
Hypothenuse = (a,b) -> sqrt(a**2+b**2)
ToRad = deg -> deg * pi / 180
```

### Calling lambdas
```
<lambda name>(<parameters>)
```
#### Examples
```
> Identity = x -> x
> Identity(2)
2
> Identity()
expected 1 parameter
> UndefinedLambda()
0
```
Note:
 - undefined lambdas return 0 by default like undefined variables 

### Variables and recursion
Lambdas can use global variables and constants and will default to global variables if 
a variable in the lambda expression isn't in the lambda variables tuple.
```
> Add = x, y -> (x + y) * not_found
> Add(5, 6)
0
> not_found = 1
> Add(5, 6)
11
```

Lambdas can be recursive but only if explicitly told when called.
```
> Factorial = x -> x * Factorial(x - 1)
> Factorial(10)
recursion is disabled
```

To specify recursion parameters use `[]` after the lambda call.

`Factorial(10)[max_recurrences, last_value, stop_condition]`
- `max-recurrences` specifies the maximum number of times the lambda can call itself during the evaluation of the current expression [Default: 0];
- `last_value` specifies the value which the last lambda automatically returns when **max_recurrences** is reached or when **stop_condition** is true [Default: 0];
- `stop_condition` is a boolean expression which can use the lambda's variables; if true, the lambda returns **last_value** and stops recurring 

```
> Factorial(10)[10]
0
> Factorial(10)[10, 1]
3628800 
> Factorial(10)[10, 1, x == 5]
151200                                    
```
#### What exactly is happening?
If we define a new lambda `Count = x -> x,Count(x-1)` which returns a tuple we can observe how the value of x changes.

```
> Count(10)[10]
(10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0)

> Count(10)[10, 100]
(10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 100)

> Count(10)[10, 1, x == 5]
(10, 9, 8, 7, 6, 5, 1)
```

#### Advanced example with `until`, `from`, `reverse`, and `nth`

```
> Fib = x -> Fib(x-1) + Fib(x-2)
> Map_ = value -> Fn(value), Map_(value-1)
> Map = value,length -> until(Map_(value)[length,0], length)
> Fn = x -> Fib(x)[1e6,1,x<3]/2
> Map(10,10)
(55, 34, 21, 13, 8, 5, 3, 2, 1, 1)

> reverse(Map(10,10))
(1, 1, 2, 3, 5, 8, 13, 21, 34, 55)

> from(reverse(Map(10,10)),1)
(1, 2, 3, 5, 8, 13, 21, 34, 55)

> from(reverse(Map(10,10)),5)
(8, 13, 21, 34, 55)

> nth(from(reverse(Map(10,10)),5),2)
21
```

## Features
- `History of expressions` and `Tab completion` of all math functions and defined variables
- All common operations
  ```
  > 1+1
  2
  > 1-20
  -19
  > 5^0 / 20
  0.05
  > 2**(2+5)
  128
  ```
- Variables
   ``` 
   > d = 12.5
   12.5
   > d * 5 + 5
   67.5
   > a * 5 + 5
   5
   ```
  **Note:** Undefined variables are equal to 0
- Comparisons `<, ==, >, <=, >=`
  ```
  > pi > phi
  true
  > 10 <=10
  true
  > 2 == 0
  false
  ```
- E, Pi, Phi
   ``` 
   > e
   2.7182818284590450907955982984276
   > pi
   3.1415926535897931159979634685442
   > phi
   1.6180339887498949025257388711907
   ```
- Single arity functions:
    - sqrt
    - ln
    - floor
    - ceil
    - exp
    - sin
    - cos
    - tan
    - round
- Two arity functions:
    - round (number, digits of precision)
   ```
   > round(1.123456789,4)
  1.123
   ```
    - log (number, base)
   ```
   > log(16,4)
  2
   ```
    - min, max
   ```
   > d = 10
   10 
   > min(d,4)
   4
   > max(d,4)
   10
   ```

## Reserved names
 
 * `quit` – If a query includes quit, the program will terminate and the query will not be saved to the history file
 * The constants `e`, `pi`, and `phi`

# TODO

- [x] Add full feature list
- [ ] Write tests
  - [x] Base tests
  - [x] Simple benchmark of a complicated expression
  - [ ] Improve tests
- [ ] Refactor to depend less on other packages
- [ ] Implement most single arity functions with `*big.Float` for improved precision
