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
	env := object.NewEnvironment()

	return Eval(program, env)
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
	t.Parallel()
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"consider (fact) { 10 }", 10},
		{"consider (cap) { 10 }", nil},
		{"consider (5 < 10) { 10 }", 10},
		{"consider (5 > 10) { 10 }", nil},
		{"consider (5 > 10) { 10 } however { 20 }", 20},
		{"consider (5 < 10) { 10 } however { 20 }", 10},
		{"consider (fact) { 10 } however { 20 }", 10},
		{"consider (cap) { 10 } however { 20 }", 20},
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

func TestErrorHandling(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected string
	}{
		{"5 + fact;", "type mismatch: INTEGER + BOOLEAN"},
		{"5 + fact; 5;", "type mismatch: INTEGER + BOOLEAN"},
		{"-fact", "unknown operator: -BOOLEAN"},
		{"fact + cap;", "unknown operator: BOOLEAN + BOOLEAN"},
		{"5; fact + cap; 5", "unknown operator: BOOLEAN + BOOLEAN"},
		{"consider (10 > 1) { fact + cap; }", "unknown operator: BOOLEAN + BOOLEAN"},
		{
			`
consider (10 > 1) {
	consider (10 > 1) {
		giving fact + cap;
	}

giving 1;
}
`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{"foo", "identifier not found: foo"},
		{"\"hello\" - \"world\"", "unknown operator: STRING - STRING"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)

			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("no error object returned. got=%T(%+v)",
					evaluated, evaluated)
				return
			}

			if errObj.Message != tt.expected {
				t.Errorf("wrong error message. expected=%q, got=%q",
					tt.expected, errObj.Message)
			}
		})
	}
}

func TestLetStatements(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected int64
	}{
		{"ask a = 5; giving a;", 5},
		{"ask a = 5 * 5; giving a;", 25},
		{"ask a = 5; ask b = a; giving b;", 5},
		{"ask a = 5; ask b = a; ask c = a + b + 5; giving c;", 15},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)
			testIntegerObject(t, evaluated, tt.expected)
		})
	}
}

func TestFunctionObject(t *testing.T) {
	t.Parallel()
	input := `funk(x) { x + 2; };`

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected int64
	}{
		{"ask identity = funk(x) { x; }; giving identity(5);", 5},
		{"ask identity = funk(x) { giving x; }; identity(5);", 5},
		{"ask double = funk(x) { x * 2; }; double(5);", 10},
		{"ask add = funk(x, y) { x + y; }; add(5, 5);", 10},
		{"ask add = funk(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"funk(x) { x; }(5)", 5},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)
			testIntegerObject(t, evaluated, tt.expected)
		})
	}
}

func TestClosures(t *testing.T) {
	t.Parallel()
	input := `
	ask newAdder = funk(x) {
		funk(y) { x + y; };
	};

	ask addTwo = newAdder(2);
	giving addTwo(3);
	`

	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 5)
}

func TestStringLiteral(t *testing.T) {
	t.Parallel()
	input := `"hello world!";`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "hello world!" {
		t.Errorf("str.Value is not %q. got=%q", "hello world!", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	t.Parallel()
	input := `"ello" + " " + "govna!";`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "ello govna!" {
		t.Errorf("str.Value is not %q. got=%q", "ello govna!", str.Value)
	}
}

func TestBuiltinFunctions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`thickness("")`, 0},
		{`thickness("govna")`, 5},
		{`thickness("hello world!")`, 12},
		{`thickness(1)`, "argument to `thickness` not supported, got INTEGER"},
		{`thickness("one", "two")`, "wrong number of arguments. got=2, want=1"},
		{`thickness([1, 2, 3])`, 3},
		{`thickness([])`, 0},
		{`first([1, 2, 3])`, 1},
		{`first([])`, nil},
		{`first(1)`, "argument to `first` must be ARRAY, got INTEGER"},
		{`last([1, 2, 3])`, 3},
		{`last([])`, nil},
		{`last(1)`, "argument to `last` must be ARRAY, got INTEGER"},
		{`bum([1, 2, 3])`, []int{2, 3}},
		{`bum([])`, nil},
		{`push([], 1)`, []int{1}},
		{`push(1, 1)`, "argument to `push` must be ARRAY, got INTEGER"},
		{`log("hello", "world!")`, nil},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)

			switch expected := tt.expected.(type) {
			case int:
				testIntegerObject(t, evaluated, int64(expected))
			case nil:
				testNullObject(t, evaluated)
			case string:
				errObj, ok := evaluated.(*object.Error)
				if !ok {
					t.Errorf("object is not Error. got=%T (%+v)",
						evaluated, evaluated)
					return
				}
				if errObj.Message != expected {
					t.Errorf("wrong error message. expected=%q, got=%q",
						expected, errObj.Message)
				}
			case []int:
				array, ok := evaluated.(*object.Array)
				if !ok {
					t.Errorf("obj not Array. got=%T (%+v)", evaluated, evaluated)
					return
				}

				if len(array.Elements) != len(expected) {
					t.Errorf("wrong num of elements. want=%d, got=%d",
						len(expected), len(array.Elements))
					return
				}

				for i, expectedElem := range expected {
					testIntegerObject(t, array.Elements[i], int64(expectedElem))
				}
			}
		})
	}
}

func TestArrayLiterals(t *testing.T) {
	t.Parallel()
	input := "[1, 2 * 2, 3 + 3];"
	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}
	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d", len(result.Elements))
	}
	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"[1, 2, 3][0]", 1},
		{"[1, 2, 3][1]", 2},
		{"[1, 2, 3][2]", 3},
		{"ask myArray = [1, 2, 3]; myArray[0];", 1},
		{"ask myArray = [1, 2, 3]; myArray[1];", 2},
		{"ask myArray = [1, 2, 3]; myArray[2];", 3},
		{"ask myArray = [1, 2, 3]; ask i = myArray[0]; myArray[i];", 2},
		{"[1, 2, 3][3]", nil},
		{"[1, 2, 3][-1]", nil},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)

			switch expected := tt.expected.(type) {
			case int:
				testIntegerObject(t, evaluated, int64(expected))
			default:
				testNullObject(t, evaluated)
			}
		})
	}
}

func TestHashLiterals(t *testing.T) {
	t.Parallel()
	input := `ask two = "two"; {
	"one": 10 - 9,
	two: 1 + 1,
	"thr" + "ee": 6 / 2,
	4: 4,
	5: 5,
	fact: 5,
	cap: 6
}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		(&object.Integer{Value: 5}).HashKey():      5,
		TRUE.HashKey():                             5,
		FALSE.HashKey():                            6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}
		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestHashIndexExpressions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`{"foo": 5}["foo"]`, 5},
		{`{"foo": 5}["bar"]`, nil},
		{`ask key = "foo"; {"foo": 5}[key]`, 5},
		{`{}["foo"]`, nil},
		{`{5: 5}[5]`, 5},
		{`{fact: 5}[fact]`, 5},
		{`{cap: 5}[cap]`, 5},
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
