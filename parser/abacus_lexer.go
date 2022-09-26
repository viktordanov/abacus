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
		"LSQPAREN", "RSQPAREN", "CONSTANT", "SCIENTIFIC_NUMBER", "NUMBER", "VARIABLE",
		"LAMBDA_VARIABLE", "DIGITS", "UPPERCASE", "LOWERCASE", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "AND", "OR", "XOR", "NOT", "EQ", "LS",
		"GR", "ARROW", "POW", "MUL", "DIV", "ADD", "SUB", "PER", "POINT", "LPAREN",
		"RPAREN", "LSQPAREN", "RSQPAREN", "CONSTANT", "SCIENTIFIC_NUMBER", "NUMBER",
		"WHOLE_NUMBER", "DECIMAL_NUMBER", "SIGN", "VARIABLE", "LAMBDA_VARIABLE",
		"VALID_ID_START", "VALID_ID_CHAR", "DIGITS", "UPPERCASE", "LOWERCASE",
		"WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 216, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 113, 8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 118, 8,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 146, 8, 23, 1, 24, 1, 24,
		1, 24, 3, 24, 151, 8, 24, 1, 24, 3, 24, 154, 8, 24, 1, 25, 1, 25, 3, 25,
		158, 8, 25, 1, 26, 4, 26, 161, 8, 26, 11, 26, 12, 26, 162, 1, 26, 1, 26,
		4, 26, 167, 8, 26, 11, 26, 12, 26, 168, 3, 26, 171, 8, 26, 1, 27, 1, 27,
		4, 27, 175, 8, 27, 11, 27, 12, 27, 176, 1, 28, 1, 28, 1, 29, 1, 29, 5,
		29, 183, 8, 29, 10, 29, 12, 29, 186, 9, 29, 1, 30, 1, 30, 5, 30, 190, 8,
		30, 10, 30, 12, 30, 193, 9, 30, 1, 31, 1, 31, 3, 31, 197, 8, 31, 1, 32,
		1, 32, 1, 32, 3, 32, 202, 8, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 36, 4, 36, 211, 8, 36, 11, 36, 12, 36, 212, 1, 36, 1, 36, 0, 0,
		37, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 0, 55, 0, 57, 0,
		59, 27, 61, 28, 63, 0, 65, 0, 67, 29, 69, 30, 71, 31, 73, 32, 1, 0, 4,
		3, 0, 33, 33, 126, 126, 172, 172, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43,
		45, 45, 3, 0, 9, 10, 13, 13, 32, 32, 227, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 67, 1, 0,
		0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 1, 75, 1,
		0, 0, 0, 3, 80, 1, 0, 0, 0, 5, 86, 1, 0, 0, 0, 7, 88, 1, 0, 0, 0, 9, 90,
		1, 0, 0, 0, 11, 93, 1, 0, 0, 0, 13, 96, 1, 0, 0, 0, 15, 100, 1, 0, 0, 0,
		17, 102, 1, 0, 0, 0, 19, 104, 1, 0, 0, 0, 21, 106, 1, 0, 0, 0, 23, 112,
		1, 0, 0, 0, 25, 117, 1, 0, 0, 0, 27, 119, 1, 0, 0, 0, 29, 121, 1, 0, 0,
		0, 31, 123, 1, 0, 0, 0, 33, 125, 1, 0, 0, 0, 35, 127, 1, 0, 0, 0, 37, 129,
		1, 0, 0, 0, 39, 131, 1, 0, 0, 0, 41, 133, 1, 0, 0, 0, 43, 135, 1, 0, 0,
		0, 45, 137, 1, 0, 0, 0, 47, 145, 1, 0, 0, 0, 49, 147, 1, 0, 0, 0, 51, 157,
		1, 0, 0, 0, 53, 160, 1, 0, 0, 0, 55, 172, 1, 0, 0, 0, 57, 178, 1, 0, 0,
		0, 59, 180, 1, 0, 0, 0, 61, 187, 1, 0, 0, 0, 63, 196, 1, 0, 0, 0, 65, 201,
		1, 0, 0, 0, 67, 203, 1, 0, 0, 0, 69, 205, 1, 0, 0, 0, 71, 207, 1, 0, 0,
		0, 73, 210, 1, 0, 0, 0, 75, 76, 5, 116, 0, 0, 76, 77, 5, 114, 0, 0, 77,
		78, 5, 117, 0, 0, 78, 79, 5, 101, 0, 0, 79, 2, 1, 0, 0, 0, 80, 81, 5, 102,
		0, 0, 81, 82, 5, 97, 0, 0, 82, 83, 5, 108, 0, 0, 83, 84, 5, 115, 0, 0,
		84, 85, 5, 101, 0, 0, 85, 4, 1, 0, 0, 0, 86, 87, 5, 58, 0, 0, 87, 6, 1,
		0, 0, 0, 88, 89, 5, 44, 0, 0, 89, 8, 1, 0, 0, 0, 90, 91, 5, 38, 0, 0, 91,
		92, 5, 38, 0, 0, 92, 10, 1, 0, 0, 0, 93, 94, 5, 124, 0, 0, 94, 95, 5, 124,
		0, 0, 95, 12, 1, 0, 0, 0, 96, 97, 5, 120, 0, 0, 97, 98, 5, 111, 0, 0, 98,
		99, 5, 114, 0, 0, 99, 14, 1, 0, 0, 0, 100, 101, 7, 0, 0, 0, 101, 16, 1,
		0, 0, 0, 102, 103, 5, 61, 0, 0, 103, 18, 1, 0, 0, 0, 104, 105, 5, 60, 0,
		0, 105, 20, 1, 0, 0, 0, 106, 107, 5, 62, 0, 0, 107, 22, 1, 0, 0, 0, 108,
		109, 5, 45, 0, 0, 109, 113, 5, 62, 0, 0, 110, 111, 5, 61, 0, 0, 111, 113,
		5, 62, 0, 0, 112, 108, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 24, 1, 0,
		0, 0, 114, 118, 5, 94, 0, 0, 115, 116, 5, 42, 0, 0, 116, 118, 5, 42, 0,
		0, 117, 114, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 26, 1, 0, 0, 0, 119,
		120, 5, 42, 0, 0, 120, 28, 1, 0, 0, 0, 121, 122, 5, 47, 0, 0, 122, 30,
		1, 0, 0, 0, 123, 124, 5, 43, 0, 0, 124, 32, 1, 0, 0, 0, 125, 126, 5, 45,
		0, 0, 126, 34, 1, 0, 0, 0, 127, 128, 5, 37, 0, 0, 128, 36, 1, 0, 0, 0,
		129, 130, 5, 46, 0, 0, 130, 38, 1, 0, 0, 0, 131, 132, 5, 40, 0, 0, 132,
		40, 1, 0, 0, 0, 133, 134, 5, 41, 0, 0, 134, 42, 1, 0, 0, 0, 135, 136, 5,
		91, 0, 0, 136, 44, 1, 0, 0, 0, 137, 138, 5, 93, 0, 0, 138, 46, 1, 0, 0,
		0, 139, 140, 5, 112, 0, 0, 140, 146, 5, 105, 0, 0, 141, 146, 5, 101, 0,
		0, 142, 143, 5, 112, 0, 0, 143, 144, 5, 104, 0, 0, 144, 146, 5, 105, 0,
		0, 145, 139, 1, 0, 0, 0, 145, 141, 1, 0, 0, 0, 145, 142, 1, 0, 0, 0, 146,
		48, 1, 0, 0, 0, 147, 153, 3, 51, 25, 0, 148, 150, 7, 1, 0, 0, 149, 151,
		3, 57, 28, 0, 150, 149, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 1,
		0, 0, 0, 152, 154, 3, 51, 25, 0, 153, 148, 1, 0, 0, 0, 153, 154, 1, 0,
		0, 0, 154, 50, 1, 0, 0, 0, 155, 158, 3, 53, 26, 0, 156, 158, 3, 55, 27,
		0, 157, 155, 1, 0, 0, 0, 157, 156, 1, 0, 0, 0, 158, 52, 1, 0, 0, 0, 159,
		161, 3, 67, 33, 0, 160, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 160,
		1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 170, 1, 0, 0, 0, 164, 166, 3, 37,
		18, 0, 165, 167, 3, 67, 33, 0, 166, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0,
		0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 171, 1, 0, 0, 0, 170,
		164, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 54, 1, 0, 0, 0, 172, 174, 3,
		37, 18, 0, 173, 175, 3, 67, 33, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0,
		0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 56, 1, 0, 0, 0,
		178, 179, 7, 2, 0, 0, 179, 58, 1, 0, 0, 0, 180, 184, 3, 63, 31, 0, 181,
		183, 3, 65, 32, 0, 182, 181, 1, 0, 0, 0, 183, 186, 1, 0, 0, 0, 184, 182,
		1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 60, 1, 0, 0, 0, 186, 184, 1, 0,
		0, 0, 187, 191, 3, 69, 34, 0, 188, 190, 3, 65, 32, 0, 189, 188, 1, 0, 0,
		0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192,
		62, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 197, 3, 71, 35, 0, 195, 197,
		5, 95, 0, 0, 196, 194, 1, 0, 0, 0, 196, 195, 1, 0, 0, 0, 197, 64, 1, 0,
		0, 0, 198, 202, 3, 63, 31, 0, 199, 202, 3, 69, 34, 0, 200, 202, 3, 67,
		33, 0, 201, 198, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 200, 1, 0, 0, 0,
		202, 66, 1, 0, 0, 0, 203, 204, 2, 48, 57, 0, 204, 68, 1, 0, 0, 0, 205,
		206, 2, 65, 90, 0, 206, 70, 1, 0, 0, 0, 207, 208, 2, 97, 122, 0, 208, 72,
		1, 0, 0, 0, 209, 211, 7, 3, 0, 0, 210, 209, 1, 0, 0, 0, 211, 212, 1, 0,
		0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0,
		214, 215, 6, 36, 0, 0, 215, 74, 1, 0, 0, 0, 16, 0, 112, 117, 145, 150,
		153, 157, 162, 168, 170, 176, 184, 191, 196, 201, 212, 1, 6, 0, 0,
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
	AbacusLexerNUMBER            = 26
	AbacusLexerVARIABLE          = 27
	AbacusLexerLAMBDA_VARIABLE   = 28
	AbacusLexerDIGITS            = 29
	AbacusLexerUPPERCASE         = 30
	AbacusLexerLOWERCASE         = 31
	AbacusLexerWHITESPACE        = 32
)
