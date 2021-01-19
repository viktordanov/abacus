// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AbacusListener is a complete listener for a parse tree produced by AbacusParser.
type AbacusListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterPow is called when entering the Pow production.
	EnterPow(c *PowContext)

	// EnterAtomExpr is called when entering the AtomExpr production.
	EnterAtomExpr(c *AtomExprContext)

	// EnterParentheses is called when entering the Parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// EnterFuncExpr is called when entering the FuncExpr production.
	EnterFuncExpr(c *FuncExprContext)

	// EnterConstant is called when entering the Constant production.
	EnterConstant(c *ConstantContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterVariable is called when entering the Variable production.
	EnterVariable(c *VariableContext)

	// EnterSqrtFunction is called when entering the SqrtFunction production.
	EnterSqrtFunction(c *SqrtFunctionContext)

	// EnterLnFunction is called when entering the LnFunction production.
	EnterLnFunction(c *LnFunctionContext)

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

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitPow is called when exiting the Pow production.
	ExitPow(c *PowContext)

	// ExitAtomExpr is called when exiting the AtomExpr production.
	ExitAtomExpr(c *AtomExprContext)

	// ExitParentheses is called when exiting the Parentheses production.
	ExitParentheses(c *ParenthesesContext)

	// ExitFuncExpr is called when exiting the FuncExpr production.
	ExitFuncExpr(c *FuncExprContext)

	// ExitConstant is called when exiting the Constant production.
	ExitConstant(c *ConstantContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitVariable is called when exiting the Variable production.
	ExitVariable(c *VariableContext)

	// ExitSqrtFunction is called when exiting the SqrtFunction production.
	ExitSqrtFunction(c *SqrtFunctionContext)

	// ExitLnFunction is called when exiting the LnFunction production.
	ExitLnFunction(c *LnFunctionContext)

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
}
