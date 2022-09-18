package expr

import (
	"errors"
	"strconv"
	"strings"
)

func parse(tokens []token) (node, error) {
	// stmt : id '=' expr OR expr
	// expr : prod [('+' OR '-') prod]*
	// prod : exp [('*' OR '/') exp]*
	// exp : paren ['^' exp]*
	// paren : ('+' OR '-' OR '') '(' expr ')' OR
	//         ('+' OR '-' OR '') number OR
	//         ('+' OR '-' OR '') id

	next := func() token {
		t := tokens[0]
		tokens = tokens[1:]
		return t
	}

	hasNext := func() bool {
		return len(tokens) > 0
	}

	peek := func() token {
		return tokens[0]
	}

	var parseExpression func() (expression, error)

	parseParen := func() (expression, error) {
		var sign rune

		if strings.Contains("+-", peek().text) {
			sign = rune(next().text[0])
		}

		var result expression
		if peek().text == "(" {
			next()
			expr, _ := parseExpression()
			next() // Skip ')'.
			result = paren(expr)
		} else if peek().kind == number {
			x, err := strconv.ParseFloat(next().text, 64)
			if err != nil {
				return nil, errors.New("number expected")
			}
			result = numberExpr{value: x}
		} else if peek().kind == identifier {
			result = id(next().text)
		} else {
			return nil, errors.New("invalid parenthesis expression")
		}

		if sign != 0 {
			result = unary(sign, result)
		}

		return result, nil
	}

	var parseExp func() (expression, error)
	parseExp = func() (expression, error) {
		if !hasNext() {
			return nil, errors.New("unexpected end")
		}

		exp, err := parseParen()
		if err != nil {
			return nil, err
		}

		for hasNext() && peek().text == "^" {
			next()
			right, err := parseExp()
			if err != nil {
				return nil, err
			}
			exp = binaryOp(exp, '^', right)
		}

		return exp, nil
	}

	parseProduct := func() (expression, error) {
		if !hasNext() {
			return nil, errors.New("unexpected end")
		}

		factor, err := parseExp()
		if err != nil {
			return nil, err
		}

		for hasNext() && strings.Contains("*/", peek().text) {
			op := rune(next().text[0])
			right, err := parseExp()
			if err != nil {
				return nil, err
			}
			factor = binaryOp(factor, op, right)
		}

		return factor, nil
	}

	parseExpression = func() (expression, error) {
		summand, err := parseProduct()
		if err != nil {
			return nil, errors.New("invalid expression start")
		}
		for hasNext() && strings.Contains("+-", peek().text) {
			op := rune(next().text[0])
			right, err := parseProduct()
			if err != nil {
				return nil, err
			}
			summand = binaryOp(summand, op, right)
		}
		return summand, nil
	}

	parseStatement := func() (node, error) {
		if len(tokens) >= 2 &&
			tokens[0].kind == identifier && tokens[1].kind == '=' {
			name := next().text
			next() // Skip '='.
			value, err := parseExpression()
			if err != nil {
				return nil, err
			}
			return assign(name, value), nil
		} else {
			return parseExpression()
		}
	}

	return parseStatement()
}

type node interface{}

type expression interface{}

type binaryOpExpr struct {
	op    rune
	left  expression
	right expression
}

func binaryOp(left expression, op rune, right expression) binaryOpExpr {
	return binaryOpExpr{
		op:    op,
		left:  left,
		right: right,
	}
}

type parenthesisExpr struct {
	expr expression
}

func paren(e expression) parenthesisExpr {
	return parenthesisExpr{expr: e}
}

type numberExpr struct {
	value float64
}

func num(x float64) numberExpr {
	return numberExpr{value: x}
}

type unaryOpExpr struct {
	op   rune
	expr expression
}

func unary(op rune, expr expression) unaryOpExpr {
	return unaryOpExpr{op: op, expr: expr}
}

type assignment struct {
	name  string
	value expression
}

func assign(name string, value expression) assignment {
	return assignment{name: name, value: value}
}

type identifierExpr string

func id(name string) identifierExpr {
	return identifierExpr(name)
}
