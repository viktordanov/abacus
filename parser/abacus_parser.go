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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 181,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 24, 10, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 48, 10,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 57, 10, 4, 3, 4, 3,
	4, 5, 4, 61, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 69, 10, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 80, 10, 5,
	12, 5, 14, 5, 83, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 89, 10, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 179, 10, 7, 3, 7, 2, 3, 8, 8, 2, 4, 6, 8, 10, 12,
	2, 4, 3, 2, 22, 23, 3, 2, 24, 25, 2, 204, 2, 23, 3, 2, 2, 2, 4, 25, 3,
	2, 2, 2, 6, 60, 3, 2, 2, 2, 8, 68, 3, 2, 2, 2, 10, 88, 3, 2, 2, 2, 12,
	178, 3, 2, 2, 2, 14, 15, 5, 4, 3, 2, 15, 16, 7, 2, 2, 3, 16, 24, 3, 2,
	2, 2, 17, 18, 5, 6, 4, 2, 18, 19, 7, 2, 2, 3, 19, 24, 3, 2, 2, 2, 20, 21,
	5, 8, 5, 2, 21, 22, 7, 2, 2, 3, 22, 24, 3, 2, 2, 2, 23, 14, 3, 2, 2, 2,
	23, 17, 3, 2, 2, 2, 23, 20, 3, 2, 2, 2, 24, 3, 3, 2, 2, 2, 25, 26, 7, 31,
	2, 2, 26, 27, 7, 18, 2, 2, 27, 28, 5, 8, 5, 2, 28, 5, 3, 2, 2, 2, 29, 30,
	5, 8, 5, 2, 30, 31, 7, 18, 2, 2, 31, 32, 7, 18, 2, 2, 32, 33, 5, 8, 5,
	2, 33, 61, 3, 2, 2, 2, 34, 35, 5, 8, 5, 2, 35, 36, 7, 19, 2, 2, 36, 37,
	5, 8, 5, 2, 37, 61, 3, 2, 2, 2, 38, 39, 5, 8, 5, 2, 39, 40, 7, 20, 2, 2,
	40, 41, 5, 8, 5, 2, 41, 61, 3, 2, 2, 2, 42, 47, 5, 8, 5, 2, 43, 44, 7,
	19, 2, 2, 44, 48, 7, 18, 2, 2, 45, 46, 7, 18, 2, 2, 46, 48, 7, 19, 2, 2,
	47, 43, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50, 5,
	8, 5, 2, 50, 61, 3, 2, 2, 2, 51, 56, 5, 8, 5, 2, 52, 53, 7, 20, 2, 2, 53,
	57, 7, 18, 2, 2, 54, 55, 7, 18, 2, 2, 55, 57, 7, 20, 2, 2, 56, 52, 3, 2,
	2, 2, 56, 54, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 5, 8, 5, 2, 59, 61,
	3, 2, 2, 2, 60, 29, 3, 2, 2, 2, 60, 34, 3, 2, 2, 2, 60, 38, 3, 2, 2, 2,
	60, 42, 3, 2, 2, 2, 60, 51, 3, 2, 2, 2, 61, 7, 3, 2, 2, 2, 62, 63, 8, 5,
	1, 2, 63, 64, 7, 27, 2, 2, 64, 65, 5, 8, 5, 2, 65, 66, 7, 28, 2, 2, 66,
	69, 3, 2, 2, 2, 67, 69, 5, 10, 6, 2, 68, 62, 3, 2, 2, 2, 68, 67, 3, 2,
	2, 2, 69, 81, 3, 2, 2, 2, 70, 71, 12, 7, 2, 2, 71, 72, 7, 21, 2, 2, 72,
	80, 5, 8, 5, 8, 73, 74, 12, 6, 2, 2, 74, 75, 9, 2, 2, 2, 75, 80, 5, 8,
	5, 7, 76, 77, 12, 5, 2, 2, 77, 78, 9, 3, 2, 2, 78, 80, 5, 8, 5, 6, 79,
	70, 3, 2, 2, 2, 79, 73, 3, 2, 2, 2, 79, 76, 3, 2, 2, 2, 80, 83, 3, 2, 2,
	2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 9, 3, 2, 2, 2, 83, 81, 3,
	2, 2, 2, 84, 89, 5, 12, 7, 2, 85, 89, 7, 29, 2, 2, 86, 89, 7, 30, 2, 2,
	87, 89, 7, 31, 2, 2, 88, 84, 3, 2, 2, 2, 88, 85, 3, 2, 2, 2, 88, 86, 3,
	2, 2, 2, 88, 87, 3, 2, 2, 2, 89, 11, 3, 2, 2, 2, 90, 91, 7, 3, 2, 2, 91,
	92, 7, 27, 2, 2, 92, 93, 5, 8, 5, 2, 93, 94, 7, 28, 2, 2, 94, 179, 3, 2,
	2, 2, 95, 96, 7, 4, 2, 2, 96, 97, 7, 27, 2, 2, 97, 98, 5, 8, 5, 2, 98,
	99, 7, 28, 2, 2, 99, 179, 3, 2, 2, 2, 100, 101, 7, 5, 2, 2, 101, 102, 7,
	27, 2, 2, 102, 103, 5, 8, 5, 2, 103, 104, 7, 28, 2, 2, 104, 179, 3, 2,
	2, 2, 105, 106, 7, 6, 2, 2, 106, 107, 7, 27, 2, 2, 107, 108, 5, 8, 5, 2,
	108, 109, 7, 28, 2, 2, 109, 179, 3, 2, 2, 2, 110, 111, 7, 7, 2, 2, 111,
	112, 7, 27, 2, 2, 112, 113, 5, 8, 5, 2, 113, 114, 7, 28, 2, 2, 114, 179,
	3, 2, 2, 2, 115, 116, 7, 8, 2, 2, 116, 117, 7, 27, 2, 2, 117, 118, 5, 8,
	5, 2, 118, 119, 7, 28, 2, 2, 119, 179, 3, 2, 2, 2, 120, 121, 7, 9, 2, 2,
	121, 122, 7, 27, 2, 2, 122, 123, 5, 8, 5, 2, 123, 124, 7, 28, 2, 2, 124,
	179, 3, 2, 2, 2, 125, 126, 7, 10, 2, 2, 126, 127, 7, 27, 2, 2, 127, 128,
	5, 8, 5, 2, 128, 129, 7, 28, 2, 2, 129, 179, 3, 2, 2, 2, 130, 131, 7, 11,
	2, 2, 131, 132, 7, 27, 2, 2, 132, 133, 5, 8, 5, 2, 133, 134, 7, 28, 2,
	2, 134, 179, 3, 2, 2, 2, 135, 136, 7, 12, 2, 2, 136, 137, 7, 27, 2, 2,
	137, 138, 5, 8, 5, 2, 138, 139, 7, 28, 2, 2, 139, 179, 3, 2, 2, 2, 140,
	141, 7, 13, 2, 2, 141, 142, 7, 27, 2, 2, 142, 143, 5, 8, 5, 2, 143, 144,
	7, 28, 2, 2, 144, 179, 3, 2, 2, 2, 145, 146, 7, 14, 2, 2, 146, 147, 7,
	27, 2, 2, 147, 148, 5, 8, 5, 2, 148, 149, 7, 28, 2, 2, 149, 179, 3, 2,
	2, 2, 150, 151, 7, 14, 2, 2, 151, 152, 7, 27, 2, 2, 152, 153, 5, 8, 5,
	2, 153, 154, 7, 15, 2, 2, 154, 155, 5, 8, 5, 2, 155, 156, 7, 28, 2, 2,
	156, 179, 3, 2, 2, 2, 157, 158, 7, 5, 2, 2, 158, 159, 7, 27, 2, 2, 159,
	160, 5, 8, 5, 2, 160, 161, 7, 15, 2, 2, 161, 162, 5, 8, 5, 2, 162, 163,
	7, 28, 2, 2, 163, 179, 3, 2, 2, 2, 164, 165, 7, 16, 2, 2, 165, 166, 7,
	27, 2, 2, 166, 167, 5, 8, 5, 2, 167, 168, 7, 15, 2, 2, 168, 169, 5, 8,
	5, 2, 169, 170, 7, 28, 2, 2, 170, 179, 3, 2, 2, 2, 171, 172, 7, 17, 2,
	2, 172, 173, 7, 27, 2, 2, 173, 174, 5, 8, 5, 2, 174, 175, 7, 15, 2, 2,
	175, 176, 5, 8, 5, 2, 176, 177, 7, 28, 2, 2, 177, 179, 3, 2, 2, 2, 178,
	90, 3, 2, 2, 2, 178, 95, 3, 2, 2, 2, 178, 100, 3, 2, 2, 2, 178, 105, 3,
	2, 2, 2, 178, 110, 3, 2, 2, 2, 178, 115, 3, 2, 2, 2, 178, 120, 3, 2, 2,
	2, 178, 125, 3, 2, 2, 2, 178, 130, 3, 2, 2, 2, 178, 135, 3, 2, 2, 2, 178,
	140, 3, 2, 2, 2, 178, 145, 3, 2, 2, 2, 178, 150, 3, 2, 2, 2, 178, 157,
	3, 2, 2, 2, 178, 164, 3, 2, 2, 2, 178, 171, 3, 2, 2, 2, 179, 13, 3, 2,
	2, 2, 11, 23, 47, 56, 60, 68, 79, 81, 88, 178,
}
var literalNames = []string{
	"", "'sqrt'", "'ln'", "'log'", "'log2'", "'log10'", "'floor'", "'ceil'",
	"'exp'", "'sin'", "'cos'", "'tan'", "'round'", "','", "'min'", "'max'",
	"'='", "'<'", "'>'", "", "'*'", "'/'", "'+'", "'-'", "'.'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "EQ", "LS",
	"GR", "POW", "MUL", "DIV", "ADD", "SUB", "POINT", "LPAREN", "RPAREN", "CONSTANT",
	"SCIENTIFIC_NUMBER", "VARIABLE", "WHITESPACE",
}

