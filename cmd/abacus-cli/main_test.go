package main

import (
	"github.com/viktordanov/abacus/abacus"
	"testing"

	"github.com/cockroachdb/apd"
)

type testArgs = []struct {
	expr    string
	want    interface{}
	visitor *abacus.Visitor
}

func Test_evaluateExpression(t *testing.T) {
	visitor, err := abacus.NewAbacusVisitor(64, false)
	if err != nil {
		t.Errorf("Error creating visitor: %v", err)
		return
	}
	decimalCtx := apd.BaseContext.WithPrecision(64)

	additionTests := testArgs{
		{expr: "2+2", want: abacus.NewNumber(4), visitor: visitor},
		{expr: "2 +2+   2+2+2+ 2", want: abacus.NewNumber(12), visitor: visitor},
		{expr: "0+ 8+(  d+4)", want: abacus.NewNumber(12), visitor: visitor},
		{expr: "(2+2)", want: abacus.NewNumber(4), visitor: visitor},
		{expr: "2-2", want: abacus.NewNumber(0), visitor: visitor},
		{expr: "2", want: abacus.NewNumber(2), visitor: visitor},
		{expr: "2.139487526 + 9.4777777", want: abacus.NewNumber(11.617265226), visitor: visitor},
	}
	expressionTests := testArgs{
		{expr: "2", want: abacus.NewNumber(2), visitor: visitor},
		{expr: "0.2", want: abacus.NewNumber(0.2), visitor: visitor},
		{expr: ".2", want: abacus.NewNumber(0.2), visitor: visitor},
	}
	multiplicationTests := testArgs{
		{expr: "2*2", want: abacus.NewNumber(4), visitor: visitor},
		{expr: "2*2*2", want: abacus.NewNumber(8), visitor: visitor},
		{expr: "0.188492164*2.444445187", want: newNumFromString(decimalCtx, "0.460758763077014668"), visitor: visitor},
		{expr: "9.00005849491657/0.18849622221346", want: newNumFromString(decimalCtx, "47.74662531286476055056365115305431970338920063933451796995514209"), visitor: visitor},
		{expr: "2*(2+9)", want: abacus.NewNumber(22), visitor: visitor},
		{expr: "(2*2)+9", want: abacus.NewNumber(13), visitor: visitor},
	}
	exponentiationTests := testArgs{
		{expr: "2^2", want: abacus.NewNumber(4), visitor: visitor},
		{expr: "2^2^2**2", want: abacus.NewNumber(256), visitor: visitor},
		{expr: "2^(2**2^2)", want: abacus.NewNumber(65536), visitor: visitor},
		{expr: "2**0", want: abacus.NewNumber(1), visitor: visitor},
	}
	variableTests := testArgs{
		{expr: "d", want: abacus.NewNumber(0), visitor: visitor},
		{expr: "d=1", want: abacus.Assignment{abacus.NewNumber(1)}, visitor: visitor},
		{expr: "d", want: abacus.NewNumber(1), visitor: visitor},
		{expr: "dd=d", want: abacus.Assignment{abacus.NewNumber(1)}, visitor: visitor},
		{expr: "dd", want: abacus.NewNumber(1), visitor: visitor},
	}
	comparisonTests := testArgs{
		{expr: "0 == 0", want: true, visitor: visitor},
		{expr: "0 == 1", want: false, visitor: visitor},
		{expr: "0 > 0", want: false, visitor: visitor},
		{expr: "0 >= 0", want: true, visitor: visitor},
		{expr: "0 >= 1", want: false, visitor: visitor},
		{expr: "1 > 0", want: true, visitor: visitor},
		{expr: "1 >= 0", want: true, visitor: visitor},
		{expr: "0 < 1", want: true, visitor: visitor},
		{expr: "0 <= 1", want: true, visitor: visitor},
		{expr: "0 < 0", want: false, visitor: visitor},
		{expr: "0 <= 0", want: true, visitor: visitor},
		{expr: "0 < 1", want: true, visitor: visitor},
		{expr: "0 <= 1", want: true, visitor: visitor},
		{expr: "1 <= 1", want: true, visitor: visitor},
		{expr: "1 <= 0", want: false, visitor: visitor},
	}

	runTestSuite(t, "Addition & Subtraction", additionTests)
	runTestSuite(t, "Expressions", expressionTests)
	runTestSuite(t, "Multiplication & Division", multiplicationTests)
	runTestSuite(t, "Exponentiation", exponentiationTests)
	runTestSuite(t, "Variables", variableTests)
	runTestSuite(t, "Comparisons", comparisonTests)
}

func BenchmarkEvaluateExpression(b *testing.B) {
	b.StopTimer()
	visitor, err := abacus.NewAbacusVisitor(64, true)
	if err != nil {
		b.Errorf("Error creating visitor: %v", err)
		return
	}

	complicatedExpression := "2+2+sin(4**5**((7/(4**(pi/1.45*e*log(97,1.1))))**8))"

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		evaluateExpression(complicatedExpression, visitor)
	}
}

func runTestSuite(t *testing.T, name string, tests testArgs) {
	t.Run(name, func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.expr, func(t *testing.T) {
				ans := evaluateExpression(tt.expr, tt.visitor)
				switch val := ans.Value.(type) {
				case abacus.Assignment:
					expected, ok := tt.want.(abacus.Assignment)
					if !ok {
						t.Errorf("Assignment %v, unexpected type", ans)
					}
					for i := 0; i < len(val); i++ {
						if !compareNum(val[i], expected[i]) {
							t.Errorf("Assignment %v, want %v", val[i], expected[i])
						}
					}
				case abacus.Number:
					expected, ok := tt.want.(abacus.Number)
					if !ok {
						t.Errorf("Number %v, unexpected type", ans)
					}
					if !compareNum(val, expected) {
						t.Errorf("Number %v, want %v", val, expected)
					}
				case abacus.Bool:
					expected, ok := tt.want.(bool)
					if !ok {
						t.Errorf("Bool %v, unexpected type", ans)
					}
					if bool(val) != expected {
						t.Errorf("Bool %v, want %v", val, expected)
					}
				}
			})
		}
	})
}

func newNumFromString(ctx *apd.Context, s string) abacus.Number {
	n, _, _ := ctx.NewFromString(s)
	return abacus.Number{Decimal: n}
}

func compareNum(left, right abacus.Number) bool {
	return left.Cmp(right.Decimal) == 0
}
