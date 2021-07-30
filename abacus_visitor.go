package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/cockroachdb/apd"
	"github.com/viktordanov/abacus/parser"
	"math"
	"strconv"
	"strings"
)

var (
	logCache   map[string]Number
	decimalCtx *apd.Context
	PI         Number
	PHI        Number
	E          Number
)

func init() {
	logCache = make(map[string]Number)
	decimalCtx = apd.BaseContext.WithPrecision(arguments.Precision)
	cachedLog(Number{apd.New(2, 0)})
	cachedLog(Number{apd.New(10, 0)})

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

	PI = Number{pi}
	PHI = Number{phi}
	E = Number{e}
}

func newNumber(f float64) Number {
	res := apd.New(0, 0)
	res.SetFloat64(f)
	return Number{res}
}

func cachedLog(n Number) Number {
	out := newNumber(0)

	var r Number
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

type LambdaDeclaration struct {
	ctx       *parser.LambdaDeclarationContext
	arguments LambdaArguments
	argSet    map[string]bool
}

type Parameter struct {
	Name  string
	Value interface{}
}

type RecursionParameters struct {
	MaxRecurrences uint
	LastValue      parser.IExpressionContext
	StopWhen       parser.IBoolExpressionContext
	Memoize        bool
}

func NewRecursionParameters() *RecursionParameters {
	return &RecursionParameters{MaxRecurrences: 0, LastValue: nil, StopWhen: nil, Memoize: false}
}

type CalledLambda struct {
	ctx       *parser.LambdaDeclarationContext
	parent    *CalledLambda
	children  []*CalledLambda
	arguments map[string]interface{}
	nested    bool
	name      string
}

func (l *CalledLambda) Signature() string {
	var b strings.Builder
	b.WriteString(l.name)
	b.WriteRune('_')
	for _, i := range l.arguments {
		switch val := i.(type) {
		case Number:
			b.WriteString(val.String())
		case String:
			b.WriteString(val.String())
		}
		b.WriteRune('_')
	}
	return b.String()
}

type LambdaCallStack struct {
	root      *CalledLambda
	trace     []*CalledLambda
	invokes   map[string]uint
	recursion map[string]*RecursionParameters
	memoized  map[string]*Result
}

func (s *LambdaCallStack) Peek() *CalledLambda {
	if len(s.trace) == 0 {
		return nil
	}
	return s.trace[len(s.trace)-1]
}

func (s *LambdaCallStack) Pop() *CalledLambda {
	if len(s.trace) == 0 {
		return nil
	}
	el := s.Peek()
	s.trace = s.trace[:len(s.trace)-1]
	return el
}

func (s *LambdaCallStack) Push(l *CalledLambda) {
	s.trace = append(s.trace, l)
}

func (s *LambdaCallStack) IsRecurring(lambdaName string) bool {
	for i := len(s.trace) - 1; i >= 0; i-- {
		if s.trace[i].name == lambdaName {
			return true
		}
	}

	return false
}

type AbacusVisitor struct {
	antlr.ParseTreeVisitor
	variables          map[string]Number
	lambdaDeclarations map[string]*LambdaDeclaration
	lambdaCallStack    *LambdaCallStack
	decimalCtx         *apd.Context
}

func NewAbacusVisitor() *AbacusVisitor {
	return &AbacusVisitor{
		ParseTreeVisitor:   &parser.BaseAbacusVisitor{},
		variables:          make(map[string]Number),
		lambdaDeclarations: make(map[string]*LambdaDeclaration),
		lambdaCallStack:    &LambdaCallStack{root: nil, invokes: map[string]uint{}, trace: []*CalledLambda{}, recursion: map[string]*RecursionParameters{}, memoized: map[string]*Result{}},
		decimalCtx:         apd.BaseContext.WithPrecision(arguments.Precision),
	}
}

func (a *AbacusVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.RootContext:
		a.lambdaCallStack.root = nil
		a.lambdaCallStack.trace = []*CalledLambda{}
		a.lambdaCallStack.invokes = map[string]uint{}
		a.lambdaCallStack.recursion = map[string]*RecursionParameters{}
		a.lambdaCallStack.memoized = map[string]*Result{}
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
	if c.LAMBDA_VARIABLE() != nil {
		if declaration, ok := a.lambdaDeclarations[c.LAMBDA_VARIABLE().GetText()]; ok {
			return NewResult(LambdaAssignment(declaration.ctx.GetText()))

		}
		return NewResult(nil).WithErrors(nil, "undefined lambda "+c.LAMBDA_VARIABLE().GetText())

	}
	return NewResult(nil).WithErrors(nil, "syntax error")
}

func (a *AbacusVisitor) visitTupleTail(c parser.ITupleContext, resultTuple *Tuple) {
	ctx, ok := c.(*parser.TupleContext)
	if !ok || ctx == nil {
		return
	}
	val := ctx.Expression().Accept(a).(*Result)

	switch v := val.Value.(type) {
	case Number:
		*resultTuple = append(*resultTuple, v)
	case Tuple:
		*resultTuple = append(*resultTuple, v...)
	}

	a.visitTupleTail(ctx.Tuple(), resultTuple)
}

func (a *AbacusVisitor) VisitTuple(c *parser.TupleContext) interface{} {
	if c.Tuple() == nil {
		return c.Expression().Accept(a)
	}

	evaledTuple := Tuple{}
	val := c.Expression().Accept(a).(*Result)

	switch v := val.Value.(type) {
	case Number:
		evaledTuple = append(evaledTuple, v)
	case Tuple:
		evaledTuple = append(evaledTuple, v...)
	}
	a.visitTupleTail(c.Tuple(), &evaledTuple)

	return NewResult(evaledTuple)
}

func (a *AbacusVisitor) visitMixedTupleTail(c parser.IMixedTupleContext, resultMixedTuple *MixedTuple) {
	ctx, ok := c.(*parser.MixedTupleContext)
	if !ok || ctx == nil {
		return
	}
	val := NewResult(nil)
	if ctx.Expression() != nil {
		val = ctx.Expression().Accept(a).(*Result)
	} else if ctx.LAMBDA_VARIABLE() != nil {
		val.Value = String(ctx.LAMBDA_VARIABLE().GetText())
	}

	switch v := val.Value.(type) {
	case Number:
		*resultMixedTuple = append(*resultMixedTuple, v)
	case String:
		*resultMixedTuple = append(*resultMixedTuple, v)
	case MixedTuple:
		*resultMixedTuple = append(*resultMixedTuple, v...)
	}

	a.visitMixedTupleTail(ctx.MixedTuple(), resultMixedTuple)
}

func (a *AbacusVisitor) VisitMixedTuple(c *parser.MixedTupleContext) interface{} {
	if c.MixedTuple() == nil {
		val := NewResult(nil)
		if c.Expression() != nil {
			val = c.Expression().Accept(a).(*Result)
		} else if c.LAMBDA_VARIABLE() != nil {
			val.Value = String(c.LAMBDA_VARIABLE().GetText())
		}
		return val
	}

	evaledMixedTuple := MixedTuple{}
	val := NewResult(nil)
	if c.Expression() != nil {
		val = c.Expression().Accept(a).(*Result)
	} else if c.LAMBDA_VARIABLE() != nil {
		val.Value = String(c.LAMBDA_VARIABLE().GetText())
	}

	switch v := val.Value.(type) {
	case Number:
		evaledMixedTuple = append(evaledMixedTuple, v)
	case String:
		evaledMixedTuple = append(evaledMixedTuple, v)
	case MixedTuple:
		evaledMixedTuple = append(evaledMixedTuple, v...)
	}
	a.visitMixedTupleTail(c.MixedTuple(), &evaledMixedTuple)

	return NewResult(evaledMixedTuple)
}

func (a *AbacusVisitor) visitVariableTupleTail(c parser.IVariablesTupleContext, resultTuple *VariablesTuple) {
	ctx, ok := c.(*parser.VariablesTupleContext)
	if !ok || ctx == nil {
		return
	}
	val := ctx.VARIABLE().GetText()
	*resultTuple = append(*resultTuple, String(val))
	a.visitVariableTupleTail(ctx.VariablesTuple(), resultTuple)
}

func (a *AbacusVisitor) VisitVariablesTuple(c *parser.VariablesTupleContext) interface{} {
	if c.VariablesTuple() == nil {
		return NewResult(String(c.VARIABLE().GetText()))
	}

	evaledTuple := VariablesTuple{}

	val := c.VARIABLE().GetText()
	evaledTuple = append(evaledTuple, String(val))
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
	case String:
		result.Value = VariablesTuple{val}
	}
}