var ruleNames = []string{
	"root", "declaration", "comparison", "expression", "atom", "function",
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
	AbacusParserT__8              = 9
	AbacusParserT__9              = 10
	AbacusParserT__10             = 11
	AbacusParserT__11             = 12
	AbacusParserT__12             = 13
	AbacusParserT__13             = 14
	AbacusParserT__14             = 15
	AbacusParserEQ                = 16
	AbacusParserLS                = 17
	AbacusParserGR                = 18
	AbacusParserPOW               = 19
	AbacusParserMUL               = 20
	AbacusParserDIV               = 21
	AbacusParserADD               = 22
	AbacusParserSUB               = 23
	AbacusParserPOINT             = 24
	AbacusParserLPAREN            = 25
	AbacusParserRPAREN            = 26
	AbacusParserCONSTANT          = 27
	AbacusParserSCIENTIFIC_NUMBER = 28
	AbacusParserVARIABLE          = 29
	AbacusParserWHITESPACE        = 30
)

// AbacusParser rules.
const (
	AbacusParserRULE_root        = 0
	AbacusParserRULE_declaration = 1
	AbacusParserRULE_comparison  = 2
	AbacusParserRULE_expression  = 3
	AbacusParserRULE_atom        = 4
	AbacusParserRULE_function    = 5
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

func (s *RootContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
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

	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)
			p.Declaration()
		}
		{
			p.SetState(13)
			p.Match(AbacusParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(15)
			p.Comparison()
		}
		{
			p.SetState(16)
			p.Match(AbacusParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(18)
			p.expression(0)
		}
		{
			p.SetState(19)
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
		p.SetState(23)
		p.Match(AbacusParserVARIABLE)
	}
	{
		p.SetState(24)
		p.Match(AbacusParserEQ)
	}
	{
		p.SetState(25)
		p.expression(0)
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) CopyFrom(ctx *ComparisonContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GreaterOrEqualComparisonContext struct {
	*ComparisonContext
}

func NewGreaterOrEqualComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterOrEqualComparisonContext {
	var p = new(GreaterOrEqualComparisonContext)

	p.ComparisonContext = NewEmptyComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparisonContext))

	return p
}

func (s *GreaterOrEqualComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterOrEqualComparisonContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *GreaterOrEqualComparisonContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GreaterOrEqualComparisonContext) GR() antlr.TerminalNode {
	return s.GetToken(AbacusParserGR, 0)
}

func (s *GreaterOrEqualComparisonContext) EQ() antlr.TerminalNode {
	return s.GetToken(AbacusParserEQ, 0)
}

func (s *GreaterOrEqualComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterGreaterOrEqualComparison(s)
	}
}

