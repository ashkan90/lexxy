package parser

import (
	"fmt"
	"new_lexxy/ast"
	"new_lexxy/lexer"
	"new_lexxy/tokens"
)

type Parser struct {
	l *lexer.Lexer

	errors       []string

	curToken  tokens.Token
	peekToken tokens.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) currentError(t tokens.TokenType) {
	msg := fmt.Sprintf("expected next token to be '%s', got '%s' instead",
		t, p.curToken.Type)
	p.errors = append(p.errors, msg)

}

func (p *Parser) peekError(t tokens.TokenType) {
	msg := fmt.Sprintf("expected next token to be '%s', got '%s' instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(tokens.EOF) {
		_struct := p.parseStruct()
		if _struct != nil {
			program.Statements = append(program.Statements, _struct)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStruct() ast.Statement {
	switch p.curToken.Type {
	case tokens.ROOT:
		return p.tStatement(nil)
	default:
		return nil
	}
}

func (p *Parser) tStatement(selfRoot *ast.StructToken) *ast.StructToken {
	if selfRoot == nil {
		selfRoot = &ast.StructToken{ Itself: &ast.FieldToken{ Token: p.curToken }}

		selfRoot.Itself.Name = &ast.Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
	}

	for !p.curTokenIs(tokens.LPAREN) {
		p.nextToken()
	}

	p.nextToken()

	for !p.curTokenIs(tokens.RPAREN) {
		// Handle deadlock
		if p.curToken.Type == tokens.EOF {
			p.currentError(tokens.RPAREN)
			break
		}

		switch p.curToken.Type {

		// catch errors at 'some' tokens.
		case tokens.COMMA:
			if p.peekTokenIs(tokens.RPAREN) {
				p.peekError(p.curToken.Type)
				return nil
			}

			//log.Println("lol: ", p.curToken)

		case tokens.FIELD:
			selfRoot.Children = append(selfRoot.Children, &ast.FieldToken{
				Token: p.curToken,
				Name: &ast.Identifier{
					Token: p.curToken,
					Value: p.curToken.Literal,
				},
			})

		case tokens.ROOT:
			newRootAsChildren := &ast.StructToken{
				Itself: &ast.FieldToken{
					Token: p.curToken,
					Name: &ast.Identifier{
						Token: p.curToken,
						Value: p.curToken.Literal,
					},
				},
			}

			// TODO: error wil be handled
			selfRoot.Children = append(selfRoot.Children, p.tStatement(newRootAsChildren))

		}
		p.nextToken()
	}

	return selfRoot
}

func (p *Parser) expectPeek(t tokens.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) curTokenIs(tokType tokens.TokenType) bool {
	return p.curToken.Type == tokType
}

func (p *Parser) peekTokenIs(tokType tokens.TokenType) bool {
	return p.peekToken.Type == tokType
}
