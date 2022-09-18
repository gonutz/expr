package expr

import (
	"strconv"
	"testing"

	"github.com/gonutz/check"
)

func TestExpressionCanBeEvaluated(t *testing.T) {
	calc := func(codeAndAnswer ...string) {
		t.Helper()

		n := len(codeAndAnswer)
		codes, want := codeAndAnswer[:n-1], codeAndAnswer[n-1]

		c := NewCalculator()
		var lastAnswer string
		for _, code := range codes {
			a, err := c.Evaluate(code)
			check.Eq(t, err, nil)
			lastAnswer = a
		}

		check.Eq(t, lastAnswer, want)
	}
	calcErr := func(code string) {
		t.Helper()

		_, err := NewCalculator().Evaluate(code)
		check.Neq(t, err, nil)
	}

	calc("1", "1")
	calc("2+3", "5")
	calc("9-7", "2")
	calc("4*5", "20")
	calc("15/5", "3")
	calcErr("15/0")
	calc("-1", "-1")
	calc("+2", "2")
	calc("-(1+2)", "-3")
	calc("-(1 + 2) * ((3+4) * 2)", "-42")
	calc(" 4^ 3^2 ", strconv.Itoa(1<<18))
	calc("123^0", "1")
	calc("x = 1", "x = 1")
	calc("y = 5*6", "y = 30")
	calc(
		"x = 2",
		"x",
		"2",
	)
	calc(
		"a = 5",
		"b = 7",
		"-a + -b",
		"-12",
	)
	calcErr("x")
}
