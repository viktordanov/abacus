// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AbacusListener is a complete listener for a parse tree produced by AbacusParser.
type AbacusListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

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

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterVariable is called when entering the Variable production.
	EnterVariable(c *VariableContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

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

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitVariable is called when exiting the Variable production.
	ExitVariable(c *VariableContext)
}