func (a *AbacusVisitor) convertLambdaArgumentsResult(result *Result) {
	switch val := result.Value.(type) {
	case String:
		result.Value = LambdaArguments{val}
	}
}

func (a *AbacusVisitor) convertTupleResult(result *Result) {
	switch val := result.Value.(type) {
	case Number:
		result.Value = Tuple{val}
	}
}

func (a *AbacusVisitor) convertMixedTupleResult(result *Result) {
	switch val := result.Value.(type) {
	case Number:
		result.Value = MixedTuple{val}
	case String:
		result.Value = MixedTuple{val}
	}
}

func (a *AbacusVisitor) VisitLambdaArguments(c *parser.LambdaArgumentsContext) interface{} {
	if c.LambdaArguments() == nil {
		val := ""
		if c.VARIABLE() != nil {
			val = c.VARIABLE().GetText()
		} else if c.LAMBDA_VARIABLE() != nil {
			val = c.LAMBDA_VARIABLE().GetText()
		}
		return NewResult(String(val))
	}

	evaledArgs := LambdaArguments{}

	val := ""
	if c.VARIABLE() != nil {
		val = c.VARIABLE().GetText()
	} else if c.LAMBDA_VARIABLE() != nil {
		val = c.LAMBDA_VARIABLE().GetText()
	}
	evaledArgs = append(evaledArgs, String(val))
	a.visitLambdaArgsTail(c.LambdaArguments(), &evaledArgs)

	foundVars := make(map[string]bool)
	for _, variable := range evaledArgs {
		if _, ok := foundVars[variable.String()]; !ok {
			foundVars[variable.String()] = true
		} else {
			firstChar := variable[0]
			if firstChar >= 'A' && firstChar <= 'Z' {
				return NewResult(evaledArgs).WithErrors(nil, "duplicate lambda name \""+variable.String()+"\"")
			} else {
				return NewResult(evaledArgs).WithErrors(nil, "duplicate variable name \""+variable.String()+"\"")
			}
		}
	}

	return NewResult(evaledArgs)
}

