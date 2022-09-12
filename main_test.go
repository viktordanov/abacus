package main

import (
	"github.com/viktordanov/abacus/abacus"
	"testing"

	"github.com/cockroachdb/apd"
)

type testArgs = []struct {
	expr    string
	want    interface{}
	visitor *abacus.AbacusVisitor
}

func Test_evaluateExpression(t *testing.T) {
	visitor := abacus.NewAbacusVisitor(64)
	abacus.decimalCtx = apd.BaseContext.WithPrecision(64)

	additionTests := testArgs{
		{expr: "2+2", want: abacus.newNumber(4), visitor: visitor},
		{expr: "2 +2+   2+2+2+ 2", want: abacus.newNumber(12), visitor: visitor},
		{expr: "0+ 8+(  d+4)", want: abacus.newNumber(12), visitor: visitor},
		{expr: "(2+2)", want: abacus.newNumber(4), visitor: visitor},
		{expr: "2-2", want: abacus.newNumber(0), visitor: visitor},
		{expr: "2", want: abacus.newNumber(2), visitor: visitor},
		{expr: "2.139487526 + 9.4777777", want: abacus.newNumber(11.617265226), visitor: visitor},
	}
	multiplicationTests := testArgs{
		{expr: "2*2", want: abacus.newNumber(4), visitor: visitor},
		{expr: "2*2*2", want: abacus.newNumber(8), visitor: visitor},
		{expr: "0.188492164*2.444445187", want: newNumFromString(abacus.decimalCtx, "0.460758763077014668"), visitor: visitor},
		{expr: "9.00005849491657/0.18849622221346", want: newNumFromString(abacus.decimalCtx, "47.74662531286476055056365115305431970338920063933451796995514209"), visitor: visitor},
		{expr: "2*(2+9)", want: abacus.newNumber(22), visitor: visitor},
		{expr: "(2*2)+9", want: abacus.newNumber(13), visitor: visitor},
	}
	exponentiationTests := testArgs{
		{expr: "2^2", want: abacus.newNumber(4), visitor: visitor},
		{expr: "2^2^2**2", want: abacus.newNumber(256), visitor: visitor},
		{expr: "2^(2**2^2)", want: abacus.newNumber(65536), visitor: visitor},
		{expr: "2**0", want: abacus.newNumber(1), visitor: visitor},
	}
	variableTests := testArgs{
		{expr: "d", want: abacus.newNumber(0), visitor: visitor},
		{expr: "d=1", want: abacus.Assignment{abacus.newNumber(1)}, visitor: visitor},
		{expr: "d", want: abacus.newNumber(1), visitor: visitor},
		{expr: "dd=d", want: abacus.Assignment{abacus.newNumber(1)}, visitor: visitor},
		{expr: "dd", want: abacus.newNumber(1), visitor: visitor},
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
	runTestSuite(t, "Multiplication & Division", multiplicationTests)
	runTestSuite(t, "Exponentiation", exponentiationTests)
	runTestSuite(t, "Variables", variableTests)
	runTestSuite(t, "Comparisons", comparisonTests)
}

func BenchmarkEvaluateExpression(b *testing.B) {
	visitor := abacus.NewAbacusVisitor(64)
	complicatedExpression := "2+2+sin(4**5**((7/(4**(pi/1.45*e*log(97,1.1))))**8))"
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
	return abacus.Number{n}
}

func compareNum(left, right abacus.Number) bool {
	return left.Cmp(right.Decimal) == 0
}
