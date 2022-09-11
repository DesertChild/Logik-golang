package token

//types
type TokenType string

type  Token struct{
	Type TokenType
	Literal string
}

const(
	STRING = "STRING"
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	//IDENTIFIERS AND LITERALS
	INT = "INT"
	FLOAT = "FLOAT"
	IDENT = "IDENT"

	//OPERATORS
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT = "<"
	GT = ">"
	EQ     = "=="
	NOT_EQ = "!="

	//Delimiters
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"let" : LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookUpId(ident string) TokenType{
	if tok,ok := keywords[ident];ok {
		return tok
	}
	return IDENT
}

//adding a function to be able to tell user-defined identifiers from the keywords
func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}