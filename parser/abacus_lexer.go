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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 23, 151,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11,
	92, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 5, 19, 114, 10, 19, 3, 20, 3, 20, 3, 21, 6, 21, 119, 10, 21, 13,
	21, 14, 21, 120, 3, 21, 3, 21, 6, 21, 125, 10, 21, 13, 21, 14, 21, 126,
	5, 21, 129, 10, 21, 3, 22, 3, 22, 7, 22, 133, 10, 22, 12, 22, 14, 22, 136,
	11, 22, 3, 23, 5, 23, 139, 10, 23, 3, 24, 3, 24, 5, 24, 143, 10, 24, 3,
	25, 6, 25, 146, 10, 25, 13, 25, 14, 25, 147, 3, 25, 3, 25, 2, 2, 26, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 2,
	43, 22, 45, 2, 47, 2, 49, 23, 3, 2, 4, 5, 2, 67, 92, 97, 97, 99, 124, 5,
	2, 11, 12, 15, 15, 34, 34, 2, 156, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2,
	2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3,
	2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37,
	3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 3,
	51, 3, 2, 2, 2, 5, 56, 3, 2, 2, 2, 7, 59, 3, 2, 2, 2, 9, 65, 3, 2, 2, 2,
	11, 70, 3, 2, 2, 2, 13, 74, 3, 2, 2, 2, 15, 78, 3, 2, 2, 2, 17, 82, 3,
	2, 2, 2, 19, 86, 3, 2, 2, 2, 21, 91, 3, 2, 2, 2, 23, 93, 3, 2, 2, 2, 25,
	95, 3, 2, 2, 2, 27, 97, 3, 2, 2, 2, 29, 99, 3, 2, 2, 2, 31, 101, 3, 2,
	2, 2, 33, 103, 3, 2, 2, 2, 35, 105, 3, 2, 2, 2, 37, 113, 3, 2, 2, 2, 39,
	115, 3, 2, 2, 2, 41, 118, 3, 2, 2, 2, 43, 130, 3, 2, 2, 2, 45, 138, 3,
	2, 2, 2, 47, 142, 3, 2, 2, 2, 49, 145, 3, 2, 2, 2, 51, 52, 7, 117, 2, 2,
	52, 53, 7, 115, 2, 2, 53, 54, 7, 116, 2, 2, 54, 55, 7, 118, 2, 2, 55, 4,
	3, 2, 2, 2, 56, 57, 7, 110, 2, 2, 57, 58, 7, 112, 2, 2, 58, 6, 3, 2, 2,
	2, 59, 60, 7, 104, 2, 2, 60, 61, 7, 110, 2, 2, 61, 62, 7, 113, 2, 2, 62,
	63, 7, 113, 2, 2, 63, 64, 7, 116, 2, 2, 64, 8, 3, 2, 2, 2, 65, 66, 7, 101,
	2, 2, 66, 67, 7, 103, 2, 2, 67, 68, 7, 107, 2, 2, 68, 69, 7, 110, 2, 2,
	69, 10, 3, 2, 2, 2, 70, 71, 7, 103, 2, 2, 71, 72, 7, 122, 2, 2, 72, 73,
	7, 114, 2, 2, 73, 12, 3, 2, 2, 2, 74, 75, 7, 117, 2, 2, 75, 76, 7, 107,
	2, 2, 76, 77, 7, 112, 2, 2, 77, 14, 3, 2, 2, 2, 78, 79, 7, 101, 2, 2, 79,
	80, 7, 113, 2, 2, 80, 81, 7, 117, 2, 2, 81, 16, 3, 2, 2, 2, 82, 83, 7,
	118, 2, 2, 83, 84, 7, 99, 2, 2, 84, 85, 7, 112, 2, 2, 85, 18, 3, 2, 2,
	2, 86, 87, 7, 63, 2, 2, 87, 20, 3, 2, 2, 2, 88, 92, 7, 96, 2, 2, 89, 90,
	7, 44, 2, 2, 90, 92, 7, 44, 2, 2, 91, 88, 3, 2, 2, 2, 91, 89, 3, 2, 2,
	2, 92, 22, 3, 2, 2, 2, 93, 94, 7, 44, 2, 2, 94, 24, 3, 2, 2, 2, 95, 96,
	7, 49, 2, 2, 96, 26, 3, 2, 2, 2, 97, 98, 7, 45, 2, 2, 98, 28, 3, 2, 2,
	2, 99, 100, 7, 47, 2, 2, 100, 30, 3, 2, 2, 2, 101, 102, 7, 48, 2, 2, 102,
	32, 3, 2, 2, 2, 103, 104, 7, 42, 2, 2, 104, 34, 3, 2, 2, 2, 105, 106, 7,
	43, 2, 2, 106, 36, 3, 2, 2, 2, 107, 108, 7, 114, 2, 2, 108, 114, 7, 107,
	2, 2, 109, 114, 7, 103, 2, 2, 110, 111, 7, 114, 2, 2, 111, 112, 7, 106,
	2, 2, 112, 114, 7, 107, 2, 2, 113, 107, 3, 2, 2, 2, 113, 109, 3, 2, 2,
	2, 113, 110, 3, 2, 2, 2, 114, 38, 3, 2, 2, 2, 115, 116, 5, 41, 21, 2, 116,
	40, 3, 2, 2, 2, 117, 119, 4, 50, 59, 2, 118, 117, 3, 2, 2, 2, 119, 120,
	3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 128, 3, 2,
	2, 2, 122, 124, 5, 31, 16, 2, 123, 125, 4, 50, 59, 2, 124, 123, 3, 2, 2,
	2, 125, 126, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127,
	129, 3, 2, 2, 2, 128, 122, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 42, 3,
	2, 2, 2, 130, 134, 5, 45, 23, 2, 131, 133, 5, 47, 24, 2, 132, 131, 3, 2,
	2, 2, 133, 136, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2,
	135, 44, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 137, 139, 9, 2, 2, 2, 138, 137,
	3, 2, 2, 2, 139, 46, 3, 2, 2, 2, 140, 143, 5, 45, 23, 2, 141, 143, 4, 50,
	59, 2, 142, 140, 3, 2, 2, 2, 142, 141, 3, 2, 2, 2, 143, 48, 3, 2, 2, 2,
	144, 146, 9, 3, 2, 2, 145, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147,
	145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 150,
	8, 25, 2, 2, 150, 50, 3, 2, 2, 2, 12, 2, 91, 113, 120, 126, 128, 134, 138,
	142, 147, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'sqrt'", "'ln'", "'floor'", "'ceil'", "'exp'", "'sin'", "'cos'", "'tan'",
	"'='", "", "'*'", "'/'", "'+'", "'-'", "'.'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "EQ", "POW", "MUL", "DIV", "ADD", "SUB",
	"POINT", "LPAREN", "RPAREN", "CONSTANT", "SCIENTIFIC_NUMBER", "VARIABLE",
	"WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "EQ", "POW",
	"MUL", "DIV", "ADD", "SUB", "POINT", "LPAREN", "RPAREN", "CONSTANT", "SCIENTIFIC_NUMBER",
	"NUMBER", "VARIABLE", "VALID_ID_START", "VALID_ID_CHAR", "WHITESPACE",
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
	AbacusLexerT__0              = 1
	AbacusLexerT__1              = 2
	AbacusLexerT__2              = 3
	AbacusLexerT__3              = 4
	AbacusLexerT__4              = 5
	AbacusLexerT__5              = 6
	AbacusLexerT__6              = 7
	AbacusLexerT__7              = 8
	AbacusLexerEQ                = 9
	AbacusLexerPOW               = 10
	AbacusLexerMUL               = 11
	AbacusLexerDIV               = 12
	AbacusLexerADD               = 13
	AbacusLexerSUB               = 14
	AbacusLexerPOINT             = 15
	AbacusLexerLPAREN            = 16
	AbacusLexerRPAREN            = 17
	AbacusLexerCONSTANT          = 18
	AbacusLexerSCIENTIFIC_NUMBER = 19
	AbacusLexerVARIABLE          = 20
	AbacusLexerWHITESPACE        = 21
)
