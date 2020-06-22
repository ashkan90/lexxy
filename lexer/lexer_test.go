package lexer

import (
	"log"
	"new_lexxy/tokens"
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `[Service: (ID,Name, Details:(ID, Dte, UUID))]`
	test := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	} {
		{tokens.LBRACKET, "["},
		{tokens.FIELD, "Service"},
		{tokens.LOL, ":"},
		{tokens.LPAREN, "("},
		{tokens.FIELD, "ID"},
		{tokens.COMMA, ","},
		{tokens.FIELD, "Name"},
		{tokens.COMMA, ","},

		{tokens.FIELD, "Details"},
		{tokens.LOL, ":"},
		{tokens.LPAREN, "("},
		{tokens.FIELD, "ID"},
		{tokens.COMMA, ","},
		{tokens.FIELD, "Dte"},
		{tokens.COMMA, ","},
		{tokens.FIELD, "UUID"},
		{tokens.RPAREN, ")"},
		{tokens.RPAREN, ")"},

		{tokens.RBRACKET, "]"},
	}




	l := New(input)

	for i, tt := range test {
		tok := l.NextToken()

		log.Println(tok)

		if tok.Type != tt.expectedType {
			log.Println(tok)
			log.Fatalf("test[%d] - tokentype wrong, expected = %q, got = %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
