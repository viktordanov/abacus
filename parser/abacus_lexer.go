// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AbacusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var abacuslexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func abacuslexerLexerInit() {
	staticData := &abacuslexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'true'", "'false'", "':'", "','", "'&&'", "'||'", "'xor'", "",
		"'='", "'<'", "'>'", "", "", "'*'", "'/'", "'+'", "'-'", "'%'", "'.'",
		"'('", "')'", "'['", "']'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "AND", "OR", "XOR", "NOT", "EQ", "LS", "GR", "ARROW",
		"POW", "MUL", "DIV", "ADD", "SUB", "PER", "POINT", "LPAREN", "RPAREN",
		"LSQPAREN", "RSQPAREN", "CONSTANT", "SCIENTIFIC_NUMBER", "VARIABLE",
		"LAMBDA_VARIABLE", "DIGITS", "UPPERCASE", "LOWERCASE", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "AND", "OR", "XOR", "NOT", "EQ", "LS",
		"GR", "ARROW", "POW", "MUL", "DIV", "ADD", "SUB", "PER", "POINT", "LPAREN",
		"RPAREN", "LSQPAREN", "RSQPAREN", "CONSTANT", "SCIENTIFIC_NUMBER", "SIGN",
		"NUMBER", "VARIABLE", "LAMBDA_VARIABLE", "VALID_ID_START", "VALID_ID_CHAR",
		"DIGITS", "UPPERCASE", "LOWERCASE", "WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 31, 202, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11,
		109, 8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 114, 8, 12, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 3, 23, 142, 8, 23, 1, 24, 1, 24, 1, 24, 3, 24, 147, 8, 24,
		1, 24, 3, 24, 150, 8, 24, 1, 25, 1, 25, 1, 26, 4, 26, 155, 8, 26, 11, 26,
		12, 26, 156, 1, 26, 1, 26, 4, 26, 161, 8, 26, 11, 26, 12, 26, 162, 3, 26,
		165, 8, 26, 1, 27, 1, 27, 5, 27, 169, 8, 27, 10, 27, 12, 27, 172, 9, 27,
		1, 28, 1, 28, 5, 28, 176, 8, 28, 10, 28, 12, 28, 179, 9, 28, 1, 29, 1,
		29, 3, 29, 183, 8, 29, 1, 30, 1, 30, 1, 30, 3, 30, 188, 8, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 4, 34, 197, 8, 34, 11, 34, 12, 34,
		198, 1, 34, 1, 34, 0, 0, 35, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		0, 53, 0, 55, 26, 57, 27, 59, 0, 61, 0, 63, 28, 65, 29, 67, 30, 69, 31,
		1, 0, 4, 3, 0, 33, 33, 126, 126, 172, 172, 2, 0, 69, 69, 101, 101, 2, 0,
		43, 43, 45, 45, 3, 0, 9, 10, 13, 13, 32, 32, 212, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 1, 71, 1, 0, 0, 0,
		3, 76, 1, 0, 0, 0, 5, 82, 1, 0, 0, 0, 7, 84, 1, 0, 0, 0, 9, 86, 1, 0, 0,
		0, 11, 89, 1, 0, 0, 0, 13, 92, 1, 0, 0, 0, 15, 96, 1, 0, 0, 0, 17, 98,
		1, 0, 0, 0, 19, 100, 1, 0, 0, 0, 21, 102, 1, 0, 0, 0, 23, 108, 1, 0, 0,
		0, 25, 113, 1, 0, 0, 0, 27, 115, 1, 0, 0, 0, 29, 117, 1, 0, 0, 0, 31, 119,
		1, 0, 0, 0, 33, 121, 1, 0, 0, 0, 35, 123, 1, 0, 0, 0, 37, 125, 1, 0, 0,
		0, 39, 127, 1, 0, 0, 0, 41, 129, 1, 0, 0, 0, 43, 131, 1, 0, 0, 0, 45, 133,
		1, 0, 0, 0, 47, 141, 1, 0, 0, 0, 49, 143, 1, 0, 0, 0, 51, 151, 1, 0, 0,
		0, 53, 154, 1, 0, 0, 0, 55, 166, 1, 0, 0, 0, 57, 173, 1, 0, 0, 0, 59, 182,
		1, 0, 0, 0, 61, 187, 1, 0, 0, 0, 63, 189, 1, 0, 0, 0, 65, 191, 1, 0, 0,
		0, 67, 193, 1, 0, 0, 0, 69, 196, 1, 0, 0, 0, 71, 72, 5, 116, 0, 0, 72,
		73, 5, 114, 0, 0, 73, 74, 5, 117, 0, 0, 74, 75, 5, 101, 0, 0, 75, 2, 1,
		0, 0, 0, 76, 77, 5, 102, 0, 0, 77, 78, 5, 97, 0, 0, 78, 79, 5, 108, 0,
		0, 79, 80, 5, 115, 0, 0, 80, 81, 5, 101, 0, 0, 81, 4, 1, 0, 0, 0, 82, 83,
		5, 58, 0, 0, 83, 6, 1, 0, 0, 0, 84, 85, 5, 44, 0, 0, 85, 8, 1, 0, 0, 0,
		86, 87, 5, 38, 0, 0, 87, 88, 5, 38, 0, 0, 88, 10, 1, 0, 0, 0, 89, 90, 5,
		124, 0, 0, 90, 91, 5, 124, 0, 0, 91, 12, 1, 0, 0, 0, 92, 93, 5, 120, 0,
		0, 93, 94, 5, 111, 0, 0, 94, 95, 5, 114, 0, 0, 95, 14, 1, 0, 0, 0, 96,
		97, 7, 0, 0, 0, 97, 16, 1, 0, 0, 0, 98, 99, 5, 61, 0, 0, 99, 18, 1, 0,
		0, 0, 100, 101, 5, 60, 0, 0, 101, 20, 1, 0, 0, 0, 102, 103, 5, 62, 0, 0,
		103, 22, 1, 0, 0, 0, 104, 105, 5, 45, 0, 0, 105, 109, 5, 62, 0, 0, 106,
		107, 5, 61, 0, 0, 107, 109, 5, 62, 0, 0, 108, 104, 1, 0, 0, 0, 108, 106,
		1, 0, 0, 0, 109, 24, 1, 0, 0, 0, 110, 114, 5, 94, 0, 0, 111, 112, 5, 42,
		0, 0, 112, 114, 5, 42, 0, 0, 113, 110, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0,
		114, 26, 1, 0, 0, 0, 115, 116, 5, 42, 0, 0, 116, 28, 1, 0, 0, 0, 117, 118,
		5, 47, 0, 0, 118, 30, 1, 0, 0, 0, 119, 120, 5, 43, 0, 0, 120, 32, 1, 0,
		0, 0, 121, 122, 5, 45, 0, 0, 122, 34, 1, 0, 0, 0, 123, 124, 5, 37, 0, 0,
		124, 36, 1, 0, 0, 0, 125, 126, 5, 46, 0, 0, 126, 38, 1, 0, 0, 0, 127, 128,
		5, 40, 0, 0, 128, 40, 1, 0, 0, 0, 129, 130, 5, 41, 0, 0, 130, 42, 1, 0,
		0, 0, 131, 132, 5, 91, 0, 0, 132, 44, 1, 0, 0, 0, 133, 134, 5, 93, 0, 0,
		134, 46, 1, 0, 0, 0, 135, 136, 5, 112, 0, 0, 136, 142, 5, 105, 0, 0, 137,
		142, 5, 101, 0, 0, 138, 139, 5, 112, 0, 0, 139, 140, 5, 104, 0, 0, 140,
		142, 5, 105, 0, 0, 141, 135, 1, 0, 0, 0, 141, 137, 1, 0, 0, 0, 141, 138,
		1, 0, 0, 0, 142, 48, 1, 0, 0, 0, 143, 149, 3, 53, 26, 0, 144, 146, 7, 1,
		0, 0, 145, 147, 3, 51, 25, 0, 146, 145, 1, 0, 0, 0, 146, 147, 1, 0, 0,
		0, 147, 148, 1, 0, 0, 0, 148, 150, 3, 53, 26, 0, 149, 144, 1, 0, 0, 0,
		149, 150, 1, 0, 0, 0, 150, 50, 1, 0, 0, 0, 151, 152, 7, 2, 0, 0, 152, 52,
		1, 0, 0, 0, 153, 155, 3, 63, 31, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1,
		0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 164, 1, 0, 0,
		0, 158, 160, 3, 37, 18, 0, 159, 161, 3, 63, 31, 0, 160, 159, 1, 0, 0, 0,
		161, 162, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163,
		165, 1, 0, 0, 0, 164, 158, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 54, 1,
		0, 0, 0, 166, 170, 3, 59, 29, 0, 167, 169, 3, 61, 30, 0, 168, 167, 1, 0,
		0, 0, 169, 172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0,
		171, 56, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 177, 3, 65, 32, 0, 174,
		176, 3, 61, 30, 0, 175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175,
		1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 58, 1, 0, 0, 0, 179, 177, 1, 0,
		0, 0, 180, 183, 3, 67, 33, 0, 181, 183, 5, 95, 0, 0, 182, 180, 1, 0, 0,
		0, 182, 181, 1, 0, 0, 0, 183, 60, 1, 0, 0, 0, 184, 188, 3, 59, 29, 0, 185,
		188, 3, 65, 32, 0, 186, 188, 3, 63, 31, 0, 187, 184, 1, 0, 0, 0, 187, 185,
		1, 0, 0, 0, 187, 186, 1, 0, 0, 0, 188, 62, 1, 0, 0, 0, 189, 190, 2, 48,
		57, 0, 190, 64, 1, 0, 0, 0, 191, 192, 2, 65, 90, 0, 192, 66, 1, 0, 0, 0,
		193, 194, 2, 97, 122, 0, 194, 68, 1, 0, 0, 0, 195, 197, 7, 3, 0, 0, 196,
		195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198, 199,
		1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 6, 34, 0, 0, 201, 70, 1, 0,
		0, 0, 14, 0, 108, 113, 141, 146, 149, 156, 162, 164, 170, 177, 182, 187,
		198, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AbacusLexerInit initializes any static state used to implement AbacusLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAbacusLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AbacusLexerInit() {
	staticData := &abacuslexerLexerStaticData
	staticData.once.Do(abacuslexerLexerInit)
}

// NewAbacusLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAbacusLexer(input antlr.CharStream) *AbacusLexer {
	AbacusLexerInit()
	l := new(AbacusLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &abacuslexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
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
	AbacusLexerAND               = 5
	AbacusLexerOR                = 6
	AbacusLexerXOR               = 7
	AbacusLexerNOT               = 8
	AbacusLexerEQ                = 9
	AbacusLexerLS                = 10
	AbacusLexerGR                = 11
	AbacusLexerARROW             = 12
	AbacusLexerPOW               = 13
	AbacusLexerMUL               = 14
	AbacusLexerDIV               = 15
	AbacusLexerADD               = 16
	AbacusLexerSUB               = 17
	AbacusLexerPER               = 18
	AbacusLexerPOINT             = 19
	AbacusLexerLPAREN            = 20
	AbacusLexerRPAREN            = 21
	AbacusLexerLSQPAREN          = 22
	AbacusLexerRSQPAREN          = 23
	AbacusLexerCONSTANT          = 24
	AbacusLexerSCIENTIFIC_NUMBER = 25
	AbacusLexerVARIABLE          = 26
	AbacusLexerLAMBDA_VARIABLE   = 27
	AbacusLexerDIGITS            = 28
	AbacusLexerUPPERCASE         = 29
	AbacusLexerLOWERCASE         = 30
	AbacusLexerWHITESPACE        = 31
)