func (a *AbacusVisitor) visitLambdaArgsTail(c parser.ILambdaArgumentsContext, resultArgs *LambdaArguments) {
	ctx, ok := c.(*parser.LambdaArgumentsContext)
	if !ok || ctx == nil {
		return
	}
	val := ""
	if ctx.VARIABLE() != nil {
		val = ctx.VARIABLE().GetText()
	} else if ctx.LAMBDA_VARIABLE() != nil {
		val = ctx.LAMBDA_VARIABLE().GetText()
	}
	*resultArgs = append(*resultArgs, String(val))
	a.visitLambdaArgsTail(ctx.LambdaArguments(), resultArgs)
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

	variables, ok := variablesRes.Value.(VariablesTuple)
	if !ok {
		panic("unable to convert variables result to its raw type")
	}
	values, ok := valuesRes.Value.(Tuple)
	if !ok {
		panic("unable to convert values result to its raw type")
	}

	for i, variable := range variables {
		a.variables[variable.String()] = values[i]
	}
	return NewResult(Assignment(values))
}

func (a *AbacusVisitor) VisitLambdaDeclaration(c *parser.LambdaDeclarationContext) interface{} {
	lambdaName := c.LAMBDA_VARIABLE().GetText()
	lambda := c.Lambda()

	// Check if 1) multiple variables && duped variables
	multipleVarLambda, ok := lambda.(*parser.VariablesLambdaContext)
	arguments := LambdaArguments{}
	argMap := map[string]bool{}
	if ok {
		argsRes := multipleVarLambda.LambdaArguments().Accept(a).(*Result)
		if hasErrors(argsRes) {
			return argsRes
		}
		a.convertLambdaArgumentsResult(argsRes)

		arguments, ok = argsRes.Value.(LambdaArguments)
		if !ok {
			panic("unable to cast argsRes to LambdaArguments")
		}
		for _, arg := range arguments {
			argMap[string(arg)] = true
		}

	}

	a.lambdaDeclarations[lambdaName] = &LambdaDeclaration{
		ctx:       c,
		argSet:    argMap,
		arguments: arguments,
	}

	formattedLambda := lambdaName + " = " + lambda.GetText()
	return NewResult(LambdaAssignment(formattedLambda))
}

