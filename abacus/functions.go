package abacus

import (
	"fmt"
	"github.com/cockroachdb/apd"
	"github.com/viktordanov/abacus/parser"
	"math"
)

type FunctionName string

const (
	sqrt    = FunctionName("sqrt")
	cbrt    = FunctionName("cbrt")
	ln      = FunctionName("ln")
	log     = FunctionName("log")
	log2    = FunctionName("log2")
	log10   = FunctionName("log10")
	floor   = FunctionName("floor")
	ceil    = FunctionName("ceil")
	exp     = FunctionName("exp")
	sin     = FunctionName("sin")
	cos     = FunctionName("cos")
	tan     = FunctionName("tan")
	sign    = FunctionName("sign")
	abs     = FunctionName("abs")
	round   = FunctionName("round")
	min     = FunctionName("min")
	max     = FunctionName("max")
	avg     = FunctionName("avg")
	from    = FunctionName("from")
	until   = FunctionName("until")
	reverse = FunctionName("reverse")
	nth     = FunctionName("nth")
)

var FunctionNames = map[FunctionName]struct{}{
	sqrt:    {},
	cbrt:    {},
	ln:      {},
	log:     {},
	log2:    {},
	log10:   {},
	floor:   {},
	ceil:    {},
	exp:     {},
	sin:     {},
	cos:     {},
	tan:     {},
	sign:    {},
	abs:     {},
	round:   {},
	min:     {},
	max:     {},
	avg:     {},
	from:    {},
	until:   {},
	reverse: {},
	nth:     {},
}

func (a *AbacusVisitor) HandleFunctionInvocation(ctx *parser.FunctionInvocationContext, decimalCtx *apd.Context) interface{} {
	functionName := ctx.VARIABLE().GetText()
	if _, ok := FunctionNames[FunctionName(functionName)]; !ok {
		return NewResult(nil).WithError(fmt.Sprintf("undefined function %s", functionName))
	}

	tupleRes := ctx.Tuple().Accept(a).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	a.convertTupleResult(tupleRes)

	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
	}

	switch FunctionName(functionName) {
	case sqrt:
		return handleSqrt(tuple, decimalCtx)
	case cbrt:
		return handleCbrt(tuple, decimalCtx)
	case ln:
		return handleLn(tuple, decimalCtx)
	case log:
		return handleLog(tuple, decimalCtx)
	case log2:
		return handleLog2(tuple, decimalCtx)
	case log10:
		return handleLog10(tuple, decimalCtx)
	case floor:
		return handleFloor(tuple, decimalCtx)
	case ceil:
		return handleCeil(tuple, decimalCtx)
	case exp:
		return handleExp(tuple, decimalCtx)
	case sin:
		return handleSin(tuple, decimalCtx)
	case cos:
		return handleCos(tuple, decimalCtx)
	case tan:
		return handleTan(tuple, decimalCtx)
	case sign:
		return handleSign(tuple, decimalCtx)
	case abs:
		return handleAbs(tuple, decimalCtx)
	case round:
		return handleRound(tuple, decimalCtx)
	case min:
		return handleMin(tuple, decimalCtx)
	case max:
		return handleMax(tuple, decimalCtx)
	case avg:
		return handleAvg(tuple, decimalCtx)
	case from:
		return handleFrom(tuple, decimalCtx)
	case until:
		return handleUntil(tuple, decimalCtx)
	case reverse:
		return handleReverse(tuple, decimalCtx)
	case nth:
		return handleNth(tuple, decimalCtx)
	default:
		panic("unhandled function")
	}
}

