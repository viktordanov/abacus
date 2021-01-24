package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/viktordanov/abacus/parser"
	"math"
	"math/big"
	"strconv"
)

type Lambda struct {
	ctx *parser.LambdaDeclarationContext
}

type AbacusVisitor struct {
	antlr.ParseTreeVisitor
	vars            map[string]*big.Float
	lambdas         map[string]*Lambda
	lambdaVars      map[string]*big.Float
	lambdaRecursion map[string]uint
	answerChannel   chan interface{}
}

func NewAbacusVisitor() *AbacusVisitor {
	return &AbacusVisitor{
		ParseTreeVisitor: &parser.BaseAbacusVisitor{},
		vars:             make(map[string]*big.Float),
		lambdas:          make(map[string]*Lambda),
		lambdaVars:       make(map[string]*big.Float),
		lambdaRecursion:  make(map[string]uint),
		answerChannel:    make(chan interface{}),
	}
}

func (a *AbacusVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.RootContext:
		a.lambdaRecursion = make(map[string]uint)
		return val.Accept(a)

	case *parser.DeclarationContext:
		return val.Accept(a)

	case *parser.ComparisonContext:
		return val.Accept(a)

	case *parser.ExpressionContext:
		return val.Accept(a)
	}
	return nil
}

func (a *AbacusVisitor) VisitRoot(c *parser.RootContext) interface{} {
	if c.Tuple() != nil {
		return c.Tuple().Accept(a)
	}
	if c.Declaration() != nil {
		return c.Declaration().Accept(a)
	}
	if c.Comparison() != nil {
		return c.Comparison().Accept(a)
	}
	return nil
}

func (a *AbacusVisitor) visitTupleTail(c parser.ITupleContext, resultTuple *ResultTuple) {
	ctx, ok := c.(*parser.TupleContext)
	if !ok || ctx == nil {
		return
	}
	val, _ := ctx.Expression().Accept(a).(*big.Float)
	resultTuple.Values = append(resultTuple.Values, val)
	a.visitTupleTail(ctx.Tuple(), resultTuple)
}

func (a *AbacusVisitor) VisitTuple(c *parser.TupleContext) interface{} {
	if c.Tuple() == nil {
		return c.Expression().Accept(a)
	}

	evaledTuple := NewResultTuple()
	val, _ := c.Expression().Accept(a).(*big.Float)
	evaledTuple.Values = append(evaledTuple.Values, val)
	a.visitTupleTail(c.Tuple(), &evaledTuple)

	return evaledTuple
}

func (a *AbacusVisitor) visitVariableTupleTail(c parser.IVariablesTupleContext, resultTuple *ResultVariablesTuple) {
	ctx, ok := c.(*parser.VariablesTupleContext)
	if !ok || ctx == nil {
		return
	}
	val := ctx.VARIABLE().GetText()
	resultTuple.Variables = append(resultTuple.Variables, val)
	a.visitVariableTupleTail(ctx.VariablesTuple(), resultTuple)
}

func (a *AbacusVisitor) VisitVariablesTuple(c *parser.VariablesTupleContext) interface{} {
	if c.VariablesTuple() == nil {
		return c.VARIABLE().GetText()
	}

	evaledTuple := NewResultVariablesTuple()

	val := c.VARIABLE().GetText()
	evaledTuple.Variables = append(evaledTuple.Variables, val)
	a.visitVariableTupleTail(c.VariablesTuple(), &evaledTuple)

	foundVars := make(map[string]bool)
	for _, variable := range evaledTuple.Variables {
		if _, ok := foundVars[variable]; !ok {
			foundVars[variable] = true
		} else {
			return ResultError("duplicate variable name \"" + variable + "\"")
		}
	}

	return evaledTuple
}

func (a *AbacusVisitor) convertVariablesTupleResult(result interface{}) (ResultVariablesTuple, *ResultError) {
	variableNames := NewResultVariablesTuple()
	switch val := result.(type) {
	case ResultError:
		a.answerChannel <- val
		return variableNames, &val
	case string:
		variableNames.Variables = append(variableNames.Variables, val)
	case ResultVariablesTuple:
		variableNames = val
	}
	return variableNames, nil
}

func (a *AbacusVisitor) convertTupleResult(result interface{}) ResultTuple {
	values := NewResultTuple()
	switch val := result.(type) {
	case *big.Float:
		values.Values = append(values.Values, val)
	case ResultTuple:
		values = val
	}
	return values
}

