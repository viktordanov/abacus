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
	LastValue      ResultNumber
	StopWhen       parser.IBoolExpressionContext
}

func NewRecursionParameters() *RecursionParameters {
	return &RecursionParameters{MaxRecurrences: 0, LastValue: newNumber(0), StopWhen: nil}
}

var (
	logCache   map[string]ResultNumber
	decimalCtx *apd.Context
	PI         ResultNumber
	PHI        ResultNumber
	E          ResultNumber
)

func init() {
	logCache = make(map[string]ResultNumber)
	decimalCtx = apd.BaseContext.WithPrecision(arguments.Precision)
	cachedLog(ResultNumber{apd.New(2, 0)})
	cachedLog(ResultNumber{apd.New(10, 0)})

	pi, _, _ := decimalCtx.NewFromString("3." +
		"14159265358979323846264338327950288419716939937510" +
		"58209749445923078164062862089986280348253421170679" +
		"82148086513282306647093844609550582231725359408128" +
		"48111745028410270193852110555964462294895493038196" +
		"44288109756659334461284756482337867831652712019091" +
		"45648566923460348610454326648213393607260249141273" +
		"72458700660631558817488152092096282925409171536444")

	e, _, _ := decimalCtx.NewFromString("2." +
		"71828182845904523536028747135266249775724709369995" +
		"95749669676277240766303535475945713821785251664274" +
		"27466391932003059921817413596629043572900334295260" +
		"59563073813232862794349076323382988075319525101901" +
		"157383418793070215408914993488416750924147614606680" +
		"82264800168477411853742345442437107539077744992069" +
		"55170276183860626133138458300075204493382656029760")

	phi, _, _ := decimalCtx.NewFromString("1." +
		"61803398874989484820458683436563811772030917980576" +
		"28621354486227052604628189024497072072041893911374" +
		"84754088075386891752126633862223536931793180060766" +
		"72635443338908659593958290563832266131992829026788" +
		"06752087668925017116962070322210432162695486262963" +
		"13614438149758701220340805887954454749246185695364" +
		"86444924104432077134494704956584678850987433944221")

	PI = ResultNumber{pi}
	PHI = ResultNumber{phi}
	E = ResultNumber{e}
}

func newNumber(f float64) ResultNumber {
	res := apd.New(0, 0)
	res.SetFloat64(f)
	return ResultNumber{res}
}

func cachedLog(n ResultNumber) ResultNumber {
	out := newNumber(0)

	var r ResultNumber
	var ok bool
	if r, ok = logCache[n.String()]; !ok {
		decimalCtx.Ln(out.Decimal, n.Decimal)
		logCache[n.String()] = newNumber(0)
		logCache[n.String()].Set(out.Decimal)
	} else {
		out.Set(r.Decimal)
	}
	return out
}

type AbacusVisitor struct {
	antlr.ParseTreeVisitor
	vars                 map[string]ResultNumber
	lambdas              map[string]*Lambda
	lambdaVars           map[string]ResultNumber
	lambdaRecursion      map[string]uint
	lambdaRecursionStack map[string]uint
	decimalCtx           *apd.Context
}