func (a *AbacusVisitor) VisitEqualComparison(c *parser.EqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(Number)
	if !ok {
		panic("unable to cast right to Number")
	}
	rightVal, ok := right.Value.(Number)
	if !ok {
		panic("unable to cast left to Number")
	}

	return NewResult(Bool(leftVal.Cmp(rightVal.Decimal) == 0))
}

func (a *AbacusVisitor) VisitLessComparison(c *parser.LessComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(Number)
	if !ok {
		panic("unable to cast right to Number")
	}
	rightVal, ok := right.Value.(Number)
	if !ok {
		panic("unable to cast left to Number")
	}

	return NewResult(Bool(leftVal.Cmp(rightVal.Decimal) == -1))
}

func (a *AbacusVisitor) VisitGreaterComparison(c *parser.GreaterComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(Number)
	if !ok {
		panic("unable to cast right to Number")
	}
	rightVal, ok := right.Value.(Number)
	if !ok {
		panic("unable to cast left to Number")
	}

	return NewResult(Bool(leftVal.Cmp(rightVal.Decimal) == 1))
}

func (a *AbacusVisitor) VisitLessOrEqualComparison(c *parser.LessOrEqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(Number)
	if !ok {
		panic("unable to cast right to Number")
	}
	rightVal, ok := right.Value.(Number)
	if !ok {
		panic("unable to cast left to Number")
	}

	return NewResult(Bool(leftVal.Cmp(rightVal.Decimal) <= 0))
}

func (a *AbacusVisitor) VisitGreaterOrEqualComparison(c *parser.GreaterOrEqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(a).(*Result)
	right := c.Expression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(Number)
	if !ok {
		panic("unable to cast right to Number")
	}
	rightVal, ok := right.Value.(Number)
	if !ok {
		panic("unable to cast left to Number")
	}

	return NewResult(Bool(leftVal.Cmp(rightVal.Decimal) >= 0))
}

func (a *AbacusVisitor) VisitAndOrXor(c *parser.AndOrXorContext) interface{} {
	left := c.BoolExpression(0).Accept(a).(*Result)
	right := c.BoolExpression(1).Accept(a).(*Result)

	if hasErrors(left) || hasErrors(right) {
		return left.WithErrors(right)
	}

	leftVal, ok := left.Value.(Bool)
	if !ok {
		panic("unable to cast right to Bool")
	}
	rightVal, ok := right.Value.(Bool)
	if !ok {
		panic("unable to cast left to Bool")
	}

	op := c.GetOp().GetTokenType()
	result := Bool(false)

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
	val, ok := valRes.Value.(Bool)
	if !ok {
		panic("unable to cast right to Bool")
	}
	return NewResult(!val)
}

func (a *AbacusVisitor) VisitParenthesesBoolean(c *parser.ParenthesesBooleanContext) interface{} {
	return c.BoolExpression().Accept(a)
}