func (a *AbacusVisitor) VisitVariableDeclaration(c *parser.VariableDeclarationContext) interface{} {
	resVariables := c.VariablesTuple().Accept(a)
	variableNames, err := a.convertVariablesTupleResult(resVariables)
	if err != nil {
		return *err
	}

	resValues := c.Tuple().Accept(a)
	values := a.convertTupleResult(resValues)

	if len(variableNames.Variables) != len(values.Values) {
		return ResultError("Wrong number of values " + strconv.FormatInt(int64(len(values.Values)), 10) + "; expected " + strconv.FormatInt(int64(len(variableNames.Variables)), 10))
	}

	for i, variable := range variableNames.Variables {
		a.vars[variable] = values.Values[i]
	}
	return ResultAssignment{Values: values.Values}
}

func (a *AbacusVisitor) VisitLambdaDeclaration(c *parser.LambdaDeclarationContext) interface{} {
	lambdaName := c.LAMBDA_VARIABLE().GetText()
	lambda := c.Lambda()

	// Check if 1) multiple vars 2) duped vars
	multipleVarLambda, ok := lambda.(*parser.MultiVariableLambdaContext)
	if ok {
		resVars := multipleVarLambda.VariablesTuple().Accept(a)
		_, err := a.convertVariablesTupleResult(resVars)
		if err != nil {
			return *err
		}
	}

	a.lambdas[lambdaName] = &Lambda{
		ctx: c,
	}

	formattedLambda := lambdaName + " = " + lambda.GetText()
	return ResultLambdaAssignment(formattedLambda)
}

func (a *AbacusVisitor) VisitEqualComparison(c *parser.EqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*big.Float)
	right := c.Expression(1).Accept(a).(*big.Float)
	return left.Cmp(right) == 0
}

func (a *AbacusVisitor) VisitLessComparison(c *parser.LessComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*big.Float)
	right := c.Expression(1).Accept(a).(*big.Float)
	return left.Cmp(right) == -1
}

func (a *AbacusVisitor) VisitGreaterComparison(c *parser.GreaterComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*big.Float)
	right := c.Expression(1).Accept(a).(*big.Float)
	return left.Cmp(right) == 1
}

func (a *AbacusVisitor) VisitLessOrEqualComparison(c *parser.LessOrEqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*big.Float)
	right := c.Expression(1).Accept(a).(*big.Float)
	return left.Cmp(right) <= 0
}

func (a *AbacusVisitor) VisitGreaterOrEqualComparison(c *parser.GreaterOrEqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*big.Float)
	right := c.Expression(1).Accept(a).(*big.Float)
	return left.Cmp(right) >= 0
}

func (a *AbacusVisitor) VisitMulDiv(c *parser.MulDivContext) interface{} {
	left := c.Expression(0).Accept(a).(*big.Float)
	right := c.Expression(1).Accept(a).(*big.Float)

	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserMUL:
		return Mul(left, right)
	case parser.AbacusLexerDIV:
		return Div(left, right)
	}
	return 0
}

func (a *AbacusVisitor) VisitAddSub(c *parser.AddSubContext) interface{} {
	left := c.Expression(0).Accept(a).(*big.Float)
	right := c.Expression(1).Accept(a).(*big.Float)

	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserADD:
		return Add(left, right)
	case parser.AbacusLexerSUB:
		return Sub(left, right)
	}
	return nil
}

func (a *AbacusVisitor) VisitPow(c *parser.PowContext) interface{} {
	left := c.Expression(0).Accept(a).(*big.Float) // TODO: This receives ResultTuple if a lambda is used; figure out a way to deal with that
	right := c.Expression(1).Accept(a).(*big.Float)
	return Pow(left, right)
}

func (a *AbacusVisitor) VisitParentheses(c *parser.ParenthesesContext) interface{} {
	return c.Expression().Accept(a)
}

func (a *AbacusVisitor) VisitAtomExpr(c *parser.AtomExprContext) interface{} {
	atomValue := c.Atom().Accept(a).(*big.Float)

	multiplier := New(1)

	if c.Sign() != nil {
		sign := c.Sign().Accept(a).(rune)
		if sign == '-' {
			multiplier = New(-1)
		}
	}

	return Zero().Mul(atomValue, multiplier)
}

