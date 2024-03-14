package evaluator

import (
	"testing"

	"github.com/Linkinlog/MagLang/lexer"
	"github.com/Linkinlog/MagLang/object"
	"github.com/Linkinlog/MagLang/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)
			testIntegerObject(t, evaluated, tt.expected)
		})
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}
	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected bool
	}{
		{"fact", true},
		{"cap", false},
		{"5 < 10", true},
		{"5 > 10", false},
		{"5 < 5", false},
		{"5 > 4", true},
		{"5 == 5", true},
		{"5 != 5", false},
		{"5 == 6", false},
		{"5 != 6", true},
		{"fact == fact", true},
		{"cap == cap", true},
		{"fact == cap", false},
		{"fact != cap", true},
		{"(5 < 10) == fact", true},
		{"(5 < 10) == cap", false},
		{"(5 > 10) == fact", false},
		{"(5 > 10) == cap", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)
			testBooleanObject(t, evaluated, tt.expected)
		})
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
		return false
	}
	return true
}

func TestBangOperator(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected bool
	}{
		{"fact", true},
		{"cap", false},
		{"!fact", false},
		{"!cap", true},
		{"!5", false},
		{"!!fact", true},
		{"!!cap", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)
			testBooleanObject(t, evaluated, tt.expected)
		})
	}
}

func TestIfElseExpressions(t *testing.T) {
	t.Skip()
	t.Parallel()
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (fact) { 10 }", 10},
		{"if (cap) { 10 }", nil},
		{"if (5 < 10) { 10 }", 10},
		{"if (5 > 10) { 10 }", nil},
		{"if (5 > 10) { 10 } else { 20 }", 20},
		{"if (5 < 10) { 10 } else { 20 }", 10},
		{"if (fact) { 10 } else { 20 }", 10},
		{"if (cap) { 10 } else { 20 }", 20},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)
			integer, ok := tt.expected.(int)
			if ok {
				testIntegerObject(t, evaluated, int64(integer))
			} else {
				testNullObject(t, evaluated)
			}
		})
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected int64
	}{
		{"giving 6;", 6},
		{"giving 10; 9;", 10},
		{"giving 2 * 5; 9;", 10},
		{"9; giving 2 * 5; 9;", 10},
		{
			`
consider (10 > 1) {
  consider (cap) {
    giving 10;
  }

  giving 1;
}
`,
			1,
		},
		{
			`
consider (10 > 1) {
  consider (10 > 1) {
    giving 10;
  }

  giving 1;
}
`,
			10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)
			testIntegerObject(t, evaluated, tt.expected)
		})
	}
}
