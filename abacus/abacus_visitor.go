package abacus

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/cockroachdb/apd"
	"github.com/viktordanov/abacus/parser"
	_ "math"
	"strconv"
)

func initConstants(store *ConstantsStore, precision uint32) {
	decimalCtx := apd.BaseContext.WithPrecision(precision)

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

	store.SetConstant(PI, Number{pi})
	store.SetConstant(PHI, Number{phi})
	store.SetConstant(E, Number{e})

	calculateLog(store, decimalCtx, Number{apd.New(2, 0)})
	calculateLog(store, decimalCtx, Number{apd.New(10, 0)})
}

func NewNumber(f float64) Number {
	res := apd.New(0, 0)
	res.SetFloat64(f)
	return Number{res}
}

func calculateLog(store *ConstantsStore, ctx *apd.Context, n Number) Number {
	out := NewNumber(0)

	var r Number
	var ok bool
	if r, ok = store.Cached(n.String()); !ok {
		ctx.Ln(out.Decimal, n.Decimal)
		newValue := NewNumber(0)
		newValue.Decimal.Set(out.Decimal)
		store.Cache(n.String(), newValue)
	} else {
		out.Set(r.Decimal)
	}
	return out
}

type Visitor struct {
	antlr.ParseTreeVisitor
	variables          map[string]Number
	lambdaDeclarations map[string]*LambdaDeclaration
	lambdaCallStack    *LambdaCallStack
	decimalCtx         *apd.Context
	strict             bool
	ConstantsStore     *ConstantsStore
}

func (v *Visitor) VariableNames() []string {
	var names []string
	for k := range v.variables {
		names = append(names, k)
	}
	return names
}

func (v *Visitor) LambdaNames() []string {
	var names []string
	for k := range v.lambdaDeclarations {
		names = append(names, k)
	}
	return names
}

func (v *Visitor) SetVariable(name string, value Number) {
	v.variables[name] = value
}

func NewAbacusVisitor(precision uint32, strict bool) (*Visitor, error) {
	constantsStore, err := NewConstantsStore()
	if err != nil {
		return nil, err
	}

	initConstants(constantsStore, precision)
	return &Visitor{
		ParseTreeVisitor:   &parser.BaseAbacusVisitor{},
		variables:          make(map[string]Number),
		lambdaDeclarations: make(map[string]*LambdaDeclaration),
		lambdaCallStack:    &LambdaCallStack{root: nil, invokes: map[string]uint{}, trace: []*CalledLambda{}, recursion: map[string]*RecursionParameters{}, memoized: map[string]*Result{}},
		decimalCtx:         apd.BaseContext.WithPrecision(precision),
		strict:             strict,

		ConstantsStore: constantsStore,
	}, err
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.RootContext:
		v.lambdaCallStack.root = nil
		v.lambdaCallStack.trace = []*CalledLambda{}
		v.lambdaCallStack.invokes = map[string]uint{}
		v.lambdaCallStack.recursion = map[string]*RecursionParameters{}
		v.lambdaCallStack.memoized = map[string]*Result{}
		return val.Accept(v)
	}
	return nil
}

func (v *Visitor) VisitRoot(c *parser.RootContext) interface{} {
	if c.Tuple() != nil {
		return c.Tuple().Accept(v)
	}
	if c.Declaration() != nil {
		return c.Declaration().Accept(v)
	}
	if c.BoolExpression() != nil {
		return c.BoolExpression().Accept(v)
	}
	if c.LAMBDA_VARIABLE() != nil {
		if declaration, ok := v.lambdaDeclarations[c.LAMBDA_VARIABLE().GetText()]; ok {
			return NewResult(LambdaAssignment(declaration.ctx.GetText()))

		}
		return NewResult(nil).WithError("undefined lambda " + c.LAMBDA_VARIABLE().GetText())

	}
	return NewResult(nil).WithError("syntax error")
}

