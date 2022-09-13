package abacus

import (
	"errors"
	"github.com/cockroachdb/apd"
	"strconv"
	"strings"
)

// Result is struct returned after evaluating each leaf in the parser tree.
// Errors are propagated up
type Result struct {
	Value  Formatter
	Errors []error
}

func NewResult(value Formatter) *Result {
	return &Result{Value: value, Errors: make([]error, 0)}
}

func (r *Result) WithError(ee ...string) *Result {
	for _, e := range ee {
		r.Errors = append(r.Errors, errors.New(e))
	}
	return r
}

func (r *Result) WithErrors(previous *Result, ee ...string) *Result {
	if previous != nil {
		for _, e := range previous.Errors {
			r.Errors = append(r.Errors, e)
		}
	}
	for _, e := range ee {
		r.Errors = append(r.Errors, errors.New(e))
	}
	return r
}

func decimalsToStringList(dd []Number) string {
	var b strings.Builder
	count := len(dd)
	if count > 1 {
		b.WriteRune('(')
	}
	for i, d := range dd {
		b.WriteString(d.Text('g'))
		if i+1 != count {
			b.WriteString(", ")
		}
	}

	if count > 1 {
		b.WriteRune(')')
	}
	return b.String()
}

func stringsToStringList(ss []String) string {
	var b strings.Builder
	count := len(ss)
	if count > 1 {
		b.WriteRune('(')
	}
	for i, s := range ss {
		b.WriteString(string(s))
		if i+1 != count {
			b.WriteString(", ")
		}
	}

	if count > 1 {
		b.WriteRune(')')
	}
	return b.String()
}

func mixedToStringList(dd MixedTuple) string {
	var b strings.Builder
	count := len(dd)
	if count > 1 {
		b.WriteRune('(')
	}
	for i, d := range dd {
		switch val := d.(type) {
		case String:
			b.WriteString(string(val))
			if i+1 != count {
				b.WriteString(", ")
			}
		case Number:
			b.WriteString(val.Text('g'))
			if i+1 != count {
				b.WriteString(", ")
			}
		}
	}

	if count > 1 {
		b.WriteRune(')')
	}
	return b.String()
}

func addColor(str string, color Color) string {
	var b strings.Builder
	b.WriteString(string(color))
	b.WriteString(str)
	b.WriteString(string(Reset))
	return b.String()
}

type Assignment []Number

func (r Assignment) String() string {
	return decimalsToStringList(r)
}

func (r Assignment) Color() string {
	return addColor(r.String(), Magenta)
}

type LambdaAssignment string

func (r LambdaAssignment) String() string {
	return string(r)
}
func (r LambdaAssignment) Color() string {
	return addColor(r.String(), Magenta)
}

type Tuple []Number

func (r Tuple) String() string {
	return decimalsToStringList(r)
}

func (r Tuple) Color() string {
	return addColor(r.String(), Green)
}

type MixedTuple []interface{}

func (r MixedTuple) String() string {
	return mixedToStringList(r)
}

func (r MixedTuple) Color() string {
	return addColor(r.String(), Green)
}

type VariablesTuple []String

func (r VariablesTuple) String() string {
	return stringsToStringList(r)
}

func (r VariablesTuple) Color() string {
	return addColor(r.String(), Magenta)
}

type LambdaArguments []String

func (r LambdaArguments) String() string {
	return stringsToStringList(r)
}

// Color variables and Lambdas differently
func (r LambdaArguments) Color() string {
	return addColor(r.String(), Magenta)
}

type Number struct {
	*apd.Decimal
}

func (r Number) String() string {
	return r.Text('g')
}

func (r Number) Color() string {
	return addColor(r.String(), Green)
}

type Bool bool

func (r Bool) String() string {
	return strconv.FormatBool(bool(r))
}
func (r Bool) Color() string {
	return addColor(r.String(), Green)
}

type String string

func (r String) String() string {
	return string(r)
}
func (r String) Color() string {
	return addColor(r.String(), Green)
}

// Length returns the length of the tuple value of result.
// If Value is not a tuple type 1 is returned
func (r *Result) Length() int {
	switch val := r.Value.(type) {
	case VariablesTuple:
		return len(val)
	case LambdaArguments:
		return len(val)
	case Tuple:
		return len(val)
	}
	return 1
}