func handleSqrt(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := newNumber(0)
	_, err = decimalCtx.Sqrt(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func handleCbrt(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := newNumber(0)
	_, err = decimalCtx.Cbrt(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func handleLn(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := newNumber(0)
	_, err = decimalCtx.Ln(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func handleLog(tuple Tuple, decimalCtx *apd.Context) *Result {
	return handleLn(tuple, decimalCtx)
}

func handleLog2(tuple Tuple, decimalCtx *apd.Context) *Result {
	result := handleLn(tuple, decimalCtx)
	if hasErrors(result) {
		return result
	}
	out := newNumber(0)

	base := cachedLog(newNumber(2))
	_, err := decimalCtx.Quo(out.Decimal, result.Value.(Number).Decimal, base.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func handleLog10(tuple Tuple, decimalCtx *apd.Context) *Result {
	result := handleLn(tuple, decimalCtx)
	if hasErrors(result) {
		return result
	}
	out := newNumber(0)

	base := cachedLog(newNumber(10))
	_, err := decimalCtx.Quo(out.Decimal, result.Value.(Number).Decimal, base.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func handleFloor(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := newNumber(0)
	_, err = decimalCtx.Floor(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func handleCeil(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := newNumber(0)
	_, err = decimalCtx.Ceil(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func handleExp(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := newNumber(0)
	_, err = decimalCtx.Exp(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func handleSin(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	toFloat, err := value.Float64()
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	v := newNumber(0)
	v.Decimal, err = v.Decimal.SetFloat64(math.Sin(toFloat))
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(v)
}

func handleCos(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	toFloat, err := value.Float64()
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	v := newNumber(0)
	v.Decimal, err = v.Decimal.SetFloat64(math.Cos(toFloat))
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(v)
}

func handleTan(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	toFloat, err := value.Float64()
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	v := newNumber(0)
	v.Decimal, err = v.Decimal.SetFloat64(math.Tan(toFloat))
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(v)
}

func handleSign(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}

	out := newNumber(0)
	out = newNumber(float64(value.Decimal.Cmp(out.Decimal)))
	return NewResult(out)
}

func handleAbs(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := newNumber(0)
	_, err = decimalCtx.Abs(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func handleRound(tuple Tuple, decimalCtx *apd.Context) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := newNumber(0)
	_, err = decimalCtx.Round(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func handleMin(tuple Tuple, decimalCtx *apd.Context) *Result {
	smallest := newNumber(0)
	smallest.Set(tuple[0].Decimal)

	for i := 1; i < len(tuple); i++ {
		curr := tuple[i]
		if curr.Cmp(smallest.Decimal) == -1 {
			smallest = curr
		}
	}
	return NewResult(smallest)
}

func handleMax(tuple Tuple, decimalCtx *apd.Context) *Result {
	biggest := newNumber(0)
	biggest.Set(tuple[0].Decimal)

	for i := 1; i < len(tuple); i++ {
		curr := tuple[i]
		if curr.Cmp(biggest.Decimal) == 1 {
			biggest = curr
		}
	}
	return NewResult(biggest)
}

func handleAvg(tuple Tuple, decimalCtx *apd.Context) *Result {
	sum := newNumber(0)
	sum.Set(tuple[0].Decimal)

	for i := 1; i < len(tuple); i++ {
		curr := tuple[i]
		// TODO: handle error
		decimalCtx.Add(sum.Decimal, sum.Decimal, curr.Decimal)
	}
	_, err := decimalCtx.Quo(sum.Decimal, sum.Decimal, apd.New(int64(len(tuple)), 0))
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(sum)
}

func handleFrom(tuple Tuple, decimalCtx *apd.Context) *Result {
	var arg Number
	var err error

	if arg, err = last(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}

	if len(tuple) == 1 {
		return NewResult(nil).WithError("from requires at least 2 arguments")
	}

	intValue, _ := arg.Int64()
	if intValue < 0 {
		return NewResult(nil).WithError("from requires a non-negative integer")
	}

	newTuple := make(Tuple, 0, len(tuple)-1)
	length := len(tuple)
	if intValue > int64(length)-1 {
		intValue = int64(length) - 1
	}

	for i := intValue; i < int64(length)-1; i++ {
		newTuple = append(newTuple, tuple[i])
	}

	return NewResult(newTuple)
}

func handleUntil(tuple Tuple, decimalCtx *apd.Context) *Result {
	var arg Number
	var err error

	if arg, err = last(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}

	if len(tuple) == 1 {
		return NewResult(nil).WithError("until requires at least 2 arguments")
	}

	intValue, _ := arg.Int64()
	if intValue < 0 {
		return NewResult(nil).WithError("until requires a non-negative integer")
	}

	newTuple := make(Tuple, 0, len(tuple)-1)
	length := len(tuple)
	if intValue > int64(length)-1 {
		intValue = int64(length) - 1
	}
	for i := 0; i < int(intValue); i++ {
		newTuple = append(newTuple, tuple[i])
	}

	return NewResult(newTuple)
}

func handleReverse(tuple Tuple, decimalCtx *apd.Context) *Result {
	newTuple := make(Tuple, 0, len(tuple))

	for i := len(tuple) - 1; i >= 0; i-- {
		newTuple = append(newTuple, tuple[i])
	}

	return NewResult(newTuple)
}

func handleNth(tuple Tuple, decimalCtx *apd.Context) *Result {
	var arg Number
	var err error

	if arg, err = last(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}

	if len(tuple) == 1 {
		return NewResult(nil).WithError("nth requires at least 2 arguments")
	}

	intValue, _ := arg.Int64()
	if intValue < 0 || intValue > int64(len(tuple)-2) {
		return NewResult(nil).WithError("nth requires a non-negative integer less than the length of the tuple")
	}

	return NewResult(tuple[intValue])
}

func first(tuple Tuple) (Number, error) {
	if len(tuple) == 0 {
		return Number{}, fmt.Errorf("tuple is empty")
	}
	return tuple[0], nil
}

func second(tuple Tuple) (Number, error) {
	if len(tuple) < 2 {
		return Number{}, fmt.Errorf("tuple has less than 2 elements")
	}
	return tuple[1], nil
}

func last(tuple Tuple) (Number, error) {
	if len(tuple) == 0 {
		return Number{}, fmt.Errorf("tuple is empty")
	}
	return tuple[len(tuple)-1], nil
}
