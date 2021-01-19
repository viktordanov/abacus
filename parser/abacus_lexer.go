// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 50, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 5, 2, 25, 10,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 6, 9, 40, 10, 9, 13, 9, 14, 9, 41, 3, 10, 6, 10, 45, 10, 10, 13,
	10, 14, 10, 46, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 3, 2, 4, 3, 2, 50, 59, 5, 2, 11, 12, 15,
	15, 34, 34, 2, 52, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 24, 3, 2, 2, 2, 5, 26, 3, 2,
	2, 2, 7, 28, 3, 2, 2, 2, 9, 30, 3, 2, 2, 2, 11, 32, 3, 2, 2, 2, 13, 34,
	3, 2, 2, 2, 15, 36, 3, 2, 2, 2, 17, 39, 3, 2, 2, 2, 19, 44, 3, 2, 2, 2,
	21, 25, 7, 96, 2, 2, 22, 23, 7, 44, 2, 2, 23, 25, 7, 44, 2, 2, 24, 21,
	3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 25, 4, 3, 2, 2, 2, 26, 27, 7, 44, 2, 2,
	27, 6, 3, 2, 2, 2, 28, 29, 7, 49, 2, 2, 29, 8, 3, 2, 2, 2, 30, 31, 7, 45,
	2, 2, 31, 10, 3, 2, 2, 2, 32, 33, 7, 47, 2, 2, 33, 12, 3, 2, 2, 2, 34,
	35, 7, 42, 2, 2, 35, 14, 3, 2, 2, 2, 36, 37, 7, 43, 2, 2, 37, 16, 3, 2,
	2, 2, 38, 40, 9, 2, 2, 2, 39, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 39,
	3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 18, 3, 2, 2, 2, 43, 45, 9, 3, 2, 2,
	44, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3,
	2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 49, 8, 10, 2, 2, 49, 20, 3, 2, 2, 2, 6,
	2, 24, 41, 46, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "POW", "MUL", "DIV", "ADD", "SUB", "LPAREN", "RPAREN", "NUMBER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"POW", "MUL", "DIV", "ADD", "SUB", "LPAREN", "RPAREN", "NUMBER", "WHITESPACE",
}

type AbacusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewAbacusLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *AbacusLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewAbacusLexer(input antlr.CharStream) *AbacusLexer {
	l := new(AbacusLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Abacus.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AbacusLexer tokens.
const (
	AbacusLexerPOW        = 1
	AbacusLexerMUL        = 2
	AbacusLexerDIV        = 3
	AbacusLexerADD        = 4
	AbacusLexerSUB        = 5
	AbacusLexerLPAREN     = 6
	AbacusLexerRPAREN     = 7
	AbacusLexerNUMBER     = 8
	AbacusLexerWHITESPACE = 9
)