func (s *GreaterOrEqualComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitGreaterOrEqualComparison(s)
	}
}

func (s *GreaterOrEqualComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitGreaterOrEqualComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type GreaterComparisonContext struct {
	*ComparisonContext
}

func NewGreaterComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterComparisonContext {
	var p = new(GreaterComparisonContext)

	p.ComparisonContext = NewEmptyComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparisonContext))

	return p
}

func (s *GreaterComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterComparisonContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *GreaterComparisonContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GreaterComparisonContext) GR() antlr.TerminalNode {
	return s.GetToken(AbacusParserGR, 0)
}

func (s *GreaterComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterGreaterComparison(s)
	}
}

func (s *GreaterComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitGreaterComparison(s)
	}
}

func (s *GreaterComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitGreaterComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type LessOrEqualComparisonContext struct {
	*ComparisonContext
}

func NewLessOrEqualComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessOrEqualComparisonContext {
	var p = new(LessOrEqualComparisonContext)

	p.ComparisonContext = NewEmptyComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparisonContext))

	return p
}

func (s *LessOrEqualComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessOrEqualComparisonContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LessOrEqualComparisonContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LessOrEqualComparisonContext) LS() antlr.TerminalNode {
	return s.GetToken(AbacusParserLS, 0)
}

