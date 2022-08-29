package main

import (
	"testing"

	"github.com/gonutz/check"
)

func TestStringIsTokenized(t *testing.T) {
	lex := func(code string, want ...string) {
		t.Helper()
		tokens := tokenize(code)
		texts := make([]string, len(tokens))
		for i := range tokens {
			texts[i] = tokens[i].text
		}
		check.Eq(t, texts, want)
	}

	lex("")
	lex(" \t\n\r ")
	lex("1", "1")
	lex("456", "456")
	lex("78.12", "78.12")
	lex("+-*/^()", "+", "-", "*", "/", "^", "(", ")")
	lex(
		"1 + (2*3.456) - -78.90 /12.45",
		"1", "+", "(", "2", "*", "3.456", ")", "-", "-", "78.90", "/", "12.45",
	)
	lex("1??2", "1", "?", "?", "2")
}

func TestTokensHaveRightKind(t *testing.T) {
	lex := func(code, wantTokens string) {
		t.Helper()
		tokens := tokenize(code)
		have := ""
		for i := range tokens {
			have += string(tokens[i].kind)
		}
		check.Eq(t, have, wantTokens)
	}

	lex("", "")
	lex(" \n\t\r ", "")
	lex("1", "n")
	lex("1 üä 2", "niin")
	lex("+-*/()^ 123", "+-*/()^n")
}
