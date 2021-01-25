// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 40, 270,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 34, 10, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 43, 10, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 63, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 5, 4, 72, 10, 4, 3, 4, 3, 4, 5, 4, 76, 10, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 86, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 99, 10, 6, 3, 6, 3, 6, 5,
	6, 103, 10, 6, 3, 6, 5, 6, 106, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 123, 10,
	6, 12, 6, 14, 6, 126, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7,
	134, 10, 7, 5, 7, 136, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 143,
	10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 148, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9,
	154, 10, 9, 3, 9, 5, 9, 157, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	163, 10, 10, 3, 11, 3, 11, 5, 11, 167, 10, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 5, 12, 268, 10, 12, 3, 12, 2, 3, 10, 13, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 2, 4, 3, 2, 26, 27, 3, 2, 28, 29, 2, 306, 2, 33, 3, 2,
	2, 2, 4, 42, 3, 2, 2, 2, 6, 75, 3, 2, 2, 2, 8, 85, 3, 2, 2, 2, 10, 105,
	3, 2, 2, 2, 12, 127, 3, 2, 2, 2, 14, 139, 3, 2, 2, 2, 16, 156, 3, 2, 2,
	2, 18, 162, 3, 2, 2, 2, 20, 166, 3, 2, 2, 2, 22, 267, 3, 2, 2, 2, 24, 25,
	5, 4, 3, 2, 25, 26, 7, 2, 2, 3, 26, 34, 3, 2, 2, 2, 27, 28, 5, 6, 4, 2,
	28, 29, 7, 2, 2, 3, 29, 34, 3, 2, 2, 2, 30, 31, 5, 14, 8, 2, 31, 32, 7,
	2, 2, 3, 32, 34, 3, 2, 2, 2, 33, 24, 3, 2, 2, 2, 33, 27, 3, 2, 2, 2, 33,
	30, 3, 2, 2, 2, 34, 3, 3, 2, 2, 2, 35, 36, 5, 16, 9, 2, 36, 37, 7, 21,
	2, 2, 37, 38, 5, 14, 8, 2, 38, 43, 3, 2, 2, 2, 39, 40, 7, 39, 2, 2, 40,
	41, 7, 21, 2, 2, 41, 43, 5, 8, 5, 2, 42, 35, 3, 2, 2, 2, 42, 39, 3, 2,
	2, 2, 43, 5, 3, 2, 2, 2, 44, 45, 5, 10, 6, 2, 45, 46, 7, 21, 2, 2, 46,
	47, 7, 21, 2, 2, 47, 48, 5, 10, 6, 2, 48, 76, 3, 2, 2, 2, 49, 50, 5, 10,
	6, 2, 50, 51, 7, 22, 2, 2, 51, 52, 5, 10, 6, 2, 52, 76, 3, 2, 2, 2, 53,
	54, 5, 10, 6, 2, 54, 55, 7, 23, 2, 2, 55, 56, 5, 10, 6, 2, 56, 76, 3, 2,
	2, 2, 57, 62, 5, 10, 6, 2, 58, 59, 7, 22, 2, 2, 59, 63, 7, 21, 2, 2, 60,
	61, 7, 21, 2, 2, 61, 63, 7, 22, 2, 2, 62, 58, 3, 2, 2, 2, 62, 60, 3, 2,
	2, 2, 63, 64, 3, 2, 2, 2, 64, 65, 5, 10, 6, 2, 65, 76, 3, 2, 2, 2, 66,
	71, 5, 10, 6, 2, 67, 68, 7, 23, 2, 2, 68, 72, 7, 21, 2, 2, 69, 70, 7, 21,
	2, 2, 70, 72, 7, 23, 2, 2, 71, 67, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72,
	73, 3, 2, 2, 2, 73, 74, 5, 10, 6, 2, 74, 76, 3, 2, 2, 2, 75, 44, 3, 2,
	2, 2, 75, 49, 3, 2, 2, 2, 75, 53, 3, 2, 2, 2, 75, 57, 3, 2, 2, 2, 75, 66,
	3, 2, 2, 2, 76, 7, 3, 2, 2, 2, 77, 78, 5, 16, 9, 2, 78, 79, 7, 24, 2, 2,
	79, 80, 5, 14, 8, 2, 80, 86, 3, 2, 2, 2, 81, 82, 7, 32, 2, 2, 82, 83, 7,
	33, 2, 2, 83, 84, 7, 24, 2, 2, 84, 86, 5, 14, 8, 2, 85, 77, 3, 2, 2, 2,
	85, 81, 3, 2, 2, 2, 86, 9, 3, 2, 2, 2, 87, 88, 8, 6, 1, 2, 88, 89, 5, 20,
	11, 2, 89, 90, 5, 10, 6, 11, 90, 106, 3, 2, 2, 2, 91, 92, 7, 32, 2, 2,
	92, 93, 5, 10, 6, 2, 93, 94, 7, 33, 2, 2, 94, 106, 3, 2, 2, 2, 95, 96,
	7, 39, 2, 2, 96, 98, 7, 32, 2, 2, 97, 99, 5, 14, 8, 2, 98, 97, 3, 2, 2,
	2, 98, 99, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 102, 7, 33, 2, 2, 101,
	103, 5, 12, 7, 2, 102, 101, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 106,
	3, 2, 2, 2, 104, 106, 5, 18, 10, 2, 105, 87, 3, 2, 2, 2, 105, 91, 3, 2,
	2, 2, 105, 95, 3, 2, 2, 2, 105, 104, 3, 2, 2, 2, 106, 124, 3, 2, 2, 2,
	107, 108, 12, 9, 2, 2, 108, 109, 7, 25, 2, 2, 109, 123, 5, 10, 6, 10, 110,
	111, 12, 8, 2, 2, 111, 112, 7, 30, 2, 2, 112, 113, 7, 30, 2, 2, 113, 123,
	5, 10, 6, 9, 114, 115, 12, 7, 2, 2, 115, 116, 9, 2, 2, 2, 116, 123, 5,
	10, 6, 8, 117, 118, 12, 6, 2, 2, 118, 119, 9, 3, 2, 2, 119, 123, 5, 10,
	6, 7, 120, 121, 12, 10, 2, 2, 121, 123, 7, 30, 2, 2, 122, 107, 3, 2, 2,
	2, 122, 110, 3, 2, 2, 2, 122, 114, 3, 2, 2, 2, 122, 117, 3, 2, 2, 2, 122,
	120, 3, 2, 2, 2, 123, 126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125,
	3, 2, 2, 2, 125, 11, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 128, 7, 34,
	2, 2, 128, 135, 5, 10, 6, 2, 129, 130, 7, 3, 2, 2, 130, 133, 5, 10, 6,
	2, 131, 132, 7, 3, 2, 2, 132, 134, 5, 6, 4, 2, 133, 131, 3, 2, 2, 2, 133,
	134, 3, 2, 2, 2, 134, 136, 3, 2, 2, 2, 135, 129, 3, 2, 2, 2, 135, 136,
	3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 138, 7, 35, 2, 2, 138, 13, 3, 2,
	2, 2, 139, 142, 5, 10, 6, 2, 140, 141, 7, 3, 2, 2, 141, 143, 5, 14, 8,
	2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 15, 3, 2, 2, 2, 144,
	147, 7, 38, 2, 2, 145, 146, 7, 3, 2, 2, 146, 148, 5, 16, 9, 2, 147, 145,
	3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 157, 3, 2, 2, 2, 149, 150, 7, 32,
	2, 2, 150, 153, 7, 38, 2, 2, 151, 152, 7, 3, 2, 2, 152, 154, 5, 16, 9,
	2, 153, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155,
	157, 7, 33, 2, 2, 156, 144, 3, 2, 2, 2, 156, 149, 3, 2, 2, 2, 157, 17,
	3, 2, 2, 2, 158, 163, 5, 22, 12, 2, 159, 163, 7, 36, 2, 2, 160, 163, 7,
	37, 2, 2, 161, 163, 7, 38, 2, 2, 162, 158, 3, 2, 2, 2, 162, 159, 3, 2,
	2, 2, 162, 160, 3, 2, 2, 2, 162, 161, 3, 2, 2, 2, 163, 19, 3, 2, 2, 2,
	164, 167, 7, 28, 2, 2, 165, 167, 7, 29, 2, 2, 166, 164, 3, 2, 2, 2, 166,
	165, 3, 2, 2, 2, 167, 21, 3, 2, 2, 2, 168, 169, 7, 4, 2, 2, 169, 170, 7,
	32, 2, 2, 170, 171, 5, 10, 6, 2, 171, 172, 7, 33, 2, 2, 172, 268, 3, 2,
	2, 2, 173, 174, 7, 5, 2, 2, 174, 175, 7, 32, 2, 2, 175, 176, 5, 10, 6,
	2, 176, 177, 7, 33, 2, 2, 177, 268, 3, 2, 2, 2, 178, 179, 7, 6, 2, 2, 179,
	180, 7, 32, 2, 2, 180, 181, 5, 10, 6, 2, 181, 182, 7, 33, 2, 2, 182, 268,
	3, 2, 2, 2, 183, 184, 7, 7, 2, 2, 184, 185, 7, 32, 2, 2, 185, 186, 5, 10,
	6, 2, 186, 187, 7, 33, 2, 2, 187, 268, 3, 2, 2, 2, 188, 189, 7, 8, 2, 2,
	189, 190, 7, 32, 2, 2, 190, 191, 5, 10, 6, 2, 191, 192, 7, 33, 2, 2, 192,
	268, 3, 2, 2, 2, 193, 194, 7, 9, 2, 2, 194, 195, 7, 32, 2, 2, 195, 196,
	5, 10, 6, 2, 196, 197, 7, 33, 2, 2, 197, 268, 3, 2, 2, 2, 198, 199, 7,
	10, 2, 2, 199, 200, 7, 32, 2, 2, 200, 201, 5, 10, 6, 2, 201, 202, 7, 33,
	2, 2, 202, 268, 3, 2, 2, 2, 203, 204, 7, 11, 2, 2, 204, 205, 7, 32, 2,
	2, 205, 206, 5, 10, 6, 2, 206, 207, 7, 33, 2, 2, 207, 268, 3, 2, 2, 2,
	208, 209, 7, 12, 2, 2, 209, 210, 7, 32, 2, 2, 210, 211, 5, 10, 6, 2, 211,
	212, 7, 33, 2, 2, 212, 268, 3, 2, 2, 2, 213, 214, 7, 13, 2, 2, 214, 215,
	7, 32, 2, 2, 215, 216, 5, 10, 6, 2, 216, 217, 7, 33, 2, 2, 217, 268, 3,
	2, 2, 2, 218, 219, 7, 14, 2, 2, 219, 220, 7, 32, 2, 2, 220, 221, 5, 10,
	6, 2, 221, 222, 7, 33, 2, 2, 222, 268, 3, 2, 2, 2, 223, 224, 7, 15, 2,
	2, 224, 225, 7, 32, 2, 2, 225, 226, 5, 10, 6, 2, 226, 227, 7, 33, 2, 2,
	227, 268, 3, 2, 2, 2, 228, 229, 7, 16, 2, 2, 229, 230, 7, 32, 2, 2, 230,
	231, 5, 10, 6, 2, 231, 232, 7, 33, 2, 2, 232, 268, 3, 2, 2, 2, 233, 234,
	7, 17, 2, 2, 234, 235, 7, 32, 2, 2, 235, 236, 5, 10, 6, 2, 236, 237, 7,
	33, 2, 2, 237, 268, 3, 2, 2, 2, 238, 239, 7, 16, 2, 2, 239, 240, 7, 32,
	2, 2, 240, 241, 5, 10, 6, 2, 241, 242, 7, 3, 2, 2, 242, 243, 5, 10, 6,
	2, 243, 244, 7, 33, 2, 2, 244, 268, 3, 2, 2, 2, 245, 246, 7, 7, 2, 2, 246,
	247, 7, 32, 2, 2, 247, 248, 5, 10, 6, 2, 248, 249, 7, 3, 2, 2, 249, 250,
	5, 10, 6, 2, 250, 251, 7, 33, 2, 2, 251, 268, 3, 2, 2, 2, 252, 253, 7,
	18, 2, 2, 253, 254, 7, 32, 2, 2, 254, 255, 5, 14, 8, 2, 255, 256, 7, 33,
	2, 2, 256, 268, 3, 2, 2, 2, 257, 258, 7, 19, 2, 2, 258, 259, 7, 32, 2,
	2, 259, 260, 5, 14, 8, 2, 260, 261, 7, 33, 2, 2, 261, 268, 3, 2, 2, 2,
	262, 263, 7, 20, 2, 2, 263, 264, 7, 32, 2, 2, 264, 265, 5, 14, 8, 2, 265,
	266, 7, 33, 2, 2, 266, 268, 3, 2, 2, 2, 267, 168, 3, 2, 2, 2, 267, 173,
	3, 2, 2, 2, 267, 178, 3, 2, 2, 2, 267, 183, 3, 2, 2, 2, 267, 188, 3, 2,
	2, 2, 267, 193, 3, 2, 2, 2, 267, 198, 3, 2, 2, 2, 267, 203, 3, 2, 2, 2,
	267, 208, 3, 2, 2, 2, 267, 213, 3, 2, 2, 2, 267, 218, 3, 2, 2, 2, 267,
	223, 3, 2, 2, 2, 267, 228, 3, 2, 2, 2, 267, 233, 3, 2, 2, 2, 267, 238,
	3, 2, 2, 2, 267, 245, 3, 2, 2, 2, 267, 252, 3, 2, 2, 2, 267, 257, 3, 2,
	2, 2, 267, 262, 3, 2, 2, 2, 268, 23, 3, 2, 2, 2, 22, 33, 42, 62, 71, 75,
	85, 98, 102, 105, 122, 124, 133, 135, 142, 147, 153, 156, 162, 166, 267,
}
var literalNames = []string{
	"", "','", "'sqrt'", "'cbrt'", "'ln'", "'log'", "'log2'", "'log10'", "'floor'",
	"'ceil'", "'exp'", "'sin'", "'cos'", "'tan'", "'round'", "'abs'", "'min'",
	"'max'", "'avg'", "'='", "'<'", "'>'", "'->'", "", "'*'", "'/'", "'+'",
	"'-'", "'%'", "'.'", "'('", "')'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "EQ", "LS", "GR", "ARROW", "POW", "MUL", "DIV", "ADD", "SUB", "PER",
	"POINT", "LPAREN", "RPAREN", "LSQPAREN", "RSQPAREN", "CONSTANT", "SCIENTIFIC_NUMBER",
	"VARIABLE", "LAMBDA_VARIABLE", "WHITESPACE",
}