func NewAbacusVisitor() *AbacusVisitor {
	return &AbacusVisitor{
		ParseTreeVisitor:     &parser.BaseAbacusVisitor{},
		vars:                 make(map[string]ResultNumber),
		lambdas:              make(map[string]*Lambda),
		lambdaVars:           make(map[string]ResultNumber),
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

	case *parser.BoolExpressionContext:
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
	if c.BoolExpression() != nil {
		return c.BoolExpression().Accept(a)
	}
	return NewResult(nil).WithErrors(nil, "edge case")
}

func (a *AbacusVisitor) visitTupleTail(c parser.ITupleContext, resultTuple *ResultTuple) {
	ctx, ok := c.(*parser.TupleContext)
	if !ok || ctx == nil {
		return
	}
	val := ctx.Expression().Accept(a).(*Result)

	switch v := val.Value.(type) {
	case ResultNumber:
		*resultTuple = append(*resultTuple, v)
	case ResultTuple:
		*resultTuple = append(*resultTuple, v...)
	}

	a.visitTupleTail(ctx.Tuple(), resultTuple)
}

func (a *AbacusVisitor) VisitTuple(c *parser.TupleContext) interface{} {
	if c.Tuple() == nil {
		return c.Expression().Accept(a)
	}

	evaledTuple := ResultTuple{}
	val := c.Expression().Accept(a).(*Result)

	switch v := val.Value.(type) {
	case ResultNumber:
		evaledTuple = append(evaledTuple, v)
	case ResultTuple:
		evaledTuple = append(evaledTuple, v...)
	}
	a.visitTupleTail(c.Tuple(), &evaledTuple)

	return NewResult(evaledTuple)
}

func (a *AbacusVisitor) visitVariableTupleTail(c parser.IVariablesTupleContext, resultTuple *ResultVariablesTuple) {
	ctx, ok := c.(*parser.VariablesTupleContext)
	if !ok || ctx == nil {
		return
	}
	val := ctx.VARIABLE().GetText()
	*resultTuple = append(*resultTuple, ResultString(val))
	a.visitVariableTupleTail(ctx.VariablesTuple(), resultTuple)
}

func (a *AbacusVisitor) VisitVariablesTuple(c *parser.VariablesTupleContext) interface{} {
	if c.VariablesTuple() == nil {
		return NewResult(ResultString(c.VARIABLE().GetText()))
	}

	evaledTuple := ResultVariablesTuple{}

	val := c.VARIABLE().GetText()
	evaledTuple = append(evaledTuple, ResultString(val))
	a.visitVariableTupleTail(c.VariablesTuple(), &evaledTuple)

	foundVars := make(map[string]bool)
	for _, variable := range evaledTuple {
		if _, ok := foundVars[variable.String()]; !ok {
			foundVars[variable.String()] = true
		} else {
			return NewResult(evaledTuple).WithErrors(nil, "duplicate variable name \""+variable.String()+"\"")
		}
	}

	return NewResult(evaledTuple)
}

func (a *AbacusVisitor) convertVariablesTupleResult(result *Result) {
	switch val := result.Value.(type) {
	case ResultString:
		result.Value = ResultVariablesTuple{val}
	}
}

func (a *AbacusVisitor) convertTupleResult(result *Result) {
	switch val := result.Value.(type) {
	case ResultNumber:
		result.Value = ResultTuple{val}
	}
}

func (a *AbacusVisitor) VisitVariableDeclaration(c *parser.VariableDeclarationContext) interface{} {
	variablesRes := c.VariablesTuple().Accept(a).(*Result)
	if hasErrors(variablesRes) {
		return variablesRes
	}
	a.convertVariablesTupleResult(variablesRes)

	valuesRes := c.Tuple().Accept(a).(*Result)
	if hasErrors(variablesRes) {
		return variablesRes
	}
	a.convertTupleResult(valuesRes)

	if variablesRes.Length() != valuesRes.Length() {
		return variablesRes.WithErrors(nil, "wrong number of valuesRes "+strconv.FormatInt(int64(valuesRes.Length()), 10)+"; expected "+strconv.FormatInt(int64(variablesRes.Length()), 10))
	}

	variables, ok := variablesRes.Value.(ResultVariablesTuple)
	if !ok {
		panic("unable to convert variables result to its raw type")
	}
	values, ok := valuesRes.Value.(ResultTuple)
	if !ok {
		panic("unable to convert values result to its raw type")
	}

	for i, variable := range variables {
		a.vars[variable.String()] = values[i]
	}
	return NewResult(ResultAssignment(values))
}

func (a *AbacusVisitor) VisitLambdaDeclaration(c *parser.LambdaDeclarationContext) interface{} {
	lambdaName := c.LAMBDA_VARIABLE().GetText()
	lambda := c.Lambda()

	// Check if 1) multiple vars && duped vars
	multipleVarLambda, ok := lambda.(*parser.VariablesLambdaContext)
	if ok {
		variablesResult := multipleVarLambda.VariablesTuple().Accept(a).(*Result)
		if hasErrors(variablesResult) {
			return variablesResult
		}
	}

	a.lambdas[lambdaName] = &Lambda{
		ctx: c,
	}

	formattedLambda := lambdaName + " = " + lambda.GetText()
	return NewResult(ResultLambdaAssignment(formattedLambda))
}

func (a *AbacusVisitor) VisitEqualComparison(c *parser.EqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultNumber)
	if !ok {
		panic("unable to cast right to ResultNumber")
	}
	rightVal, ok := right.Value.(ResultNumber)
	if !ok {
		panic("unable to cast left to ResultNumber")
	}

	return NewResult(ResultBool(leftVal.Cmp(rightVal.Decimal) == 0))
}

