package main

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

func decimalsToStringList(dd []ResultNumber) string {
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

func stringsToStringList(ss []ResultString) string {
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

func mixedToStringList(dd ResultMixedTuple) string {
	var b strings.Builder
	count := len(dd)
	if count > 1 {
		b.WriteRune('(')
	}
	for i, d := range dd {
		switch val := d.(type) {
		case ResultString:
			b.WriteString(string(val))
			if i+1 != count {
				b.WriteString(", ")
			}
		case ResultNumber:
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

type ResultAssignment []ResultNumber

func (r ResultAssignment) String() string {
	return decimalsToStringList(r)
}

func (r ResultAssignment) Color() string {
	return addColor(r.String(), Magenta)
}

type ResultLambdaAssignment string

func (r ResultLambdaAssignment) String() string {
	return string(r)
}
func (r ResultLambdaAssignment) Color() string {
	return addColor(r.String(), Magenta)
}

type ResultTuple []ResultNumber

func (r ResultTuple) String() string {
	return decimalsToStringList(r)
}

func (r ResultTuple) Color() string {
	return addColor(r.String(), Green)
}

type ResultMixedTuple []interface{}

func (r ResultMixedTuple) String() string {
	return mixedToStringList(r)
}

func (r ResultMixedTuple) Color() string {
	return addColor(r.String(), Green)
}

type ResultVariablesTuple []ResultString

func (r ResultVariablesTuple) String() string {
	return stringsToStringList(r)
}

func (r ResultVariablesTuple) Color() string {
	return addColor(r.String(), Magenta)
}

type ResultLambdaArguments []ResultString

func (r ResultLambdaArguments) String() string {
	return stringsToStringList(r)
}

// Color variables and Lambdas differently
func (r ResultLambdaArguments) Color() string {
	return addColor(r.String(), Magenta)
}

type ResultNumber struct {
	*apd.Decimal
}

func (r ResultNumber) String() string {
	return r.Text('g')
}

func (r ResultNumber) Color() string {
	return addColor(r.String(), Green)
}

type ResultBool bool

func (r ResultBool) String() string {
	return strconv.FormatBool(bool(r))
}
func (r ResultBool) Color() string {
	return addColor(r.String(), Green)
}

type ResultString string

func (r ResultString) String() string {
	return string(r)
}
func (r ResultString) Color() string {
	return addColor(r.String(), Green)
}

// Length returns the length of the tuple value of result.
// If Value is not a tuple type 1 is returned
func (r *Result) Length() int {
	switch val := r.Value.(type) {
	case ResultVariablesTuple:
		return len(val)
	case ResultLambdaArguments:
		return len(val)
	case ResultTuple:
		return len(val)
	}
	return 1
}
