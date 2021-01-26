package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/cockroachdb/apd"
	"github.com/viktordanov/abacus/parser"
	"math"
	"strconv"
)

type Lambda struct {
	ctx        *parser.LambdaDeclarationContext
	parameters *RecursionParameters
}

type RecursionParameters struct {
	MaxRecurrences uint
	LastValue      *apd.Decimal
	StopWhen       parser.IComparisonContext
}

func NewRecursionParameters() *RecursionParameters {
	return &RecursionParameters{MaxRecurrences: 0, LastValue: newDecimal(0), StopWhen: nil}
}

var (
	logCache   map[string]*apd.Decimal
	decimalCtx *apd.Context
	PI         *apd.Decimal
	PHI        *apd.Decimal
	E          *apd.Decimal
)

func init() {
	logCache = make(map[string]*apd.Decimal)
	decimalCtx = apd.BaseContext.WithPrecision(arguments.Precision)
	cachedLog(apd.New(2, 0))
	cachedLog(apd.New(10, 0))

	PI, _, _ = decimalCtx.NewFromString("3." +
		"14159265358979323846264338327950288419716939937510" +
		"58209749445923078164062862089986280348253421170679" +
		"82148086513282306647093844609550582231725359408128" +
		"48111745028410270193852110555964462294895493038196" +
		"44288109756659334461284756482337867831652712019091" +
		"45648566923460348610454326648213393607260249141273" +
		"72458700660631558817488152092096282925409171536444")

	E, _, _ = decimalCtx.NewFromString("2." +
		"71828182845904523536028747135266249775724709369995" +
		"95749669676277240766303535475945713821785251664274" +
		"27466391932003059921817413596629043572900334295260" +
		"59563073813232862794349076323382988075319525101901" +
		"157383418793070215408914993488416750924147614606680" +
		"82264800168477411853742345442437107539077744992069" +
		"55170276183860626133138458300075204493382656029760")

	PHI, _, _ = decimalCtx.NewFromString("1." +
		"61803398874989484820458683436563811772030917980576" +
		"28621354486227052604628189024497072072041893911374" +
		"84754088075386891752126633862223536931793180060766" +
		"72635443338908659593958290563832266131992829026788" +
		"06752087668925017116962070322210432162695486262963" +
		"13614438149758701220340805887954454749246185695364" +
		"86444924104432077134494704956584678850987433944221")
}
func newDecimal(f float64) *apd.Decimal {
	res := apd.New(0, 0)
	res.SetFloat64(f)
	return res
}

func cachedLog(n *apd.Decimal) *apd.Decimal {
	out := newDecimal(0)

	var r *apd.Decimal
	var ok bool
	if r, ok = logCache[n.String()]; !ok {
		decimalCtx.Ln(out, n)
		logCache[n.String()] = newDecimal(0)
		logCache[n.String()].Set(out)
	} else {
		out.Set(r)
	}
	return out
}

type AbacusVisitor struct {
	antlr.ParseTreeVisitor
	vars                 map[string]*apd.Decimal
	lambdas              map[string]*Lambda
	lambdaVars           map[string]*apd.Decimal
	lambdaRecursion      map[string]uint
	lambdaRecursionStack map[string]uint
	decimalCtx           *apd.Context
}

func NewAbacusVisitor() *AbacusVisitor {
	return &AbacusVisitor{
		ParseTreeVisitor:     &parser.BaseAbacusVisitor{},
		vars:                 make(map[string]*apd.Decimal),
		lambdas:              make(map[string]*Lambda),
		lambdaVars:           make(map[string]*apd.Decimal),
		lambdaRecursion:      make(map[string]uint),
		lambdaRecursionStack: make(map[string]uint),
		decimalCtx:           apd.BaseContext.WithPrecision(arguments.Precision),
	}
}

func (a *AbacusVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.RootContext:
		a.lambdaRecursion = make(map[string]uint)
		a.lambdaRecursionStack = make(map[string]uint)
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
	val := ctx.Expression().Accept(a)

	switch v := val.(type) {
	case *apd.Decimal:
		resultTuple.Values = append(resultTuple.Values, v)
	case ResultTuple:
		resultTuple.Values = append(resultTuple.Values, v.Values...)
	}

	a.visitTupleTail(ctx.Tuple(), resultTuple)
}

