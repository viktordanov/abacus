// Code generated from Abacus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Abacus

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AbacusParser struct {
	*antlr.BaseParser
}

var abacusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func abacusParserInit() {
	staticData := &abacusParserStaticData
	staticData.literalNames = []string{
		"", "'true'", "'false'", "':'", "','", "'sqrt'", "'cbrt'", "'ln'", "'log'",
		"'log2'", "'log10'", "'floor'", "'ceil'", "'exp'", "'sin'", "'cos'",
		"'tan'", "'round'", "'sign'", "'abs'", "'min'", "'max'", "'avg'", "'until'",
		"'from'", "'reverse'", "'nth'", "'&&'", "'||'", "'xor'", "", "'='",
		"'<'", "'>'", "", "", "'*'", "'/'", "'+'", "'-'", "'%'", "'.'", "'('",
		"')'", "'['", "']'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "AND", "OR", "XOR", "NOT", "EQ",
		"LS", "GR", "ARROW", "POW", "MUL", "DIV", "ADD", "SUB", "PER", "POINT",
		"LPAREN", "RPAREN", "LSQPAREN", "RSQPAREN", "CONSTANT", "SCIENTIFIC_NUMBER",
		"VARIABLE", "LAMBDA_VARIABLE", "DIGITS", "UPPERCASE", "LOWERCASE", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"root", "declaration", "boolExpression", "boolAtom", "lambda", "expression",
		"parameter", "recursionParameters", "mixedTuple", "tuple", "lambdaArguments",
		"variablesTuple", "atom", "sign", "function",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 53, 354, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 42, 8, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 51, 8, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 72, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 81, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 92, 8, 2, 1, 2, 1, 2, 1, 2, 5, 2, 97, 8, 2, 10, 2, 12, 2, 100,
		9, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		112, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 3, 5, 125, 8, 5, 1, 5, 1, 5, 3, 5, 129, 8, 5, 1, 5, 3, 5, 132, 8,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 5, 5, 149, 8, 5, 10, 5, 12, 5, 152, 9, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 158, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 164, 8, 7,
		10, 7, 12, 7, 167, 9, 7, 3, 7, 169, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8,
		175, 8, 8, 1, 8, 1, 8, 3, 8, 179, 8, 8, 1, 9, 1, 9, 1, 9, 3, 9, 184, 8,
		9, 1, 10, 1, 10, 1, 10, 3, 10, 189, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 195, 8, 10, 1, 10, 3, 10, 198, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11,
		203, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 209, 8, 11, 1, 11, 3, 11,
		212, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 218, 8, 12, 1, 13, 1, 13,
		3, 13, 222, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 352, 8, 14, 1, 14, 0, 2, 4, 10,
		15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 0, 6, 1, 0,
		27, 29, 1, 0, 1, 2, 1, 0, 36, 37, 1, 0, 38, 39, 2, 0, 3, 3, 31, 31, 1,
		0, 48, 49, 402, 0, 41, 1, 0, 0, 0, 2, 50, 1, 0, 0, 0, 4, 91, 1, 0, 0, 0,
		6, 101, 1, 0, 0, 0, 8, 111, 1, 0, 0, 0, 10, 131, 1, 0, 0, 0, 12, 153, 1,
		0, 0, 0, 14, 159, 1, 0, 0, 0, 16, 174, 1, 0, 0, 0, 18, 180, 1, 0, 0, 0,
		20, 197, 1, 0, 0, 0, 22, 211, 1, 0, 0, 0, 24, 217, 1, 0, 0, 0, 26, 221,
		1, 0, 0, 0, 28, 351, 1, 0, 0, 0, 30, 31, 3, 2, 1, 0, 31, 32, 5, 0, 0, 1,
		32, 42, 1, 0, 0, 0, 33, 34, 3, 4, 2, 0, 34, 35, 5, 0, 0, 1, 35, 42, 1,
		0, 0, 0, 36, 37, 3, 18, 9, 0, 37, 38, 5, 0, 0, 1, 38, 42, 1, 0, 0, 0, 39,
		40, 5, 49, 0, 0, 40, 42, 5, 0, 0, 1, 41, 30, 1, 0, 0, 0, 41, 33, 1, 0,
		0, 0, 41, 36, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 1, 1, 0, 0, 0, 43, 44,
		3, 22, 11, 0, 44, 45, 5, 31, 0, 0, 45, 46, 3, 18, 9, 0, 46, 51, 1, 0, 0,
		0, 47, 48, 5, 49, 0, 0, 48, 49, 5, 31, 0, 0, 49, 51, 3, 8, 4, 0, 50, 43,
		1, 0, 0, 0, 50, 47, 1, 0, 0, 0, 51, 3, 1, 0, 0, 0, 52, 53, 6, 2, -1, 0,
		53, 54, 3, 10, 5, 0, 54, 55, 5, 31, 0, 0, 55, 56, 5, 31, 0, 0, 56, 57,
		3, 10, 5, 0, 57, 92, 1, 0, 0, 0, 58, 59, 3, 10, 5, 0, 59, 60, 5, 32, 0,
		0, 60, 61, 3, 10, 5, 0, 61, 92, 1, 0, 0, 0, 62, 63, 3, 10, 5, 0, 63, 64,
		5, 33, 0, 0, 64, 65, 3, 10, 5, 0, 65, 92, 1, 0, 0, 0, 66, 71, 3, 10, 5,
		0, 67, 68, 5, 32, 0, 0, 68, 72, 5, 31, 0, 0, 69, 70, 5, 31, 0, 0, 70, 72,
		5, 32, 0, 0, 71, 67, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0,
		73, 74, 3, 10, 5, 0, 74, 92, 1, 0, 0, 0, 75, 80, 3, 10, 5, 0, 76, 77, 5,
		33, 0, 0, 77, 81, 5, 31, 0, 0, 78, 79, 5, 31, 0, 0, 79, 81, 5, 33, 0, 0,
		80, 76, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 3,
		10, 5, 0, 83, 92, 1, 0, 0, 0, 84, 85, 5, 30, 0, 0, 85, 92, 3, 4, 2, 3,
		86, 87, 5, 42, 0, 0, 87, 88, 3, 4, 2, 0, 88, 89, 5, 43, 0, 0, 89, 92, 1,
		0, 0, 0, 90, 92, 3, 6, 3, 0, 91, 52, 1, 0, 0, 0, 91, 58, 1, 0, 0, 0, 91,
		62, 1, 0, 0, 0, 91, 66, 1, 0, 0, 0, 91, 75, 1, 0, 0, 0, 91, 84, 1, 0, 0,
		0, 91, 86, 1, 0, 0, 0, 91, 90, 1, 0, 0, 0, 92, 98, 1, 0, 0, 0, 93, 94,
		10, 4, 0, 0, 94, 95, 7, 0, 0, 0, 95, 97, 3, 4, 2, 5, 96, 93, 1, 0, 0, 0,
		97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 5, 1,
		0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 102, 7, 1, 0, 0, 102, 7, 1, 0, 0, 0,
		103, 104, 3, 20, 10, 0, 104, 105, 5, 34, 0, 0, 105, 106, 3, 18, 9, 0, 106,
		112, 1, 0, 0, 0, 107, 108, 5, 42, 0, 0, 108, 109, 5, 43, 0, 0, 109, 110,
		5, 34, 0, 0, 110, 112, 3, 18, 9, 0, 111, 103, 1, 0, 0, 0, 111, 107, 1,
		0, 0, 0, 112, 9, 1, 0, 0, 0, 113, 114, 6, 5, -1, 0, 114, 115, 3, 26, 13,
		0, 115, 116, 3, 10, 5, 9, 116, 132, 1, 0, 0, 0, 117, 118, 5, 42, 0, 0,
		118, 119, 3, 10, 5, 0, 119, 120, 5, 43, 0, 0, 120, 132, 1, 0, 0, 0, 121,
		122, 5, 49, 0, 0, 122, 124, 5, 42, 0, 0, 123, 125, 3, 16, 8, 0, 124, 123,
		1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 128, 5, 43,
		0, 0, 127, 129, 3, 14, 7, 0, 128, 127, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0,
		129, 132, 1, 0, 0, 0, 130, 132, 3, 24, 12, 0, 131, 113, 1, 0, 0, 0, 131,
		117, 1, 0, 0, 0, 131, 121, 1, 0, 0, 0, 131, 130, 1, 0, 0, 0, 132, 150,
		1, 0, 0, 0, 133, 134, 10, 7, 0, 0, 134, 135, 5, 35, 0, 0, 135, 149, 3,
		10, 5, 8, 136, 137, 10, 6, 0, 0, 137, 138, 5, 40, 0, 0, 138, 139, 5, 40,
		0, 0, 139, 149, 3, 10, 5, 7, 140, 141, 10, 5, 0, 0, 141, 142, 7, 2, 0,
		0, 142, 149, 3, 10, 5, 6, 143, 144, 10, 4, 0, 0, 144, 145, 7, 3, 0, 0,
		145, 149, 3, 10, 5, 5, 146, 147, 10, 8, 0, 0, 147, 149, 5, 40, 0, 0, 148,
		133, 1, 0, 0, 0, 148, 136, 1, 0, 0, 0, 148, 140, 1, 0, 0, 0, 148, 143,
		1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0,
		0, 0, 150, 151, 1, 0, 0, 0, 151, 11, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0,
		153, 154, 5, 48, 0, 0, 154, 157, 7, 4, 0, 0, 155, 158, 3, 4, 2, 0, 156,
		158, 3, 10, 5, 0, 157, 155, 1, 0, 0, 0, 157, 156, 1, 0, 0, 0, 158, 13,
		1, 0, 0, 0, 159, 168, 5, 44, 0, 0, 160, 165, 3, 12, 6, 0, 161, 162, 5,
		4, 0, 0, 162, 164, 3, 12, 6, 0, 163, 161, 1, 0, 0, 0, 164, 167, 1, 0, 0,
		0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167,
		165, 1, 0, 0, 0, 168, 160, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170,
		1, 0, 0, 0, 170, 171, 5, 45, 0, 0, 171, 15, 1, 0, 0, 0, 172, 175, 3, 10,
		5, 0, 173, 175, 5, 49, 0, 0, 174, 172, 1, 0, 0, 0, 174, 173, 1, 0, 0, 0,
		175, 178, 1, 0, 0, 0, 176, 177, 5, 4, 0, 0, 177, 179, 3, 16, 8, 0, 178,
		176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 17, 1, 0, 0, 0, 180, 183, 3,
		10, 5, 0, 181, 182, 5, 4, 0, 0, 182, 184, 3, 18, 9, 0, 183, 181, 1, 0,
		0, 0, 183, 184, 1, 0, 0, 0, 184, 19, 1, 0, 0, 0, 185, 188, 7, 5, 0, 0,
		186, 187, 5, 4, 0, 0, 187, 189, 3, 20, 10, 0, 188, 186, 1, 0, 0, 0, 188,
		189, 1, 0, 0, 0, 189, 198, 1, 0, 0, 0, 190, 191, 5, 42, 0, 0, 191, 194,
		7, 5, 0, 0, 192, 193, 5, 4, 0, 0, 193, 195, 3, 20, 10, 0, 194, 192, 1,
		0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 198, 5, 43, 0,
		0, 197, 185, 1, 0, 0, 0, 197, 190, 1, 0, 0, 0, 198, 21, 1, 0, 0, 0, 199,
		202, 5, 48, 0, 0, 200, 201, 5, 4, 0, 0, 201, 203, 3, 22, 11, 0, 202, 200,
		1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 212, 1, 0, 0, 0, 204, 205, 5, 42,
		0, 0, 205, 208, 5, 48, 0, 0, 206, 207, 5, 4, 0, 0, 207, 209, 3, 22, 11,
		0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210,
		212, 5, 43, 0, 0, 211, 199, 1, 0, 0, 0, 211, 204, 1, 0, 0, 0, 212, 23,
		1, 0, 0, 0, 213, 218, 3, 28, 14, 0, 214, 218, 5, 46, 0, 0, 215, 218, 5,
		47, 0, 0, 216, 218, 5, 48, 0, 0, 217, 213, 1, 0, 0, 0, 217, 214, 1, 0,
		0, 0, 217, 215, 1, 0, 0, 0, 217, 216, 1, 0, 0, 0, 218, 25, 1, 0, 0, 0,
		219, 222, 5, 38, 0, 0, 220, 222, 5, 39, 0, 0, 221, 219, 1, 0, 0, 0, 221,
		220, 1, 0, 0, 0, 222, 27, 1, 0, 0, 0, 223, 224, 5, 5, 0, 0, 224, 225, 5,
		42, 0, 0, 225, 226, 3, 10, 5, 0, 226, 227, 5, 43, 0, 0, 227, 352, 1, 0,
		0, 0, 228, 229, 5, 6, 0, 0, 229, 230, 5, 42, 0, 0, 230, 231, 3, 10, 5,
		0, 231, 232, 5, 43, 0, 0, 232, 352, 1, 0, 0, 0, 233, 234, 5, 7, 0, 0, 234,
		235, 5, 42, 0, 0, 235, 236, 3, 10, 5, 0, 236, 237, 5, 43, 0, 0, 237, 352,
		1, 0, 0, 0, 238, 239, 5, 8, 0, 0, 239, 240, 5, 42, 0, 0, 240, 241, 3, 10,
		5, 0, 241, 242, 5, 43, 0, 0, 242, 352, 1, 0, 0, 0, 243, 244, 5, 9, 0, 0,
		244, 245, 5, 42, 0, 0, 245, 246, 3, 10, 5, 0, 246, 247, 5, 43, 0, 0, 247,
		352, 1, 0, 0, 0, 248, 249, 5, 10, 0, 0, 249, 250, 5, 42, 0, 0, 250, 251,
		3, 10, 5, 0, 251, 252, 5, 43, 0, 0, 252, 352, 1, 0, 0, 0, 253, 254, 5,
		11, 0, 0, 254, 255, 5, 42, 0, 0, 255, 256, 3, 10, 5, 0, 256, 257, 5, 43,
		0, 0, 257, 352, 1, 0, 0, 0, 258, 259, 5, 12, 0, 0, 259, 260, 5, 42, 0,
		0, 260, 261, 3, 10, 5, 0, 261, 262, 5, 43, 0, 0, 262, 352, 1, 0, 0, 0,
		263, 264, 5, 13, 0, 0, 264, 265, 5, 42, 0, 0, 265, 266, 3, 10, 5, 0, 266,
		267, 5, 43, 0, 0, 267, 352, 1, 0, 0, 0, 268, 269, 5, 14, 0, 0, 269, 270,
		5, 42, 0, 0, 270, 271, 3, 10, 5, 0, 271, 272, 5, 43, 0, 0, 272, 352, 1,
		0, 0, 0, 273, 274, 5, 15, 0, 0, 274, 275, 5, 42, 0, 0, 275, 276, 3, 10,
		5, 0, 276, 277, 5, 43, 0, 0, 277, 352, 1, 0, 0, 0, 278, 279, 5, 16, 0,
		0, 279, 280, 5, 42, 0, 0, 280, 281, 3, 10, 5, 0, 281, 282, 5, 43, 0, 0,
		282, 352, 1, 0, 0, 0, 283, 284, 5, 17, 0, 0, 284, 285, 5, 42, 0, 0, 285,
		286, 3, 10, 5, 0, 286, 287, 5, 43, 0, 0, 287, 352, 1, 0, 0, 0, 288, 289,
		5, 18, 0, 0, 289, 290, 5, 42, 0, 0, 290, 291, 3, 10, 5, 0, 291, 292, 5,
		43, 0, 0, 292, 352, 1, 0, 0, 0, 293, 294, 5, 19, 0, 0, 294, 295, 5, 42,
		0, 0, 295, 296, 3, 10, 5, 0, 296, 297, 5, 43, 0, 0, 297, 352, 1, 0, 0,
		0, 298, 299, 5, 17, 0, 0, 299, 300, 5, 42, 0, 0, 300, 301, 3, 10, 5, 0,
		301, 302, 5, 4, 0, 0, 302, 303, 3, 10, 5, 0, 303, 304, 5, 43, 0, 0, 304,
		352, 1, 0, 0, 0, 305, 306, 5, 8, 0, 0, 306, 307, 5, 42, 0, 0, 307, 308,
		3, 10, 5, 0, 308, 309, 5, 4, 0, 0, 309, 310, 3, 10, 5, 0, 310, 311, 5,
		43, 0, 0, 311, 352, 1, 0, 0, 0, 312, 313, 5, 20, 0, 0, 313, 314, 5, 42,
		0, 0, 314, 315, 3, 18, 9, 0, 315, 316, 5, 43, 0, 0, 316, 352, 1, 0, 0,
		0, 317, 318, 5, 21, 0, 0, 318, 319, 5, 42, 0, 0, 319, 320, 3, 18, 9, 0,
		320, 321, 5, 43, 0, 0, 321, 352, 1, 0, 0, 0, 322, 323, 5, 22, 0, 0, 323,
		324, 5, 42, 0, 0, 324, 325, 3, 18, 9, 0, 325, 326, 5, 43, 0, 0, 326, 352,
		1, 0, 0, 0, 327, 328, 5, 23, 0, 0, 328, 329, 5, 42, 0, 0, 329, 330, 3,
		18, 9, 0, 330, 331, 5, 4, 0, 0, 331, 332, 3, 10, 5, 0, 332, 333, 5, 43,
		0, 0, 333, 352, 1, 0, 0, 0, 334, 335, 5, 24, 0, 0, 335, 336, 5, 42, 0,
		0, 336, 337, 3, 18, 9, 0, 337, 338, 5, 43, 0, 0, 338, 352, 1, 0, 0, 0,
		339, 340, 5, 25, 0, 0, 340, 341, 5, 42, 0, 0, 341, 342, 3, 18, 9, 0, 342,
		343, 5, 43, 0, 0, 343, 352, 1, 0, 0, 0, 344, 345, 5, 26, 0, 0, 345, 346,
		5, 42, 0, 0, 346, 347, 3, 18, 9, 0, 347, 348, 5, 4, 0, 0, 348, 349, 3,
		10, 5, 0, 349, 350, 5, 43, 0, 0, 350, 352, 1, 0, 0, 0, 351, 223, 1, 0,
		0, 0, 351, 228, 1, 0, 0, 0, 351, 233, 1, 0, 0, 0, 351, 238, 1, 0, 0, 0,
		351, 243, 1, 0, 0, 0, 351, 248, 1, 0, 0, 0, 351, 253, 1, 0, 0, 0, 351,
		258, 1, 0, 0, 0, 351, 263, 1, 0, 0, 0, 351, 268, 1, 0, 0, 0, 351, 273,
		1, 0, 0, 0, 351, 278, 1, 0, 0, 0, 351, 283, 1, 0, 0, 0, 351, 288, 1, 0,
		0, 0, 351, 293, 1, 0, 0, 0, 351, 298, 1, 0, 0, 0, 351, 305, 1, 0, 0, 0,
		351, 312, 1, 0, 0, 0, 351, 317, 1, 0, 0, 0, 351, 322, 1, 0, 0, 0, 351,
		327, 1, 0, 0, 0, 351, 334, 1, 0, 0, 0, 351, 339, 1, 0, 0, 0, 351, 344,
		1, 0, 0, 0, 352, 29, 1, 0, 0, 0, 27, 41, 50, 71, 80, 91, 98, 111, 124,
		128, 131, 148, 150, 157, 165, 168, 174, 178, 183, 188, 194, 197, 202, 208,
		211, 217, 221, 351,
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

// AbacusParserInit initializes any static state used to implement AbacusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAbacusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AbacusParserInit() {
	staticData := &abacusParserStaticData
	staticData.once.Do(abacusParserInit)
}

// NewAbacusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAbacusParser(input antlr.TokenStream) *AbacusParser {
	AbacusParserInit()
	this := new(AbacusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &abacusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
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
	AbacusParserT__15             = 16
	AbacusParserT__16             = 17
	AbacusParserT__17             = 18
	AbacusParserT__18             = 19
	AbacusParserT__19             = 20
	AbacusParserT__20             = 21
	AbacusParserT__21             = 22
	AbacusParserT__22             = 23
	AbacusParserT__23             = 24
	AbacusParserT__24             = 25
	AbacusParserT__25             = 26
	AbacusParserAND               = 27
	AbacusParserOR                = 28
	AbacusParserXOR               = 29
	AbacusParserNOT               = 30
	AbacusParserEQ                = 31
	AbacusParserLS                = 32
	AbacusParserGR                = 33
	AbacusParserARROW             = 34
	AbacusParserPOW               = 35
	AbacusParserMUL               = 36
	AbacusParserDIV               = 37
	AbacusParserADD               = 38
	AbacusParserSUB               = 39
	AbacusParserPER               = 40
	AbacusParserPOINT             = 41
	AbacusParserLPAREN            = 42
	AbacusParserRPAREN            = 43
	AbacusParserLSQPAREN          = 44
	AbacusParserRSQPAREN          = 45
	AbacusParserCONSTANT          = 46
	AbacusParserSCIENTIFIC_NUMBER = 47
	AbacusParserVARIABLE          = 48
	AbacusParserLAMBDA_VARIABLE   = 49
	AbacusParserDIGITS            = 50
	AbacusParserUPPERCASE         = 51
	AbacusParserLOWERCASE         = 52
	AbacusParserWHITESPACE        = 53
)

// AbacusParser rules.
const (
	AbacusParserRULE_root                = 0
	AbacusParserRULE_declaration         = 1
	AbacusParserRULE_boolExpression      = 2
	AbacusParserRULE_boolAtom            = 3
	AbacusParserRULE_lambda              = 4
	AbacusParserRULE_expression          = 5
	AbacusParserRULE_parameter           = 6
	AbacusParserRULE_recursionParameters = 7
	AbacusParserRULE_mixedTuple          = 8
	AbacusParserRULE_tuple               = 9
	AbacusParserRULE_lambdaArguments     = 10
	AbacusParserRULE_variablesTuple      = 11
	AbacusParserRULE_atom                = 12
	AbacusParserRULE_sign                = 13
	AbacusParserRULE_function            = 14
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(AbacusParserEOF, 0)
}

func (s *RootContext) BoolExpression() IBoolExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExpressionContext)
}

