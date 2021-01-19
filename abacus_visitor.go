package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/viktordanov/abacus/parser"
	"math"
	"strconv"
)

type AbacusVisitor struct {
	antlr.ParseTreeVisitor

	trace []string
}

func NewAbacusVisitor() *AbacusVisitor {
	return &AbacusVisitor{
		ParseTreeVisitor: &parser.BaseAbacusVisitor{},
		trace:            []string{},
	}
}

func (a *AbacusVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.StartContext:
		a.trace = append(a.trace, "root")
		return val.Accept(a)
	}
	return nil
}

func (a *AbacusVisitor) VisitStart(c *parser.StartContext) interface{} {
	a.trace = append(a.trace, "start")
	return c.Expression().Accept(a)
}

func (a *AbacusVisitor) VisitExpression(c *parser.ExpressionContext) interface{} {
	a.trace = append(a.trace, "expression")
	return c.Accept(a)
}

func (a *AbacusVisitor) VisitMulDiv(c *parser.MulDivContext) interface{} {
	a.trace = append(a.trace, "muldiv")
	first := c.Expression(0).Accept(a).(float64)
	second := c.Expression(1).Accept(a).(float64)

	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserMUL:
		return first * second
	case parser.AbacusLexerDIV:
		return first / second
	}
	return 0
}

func (a *AbacusVisitor) VisitAddSub(c *parser.AddSubContext) interface{} {
	a.trace = append(a.trace, "addsub")
	first := c.Expression(0).Accept(a).(float64)
	second := c.Expression(1).Accept(a).(float64)

	switch c.GetOp().GetTokenType() {
	case parser.AbacusParserADD:
		return first + second
	case parser.AbacusLexerSUB:
		return first - second
	}
	return nil
}

func (a *AbacusVisitor) VisitPow(c *parser.PowContext) interface{} {
	a.trace = append(a.trace, "pow")
	first := c.Expression(0).Accept(a).(float64)
	second := c.Expression(1).Accept(a).(float64)
	return math.Pow(first, second)
}

func (a *AbacusVisitor) VisitParentheses(c *parser.ParenthesesContext) interface{} {
	a.trace = append(a.trace, "parentheses")
	return c.Expression().Accept(a)
}

func (a *AbacusVisitor) VisitNumber(c *parser.NumberContext) interface{} {
	a.trace = append(a.trace, "num")
	first := c.NUMBER().GetText()
	out, err := strconv.ParseFloat(first, 64)
	if err != nil {
		panic(err)
	}
	return out
}
