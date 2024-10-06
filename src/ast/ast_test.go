package ast

import (
	"testing"

	"github.com/batt0s/rizzy/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&DefStatement{
				Token: token.Token{Type: token.DEF, Literal: "def"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "otherVar"},
					Value: "otherVar",
				},
			},
		},
	}

	if program.String() != "def myVar = otherVar;" {
		t.Errorf("program.String() wrong, got %q", program.String())
	}
}