func (a *AbacusVisitor) VisitTuple(c *parser.TupleContext) interface{} {
	if c.Tuple() == nil {
		return c.Expression().Accept(a)
	}

	evaledTuple := NewResultTuple()
	val := c.Expression().Accept(a)

	switch v := val.(type) {
	case *apd.Decimal:
		evaledTuple.Values = append(evaledTuple.Values, v)
	case ResultTuple:
		evaledTuple.Values = append(evaledTuple.Values, v.Values...)
	}
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
	case *apd.Decimal:
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
	multipleVarLambda, ok := lambda.(*parser.VariablesLambdaContext)
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
	left := c.Expression(0).Accept(a).(*apd.Decimal)
	right := c.Expression(1).Accept(a).(*apd.Decimal)
	return left.Cmp(right) == 0
}

func (a *AbacusVisitor) VisitLessComparison(c *parser.LessComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*apd.Decimal)
	right := c.Expression(1).Accept(a).(*apd.Decimal)
	return left.Cmp(right) == -1
}

func (a *AbacusVisitor) VisitGreaterComparison(c *parser.GreaterComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*apd.Decimal)
	right := c.Expression(1).Accept(a).(*apd.Decimal)
	return left.Cmp(right) == 1
}

func (a *AbacusVisitor) VisitLessOrEqualComparison(c *parser.LessOrEqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*apd.Decimal)
	right := c.Expression(1).Accept(a).(*apd.Decimal)
	return left.Cmp(right) <= 0
}

func (a *AbacusVisitor) VisitGreaterOrEqualComparison(c *parser.GreaterOrEqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*apd.Decimal)
	right := c.Expression(1).Accept(a).(*apd.Decimal)
	return left.Cmp(right) >= 0
}

func (a *AbacusVisitor) VisitMulDiv(c *parser.MulDivContext) interface{} {
	left := c.Expression(0).Accept(a).(*apd.Decimal)
	right := c.Expression(1).Accept(a).(*apd.Decimal)

	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserMUL:
		res := newDecimal(0)
		a.decimalCtx.Mul(res, left, right)
		return res
	case parser.AbacusLexerDIV:
		res := newDecimal(0)
		a.decimalCtx.Quo(res, left, right)
		return res
	}
	return 0
}

func (a *AbacusVisitor) VisitAddSub(c *parser.AddSubContext) interface{} {
	left := c.Expression(0).Accept(a).(*apd.Decimal)
	right := c.Expression(1).Accept(a).(*apd.Decimal)

	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserADD:
		res := newDecimal(0)
		a.decimalCtx.Add(res, left, right)
		return res
	case parser.AbacusLexerSUB:
		res := newDecimal(0)
		a.decimalCtx.Sub(res, left, right)
		return res
	}
	return nil
}

func (a *AbacusVisitor) VisitPow(c *parser.PowContext) interface{} {
	left := c.Expression(0).Accept(a).(*apd.Decimal)
	right := c.Expression(1).Accept(a).(*apd.Decimal)
	res := newDecimal(0)
	a.decimalCtx.Pow(res, left, right)
	return res
}

func (a *AbacusVisitor) VisitMod(c *parser.ModContext) interface{} {
	left := c.Expression(0).Accept(a).(*apd.Decimal)
	right := c.Expression(1).Accept(a).(*apd.Decimal)
	res := newDecimal(0)
	a.decimalCtx.Rem(res, left, right)
	return res
}

func (a *AbacusVisitor) VisitSignedExpr(c *parser.SignedExprContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)

	sign := c.Sign().Accept(a).(rune)
	if sign == '-' {
		val.Negative = !val.Negative
	}

	return val
}

func (a *AbacusVisitor) VisitParentheses(c *parser.ParenthesesContext) interface{} {
	return c.Expression().Accept(a)
}

func (a *AbacusVisitor) VisitAtomExpr(c *parser.AtomExprContext) interface{} {
	return c.Atom().Accept(a)
}

func (a *AbacusVisitor) VisitFuncExpr(c *parser.FuncExprContext) interface{} {
	return c.Function().Accept(a)
}

