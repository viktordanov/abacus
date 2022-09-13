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

func (v *Visitor) HandleFunctionInvocation(ctx *parser.FunctionInvocationContext) interface{} {
	functionName := ctx.VARIABLE().GetText()
	if _, ok := FunctionNames[FunctionName(functionName)]; !ok {
		return NewResult(nil).WithError(fmt.Sprintf("undefined function %s", functionName))
	}

	tupleRes := ctx.Tuple().Accept(v).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	v.convertTupleResult(tupleRes)

	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
	}

	switch FunctionName(functionName) {
	case sqrt:
		return v.handleSqrt(tuple)
	case cbrt:
		return v.handleCbrt(tuple)
	case ln:
		return v.handleLn(tuple)
	case log:
		return v.handleLog(tuple)
	case log2:
		return v.handleLog2(tuple)
	case log10:
		return v.handleLog10(tuple)
	case floor:
		return v.handleFloor(tuple)
	case ceil:
		return v.handleCeil(tuple)
	case exp:
		return v.handleExp(tuple)
	case sin:
		return v.handleSin(tuple)
	case cos:
		return v.handleCos(tuple)
	case tan:
		return v.handleTan(tuple)
	case sign:
		return v.handleSign(tuple)
	case abs:
		return v.handleAbs(tuple)
	case round:
		return v.handleRound(tuple)
	case min:
		return v.handleMin(tuple)
	case max:
		return v.handleMax(tuple)
	case avg:
		return v.handleAvg(tuple)
	case from:
		return v.handleFrom(tuple)
	case until:
		return v.handleUntil(tuple)
	case reverse:
		return v.handleReverse(tuple)
	case nth:
		return v.handleNth(tuple)
	default:
		panic("unhandled function")
	}
}

func (v *Visitor) handleSqrt(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := NewNumber(0)
	_, err = v.decimalCtx.Sqrt(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleCbrt(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := NewNumber(0)
	_, err = v.decimalCtx.Cbrt(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleLn(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := NewNumber(0)
	_, err = v.decimalCtx.Ln(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleLog(tuple Tuple) *Result {
	return v.handleLn(tuple)
}

func (v *Visitor) handleLog2(tuple Tuple) *Result {
	result := v.handleLn(tuple)
	if hasErrors(result) {
		return result
	}
	out := NewNumber(0)

	base := calculateLog(v.ConstantsStore, v.decimalCtx, NewNumber(2))
	_, err := v.decimalCtx.Quo(out.Decimal, result.Value.(Number).Decimal, base.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleLog10(tuple Tuple) *Result {
	result := v.handleLn(tuple)
	if hasErrors(result) {
		return result
	}
	out := NewNumber(0)

	base := calculateLog(v.ConstantsStore, v.decimalCtx, NewNumber(10))
	_, err := v.decimalCtx.Quo(out.Decimal, result.Value.(Number).Decimal, base.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleFloor(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := NewNumber(0)
	_, err = v.decimalCtx.Floor(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleCeil(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := NewNumber(0)
	_, err = v.decimalCtx.Ceil(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleExp(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := NewNumber(0)
	_, err = v.decimalCtx.Exp(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleSin(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	toFloat, err := value.Float64()
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := NewNumber(0)
	out.Decimal, err = out.Decimal.SetFloat64(math.Sin(toFloat))
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleCos(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	toFloat, err := value.Float64()
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := NewNumber(0)
	out.Decimal, err = out.Decimal.SetFloat64(math.Cos(toFloat))
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleTan(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	toFloat, err := value.Float64()
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := NewNumber(0)
	out.Decimal, err = out.Decimal.SetFloat64(math.Tan(toFloat))
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleSign(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}

	out := NewNumber(0)
	out = NewNumber(float64(value.Decimal.Cmp(out.Decimal)))
	return NewResult(out)
}

func (v *Visitor) handleAbs(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := NewNumber(0)
	_, err = v.decimalCtx.Abs(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleRound(tuple Tuple) *Result {
	var value Number
	var err error
	if value, err = first(tuple); err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	out := NewNumber(0)
	_, err = v.decimalCtx.Round(out.Decimal, value.Decimal)
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(out)
}

func (v *Visitor) handleMin(tuple Tuple) *Result {
	smallest := NewNumber(0)
	smallest.Set(tuple[0].Decimal)

	for i := 1; i < len(tuple); i++ {
		curr := tuple[i]
		if curr.Cmp(smallest.Decimal) == -1 {
			smallest = curr
		}
	}
	return NewResult(smallest)
}

func (v *Visitor) handleMax(tuple Tuple) *Result {
	biggest := NewNumber(0)
	biggest.Set(tuple[0].Decimal)

	for i := 1; i < len(tuple); i++ {
		curr := tuple[i]
		if curr.Cmp(biggest.Decimal) == 1 {
			biggest = curr
		}
	}
	return NewResult(biggest)
}

func (v *Visitor) handleAvg(tuple Tuple) *Result {
	sum := NewNumber(0)
	sum.Set(tuple[0].Decimal)

	for i := 1; i < len(tuple); i++ {
		curr := tuple[i]
		// TODO: handle error
		v.decimalCtx.Add(sum.Decimal, sum.Decimal, curr.Decimal)
	}
	_, err := v.decimalCtx.Quo(sum.Decimal, sum.Decimal, apd.New(int64(len(tuple)), 0))
	if err != nil {
		return NewResult(nil).WithError(err.Error())
	}
	return NewResult(sum)
}

func (v *Visitor) handleFrom(tuple Tuple) *Result {
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

func (v *Visitor) handleUntil(tuple Tuple) *Result {
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

func (v *Visitor) handleReverse(tuple Tuple) *Result {
	newTuple := make(Tuple, 0, len(tuple))

	for i := len(tuple) - 1; i >= 0; i-- {
		newTuple = append(newTuple, tuple[i])
	}

	return NewResult(newTuple)
}

func (v *Visitor) handleNth(tuple Tuple) *Result {
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

func last(tuple Tuple) (Number, error) {
	if len(tuple) == 0 {
		return Number{}, fmt.Errorf("tuple is empty")
	}
	return tuple[len(tuple)-1], nil
}
