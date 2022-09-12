package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/peterh/liner"
	"github.com/thecodeteam/goodbye"
	"github.com/viktordanov/abacus/parser"
)

var (
	arguments   args
	homeDir, _  = os.UserHomeDir()
	historyFile = filepath.Join(homeDir, ".abacus_history")
	funcs       = []string{
		"sqrt(", "cbrt(", "ln(", "log(", "log2(", "log10(", "floor(", "ceil(", "exp(", "sin(", "cos(", "tan(", "sign(", "abs(", "round(", "min(", "max(", "avg(", "from(", "until(", "reverse(", "nth(", "pi", "e", "phi",
	}
)

type StringSliceArg []string

func (s *StringSliceArg) UnmarshalText(text []byte) error {
	raw := string(text)
	raw = strings.ReplaceAll(raw, " ", "")

	*s = strings.Split(raw, ",")
	variableNameMatcher := regexp.MustCompile(`^[a-z]\w*$`).MatchString
	for _, variableName := range *s {
		if !variableNameMatcher(variableName) {
			return errors.New("variable names must begin with a latin lowercase letter and consist of a-zA-Z0-9")
		}
	}
	return nil
}

type args struct {
	IgnoreColor           bool           `arg:"-n,--no-color" help:"disable color in output" default:"false"`
	AllowCopy             bool           `arg:"--allow-copy" help:"Ctrl-C will copy current expression (if present) or last answer instead of aborting" default:"false"`
	Strict                bool           `arg:"--strict" help:"prohibit use of undefined lambdas and variables" default:"false"`
	Precision             uint32         `arg:"-p,--precision" help:"precision for calculations" default:"64"`
	EvalExpressionAndExit string         `arg:"-e,--eval" help:"evaluate expression and exit"`
	ImportDefinitions     string         `arg:"-i,--import" help:"import statements from file and continue"`
	CustomPromptSymbol    string         `arg:"--prompt-symbol" help:"custom prompt symbol" default:"> " `
	LastAnswerVariables   StringSliceArg `arg:"--answer-vars" help:"custom last answer variable names" default:"ans,answer"`
}

func (args) Version() string {
	return "v1.3.0\n"
}
func (args) Description() string {
	return "abacus - a simple interactive calculator CLI with support for variables, lambdas, comparison checks, and math functions\n"
}

func main() {
	if err := run(); err != nil {
		_, err = fmt.Fprintf(os.Stderr, "%s\n", err)
		if err != nil {
			log.Fatal(err.Error())
		}
		os.Exit(1)
	}
}

func run() error {
	arg.MustParse(&arguments)
	abacusVisitor := NewAbacusVisitor(arguments.Precision)
	line := liner.NewLiner()

	defer line.Close()
	line.SetCtrlCAborts(!arguments.AllowCopy)

	if _, err := os.Stat(historyFile); os.IsNotExist(err) {
		_, err = os.Create(historyFile)
		if err != nil {
			return err
		}
	}

	f, err := os.Open(historyFile)
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = line.ReadHistory(f)
	if err != nil {
		return err
	}

	handleAndPrintAnswer := func(res *Result) {
		if len(res.Errors) != 0 {
			for _, e := range res.Errors {
				fmt.Print(Red)
				fmt.Print(e.Error())
				fmt.Println(Reset)
			}
			return
		}

		switch rawValue := res.Value.(type) {
		case Assignment:
			updateCompletions(line, abacusVisitor)
		case LambdaAssignment:
			updateCompletions(line, abacusVisitor)
		case Number:
			for _, variableName := range arguments.LastAnswerVariables {
				abacusVisitor.variables[variableName] = rawValue
			}
			updateCompletions(line, abacusVisitor)
		}

		if arguments.IgnoreColor {
			fmt.Println(res.Value.String())
		} else {
			fmt.Println(res.Value.Color())
		}
	}

	if len(arguments.ImportDefinitions) != 0 {
		dat, err := ioutil.ReadFile(arguments.ImportDefinitions)
		if err != nil {
			return err
		}
		evaluateExpression(string(dat), abacusVisitor)
	}

	if len(arguments.EvalExpressionAndExit) != 0 {
		handleAndPrintAnswer(evaluateExpression(arguments.EvalExpressionAndExit, abacusVisitor))
		return nil
	}

	updateCompletions(line, abacusVisitor)

	// Ensure files and REPL are closed in the event of an unexpected exit
	ctx := context.Background()
	defer goodbye.Exit(ctx, -1)
	goodbye.Notify(ctx)
	goodbye.Register(func(ctx context.Context, sig os.Signal) {
		_ = writeHistoryFile(line)
		_ = line.Close()
		_ = f.Close()
	})

	// Read Eval Print Loop
	for {
		savedPrecision := arguments.Precision
		input, err := line.Prompt(arguments.CustomPromptSymbol)
		if err != nil {
			if errors.Is(err, liner.ErrPromptAborted) || errors.Is(err, io.EOF) {
				return writeHistoryFile(line)
			}
			log.Print("Error reading line: ", err)
		}

		if strings.Contains(input, "quit") {
			return writeHistoryFile(line)
		}

		if input != "" {
			line.AppendHistory(input)
			handleAndPrintAnswer(evaluateExpression(input, abacusVisitor))
			arguments.Precision = savedPrecision
		}
	}
}

func evaluateExpression(expr string, visitor *AbacusVisitor) *Result {
	result := NewResult(nil).WithErrors(nil, "expression did not yield a result")
	ok := false

	expressions := strings.Split(expr, ";")
	for i, expression := range expressions {
		var b strings.Builder
		for _, char := range expression {
			if char != '\n' {
				b.WriteRune(char)
			}
		}
		expressions[i] = b.String()
	}

	for _, e := range expressions {
		if len(e) == 0 {
			continue
		}
		is := antlr.NewInputStream(e)
		lexer := parser.NewAbacusLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := parser.NewAbacusParser(stream)
		p.BuildParseTrees = true
		tree := p.Root()
		t := visitor.Visit(tree)
		result, ok = t.(*Result)
		if !ok {
			return NewResult(nil).WithErrors(nil, "expression did not yield a result")
		}
	}
	return result
}

func writeHistoryFile(line *liner.State) error {
	f, err := os.Create(historyFile)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = line.WriteHistory(f)
	return err
}

func updateCompletions(line *liner.State, a *AbacusVisitor) {
	completions := make([]string, 0)
	completions = append(completions, funcs...)
	for k := range a.variables {
		completions = append(completions, k)
	}

	for k := range a.lambdaDeclarations {
		completions = append(completions, k+"(")
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
			if strings.HasPrefix(n, line[idx:]) {
				c = append(c, line[0:idx]+n)
			}
		}
		return
	})
}
