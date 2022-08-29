package main

import (
	"testing"

	"github.com/gonutz/check"
)

func TestTokensMakeTree(t *testing.T) {
	checkParse := func(code string, want expression) {
		t.Helper()
		expr, err := parse(tokenize(code))
		check.Eq(t, err, nil)
		check.Eq(t, expr, want)
	}

	checkParse("1", num(1))

	checkParse("2+3", binaryOp(num(2), '+', num(3)))

	checkParse(
		"4-5+6",
		binaryOp(
			binaryOp(num(4), '-', num(5)),
			'+',
			num(6),
		),
	)

	checkParse("1*2", binaryOp(num(1), '*', num(2)))

	checkParse(
		"1 + 2*3 - 4/5",
		binaryOp(
			binaryOp(
				num(1),
				'+',
				binaryOp(num(2), '*', num(3)),
			),
			'-',
			binaryOp(num(4), '/', num(5)),
		),
	)

	checkParse(
		"(1+2)*3",
		binaryOp(
			paren(binaryOp(num(1), '+', num(2))),
			'*',
			num(3),
		),
	)

	checkParse("+1", unary('+', num(1)))

	checkParse("-1", unary('-', num(1)))

	checkParse(
		"-(+1 - 2)",
		unary(
			'-',
			paren(binaryOp(
				unary('+', num(1)),
				'-',
				num(2),
			)),
		),
	)

	checkParse("1^2", binaryOp(num(1), '^', num(2)))

	checkParse("1^2^3", binaryOp(
		num(1),
		'^',
		binaryOp(num(2), '^', num(3)),
	))
}