func (s *RootContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *RootContext) LAMBDA_VARIABLE() antlr.TerminalNode {
	return s.GetToken(AbacusParserLAMBDA_VARIABLE, 0)
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
	this := p
	_ = this

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

	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Declaration()
		}
		{
			p.SetState(31)
			p.Match(AbacusParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.boolExpression(0)
		}
		{
			p.SetState(34)
			p.Match(AbacusParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Tuple()
		}
		{
			p.SetState(37)
			p.Match(AbacusParserEOF)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.Match(AbacusParserLAMBDA_VARIABLE)
		}
		{
			p.SetState(40)
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

func (s *DeclarationContext) CopyFrom(ctx *DeclarationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LambdaDeclarationContext struct {
	*DeclarationContext
}

func NewLambdaDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LambdaDeclarationContext {
	var p = new(LambdaDeclarationContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *LambdaDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaDeclarationContext) LAMBDA_VARIABLE() antlr.TerminalNode {
	return s.GetToken(AbacusParserLAMBDA_VARIABLE, 0)
}

func (s *LambdaDeclarationContext) EQ() antlr.TerminalNode {
	return s.GetToken(AbacusParserEQ, 0)
}

func (s *LambdaDeclarationContext) Lambda() ILambdaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaContext)
}

func (s *LambdaDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterLambdaDeclaration(s)
	}
}

func (s *LambdaDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitLambdaDeclaration(s)
	}
}

func (s *LambdaDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitLambdaDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableDeclarationContext struct {
	*DeclarationContext
}

func NewVariableDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.DeclarationContext = NewEmptyDeclarationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DeclarationContext))

	return p
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) VariablesTuple() IVariablesTupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariablesTupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariablesTupleContext)
}