func (a *AbacusVisitor) VisitSqrtFunction(c *parser.SqrtFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)

	v := newDecimal(0)
	a.decimalCtx.Sqrt(v, val)
	return v
}

func (a *AbacusVisitor) VisitCbrtFunction(c *parser.CbrtFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)

	v := newDecimal(0)
	a.decimalCtx.Cbrt(v, val)
	return v
}

func (a *AbacusVisitor) VisitLnFunction(c *parser.LnFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)
	v := newDecimal(0)
	a.decimalCtx.Ln(v, val)
	return v
}

func (a *AbacusVisitor) VisitLogDefFunction(c *parser.LogDefFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)
	v := newDecimal(0)
	a.decimalCtx.Ln(v, val)
	return v
}

func (a *AbacusVisitor) VisitLog2Function(c *parser.Log2FunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)

	v := newDecimal(0)
	a.decimalCtx.Ln(v, val)
	a.decimalCtx.Quo(v, v, logCache["2"])
	return v
}

func (a *AbacusVisitor) VisitLog10Function(c *parser.Log10FunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)
	v := newDecimal(0)
	a.decimalCtx.Ln(v, val)
	base := cachedLog(newDecimal(10))
	a.decimalCtx.Quo(v, v, base)
	return v
}

func (a *AbacusVisitor) VisitFloorFunction(c *parser.FloorFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)
	v := newDecimal(0)
	a.decimalCtx.Floor(v, val)
	return v
}

func (a *AbacusVisitor) VisitCeilFunction(c *parser.CeilFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)
	v := newDecimal(0)
	a.decimalCtx.Ceil(v, val)
	return v
}
func (a *AbacusVisitor) VisitSinFunction(c *parser.SinFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)
	toFloat, _ := val.Float64()
	res := newDecimal(0)
	res, _ = res.SetFloat64(math.Sin(toFloat))
	return res
}
func (a *AbacusVisitor) VisitCosFunction(c *parser.CosFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)
	toFloat, _ := val.Float64()
	res := newDecimal(0)
	res, _ = res.SetFloat64(math.Cos(toFloat))
	return res
}
func (a *AbacusVisitor) VisitTanFunction(c *parser.TanFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)
	toFloat, _ := val.Float64()
	res := newDecimal(0)
	res, _ = res.SetFloat64(math.Tan(toFloat))
	return res
}
func (a *AbacusVisitor) VisitExpFunction(c *parser.ExpFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)
	v := newDecimal(0)
	a.decimalCtx.Exp(v, val)
	return v
}

func (a *AbacusVisitor) VisitAbsFunction(c *parser.AbsFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)
	v := newDecimal(0)
	a.decimalCtx.Abs(v, val)
	return v
}

func (a *AbacusVisitor) VisitRoundDefFunction(c *parser.RoundDefFunctionContext) interface{} {
	val := c.Expression().Accept(a).(*apd.Decimal)
	v := newDecimal(0)
	a.decimalCtx.Round(v, val)
	return v
}

func (a *AbacusVisitor) VisitRound2Function(c *parser.Round2FunctionContext) interface{} {
	left := c.Expression(0).Accept(a).(*apd.Decimal)
	right := c.Expression(1).Accept(a).(*apd.Decimal)

	intValue, _ := right.Int64()
	exponent := apd.New(10, int32(intValue))

	v := newDecimal(0)
	a.decimalCtx.Mul(v, left, exponent)
	a.decimalCtx.Round(v, v)
	a.decimalCtx.Quo(v, v, exponent)

	return v
}

func (a *AbacusVisitor) VisitLogFunction(c *parser.LogFunctionContext) interface{} {
	left := c.Expression(0).Accept(a).(*apd.Decimal)
	right := c.Expression(1).Accept(a).(*apd.Decimal)

	v := newDecimal(0)

	a.decimalCtx.Ln(v, left)
	base := cachedLog(right)
	a.decimalCtx.Quo(v, v, base)
	return v
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

	biggest := newDecimal(0)
	biggest.Set(tuple.Values[0])

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

	sum := newDecimal(0)
	sum.Set(tuple.Values[0])

	for i := 1; i < len(tuple.Values); i++ {
		curr := tuple.Values[i]
		a.decimalCtx.Add(sum, sum, curr)
	}
	a.decimalCtx.Quo(sum, sum, apd.New(int64(len(tuple.Values)), 0))
	return sum
}

