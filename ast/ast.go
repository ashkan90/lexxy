package ast

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"lexxy/tokens"
	"log"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
	Json() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type SearchParameter struct {
	ColumnNumber int
	ColumnName   string
}

func (p *Program) TryToRun() {

	//var ctx context.Context
	var db *sql.DB
	var err error

	db, err = sql.Open("mysql", "root:@/bnew")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("select id from companies")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var out int64
	for rows.Next() {
		err = rows.Scan(&out)
		if err != nil {
			panic(err)
		}

		fmt.Println(out)
	}

}

func (p *Program) String() string {
	var out bytes.Buffer

	out.WriteString("[")

	for _, s := range p.Statements {
		out.WriteString(s.String() + ",")
	}

	trim := strings.TrimRight(out.String(), ",")

	out.Reset()
	out.WriteString(trim)
	out.WriteString("]")

	return out.String()
}

func (p *Program) Json() string {
	var out bytes.Buffer

	out.WriteString("{")

	for _, s := range p.Statements {
		out.WriteString(s.Json() + ",")
	}

	trim := strings.TrimRight(out.String(), ",")

	out.Reset()
	out.WriteString(trim)
	out.WriteString("}")

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}
func (p *Program) SearchOn(args ...*SearchParameter) {
	log.Println(p.Statements)
	//for _, statement := range p.Statements {
	//	SearchParameter{
	//		ColumnNumber: 0,
	//		ColumnName:   "",
	//	}
	//}
}

type StructToken struct {
	Itself   *FieldToken
	Children []interface{}
}

func (st *StructToken) statementNode() {}
func (st *StructToken) TokenLiteralJSON() string {
	return fmt.Sprintf("\"%s\"", st.Itself.Token.Literal)
}
func (st *StructToken) TokenLiteral() string { return st.Itself.Token.Literal }
func (st *StructToken) String() string {
	var out bytes.Buffer

	out.WriteString(st.Itself.Name.TokenLiteral() + ":(")
	inf(st.Children, &out)

	return out.String()
}
func (st *StructToken) Json() string {
	var out bytes.Buffer

	out.WriteString(st.Itself.Name.TokenLiteralJSON() + ":")
	out.WriteString("{")
	inf(st.Children, &out)

	return out.String()
}

type FieldToken struct {
	Token tokens.Token
	Name  *Identifier
	Value Expression
}

func (ft *FieldToken) statementNode()           {}
func (ft *FieldToken) TokenLiteral() string     { return ft.Token.Literal }
func (ft *FieldToken) TokenLiteralJSON() string { return fmt.Sprintf("\"%s\"", ft.Token.Literal) }
func (ft *FieldToken) String() string {
	var out bytes.Buffer
	out.WriteString(ft.TokenLiteral())

	return out.String()
}

func (ft *FieldToken) Json() string {
	var out bytes.Buffer
	out.WriteString(ft.TokenLiteralJSON())

	return out.String()
}

type Identifier struct {
	Token tokens.Token // tokens.FIELD
	Value string
}

func (i *Identifier) expressionNode()          {}
func (i *Identifier) TokenLiteral() string     { return i.Token.Literal }
func (i *Identifier) TokenLiteralJSON() string { return fmt.Sprintf("\"%s\"", i.Token.Literal) }

func inf(children []interface{}, writer *bytes.Buffer) {
	for _, ch := range children {
		switch ch.(type) {
		case *StructToken:
			var st *StructToken
			st = ch.(*StructToken)

			(*writer).WriteString(st.Itself.Name.TokenLiteral() + ":(")
			if len(st.Children) > 0 {
				inf(st.Children, writer)
			}
		case *FieldToken:
			var ft *FieldToken
			ft = ch.(*FieldToken)

			(*writer).WriteString(ft.String() + ",")

		}
	}
	trim := strings.TrimRight((*writer).String(), ",")

	(*writer).Reset()
	(*writer).WriteString(trim)
	(*writer).WriteString(")")

}