func (s *VariableDeclarationContext) EQ() antlr.TerminalNode {
	return s.GetToken(AbacusParserEQ, 0)
}

func (s *VariableDeclarationContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

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

	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserLPAREN, AbacusParserVARIABLE:
		localctx = NewVariableDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.VariablesTuple()
		}
		{
			p.SetState(44)
			p.Match(AbacusParserEQ)
		}
		{
			p.SetState(45)
			p.Tuple()
		}

	case AbacusParserLAMBDA_VARIABLE:
		localctx = NewLambdaDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Match(AbacusParserLAMBDA_VARIABLE)
		}
		{
			p.SetState(48)
			p.Match(AbacusParserEQ)
		}
		{
			p.SetState(49)
			p.Lambda()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBoolExpressionContext is an interface to support dynamic dispatch.
type IBoolExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolExpressionContext differentiates from other interfaces.
	IsBoolExpressionContext()
}

type BoolExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExpressionContext() *BoolExpressionContext {
	var p = new(BoolExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_boolExpression
	return p
}

func (*BoolExpressionContext) IsBoolExpressionContext() {}

func NewBoolExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExpressionContext {
	var p = new(BoolExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_boolExpression

	return p
}

func (s *BoolExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExpressionContext) CopyFrom(ctx *BoolExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BoolExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotContext struct {
	*BoolExpressionContext
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	p.BoolExpressionContext = NewEmptyBoolExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressionContext))

	return p
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) NOT() antlr.TerminalNode {
	return s.GetToken(AbacusParserNOT, 0)
}

func (s *NotContext) BoolExpression() IBoolExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExpressionContext)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitNot(s)
	}
}

