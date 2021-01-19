// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Abacus

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 95, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 19, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 31, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 42, 10, 4, 12, 4, 14, 4, 45, 11, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 5, 5, 51, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 93,
	10, 6, 3, 6, 2, 3, 6, 7, 2, 4, 6, 8, 10, 2, 4, 3, 2, 13, 14, 3, 2, 15,
	16, 2, 104, 2, 18, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 30, 3, 2, 2, 2, 8,
	50, 3, 2, 2, 2, 10, 92, 3, 2, 2, 2, 12, 13, 5, 4, 3, 2, 13, 14, 7, 2, 2,
	3, 14, 19, 3, 2, 2, 2, 15, 16, 5, 6, 4, 2, 16, 17, 7, 2, 2, 3, 17, 19,
	3, 2, 2, 2, 18, 12, 3, 2, 2, 2, 18, 15, 3, 2, 2, 2, 19, 3, 3, 2, 2, 2,
	20, 21, 7, 22, 2, 2, 21, 22, 7, 11, 2, 2, 22, 23, 5, 6, 4, 2, 23, 5, 3,
	2, 2, 2, 24, 25, 8, 4, 1, 2, 25, 26, 7, 18, 2, 2, 26, 27, 5, 6, 4, 2, 27,
	28, 7, 19, 2, 2, 28, 31, 3, 2, 2, 2, 29, 31, 5, 8, 5, 2, 30, 24, 3, 2,
	2, 2, 30, 29, 3, 2, 2, 2, 31, 43, 3, 2, 2, 2, 32, 33, 12, 7, 2, 2, 33,
	34, 7, 12, 2, 2, 34, 42, 5, 6, 4, 8, 35, 36, 12, 6, 2, 2, 36, 37, 9, 2,
	2, 2, 37, 42, 5, 6, 4, 7, 38, 39, 12, 5, 2, 2, 39, 40, 9, 3, 2, 2, 40,
	42, 5, 6, 4, 6, 41, 32, 3, 2, 2, 2, 41, 35, 3, 2, 2, 2, 41, 38, 3, 2, 2,
	2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 7, 3,
	2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 51, 5, 10, 6, 2, 47, 51, 7, 20, 2, 2,
	48, 51, 7, 21, 2, 2, 49, 51, 7, 22, 2, 2, 50, 46, 3, 2, 2, 2, 50, 47, 3,
	2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 49, 3, 2, 2, 2, 51, 9, 3, 2, 2, 2, 52,
	53, 7, 3, 2, 2, 53, 54, 7, 18, 2, 2, 54, 55, 5, 6, 4, 2, 55, 56, 7, 19,
	2, 2, 56, 93, 3, 2, 2, 2, 57, 58, 7, 4, 2, 2, 58, 59, 7, 18, 2, 2, 59,
	60, 5, 6, 4, 2, 60, 61, 7, 19, 2, 2, 61, 93, 3, 2, 2, 2, 62, 63, 7, 5,
	2, 2, 63, 64, 7, 18, 2, 2, 64, 65, 5, 6, 4, 2, 65, 66, 7, 19, 2, 2, 66,
	93, 3, 2, 2, 2, 67, 68, 7, 6, 2, 2, 68, 69, 7, 18, 2, 2, 69, 70, 5, 6,
	4, 2, 70, 71, 7, 19, 2, 2, 71, 93, 3, 2, 2, 2, 72, 73, 7, 7, 2, 2, 73,
	74, 7, 18, 2, 2, 74, 75, 5, 6, 4, 2, 75, 76, 7, 19, 2, 2, 76, 93, 3, 2,
	2, 2, 77, 78, 7, 8, 2, 2, 78, 79, 7, 18, 2, 2, 79, 80, 5, 6, 4, 2, 80,
	81, 7, 19, 2, 2, 81, 93, 3, 2, 2, 2, 82, 83, 7, 9, 2, 2, 83, 84, 7, 18,
	2, 2, 84, 85, 5, 6, 4, 2, 85, 86, 7, 19, 2, 2, 86, 93, 3, 2, 2, 2, 87,
	88, 7, 10, 2, 2, 88, 89, 7, 18, 2, 2, 89, 90, 5, 6, 4, 2, 90, 91, 7, 19,
	2, 2, 91, 93, 3, 2, 2, 2, 92, 52, 3, 2, 2, 2, 92, 57, 3, 2, 2, 2, 92, 62,
	3, 2, 2, 2, 92, 67, 3, 2, 2, 2, 92, 72, 3, 2, 2, 2, 92, 77, 3, 2, 2, 2,
	92, 82, 3, 2, 2, 2, 92, 87, 3, 2, 2, 2, 93, 11, 3, 2, 2, 2, 8, 18, 30,
	41, 43, 50, 92,
}
var literalNames = []string{
	"", "'sqrt'", "'ln'", "'floor'", "'ceil'", "'exp'", "'sin'", "'cos'", "'tan'",
	"'='", "", "'*'", "'/'", "'+'", "'-'", "'.'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "EQ", "POW", "MUL", "DIV", "ADD", "SUB",
	"POINT", "LPAREN", "RPAREN", "CONSTANT", "SCIENTIFIC_NUMBER", "VARIABLE",
	"WHITESPACE",
}

