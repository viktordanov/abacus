// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// AbacusListener is a complete listener for a parse tree produced by AbacusParser.
type AbacusListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterVariableDeclaration is called when entering the VariableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterLambdaDeclaration is called when entering the LambdaDeclaration production.
	EnterLambdaDeclaration(c *LambdaDeclarationContext)

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterGreaterOrEqualComparison is called when entering the GreaterOrEqualComparison production.
	EnterGreaterOrEqualComparison(c *GreaterOrEqualComparisonContext)

	// EnterGreaterComparison is called when entering the GreaterComparison production.
	EnterGreaterComparison(c *GreaterComparisonContext)

	// EnterLessOrEqualComparison is called when entering the LessOrEqualComparison production.
	EnterLessOrEqualComparison(c *LessOrEqualComparisonContext)

	// EnterBooleanAtom is called when entering the BooleanAtom production.
	EnterBooleanAtom(c *BooleanAtomContext)

	// EnterLessComparison is called when entering the LessComparison production.
	EnterLessComparison(c *LessComparisonContext)

	// EnterParenthesesBoolean is called when entering the ParenthesesBoolean production.
	EnterParenthesesBoolean(c *ParenthesesBooleanContext)

	// EnterAndOrXor is called when entering the AndOrXor production.
	EnterAndOrXor(c *AndOrXorContext)

	// EnterEqualComparison is called when entering the EqualComparison production.
	EnterEqualComparison(c *EqualComparisonContext)

	// EnterBoolAtom is called when entering the boolAtom production.
	EnterBoolAtom(c *BoolAtomContext)

	// EnterVariablesLambda is called when entering the VariablesLambda production.
	EnterVariablesLambda(c *VariablesLambdaContext)

	// EnterNullArityLambda is called when entering the NullArityLambda production.
	EnterNullArityLambda(c *NullArityLambdaContext)

	// EnterSignedExpr is called when entering the SignedExpr production.
	EnterSignedExpr(c *SignedExprContext)

	// EnterMod is called when entering the Mod production.
	EnterMod(c *ModContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterPercent is called when entering the Percent production.
	EnterPercent(c *PercentContext)

	// EnterLambdaExpr is called when entering the LambdaExpr production.
	EnterLambdaExpr(c *LambdaExprContext)

	// EnterPow is called when entering the Pow production.
	EnterPow(c *PowContext)

	// EnterAtomExpr is called when entering the AtomExpr production.
	EnterAtomExpr(c *AtomExprContext)

	// EnterParentheses is called when entering the Parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterRecursionParameters is called when entering the recursionParameters production.
	EnterRecursionParameters(c *RecursionParametersContext)

	// EnterMixedTuple is called when entering the mixedTuple production.
	EnterMixedTuple(c *MixedTupleContext)

	// EnterTuple is called when entering the tuple production.
	EnterTuple(c *TupleContext)

	// EnterLambdaArguments is called when entering the lambdaArguments production.
	EnterLambdaArguments(c *LambdaArgumentsContext)

	// EnterVariablesTuple is called when entering the variablesTuple production.
	EnterVariablesTuple(c *VariablesTupleContext)

	// EnterFuncExpr is called when entering the FuncExpr production.
	EnterFuncExpr(c *FuncExprContext)

	// EnterConstant is called when entering the Constant production.
	EnterConstant(c *ConstantContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterVariable is called when entering the Variable production.
	EnterVariable(c *VariableContext)

	// EnterPlusSign is called when entering the PlusSign production.
	EnterPlusSign(c *PlusSignContext)

	// EnterMinusSign is called when entering the MinusSign production.
	EnterMinusSign(c *MinusSignContext)

	// EnterFunctionInvocation is called when entering the FunctionInvocation production.
	EnterFunctionInvocation(c *FunctionInvocationContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitVariableDeclaration is called when exiting the VariableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitLambdaDeclaration is called when exiting the LambdaDeclaration production.
	ExitLambdaDeclaration(c *LambdaDeclarationContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitGreaterOrEqualComparison is called when exiting the GreaterOrEqualComparison production.
	ExitGreaterOrEqualComparison(c *GreaterOrEqualComparisonContext)

	// ExitGreaterComparison is called when exiting the GreaterComparison production.
	ExitGreaterComparison(c *GreaterComparisonContext)

	// ExitLessOrEqualComparison is called when exiting the LessOrEqualComparison production.
	ExitLessOrEqualComparison(c *LessOrEqualComparisonContext)

	// ExitBooleanAtom is called when exiting the BooleanAtom production.
	ExitBooleanAtom(c *BooleanAtomContext)

	// ExitLessComparison is called when exiting the LessComparison production.
	ExitLessComparison(c *LessComparisonContext)

	// ExitParenthesesBoolean is called when exiting the ParenthesesBoolean production.
	ExitParenthesesBoolean(c *ParenthesesBooleanContext)

	// ExitAndOrXor is called when exiting the AndOrXor production.
	ExitAndOrXor(c *AndOrXorContext)

	// ExitEqualComparison is called when exiting the EqualComparison production.
	ExitEqualComparison(c *EqualComparisonContext)

	// ExitBoolAtom is called when exiting the boolAtom production.
	ExitBoolAtom(c *BoolAtomContext)

	// ExitVariablesLambda is called when exiting the VariablesLambda production.
	ExitVariablesLambda(c *VariablesLambdaContext)

	// ExitNullArityLambda is called when exiting the NullArityLambda production.
	ExitNullArityLambda(c *NullArityLambdaContext)

	// ExitSignedExpr is called when exiting the SignedExpr production.
	ExitSignedExpr(c *SignedExprContext)

	// ExitMod is called when exiting the Mod production.
	ExitMod(c *ModContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitPercent is called when exiting the Percent production.
	ExitPercent(c *PercentContext)

	// ExitLambdaExpr is called when exiting the LambdaExpr production.
	ExitLambdaExpr(c *LambdaExprContext)

	// ExitPow is called when exiting the Pow production.
	ExitPow(c *PowContext)

	// ExitAtomExpr is called when exiting the AtomExpr production.
	ExitAtomExpr(c *AtomExprContext)

	// ExitParentheses is called when exiting the Parentheses production.
	ExitParentheses(c *ParenthesesContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitRecursionParameters is called when exiting the recursionParameters production.
	ExitRecursionParameters(c *RecursionParametersContext)

	// ExitMixedTuple is called when exiting the mixedTuple production.
	ExitMixedTuple(c *MixedTupleContext)

	// ExitTuple is called when exiting the tuple production.
	ExitTuple(c *TupleContext)

	// ExitLambdaArguments is called when exiting the lambdaArguments production.
	ExitLambdaArguments(c *LambdaArgumentsContext)

	// ExitVariablesTuple is called when exiting the variablesTuple production.
	ExitVariablesTuple(c *VariablesTupleContext)

	// ExitFuncExpr is called when exiting the FuncExpr production.
	ExitFuncExpr(c *FuncExprContext)

	// ExitConstant is called when exiting the Constant production.
	ExitConstant(c *ConstantContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitVariable is called when exiting the Variable production.
	ExitVariable(c *VariableContext)

	// ExitPlusSign is called when exiting the PlusSign production.
	ExitPlusSign(c *PlusSignContext)

	// ExitMinusSign is called when exiting the MinusSign production.
	ExitMinusSign(c *MinusSignContext)

	// ExitFunctionInvocation is called when exiting the FunctionInvocation production.
	ExitFunctionInvocation(c *FunctionInvocationContext)
}