func (v *Visitor) visitTupleTail(c parser.ITupleContext, resultTuple *Tuple) {
	ctx, ok := c.(*parser.TupleContext)
	if !ok || ctx == nil {
		return
	}
	val := ctx.Expression().Accept(v).(*Result)

	switch v := val.Value.(type) {
	case Number:
		*resultTuple = append(*resultTuple, v)
	case Tuple:
		*resultTuple = append(*resultTuple, v...)
	}

	v.visitTupleTail(ctx.Tuple(), resultTuple)
}

func (v *Visitor) VisitTuple(c *parser.TupleContext) interface{} {
	if c.Tuple() == nil {
		return c.Expression().Accept(v)
	}

	evaledTuple := Tuple{}
	val := c.Expression().Accept(v).(*Result)

	switch v := val.Value.(type) {
	case Number:
		evaledTuple = append(evaledTuple, v)
	case Tuple:
		evaledTuple = append(evaledTuple, v...)
	}
	v.visitTupleTail(c.Tuple(), &evaledTuple)

	return NewResult(evaledTuple)
}

func (v *Visitor) visitMixedTupleTail(c parser.IMixedTupleContext, resultMixedTuple *MixedTuple) {
	ctx, ok := c.(*parser.MixedTupleContext)
	if !ok || ctx == nil {
		return
	}
	val := NewResult(nil)
	if ctx.Expression() != nil {
		val = ctx.Expression().Accept(v).(*Result)
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

	v.visitMixedTupleTail(ctx.MixedTuple(), resultMixedTuple)
}

func (v *Visitor) VisitMixedTuple(c *parser.MixedTupleContext) interface{} {
	if c.MixedTuple() == nil {
		val := NewResult(nil)
		if c.Expression() != nil {
			val = c.Expression().Accept(v).(*Result)
		} else if c.LAMBDA_VARIABLE() != nil {
			val.Value = String(c.LAMBDA_VARIABLE().GetText())
		}
		return val
	}

	evaledMixedTuple := MixedTuple{}
	val := NewResult(nil)
	if c.Expression() != nil {
		val = c.Expression().Accept(v).(*Result)
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
	v.visitMixedTupleTail(c.MixedTuple(), &evaledMixedTuple)

	return NewResult(evaledMixedTuple)
}

func (v *Visitor) visitVariableTupleTail(c parser.IVariablesTupleContext, resultTuple *VariablesTuple) {
	ctx, ok := c.(*parser.VariablesTupleContext)
	if !ok || ctx == nil {
		return
	}
	val := ctx.VARIABLE().GetText()
	*resultTuple = append(*resultTuple, String(val))
	v.visitVariableTupleTail(ctx.VariablesTuple(), resultTuple)
}

func (v *Visitor) VisitVariablesTuple(c *parser.VariablesTupleContext) interface{} {
	if c.VariablesTuple() == nil {
		return NewResult(String(c.VARIABLE().GetText()))
	}

	evaledTuple := VariablesTuple{}

	val := c.VARIABLE().GetText()
	evaledTuple = append(evaledTuple, String(val))
	v.visitVariableTupleTail(c.VariablesTuple(), &evaledTuple)

	foundVars := make(map[string]bool)
	for _, variable := range evaledTuple {
		if _, ok := foundVars[variable.String()]; !ok {
			foundVars[variable.String()] = true
		} else {
			return NewResult(evaledTuple).WithError("duplicate variable name \"" + variable.String() + "\"")
		}
	}

	return NewResult(evaledTuple)
}

func (v *Visitor) convertVariablesTupleResult(result *Result) {
	switch val := result.Value.(type) {
	case String:
		result.Value = VariablesTuple{val}
	}
}

func (v *Visitor) convertLambdaArgumentsResult(result *Result) {
	switch val := result.Value.(type) {
	case String:
		result.Value = LambdaArguments{val}
	}
}

func (v *Visitor) convertTupleResult(result *Result) {
	switch val := result.Value.(type) {
	case Number:
		result.Value = Tuple{val}
	}
}

func (v *Visitor) convertMixedTupleResult(result *Result) {
	switch val := result.Value.(type) {
	case Number:
		result.Value = MixedTuple{val}
	case String:
		result.Value = MixedTuple{val}
	}
}

func (v *Visitor) VisitLambdaArguments(c *parser.LambdaArgumentsContext) interface{} {
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
	v.visitLambdaArgsTail(c.LambdaArguments(), &evaledArgs)

	foundVars := make(map[string]bool)
	for _, variable := range evaledArgs {
		if _, ok := foundVars[variable.String()]; !ok {
			foundVars[variable.String()] = true
		} else {
			firstChar := variable[0]
			if firstChar >= 'A' && firstChar <= 'Z' {
				return NewResult(evaledArgs).WithError("duplicate lambda name \"" + variable.String() + "\"")
			} else {
				return NewResult(evaledArgs).WithError("duplicate variable name \"" + variable.String() + "\"")
			}
		}
	}

	return NewResult(evaledArgs)
}

func (v *Visitor) visitLambdaArgsTail(c parser.ILambdaArgumentsContext, resultArgs *LambdaArguments) {
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
	v.visitLambdaArgsTail(ctx.LambdaArguments(), resultArgs)
}

func (v *Visitor) VisitVariableDeclaration(c *parser.VariableDeclarationContext) interface{} {
	variablesRes := c.VariablesTuple().Accept(v).(*Result)
	if hasErrors(variablesRes) {
		return variablesRes
	}
	v.convertVariablesTupleResult(variablesRes)

	valuesRes := c.Tuple().Accept(v).(*Result)
	if hasErrors(variablesRes) {
		return variablesRes
	}
	v.convertTupleResult(valuesRes)

	if variablesRes.Length() != valuesRes.Length() {
		return variablesRes.WithError("wrong number of valuesRes " + strconv.FormatInt(int64(valuesRes.Length()), 10) + "; expected " + strconv.FormatInt(int64(variablesRes.Length()), 10))
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
		v.variables[variable.String()] = values[i]
	}
	return NewResult(Assignment(values))
}

func (v *Visitor) VisitLambdaDeclaration(c *parser.LambdaDeclarationContext) interface{} {
	lambdaName := c.LAMBDA_VARIABLE().GetText()
	lambda := c.Lambda()

	// Check if 1) multiple variables && duped variables
	multipleVarLambda, ok := lambda.(*parser.VariablesLambdaContext)
	arguments := LambdaArguments{}
	argMap := map[string]bool{}
	if ok {
		argsRes := multipleVarLambda.LambdaArguments().Accept(v).(*Result)
		if hasErrors(argsRes) {
			return argsRes
		}
		v.convertLambdaArgumentsResult(argsRes)

		arguments, ok = argsRes.Value.(LambdaArguments)
		if !ok {
			panic("unable to cast argsRes to LambdaArguments")
		}
		for _, arg := range arguments {
			argMap[string(arg)] = true
		}

	}

	v.lambdaDeclarations[lambdaName] = &LambdaDeclaration{
		ctx:       c,
		argSet:    argMap,
		arguments: arguments,
	}

	formattedLambda := lambdaName + " = " + lambda.GetText()
	return NewResult(LambdaAssignment(formattedLambda))
}

func (v *Visitor) VisitEqualComparison(c *parser.EqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(v).(*Result)
	right := c.Expression(1).Accept(v).(*Result)

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

func (v *Visitor) VisitLessComparison(c *parser.LessComparisonContext) interface{} {
	left := c.Expression(0).Accept(v).(*Result)
	right := c.Expression(1).Accept(v).(*Result)

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

func (v *Visitor) VisitGreaterComparison(c *parser.GreaterComparisonContext) interface{} {
	left := c.Expression(0).Accept(v).(*Result)
	right := c.Expression(1).Accept(v).(*Result)

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

func (v *Visitor) VisitLessOrEqualComparison(c *parser.LessOrEqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(v).(*Result)
	right := c.Expression(1).Accept(v).(*Result)

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

func (v *Visitor) VisitGreaterOrEqualComparison(c *parser.GreaterOrEqualComparisonContext) interface{} {
	left := c.Expression(0).Accept(v).(*Result)
	right := c.Expression(1).Accept(v).(*Result)

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

func (v *Visitor) VisitAndOrXor(c *parser.AndOrXorContext) interface{} {
	left := c.BoolExpression(0).Accept(v).(*Result)
	right := c.BoolExpression(1).Accept(v).(*Result)

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

func (v *Visitor) VisitNot(c *parser.NotContext) interface{} {
	valRes := c.BoolExpression().Accept(v).(*Result)
	if hasErrors(valRes) {
		return valRes
	}
	val, ok := valRes.Value.(Bool)
	if !ok {
		panic("unable to cast right to Bool")
	}
	return NewResult(!val)
}

func (v *Visitor) VisitParenthesesBoolean(c *parser.ParenthesesBooleanContext) interface{} {
	return c.BoolExpression().Accept(v)
}

func (v *Visitor) VisitBoolAtom(c *parser.BoolAtomContext) interface{} {
	val := c.GetText()
	if val == "true" {
		return NewResult(Bool(true))
	}
	return NewResult(Bool(false))
}

func (v *Visitor) VisitBooleanAtom(c *parser.BooleanAtomContext) interface{} {
	return c.BoolAtom().Accept(v)
}

func (v *Visitor) VisitMulDiv(c *parser.MulDivContext) interface{} {
	left := c.Expression(0).Accept(v).(*Result)
	right := c.Expression(1).Accept(v).(*Result)

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

	res := NewNumber(0)
	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserMUL:
		v.decimalCtx.Mul(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	case parser.AbacusLexerDIV:
		v.decimalCtx.Quo(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	}
	return NewResult(res)
}

func (v *Visitor) VisitAddSub(c *parser.AddSubContext) interface{} {
	left := c.Expression(0).Accept(v).(*Result)
	right := c.Expression(1).Accept(v).(*Result)

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

	res := NewNumber(0)
	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserADD:
		v.decimalCtx.Add(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	case parser.AbacusLexerSUB:
		v.decimalCtx.Sub(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	}
	return NewResult(res)
}

func (v *Visitor) VisitPow(c *parser.PowContext) interface{} {
	left := c.Expression(0).Accept(v).(*Result)
	right := c.Expression(1).Accept(v).(*Result)

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
	res := NewNumber(0)
	v.decimalCtx.Pow(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	return NewResult(res)
}

func (v *Visitor) VisitMod(c *parser.ModContext) interface{} {
	left := c.Expression(0).Accept(v).(*Result)
	right := c.Expression(1).Accept(v).(*Result)

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
	res := NewNumber(0)
	v.decimalCtx.Rem(res.Decimal, leftVal.Decimal, rightVal.Decimal)
	return NewResult(res)
}

func (v *Visitor) VisitSignedExpr(c *parser.SignedExprContext) interface{} {
	valRes := c.Expression().Accept(v).(*Result)
	if hasErrors(valRes) {
		return valRes
	}

	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast val to Number")
	}

	res := NewNumber(0)

	res.Set(val.Decimal)
	sign := c.Sign().Accept(v).(rune)
	if sign == '-' {
		res.Negative = !res.Negative
	}

	return NewResult(res)
}

func (v *Visitor) VisitParentheses(c *parser.ParenthesesContext) interface{} {
	return c.Expression().Accept(v)
}

func (v *Visitor) VisitAtomExpr(c *parser.AtomExprContext) interface{} {
	return c.Atom().Accept(v)
}

func (v *Visitor) VisitFuncExpr(c *parser.FuncExprContext) interface{} {
	return c.Function().Accept(v)
}

func (v *Visitor) VisitFunctionInvocation(c *parser.FunctionInvocationContext) interface{} {
	return v.HandleFunctionInvocation(c)
}

func (v *Visitor) VisitConstant(c *parser.ConstantContext) interface{} {
	switch c.CONSTANT().GetText() {
	case "pi":
		if pi, ok := v.ConstantsStore.Constant(PI); !ok {
			return NewResult(nil).WithError("pi is not defined")
		} else {
			return NewResult(pi)
		}
	case "phi":
		if phi, ok := v.ConstantsStore.Constant(PHI); !ok {
			return NewResult(nil).WithError("phi is not defined")
		} else {
			return NewResult(phi)
		}
	case "e":
		if e, ok := v.ConstantsStore.Constant(E); !ok {
			return NewResult(nil).WithError("e is not defined")
		} else {
			return NewResult(e)
		}
	}
	return NewResult(NewNumber(0))
}

func (v *Visitor) VisitPercent(c *parser.PercentContext) interface{} {
	valRes := c.Expression().Accept(v).(*Result)
	if hasErrors(valRes) {
		return valRes
	}
	val, ok := valRes.Value.(Number)
	if !ok {
		panic("unable to cast valRes to Number")
	}

	v.decimalCtx.Quo(val.Decimal, val.Decimal, NewNumber(100).Decimal)
	return NewResult(val)
}

func (v *Visitor) VisitNumber(c *parser.NumberContext) interface{} {
	numberString := c.SCIENTIFIC_NUMBER().GetText()

	out, _, err := apd.NewFromString(numberString)
	if err != nil {
		panic(err)
	}
	return NewResult(Number{out})
}

func (v *Visitor) VisitPlusSign(c *parser.PlusSignContext) interface{} {
	return '+'
}

func (v *Visitor) VisitMinusSign(c *parser.MinusSignContext) interface{} {
	return '-'
}

func (v *Visitor) VisitVariablesLambda(c *parser.VariablesLambdaContext) interface{} {
	tupleRes := c.Tuple().Accept(v).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	v.convertTupleResult(tupleRes)
	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
	}

	return NewResult(tuple)
}

func (v *Visitor) VisitNullArityLambda(c *parser.NullArityLambdaContext) interface{} {
	tupleRes := c.Tuple().Accept(v).(*Result)
	if hasErrors(tupleRes) {
		return tupleRes
	}
	v.convertTupleResult(tupleRes)
	tuple, ok := tupleRes.Value.(Tuple)
	if !ok {
		panic("unable to cast tupleRes to Tuple")
	}

	return NewResult(tuple)
}

func (v *Visitor) VisitParameter(c *parser.ParameterContext) interface{} {
	paramName := c.VARIABLE().GetText()

	if paramName == "last" {
		return Parameter{
			Name:  paramName,
			Value: c.Expression(),
		}
	}

	if c.Expression() != nil {
		expRes := c.Expression().Accept(v).(*Result)
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
func (v *Visitor) VisitRecursionParameters(c *parser.RecursionParametersContext) interface{} {
	recursionParameters := NewRecursionParameters()

	for i := 0; i < len(c.AllParameter()); i++ {
		valRes := c.Parameter(i).Accept(v)
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
					res := parameterValue.Accept(v).(*Result)
					if !hasErrors(res) {
						recursionParameters.Memoize = bool(res.Value.(Bool))
					}
				}
			}
		}
	}

	return recursionParameters
}

func (v *Visitor) VisitLambdaExpr(c *parser.LambdaExprContext) interface{} {
	stack := v.lambdaCallStack
	lambda := &CalledLambda{
		arguments: nil,
		name:      c.LAMBDA_VARIABLE().GetText(),
		children:  []*CalledLambda{},
		nested:    false,
	}

getDeclaration:
	declaration, found := v.lambdaDeclarations[lambda.name]

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
					if _, ok := v.lambdaDeclarations[lambda.name]; ok && !isLocalToGrandparent {
						goto getDeclaration
					}
				}
			} else {
				break
			}
		}
	}
	if !found {
		if v.strict {
			return NewResult(nil).WithError("undefined lambda " + lambda.name)
		}
		return NewResult(NewNumber(0))
	}

	// Init arguments
	if c.MixedTuple() != nil {
		valuesRes := c.MixedTuple().Accept(v).(*Result)
		if hasErrors(valuesRes) {
			return valuesRes
		}
		v.convertMixedTupleResult(valuesRes)
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
				return NewResult(nil).WithError("expected " + strconv.FormatInt(int64(count), 10) + " parameter" + s)
			}
		}

		// Create argName -> value map
		lambda.arguments = map[string]interface{}{}
		for i, argument := range declaration.arguments {
			// Check type of received parameter and check against expected argument
			if isLambdaName(argument.String()) {
				if _, ok := tuple[i].(String); !ok {
					return NewResult(nil).WithError("[" + lambda.name + "] expected lambda for parameter " + strconv.Itoa(i+1) + ", got number")
				}
			} else {
				if _, ok := tuple[i].(Number); !ok {
					return NewResult(nil).WithError("[" + lambda.name + "] expected number for parameter " + strconv.Itoa(i+1) + ", got lambda")
				}
			}
			lambda.arguments[argument.String()] = tuple[i]
		}
	}

	// Init recursion parameters
	if _, ok := stack.recursion[lambda.name]; !ok {
		if c.RecursionParameters() != nil {
			stack.recursion[lambda.name] = c.RecursionParameters().Accept(v).(*RecursionParameters)
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
		conditionRes := stack.recursion[lambda.name].StopWhen.Accept(v).(*Result)
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
		val := NewNumber(0)
		if stack.recursion[lambda.name].LastValue == nil {
			stack.Pop()
			return NewResult(val)
		}
		// Eval last value
		lastValRes := stack.recursion[lambda.name].LastValue.Accept(v).(*Result)
		if hasErrors(lastValRes) {
			return lastValRes
		}
		lastVal, ok := lastValRes.Value.(Number)
		if !ok {
			panic("unable to cast lastValRes to Number")
		}

		val.Set(lastVal.Decimal)
		stack.Pop()
		return NewResult(val)
	}

	// Handle recursion
	if recurring {
		recParameters := stack.recursion[lambda.name]
		invokes := stack.invokes[lambda.name]

		if recParameters.MaxRecurrences == 0 {
			return NewResult(nil).WithError("recursion is disabled")
		}

		if shouldStop {
			val := NewNumber(0)
			if stack.recursion[lambda.name].LastValue == nil {
				stack.Pop()
				return NewResult(val)
			}
			// Eval last value
			lastValRes := stack.recursion[lambda.name].LastValue.Accept(v).(*Result)
			if hasErrors(lastValRes) {
				return lastValRes
			}
			lastVal, ok := lastValRes.Value.(Number)
			if !ok {
				panic("unable to cast lastValRes to Number")
			}

			val.Set(lastVal.Decimal)
			stack.Pop()
			return NewResult(val)
		}
		if invokes > recParameters.MaxRecurrences {
			val := NewNumber(0)
			if stack.recursion[lambda.name].LastValue == nil {
				stack.Pop()
				return NewResult(val)
			}
			// Eval last value
			lastValRes := stack.recursion[lambda.name].LastValue.Accept(v).(*Result)
			if hasErrors(lastValRes) {
				return lastValRes
			}
			lastVal, ok := lastValRes.Value.(Number)
			if !ok {
				panic("unable to cast lastValRes to Number")
			}

			val.Set(lastVal.Decimal)
			stack.Pop()
			return NewResult(val)
		}
	}

	// Evaluate lambda
	declaredLambda := lambda.ctx.Lambda()
	result := declaredLambda.Accept(v).(*Result)
	stack.Pop()

	switch value := result.Value.(type) {
	case Tuple:
		if len(value) == 1 {
			v := NewNumber(0)
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

func (v *Visitor) VisitVariable(c *parser.VariableContext) interface{} {
	name := c.VARIABLE().GetText()

	lambda := v.lambdaCallStack.Peek()
	if lambda != nil {
		if value, ok := lambda.arguments[name]; ok {
			return NewResult(value.(Number))
		}
	}

	if value, ok := v.variables[name]; ok {
		return NewResult(value)
	}

	if v.strict {
		return NewResult(nil).WithError("undefined global variable " + name)
	}
	return NewResult(NewNumber(0))
}

func (v *Visitor) checkParentCtxForLambda(c antlr.Tree) (bool, string) {
	ctx, ok := interface{}(c).(*parser.LambdaDeclarationContext)

	lambdaName := ""
	if ok {
		lambdaName = ctx.LAMBDA_VARIABLE().GetText()
	}

	if !ok && c.GetParent() != nil {
		return v.checkParentCtxForLambda(c.GetParent())
	}
	return ok, lambdaName
}

func hasErrors(r *Result) bool {
	return len(r.Errors) != 0
}

func isLambdaName(name string) bool {
	return name[0] >= 'A' && name[0] <= 'Z'
}
