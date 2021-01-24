package main

import (
	"math"
	"math/big"
	"testing"
)

func newBigFloat(val float64) *big.Float {
	return big.NewFloat(val)
}

type testArgs = []struct {
	expr    string
	want    interface{}
	visitor *AbacusVisitor
}

func Test_evaluateExpression(t *testing.T) {
	visitor := NewAbacusVisitor()

	additionTests := testArgs{
		{expr: "2+2", want: newBigFloat(4), visitor: visitor},
		{expr: "2 +2+   2+2+2+ 2", want: newBigFloat(12), visitor: visitor},
		{expr: "0+ 8+(  d+4)", want: newBigFloat(12), visitor: visitor},
		{expr: "(2+2)", want: newBigFloat(4), visitor: visitor},
		{expr: "2-2", want: newBigFloat(0), visitor: visitor},
		{expr: "2", want: newBigFloat(2), visitor: visitor},
		{expr: "2.139487526 + 9.4777777", want: newBigFloat(11.617265226), visitor: visitor},
	}
	multiplicationTests := testArgs{
		{expr: "2*2", want: newBigFloat(4), visitor: visitor},
		{expr: "2*2*2", want: newBigFloat(8), visitor: visitor},
		{expr: "0.188492164*2.444445187", want: newBigFloat(0.460758763077014668), visitor: visitor},
		{expr: "9.00005849491657/0.18849622221346", want: newBigFloat(47.746625312864760550563651153054), visitor: visitor},
		{expr: "2*(2+9)", want: newBigFloat(22), visitor: visitor},
		{expr: "(2*2)+9", want: newBigFloat(13), visitor: visitor},
	}
	exponentiationTests := testArgs{
		{expr: "2^2", want: newBigFloat(4), visitor: visitor},
		{expr: "2^2^2**2", want: newBigFloat(256), visitor: visitor},
		{expr: "2^(2**2^2)", want: newBigFloat(65536), visitor: visitor},
		{expr: "2**0", want: newBigFloat(1), visitor: visitor},
	}
	variableTests := testArgs{
		{expr: "d", want: newBigFloat(0), visitor: visitor},
		{expr: "d=1", want: ResultAssignment{[]*big.Float{New(1)}}, visitor: visitor},
		{expr: "d", want: newBigFloat(1), visitor: visitor},
		{expr: "dd=d", want: ResultAssignment{[]*big.Float{New(1)}}, visitor: visitor},
		{expr: "dd", want: newBigFloat(1), visitor: visitor},
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
	visitor := NewAbacusVisitor()
	complicatedExpression := "2+2+sin(4**5**((7/(4**(pi/1.45*e*log(97,1.1))))**8))"
	for i := 0; i < b.N; i++ {
		evaluateExpression(complicatedExpression, visitor)
	}
}

func runTestSuite(t *testing.T, name string, tests testArgs) {
	t.Run(name, func(t *testing.T) {
		for _, tt := range tests {
			t.Run("", func(t *testing.T) {
				ans := evaluateExpression(tt.expr, tt.visitor)
				switch val := ans.(type) {
				case ResultAssignment:
					expected, ok := tt.want.(ResultAssignment)
					if !ok {
						t.Errorf("VarAssign %v, unexpected type", ans)
					}
					for i := 0; i < len(val.Values); i++ {
						if !compareBigFloat(val.Values[i], expected.Values[i]) {
							t.Errorf("VarAssign %v, want %v", val.Values[i], expected.Values[i])
						}
					}

				case *big.Float:
					expected, ok := tt.want.(*big.Float)
					if !ok {
						t.Errorf("*big.Float %v, unexpected type", ans)
					}
					if !compareBigFloat(val, expected) {
						t.Errorf("*big.Float %v, want %v", val, expected)
					}
				case bool:
					expected, ok := tt.want.(bool)
					if !ok {
						t.Errorf("bool %v, unexpected type", ans)
					}
					if val != expected {
						t.Errorf("bool %v, want %v", val, expected)
					}
				}

			})
		}
	})
}

func compareBigFloat(left *big.Float, right *big.Float) bool {
	a := Zero()
	a.Sub(right, left)
	floatValue, _ := a.Float64()
	return a.Cmp(Zero()) == 0 || (math.Round(floatValue*1e10) == 0)
}
