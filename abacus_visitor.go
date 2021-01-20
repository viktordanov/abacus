package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/viktordanov/abacus/parser"
	"math"
	"math/big"
)

type AbacusVisitor struct {
	antlr.ParseTreeVisitor
	vars map[string]*big.Float
}

func NewAbacusVisitor() *AbacusVisitor {
	return &AbacusVisitor{
		ParseTreeVisitor: &parser.BaseAbacusVisitor{},
		vars:             make(map[string]*big.Float),
	}
}

func (a *AbacusVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.RootContext:
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

	if c.Expression() != nil {
		return c.Expression().Accept(a)
	}
	if c.Declaration() != nil {
		return c.Declaration().Accept(a)
	}
	if c.Comparison() != nil {
		return c.Comparison().Accept(a)
	}
	return nil
}

func (a *AbacusVisitor) VisitDeclaration(c *parser.DeclarationContext) interface{} {
	variableName := c.VARIABLE().GetText()
	value := c.Expression().Accept(a).(*big.Float)

	a.vars[variableName] = value
	return variableAssignment{newValue: value}
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
	first := c.Expression(0).Accept(a).(*big.Float)
	second := c.Expression(1).Accept(a).(*big.Float)

	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserMUL:
		return Mul(first, second)
	case parser.AbacusLexerDIV:
		return Div(first, second)
	}
	return 0
}

func (a *AbacusVisitor) VisitAddSub(c *parser.AddSubContext) interface{} {
	first := c.Expression(0).Accept(a).(*big.Float)
	second := c.Expression(1).Accept(a).(*big.Float)

	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserADD:
		return Add(first, second)
	case parser.AbacusLexerSUB:
		return Sub(first, second)
	}
	return nil
}

func (a *AbacusVisitor) VisitPow(c *parser.PowContext) interface{} {
	first := c.Expression(0).Accept(a).(*big.Float)
	second := c.Expression(1).Accept(a).(*big.Float)
	return Pow(first, second)
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
	arg := c.Expression().Accept(a).(*big.Float)
	return Sqrt(arg)
}

func (a *AbacusVisitor) VisitLnFunction(c *parser.LnFunctionContext) interface{} {
	arg := c.Expression().Accept(a).(*big.Float)
	return Log(arg)
}

func (a *AbacusVisitor) VisitLogDefFunction(c *parser.LogDefFunctionContext) interface{} {
	arg := c.Expression().Accept(a).(*big.Float)
	return Log(arg)
}

func (a *AbacusVisitor) VisitLog2Function(c *parser.Log2FunctionContext) interface{} {
	arg := c.Expression().Accept(a).(*big.Float)
	return Div(Log(arg), Log(New(2)))
}

func (a *AbacusVisitor) VisitLog10Function(c *parser.Log10FunctionContext) interface{} {
	arg := c.Expression().Accept(a).(*big.Float)
	return Div(Log(arg), Log(New(10)))
}

func (a *AbacusVisitor) VisitFloorFunction(c *parser.FloorFunctionContext) interface{} {
	arg := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := arg.Float64()
	return big.NewFloat(math.Floor(toFloat))
}

func (a *AbacusVisitor) VisitCeilFunction(c *parser.CeilFunctionContext) interface{} {
	arg := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := arg.Float64()
	return big.NewFloat(math.Ceil(toFloat))
}
func (a *AbacusVisitor) VisitSinFunction(c *parser.SinFunctionContext) interface{} {
	arg := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := arg.Float64()
	return big.NewFloat(math.Sin(toFloat))
}
func (a *AbacusVisitor) VisitCosFunction(c *parser.CosFunctionContext) interface{} {
	arg := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := arg.Float64()
	return big.NewFloat(math.Cos(toFloat))
}
func (a *AbacusVisitor) VisitTanFunction(c *parser.TanFunctionContext) interface{} {
	arg := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := arg.Float64()
	return big.NewFloat(math.Tan(toFloat))
}
func (a *AbacusVisitor) VisitExpFunction(c *parser.ExpFunctionContext) interface{} {
	arg := c.Expression().Accept(a).(*big.Float)
	return Exp(arg)
}

func (a *AbacusVisitor) VisitRoundDefFunction(c *parser.RoundDefFunctionContext) interface{} {
	arg := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := arg.Float64()
	return big.NewFloat(math.Round(toFloat))
}

func (a *AbacusVisitor) VisitRound2Function(c *parser.Round2FunctionContext) interface{} {
	arg := c.Expression(0).Accept(a).(*big.Float)
	arg2 := c.Expression(1).Accept(a).(*big.Float)
	num, _ := arg.Float64()
	digits, _ := arg2.Float64()
	mult := math.Pow(10, digits+1)
	precision = uint(digits)
	return big.NewFloat(math.Round(num*mult) / mult)
}

func (a *AbacusVisitor) VisitLogFunction(c *parser.LogFunctionContext) interface{} {
	arg := c.Expression(0).Accept(a).(*big.Float)
	arg2 := c.Expression(1).Accept(a).(*big.Float)

	return Div(Log(arg), Log(arg2))
}

func (a *AbacusVisitor) VisitMinFunction(c *parser.MinFunctionContext) interface{} {
	arg := c.Expression(0).Accept(a).(*big.Float)
	arg2 := c.Expression(1).Accept(a).(*big.Float)
	if arg.Cmp(arg2) == -1 {
		return arg
	}
	return arg2
}

func (a *AbacusVisitor) VisitMaxFunction(c *parser.MaxFunctionContext) interface{} {
	arg := c.Expression(0).Accept(a).(*big.Float)
	arg2 := c.Expression(1).Accept(a).(*big.Float)
	if arg.Cmp(arg2) == 1 {
		return arg
	}
	return arg2
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
	out, _, err := big.ParseFloat(c.SCIENTIFIC_NUMBER().GetText(), 10, precision, big.ToNearestEven)
	if err != nil {
		panic(err)
	}
	return out
}

func (a *AbacusVisitor) VisitVariable(c *parser.VariableContext) interface{} {
	var value *big.Float
	ok := false
	if value, ok = a.vars[c.VARIABLE().GetText()]; !ok {
		return big.NewFloat(0)
	}
	return value
}
