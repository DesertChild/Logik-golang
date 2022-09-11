package lexer

//imports
import (
	"logic/token"
	"testing"
)

//functions
func TestNextToken(t *testing.T){
	input := `let seven = 7;
	let twenty = 20;
	
	let add = function(x, y) {
	  x + y;
	};
	
	let result = add(seven, twenty);
	!-/*7;
	7 < 20 > 7;
	
	if (7 < 20) {
		return true;
	} else {
		return false;
	}
	
	20 == 20;
	20 != 19;
	`
	tests := []struct{ 
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "seven"},
		{token.ASSIGN, "="},
		{token.INT, "7"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "twenty"},
		{token.ASSIGN, "="},
		{token.INT, "20"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "function"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "seven"},
		{token.COMMA, ","},
		{token.IDENT, "twenty"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "7"},
		{token.SEMICOLON, ";"},
		{token.INT, "7"},
		{token.LT, "<"},
		{token.INT, "20"},
		{token.GT, ">"},
		{token.INT, "7"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "7"},
		{token.LT, "<"},
		{token.INT, "20"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "20"},
		{token.EQ, "=="},
		{token.INT, "20"},
		{token.SEMICOLON, ";"},
		{token.INT, "20"},
		{token.NOT_EQ, "!="},
		{token.INT, "19"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},


	}
	l := New(input)
	for i,tt:= range tests{
		tok:= l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			i, tt.expectedLiteral, tok.Literal)
		}
	}
}