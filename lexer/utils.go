package lexer

import (
	"unicode"

	"github.com/qianxi-lang/token"
)

func isLetter(ch rune) bool {
	return 'a' <= ch && 'z' >= ch || 'A' <= ch && 'Z' >= ch || ch == '_' ||
		unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