func (a *AbacusVisitor) VisitUntilFunction(c *parser.UntilFunctionContext) interface{} {
	resValues := c.Tuple().Accept(a)
	tuple := a.convertTupleResult(resValues)

	arg := c.Expression().Accept(a).(*apd.Decimal)
	intValue, _ := arg.Int64()

	newTuple := NewResultTuple()
	length := int64(len(tuple.Values))

	if intValue > length {
		intValue = length
	}
	for i := 0; i < int(intValue); i++ {
		newTuple.Values = append(newTuple.Values, tuple.Values[i])
	}

	if len(newTuple.Values) == 0 {
		return newDecimal(0)
	}
	return newTuple
}

func (a *AbacusVisitor) VisitFromFunction(c *parser.FromFunctionContext) interface{} {
	resValues := c.Tuple().Accept(a)
	tuple := a.convertTupleResult(resValues)
	arg := tuple.Values[len(tuple.Values)-1]
	intValue, _ := arg.Int64()

	newTuple := NewResultTuple()
	length := int64(len(tuple.Values))

	if intValue < 0 {
		intValue = 0
	}
	for i := intValue; i < length-1; i++ {
		newTuple.Values = append(newTuple.Values, tuple.Values[i])
	}

	if len(newTuple.Values) == 0 {
		return newDecimal(0)
	}
	return newTuple
}

func (a *AbacusVisitor) VisitReverseFunction(c *parser.ReverseFunctionContext) interface{} {
	resValues := c.Tuple().Accept(a)
	tuple := a.convertTupleResult(resValues)

	values := make([]*apd.Decimal, 0)
	for i := len(tuple.Values) - 1; i >= 0; i-- {
		values = append(values, tuple.Values[i])
	}
	tuple.Values = values
	return tuple
}

func (a *AbacusVisitor) VisitNthFunction(c *parser.NthFunctionContext) interface{} {
	resValues := c.Tuple().Accept(a)
	tuple := a.convertTupleResult(resValues)

	arg1 := c.Expression().Accept(a).(*apd.Decimal)
	intValue1, _ := arg1.Int64()

	if intValue1 >= int64(len(tuple.Values)) || intValue1 < 0 {
		return newDecimal(0)
	}
	return tuple.Values[intValue1]
}

func (a *AbacusVisitor) VisitConstant(c *parser.ConstantContext) interface{} {
	switch c.CONSTANT().GetText() {
	case "pi":
		return PI
	case "phi":
		return PHI
	case "e":
		return E
	}
	return 0
}

func (a *AbacusVisitor) VisitPercent(c *parser.PercentContext) interface{} {
	numberString := c.Expression().Accept(a).(*apd.Decimal)

	a.decimalCtx.Quo(numberString, numberString, newDecimal(100))
	return numberString
}

