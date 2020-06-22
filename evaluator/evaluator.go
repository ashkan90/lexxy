package evaluator

import (
	"new_lexxy/ast"
	"new_lexxy/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements)
	case *ast.StructToken:
		if node.Children != nil && len(node.Children) > 0 {

			_struct := &object.Struct{ Itself: node.Itself.TokenLiteral() }
			evalFields(&_struct.Fields, node.Children)

			return _struct
		}
		return &object.Struct{ Itself: node.Itself.TokenLiteral(), Fields: nil}
	}

	return nil
}

func evalProgram(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}

func evalFields(outFields *[]interface{}, fields []interface{})  {
	for _, field := range fields {
		switch field := field.(type) {
		case *ast.StructToken:
			eva := Eval(field).(*object.Struct)
			*outFields = append(*outFields, eva)
		case *ast.FieldToken:
			*outFields = append(*outFields, &object.Field{ Name: field.TokenLiteral() })
		}
	}
}