package main

import (
	"flag"
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
var precision *int

func main() {
	precision = flag.Int("prec", 32, "precision bits to calculate for")
	isColored := flag.Bool("color", true, "color the output")
	flag.Parse()

	visitor := NewAbacusVisitor()
	line := liner.NewLiner()
	numberPrinter := color.New(color.FgGreen)
	booleanPrinter := color.New(color.FgMagenta)
	defaultPrinter := color.New(color.FgWhite)
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
		savedPrecision := *precision
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
				if *isColored {
					numberPrinter.Println(val.Text('g', *precision))
				} else {
					defaultPrinter.Println(val.Text('g', *precision))
				}
			case string:
				if *isColored {
					numberPrinter.Println(val)
				} else {
					defaultPrinter.Println(val)
				}
			case bool:
				if *isColored {
					booleanPrinter.Println(val)
				} else {
					defaultPrinter.Println(val)
				}
			}
		} else if err == liner.ErrPromptAborted {
			os.Exit(0)
		} else {
			log.Print("Error reading line: ", err)
		}
		*precision = savedPrecision
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
