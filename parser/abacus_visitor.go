// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by AbacusParser.
type AbacusVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AbacusParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by AbacusParser#VariableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by AbacusParser#LambdaDeclaration.
	VisitLambdaDeclaration(ctx *LambdaDeclarationContext) interface{}

	// Visit a parse tree produced by AbacusParser#Not.
	VisitNot(ctx *NotContext) interface{}

	// Visit a parse tree produced by AbacusParser#GreaterOrEqualComparison.
	VisitGreaterOrEqualComparison(ctx *GreaterOrEqualComparisonContext) interface{}

	// Visit a parse tree produced by AbacusParser#GreaterComparison.
	VisitGreaterComparison(ctx *GreaterComparisonContext) interface{}

	// Visit a parse tree produced by AbacusParser#LessOrEqualComparison.
	VisitLessOrEqualComparison(ctx *LessOrEqualComparisonContext) interface{}

	// Visit a parse tree produced by AbacusParser#BooleanAtom.
	VisitBooleanAtom(ctx *BooleanAtomContext) interface{}

	// Visit a parse tree produced by AbacusParser#LessComparison.
	VisitLessComparison(ctx *LessComparisonContext) interface{}

	// Visit a parse tree produced by AbacusParser#ParenthesesBoolean.
	VisitParenthesesBoolean(ctx *ParenthesesBooleanContext) interface{}

	// Visit a parse tree produced by AbacusParser#AndOrXor.
	VisitAndOrXor(ctx *AndOrXorContext) interface{}

	// Visit a parse tree produced by AbacusParser#EqualComparison.
	VisitEqualComparison(ctx *EqualComparisonContext) interface{}

	// Visit a parse tree produced by AbacusParser#boolAtom.
	VisitBoolAtom(ctx *BoolAtomContext) interface{}

	// Visit a parse tree produced by AbacusParser#VariablesLambda.
	VisitVariablesLambda(ctx *VariablesLambdaContext) interface{}

	// Visit a parse tree produced by AbacusParser#NullArityLambda.
	VisitNullArityLambda(ctx *NullArityLambdaContext) interface{}

	// Visit a parse tree produced by AbacusParser#SignedExpr.
	VisitSignedExpr(ctx *SignedExprContext) interface{}

	// Visit a parse tree produced by AbacusParser#Mod.
	VisitMod(ctx *ModContext) interface{}

	// Visit a parse tree produced by AbacusParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by AbacusParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by AbacusParser#Percent.
	VisitPercent(ctx *PercentContext) interface{}

	// Visit a parse tree produced by AbacusParser#LambdaExpr.
	VisitLambdaExpr(ctx *LambdaExprContext) interface{}

	// Visit a parse tree produced by AbacusParser#Pow.
	VisitPow(ctx *PowContext) interface{}

	// Visit a parse tree produced by AbacusParser#AtomExpr.
	VisitAtomExpr(ctx *AtomExprContext) interface{}

	// Visit a parse tree produced by AbacusParser#Parentheses.
	VisitParentheses(ctx *ParenthesesContext) interface{}

	// Visit a parse tree produced by AbacusParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by AbacusParser#recursionParameters.
	VisitRecursionParameters(ctx *RecursionParametersContext) interface{}

	// Visit a parse tree produced by AbacusParser#mixedTuple.
	VisitMixedTuple(ctx *MixedTupleContext) interface{}

	// Visit a parse tree produced by AbacusParser#tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by AbacusParser#lambdaArguments.
	VisitLambdaArguments(ctx *LambdaArgumentsContext) interface{}

	// Visit a parse tree produced by AbacusParser#variablesTuple.
	VisitVariablesTuple(ctx *VariablesTupleContext) interface{}

	// Visit a parse tree produced by AbacusParser#FuncExpr.
	VisitFuncExpr(ctx *FuncExprContext) interface{}

	// Visit a parse tree produced by AbacusParser#Constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by AbacusParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by AbacusParser#Variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by AbacusParser#PlusSign.
	VisitPlusSign(ctx *PlusSignContext) interface{}

	// Visit a parse tree produced by AbacusParser#MinusSign.
	VisitMinusSign(ctx *MinusSignContext) interface{}

	// Visit a parse tree produced by AbacusParser#FunctionInvocation.
	VisitFunctionInvocation(ctx *FunctionInvocationContext) interface{}
}
