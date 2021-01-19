package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/fatih/color"
	"github.com/peterh/liner"
	"github.com/viktordanov/abacus/parser"
	"log"
	"os"
	"path/filepath"
	"strings"

	"math/big"
)

var historyFile = filepath.Join(os.TempDir(), ".abacus_history")

func main() {
	visitor := NewAbacusVisitor()
	line := liner.NewLiner()
	answerPrinter := color.New(color.FgGreen)
	defer line.Close()

	line.SetCtrlCAborts(true)
	if f, err := os.Open(historyFile); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	if f, err := os.Create(historyFile); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}

	for {
		if name, err := line.Prompt("> "); err == nil {
			line.AppendHistory(name)
			is := antlr.NewInputStream(name)
			lexer := parser.NewAbacusLexer(is)
			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

			p := parser.NewAbacusParser(stream)
			tree := p.Root()
			ans := visitor.Visit(tree)
			switch val := ans.(type) {
			case *big.Float:
				answerPrinter.Println(val.Text('g', 256))
			}
		} else if err == liner.ErrPromptAborted {
			os.Exit(0)
		} else {
			log.Print("Error reading line: ", err)
		}

	}
}

func updateCompletions(line *liner.State, completions []string) {
	line.SetCompleter(func(line string) (c []string) {
		for _, n := range completions {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})
}
