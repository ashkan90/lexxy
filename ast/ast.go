package ast

import (
	"bytes"
	"log"
	"new_lexxy/tokens"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
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
	Statements      []Statement
}

type (
	Builder interface {
		AddRow(key string, value interface{}) Builder
		GetRowValue(key string) *fieldImplementation
	}

	fieldImplementation struct {
		content interface{}
	}

	mapImplementation struct {
		definition map[string]*fieldImplementation
	}
)

func newMapImplementationV(key string, value interface{}) mapImplementation {
	return mapImplementation{definition: map[string]*fieldImplementation{
		key: {content: value},
	}}
}

func newMapImplementation() mapImplementation {
	return mapImplementation{definition: map[string]*fieldImplementation{}}
}

func NewMapper() Builder {
	return &mapImplementation{definition: map[string]*fieldImplementation{}}
}

func (m *mapImplementation) AddRow(key string, value interface{}) Builder {
	m.definition[key] = &fieldImplementation{
		content: value,
	}
	return m
}

func (m *mapImplementation) GetRowValue(key string) *fieldImplementation {
	return m.definition[key]
}

// console test text: [Service:(ID)]
var mapper Builder

func (p *Program) Serialize() {
	mapper = NewMapper()
	for _, statement := range p.Statements {
		p.serialize(statement)
	}

	log.Println(mapper)
}

func (p *Program) serialize(stmt Statement) *fieldImplementation {

	switch statement := stmt.(type) {
	case *StructToken:



		mapper.AddRow(statement.Itself.TokenLiteral(), newMapImplementation())

		if statement != nil && len(statement.Children) > 0 {
			p.evalFields(mapper.GetRowValue(statement.Itself.String()), statement.Children)
		}

		return mapper.GetRowValue(statement.Itself.String())
		//if statement.Children != nil && len(statement.Children) > 0 {
		//	if p.latestStatement != nil && p.latestStatement.String() == "" {
		//		mapper.AddRow(statement.Itself.TokenLiteral(), newMapImplementation())
		//		p.evalFields(mapper.GetRowValue(statement.Itself.String()), statement.Children)
		//	} else {
		//		mp := mapper.GetRowValue(p.latestStatement.String()).content.(mapImplementation)
		//		mp.AddRow(statement.Itself.TokenLiteral(), newMapImplementation())
		//		//mapper.AddRow(statement.Itself.TokenLiteral(), newMapImplementation())
		//		p.evalFields(mp.GetRowValue(statement.Itself.TokenLiteral()), statement.Children)
		//	}
		//
		//} else {
		//	mapper.AddRow(statement.Itself.TokenLiteral(), nil)
		//}

	}
}

// test string: [Service:(ID,Name,Details:(Id,ServiceName,ServiceAddress))]
func (p *Program) evalFields(fieldImp *fieldImplementation, fields []interface{}) {

	for _, field := range fields {
		switch field := field.(type) {
		case *StructToken:
			fieldImp.content = newMapImplementationV(field.TokenLiteral(), nil)
			p.serialize(field)

			//mapper.
			//	AddRow("ServiceID", 9).
			//	AddRow("Details", mapImplementation{definition: map[string]*fieldImplementation{
			//		"DetailID": { content: 9 },
			//		"Dte": { content: 7 },
			//		"UUID": { content: "129487912-345689734-123789SXMÖÇ-123789" },
			//		"MoreDetails": { content: mapImplementation{definition: map[string]*fieldImplementation{
			//			"MoreDetailID": { content: 18 },
			//			"Bla": { content: "bla" },
			//			"And": { content: "and" },
			//		}}},
			//	}})
		case *FieldToken:
			mapImp := fieldImp.content.(mapImplementation)
			mapImp.AddRow(field.TokenLiteral(), nil)
		}
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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

type StructToken struct {
	Itself   *FieldToken
	Children []interface{}
}

func (st *StructToken) statementNode()       {}
func (st *StructToken) TokenLiteral() string { return st.Itself.Token.Literal }
func (st *StructToken) String() string {
	var out bytes.Buffer

	out.WriteString(st.Itself.Name.TokenLiteral() + ":")
	out.WriteString("(")
	inf(st.Children, &out)

	return out.String()
}

type FieldToken struct {
	Token tokens.Token
	Name  *Identifier
	Value Expression
}

func (ft *FieldToken) statementNode()       {}
func (ft *FieldToken) TokenLiteral() string { return ft.Token.Literal }
func (ft *FieldToken) String() string {
	var out bytes.Buffer
	out.WriteString(ft.TokenLiteral())

	return out.String()
}

type Identifier struct {
	Token tokens.Token // tokens.FIELD
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

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
