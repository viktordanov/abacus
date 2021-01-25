// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseAbacusVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAbacusVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLambdaDeclaration(ctx *LambdaDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitEqualComparison(ctx *EqualComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLessComparison(ctx *LessComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitGreaterComparison(ctx *GreaterComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLessOrEqualComparison(ctx *LessOrEqualComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitGreaterOrEqualComparison(ctx *GreaterOrEqualComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitVariablesLambda(ctx *VariablesLambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitNullArityLambda(ctx *NullArityLambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitSignedExpr(ctx *SignedExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitPercent(ctx *PercentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLambdaExpr(ctx *LambdaExprContext) interface{} {
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

func (v *BaseAbacusVisitor) VisitRecursionParameters(ctx *RecursionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitVariablesTuple(ctx *VariablesTupleContext) interface{} {
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

func (v *BaseAbacusVisitor) VisitPlusSign(ctx *PlusSignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitMinusSign(ctx *MinusSignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitSqrtFunction(ctx *SqrtFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitCbrtFunction(ctx *CbrtFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLnFunction(ctx *LnFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLogDefFunction(ctx *LogDefFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLog2Function(ctx *Log2FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLog10Function(ctx *Log10FunctionContext) interface{} {
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

func (v *BaseAbacusVisitor) VisitRoundDefFunction(ctx *RoundDefFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitAbsFunction(ctx *AbsFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitRound2Function(ctx *Round2FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLogFunction(ctx *LogFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitMinFunction(ctx *MinFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitMaxFunction(ctx *MaxFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitAvgFunction(ctx *AvgFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}