func (s *LessOrEqualComparisonContext) EQ() antlr.TerminalNode {
	return s.GetToken(AbacusParserEQ, 0)
}

func (s *LessOrEqualComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterLessOrEqualComparison(s)
	}
}

func (s *LessOrEqualComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitLessOrEqualComparison(s)
	}
}

func (s *LessOrEqualComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitLessOrEqualComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type LessComparisonContext struct {
	*ComparisonContext
}

func NewLessComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessComparisonContext {
	var p = new(LessComparisonContext)

	p.ComparisonContext = NewEmptyComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparisonContext))

	return p
}

func (s *LessComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessComparisonContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LessComparisonContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LessComparisonContext) LS() antlr.TerminalNode {
	return s.GetToken(AbacusParserLS, 0)
}

func (s *LessComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterLessComparison(s)
	}
}

func (s *LessComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitLessComparison(s)
	}
}

func (s *LessComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitLessComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualComparisonContext struct {
	*ComparisonContext
}

func NewEqualComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualComparisonContext {
	var p = new(EqualComparisonContext)

	p.ComparisonContext = NewEmptyComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparisonContext))

	return p
}

func (s *EqualComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualComparisonContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqualComparisonContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualComparisonContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(AbacusParserEQ)
}

func (s *EqualComparisonContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(AbacusParserEQ, i)
}

func (s *EqualComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterEqualComparison(s)
	}
}

func (s *EqualComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitEqualComparison(s)
	}
}

func (s *EqualComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitEqualComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AbacusParserRULE_comparison)

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

	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEqualComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.expression(0)
		}
		{
			p.SetState(28)
			p.Match(AbacusParserEQ)
		}
		{
			p.SetState(29)
			p.Match(AbacusParserEQ)
		}
		{
			p.SetState(30)
			p.expression(0)
		}

	case 2:
		localctx = NewLessComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.expression(0)
		}
		{
			p.SetState(33)
			p.Match(AbacusParserLS)
		}
		{
			p.SetState(34)
			p.expression(0)
		}

	case 3:
		localctx = NewGreaterComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.expression(0)
		}
		{
			p.SetState(37)
			p.Match(AbacusParserGR)
		}
		{
			p.SetState(38)
			p.expression(0)
		}

	case 4:
		localctx = NewLessOrEqualComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.expression(0)
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AbacusParserLS:
			{
				p.SetState(41)
				p.Match(AbacusParserLS)
			}
			{
				p.SetState(42)
				p.Match(AbacusParserEQ)
			}

		case AbacusParserEQ:
			{
				p.SetState(43)
				p.Match(AbacusParserEQ)
			}
			{
				p.SetState(44)
				p.Match(AbacusParserLS)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(47)
			p.expression(0)
		}

	case 5:
		localctx = NewGreaterOrEqualComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(49)
			p.expression(0)
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AbacusParserGR:
			{
				p.SetState(50)
				p.Match(AbacusParserGR)
			}
			{
				p.SetState(51)
				p.Match(AbacusParserEQ)
			}

		case AbacusParserEQ:
			{
				p.SetState(52)
				p.Match(AbacusParserEQ)
			}
			{
				p.SetState(53)
				p.Match(AbacusParserGR)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(56)
			p.expression(0)
		}

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
	_startState := 6
	p.EnterRecursionRule(localctx, 6, AbacusParserRULE_expression, _p)
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
	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserLPAREN:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

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

	case AbacusParserT__0, AbacusParserT__1, AbacusParserT__2, AbacusParserT__3, AbacusParserT__4, AbacusParserT__5, AbacusParserT__6, AbacusParserT__7, AbacusParserT__8, AbacusParserT__9, AbacusParserT__10, AbacusParserT__11, AbacusParserT__13, AbacusParserT__14, AbacusParserCONSTANT, AbacusParserSCIENTIFIC_NUMBER, AbacusParserVARIABLE:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(69)
					p.Match(AbacusParserPOW)
				}
				{
					p.SetState(70)
					p.expression(6)
				}

			case 2:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(72)

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
					p.SetState(73)
					p.expression(5)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(75)

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
					p.SetState(76)
					p.expression(4)
				}

			}

		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 8, AbacusParserRULE_atom)

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

	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserT__0, AbacusParserT__1, AbacusParserT__2, AbacusParserT__3, AbacusParserT__4, AbacusParserT__5, AbacusParserT__6, AbacusParserT__7, AbacusParserT__8, AbacusParserT__9, AbacusParserT__10, AbacusParserT__11, AbacusParserT__13, AbacusParserT__14:
		localctx = NewFuncExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Function()
		}

	case AbacusParserCONSTANT:
		localctx = NewConstantContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Match(AbacusParserCONSTANT)
		}

	case AbacusParserSCIENTIFIC_NUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.Match(AbacusParserSCIENTIFIC_NUMBER)
		}

	case AbacusParserVARIABLE:
		localctx = NewVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(85)
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

