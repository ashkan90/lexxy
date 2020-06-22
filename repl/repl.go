package repl

import (
	"bufio"
	"fmt"
	"io"
	"new_lexxy/evaluator"
	"new_lexxy/lexer"
	"new_lexxy/parser"
	"new_lexxy/tokens"
)

const PROPMPT = ">> "

type Company struct {
	CompanyID uint
	Address   string
}

type Service struct {
	ServiceID uint
	Name      string
	Details   Detail
}

type Detail struct {
	DetailID    uint
	Dte         string
	UUID        string
	MoreDetails MoreDetail
}

type MoreDetail struct {
	MoreDetailID uint
	Bla          string
	And          string
}

/*
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
*/

func Start(in io.Reader, out io.Writer) {

	//initialData := &Service{
	//	ServiceID: 14,
	//	Name:      "Custom Service",
	//	Details:   Detail{
	//		DetailID:    9,
	//		Dte:         "2020-04-04",
	//		UUID:        "87897894-454979877-156454",
	//		MoreDetails: MoreDetail{
	//			MoreDetailID: 78,
	//			Bla:          "bla",
	//			And:          "andand",
	//		},
	//	},
	//}


	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROPMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		p := parser.New(l)
		program := p.ParseProgram()
		program.Serialize()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

		for tok := l.NextToken(); tok.Type != tokens.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
