package parser

import "github.com/qianxi-lang/token"

// some token type's priority
const (
	_ int = iota
	LOWEST
	OROR        // ||
	ANDAND      // &&
	EQUALS      // ==
	LESSGREATER // < | >
	SUM         // +
	PRODUCT     // *
	REMINDER    // %
	PREFIX      // - | !
	CALL        // call fn
	INDEX       // array[index]
)

var precedences = map[token.TokenType]int{
	token.OROR:       OROR,
	token.ANDAND:     ANDAND,
	token.EQ:         EQUALS,
	token.NOT_EQ:     EQUALS,
	token.LT:         LESSGREATER,
	token.GT:         LESSGREATER,
	token.LE:         LESSGREATER,
	token.GE:         LESSGREATER,
	token.PLUS:       SUM,
	token.MINUS:      SUM,
	token.OR:         SUM,
	token.XOR:        SUM,
	token.AND:        PRODUCT,
	token.SHITFLEFT:  PRODUCT,
	token.SHITFRIGHT: PRODUCT,
	token.SLASH:      PRODUCT,
	token.ASTERISK:   PRODUCT,
	token.REMAINDER:  REMINDER,
	token.NOT:        PREFIX,
	token.LPAREN:     CALL,
	token.LBRACKET:   INDEX,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}