var ruleNames = []string{
	"root", "declaration", "comparison", "lambda", "expression", "recursionParameters",
	"tuple", "variablesTuple", "atom", "sign", "function",
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
	AbacusParserT__15             = 16
	AbacusParserT__16             = 17
	AbacusParserT__17             = 18
	AbacusParserEQ                = 19
	AbacusParserLS                = 20
	AbacusParserGR                = 21
	AbacusParserARROW             = 22
	AbacusParserPOW               = 23
	AbacusParserMUL               = 24
	AbacusParserDIV               = 25
	AbacusParserADD               = 26
	AbacusParserSUB               = 27
	AbacusParserPER               = 28
	AbacusParserPOINT             = 29
	AbacusParserLPAREN            = 30
	AbacusParserRPAREN            = 31
	AbacusParserLSQPAREN          = 32
	AbacusParserRSQPAREN          = 33
	AbacusParserCONSTANT          = 34
	AbacusParserSCIENTIFIC_NUMBER = 35
	AbacusParserVARIABLE          = 36
	AbacusParserLAMBDA_VARIABLE   = 37
	AbacusParserWHITESPACE        = 38
)

// AbacusParser rules.
const (
	AbacusParserRULE_root                = 0
	AbacusParserRULE_declaration         = 1
	AbacusParserRULE_comparison          = 2
	AbacusParserRULE_lambda              = 3
	AbacusParserRULE_expression          = 4
	AbacusParserRULE_recursionParameters = 5
	AbacusParserRULE_tuple               = 6
	AbacusParserRULE_variablesTuple      = 7
	AbacusParserRULE_atom                = 8
	AbacusParserRULE_sign                = 9
	AbacusParserRULE_function            = 10
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

func (s *RootContext) Tuple() ITupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
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

	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Declaration()
		}
		{
			p.SetState(23)
			p.Match(AbacusParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Comparison()
		}
		{
			p.SetState(26)
			p.Match(AbacusParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(28)
			p.Tuple()
		}
		{
			p.SetState(29)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILambdaContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariablesTupleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariablesTupleContext)
}

func (s *VariableDeclarationContext) EQ() antlr.TerminalNode {
	return s.GetToken(AbacusParserEQ, 0)
}

func (s *VariableDeclarationContext) Tuple() ITupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), 0)

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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserLPAREN, AbacusParserVARIABLE:
		localctx = NewVariableDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.VariablesTuple()
		}
		{
			p.SetState(34)
			p.Match(AbacusParserEQ)
		}
		{
			p.SetState(35)
			p.Tuple()
		}

	case AbacusParserLAMBDA_VARIABLE:
		localctx = NewLambdaDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(AbacusParserLAMBDA_VARIABLE)
		}
		{
			p.SetState(38)
			p.Match(AbacusParserEQ)
		}
		{
			p.SetState(39)
			p.Lambda()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEqualComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.expression(0)
		}
		{
			p.SetState(43)
			p.Match(AbacusParserEQ)
		}
		{
			p.SetState(44)
			p.Match(AbacusParserEQ)
		}
		{
			p.SetState(45)
			p.expression(0)
		}

	case 2:
		localctx = NewLessComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.expression(0)
		}
		{
			p.SetState(48)
			p.Match(AbacusParserLS)
		}
		{
			p.SetState(49)
			p.expression(0)
		}

	case 3:
		localctx = NewGreaterComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(51)
			p.expression(0)
		}
		{
			p.SetState(52)
			p.Match(AbacusParserGR)
		}
		{
			p.SetState(53)
			p.expression(0)
		}

	case 4:
		localctx = NewLessOrEqualComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(55)
			p.expression(0)
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AbacusParserLS:
			{
				p.SetState(56)
				p.Match(AbacusParserLS)
			}
			{
				p.SetState(57)
				p.Match(AbacusParserEQ)
			}

		case AbacusParserEQ:
			{
				p.SetState(58)
				p.Match(AbacusParserEQ)
			}
			{
				p.SetState(59)
				p.Match(AbacusParserLS)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(62)
			p.expression(0)
		}

	case 5:
		localctx = NewGreaterOrEqualComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(64)
			p.expression(0)
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AbacusParserGR:
			{
				p.SetState(65)
				p.Match(AbacusParserGR)
			}
			{
				p.SetState(66)
				p.Match(AbacusParserEQ)
			}

		case AbacusParserEQ:
			{
				p.SetState(67)
				p.Match(AbacusParserEQ)
			}
			{
				p.SetState(68)
				p.Match(AbacusParserGR)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(71)
			p.expression(0)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), 0)

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

