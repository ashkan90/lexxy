package tokens

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACKET = "["
	RBRACKET = "]"
	LBRACE   = "{"
	RBRACE   = "}"

	ROOT  = "ROOT"
	FIELD = "FIELD"

	COMMA = ","
	LOL   = ":"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