func (a *AbacusVisitor) VisitBoolAtom(c *parser.BoolAtomContext) interface{} {
	val := c.GetText()
	if val == "true" {
		return NewResult(Bool(true))
	}
	return NewResult(Bool(false))
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

	leftVal, ok := left.Value.(Number)
	if !ok {
		panic("unable to cast right to Number")
	}
	rightVal, ok := right.Value.(Number)
	if !ok {
		panic("unable to cast left to Number")
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

	leftVal, ok := left.Value.(Number)
	if !ok {
		panic("unable to cast right to Number")
	}
	rightVal, ok := right.Value.(Number)
	if !ok {
		panic("unable to cast left to Number")
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

	leftVal, ok := left.Value.(Number)
	if !ok {
		panic("unable to cast right to Number")
	}
	rightVal, ok := right.Value.(Number)
	if !ok {
		panic("unable to cast left to Number")
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

	leftVal, ok := left.Value.(Number)
	if !ok {
		panic("unable to cast right to Number")
	}
	rightVal, ok := right.Value.(Number)
	if !ok {
		panic("unable to cast left to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
	}

	res := newNumber(0)

	res.Set(val.Decimal)
	sign := c.Sign().Accept(a).(rune)
	if sign == '-' {
		res.Negative = !res.Negative
	}

	return NewResult(res)
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
	}

	v := newNumber(0)
	a.decimalCtx.Exp(v.Decimal, val.Decimal)
	return NewResult(v)
}

func (a *AbacusVisitor) VisitSignFunction(c *parser.SignFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
	}

	v := newNumber(0)
	v = newNumber(float64(val.Decimal.Cmp(v.Decimal)))
	return NewResult(v)
}

func (a *AbacusVisitor) VisitAbsFunction(c *parser.AbsFunctionContext) interface{} {
	valRes := c.Expression().Accept(a).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
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

	leftVal, ok := left.Value.(Number)
	if !ok {
		panic("unable to cast right to Number")
	}
	rightVal, ok := right.Value.(Number)
	if !ok {
		panic("unable to cast left to Number")
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

	leftVal, ok := left.Value.(Number)
	if !ok {
		panic("unable to cast right to Number")
	}
	rightVal, ok := right.Value.(Number)
	if !ok {
		panic("unable to cast left to Number")
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

	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
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

	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
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

	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
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

	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
	}

	argRes := c.Expression().Accept(a).(*Result)
	if hasErrors(argRes) {
		return argRes
	}

	arg, ok := argRes.Value.(Number)
	if !ok {
		panic("unable to cast argRes to Number")
	}
	intValue, _ := arg.Int64()

	newTuple := Tuple{}
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

	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
	}
	arg := tuple[len(tuple)-1]
	intValue, _ := arg.Int64()

	newTuple := Tuple{}
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

	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
	}

	newTuple := Tuple{}
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

	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
	}

	argRes := c.Expression().Accept(a).(*Result)
	if hasErrors(argRes) {
		return argRes
	}
	arg, ok := argRes.Value.(Number)
	if !ok {
		panic("unable to cast argRes to Number")
	}
	intValue1, _ := arg.Int64()

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
	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast valRes to Number")
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
	return NewResult(Number{out})
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
	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
	}

	return NewResult(tuple)
}

func (a *AbacusVisitor) VisitNullArityLambda(c *parser.NullArityLambdaContext) interface{} {
	tupleRes := c.Tuple().Accept(a).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	a.convertTupleResult(tupleRes)
	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
	}

	return NewResult(tuple)
}

func (a *AbacusVisitor) VisitParameter(c *parser.ParameterContext) interface{} {
	paramName := c.VARIABLE().GetText()

	if paramName == "last" {
		return Parameter{
			Name:  paramName,
			Value: c.Expression(),
		}
	}

	if c.Expression() != nil {
		expRes := c.Expression().Accept(a).(*Result)
		if hasErrors(expRes) {
			return expRes
		}
		number := expRes.Value.(Number)
		return Parameter{
			Name:  paramName,
			Value: number,
		}
	} else if c.BoolExpression() != nil {
		return Parameter{
			Name:  paramName,
			Value: c.BoolExpression(),
		}
	}

	// shouldn't be reached
	return nil
}

// TODO: Errors aren't handled
func (a *AbacusVisitor) VisitRecursionParameters(c *parser.RecursionParametersContext) interface{} {
	recursionParameters := NewRecursionParameters()

	for i := 0; i < len(c.AllParameter()); i++ {
		valRes := c.Parameter(i).Accept(a)
		switch val := valRes.(type) {
		case *Result:
			return val
		case Parameter:
			switch parameterValue := val.Value.(type) {
			case Number:
				if val.Name == "rec" {
					intValue, _ := parameterValue.Abs(parameterValue.Decimal).Int64()
					recursionParameters.MaxRecurrences = uint(intValue)
				}
			case parser.IExpressionContext:
				if val.Name == "last" {
					recursionParameters.LastValue = parameterValue
				}
			case parser.IBoolExpressionContext:
				if val.Name == "stop" {
					recursionParameters.StopWhen = parameterValue
				} else if val.Name == "mem" {
					res := parameterValue.Accept(a).(*Result)
					if !hasErrors(res) {
						recursionParameters.Memoize = bool(res.Value.(Bool))
					}
				}
			}
		}
	}

	return recursionParameters
}

