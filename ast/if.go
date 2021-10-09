package ast

import (
	"bytes"

	"github.com/monkey-lang/token"
)

type IfExpression struct {
	// token.IF
	Token       token.Token
	Condition   []Expression
	Consequence []*BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition[0].String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence[0].String())

	if len(ie.Condition) > 1 {
		for i := 1; i < len(ie.Condition); i++ {
			out.WriteString("elif")
			out.WriteString(ie.Condition[i].String())
			out.WriteString(" ")
			out.WriteString(ie.Consequence[i].String())
		}
	}

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}
