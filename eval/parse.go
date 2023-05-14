package eval

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

type lexerPanic string

type lexer struct {
	scan  scanner.Scanner
	token rune // lookahead token
}

func (l *lexer) describe() string {
	switch l.token {
	case scanner.EOF:
		return "End of file"
	case scanner.Ident:
		return fmt.Sprintf("Identifier %s", l.text())
	case scanner.Float, scanner.Int:
		return fmt.Sprintf("Number %s", l.text())
	default:
		return fmt.Sprintf("Other rune %q", l.token)
	}
}

func (l *lexer) text() string {
	return l.scan.TokenText()
}

func (l *lexer) next() {
	l.token = l.scan.Scan()
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '/', '*':
		return 2
	}
	return 0
}

func parseExpr(lex *lexer) Expr {
	return parseBinary(lex, 1)
}

func parseUnary(lex *lexer) Expr {
	if lex.token == '+' || lex.token == '-' {
		op := lex.token
		lex.next()
		return unary{op: op, x: parseUnary(lex)}
	}
	return parsePrimary(lex)
}

func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	case scanner.Ident:
		x := lex.text()
		lex.next()
		if lex.token != '(' {
			return Var(x)
		}
		lex.next()
		var args []Expr
		if lex.token != ')' {
			for {
				args = append(args, parseExpr(lex))
				if lex.token != ',' {
					break
				}
				lex.next()
				if lex.token != ')' {
					msg := fmt.Sprintf("Got %s, want ')'", lex.describe())
					panic(lexerPanic(msg))
				}
			}
		}
		lex.next()
		return call{fn: x, args: args}
	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexerPanic(err.Error()))
		}
		lex.next()
		return literal(f)
	case ')':
		lex.next()
		x := parseExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("Got %s, want ')'", lex.describe())
			panic(lexerPanic(msg))
		}
		lex.next()
		return x
	}
	msg := fmt.Sprintf("Unexpected token %s", lex.describe())
	panic(lexerPanic(msg))
}

func parseBinary(lex *lexer, prec int) Expr {
	lhs := parseUnary(lex)
	for curPrec := precedence(lex.token); curPrec >= prec; curPrec-- {
		for precedence(lex.token) == curPrec {
			op := lex.token
			lex.next() // consume operator
			rhs := parseBinary(lex, curPrec+1)
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs
}

func Parse(input string) (_ Expr, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
		// no panic
		case lexerPanic:
			err = fmt.Errorf("%s", x)
		default:
			panic(x)
		}
	}()

	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanInts
	lex.next()
	x := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return x, nil
}
