package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/viktordanov/abacus/parser"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("(1 + 2)^(3+1) * 3")

	// Create the Lexer
	lexer := parser.NewAbacusLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewAbacusParser(stream)

	// Finally parse the expression
	visitor := NewAbacusVisitor()
	tree := p.Start()
	fmt.Println(visitor.Visit(tree))
	fmt.Println(visitor.trace)
}