func (s *NotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type GreaterOrEqualComparisonContext struct {
	*BoolExpressionContext
}

func NewGreaterOrEqualComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterOrEqualComparisonContext {
	var p = new(GreaterOrEqualComparisonContext)

	p.BoolExpressionContext = NewEmptyBoolExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressionContext))

	return p
}

func (s *GreaterOrEqualComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterOrEqualComparisonContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *GreaterOrEqualComparisonContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	*BoolExpressionContext
}

func NewGreaterComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterComparisonContext {
	var p = new(GreaterComparisonContext)

	p.BoolExpressionContext = NewEmptyBoolExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressionContext))

	return p
}

func (s *GreaterComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterComparisonContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *GreaterComparisonContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	*BoolExpressionContext
}

func NewLessOrEqualComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessOrEqualComparisonContext {
	var p = new(LessOrEqualComparisonContext)

	p.BoolExpressionContext = NewEmptyBoolExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressionContext))

	return p
}

func (s *LessOrEqualComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessOrEqualComparisonContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LessOrEqualComparisonContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

type BooleanAtomContext struct {
	*BoolExpressionContext
}

func NewBooleanAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanAtomContext {
	var p = new(BooleanAtomContext)

	p.BoolExpressionContext = NewEmptyBoolExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressionContext))

	return p
}

func (s *BooleanAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanAtomContext) BoolAtom() IBoolAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolAtomContext)
}

func (s *BooleanAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitBooleanAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type LessComparisonContext struct {
	*BoolExpressionContext
}

func NewLessComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessComparisonContext {
	var p = new(LessComparisonContext)

	p.BoolExpressionContext = NewEmptyBoolExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressionContext))

	return p
}

func (s *LessComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessComparisonContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LessComparisonContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

type ParenthesesBooleanContext struct {
	*BoolExpressionContext
}

func NewParenthesesBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesesBooleanContext {
	var p = new(ParenthesesBooleanContext)

	p.BoolExpressionContext = NewEmptyBoolExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressionContext))

	return p
}

func (s *ParenthesesBooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesesBooleanContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *ParenthesesBooleanContext) BoolExpression() IBoolExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExpressionContext)
}

func (s *ParenthesesBooleanContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *ParenthesesBooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterParenthesesBoolean(s)
	}
}

func (s *ParenthesesBooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitParenthesesBoolean(s)
	}
}

func (s *ParenthesesBooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitParenthesesBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndOrXorContext struct {
	*BoolExpressionContext
	op antlr.Token
}

func NewAndOrXorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndOrXorContext {
	var p = new(AndOrXorContext)

	p.BoolExpressionContext = NewEmptyBoolExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressionContext))

	return p
}

func (s *AndOrXorContext) GetOp() antlr.Token { return s.op }

func (s *AndOrXorContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndOrXorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndOrXorContext) AllBoolExpression() []IBoolExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolExpressionContext); ok {
			len++
		}
	}

	tst := make([]IBoolExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolExpressionContext); ok {
			tst[i] = t.(IBoolExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndOrXorContext) BoolExpression(i int) IBoolExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExpressionContext)
}

func (s *AndOrXorContext) AND() antlr.TerminalNode {
	return s.GetToken(AbacusParserAND, 0)
}

func (s *AndOrXorContext) OR() antlr.TerminalNode {
	return s.GetToken(AbacusParserOR, 0)
}

func (s *AndOrXorContext) XOR() antlr.TerminalNode {
	return s.GetToken(AbacusParserXOR, 0)
}

func (s *AndOrXorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterAndOrXor(s)
	}
}

func (s *AndOrXorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitAndOrXor(s)
	}
}

func (s *AndOrXorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitAndOrXor(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualComparisonContext struct {
	*BoolExpressionContext
}

func NewEqualComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualComparisonContext {
	var p = new(EqualComparisonContext)

	p.BoolExpressionContext = NewEmptyBoolExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExpressionContext))

	return p
}

func (s *EqualComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualComparisonContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualComparisonContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

func (p *AbacusParser) BoolExpression() (localctx IBoolExpressionContext) {
	return p.boolExpression(0)
}

func (p *AbacusParser) boolExpression(_p int) (localctx IBoolExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBoolExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoolExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, AbacusParserRULE_boolExpression, _p)
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
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEqualComparisonContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(53)
			p.expression(0)
		}
		{
			p.SetState(54)
			p.Match(AbacusParserEQ)
		}
		{
			p.SetState(55)
			p.Match(AbacusParserEQ)
		}
		{
			p.SetState(56)
			p.expression(0)
		}

	case 2:
		localctx = NewLessComparisonContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.expression(0)
		}
		{
			p.SetState(59)
			p.Match(AbacusParserLS)
		}
		{
			p.SetState(60)
			p.expression(0)
		}

	case 3:
		localctx = NewGreaterComparisonContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.expression(0)
		}
		{
			p.SetState(63)
			p.Match(AbacusParserGR)
		}
		{
			p.SetState(64)
			p.expression(0)
		}

	case 4:
		localctx = NewLessOrEqualComparisonContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)
			p.expression(0)
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AbacusParserLS:
			{
				p.SetState(67)
				p.Match(AbacusParserLS)
			}
			{
				p.SetState(68)
				p.Match(AbacusParserEQ)
			}

		case AbacusParserEQ:
			{
				p.SetState(69)
				p.Match(AbacusParserEQ)
			}
			{
				p.SetState(70)
				p.Match(AbacusParserLS)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(73)
			p.expression(0)
		}

	case 5:
		localctx = NewGreaterOrEqualComparisonContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(75)
			p.expression(0)
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AbacusParserGR:
			{
				p.SetState(76)
				p.Match(AbacusParserGR)
			}
			{
				p.SetState(77)
				p.Match(AbacusParserEQ)
			}

		case AbacusParserEQ:
			{
				p.SetState(78)
				p.Match(AbacusParserEQ)
			}
			{
				p.SetState(79)
				p.Match(AbacusParserGR)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(82)
			p.expression(0)
		}

	case 6:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(84)
			p.Match(AbacusParserNOT)
		}
		{
			p.SetState(85)
			p.boolExpression(3)
		}

	case 7:
		localctx = NewParenthesesBooleanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(86)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(87)
			p.boolExpression(0)
		}
		{
			p.SetState(88)
			p.Match(AbacusParserRPAREN)
		}

	case 8:
		localctx = NewBooleanAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.BoolAtom()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndOrXorContext(p, NewBoolExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_boolExpression)
			p.SetState(93)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(94)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AndOrXorContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AbacusParserAND)|(1<<AbacusParserOR)|(1<<AbacusParserXOR))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AndOrXorContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(95)
				p.boolExpression(5)
			}

		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IBoolAtomContext is an interface to support dynamic dispatch.
type IBoolAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolAtomContext differentiates from other interfaces.
	IsBoolAtomContext()
}

type BoolAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolAtomContext() *BoolAtomContext {
	var p = new(BoolAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_boolAtom
	return p
}

func (*BoolAtomContext) IsBoolAtomContext() {}

func NewBoolAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolAtomContext {
	var p = new(BoolAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_boolAtom

	return p
}

func (s *BoolAtomContext) GetParser() antlr.Parser { return s.parser }
func (s *BoolAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterBoolAtom(s)
	}
}

func (s *BoolAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitBoolAtom(s)
	}
}

func (s *BoolAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitBoolAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) BoolAtom() (localctx IBoolAtomContext) {
	this := p
	_ = this

	localctx = NewBoolAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AbacusParserRULE_boolAtom)
	var _la int

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
		p.SetState(101)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AbacusParserT__0 || _la == AbacusParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_lambda
	return p
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) CopyFrom(ctx *LambdaContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullArityLambdaContext struct {
	*LambdaContext
}

func NewNullArityLambdaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullArityLambdaContext {
	var p = new(NullArityLambdaContext)

	p.LambdaContext = NewEmptyLambdaContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LambdaContext))

	return p
}

func (s *NullArityLambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullArityLambdaContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *NullArityLambdaContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *NullArityLambdaContext) ARROW() antlr.TerminalNode {
	return s.GetToken(AbacusParserARROW, 0)
}

func (s *NullArityLambdaContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *NullArityLambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterNullArityLambda(s)
	}
}

func (s *NullArityLambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitNullArityLambda(s)
	}
}

func (s *NullArityLambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitNullArityLambda(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariablesLambdaContext struct {
	*LambdaContext
}

func NewVariablesLambdaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariablesLambdaContext {
	var p = new(VariablesLambdaContext)

	p.LambdaContext = NewEmptyLambdaContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LambdaContext))

	return p
}

func (s *VariablesLambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablesLambdaContext) LambdaArguments() ILambdaArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaArgumentsContext)
}

func (s *VariablesLambdaContext) ARROW() antlr.TerminalNode {
	return s.GetToken(AbacusParserARROW, 0)
}

func (s *VariablesLambdaContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *VariablesLambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterVariablesLambda(s)
	}
}

func (s *VariablesLambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitVariablesLambda(s)
	}
}

func (s *VariablesLambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitVariablesLambda(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Lambda() (localctx ILambdaContext) {
	this := p
	_ = this

	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AbacusParserRULE_lambda)

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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVariablesLambdaContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.LambdaArguments()
		}
		{
			p.SetState(104)
			p.Match(AbacusParserARROW)
		}
		{
			p.SetState(105)
			p.Tuple()
		}

	case 2:
		localctx = NewNullArityLambdaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(108)
			p.Match(AbacusParserRPAREN)
		}
		{
			p.SetState(109)
			p.Match(AbacusParserARROW)
		}
		{
			p.SetState(110)
			p.Tuple()
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

type SignedExprContext struct {
	*ExpressionContext
}

func NewSignedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SignedExprContext {
	var p = new(SignedExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SignedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignedExprContext) Sign() ISignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignContext)
}

func (s *SignedExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SignedExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterSignedExpr(s)
	}
}

func (s *SignedExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitSignedExpr(s)
	}
}

func (s *SignedExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitSignedExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModContext struct {
	*ExpressionContext
}

func NewModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModContext {
	var p = new(ModContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ModContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ModContext) AllPER() []antlr.TerminalNode {
	return s.GetTokens(AbacusParserPER)
}

func (s *ModContext) PER(i int) antlr.TerminalNode {
	return s.GetToken(AbacusParserPER, i)
}

func (s *ModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterMod(s)
	}
}

func (s *ModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitMod(s)
	}
}

func (s *ModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitMod(s)

	default:
		return t.VisitChildren(s)
	}
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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

type PercentContext struct {
	*ExpressionContext
}

func NewPercentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PercentContext {
	var p = new(PercentContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PercentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PercentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PercentContext) PER() antlr.TerminalNode {
	return s.GetToken(AbacusParserPER, 0)
}

func (s *PercentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterPercent(s)
	}
}

func (s *PercentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitPercent(s)
	}
}

func (s *PercentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitPercent(s)

	default:
		return t.VisitChildren(s)
	}
}

type LambdaExprContext struct {
	*ExpressionContext
}

func NewLambdaExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LambdaExprContext {
	var p = new(LambdaExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LambdaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaExprContext) LAMBDA_VARIABLE() antlr.TerminalNode {
	return s.GetToken(AbacusParserLAMBDA_VARIABLE, 0)
}

func (s *LambdaExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *LambdaExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *LambdaExprContext) MixedTuple() IMixedTupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixedTupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixedTupleContext)
}

func (s *LambdaExprContext) RecursionParameters() IRecursionParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecursionParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecursionParametersContext)
}

func (s *LambdaExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterLambdaExpr(s)
	}
}

func (s *LambdaExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitLambdaExpr(s)
	}
}

func (s *LambdaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitLambdaExpr(s)

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PowContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, AbacusParserRULE_expression, _p)
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
	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserADD, AbacusParserSUB:
		localctx = NewSignedExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(114)
			p.Sign()
		}
		{
			p.SetState(115)
			p.expression(9)
		}

	case AbacusParserLPAREN:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(117)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(118)
			p.expression(0)
		}
		{
			p.SetState(119)
			p.Match(AbacusParserRPAREN)
		}

	case AbacusParserLAMBDA_VARIABLE:
		localctx = NewLambdaExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(121)
			p.Match(AbacusParserLAMBDA_VARIABLE)
		}
		{
			p.SetState(122)
			p.Match(AbacusParserLPAREN)
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AbacusParserT__4)|(1<<AbacusParserT__5)|(1<<AbacusParserT__6)|(1<<AbacusParserT__7)|(1<<AbacusParserT__8)|(1<<AbacusParserT__9)|(1<<AbacusParserT__10)|(1<<AbacusParserT__11)|(1<<AbacusParserT__12)|(1<<AbacusParserT__13)|(1<<AbacusParserT__14)|(1<<AbacusParserT__15)|(1<<AbacusParserT__16)|(1<<AbacusParserT__17)|(1<<AbacusParserT__18)|(1<<AbacusParserT__19)|(1<<AbacusParserT__20)|(1<<AbacusParserT__21)|(1<<AbacusParserT__22)|(1<<AbacusParserT__23)|(1<<AbacusParserT__24)|(1<<AbacusParserT__25))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(AbacusParserADD-38))|(1<<(AbacusParserSUB-38))|(1<<(AbacusParserLPAREN-38))|(1<<(AbacusParserCONSTANT-38))|(1<<(AbacusParserSCIENTIFIC_NUMBER-38))|(1<<(AbacusParserVARIABLE-38))|(1<<(AbacusParserLAMBDA_VARIABLE-38)))) != 0) {
			{
				p.SetState(123)
				p.MixedTuple()
			}

		}
		{
			p.SetState(126)
			p.Match(AbacusParserRPAREN)
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(127)
				p.RecursionParameters()
			}

		}

	case AbacusParserT__4, AbacusParserT__5, AbacusParserT__6, AbacusParserT__7, AbacusParserT__8, AbacusParserT__9, AbacusParserT__10, AbacusParserT__11, AbacusParserT__12, AbacusParserT__13, AbacusParserT__14, AbacusParserT__15, AbacusParserT__16, AbacusParserT__17, AbacusParserT__18, AbacusParserT__19, AbacusParserT__20, AbacusParserT__21, AbacusParserT__22, AbacusParserT__23, AbacusParserT__24, AbacusParserT__25, AbacusParserCONSTANT, AbacusParserSCIENTIFIC_NUMBER, AbacusParserVARIABLE:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(130)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(133)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(134)
					p.Match(AbacusParserPOW)
				}
				{
					p.SetState(135)
					p.expression(8)
				}

			case 2:
				localctx = NewModContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(136)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(137)
					p.Match(AbacusParserPER)
				}
				{
					p.SetState(138)
					p.Match(AbacusParserPER)
				}
				{
					p.SetState(139)
					p.expression(7)
				}

			case 3:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(141)

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
					p.SetState(142)
					p.expression(6)
				}

			case 4:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(144)

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
					p.SetState(145)
					p.expression(5)
				}

			case 5:
				localctx = NewPercentContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(147)
					p.Match(AbacusParserPER)
				}

			}

		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(AbacusParserVARIABLE, 0)
}

