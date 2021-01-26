// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr"

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

	// EnterRecursionParameters is called when entering the recursionParameters production.
	EnterRecursionParameters(c *RecursionParametersContext)

	// EnterTuple is called when entering the tuple production.
	EnterTuple(c *TupleContext)

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

	// EnterSqrtFunction is called when entering the SqrtFunction production.
	EnterSqrtFunction(c *SqrtFunctionContext)

	// EnterCbrtFunction is called when entering the CbrtFunction production.
	EnterCbrtFunction(c *CbrtFunctionContext)

	// EnterLnFunction is called when entering the LnFunction production.
	EnterLnFunction(c *LnFunctionContext)

	// EnterLogDefFunction is called when entering the LogDefFunction production.
	EnterLogDefFunction(c *LogDefFunctionContext)

	// EnterLog2Function is called when entering the Log2Function production.
	EnterLog2Function(c *Log2FunctionContext)

	// EnterLog10Function is called when entering the Log10Function production.
	EnterLog10Function(c *Log10FunctionContext)

	// EnterFloorFunction is called when entering the FloorFunction production.
	EnterFloorFunction(c *FloorFunctionContext)

	// EnterCeilFunction is called when entering the CeilFunction production.
	EnterCeilFunction(c *CeilFunctionContext)

	// EnterExpFunction is called when entering the ExpFunction production.
	EnterExpFunction(c *ExpFunctionContext)

	// EnterSinFunction is called when entering the SinFunction production.
	EnterSinFunction(c *SinFunctionContext)

	// EnterCosFunction is called when entering the CosFunction production.
	EnterCosFunction(c *CosFunctionContext)

	// EnterTanFunction is called when entering the TanFunction production.
	EnterTanFunction(c *TanFunctionContext)

	// EnterRoundDefFunction is called when entering the RoundDefFunction production.
	EnterRoundDefFunction(c *RoundDefFunctionContext)

	// EnterAbsFunction is called when entering the AbsFunction production.
	EnterAbsFunction(c *AbsFunctionContext)

	// EnterRound2Function is called when entering the Round2Function production.
	EnterRound2Function(c *Round2FunctionContext)

	// EnterLogFunction is called when entering the LogFunction production.
	EnterLogFunction(c *LogFunctionContext)

	// EnterMinFunction is called when entering the MinFunction production.
	EnterMinFunction(c *MinFunctionContext)

	// EnterMaxFunction is called when entering the MaxFunction production.
	EnterMaxFunction(c *MaxFunctionContext)

	// EnterAvgFunction is called when entering the AvgFunction production.
	EnterAvgFunction(c *AvgFunctionContext)

	// EnterUntilFunction is called when entering the UntilFunction production.
	EnterUntilFunction(c *UntilFunctionContext)

	// EnterFromFunction is called when entering the FromFunction production.
	EnterFromFunction(c *FromFunctionContext)

	// EnterReverseFunction is called when entering the ReverseFunction production.
	EnterReverseFunction(c *ReverseFunctionContext)

	// EnterNthFunction is called when entering the NthFunction production.
	EnterNthFunction(c *NthFunctionContext)

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

	// ExitRecursionParameters is called when exiting the recursionParameters production.
	ExitRecursionParameters(c *RecursionParametersContext)

	// ExitTuple is called when exiting the tuple production.
	ExitTuple(c *TupleContext)

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

	// ExitSqrtFunction is called when exiting the SqrtFunction production.
	ExitSqrtFunction(c *SqrtFunctionContext)

	// ExitCbrtFunction is called when exiting the CbrtFunction production.
	ExitCbrtFunction(c *CbrtFunctionContext)

	// ExitLnFunction is called when exiting the LnFunction production.
	ExitLnFunction(c *LnFunctionContext)

	// ExitLogDefFunction is called when exiting the LogDefFunction production.
	ExitLogDefFunction(c *LogDefFunctionContext)

	// ExitLog2Function is called when exiting the Log2Function production.
	ExitLog2Function(c *Log2FunctionContext)

	// ExitLog10Function is called when exiting the Log10Function production.
	ExitLog10Function(c *Log10FunctionContext)

	// ExitFloorFunction is called when exiting the FloorFunction production.
	ExitFloorFunction(c *FloorFunctionContext)

	// ExitCeilFunction is called when exiting the CeilFunction production.
	ExitCeilFunction(c *CeilFunctionContext)

	// ExitExpFunction is called when exiting the ExpFunction production.
	ExitExpFunction(c *ExpFunctionContext)

	// ExitSinFunction is called when exiting the SinFunction production.
	ExitSinFunction(c *SinFunctionContext)

	// ExitCosFunction is called when exiting the CosFunction production.
	ExitCosFunction(c *CosFunctionContext)

	// ExitTanFunction is called when exiting the TanFunction production.
	ExitTanFunction(c *TanFunctionContext)

	// ExitRoundDefFunction is called when exiting the RoundDefFunction production.
	ExitRoundDefFunction(c *RoundDefFunctionContext)

	// ExitAbsFunction is called when exiting the AbsFunction production.
	ExitAbsFunction(c *AbsFunctionContext)

	// ExitRound2Function is called when exiting the Round2Function production.
	ExitRound2Function(c *Round2FunctionContext)

	// ExitLogFunction is called when exiting the LogFunction production.
	ExitLogFunction(c *LogFunctionContext)

	// ExitMinFunction is called when exiting the MinFunction production.
	ExitMinFunction(c *MinFunctionContext)

	// ExitMaxFunction is called when exiting the MaxFunction production.
	ExitMaxFunction(c *MaxFunctionContext)

	// ExitAvgFunction is called when exiting the AvgFunction production.
	ExitAvgFunction(c *AvgFunctionContext)

	// ExitUntilFunction is called when exiting the UntilFunction production.
	ExitUntilFunction(c *UntilFunctionContext)

	// ExitFromFunction is called when exiting the FromFunction production.
	ExitFromFunction(c *FromFunctionContext)

	// ExitReverseFunction is called when exiting the ReverseFunction production.
	ExitReverseFunction(c *ReverseFunctionContext)

	// ExitNthFunction is called when exiting the NthFunction production.
	ExitNthFunction(c *NthFunctionContext)
}
