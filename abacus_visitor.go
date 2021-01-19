package main

import (
	"github.com/ALTree/bigfloat"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/viktordanov/abacus/parser"
	"math"
	"math/big"
)

type AbacusVisitor struct {
	antlr.ParseTreeVisitor

	trace []string
	vars  map[string]*big.Float
}

func NewAbacusVisitor() *AbacusVisitor {
	return &AbacusVisitor{
		ParseTreeVisitor: &parser.BaseAbacusVisitor{},
		trace:            []string{},
		vars:             make(map[string]*big.Float),
	}
}

func (a *AbacusVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.RootContext:
		a.trace = append(a.trace, "root")
		return val.Accept(a)

	case *parser.DeclarationContext:
		a.trace = append(a.trace, "declaration")
		return val.Accept(a)

	case *parser.ExpressionContext:
		a.trace = append(a.trace, "expression")
		return val.Accept(a)
	}
	return nil
}

func (a *AbacusVisitor) VisitRoot(c *parser.RootContext) interface{} {
	a.trace = append(a.trace, "root")

	if c.Expression() != nil {
		return c.Expression().Accept(a)
	}
	if c.Declaration() != nil {
		return c.Declaration().Accept(a)
	}
	return nil
}

func (a *AbacusVisitor) VisitDeclaration(c *parser.DeclarationContext) interface{} {
	a.trace = append(a.trace, "declaration")
	variableName := c.VARIABLE().GetText()
	value := c.Expression().Accept(a).(*big.Float)

	a.vars[variableName] = value
	return value
}

func (a *AbacusVisitor) VisitMulDiv(c *parser.MulDivContext) interface{} {
	a.trace = append(a.trace, "muldiv")
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
	a.trace = append(a.trace, "addsub")
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
	a.trace = append(a.trace, "pow")
	first := c.Expression(0).Accept(a).(*big.Float)
	second, _ := c.Expression(1).Accept(a).(*big.Float).Int64()
	return Pow(first, uint64(second))
}

func (a *AbacusVisitor) VisitParentheses(c *parser.ParenthesesContext) interface{} {
	a.trace = append(a.trace, "parentheses")
	return c.Expression().Accept(a)
}

func (a *AbacusVisitor) VisitAtomExpr(c *parser.AtomExprContext) interface{} {
	return c.Atom().Accept(a)
}

func (a *AbacusVisitor) VisitFuncExpr(c *parser.FuncExprContext) interface{} {
	return c.Function().Accept(a)
}

func (a *AbacusVisitor) VisitSqrtFunction(c *parser.SqrtFunctionContext) interface{} {
	a.trace = append(a.trace, "sqrt")
	arg := c.Expression().Accept(a).(*big.Float)
	return bigfloat.Sqrt(arg)
}

func (a *AbacusVisitor) VisitLnFunction(c *parser.LnFunctionContext) interface{} {
	a.trace = append(a.trace, "sqrt")
	arg := c.Expression().Accept(a).(*big.Float)
	return bigfloat.Log(arg)
}
func (a *AbacusVisitor) VisitFloorFunction(c *parser.FloorFunctionContext) interface{} {
	a.trace = append(a.trace, "sqrt")
	arg := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := arg.Float64()
	return big.NewFloat(math.Floor(toFloat))
}

func (a *AbacusVisitor) VisitCeilFunction(c *parser.CeilFunctionContext) interface{} {
	a.trace = append(a.trace, "sqrt")
	arg := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := arg.Float64()
	return big.NewFloat(math.Ceil(toFloat))
}
func (a *AbacusVisitor) VisitSinFunction(c *parser.SinFunctionContext) interface{} {
	a.trace = append(a.trace, "sqrt")
	arg := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := arg.Float64()
	return big.NewFloat(math.Sin(toFloat))
}
func (a *AbacusVisitor) VisitCosFunction(c *parser.CosFunctionContext) interface{} {
	a.trace = append(a.trace, "sqrt")
	arg := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := arg.Float64()
	return big.NewFloat(math.Cos(toFloat))
}
func (a *AbacusVisitor) VisitTanFunction(c *parser.TanFunctionContext) interface{} {
	a.trace = append(a.trace, "sqrt")
	arg := c.Expression().Accept(a).(*big.Float)
	toFloat, _ := arg.Float64()
	return big.NewFloat(math.Tan(toFloat))
}
func (a *AbacusVisitor) VisitExpFunction(c *parser.ExpFunctionContext) interface{} {
	a.trace = append(a.trace, "sqrt")
	arg := c.Expression().Accept(a).(*big.Float)
	return bigfloat.Exp(arg)
}

func (a *AbacusVisitor) VisitConstant(c *parser.ConstantContext) interface{} {
	a.trace = append(a.trace, "const")

	switch c.CONSTANT().GetText() {
	case "pi":
		return big.NewFloat(math.Pi)
	case "phi":
		return big.NewFloat(math.Phi)
	case "e":
		return big.NewFloat(math.E)

	}

	return 0
}

func (a *AbacusVisitor) VisitNumber(c *parser.NumberContext) interface{} {
	a.trace = append(a.trace, "num")

	out, _, err := big.ParseFloat(c.SCIENTIFIC_NUMBER().GetText(), 10, 256, big.ToNearestEven)
	if err != nil {
		panic(err)
	}
	return out
}

func (a *AbacusVisitor) VisitVariable(c *parser.VariableContext) interface{} {
	a.trace = append(a.trace, "var")

	var value *big.Float
	ok := false
	if value, ok = a.vars[c.VARIABLE().GetText()]; !ok {
		return big.NewFloat(0)
	}

	return value
}
