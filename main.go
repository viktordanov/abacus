package main

import (
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/viktordanov/abacus/parser"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	visitor := NewAbacusVisitor()
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')

		is := antlr.NewInputStream(input)
		lexer := parser.NewAbacusLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := parser.NewAbacusParser(stream)
		tree := p.Start()
		fmt.Println(visitor.Visit(tree))
	}
}
