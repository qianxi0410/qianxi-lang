package lexer

import (
	"fmt"
	"testing"

	"github.com/monkey-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `
	!
	let five = 5;
	let ten = 10;
	let add = fn(x, y) { x + y;
	};
	if (5 < 10) { return true;
	} else { return false;
	}
	10 == 10; 10 != 9;
	"foobar"; "foo bar";
	[1, 2];
	>=
	<=
		`

	l := New(input)
	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			return
		}
		fmt.Println(tok)
	}
}