func (a *AbacusVisitor) VisitLambdaExpr(c *parser.LambdaExprContext) interface{} {
	stack := a.lambdaCallStack
	lambda := &CalledLambda{
		arguments: nil,
		name:      c.LAMBDA_VARIABLE().GetText(),
		children:  []*CalledLambda{},
		nested:    false,
	}

getDeclaration:
	declaration, found := a.lambdaDeclarations[lambda.name]

	if found {
		lambda.ctx = declaration.ctx
	}

	// Handle local "composed" lambdas, e.g.    E = x,Fn -> Fn(x)  // Fn is local
	parent := stack.Peek()
	isLocal := false

	if parent != nil {
		_, isLocal = parent.arguments[lambda.name]
	}
	if isLocal {
		var parent *CalledLambda
		for {
			if parent == nil {
				parent = stack.Peek()
			} else {
				parent = parent.parent
			}
			if parent != nil {
				if value, ok := parent.arguments[lambda.name]; ok {
					lambda.name = value.(String).String()

					isLocalToGrandparent := false
					if parent.parent != nil {
						_, isLocalToGrandparent = parent.parent.arguments[lambda.name]
					}
					if _, ok := a.lambdaDeclarations[lambda.name]; ok && !isLocalToGrandparent {
						goto getDeclaration
					}
				}
			} else {
				break
			}
		}
	}
	if !found {
		if arguments.Strict {
			return NewResult(nil).WithErrors(nil, "undefined lambda "+lambda.name)
		}
		return NewResult(newNumber(0))
	}

	// Init arguments
	if c.MixedTuple() != nil {
		valuesRes := c.MixedTuple().Accept(a).(*Result)
		if hasErrors(valuesRes) {
			return valuesRes
		}
		a.convertMixedTupleResult(valuesRes)
		tuple, ok := valuesRes.Value.(MixedTuple)
		if !ok {
			panic("unable to cast valuesRes to MixedTuple")
		}

		// Check arity
		if len(tuple) != len(declaration.arguments) {
			count := len(declaration.arguments)
			s := ""
			if count > 1 {
				s = "s"
			}
			if len(tuple) < count {
				return NewResult(nil).WithErrors(nil, "expected "+strconv.FormatInt(int64(count), 10)+" parameter"+s)
			}
		}

		// Create argName -> value map
		lambda.arguments = map[string]interface{}{}
		for i, argument := range declaration.arguments {
			// Check type of received parameter and check against expected argument
			if isLambdaName(argument.String()) {
				if _, ok := tuple[i].(String); !ok {
					return NewResult(nil).WithErrors(nil, "["+lambda.name+"] expected lambda for parameter "+strconv.Itoa(i+1)+", got number")
				}
			} else {
				if _, ok := tuple[i].(Number); !ok {
					return NewResult(nil).WithErrors(nil, "["+lambda.name+"] expected number for parameter "+strconv.Itoa(i+1)+", got lambda")
				}
			}
			lambda.arguments[argument.String()] = tuple[i]
		}
	}

	// Init recursion parameters
	if _, ok := stack.recursion[lambda.name]; !ok {
		if c.RecursionParameters() != nil {
			stack.recursion[lambda.name] = c.RecursionParameters().Accept(a).(*RecursionParameters)
		} else {
			stack.recursion[lambda.name] = NewRecursionParameters()
		}
	}

	// Init leaf
	if stack.root == nil {
		stack.root = lambda
	}
	parent = stack.Peek()
	if parent != nil {
		lambda.parent = parent
		parent.children = append(parent.children, lambda)
	}

	recurring := stack.IsRecurring(lambda.name)
	stack.Push(lambda)

	// Count times lambda has been invoked
	if _, ok := stack.invokes[lambda.name]; !ok {
		stack.invokes[lambda.name] = 1
	} else {
		stack.invokes[lambda.name]++
	}

	// Check if memoization is enabled and if a value exists to prevent branching
	if stack.recursion[lambda.name].Memoize {
		cachedResult, ok := stack.memoized[lambda.Signature()]
		if ok {
			stack.Pop()
			return cachedResult
		}
	}

	// Calculate and the stop condition if present
	shouldStop := false
	if stack.recursion[lambda.name].StopWhen != nil {
		conditionRes := stack.recursion[lambda.name].StopWhen.Accept(a).(*Result)
		if hasErrors(conditionRes) {
			return conditionRes
		}
		condition, ok := conditionRes.Value.(Bool)
		if !ok {
			panic("unable to cast conditionRes to Bool")
		}
		shouldStop = bool(condition)
	}

	if shouldStop {
		v := newNumber(0)
		if stack.recursion[lambda.name].LastValue == nil {
			stack.Pop()
			return NewResult(v)
		}
		// Eval last value
		lastValRes := stack.recursion[lambda.name].LastValue.Accept(a).(*Result)
		if hasErrors(lastValRes) {
			return lastValRes
		}
		lastVal, ok := lastValRes.Value.(Number)
		if !ok {
			panic("unable to cast lastValRes to Number")
		}

		v.Set(lastVal.Decimal)
		stack.Pop()
		return NewResult(v)
	}

	// Handle recursion
	if recurring {
		recParameters := stack.recursion[lambda.name]
		invokes := stack.invokes[lambda.name]

		if recParameters.MaxRecurrences == 0 {
			return NewResult(nil).WithErrors(nil, "recursion is disabled")
		}

		if shouldStop {
			v := newNumber(0)
			if stack.recursion[lambda.name].LastValue == nil {
				stack.Pop()
				return NewResult(v)
			}
			// Eval last value
			lastValRes := stack.recursion[lambda.name].LastValue.Accept(a).(*Result)
			if hasErrors(lastValRes) {
				return lastValRes
			}
			lastVal, ok := lastValRes.Value.(Number)
			if !ok {
				panic("unable to cast lastValRes to Number")
			}

			v.Set(lastVal.Decimal)
			stack.Pop()
			return NewResult(v)
		}
		if invokes > recParameters.MaxRecurrences {
			v := newNumber(0)
			if stack.recursion[lambda.name].LastValue == nil {
				stack.Pop()
				return NewResult(v)
			}
			// Eval last value
			lastValRes := stack.recursion[lambda.name].LastValue.Accept(a).(*Result)
			if hasErrors(lastValRes) {
				return lastValRes
			}
			lastVal, ok := lastValRes.Value.(Number)
			if !ok {
				panic("unable to cast lastValRes to Number")
			}

			v.Set(lastVal.Decimal)
			stack.Pop()
			return NewResult(v)
		}
	}

	// Evaluate lambda
	declaredLambda := lambda.ctx.Lambda()
	result := declaredLambda.Accept(a).(*Result)
	stack.Pop()

	switch value := result.Value.(type) {
	case Tuple:
		if len(value) == 1 {
			v := newNumber(0)
			v.Set(value[0].Decimal)
			result = NewResult(v)
			if stack.recursion[lambda.name].Memoize {
				stack.memoized[lambda.Signature()] = result
			}
			return result
		}
	}

	if stack.recursion[lambda.name].Memoize {
		stack.memoized[lambda.Signature()] = result
	}
	return result
}

func (a *AbacusVisitor) VisitVariable(c *parser.VariableContext) interface{} {
	name := c.VARIABLE().GetText()

	lambda := a.lambdaCallStack.Peek()
	if lambda != nil {
		if value, ok := lambda.arguments[name]; ok {
			return NewResult(value.(Number))
		}
	}

	if value, ok := a.variables[name]; ok {
		return NewResult(value)
	}

	if arguments.Strict {
		return NewResult(nil).WithErrors(nil, "undefined global variable "+name)
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

func hasErrors(r *Result) bool {
	return len(r.Errors) != 0
}

func isLambdaName(name string) bool {
	return name[0] >= 'A' && name[0] <= 'Z'
}