func (s *ParameterContext) EQ() antlr.TerminalNode {
	return s.GetToken(AbacusParserEQ, 0)
}

func (s *ParameterContext) BoolExpression() IBoolExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolExpressionContext)
}

func (s *ParameterContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Parameter() (localctx IParameterContext) {
	this := p
	_ = this

	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AbacusParserRULE_parameter)
	var _la int

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
		p.SetState(153)
		p.Match(AbacusParserVARIABLE)
	}
	{
		p.SetState(154)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AbacusParserT__2 || _la == AbacusParserEQ) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(155)
			p.boolExpression(0)
		}

	case 2:
		{
			p.SetState(156)
			p.expression(0)
		}

	}

	return localctx
}

// IRecursionParametersContext is an interface to support dynamic dispatch.
type IRecursionParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecursionParametersContext differentiates from other interfaces.
	IsRecursionParametersContext()
}

type RecursionParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecursionParametersContext() *RecursionParametersContext {
	var p = new(RecursionParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_recursionParameters
	return p
}

func (*RecursionParametersContext) IsRecursionParametersContext() {}

func NewRecursionParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecursionParametersContext {
	var p = new(RecursionParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_recursionParameters

	return p
}

func (s *RecursionParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *RecursionParametersContext) LSQPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLSQPAREN, 0)
}

func (s *RecursionParametersContext) RSQPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRSQPAREN, 0)
}

func (s *RecursionParametersContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *RecursionParametersContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *RecursionParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecursionParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecursionParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterRecursionParameters(s)
	}
}

func (s *RecursionParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitRecursionParameters(s)
	}
}

func (s *RecursionParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitRecursionParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) RecursionParameters() (localctx IRecursionParametersContext) {
	this := p
	_ = this

	localctx = NewRecursionParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AbacusParserRULE_recursionParameters)
	var _la int

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
		p.SetState(159)
		p.Match(AbacusParserLSQPAREN)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbacusParserVARIABLE {
		{
			p.SetState(160)
			p.Parameter()
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AbacusParserT__3 {
			{
				p.SetState(161)
				p.Match(AbacusParserT__3)
			}
			{
				p.SetState(162)
				p.Parameter()
			}

			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(170)
		p.Match(AbacusParserRSQPAREN)
	}

	return localctx
}

// IMixedTupleContext is an interface to support dynamic dispatch.
type IMixedTupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixedTupleContext differentiates from other interfaces.
	IsMixedTupleContext()
}

type MixedTupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixedTupleContext() *MixedTupleContext {
	var p = new(MixedTupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_mixedTuple
	return p
}

func (*MixedTupleContext) IsMixedTupleContext() {}

func NewMixedTupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MixedTupleContext {
	var p = new(MixedTupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_mixedTuple

	return p
}

func (s *MixedTupleContext) GetParser() antlr.Parser { return s.parser }

func (s *MixedTupleContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MixedTupleContext) LAMBDA_VARIABLE() antlr.TerminalNode {
	return s.GetToken(AbacusParserLAMBDA_VARIABLE, 0)
}

func (s *MixedTupleContext) MixedTuple() IMixedTupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixedTupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixedTupleContext)
}

func (s *MixedTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixedTupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MixedTupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterMixedTuple(s)
	}
}

func (s *MixedTupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitMixedTuple(s)
	}
}

func (s *MixedTupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitMixedTuple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) MixedTuple() (localctx IMixedTupleContext) {
	this := p
	_ = this

	localctx = NewMixedTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AbacusParserRULE_mixedTuple)
	var _la int

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
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(172)
			p.expression(0)
		}

	case 2:
		{
			p.SetState(173)
			p.Match(AbacusParserLAMBDA_VARIABLE)
		}

	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbacusParserT__3 {
		{
			p.SetState(176)
			p.Match(AbacusParserT__3)
		}
		{
			p.SetState(177)
			p.MixedTuple()
		}

	}

	return localctx
}

// ITupleContext is an interface to support dynamic dispatch.
type ITupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTupleContext differentiates from other interfaces.
	IsTupleContext()
}

type TupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTupleContext() *TupleContext {
	var p = new(TupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_tuple
	return p
}

func (*TupleContext) IsTupleContext() {}

func NewTupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleContext {
	var p = new(TupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_tuple

	return p
}

func (s *TupleContext) GetParser() antlr.Parser { return s.parser }

func (s *TupleContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TupleContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *TupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterTuple(s)
	}
}

func (s *TupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitTuple(s)
	}
}

func (s *TupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitTuple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Tuple() (localctx ITupleContext) {
	this := p
	_ = this

	localctx = NewTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AbacusParserRULE_tuple)

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
		p.SetState(180)
		p.expression(0)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(181)
			p.Match(AbacusParserT__3)
		}
		{
			p.SetState(182)
			p.Tuple()
		}

	}

	return localctx
}

// ILambdaArgumentsContext is an interface to support dynamic dispatch.
type ILambdaArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambdaArgumentsContext differentiates from other interfaces.
	IsLambdaArgumentsContext()
}

type LambdaArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaArgumentsContext() *LambdaArgumentsContext {
	var p = new(LambdaArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_lambdaArguments
	return p
}

func (*LambdaArgumentsContext) IsLambdaArgumentsContext() {}

func NewLambdaArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaArgumentsContext {
	var p = new(LambdaArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_lambdaArguments

	return p
}

func (s *LambdaArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaArgumentsContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(AbacusParserVARIABLE, 0)
}

func (s *LambdaArgumentsContext) LAMBDA_VARIABLE() antlr.TerminalNode {
	return s.GetToken(AbacusParserLAMBDA_VARIABLE, 0)
}

func (s *LambdaArgumentsContext) LambdaArguments() ILambdaArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILambdaArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILambdaArgumentsContext)
}

func (s *LambdaArgumentsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *LambdaArgumentsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *LambdaArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterLambdaArguments(s)
	}
}

func (s *LambdaArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitLambdaArguments(s)
	}
}

