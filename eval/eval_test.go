package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestVar_Eval(t *testing.T) {
	tests := []struct {
		name string
		expr string
		env  Env
		want string
	}{
		{"Sqrt",
			"sqrt(a / pi)",
			Env{"a": 87616, "pi": math.Pi},
			"167"},
		{
			"pow",
			"pow(x,3)+pow(y,3)",
			Env{"x": 12, "y": 1},
			"1729",
		},
		{
			"pow",
			"pow(x,3)+pow(y,3)",
			Env{"x": 9, "y": 10},
			"1729",
		},
		{
			"degree",
			"5/9*(F-32)",
			Env{"F": -40},
			"-40",
		},
		{
			"degree",
			"5/9*(F-32)",
			Env{"F": 32},
			"0",
		},
		{
			"degree",
			"5/9*(F-32)",
			Env{"F": 212},
			"100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := Parse(tt.expr)
			if err != nil {
				t.Errorf("Parse invalid %s", err)
			} else {
				fmt.Printf("\t%v => %s\n", tt.env, tt.want)
				got := fmt.Sprintf("%.6g", expr.Eval(tt.env))
				if got != tt.want {
					t.Errorf("Eval() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
