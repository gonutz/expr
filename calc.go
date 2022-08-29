package main

import (
	"errors"
	"math"
)

func calculate(e expression) (float64, error) {
	switch x := e.(type) {
	case numberExpr:
		return x.value, nil
	case parenthesisExpr:
		return calculate(x.expr)
	case unaryOpExpr:
		value, err := calculate(x.expr)
		if err != nil {
			return 0, err
		}
		if x.op == '-' {
			value = -value
		}
		return value, nil
	case binaryOpExpr:
		a, err := calculate(x.left)
		if err != nil {
			return 0, err
		}
		b, err := calculate(x.right)
		if err != nil {
			return 0, err
		}

		switch x.op {
		case '+':
			return a + b, nil
		case '-':
			return a - b, nil
		case '*':
			return a * b, nil
		case '/':
			if b == 0 {
				return 0, errors.New("division by zero")
			}
			return a / b, nil
		case '^':
			return math.Pow(a, b), nil
		default:
			return 0, errors.New("unhandled binary operator")
		}
	default:
		return 0, errors.New("unhandled expression type")
	}
}
