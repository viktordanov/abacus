package abacus

import (
	"github.com/viktordanov/abacus/parser"
	"strings"
)

type LambdaDeclaration struct {
	ctx       *parser.LambdaDeclarationContext
	arguments LambdaArguments
	argSet    map[string]bool
}

type Parameter struct {
	Name  string
	Value interface{}
}

type RecursionParameters struct {
	MaxRecurrences uint
	LastValue      parser.IExpressionContext
	StopWhen       parser.IBoolExpressionContext
	Memoize        bool
}

func NewRecursionParameters() *RecursionParameters {
	return &RecursionParameters{MaxRecurrences: 0, LastValue: nil, StopWhen: nil, Memoize: false}
}

type CalledLambda struct {
	ctx       *parser.LambdaDeclarationContext
	parent    *CalledLambda
	children  []*CalledLambda
	arguments map[string]interface{}
	nested    bool
	name      string
}

func (l *CalledLambda) Signature() string {
	var b strings.Builder
	b.WriteString(l.name)
	b.WriteRune('_')
	for _, i := range l.arguments {
		switch val := i.(type) {
		case Number:
			b.WriteString(val.String())
		case String:
			b.WriteString(val.String())
		}
		b.WriteRune('_')
	}
	return b.String()
}

type LambdaCallStack struct {
	root      *CalledLambda
	trace     []*CalledLambda
	invokes   map[string]uint
	recursion map[string]*RecursionParameters
	memoized  map[string]*Result
}

func (s *LambdaCallStack) Peek() *CalledLambda {
	if len(s.trace) == 0 {
		return nil
	}
	return s.trace[len(s.trace)-1]
}

func (s *LambdaCallStack) Pop() *CalledLambda {
	if len(s.trace) == 0 {
		return nil
	}
	el := s.Peek()
	s.trace = s.trace[:len(s.trace)-1]
	return el
}

func (s *LambdaCallStack) Push(l *CalledLambda) {
	s.trace = append(s.trace, l)
}

func (s *LambdaCallStack) IsRecurring(lambdaName string) bool {
	for i := len(s.trace) - 1; i >= 0; i-- {
		if s.trace[i].name == lambdaName {
			return true
		}
	}

	return false
}
