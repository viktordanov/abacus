// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

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

func (v *BaseAbacusVisitor) VisitNot(ctx *NotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitGreaterOrEqualComparison(ctx *GreaterOrEqualComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitGreaterComparison(ctx *GreaterComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLessOrEqualComparison(ctx *LessOrEqualComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitBooleanAtom(ctx *BooleanAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLessComparison(ctx *LessComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitParenthesesBoolean(ctx *ParenthesesBooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitAndOrXor(ctx *AndOrXorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitEqualComparison(ctx *EqualComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitBoolAtom(ctx *BoolAtomContext) interface{} {
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

func (v *BaseAbacusVisitor) VisitMod(ctx *ModContext) interface{} {
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

func (v *BaseAbacusVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitRecursionParameters(ctx *RecursionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitMixedTuple(ctx *MixedTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAbacusVisitor) VisitLambdaArguments(ctx *LambdaArgumentsContext) interface{} {
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

func (v *BaseAbacusVisitor) VisitFunctionInvocation(ctx *FunctionInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}