func (a *AbacusVisitor) VisitLessComparison(c *parser.LessComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultNumber)
	if !ok {
		panic("unable to cast right to ResultNumber")
	}
	rightVal, ok := right.Value.(ResultNumber)
	if !ok {
		panic("unable to cast left to ResultNumber")
	}

	return NewResult(ResultBool(leftVal.Cmp(rightVal.Decimal) == -1))
}

func (a *AbacusVisitor) VisitGreaterComparison(c *parser.GreaterComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultNumber)
	if !ok {
		panic("unable to cast right to ResultNumber")
	}
	rightVal, ok := right.Value.(ResultNumber)
	if !ok {
		panic("unable to cast left to ResultNumber")
	}

	return NewResult(ResultBool(leftVal.Cmp(rightVal.Decimal) == 1))
}

func (a *AbacusVisitor) VisitLessOrEqualComparison(c *parser.LessOrEqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultNumber)
	if !ok {
		panic("unable to cast right to ResultNumber")
	}
	rightVal, ok := right.Value.(ResultNumber)
	if !ok {
		panic("unable to cast left to ResultNumber")
	}

	return NewResult(ResultBool(leftVal.Cmp(rightVal.Decimal) <= 0))
}

func (a *AbacusVisitor) VisitGreaterOrEqualComparison(c *parser.GreaterOrEqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultNumber)
	if !ok {
		panic("unable to cast right to ResultNumber")
	}
	rightVal, ok := right.Value.(ResultNumber)
	if !ok {
		panic("unable to cast left to ResultNumber")
	}

	return NewResult(ResultBool(leftVal.Cmp(rightVal.Decimal) >= 0))
}

func (a *AbacusVisitor) VisitAndOrXor(c *parser.AndOrXorContext) interface{} {
	left := c.BoolExpression(0).Accept(a).(*Result)
	right := c.BoolExpression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultBool)
	if !ok {
		panic("unable to cast right to ResultBool")
	}
	rightVal, ok := right.Value.(ResultBool)
	if !ok {
		panic("unable to cast left to ResultBool")
	}

	op := c.GetOp().GetTokenType()
	result := ResultBool(false)

	switch op {
	case parser.AbacusParserAND:
		result = leftVal && rightVal
	case parser.AbacusParserOR:
		result = leftVal || rightVal
	case parser.AbacusParserXOR:
		result = leftVal != rightVal

	}
	return NewResult(result)
}

func (a *AbacusVisitor) VisitNot(c *parser.NotContext) interface{} {
	valRes := c.BoolExpression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}
	val, ok := valRes.Value.(ResultBool)
	if !ok {
		panic("unable to cast right to ResultBool")
	}
	return NewResult(!val)
}

func (a *AbacusVisitor) VisitParenthesesBoolean(c *parser.ParenthesesBooleanContext) interface{} {
	return c.BoolExpression().Accept(a)
}

func (a *AbacusVisitor) VisitBoolAtom(c *parser.BoolAtomContext) interface{} {
	val := c.GetText()
	if val == "true" {
		return ResultBool(true)
	}
	return ResultBool(false)
}
func (a *AbacusVisitor) VisitBooleanAtom(c *parser.BooleanAtomContext) interface{} {
	return c.BoolAtom().Accept(a)
}

func (a *AbacusVisitor) VisitMulDiv(c *parser.MulDivContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultNumber)
	if !ok {
		panic("unable to cast right to ResultNumber")
	}
	rightVal, ok := right.Value.(ResultNumber)
	if !ok {
		panic("unable to cast left to ResultNumber")
	}

	res := newNumber(0)
	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserMUL:
		a.decimalCtx.Mul(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	case parser.AbacusLexerDIV:
		a.decimalCtx.Quo(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	}
	return NewResult(res)
}

func (a *AbacusVisitor) VisitAddSub(c *parser.AddSubContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultNumber)
	if !ok {
		panic("unable to cast right to ResultNumber")
	}
	rightVal, ok := right.Value.(ResultNumber)
	if !ok {
		panic("unable to cast left to ResultNumber")
	}

	res := newNumber(0)
	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserADD:
		a.decimalCtx.Add(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	case parser.AbacusLexerSUB:
		a.decimalCtx.Sub(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	}
	return NewResult(res)
}

