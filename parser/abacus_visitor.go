// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by AbacusParser.
type AbacusVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AbacusParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by AbacusParser#VariableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by AbacusParser#LambdaDeclaration.
	VisitLambdaDeclaration(ctx *LambdaDeclarationContext) interface{}

	// Visit a parse tree produced by AbacusParser#EqualComparison.
	VisitEqualComparison(ctx *EqualComparisonContext) interface{}

	// Visit a parse tree produced by AbacusParser#LessComparison.
	VisitLessComparison(ctx *LessComparisonContext) interface{}

	// Visit a parse tree produced by AbacusParser#GreaterComparison.
	VisitGreaterComparison(ctx *GreaterComparisonContext) interface{}

	// Visit a parse tree produced by AbacusParser#LessOrEqualComparison.
	VisitLessOrEqualComparison(ctx *LessOrEqualComparisonContext) interface{}

	// Visit a parse tree produced by AbacusParser#GreaterOrEqualComparison.
	VisitGreaterOrEqualComparison(ctx *GreaterOrEqualComparisonContext) interface{}

	// Visit a parse tree produced by AbacusParser#VariablesLambda.
	VisitVariablesLambda(ctx *VariablesLambdaContext) interface{}

	// Visit a parse tree produced by AbacusParser#NullArityLambda.
	VisitNullArityLambda(ctx *NullArityLambdaContext) interface{}

	// Visit a parse tree produced by AbacusParser#SignedExpr.
	VisitSignedExpr(ctx *SignedExprContext) interface{}

	// Visit a parse tree produced by AbacusParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by AbacusParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by AbacusParser#LambdaExpr.
	VisitLambdaExpr(ctx *LambdaExprContext) interface{}

	// Visit a parse tree produced by AbacusParser#Pow.
	VisitPow(ctx *PowContext) interface{}

	// Visit a parse tree produced by AbacusParser#AtomExpr.
	VisitAtomExpr(ctx *AtomExprContext) interface{}

	// Visit a parse tree produced by AbacusParser#Parentheses.
	VisitParentheses(ctx *ParenthesesContext) interface{}

	// Visit a parse tree produced by AbacusParser#recursionParameters.
	VisitRecursionParameters(ctx *RecursionParametersContext) interface{}

	// Visit a parse tree produced by AbacusParser#tuple.
	VisitTuple(ctx *TupleContext) interface{}

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

	// Visit a parse tree produced by AbacusParser#SqrtFunction.
	VisitSqrtFunction(ctx *SqrtFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#LnFunction.
	VisitLnFunction(ctx *LnFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#LogDefFunction.
	VisitLogDefFunction(ctx *LogDefFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#Log2Function.
	VisitLog2Function(ctx *Log2FunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#Log10Function.
	VisitLog10Function(ctx *Log10FunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#FloorFunction.
	VisitFloorFunction(ctx *FloorFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#CeilFunction.
	VisitCeilFunction(ctx *CeilFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#ExpFunction.
	VisitExpFunction(ctx *ExpFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#SinFunction.
	VisitSinFunction(ctx *SinFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#CosFunction.
	VisitCosFunction(ctx *CosFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#TanFunction.
	VisitTanFunction(ctx *TanFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#RoundDefFunction.
	VisitRoundDefFunction(ctx *RoundDefFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#AbsFunction.
	VisitAbsFunction(ctx *AbsFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#Round2Function.
	VisitRound2Function(ctx *Round2FunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#LogFunction.
	VisitLogFunction(ctx *LogFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#MinFunction.
	VisitMinFunction(ctx *MinFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#MaxFunction.
	VisitMaxFunction(ctx *MaxFunctionContext) interface{}

	// Visit a parse tree produced by AbacusParser#AvgFunction.
	VisitAvgFunction(ctx *AvgFunctionContext) interface{}
}