var ruleNames = []string{
	"root", "declaration", "expression", "atom", "function",
}

type AbacusParser struct {
	*antlr.BaseParser
}

// NewAbacusParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *AbacusParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewAbacusParser(input antlr.TokenStream) *AbacusParser {
	this := new(AbacusParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Abacus.g4"

	return this
}

// AbacusParser tokens.
const (
	AbacusParserEOF               = antlr.TokenEOF
	AbacusParserT__0              = 1
	AbacusParserT__1              = 2
	AbacusParserT__2              = 3
	AbacusParserT__3              = 4
	AbacusParserT__4              = 5
	AbacusParserT__5              = 6
	AbacusParserT__6              = 7
	AbacusParserT__7              = 8
	AbacusParserEQ                = 9
	AbacusParserPOW               = 10
	AbacusParserMUL               = 11
	AbacusParserDIV               = 12
	AbacusParserADD               = 13
	AbacusParserSUB               = 14
	AbacusParserPOINT             = 15
	AbacusParserLPAREN            = 16
	AbacusParserRPAREN            = 17
	AbacusParserCONSTANT          = 18
	AbacusParserSCIENTIFIC_NUMBER = 19
	AbacusParserVARIABLE          = 20
	AbacusParserWHITESPACE        = 21
)

// AbacusParser rules.
const (
	AbacusParserRULE_root        = 0
	AbacusParserRULE_declaration = 1
	AbacusParserRULE_expression  = 2
	AbacusParserRULE_atom        = 3
	AbacusParserRULE_function    = 4
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(AbacusParserEOF, 0)
}

