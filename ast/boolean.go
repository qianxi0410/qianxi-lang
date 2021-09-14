package ast

import "github.com/monkey-lang/token"

type Boolean struct {
	// token.TRUE or token.FALSE
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }
