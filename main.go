package main

import (
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/viktordanov/abacus/parser"
	"math/big"
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
		ans := visitor.Visit(tree)
		switch val := ans.(type) {
		case *big.Float:
			fmt.Println(val.Text('g', 256))
		}
	}
}
