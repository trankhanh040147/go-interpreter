package token

type TokType string

const (
	IDENT = "IDENT"
	EOF   = "EOF"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIV      = "/"

	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	ILLEGAL = "ILLEGAL"
)

var keywords = map[string]TokType{
	"let": LET,
	"fn":  FUNCTION,
}

type Token struct {
	Type    TokType
	Literal string
}

func LookUpIdentType(ident string) TokType {
	if val, ok := keywords[ident]; ok {
		return val
	}
	return IDENT
}