func (a *AbacusVisitor) VisitNumber(c *parser.NumberContext) interface{} {
	numberString := c.SCIENTIFIC_NUMBER().GetText()

	out, _, err := apd.NewFromString(numberString)
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

func (a *AbacusVisitor) VisitVariablesLambda(c *parser.VariablesLambdaContext) interface{} {
	resValues := c.Tuple().Accept(a)
	tuple := a.convertTupleResult(resValues)
	return tuple
}

func (a *AbacusVisitor) VisitNullArityLambda(c *parser.NullArityLambdaContext) interface{} {
	resValues := c.Tuple().Accept(a)
	tuple := a.convertTupleResult(resValues)
	return tuple
}

func (a *AbacusVisitor) VisitRecursionParameters(c *parser.RecursionParametersContext) interface{} {
	recursionParameters := NewRecursionParameters()
	inLambda, lambda := a.checkParentCtxForLambda(c.GetParent())

	for i := 0; i < len(c.AllExpression()); i++ {
		if inLambda {
			c.Expression(i).SetParent(a.lambdas[lambda].ctx)
		}
		val := c.Expression(i).Accept(a).(*apd.Decimal)
		switch i {
		case 0:
			intValue, _ := val.Int64()
			recursionParameters.MaxRecurrences = uint(intValue)
		case 1:
			v := newDecimal(0)
			v.Set(val)
			recursionParameters.LastValue = v
		}
	}
	recursionParameters.StopWhen = c.Comparison()
	return recursionParameters
}

func (a *AbacusVisitor) VisitLambdaExpr(c *parser.LambdaExprContext) interface{} {
	lambdaName := c.LAMBDA_VARIABLE().GetText()

	lambda, found := a.lambdas[lambdaName]
	if !found {
		return newDecimal(0)
	}

	parameters := make([]*apd.Decimal, 0)
	if c.Tuple() != nil {
		resValues := c.Tuple().Accept(a)
		tuple := a.convertTupleResult(resValues)
		parameters = tuple.Values
	}

	// Handle recursion
	inLambda, nested := a.checkParentCtxForLambda(c.GetParent())

	recursionParameters := NewRecursionParameters()
	if c.RecursionParameters() != nil {
		recursionParameters = c.RecursionParameters().Accept(a).(*RecursionParameters)
	}
	if recursionParameters.StopWhen != nil {
		recursionParameters.StopWhen.SetParent(lambda.ctx)
	}
	if lambda.parameters == nil {
		lambda.parameters = recursionParameters
	}

	if inLambda {
		// Recurrs
		if lambdaName == nested {
			if lambda.parameters.MaxRecurrences == 0 {
				return ResultError("recursion is disabled")
			}
			recurrences, ok := uint(0), false

			if recurrences, ok = a.lambdaRecursion[lambdaName]; !ok {
				a.lambdaRecursion[lambdaName] = 1
			}
			if _, ok = a.lambdaRecursionStack[lambdaName]; !ok {
				a.lambdaRecursionStack[lambdaName] = 1
			}
			shouldStop := false
			if lambda.parameters.StopWhen != nil {
				condition := lambda.parameters.StopWhen.Accept(a)
				shouldStop = condition.(bool)
			}
			if shouldStop {
				v := newDecimal(0)
				v.Set(lambda.parameters.LastValue)
				return 0
			}
			if recurrences == lambda.parameters.MaxRecurrences {
				v := newDecimal(0)
				v.Set(lambda.parameters.LastValue)
				return v
			} else {
				a.lambdaRecursion[lambdaName]++
				a.lambdaRecursionStack[lambdaName]++
			}
		} else {
			a.lambdaRecursionStack[lambdaName] = 1
		}
	}

	//log.Printf("[%s] %v %v\n", lambdaName, inLambda, parameters)

	switch val := lambda.ctx.Lambda().(type) {
	case *parser.VariablesLambdaContext:
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
			a.lambdaVars[lambdaVarName(lambdaName, varName, a.lambdaRecursionStack[lambdaName])] = parameters[i]
		}
		r := val.Accept(a)
		a.lambdaRecursionStack[lambdaName]--

		switch rr := r.(type) {
		case ResultTuple:
			if len(rr.Values) == 1 {
				v := newDecimal(0)
				v.Set(rr.Values[0])
				return v
			}
		}
		//log.Printf("[%s] %+v %+v\n", lambdaName, r, parameters)
		return r
	case *parser.NullArityLambdaContext:
		r := val.Accept(a)
		a.lambdaRecursionStack[lambdaName]--
		//log.Printf("[%s] %+v %+v\n", lambdaName, r, parameters)

		switch rr := r.(type) {
		case ResultTuple:
			if len(rr.Values) == 1 {
				v := newDecimal(0)
				v.Set(rr.Values[0])
				return v
			}
		}
		return r
	}
	return newDecimal(0)
}

func (a *AbacusVisitor) VisitVariable(c *parser.VariableContext) interface{} {
	var value *apd.Decimal
	ok := false

	name := c.VARIABLE().GetText()

	inLambda, lambdaName := a.checkParentCtxForLambda(c.GetParent())
	if inLambda {
		if value, ok = a.lambdaVars[lambdaVarName(lambdaName, name, a.lambdaRecursionStack[lambdaName])]; ok {
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
	return newDecimal(0)
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

func lambdaVarName(lambdaName, varName string, stack uint) string {
	out := ""
	if stack == 0 {
		stack = 1
	}
	for i := uint(0); i < stack; i++ {
		out += "$" + lambdaName
	}
	return out + "$" + varName
}
