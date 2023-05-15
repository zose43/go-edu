package eval

import (
	"fmt"
	"math"
	"strings"
)

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

type Env map[Var]float64

type Var string

type literal float64

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}

var numParams = map[string]int{"pow": 2, "sqrt": 1, "sin": 1}

func (c call) Check(vars map[Var]bool) error {
	arg, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("invalid call %q", c.fn)
	}
	if arg != len(c.args) {
		return fmt.Errorf("not equal args count %d vs %d in %q", len(c.args), arg, c.fn)
	}
	for _, expr := range c.args {
		if err := expr.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("operator %q not exist", u.op)
	}
	return u.x.Check(vars)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	default:
		panic(fmt.Sprintf("Invalid operator type %q", u.op))
	}
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("/*-+", b.op) {
		return fmt.Errorf("operator %q not exist", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	default:
		panic(fmt.Sprintf("Invalid operator type %q or %q", b.x, b.y))
	}
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	default:
		panic(fmt.Sprintf("Invalid function %q", c.fn))
	}
}
