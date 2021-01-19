// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseAbacusVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAbacusVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
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

func (v *BaseAbacusVisitor) VisitAtomExpr(ctx *AtomExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitParentheses(ctx *ParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitFuncExpr(ctx *FuncExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitSqrtFunction(ctx *SqrtFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLnFunction(ctx *LnFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitFloorFunction(ctx *FloorFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitCeilFunction(ctx *CeilFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitExpFunction(ctx *ExpFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitSinFunction(ctx *SinFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitCosFunction(ctx *CosFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitTanFunction(ctx *TanFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}
