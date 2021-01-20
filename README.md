# abacus

Abacus is a simple interactive calculator CLI with support for variables, comparison checks, and math functions

```
λ abacus -h
abacus - a simple interactive calculator CLI with support for variables, comparison 
checks, and math functions
Usage: 
  -color
        color the output (default true)
  -prec int
        precision bits to calculate for (default 32)
```

## Install

`go get github.com/viktordanov/abacus`

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

# TODO

- [x] Add full feature list
- [ ] Write tests
- [ ] Implement most single arity functions with `*big.Float` for improved precision
