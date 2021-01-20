package main

import (
	"flag"
	"fmt"
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

var (
	precision   *int
	homeDir, _  = os.UserHomeDir()
	historyFile = filepath.Join(homeDir, ".abacus_history")
	funcs       = []string{
		"sqrt(", "ln(", "log(", "log2(", "log10(", "floor(", "ceil(", "exp(", "sin(", "cos(", "tan(", "round(", "min(", "max(", "pi", "e", "phi",
	}
)

type variableAssignment struct {
	newValue *big.Float
}

func main() {
	precision = flag.Int("prec", 32, "precision bits to calculate for")
	isColored := flag.Bool("color", true, "color the output")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "abacus - a simple interactive calculator CLI with support for variables, comparison checks, and math functions\nUsage: ")
		flag.PrintDefaults()
	}
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

	writeHistoryFile(line)
	updateCompletions(line, visitor)

	for {
		savedPrecision := *precision
		if input, err := line.Prompt("> "); err == nil {
			line.AppendHistory(input)
			ans := evaluateExpression(input, visitor)
			switch val := ans.(type) {
			case variableAssignment:
				updateCompletions(line, visitor)
				if *isColored {
					numberPrinter.Println(val.newValue.Text('g', *precision))
				} else {
					defaultPrinter.Println(val.newValue.Text('g', *precision))
				}
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
			writeHistoryFile(line)
			os.Exit(0)
		} else {
			log.Print("Error reading line: ", err)
		}
		*precision = savedPrecision
	}
}

func evaluateExpression(expr string, visitor *AbacusVisitor) (ans interface{}) {
	is := antlr.NewInputStream(expr)
	lexer := parser.NewAbacusLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewAbacusParser(stream)
	tree := p.Root()
	ans = visitor.Visit(tree)
	return
}

func writeHistoryFile(line *liner.State) {
	if f, err := os.Create(historyFile); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
}

func updateCompletions(line *liner.State, a *AbacusVisitor) {
	completions := make([]string, 0)
	completions = append(completions, funcs...)
	for k := range a.vars {
		completions = append(completions, k)
	}

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range completions {

			var idx int
			for idx = len(line) - 1; idx >= 0; idx-- {
				r := rune(line[idx])
				if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
					continue
				}
				idx++
				break
			}
			if len(line) == 0 {
				c = append(c, n)
				continue
			}
			if idx == -1 {
				idx = 0
			}
			if strings.HasPrefix(n, strings.ToLower(line[idx:])) {
				c = append(c, line[0:idx]+n)
			}
		}
		return
	})
}
