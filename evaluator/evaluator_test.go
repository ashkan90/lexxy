package evaluator

import (
	"lexxy/lexer"
	"lexxy/object"
	"lexxy/parser"
	"log"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "Company:(Test,Testtttt,Service:(ID, D))", expected: "Company"},
		//{input: "Service:()", expected: "Service" },

	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		//if  evaluated := evaluated.(*object.Struct); len(evaluated.Fields) > 0{
		//	log.Println("evaluated: ", evaluated.Fields[0].(*object.Field))
		//}
		fields := evaluated.(*object.Struct).Fields
		log.Printf("ev1: %q\n", evaluated.(*object.Struct))
		log.Printf("ev: %q\n", fields[len(fields)-1].(*object.Struct))
		testStructObject(t, evaluated, tt.expected)
	}

}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	return Eval(program)
}

func testStructObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.Struct)

	if !ok {
		t.Errorf("Object is not Struct, got: %q", result)
		return false
	}

	if result.Itself != expected {
		t.Errorf("Struct's itself is not expected. got: %q (%+v)", result.Itself, expected)
		return false
	}

	return true
}