type Log10FunctionContext struct {
	*FunctionContext
}

func NewLog10FunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Log10FunctionContext {
	var p = new(Log10FunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *Log10FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Log10FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *Log10FunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Log10FunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *Log10FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterLog10Function(s)
	}
}

func (s *Log10FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitLog10Function(s)
	}
}

func (s *Log10FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitLog10Function(s)

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

type MaxFunctionContext struct {
	*FunctionContext
}

func NewMaxFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MaxFunctionContext {
	var p = new(MaxFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *MaxFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaxFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *MaxFunctionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MaxFunctionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MaxFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *MaxFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterMaxFunction(s)
	}
}

func (s *MaxFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitMaxFunction(s)
	}
}

func (s *MaxFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitMaxFunction(s)

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

type MinFunctionContext struct {
	*FunctionContext
}

func NewMinFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinFunctionContext {
	var p = new(MinFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *MinFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *MinFunctionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MinFunctionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MinFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *MinFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterMinFunction(s)
	}
}

func (s *MinFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitMinFunction(s)
	}
}

func (s *MinFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitMinFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogDefFunctionContext struct {
	*FunctionContext
}

func NewLogDefFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogDefFunctionContext {
	var p = new(LogDefFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *LogDefFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogDefFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *LogDefFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogDefFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *LogDefFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterLogDefFunction(s)
	}
}

func (s *LogDefFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitLogDefFunction(s)
	}
}

func (s *LogDefFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitLogDefFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogFunctionContext struct {
	*FunctionContext
}

func NewLogFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogFunctionContext {
	var p = new(LogFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *LogFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *LogFunctionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LogFunctionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *LogFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterLogFunction(s)
	}
}

func (s *LogFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitLogFunction(s)
	}
}

func (s *LogFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitLogFunction(s)

	default:
		return t.VisitChildren(s)
	}
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

type Round2FunctionContext struct {
	*FunctionContext
}

func NewRound2FunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Round2FunctionContext {
	var p = new(Round2FunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *Round2FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Round2FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *Round2FunctionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Round2FunctionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Round2FunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *Round2FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterRound2Function(s)
	}
}

func (s *Round2FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitRound2Function(s)
	}
}

func (s *Round2FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitRound2Function(s)

	default:
		return t.VisitChildren(s)
	}
}

type Log2FunctionContext struct {
	*FunctionContext
}

func NewLog2FunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Log2FunctionContext {
	var p = new(Log2FunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *Log2FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Log2FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *Log2FunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Log2FunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *Log2FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterLog2Function(s)
	}
}

func (s *Log2FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitLog2Function(s)
	}
}

func (s *Log2FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitLog2Function(s)

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

type RoundDefFunctionContext struct {
	*FunctionContext
}

func NewRoundDefFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoundDefFunctionContext {
	var p = new(RoundDefFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *RoundDefFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundDefFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *RoundDefFunctionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RoundDefFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *RoundDefFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterRoundDefFunction(s)
	}
}

func (s *RoundDefFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitRoundDefFunction(s)
	}
}