func (s *RootContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AbacusParserRULE_root)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.Declaration()
		}
		{
			p.SetState(11)
			p.Match(AbacusParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.expression(0)
		}
		{
			p.SetState(14)
			p.Match(AbacusParserEOF)
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(AbacusParserVARIABLE, 0)
}

func (s *DeclarationContext) EQ() antlr.TerminalNode {
	return s.GetToken(AbacusParserEQ, 0)
}

func (s *DeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AbacusParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Match(AbacusParserVARIABLE)
	}
	{
		p.SetState(19)
		p.Match(AbacusParserEQ)
	}
	{
		p.SetState(20)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulDivContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MulDivContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(AbacusParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(AbacusParserDIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(AbacusParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(AbacusParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowContext struct {
	*ExpressionContext
}

func NewPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowContext {
	var p = new(PowContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *PowContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowContext) POW() antlr.TerminalNode {
	return s.GetToken(AbacusParserPOW, 0)
}

func (s *PowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterPow(s)
	}
}

func (s *PowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitPow(s)
	}
}

func (s *PowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitPow(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtomExprContext struct {
	*ExpressionContext
}

func NewAtomExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomExprContext {
	var p = new(AtomExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AtomExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExprContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterAtomExpr(s)
	}
}

func (s *AtomExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitAtomExpr(s)
	}
}

func (s *AtomExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitAtomExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenthesesContext struct {
	*ExpressionContext
}

func NewParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesesContext {
	var p = new(ParenthesesContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *ParenthesesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *ParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterParentheses(s)
	}
}

func (s *ParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitParentheses(s)
	}
}

func (s *ParenthesesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitParentheses(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *AbacusParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, AbacusParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserLPAREN:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(23)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(24)
			p.expression(0)
		}
		{
			p.SetState(25)
			p.Match(AbacusParserRPAREN)
		}

	case AbacusParserT__0, AbacusParserT__1, AbacusParserT__2, AbacusParserT__3, AbacusParserT__4, AbacusParserT__5, AbacusParserT__6, AbacusParserT__7, AbacusParserCONSTANT, AbacusParserSCIENTIFIC_NUMBER, AbacusParserVARIABLE:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(27)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(30)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(31)
					p.Match(AbacusParserPOW)
				}
				{
					p.SetState(32)
					p.expression(6)
				}

			case 2:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(34)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == AbacusParserMUL || _la == AbacusParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(35)
					p.expression(5)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(37)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == AbacusParserADD || _la == AbacusParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(38)
					p.expression(4)
				}

			}

		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyFrom(ctx *AtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VariableContext struct {
	*AtomContext
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(AbacusParserVARIABLE, 0)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberContext struct {
	*AtomContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) SCIENTIFIC_NUMBER() antlr.TerminalNode {
	return s.GetToken(AbacusParserSCIENTIFIC_NUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstantContext struct {
	*AtomContext
}

func NewConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantContext {
	var p = new(ConstantContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) CONSTANT() antlr.TerminalNode {
	return s.GetToken(AbacusParserCONSTANT, 0)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncExprContext struct {
	*AtomContext
}

func NewFuncExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncExprContext {
	var p = new(FuncExprContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *FuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExprContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FuncExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterFuncExpr(s)
	}
}

func (s *FuncExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitFuncExpr(s)
	}
}

func (s *FuncExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitFuncExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AbacusParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserT__0, AbacusParserT__1, AbacusParserT__2, AbacusParserT__3, AbacusParserT__4, AbacusParserT__5, AbacusParserT__6, AbacusParserT__7:
		localctx = NewFuncExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Function()
		}

	case AbacusParserCONSTANT:
		localctx = NewConstantContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(AbacusParserCONSTANT)
		}

	case AbacusParserSCIENTIFIC_NUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(46)
			p.Match(AbacusParserSCIENTIFIC_NUMBER)
		}

	case AbacusParserVARIABLE:
		localctx = NewVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(47)
			p.Match(AbacusParserVARIABLE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) CopyFrom(ctx *FunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SqrtFunctionContext struct {
	*FunctionContext
}

func NewSqrtFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SqrtFunctionContext {
	var p = new(SqrtFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *SqrtFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqrtFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *SqrtFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SqrtFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *SqrtFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterSqrtFunction(s)
	}
}

func (s *SqrtFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitSqrtFunction(s)
	}
}

func (s *SqrtFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitSqrtFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type LnFunctionContext struct {
	*FunctionContext
}

func NewLnFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LnFunctionContext {
	var p = new(LnFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *LnFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LnFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *LnFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LnFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *LnFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterLnFunction(s)
	}
}

func (s *LnFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitLnFunction(s)
	}
}

func (s *LnFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitLnFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type CosFunctionContext struct {
	*FunctionContext
}

func NewCosFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CosFunctionContext {
	var p = new(CosFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *CosFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CosFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *CosFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CosFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *CosFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterCosFunction(s)
	}
}

func (s *CosFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitCosFunction(s)
	}
}

func (s *CosFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitCosFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type SinFunctionContext struct {
	*FunctionContext
}

func NewSinFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SinFunctionContext {
	var p = new(SinFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *SinFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *SinFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SinFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *SinFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterSinFunction(s)
	}
}

func (s *SinFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitSinFunction(s)
	}
}

func (s *SinFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitSinFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type TanFunctionContext struct {
	*FunctionContext
}

func NewTanFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TanFunctionContext {
	var p = new(TanFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *TanFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TanFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *TanFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TanFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *TanFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterTanFunction(s)
	}
}

func (s *TanFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitTanFunction(s)
	}
}

func (s *TanFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitTanFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type CeilFunctionContext struct {
	*FunctionContext
}

func NewCeilFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CeilFunctionContext {
	var p = new(CeilFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *CeilFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CeilFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *CeilFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CeilFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *CeilFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterCeilFunction(s)
	}
}

func (s *CeilFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitCeilFunction(s)
	}
}

func (s *CeilFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitCeilFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloorFunctionContext struct {
	*FunctionContext
}

func NewFloorFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloorFunctionContext {
	var p = new(FloorFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *FloorFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloorFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *FloorFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FloorFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *FloorFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterFloorFunction(s)
	}
}

func (s *FloorFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitFloorFunction(s)
	}
}

func (s *FloorFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitFloorFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpFunctionContext struct {
	*FunctionContext
}

func NewExpFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpFunctionContext {
	var p = new(ExpFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *ExpFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *ExpFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *ExpFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterExpFunction(s)
	}
}

func (s *ExpFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitExpFunction(s)
	}
}

func (s *ExpFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitExpFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AbacusParserRULE_function)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserT__0:
		localctx = NewSqrtFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(AbacusParserT__0)
		}
		{
			p.SetState(51)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(52)
			p.expression(0)
		}
		{
			p.SetState(53)
			p.Match(AbacusParserRPAREN)
		}

	case AbacusParserT__1:
		localctx = NewLnFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Match(AbacusParserT__1)
		}
		{
			p.SetState(56)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(57)
			p.expression(0)
		}
		{
			p.SetState(58)
			p.Match(AbacusParserRPAREN)
		}

	case AbacusParserT__2:
		localctx = NewFloorFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.Match(AbacusParserT__2)
		}
		{
			p.SetState(61)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(62)
			p.expression(0)
		}
		{
			p.SetState(63)
			p.Match(AbacusParserRPAREN)
		}

	case AbacusParserT__3:
		localctx = NewCeilFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.Match(AbacusParserT__3)
		}
		{
			p.SetState(66)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(67)
			p.expression(0)
		}
		{
			p.SetState(68)
			p.Match(AbacusParserRPAREN)
		}

	case AbacusParserT__4:
		localctx = NewExpFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(70)
			p.Match(AbacusParserT__4)
		}
		{
			p.SetState(71)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(72)
			p.expression(0)
		}
		{
			p.SetState(73)
			p.Match(AbacusParserRPAREN)
		}

	case AbacusParserT__5:
		localctx = NewSinFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(75)
			p.Match(AbacusParserT__5)
		}
		{
			p.SetState(76)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(77)
			p.expression(0)
		}
		{
			p.SetState(78)
			p.Match(AbacusParserRPAREN)
		}

	case AbacusParserT__6:
		localctx = NewCosFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(80)
			p.Match(AbacusParserT__6)
		}
		{
			p.SetState(81)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(82)
			p.expression(0)
		}
		{
			p.SetState(83)
			p.Match(AbacusParserRPAREN)
		}

	case AbacusParserT__7:
		localctx = NewTanFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(85)
			p.Match(AbacusParserT__7)
		}
		{
			p.SetState(86)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(87)
			p.expression(0)
		}
		{
			p.SetState(88)
			p.Match(AbacusParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *AbacusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AbacusParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
