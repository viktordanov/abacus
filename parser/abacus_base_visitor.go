// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseAbacusVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAbacusVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitPow(ctx *PowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitParentheses(ctx *ParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}