func (a *AbacusVisitor) VisitPow(c *parser.PowContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultNumber)
	if !ok {
		panic("unable to cast right to ResultNumber")
	}
	rightVal, ok := right.Value.(ResultNumber)
	if !ok {
		panic("unable to cast left to ResultNumber")
	}
	res := newNumber(0)
	a.decimalCtx.Pow(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	return NewResult(res)
}

func (a *AbacusVisitor) VisitMod(c *parser.ModContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultNumber)
	if !ok {
		panic("unable to cast right to ResultNumber")
	}
	rightVal, ok := right.Value.(ResultNumber)
	if !ok {
		panic("unable to cast left to ResultNumber")
	}
	res := newNumber(0)
	a.decimalCtx.Rem(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	return NewResult(res)
}

func (a *AbacusVisitor) VisitSignedExpr(c *parser.SignedExprContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	sign := c.Sign().Accept(a).(rune)
	if sign == '-' {
		val.Negative = !val.Negative
	}

	return NewResult(val)
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
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	v := newNumber(0)
	a.decimalCtx.Sqrt(v.Decimal, val.Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitCbrtFunction(c *parser.CbrtFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	v := newNumber(0)
	a.decimalCtx.Cbrt(v.Decimal, val.Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitLnFunction(c *parser.LnFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	v := newNumber(0)
	a.decimalCtx.Ln(v.Decimal, val.Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitLogDefFunction(c *parser.LogDefFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	v := newNumber(0)
	a.decimalCtx.Ln(v.Decimal, val.Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitLog2Function(c *parser.Log2FunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	v := newNumber(0)
	a.decimalCtx.Ln(v.Decimal, val.Decimal)
	a.decimalCtx.Quo(v.Decimal, v.Decimal, cachedLog(newNumber(2)).Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitLog10Function(c *parser.Log10FunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	v := newNumber(0)
	a.decimalCtx.Ln(v.Decimal, val.Decimal)
	base := cachedLog(newNumber(10))
	a.decimalCtx.Quo(v.Decimal, v.Decimal, base.Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitFloorFunction(c *parser.FloorFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	v := newNumber(0)
	a.decimalCtx.Floor(v.Decimal, val.Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitCeilFunction(c *parser.CeilFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	v := newNumber(0)
	a.decimalCtx.Ceil(v.Decimal, val.Decimal)
	return NewResult(v)
}
func (a *AbacusVisitor) VisitSinFunction(c *parser.SinFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	toFloat, _ := val.Float64()
	v := newNumber(0)
	v.Decimal, _ = v.Decimal.SetFloat64(math.Sin(toFloat))
	return NewResult(v)
}
func (a *AbacusVisitor) VisitCosFunction(c *parser.CosFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	toFloat, _ := val.Float64()
	v := newNumber(0)
	v.Decimal, _ = v.Decimal.SetFloat64(math.Cos(toFloat))
	return NewResult(v)
}
func (a *AbacusVisitor) VisitTanFunction(c *parser.TanFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	toFloat, _ := val.Float64()
	v := newNumber(0)
	v.Decimal, _ = v.Decimal.SetFloat64(math.Tan(toFloat))
	return NewResult(v)
}
func (a *AbacusVisitor) VisitExpFunction(c *parser.ExpFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	v := newNumber(0)
	a.decimalCtx.Exp(v.Decimal, val.Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitAbsFunction(c *parser.AbsFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	v := newNumber(0)
	a.decimalCtx.Abs(v.Decimal, val.Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitRoundDefFunction(c *parser.RoundDefFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast val to ResultNumber")
	}

	v := newNumber(0)
	a.decimalCtx.Round(v.Decimal, val.Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitRound2Function(c *parser.Round2FunctionContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultNumber)
	if !ok {
		panic("unable to cast right to ResultNumber")
	}
	rightVal, ok := right.Value.(ResultNumber)
	if !ok {
		panic("unable to cast left to ResultNumber")
	}

	intValue, _ := rightVal.Decimal.Int64()
	exponent := apd.New(10, int32(intValue))

	v := newNumber(0)
	a.decimalCtx.Mul(v.Decimal, leftVal.Decimal, exponent)
	a.decimalCtx.Round(v.Decimal, v.Decimal)
	a.decimalCtx.Quo(v.Decimal, v.Decimal, exponent)

	return NewResult(v)
}

func (a *AbacusVisitor) VisitLogFunction(c *parser.LogFunctionContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(ResultNumber)
	if !ok {
		panic("unable to cast right to ResultNumber")
	}
	rightVal, ok := right.Value.(ResultNumber)
	if !ok {
		panic("unable to cast left to ResultNumber")
	}

	v := newNumber(0)

	a.decimalCtx.Ln(v.Decimal, leftVal.Decimal)
	base := cachedLog(rightVal)
	a.decimalCtx.Quo(v.Decimal, v.Decimal, base.Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitMinFunction(c *parser.MinFunctionContext) interface{} {
	tupleRes := c.Tuple().Accept(a).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	a.convertTupleResult(tupleRes)

	tuple, ok := tupleRes.Value.(ResultTuple)
	if !ok {
		panic("unable to cast tupleRes to ResultTuple")
	}

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

func (a *AbacusVisitor) VisitMaxFunction(c *parser.MaxFunctionContext) interface{} {
	tupleRes := c.Tuple().Accept(a).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	a.convertTupleResult(tupleRes)

	tuple, ok := tupleRes.Value.(ResultTuple)
	if !ok {
		panic("unable to cast tupleRes to ResultTuple")
	}

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

func (a *AbacusVisitor) VisitAvgFunction(c *parser.AvgFunctionContext) interface{} {
	tupleRes := c.Tuple().Accept(a).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	a.convertTupleResult(tupleRes)

	tuple, ok := tupleRes.Value.(ResultTuple)
	if !ok {
		panic("unable to cast tupleRes to ResultTuple")
	}

	sum := newNumber(0)
	sum.Set(tuple[0].Decimal)

	for i := 1; i < len(tuple); i++ {
		curr := tuple[i]
		a.decimalCtx.Add(sum.Decimal, sum.Decimal, curr.Decimal)
	}
	a.decimalCtx.Quo(sum.Decimal, sum.Decimal, apd.New(int64(len(tuple)), 0))
	return NewResult(sum)
}

func (a *AbacusVisitor) VisitUntilFunction(c *parser.UntilFunctionContext) interface{} {
	tupleRes := c.Tuple().Accept(a).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	a.convertTupleResult(tupleRes)

	tuple, ok := tupleRes.Value.(ResultTuple)
	if !ok {
		panic("unable to cast tupleRes to ResultTuple")
	}

	argRes := c.Expression().Accept(a).(*Result)
	if hasErrors(argRes) {
		return argRes
	}

	arg, ok := argRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast argRes to ResultNumber")
	}
	intValue, _ := arg.Int64()

	newTuple := ResultTuple{}
	length := int64(len(tuple))

	if intValue > length {
		intValue = length
	}
	for i := 0; i < int(intValue); i++ {
		newTuple = append(newTuple, tuple[i])
	}

	if len(newTuple) == 0 {
		return NewResult(newNumber(0))
	}
	return NewResult(newTuple)
}

func (a *AbacusVisitor) VisitFromFunction(c *parser.FromFunctionContext) interface{} {
	tupleRes := c.Tuple().Accept(a).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	a.convertTupleResult(tupleRes)

	tuple, ok := tupleRes.Value.(ResultTuple)
	if !ok {
		panic("unable to cast tupleRes to ResultTuple")
	}
	arg := tuple[len(tuple)-1]
	intValue, _ := arg.Int64()

	newTuple := ResultTuple{}
	length := int64(len(tuple))

	if intValue < 0 {
		intValue = 0
	}
	for i := intValue; i < length-1; i++ {
		newTuple = append(newTuple, tuple[i])
	}

	if len(newTuple) == 0 {
		return NewResult(newNumber(0))
	}
	return NewResult(newTuple)
}

func (a *AbacusVisitor) VisitReverseFunction(c *parser.ReverseFunctionContext) interface{} {
	tupleRes := c.Tuple().Accept(a).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	a.convertTupleResult(tupleRes)

	tuple, ok := tupleRes.Value.(ResultTuple)
	if !ok {
		panic("unable to cast tupleRes to ResultTuple")
	}

	newTuple := ResultTuple{}
	for i := len(tuple) - 1; i >= 0; i-- {
		newTuple = append(newTuple, tuple[i])
	}
	return NewResult(newTuple)
}

func (a *AbacusVisitor) VisitNthFunction(c *parser.NthFunctionContext) interface{} {
	tupleRes := c.Tuple().Accept(a).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	a.convertTupleResult(tupleRes)

	tuple, ok := tupleRes.Value.(ResultTuple)
	if !ok {
		panic("unable to cast tupleRes to ResultTuple")
	}

	arg1 := c.Expression().Accept(a).(*apd.Decimal)
	intValue1, _ := arg1.Int64()

	if intValue1 >= int64(len(tuple)) || intValue1 < 0 {
		return NewResult(newNumber(0))
	}
	return NewResult(tuple[intValue1])
}

func (a *AbacusVisitor) VisitConstant(c *parser.ConstantContext) interface{} {
	switch c.CONSTANT().GetText() {
	case "pi":
		return NewResult(PI)
	case "phi":
		return NewResult(PHI)
	case "e":
		return NewResult(E)
	}
	return NewResult(newNumber(0))
}

func (a *AbacusVisitor) VisitPercent(c *parser.PercentContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}
	val, ok := valRes.Value.(ResultNumber)
	if !ok {
		panic("unable to cast valRes to ResultNumber")
	}

	a.decimalCtx.Quo(val.Decimal, val.Decimal, newNumber(100).Decimal)
	return NewResult(val)
}

func (a *AbacusVisitor) VisitNumber(c *parser.NumberContext) interface{} {
	numberString := c.SCIENTIFIC_NUMBER().GetText()

	out, _, err := apd.NewFromString(numberString)
	if err != nil {
		panic(err)
	}
	return NewResult(ResultNumber{out})
}

func (a *AbacusVisitor) VisitPlusSign(c *parser.PlusSignContext) interface{} {
	return '+'
}

func (a *AbacusVisitor) VisitMinusSign(c *parser.MinusSignContext) interface{} {
	return '-'
}

func (a *AbacusVisitor) VisitVariablesLambda(c *parser.VariablesLambdaContext) interface{} {
	tupleRes := c.Tuple().Accept(a).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	a.convertTupleResult(tupleRes)
	tuple, ok := tupleRes.Value.(ResultTuple)
	if !ok {
		panic("unable to cast tupleRes to ResultTuple")
	}

	return NewResult(tuple)
}

func (a *AbacusVisitor) VisitNullArityLambda(c *parser.NullArityLambdaContext) interface{} {
	tupleRes := c.Tuple().Accept(a).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	a.convertTupleResult(tupleRes)
	tuple, ok := tupleRes.Value.(ResultTuple)
	if !ok {
		panic("unable to cast tupleRes to ResultTuple")
	}

	return NewResult(tuple)
}

// TODO: Errors aren't handled
func (a *AbacusVisitor) VisitRecursionParameters(c *parser.RecursionParametersContext) interface{} {
	recursionParameters := NewRecursionParameters()
	inLambda, lambda := a.checkParentCtxForLambda(c.GetParent())

	for i := 0; i < len(c.AllExpression()); i++ {
		if inLambda {
			c.Expression(i).SetParent(a.lambdas[lambda].ctx)
		}
		valRes := c.Expression(i).Accept(a).(*Result)
		if hasErrors(valRes) {
			return valRes
		}
		val, ok := valRes.Value.(ResultNumber)
		if !ok {
			panic("unable to case valRes to ResultNumber")
		}
		switch i {
		case 0:
			intValue, _ := val.Abs(val.Decimal).Int64()
			recursionParameters.MaxRecurrences = uint(intValue)
		case 1:
			v := newNumber(0)
			v.Set(val.Decimal)
			recursionParameters.LastValue = v
		}
	}
	recursionParameters.StopWhen = c.BoolExpression()
	return recursionParameters
}

func (a *AbacusVisitor) VisitLambdaExpr(c *parser.LambdaExprContext) interface{} {
	lambdaName := c.LAMBDA_VARIABLE().GetText()

	lambda, found := a.lambdas[lambdaName]
	if !found {
		return NewResult(newNumber(0))
	}

	parameters := ResultTuple{}
	if c.Tuple() != nil {
		valuesRes := c.Tuple().Accept(a).(*Result)
		if hasErrors(valuesRes) {
			return valuesRes
		}
		a.convertTupleResult(valuesRes)
		tuple, ok := valuesRes.Value.(ResultTuple)
		if !ok {
			panic("unable to cast valuesRes to ResultTuple")
		}
		parameters = tuple
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

	if c.RecursionParameters() != nil {
		lambda.parameters = recursionParameters
	} else if !inLambda {
		lambda.parameters = recursionParameters
	}

	if inLambda {
		// Recurrs
		if lambdaName == nested {
			if lambda.parameters.MaxRecurrences == 0 {
				return NewResult(nil).WithErrors(nil, "recursion is disabled")
			}
			recurrences, ok := uint(0), false

			if recurrences, ok = a.lambdaRecursion[lambdaName]; !ok {
				recurrences = 1
				a.lambdaRecursion[lambdaName] = 1
			}
			if _, ok = a.lambdaRecursionStack[lambdaName]; !ok {
				a.lambdaRecursionStack[lambdaName] = 1
			}
			shouldStop := false
			if lambda.parameters.StopWhen != nil {
				conditionRes := lambda.parameters.StopWhen.Accept(a).(*Result)
				if hasErrors(conditionRes) {
					return conditionRes
				}
				condition, ok := conditionRes.Value.(ResultBool)
				if !ok {
					panic("unable to cast conditionRes to ResultBool")
				}
				shouldStop = bool(condition)
			}
			if shouldStop {
				v := newNumber(0)
				v.Set(lambda.parameters.LastValue.Decimal)
				return NewResult(v)
			}
			if recurrences >= lambda.parameters.MaxRecurrences {
				v := newNumber(0)
				v.Set(lambda.parameters.LastValue.Decimal)
				return NewResult(v)
			}
			a.lambdaRecursion[lambdaName]++

		}
	}
	if _, ok := a.lambdaRecursionStack[lambdaName]; !ok {
		a.lambdaRecursionStack[lambdaName] = 1
	} else {
		a.lambdaRecursionStack[lambdaName]++
	}

	//log.Printf("[%s] %v %v\n", lambdaName, inLambda, parameters)

	switch val := lambda.ctx.Lambda().(type) {
	case *parser.VariablesLambdaContext:
		variablesRes := val.VariablesTuple().Accept(a).(*Result)
		if hasErrors(variablesRes) {
			return variablesRes
		}
		a.convertVariablesTupleResult(variablesRes)

		count := variablesRes.Length()
		s := ""
		if count > 1 {
			s = "s"
		}
		if len(parameters) < count {
			return NewResult(nil).WithErrors(nil, "expected "+strconv.FormatInt(int64(count), 10)+" parameter"+s)
		}

		variableNames, ok := variablesRes.Value.(ResultVariablesTuple)
		if !ok {
			panic("unable to cast variablesRes to ResultVariablesTuple")
		}

		for i, varName := range variableNames {
			a.lambdaVars[lambdaVarName(lambdaName, varName.String(), a.lambdaRecursionStack[lambdaName])] = parameters[i]
		}
		result := val.Accept(a).(*Result)
		a.lambdaRecursionStack[lambdaName]--

		switch value := result.Value.(type) {
		case ResultTuple:
			if len(value) == 1 {
				v := newNumber(0)
				v.Set(value[0].Decimal)
				return NewResult(v)
			}
		}
		//log.Printf("[%s] %+v %+v\n", lambdaName, r, parameters)
		return result
	case *parser.NullArityLambdaContext:
		result := val.Accept(a).(*Result)
		a.lambdaRecursionStack[lambdaName]--
		//log.Printf("[%s] %+v %+v\n", lambdaName, r, parameters)

		switch value := result.Value.(type) {
		case ResultTuple:
			if len(value) == 1 {
				v := newNumber(0)
				v.Set(value[0].Decimal)
				NewResult(v)
			}
		}
		return result
	}
	return NewResult(newNumber(0))
}

func (a *AbacusVisitor) VisitVariable(c *parser.VariableContext) interface{} {
	var value ResultNumber
	ok := false

	name := c.VARIABLE().GetText()

	inLambda, lambdaName := a.checkParentCtxForLambda(c.GetParent())
	if inLambda {
		if value, ok = a.lambdaVars[lambdaVarName(lambdaName, name, a.lambdaRecursionStack[lambdaName])]; ok {
			return NewResult(value)
		}

		if value, ok = a.vars[name]; ok {
			return NewResult(value)
		}
	} else {
		if value, ok = a.vars[name]; ok {
			return NewResult(value)
		}
	}
	return NewResult(newNumber(0))
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

func hasErrors(r *Result) bool {
	return len(r.Errors) != 0
}
