package eval

import (
	"math"
	"testing"
)

func TestVar_Eval(t *testing.T) {
	// todo no parse method
	tests := []struct {
		name string
		v    Var
		env  Env
		want float64
	}{
		{"Sqrt",
			"sqrt(a / pi)",
			Env{"a": 87616, "pi": math.Pi},
			167},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Eval(tt.env); got != tt.want {
				t.Errorf("Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