func (a *AbacusVisitor) VisitFuncExpr(c *parser.FuncExprContext) interface{} {
	return c.Function().Accept(a)
}

func (a *AbacusVisitor) VisitSqrtFunction(c *parser.SqrtFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	return Sqrt(val)
}

func (a *AbacusVisitor) VisitLnFunction(c *parser.LnFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	return Log(val)
}

func (a *AbacusVisitor) VisitLogDefFunction(c *parser.LogDefFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	return Log(val)
}

func (a *AbacusVisitor) VisitLog2Function(c *parser.Log2FunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	return Div(Log(val), Log(New(2)))
}

func (a *AbacusVisitor) VisitLog10Function(c *parser.Log10FunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	return Div(Log(val), Log(New(10)))
}

func (a *AbacusVisitor) VisitFloorFunction(c *parser.FloorFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := val.Float64()
	return big.NewFloat(math.Floor(toFloat))
}

func (a *AbacusVisitor) VisitCeilFunction(c *parser.CeilFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := val.Float64()
	return big.NewFloat(math.Ceil(toFloat))
}
func (a *AbacusVisitor) VisitSinFunction(c *parser.SinFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := val.Float64()
	return big.NewFloat(math.Sin(toFloat))
}
func (a *AbacusVisitor) VisitCosFunction(c *parser.CosFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := val.Float64()
	return big.NewFloat(math.Cos(toFloat))
}
func (a *AbacusVisitor) VisitTanFunction(c *parser.TanFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := val.Float64()
	return big.NewFloat(math.Tan(toFloat))
}
func (a *AbacusVisitor) VisitExpFunction(c *parser.ExpFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	return Exp(val)
}

func (a *AbacusVisitor) VisitRoundDefFunction(c *parser.RoundDefFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := val.Float64()
	return big.NewFloat(math.Round(toFloat))
}

func (a *AbacusVisitor) VisitRound2Function(c *parser.Round2FunctionContext) interface{} {
	left := c.Expression(0).Accept(a).(*big.Float)
	right := c.Expression(1).Accept(a).(*big.Float)
	num, _ := left.Float64()
	digits, _ := right.Float64()
	mult := math.Pow(10, digits+1)
	precision = uint(digits)
	return big.NewFloat(math.Round(num*mult) / mult)
}

func (a *AbacusVisitor) VisitLogFunction(c *parser.LogFunctionContext) interface{} {
	left := c.Expression(0).Accept(a).(*big.Float)
	right := c.Expression(1).Accept(a).(*big.Float)

	return Div(Log(left), Log(right))
}

func (a *AbacusVisitor) VisitMinFunction(c *parser.MinFunctionContext) interface{} {
	resValues := c.Tuple().Accept(a)
	tuple := a.convertTupleResult(resValues)

	smallest := tuple.Values[0]

	for i := 1; i < len(tuple.Values); i++ {
		curr := tuple.Values[i]
		if curr.Cmp(smallest) == -1 {
			smallest = curr
		}
	}

	return smallest
}

func (a *AbacusVisitor) VisitMaxFunction(c *parser.MaxFunctionContext) interface{} {
	resValues := c.Tuple().Accept(a)
	tuple := a.convertTupleResult(resValues)

	biggest := tuple.Values[0]

	for i := 1; i < len(tuple.Values); i++ {
		curr := tuple.Values[i]
		if curr.Cmp(biggest) == 1 {
			biggest = curr
		}
	}

	return biggest
}

func (a *AbacusVisitor) VisitAvgFunction(c *parser.AvgFunctionContext) interface{} {
	resValues := c.Tuple().Accept(a)
	tuple := a.convertTupleResult(resValues)

	sum := tuple.Values[0]

	for i := 1; i < len(tuple.Values); i++ {
		curr := tuple.Values[i]
		sum = Add(sum, curr)
	}

	return Div(sum, New(float64(len(tuple.Values))))
}

func (a *AbacusVisitor) VisitConstant(c *parser.ConstantContext) interface{} {
	switch c.CONSTANT().GetText() {
	case "pi":
		return pi(precision)
	case "phi":
		return phi(precision)
	case "e":
		return e(precision)
	}
	return 0
}

