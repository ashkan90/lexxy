package object

import "fmt"

type Type string

const (
	STRUCT_OBJ = "STRUCT"
	FIELD_OBJ  = "FIELD"
	NULL_OBJ   = "NULL"
)

type Object interface {
	Type() Type
	Inspect() string
}

type Null struct{}

type Struct struct {
	Itself string
	Fields []interface{} // hem struct, hem field olabilir.
}

type Field struct {
	Name string
}

func (s *Struct) Type() Type {
	return STRUCT_OBJ
}

func (s *Struct) Inspect() string {
	return fmt.Sprintf("%#q %#v", s.Itself, s.Fields)
}

func (s *Field) Type() Type {
	return FIELD_OBJ
}

func (s *Field) Inspect() string {
	return fmt.Sprintf("%q", s.Name)
}

func (s *Null) Type() Type {
	return NULL_OBJ
}

func (s *Null) Inspect() string {
	return "null"
}
