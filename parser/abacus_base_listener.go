// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Abacus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAbacusListener is a complete listener for a parse tree produced by AbacusParser.
type BaseAbacusListener struct{}

var _ AbacusListener = &BaseAbacusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAbacusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAbacusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAbacusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAbacusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseAbacusListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseAbacusListener) ExitRoot(ctx *RootContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseAbacusListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseAbacusListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseAbacusListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseAbacusListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseAbacusListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseAbacusListener) ExitAddSub(ctx *AddSubContext) {}

// EnterPow is called when production Pow is entered.
func (s *BaseAbacusListener) EnterPow(ctx *PowContext) {}

// ExitPow is called when production Pow is exited.
func (s *BaseAbacusListener) ExitPow(ctx *PowContext) {}

// EnterAtomExpr is called when production AtomExpr is entered.
func (s *BaseAbacusListener) EnterAtomExpr(ctx *AtomExprContext) {}

// ExitAtomExpr is called when production AtomExpr is exited.
func (s *BaseAbacusListener) ExitAtomExpr(ctx *AtomExprContext) {}

// EnterParentheses is called when production Parentheses is entered.
func (s *BaseAbacusListener) EnterParentheses(ctx *ParenthesesContext) {}

// ExitParentheses is called when production Parentheses is exited.
func (s *BaseAbacusListener) ExitParentheses(ctx *ParenthesesContext) {}

// EnterFuncExpr is called when production FuncExpr is entered.
func (s *BaseAbacusListener) EnterFuncExpr(ctx *FuncExprContext) {}

// ExitFuncExpr is called when production FuncExpr is exited.
func (s *BaseAbacusListener) ExitFuncExpr(ctx *FuncExprContext) {}

// EnterConstant is called when production Constant is entered.
func (s *BaseAbacusListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production Constant is exited.
func (s *BaseAbacusListener) ExitConstant(ctx *ConstantContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseAbacusListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseAbacusListener) ExitNumber(ctx *NumberContext) {}

// EnterVariable is called when production Variable is entered.
func (s *BaseAbacusListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production Variable is exited.
func (s *BaseAbacusListener) ExitVariable(ctx *VariableContext) {}

// EnterSqrtFunction is called when production SqrtFunction is entered.
func (s *BaseAbacusListener) EnterSqrtFunction(ctx *SqrtFunctionContext) {}

// ExitSqrtFunction is called when production SqrtFunction is exited.
func (s *BaseAbacusListener) ExitSqrtFunction(ctx *SqrtFunctionContext) {}

// EnterLnFunction is called when production LnFunction is entered.
func (s *BaseAbacusListener) EnterLnFunction(ctx *LnFunctionContext) {}

// ExitLnFunction is called when production LnFunction is exited.
func (s *BaseAbacusListener) ExitLnFunction(ctx *LnFunctionContext) {}

// EnterFloorFunction is called when production FloorFunction is entered.
func (s *BaseAbacusListener) EnterFloorFunction(ctx *FloorFunctionContext) {}

// ExitFloorFunction is called when production FloorFunction is exited.
func (s *BaseAbacusListener) ExitFloorFunction(ctx *FloorFunctionContext) {}

// EnterCeilFunction is called when production CeilFunction is entered.
func (s *BaseAbacusListener) EnterCeilFunction(ctx *CeilFunctionContext) {}

// ExitCeilFunction is called when production CeilFunction is exited.
func (s *BaseAbacusListener) ExitCeilFunction(ctx *CeilFunctionContext) {}

// EnterExpFunction is called when production ExpFunction is entered.
func (s *BaseAbacusListener) EnterExpFunction(ctx *ExpFunctionContext) {}

// ExitExpFunction is called when production ExpFunction is exited.
func (s *BaseAbacusListener) ExitExpFunction(ctx *ExpFunctionContext) {}

// EnterSinFunction is called when production SinFunction is entered.
func (s *BaseAbacusListener) EnterSinFunction(ctx *SinFunctionContext) {}

// ExitSinFunction is called when production SinFunction is exited.
func (s *BaseAbacusListener) ExitSinFunction(ctx *SinFunctionContext) {}

// EnterCosFunction is called when production CosFunction is entered.
func (s *BaseAbacusListener) EnterCosFunction(ctx *CosFunctionContext) {}

// ExitCosFunction is called when production CosFunction is exited.
func (s *BaseAbacusListener) ExitCosFunction(ctx *CosFunctionContext) {}

// EnterTanFunction is called when production TanFunction is entered.
func (s *BaseAbacusListener) EnterTanFunction(ctx *TanFunctionContext) {}

// ExitTanFunction is called when production TanFunction is exited.
func (s *BaseAbacusListener) ExitTanFunction(ctx *TanFunctionContext) {}
