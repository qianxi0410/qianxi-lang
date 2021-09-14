package ast

import "github.com/monkey-lang/token"

type Identifier struct {
	// token.IDENT
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string { return i.Value }
