package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/peterh/liner"
	"github.com/thecodeteam/goodbye"
	"github.com/viktordanov/abacus/parser"

	"math/big"
)

var (
	precision   uint
	homeDir, _  = os.UserHomeDir()
	historyFile = filepath.Join(homeDir, ".abacus_history")
	funcs       = []string{
		"sqrt(", "ln(", "log(", "log2(", "log10(", "floor(", "ceil(", "exp(", "sin(", "cos(", "tan(", "round(", "min(", "max(", "pi", "e", "phi",
	}
)

type variableAssignment struct {
	newValue *big.Float
}

type args struct {
	IgnoreColor bool   `arg:"-n,--no-color" help:"disable color in output" default:"false"`
	Precision   uint   `arg:"-p,--precision" help:"precision for calculations" default:"64"`
	Expression  string `arg:"-e,--eval" help:"evaluate expression and exit"`
}

func (args) Version() string {
	return "v1.0.0\n"
}
func (args) Description() string {
	return "abacus - a simple interactive calculator CLI with support for variables, comparison checks, and math functions\n"
}

func green(arg string) {
	fmt.Println(string("\033[32m") + arg + string("\033[0m"))
}
func magenta(arg string) {
	fmt.Println(string("\033[35m") + arg + string("\033[0m"))
}
func white(arg string) {
	fmt.Println(string("\033[37m") + arg + string("\033[0m"))
}

func main() {
	var arguments args
	arg.MustParse(&arguments)
	precision = arguments.Precision

	visitor := NewAbacusVisitor()
	line := liner.NewLiner()

	defer line.Close()
	line.SetCtrlCAborts(true)

	if f, err := os.Open(historyFile); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	printAnswer := func(ans interface{}) {
		switch val := ans.(type) {
		case variableAssignment:
			updateCompletions(line, visitor)
			if !arguments.IgnoreColor {
				green(val.newValue.Text('g', int(precision)))
			} else {
				white(val.newValue.Text('g', int(precision)))
			}
		case *big.Float:
			if !arguments.IgnoreColor {
				green(val.Text('g', int(precision)))
			} else {
				white(val.Text('g', int(precision)))
			}
		case string:
			if !arguments.IgnoreColor {
				green(val)
			} else {
				white(val)
			}
		case bool:
			if !arguments.IgnoreColor {
				magenta(strconv.FormatBool(val))
			} else {
				white(strconv.FormatBool(val))
			}
		}
	}

	if len(arguments.Expression) != 0 {
		printAnswer(evaluateExpression(arguments.Expression, visitor))
		os.Exit(0)
	}
	updateCompletions(line, visitor)

	ctx := context.Background()
	defer goodbye.Exit(ctx, -1)
	goodbye.Notify(ctx)
	goodbye.Register(func(ctx context.Context, sig os.Signal) {
		writeHistoryFile(line)
	})

	for {
		savedPrecision := precision
		if input, err := line.Prompt("> "); err == nil {
			if strings.Index(input, "quit") != -1 {
				writeHistoryFile(line)
				os.Exit(0)
			}
			line.AppendHistory(input)
			ans := evaluateExpression(input, visitor)
			printAnswer(ans)
		} else if errors.Is(err, liner.ErrPromptAborted) || errors.Is(err, io.EOF) {
			writeHistoryFile(line)
			os.Exit(0)
		} else {
			log.Print("Error reading line: ", err)
		}
		precision = savedPrecision
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
