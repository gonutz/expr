package expr

import (
	"strings"
	"unicode"
)

func tokenize(code string) []token {
	runes := []rune(code)
	var tokens []token
	tokenStart := 0
	runePos := -1
	var cur rune

	atEnd := func() bool {
		return runePos >= len(runes)
	}

	next := func() {
		runePos++
		cur = 0
		if runePos < len(runes) {
			cur = runes[runePos]
		}
	}

	is := func(r rune) bool {
		return cur == r
	}

	isOneOf := func(s string) bool {
		return strings.ContainsRune(s, cur)
	}

	isSpace := func() bool {
		return unicode.IsSpace(cur)
	}

	isDigit := func() bool {
		return unicode.IsDigit(cur)
	}

	isLetter := func() bool {
		return cur == '_' || unicode.IsLetter(cur)
	}

	emit := func(kind tokenKind) {
		tokens = append(tokens, token{
			text: string(runes[tokenStart:runePos]),
			kind: kind,
		})
		tokenStart = runePos
	}

	skip := func() {
		next()
		tokenStart = runePos
	}

	next()
	for !atEnd() {
		if isOneOf("+-*/()^=") {
			kind := tokenKind(cur)
			next()
			emit(kind)
		} else if isSpace() {
			skip()
		} else if isDigit() {
			next()
			for isDigit() {
				next()
			}
			if is('.') {
				next()
				for isDigit() {
					next()
				}
			}
			emit(number)
		} else if isLetter() {
			next()
			for isLetter() || isDigit() {
				next()
			}
			emit(identifier)
		} else {
			next()
			emit(illegal)
		}
	}

	return tokens
}

type token struct {
	text string
	kind tokenKind
}

type tokenKind = rune

const (
	number     tokenKind = 'n'
	identifier tokenKind = 'I'
	illegal    tokenKind = 'i'
)