func (s *RoundDefFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitRoundDefFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AbacusParserRULE_function)

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

	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSqrtFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(AbacusParserT__0)
		}
		{
			p.SetState(89)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(90)
			p.expression(0)
		}
		{
			p.SetState(91)
			p.Match(AbacusParserRPAREN)
		}

	case 2:
		localctx = NewLnFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(AbacusParserT__1)
		}
		{
			p.SetState(94)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(95)
			p.expression(0)
		}
		{
			p.SetState(96)
			p.Match(AbacusParserRPAREN)
		}

	case 3:
		localctx = NewLogDefFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Match(AbacusParserT__2)
		}
		{
			p.SetState(99)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(100)
			p.expression(0)
		}
		{
			p.SetState(101)
			p.Match(AbacusParserRPAREN)
		}

	case 4:
		localctx = NewLog2FunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.Match(AbacusParserT__3)
		}
		{
			p.SetState(104)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(105)
			p.expression(0)
		}
		{
			p.SetState(106)
			p.Match(AbacusParserRPAREN)
		}

	case 5:
		localctx = NewLog10FunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(108)
			p.Match(AbacusParserT__4)
		}
		{
			p.SetState(109)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(110)
			p.expression(0)
		}
		{
			p.SetState(111)
			p.Match(AbacusParserRPAREN)
		}

	case 6:
		localctx = NewFloorFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(113)
			p.Match(AbacusParserT__5)
		}
		{
			p.SetState(114)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(115)
			p.expression(0)
		}
		{
			p.SetState(116)
			p.Match(AbacusParserRPAREN)
		}

	case 7:
		localctx = NewCeilFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(118)
			p.Match(AbacusParserT__6)
		}
		{
			p.SetState(119)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(120)
			p.expression(0)
		}
		{
			p.SetState(121)
			p.Match(AbacusParserRPAREN)
		}

	case 8:
		localctx = NewExpFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(123)
			p.Match(AbacusParserT__7)
		}
		{
			p.SetState(124)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(125)
			p.expression(0)
		}
		{
			p.SetState(126)
			p.Match(AbacusParserRPAREN)
		}

	case 9:
		localctx = NewSinFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(128)
			p.Match(AbacusParserT__8)
		}
		{
			p.SetState(129)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(130)
			p.expression(0)
		}
		{
			p.SetState(131)
			p.Match(AbacusParserRPAREN)
		}

	case 10:
		localctx = NewCosFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(133)
			p.Match(AbacusParserT__9)
		}
		{
			p.SetState(134)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(135)
			p.expression(0)
		}
		{
			p.SetState(136)
			p.Match(AbacusParserRPAREN)
		}

	case 11:
		localctx = NewTanFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(138)
			p.Match(AbacusParserT__10)
		}
		{
			p.SetState(139)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(140)
			p.expression(0)
		}
		{
			p.SetState(141)
			p.Match(AbacusParserRPAREN)
		}

	case 12:
		localctx = NewRoundDefFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(143)
			p.Match(AbacusParserT__11)
		}
		{
			p.SetState(144)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(145)
			p.expression(0)
		}
		{
			p.SetState(146)
			p.Match(AbacusParserRPAREN)
		}

	case 13:
		localctx = NewRound2FunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(148)
			p.Match(AbacusParserT__11)
		}
		{
			p.SetState(149)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(150)
			p.expression(0)
		}
		{
			p.SetState(151)
			p.Match(AbacusParserT__12)
		}
		{
			p.SetState(152)
			p.expression(0)
		}
		{
			p.SetState(153)
			p.Match(AbacusParserRPAREN)
		}

	case 14:
		localctx = NewLogFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(155)
			p.Match(AbacusParserT__2)
		}
		{
			p.SetState(156)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(157)
			p.expression(0)
		}
		{
			p.SetState(158)
			p.Match(AbacusParserT__12)
		}
		{
			p.SetState(159)
			p.expression(0)
		}
		{
			p.SetState(160)
			p.Match(AbacusParserRPAREN)
		}

	case 15:
		localctx = NewMinFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(162)
			p.Match(AbacusParserT__13)
		}
		{
			p.SetState(163)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(164)
			p.expression(0)
		}
		{
			p.SetState(165)
			p.Match(AbacusParserT__12)
		}
		{
			p.SetState(166)
			p.expression(0)
		}
		{
			p.SetState(167)
			p.Match(AbacusParserRPAREN)
		}

	case 16:
		localctx = NewMaxFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(169)
			p.Match(AbacusParserT__14)
		}
		{
			p.SetState(170)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(171)
			p.expression(0)
		}
		{
			p.SetState(172)
			p.Match(AbacusParserT__12)
		}
		{
			p.SetState(173)
			p.expression(0)
		}
		{
			p.SetState(174)
			p.Match(AbacusParserRPAREN)
		}

	}

	return localctx
}

func (p *AbacusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
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