func (s *VariablesLambdaContext) VariablesTuple() IVariablesTupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariablesTupleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariablesTupleContext)
}

func (s *VariablesLambdaContext) ARROW() antlr.TerminalNode {
	return s.GetToken(AbacusParserARROW, 0)
}

func (s *VariablesLambdaContext) Tuple() ITupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), 0)

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
	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AbacusParserRULE_lambda)

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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVariablesLambdaContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.VariablesTuple()
		}
		{
			p.SetState(76)
			p.Match(AbacusParserARROW)
		}
		{
			p.SetState(77)
			p.Tuple()
		}

	case 2:
		localctx = NewNullArityLambdaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(80)
			p.Match(AbacusParserRPAREN)
		}
		{
			p.SetState(81)
			p.Match(AbacusParserARROW)
		}
		{
			p.SetState(82)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignContext)
}

func (s *SignedExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ModContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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

func (s *LambdaExprContext) Tuple() ITupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *LambdaExprContext) RecursionParameters() IRecursionParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecursionParametersContext)(nil)).Elem(), 0)

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
	_startState := 8
	p.EnterRecursionRule(localctx, 8, AbacusParserRULE_expression, _p)
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
	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserADD, AbacusParserSUB:
		localctx = NewSignedExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(86)
			p.Sign()
		}
		{
			p.SetState(87)
			p.expression(9)
		}

	case AbacusParserLPAREN:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
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

	case AbacusParserLAMBDA_VARIABLE:
		localctx = NewLambdaExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(93)
			p.Match(AbacusParserLAMBDA_VARIABLE)
		}
		{
			p.SetState(94)
			p.Match(AbacusParserLPAREN)
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AbacusParserT__1)|(1<<AbacusParserT__2)|(1<<AbacusParserT__3)|(1<<AbacusParserT__4)|(1<<AbacusParserT__5)|(1<<AbacusParserT__6)|(1<<AbacusParserT__7)|(1<<AbacusParserT__8)|(1<<AbacusParserT__9)|(1<<AbacusParserT__10)|(1<<AbacusParserT__11)|(1<<AbacusParserT__12)|(1<<AbacusParserT__13)|(1<<AbacusParserT__14)|(1<<AbacusParserT__15)|(1<<AbacusParserT__16)|(1<<AbacusParserT__17)|(1<<AbacusParserADD)|(1<<AbacusParserSUB)|(1<<AbacusParserLPAREN))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(AbacusParserCONSTANT-34))|(1<<(AbacusParserSCIENTIFIC_NUMBER-34))|(1<<(AbacusParserVARIABLE-34))|(1<<(AbacusParserLAMBDA_VARIABLE-34)))) != 0) {
			{
				p.SetState(95)
				p.Tuple()
			}

		}
		{
			p.SetState(98)
			p.Match(AbacusParserRPAREN)
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(99)
				p.RecursionParameters()
			}

		}

	case AbacusParserT__1, AbacusParserT__2, AbacusParserT__3, AbacusParserT__4, AbacusParserT__5, AbacusParserT__6, AbacusParserT__7, AbacusParserT__8, AbacusParserT__9, AbacusParserT__10, AbacusParserT__11, AbacusParserT__12, AbacusParserT__13, AbacusParserT__14, AbacusParserT__15, AbacusParserT__16, AbacusParserT__17, AbacusParserCONSTANT, AbacusParserSCIENTIFIC_NUMBER, AbacusParserVARIABLE:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(102)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(106)
					p.Match(AbacusParserPOW)
				}
				{
					p.SetState(107)
					p.expression(8)
				}

			case 2:
				localctx = NewModContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(109)
					p.Match(AbacusParserPER)
				}
				{
					p.SetState(110)
					p.Match(AbacusParserPER)
				}
				{
					p.SetState(111)
					p.expression(7)
				}

			case 3:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(112)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(113)

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
					p.SetState(114)
					p.expression(6)
				}

			case 4:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(115)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(116)

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
					p.SetState(117)
					p.expression(5)
				}

			case 5:
				localctx = NewPercentContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AbacusParserRULE_expression)
				p.SetState(118)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(119)
					p.Match(AbacusParserPER)
				}

			}

		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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

func (s *RecursionParametersContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *RecursionParametersContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RecursionParametersContext) RSQPAREN() antlr.TerminalNode {
	return s.GetToken(AbacusParserRSQPAREN, 0)
}

func (s *RecursionParametersContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
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
	localctx = NewRecursionParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AbacusParserRULE_recursionParameters)
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
		p.SetState(125)
		p.Match(AbacusParserLSQPAREN)
	}
	{
		p.SetState(126)
		p.expression(0)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbacusParserT__0 {
		{
			p.SetState(127)
			p.Match(AbacusParserT__0)
		}
		{
			p.SetState(128)
			p.expression(0)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbacusParserT__0 {
			{
				p.SetState(129)
				p.Match(AbacusParserT__0)
			}
			{
				p.SetState(130)
				p.Comparison()
			}

		}

	}
	{
		p.SetState(135)
		p.Match(AbacusParserRSQPAREN)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TupleContext) Tuple() ITupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), 0)

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
	localctx = NewTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AbacusParserRULE_tuple)
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
		p.SetState(137)
		p.expression(0)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AbacusParserT__0 {
		{
			p.SetState(138)
			p.Match(AbacusParserT__0)
		}
		{
			p.SetState(139)
			p.Tuple()
		}

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariablesTupleContext)(nil)).Elem(), 0)

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
	localctx = NewVariablesTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AbacusParserRULE_variablesTuple)
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

	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserVARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Match(AbacusParserVARIABLE)
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbacusParserT__0 {
			{
				p.SetState(143)
				p.Match(AbacusParserT__0)
			}
			{
				p.SetState(144)
				p.VariablesTuple()
			}

		}

	case AbacusParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(148)
			p.Match(AbacusParserVARIABLE)
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AbacusParserT__0 {
			{
				p.SetState(149)
				p.Match(AbacusParserT__0)
			}
			{
				p.SetState(150)
				p.VariablesTuple()
			}

		}
		{
			p.SetState(153)
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
	p.EnterRule(localctx, 16, AbacusParserRULE_atom)

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

	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserT__1, AbacusParserT__2, AbacusParserT__3, AbacusParserT__4, AbacusParserT__5, AbacusParserT__6, AbacusParserT__7, AbacusParserT__8, AbacusParserT__9, AbacusParserT__10, AbacusParserT__11, AbacusParserT__12, AbacusParserT__13, AbacusParserT__14, AbacusParserT__15, AbacusParserT__16, AbacusParserT__17:
		localctx = NewFuncExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.Function()
		}

	case AbacusParserCONSTANT:
		localctx = NewConstantContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)
			p.Match(AbacusParserCONSTANT)
		}

	case AbacusParserSCIENTIFIC_NUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(158)
			p.Match(AbacusParserSCIENTIFIC_NUMBER)
		}

	case AbacusParserVARIABLE:
		localctx = NewVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(159)
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
	localctx = NewSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AbacusParserRULE_sign)

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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AbacusParserADD:
		localctx = NewPlusSignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(AbacusParserADD)
		}

	case AbacusParserSUB:
		localctx = NewMinusSignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
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

