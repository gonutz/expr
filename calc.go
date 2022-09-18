package expr

import (
	"errors"
	"fmt"
	"math"
)

func NewCalculator() *Calculator {
	return &Calculator{
		symbols: make(map[string]float64),
	}
}

type Calculator struct {
	symbols map[string]float64
}

func (c *Calculator) Evaluate(code string) (string, error) {
	n, err := parse(tokenize(code))
	if err != nil {
		return "", err
	}

	switch x := n.(type) {
	case assignment:
		value, err := c.calculate(x.value)
		if err != nil {
			return "", err
		}
		c.symbols[x.name] = value
		return fmt.Sprint(x.name, " = ", value), nil
	default:
		res, err := c.calculate(x)
		return fmt.Sprint(res), err
	}
}

func (c *Calculator) calculate(e expression) (float64, error) {
	switch x := e.(type) {
	case numberExpr:
		return x.value, nil
	case identifierExpr:
		name := string(x)
		if val, ok := c.symbols[name]; ok {
			return val, nil
		}
		return 0, fmt.Errorf("unknown variable: %s", name)
	case parenthesisExpr:
		return c.calculate(x.expr)
	case unaryOpExpr:
		value, err := c.calculate(x.expr)
		if err != nil {
			return 0, err
		}
		if x.op == '-' {
			value = -value
		}
		return value, nil
	case binaryOpExpr:
		a, err := c.calculate(x.left)
		if err != nil {
			return 0, err
		}
		b, err := c.calculate(x.right)
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
