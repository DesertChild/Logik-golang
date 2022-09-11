package ast

//imports

import (
	"logic/token"
	"testing"
)

//performing some simpel tests

func TestingString(t *testing.T){
	program := &Program{
		Statements : []Statement{
			&LetStatement{
				Token : token.Token{Type : token.LET, Literal: "let"},
				Name: &Identifier{
					Token : token.Token{Type: token.IDENT, Literal: "my"},
					Value : "my",
				},
				Value : &Identifier{
					Token : token.Token{Type: token.IDENT, Literal: "your"},
					Value : "your",
				},
			},
		},
	}
		if program.String() != "let my = your;" {
			t.Errorf("program.String() wrong. got=%q", program.String())
		}
}