func (a *AbacusVisitor) VisitNumber(c *parser.NumberContext) interface{} {
	numberString := c.SCIENTIFIC_NUMBER().GetText()

	out, _, err := big.ParseFloat(numberString, 10, precision, big.ToNearestEven)
	if err != nil {
		panic(err)
	}
	return out
}

func (a *AbacusVisitor) VisitPlusSign(c *parser.PlusSignContext) interface{} {
	return '+'
}

func (a *AbacusVisitor) VisitMinusSign(c *parser.MinusSignContext) interface{} {
	return '-'
}

func (a *AbacusVisitor) VisitSingleVariableLambda(c *parser.SingleVariableLambdaContext) interface{} {
	resValues := c.Tuple().Accept(a)
	tuple := a.convertTupleResult(resValues)
	return tuple
}

func (a *AbacusVisitor) VisitMultiVariableLambda(c *parser.MultiVariableLambdaContext) interface{} {
	resValues := c.Tuple().Accept(a)
	tuple := a.convertTupleResult(resValues)
	return tuple
}

func (a *AbacusVisitor) VisitLambdaExpr(c *parser.LambdaExprContext) interface{} {
	lambdaName := c.LAMBDA_VARIABLE().GetText()

	lambda, found := a.lambdas[lambdaName]
	if !found {
		return New(0)
	}

	// Handle recursion
	inLambda, nested := a.checkParentCtxForLambda(c.GetParent())
	if inLambda {
		// Recurrs
		if lambdaName == nested {
			if arguments.MaxRecurrences == 0 {
				return ResultError("recursion is disabled")
			}
			recurrences, ok := uint(0), false

			if recurrences, ok = a.lambdaRecursion[lambdaName]; !ok {
				a.lambdaRecursion[lambdaName] = 1
			} else if recurrences == arguments.MaxRecurrences {
				return New(float64(arguments.LastValueInRecursion))
			} else {
				a.lambdaRecursion[lambdaName] = recurrences + 1
			}
		}
	}

	parameters := make([]*big.Float, 0)
	if c.Tuple() != nil {
		resValues := c.Tuple().Accept(a)
		tuple := a.convertTupleResult(resValues)
		parameters = tuple.Values
	}

	switch val := lambda.ctx.Lambda().(type) {
	case *parser.SingleVariableLambdaContext:
		if len(parameters) < 1 {
			return ResultError("expected 1 parameter")
		}
		varName := val.VARIABLE().GetText()
		a.lambdaVars[lambdaVarName(lambdaName, varName)] = parameters[0]
		return val.Accept(a)
	case *parser.MultiVariableLambdaContext:
		resVars := val.VariablesTuple().Accept(a)
		variableNames, err := a.convertVariablesTupleResult(resVars)
		if err != nil {
			return *err
		}
		count := len(variableNames.Variables)
		s := ""
		if count > 1 {
			s = "s"
		}
		if len(parameters) < count {
			return ResultError("expected " + strconv.FormatInt(int64(count), 10) + " parameter" + s)
		}
		for i, varName := range variableNames.Variables {
			a.lambdaVars[lambdaVarName(lambdaName, varName)] = parameters[i]
		}
		return val.Accept(a)
	}
	return New(0)
}

func (a *AbacusVisitor) VisitVariable(c *parser.VariableContext) interface{} {
	var value *big.Float
	ok := false

	name := c.VARIABLE().GetText()

	inLambda, lambdaName := a.checkParentCtxForLambda(c.GetParent())
	if inLambda {
		if value, ok = a.lambdaVars[lambdaVarName(lambdaName, name)]; ok {
			return value
		}

		if value, ok = a.vars[name]; ok {
			return value
		}
	} else {
		if value, ok = a.vars[name]; ok {
			return value
		}
	}
	return big.NewFloat(0)
}

func (a *AbacusVisitor) checkParentCtxForLambda(c antlr.Tree) (bool, string) {
	ctx, ok := interface{}(c).(*parser.LambdaDeclarationContext)

	lambdaName := ""
	if ok {
		lambdaName = ctx.LAMBDA_VARIABLE().GetText()
	}

	if !ok && c.GetParent() != nil {
		return a.checkParentCtxForLambda(c.GetParent())
	}
	return ok, lambdaName
}

func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func lambdaVarName(lambdaName, varName string) string {
	return "$" + lambdaName + "$" + varName
}