func (s *LambdaArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitLambdaArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) LambdaArguments() (localctx ILambdaArgumentsContext) {
	this := p
	_ = this

	localctx = NewLambdaArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AbacusParserRULE_lambdaArguments)
	var _la int

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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserVARIABLE, AbacusParserLAMBDA_VARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			_la = p.GetTokenStream().LA(1)

			if !(_la == AbacusParserVARIABLE || _la == AbacusParserLAMBDA_VARIABLE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbacusParserT__3 {
			{
				p.SetState(186)
				p.Match(AbacusParserT__3)
			}
			{
				p.SetState(187)
				p.LambdaArguments()
			}

		}

	case AbacusParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(191)
			_la = p.GetTokenStream().LA(1)

			if !(_la == AbacusParserVARIABLE || _la == AbacusParserLAMBDA_VARIABLE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbacusParserT__3 {
			{
				p.SetState(192)
				p.Match(AbacusParserT__3)
			}
			{
				p.SetState(193)
				p.LambdaArguments()
			}

		}
		{
			p.SetState(196)
			p.Match(AbacusParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariablesTupleContext is an interface to support dynamic dispatch.
type IVariablesTupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariablesTupleContext differentiates from other interfaces.
	IsVariablesTupleContext()
}

type VariablesTupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariablesTupleContext() *VariablesTupleContext {
	var p = new(VariablesTupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_variablesTuple
	return p
}

func (*VariablesTupleContext) IsVariablesTupleContext() {}

func NewVariablesTupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablesTupleContext {
	var p = new(VariablesTupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_variablesTuple

	return p
}

func (s *VariablesTupleContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablesTupleContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(AbacusParserVARIABLE, 0)
}

func (s *VariablesTupleContext) VariablesTuple() IVariablesTupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariablesTupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariablesTupleContext)
}

func (s *VariablesTupleContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *VariablesTupleContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *VariablesTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablesTupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariablesTupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterVariablesTuple(s)
	}
}

func (s *VariablesTupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitVariablesTuple(s)
	}
}

func (s *VariablesTupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitVariablesTuple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) VariablesTuple() (localctx IVariablesTupleContext) {
	this := p
	_ = this

	localctx = NewVariablesTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AbacusParserRULE_variablesTuple)
	var _la int

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

	p.SetState(211)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserVARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(AbacusParserVARIABLE)
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbacusParserT__3 {
			{
				p.SetState(200)
				p.Match(AbacusParserT__3)
			}
			{
				p.SetState(201)
				p.VariablesTuple()
			}

		}

	case AbacusParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(205)
			p.Match(AbacusParserVARIABLE)
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbacusParserT__3 {
			{
				p.SetState(206)
				p.Match(AbacusParserT__3)
			}
			{
				p.SetState(207)
				p.VariablesTuple()
			}

		}
		{
			p.SetState(210)
			p.Match(AbacusParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AbacusParserRULE_atom)

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

	p.SetState(217)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserT__4, AbacusParserT__5, AbacusParserT__6, AbacusParserT__7, AbacusParserT__8, AbacusParserT__9, AbacusParserT__10, AbacusParserT__11, AbacusParserT__12, AbacusParserT__13, AbacusParserT__14, AbacusParserT__15, AbacusParserT__16, AbacusParserT__17, AbacusParserT__18, AbacusParserT__19, AbacusParserT__20, AbacusParserT__21, AbacusParserT__22, AbacusParserT__23, AbacusParserT__24, AbacusParserT__25:
		localctx = NewFuncExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)
			p.Function()
		}

	case AbacusParserCONSTANT:
		localctx = NewConstantContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Match(AbacusParserCONSTANT)
		}

	case AbacusParserSCIENTIFIC_NUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(215)
			p.Match(AbacusParserSCIENTIFIC_NUMBER)
		}

	case AbacusParserVARIABLE:
		localctx = NewVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(216)
			p.Match(AbacusParserVARIABLE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISignContext is an interface to support dynamic dispatch.
type ISignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSignContext differentiates from other interfaces.
	IsSignContext()
}

type SignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignContext() *SignContext {
	var p = new(SignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbacusParserRULE_sign
	return p
}

func (*SignContext) IsSignContext() {}

func NewSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignContext {
	var p = new(SignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbacusParserRULE_sign

	return p
}

func (s *SignContext) GetParser() antlr.Parser { return s.parser }

func (s *SignContext) CopyFrom(ctx *SignContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PlusSignContext struct {
	*SignContext
}

func NewPlusSignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusSignContext {
	var p = new(PlusSignContext)

	p.SignContext = NewEmptySignContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SignContext))

	return p
}

func (s *PlusSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusSignContext) ADD() antlr.TerminalNode {
	return s.GetToken(AbacusParserADD, 0)
}

func (s *PlusSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterPlusSign(s)
	}
}

func (s *PlusSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitPlusSign(s)
	}
}

func (s *PlusSignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitPlusSign(s)

	default:
		return t.VisitChildren(s)
	}
}

type MinusSignContext struct {
	*SignContext
}

func NewMinusSignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinusSignContext {
	var p = new(MinusSignContext)

	p.SignContext = NewEmptySignContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SignContext))

	return p
}

func (s *MinusSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusSignContext) SUB() antlr.TerminalNode {
	return s.GetToken(AbacusParserSUB, 0)
}

func (s *MinusSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterMinusSign(s)
	}
}

func (s *MinusSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitMinusSign(s)
	}
}

func (s *MinusSignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitMinusSign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AbacusParser) Sign() (localctx ISignContext) {
	this := p
	_ = this

	localctx = NewSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AbacusParserRULE_sign)

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

	p.SetState(221)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserADD:
		localctx = NewPlusSignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.Match(AbacusParserADD)
		}

	case AbacusParserSUB:
		localctx = NewMinusSignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.Match(AbacusParserSUB)
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

func (s *MaxFunctionContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

type UntilFunctionContext struct {
	*FunctionContext
}

func NewUntilFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntilFunctionContext {
	var p = new(UntilFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *UntilFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntilFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *UntilFunctionContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *UntilFunctionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UntilFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *UntilFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterUntilFunction(s)
	}
}

func (s *UntilFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitUntilFunction(s)
	}
}

func (s *UntilFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitUntilFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type FromFunctionContext struct {
	*FunctionContext
}

func NewFromFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FromFunctionContext {
	var p = new(FromFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *FromFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *FromFunctionContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *FromFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *FromFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterFromFunction(s)
	}
}

func (s *FromFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitFromFunction(s)
	}
}

func (s *FromFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitFromFunction(s)

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

func (s *MinFunctionContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

type CbrtFunctionContext struct {
	*FunctionContext
}

func NewCbrtFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CbrtFunctionContext {
	var p = new(CbrtFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *CbrtFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CbrtFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *CbrtFunctionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CbrtFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *CbrtFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterCbrtFunction(s)
	}
}

func (s *CbrtFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitCbrtFunction(s)
	}
}

func (s *CbrtFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitCbrtFunction(s)

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogFunctionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

type AvgFunctionContext struct {
	*FunctionContext
}

func NewAvgFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AvgFunctionContext {
	var p = new(AvgFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *AvgFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AvgFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *AvgFunctionContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *AvgFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *AvgFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterAvgFunction(s)
	}
}

func (s *AvgFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitAvgFunction(s)
	}
}

func (s *AvgFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitAvgFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type SignFunctionContext struct {
	*FunctionContext
}

func NewSignFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SignFunctionContext {
	var p = new(SignFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *SignFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *SignFunctionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SignFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *SignFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterSignFunction(s)
	}
}

func (s *SignFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitSignFunction(s)
	}
}

func (s *SignFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitSignFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type NthFunctionContext struct {
	*FunctionContext
}

func NewNthFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NthFunctionContext {
	var p = new(NthFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *NthFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NthFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *NthFunctionContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *NthFunctionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NthFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *NthFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterNthFunction(s)
	}
}

func (s *NthFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitNthFunction(s)
	}
}

func (s *NthFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitNthFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReverseFunctionContext struct {
	*FunctionContext
}

func NewReverseFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReverseFunctionContext {
	var p = new(ReverseFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *ReverseFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReverseFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *ReverseFunctionContext) Tuple() ITupleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *ReverseFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *ReverseFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterReverseFunction(s)
	}
}

func (s *ReverseFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitReverseFunction(s)
	}
}

