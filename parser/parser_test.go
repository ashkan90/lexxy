package parser

import (
	"lexxy/ast"
	"lexxy/lexer"
	"log"
	"strings"
	"testing"
	"unicode"
)

func TestFieldToken_TokenLiteral(t *testing.T) {
	input := `
		[
			Company: (
				CompanyID,
				Address
			),
			Service: (
				ServiceID,
				Name,
				Details:(
					DetailsID,
					Dte,
					UUID,
					MoreDetails:(
						MoreDetailsID,
						MoreDetail,
						Bla,
						And
					)
				)
			)
		]
`

	// Statements toString() fonksiyonunda, verileri json tipinde stringleyip
	// unmarshall ile açacağız.

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	//program.TryToRun()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatal("ParseProgram() returned 'nil'")
	}

	log.Println("Q: ", Q.String())

	log.Printf("v: %#v\n", program.Statements[0].(*ast.StructToken))
	log.Println("json: ", program.Json())

	service := program.Statements[1].(*ast.StructToken)
	serviceDetails := service.Children[len(service.Children)-1].(*ast.StructToken)
	serviceDetailsMoreDetails := serviceDetails.Children[len(serviceDetails.Children)-1].(*ast.StructToken)

	log.Println("Statements: -->> ", program.Statements)
	log.Println("First Statement: -->>", program.Statements[0])
	log.Println("Second Statement:-->(Service): -->>", service)
	log.Println("Second Statement:-->(Service)-->(Details): -->>", serviceDetails)
	log.Println("Second Statement:-->(Service)-->(Details)-->(MoreDetails): -->>", serviceDetailsMoreDetails)
	//log.Println("str: ", program.String())
	//log.Println("str: ", spaceStringsBuilder(input))

	if program.String() != spaceStringsBuilder(input) {
		t.Errorf("program.String() or 'input' wrong. got '%s', expected: '%s'", spaceStringsBuilder(input), program.String())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has errors: %d\n", len(errors))
	for _, err := range errors {
		t.Errorf("parser error: %q\n", err)
	}

	t.FailNow()
}

func spaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
