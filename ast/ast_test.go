package ast

import (
	"github.com/Linkinlog/MagLang/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&AskStatement{
				Token: token.Token{Type: token.LET, Literal: "ask"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "ask myVar = anotherVar;" {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}