func (s *ReverseFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitReverseFunction(s)

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Round2FunctionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

type AbsFunctionContext struct {
	*FunctionContext
}

func NewAbsFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AbsFunctionContext {
	var p = new(AbsFunctionContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *AbsFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserLPAREN, 0)
}

func (s *AbsFunctionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AbsFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRPAREN, 0)
}

func (s *AbsFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.EnterAbsFunction(s)
	}
}

func (s *AbsFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbacusListener); ok {
		listenerT.ExitAbsFunction(s)
	}
}

func (s *AbsFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AbacusVisitor:
		return t.VisitAbsFunction(s)

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AbacusParserRULE_function)

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

	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSqrtFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Match(AbacusParserT__4)
		}
		{
			p.SetState(224)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(225)
			p.expression(0)
		}
		{
			p.SetState(226)
			p.Match(AbacusParserRPAREN)
		}

	case 2:
		localctx = NewCbrtFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.Match(AbacusParserT__5)
		}
		{
			p.SetState(229)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(230)
			p.expression(0)
		}
		{
			p.SetState(231)
			p.Match(AbacusParserRPAREN)
		}

	case 3:
		localctx = NewLnFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(233)
			p.Match(AbacusParserT__6)
		}
		{
			p.SetState(234)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(235)
			p.expression(0)
		}
		{
			p.SetState(236)
			p.Match(AbacusParserRPAREN)
		}

	case 4:
		localctx = NewLogDefFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(238)
			p.Match(AbacusParserT__7)
		}
		{
			p.SetState(239)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(240)
			p.expression(0)
		}
		{
			p.SetState(241)
			p.Match(AbacusParserRPAREN)
		}

	case 5:
		localctx = NewLog2FunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(243)
			p.Match(AbacusParserT__8)
		}
		{
			p.SetState(244)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(245)
			p.expression(0)
		}
		{
			p.SetState(246)
			p.Match(AbacusParserRPAREN)
		}

	case 6:
		localctx = NewLog10FunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(248)
			p.Match(AbacusParserT__9)
		}
		{
			p.SetState(249)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(250)
			p.expression(0)
		}
		{
			p.SetState(251)
			p.Match(AbacusParserRPAREN)
		}

	case 7:
		localctx = NewFloorFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(253)
			p.Match(AbacusParserT__10)
		}
		{
			p.SetState(254)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(255)
			p.expression(0)
		}
		{
			p.SetState(256)
			p.Match(AbacusParserRPAREN)
		}

	case 8:
		localctx = NewCeilFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(258)
			p.Match(AbacusParserT__11)
		}
		{
			p.SetState(259)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(260)
			p.expression(0)
		}
		{
			p.SetState(261)
			p.Match(AbacusParserRPAREN)
		}

	case 9:
		localctx = NewExpFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(263)
			p.Match(AbacusParserT__12)
		}
		{
			p.SetState(264)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(265)
			p.expression(0)
		}
		{
			p.SetState(266)
			p.Match(AbacusParserRPAREN)
		}

	case 10:
		localctx = NewSinFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(268)
			p.Match(AbacusParserT__13)
		}
		{
			p.SetState(269)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(270)
			p.expression(0)
		}
		{
			p.SetState(271)
			p.Match(AbacusParserRPAREN)
		}

	case 11:
		localctx = NewCosFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(273)
			p.Match(AbacusParserT__14)
		}
		{
			p.SetState(274)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(275)
			p.expression(0)
		}
		{
			p.SetState(276)
			p.Match(AbacusParserRPAREN)
		}

	case 12:
		localctx = NewTanFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(278)
			p.Match(AbacusParserT__15)
		}
		{
			p.SetState(279)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(280)
			p.expression(0)
		}
		{
			p.SetState(281)
			p.Match(AbacusParserRPAREN)
		}

	case 13:
		localctx = NewRoundDefFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(283)
			p.Match(AbacusParserT__16)
		}
		{
			p.SetState(284)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(285)
			p.expression(0)
		}
		{
			p.SetState(286)
			p.Match(AbacusParserRPAREN)
		}

	case 14:
		localctx = NewSignFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(288)
			p.Match(AbacusParserT__17)
		}
		{
			p.SetState(289)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(290)
			p.expression(0)
		}
		{
			p.SetState(291)
			p.Match(AbacusParserRPAREN)
		}

	case 15:
		localctx = NewAbsFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(293)
			p.Match(AbacusParserT__18)
		}
		{
			p.SetState(294)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(295)
			p.expression(0)
		}
		{
			p.SetState(296)
			p.Match(AbacusParserRPAREN)
		}

	case 16:
		localctx = NewRound2FunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(298)
			p.Match(AbacusParserT__16)
		}
		{
			p.SetState(299)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(300)
			p.expression(0)
		}
		{
			p.SetState(301)
			p.Match(AbacusParserT__3)
		}
		{
			p.SetState(302)
			p.expression(0)
		}
		{
			p.SetState(303)
			p.Match(AbacusParserRPAREN)
		}

	case 17:
		localctx = NewLogFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(305)
			p.Match(AbacusParserT__7)
		}
		{
			p.SetState(306)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(307)
			p.expression(0)
		}
		{
			p.SetState(308)
			p.Match(AbacusParserT__3)
		}
		{
			p.SetState(309)
			p.expression(0)
		}
		{
			p.SetState(310)
			p.Match(AbacusParserRPAREN)
		}

	case 18:
		localctx = NewMinFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(312)
			p.Match(AbacusParserT__19)
		}
		{
			p.SetState(313)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(314)
			p.Tuple()
		}
		{
			p.SetState(315)
			p.Match(AbacusParserRPAREN)
		}

	case 19:
		localctx = NewMaxFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(317)
			p.Match(AbacusParserT__20)
		}
		{
			p.SetState(318)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(319)
			p.Tuple()
		}
		{
			p.SetState(320)
			p.Match(AbacusParserRPAREN)
		}

	case 20:
		localctx = NewAvgFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(322)
			p.Match(AbacusParserT__21)
		}
		{
			p.SetState(323)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(324)
			p.Tuple()
		}
		{
			p.SetState(325)
			p.Match(AbacusParserRPAREN)
		}

	case 21:
		localctx = NewUntilFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(327)
			p.Match(AbacusParserT__22)
		}
		{
			p.SetState(328)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(329)
			p.Tuple()
		}
		{
			p.SetState(330)
			p.Match(AbacusParserT__3)
		}
		{
			p.SetState(331)
			p.expression(0)
		}
		{
			p.SetState(332)
			p.Match(AbacusParserRPAREN)
		}

	case 22:
		localctx = NewFromFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(334)
			p.Match(AbacusParserT__23)
		}
		{
			p.SetState(335)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(336)
			p.Tuple()
		}
		{
			p.SetState(337)
			p.Match(AbacusParserRPAREN)
		}

	case 23:
		localctx = NewReverseFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(339)
			p.Match(AbacusParserT__24)
		}
		{
			p.SetState(340)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(341)
			p.Tuple()
		}
		{
			p.SetState(342)
			p.Match(AbacusParserRPAREN)
		}

	case 24:
		localctx = NewNthFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(344)
			p.Match(AbacusParserT__25)
		}
		{
			p.SetState(345)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(346)
			p.Tuple()
		}
		{
			p.SetState(347)
			p.Match(AbacusParserT__3)
		}
		{
			p.SetState(348)
			p.expression(0)
		}
		{
			p.SetState(349)
			p.Match(AbacusParserRPAREN)
		}

	}

	return localctx
}

func (p *AbacusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *BoolExpressionContext = nil
		if localctx != nil {
			t = localctx.(*BoolExpressionContext)
		}
		return p.BoolExpression_Sempred(t, predIndex)

	case 5:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AbacusParser) BoolExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *AbacusParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
