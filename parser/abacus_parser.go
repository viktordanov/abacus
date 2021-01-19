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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 38, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 18, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 4, 3, 4, 5,
	4, 36, 10, 4, 3, 4, 2, 3, 4, 5, 2, 4, 6, 2, 4, 3, 2, 4, 5, 3, 2, 6, 7,
	2, 39, 2, 8, 3, 2, 2, 2, 4, 17, 3, 2, 2, 2, 6, 35, 3, 2, 2, 2, 8, 9, 5,
	4, 3, 2, 9, 10, 7, 2, 2, 3, 10, 3, 3, 2, 2, 2, 11, 12, 8, 3, 1, 2, 12,
	13, 7, 9, 2, 2, 13, 14, 5, 4, 3, 2, 14, 15, 7, 10, 2, 2, 15, 18, 3, 2,
	2, 2, 16, 18, 5, 6, 4, 2, 17, 11, 3, 2, 2, 2, 17, 16, 3, 2, 2, 2, 18, 30,
	3, 2, 2, 2, 19, 20, 12, 7, 2, 2, 20, 21, 7, 3, 2, 2, 21, 29, 5, 4, 3, 8,
	22, 23, 12, 6, 2, 2, 23, 24, 9, 2, 2, 2, 24, 29, 5, 4, 3, 7, 25, 26, 12,
	5, 2, 2, 26, 27, 9, 3, 2, 2, 27, 29, 5, 4, 3, 6, 28, 19, 3, 2, 2, 2, 28,
	22, 3, 2, 2, 2, 28, 25, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3, 2, 2,
	2, 30, 31, 3, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 33, 36, 7,
	11, 2, 2, 34, 36, 7, 12, 2, 2, 35, 33, 3, 2, 2, 2, 35, 34, 3, 2, 2, 2,
	36, 7, 3, 2, 2, 2, 6, 17, 28, 30, 35,
}
var literalNames = []string{
	"", "", "'*'", "'/'", "'+'", "'-'", "'.'", "'('", "')'",
}
var symbolicNames = []string{
	"", "POW", "MUL", "DIV", "ADD", "SUB", "POINT", "LPAREN", "RPAREN", "SCIENTIFIC_NUMBER",
	"VARIABLE", "WHITESPACE",
}

var ruleNames = []string{
	"start", "expression", "atom",
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
	AbacusParserPOW               = 1
	AbacusParserMUL               = 2
	AbacusParserDIV               = 3
	AbacusParserADD               = 4
	AbacusParserSUB               = 5
	AbacusParserPOINT             = 6
	AbacusParserLPAREN            = 7
	AbacusParserRPAREN            = 8
	AbacusParserSCIENTIFIC_NUMBER = 9
	AbacusParserVARIABLE          = 10
	AbacusParserWHITESPACE        = 11
)

// AbacusParser rules.
const (
	AbacusParserRULE_start      = 0
	AbacusParserRULE_expression = 1
	AbacusParserRULE_atom       = 2
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(AbacusParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AbacusParserRULE_start)

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
		p.SetState(6)
		p.expression(0)
	}
	{
		p.SetState(7)
		p.Match(AbacusParserEOF)
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
	_startState := 2
	p.EnterRecursionRule(localctx, 2, AbacusParserRULE_expression, _p)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserLPAREN:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(10)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(11)
			p.expression(0)
		}
		{
			p.SetState(12)
			p.Match(AbacusParserRPAREN)
		}

	case AbacusParserSCIENTIFIC_NUMBER, AbacusParserVARIABLE:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(14)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(17)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(18)
					p.Match(AbacusParserPOW)
				}
				{
					p.SetState(19)
					p.expression(6)
				}

			case 2:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(21)

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
					p.SetState(22)
					p.expression(5)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(24)

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
					p.SetState(25)
					p.expression(4)
				}

			}

		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
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

func (p *AbacusParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AbacusParserRULE_atom)

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

	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserSCIENTIFIC_NUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Match(AbacusParserSCIENTIFIC_NUMBER)
		}

	case AbacusParserVARIABLE:
		localctx = NewVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Match(AbacusParserVARIABLE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *AbacusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
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