func (s *MaxFunctionContext) Tuple() ITupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), 0)

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

func (s *MinFunctionContext) Tuple() ITupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 20, AbacusParserRULE_function)

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

	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSqrtFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Match(AbacusParserT__1)
		}
		{
			p.SetState(167)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(168)
			p.expression(0)
		}
		{
			p.SetState(169)
			p.Match(AbacusParserRPAREN)
		}

	case 2:
		localctx = NewCbrtFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Match(AbacusParserT__2)
		}
		{
			p.SetState(172)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(173)
			p.expression(0)
		}
		{
			p.SetState(174)
			p.Match(AbacusParserRPAREN)
		}

	case 3:
		localctx = NewLnFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(176)
			p.Match(AbacusParserT__3)
		}
		{
			p.SetState(177)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(178)
			p.expression(0)
		}
		{
			p.SetState(179)
			p.Match(AbacusParserRPAREN)
		}

	case 4:
		localctx = NewLogDefFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(181)
			p.Match(AbacusParserT__4)
		}
		{
			p.SetState(182)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(183)
			p.expression(0)
		}
		{
			p.SetState(184)
			p.Match(AbacusParserRPAREN)
		}

	case 5:
		localctx = NewLog2FunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(186)
			p.Match(AbacusParserT__5)
		}
		{
			p.SetState(187)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(188)
			p.expression(0)
		}
		{
			p.SetState(189)
			p.Match(AbacusParserRPAREN)
		}

	case 6:
		localctx = NewLog10FunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(191)
			p.Match(AbacusParserT__6)
		}
		{
			p.SetState(192)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(193)
			p.expression(0)
		}
		{
			p.SetState(194)
			p.Match(AbacusParserRPAREN)
		}

	case 7:
		localctx = NewFloorFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(196)
			p.Match(AbacusParserT__7)
		}
		{
			p.SetState(197)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(198)
			p.expression(0)
		}
		{
			p.SetState(199)
			p.Match(AbacusParserRPAREN)
		}

	case 8:
		localctx = NewCeilFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(201)
			p.Match(AbacusParserT__8)
		}
		{
			p.SetState(202)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(203)
			p.expression(0)
		}
		{
			p.SetState(204)
			p.Match(AbacusParserRPAREN)
		}

	case 9:
		localctx = NewExpFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(206)
			p.Match(AbacusParserT__9)
		}
		{
			p.SetState(207)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(208)
			p.expression(0)
		}
		{
			p.SetState(209)
			p.Match(AbacusParserRPAREN)
		}

	case 10:
		localctx = NewSinFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(211)
			p.Match(AbacusParserT__10)
		}
		{
			p.SetState(212)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(213)
			p.expression(0)
		}
		{
			p.SetState(214)
			p.Match(AbacusParserRPAREN)
		}

	case 11:
		localctx = NewCosFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(216)
			p.Match(AbacusParserT__11)
		}
		{
			p.SetState(217)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(218)
			p.expression(0)
		}
		{
			p.SetState(219)
			p.Match(AbacusParserRPAREN)
		}

	case 12:
		localctx = NewTanFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(221)
			p.Match(AbacusParserT__12)
		}
		{
			p.SetState(222)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(223)
			p.expression(0)
		}
		{
			p.SetState(224)
			p.Match(AbacusParserRPAREN)
		}

	case 13:
		localctx = NewRoundDefFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(226)
			p.Match(AbacusParserT__13)
		}
		{
			p.SetState(227)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(228)
			p.expression(0)
		}
		{
			p.SetState(229)
			p.Match(AbacusParserRPAREN)
		}

	case 14:
		localctx = NewAbsFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(231)
			p.Match(AbacusParserT__14)
		}
		{
			p.SetState(232)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(233)
			p.expression(0)
		}
		{
			p.SetState(234)
			p.Match(AbacusParserRPAREN)
		}

	case 15:
		localctx = NewRound2FunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(236)
			p.Match(AbacusParserT__13)
		}
		{
			p.SetState(237)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(238)
			p.expression(0)
		}
		{
			p.SetState(239)
			p.Match(AbacusParserT__0)
		}
		{
			p.SetState(240)
			p.expression(0)
		}
		{
			p.SetState(241)
			p.Match(AbacusParserRPAREN)
		}

	case 16:
		localctx = NewLogFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(243)
			p.Match(AbacusParserT__4)
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
			p.Match(AbacusParserT__0)
		}
		{
			p.SetState(247)
			p.expression(0)
		}
		{
			p.SetState(248)
			p.Match(AbacusParserRPAREN)
		}

	case 17:
		localctx = NewMinFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(250)
			p.Match(AbacusParserT__15)
		}
		{
			p.SetState(251)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(252)
			p.Tuple()
		}
		{
			p.SetState(253)
			p.Match(AbacusParserRPAREN)
		}

	case 18:
		localctx = NewMaxFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(255)
			p.Match(AbacusParserT__16)
		}
		{
			p.SetState(256)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(257)
			p.Tuple()
		}
		{
			p.SetState(258)
			p.Match(AbacusParserRPAREN)
		}

	case 19:
		localctx = NewAvgFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(260)
			p.Match(AbacusParserT__17)
		}
		{
			p.SetState(261)
			p.Match(AbacusParserLPAREN)
		}
		{
			p.SetState(262)
			p.Tuple()
		}
		{
			p.SetState(263)
			p.Match(AbacusParserRPAREN)
		}

	}

	return localctx
}

func (p *AbacusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
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
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
