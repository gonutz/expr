package main

import (
	"testing"

	"github.com/gonutz/check"
)

func TestExpressionCanBeEvaluated(t *testing.T) {
	calc := func(code string, want float64) {
		expr, err := parse(tokenize(code))
		check.Eq(t, err, nil)

		x, err := calculate(expr)
		check.Eq(t, err, nil)
		check.Eq(t, x, want)
	}
	calcErr := func(code string) {
		expr, err := parse(tokenize(code))
		check.Eq(t, err, nil)

		_, err = calculate(expr)
		check.Neq(t, err, nil)
	}

	calc("1", 1)
	calc("2+3", 5)
	calc("9-7", 2)
	calc("4*5", 20)
	calc("15/5", 3)
	calcErr("15/0")
	calc("-1", -1)
	calc("+2", 2)
	calc("-(1+2)", -3)
	calc("-(1 + 2) * ((3+4) * 2)", -42)
	calc(" 4^ 3^2 ", 1<<18)
	calc("123^0", 1)
}
