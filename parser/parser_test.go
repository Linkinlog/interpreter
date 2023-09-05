package parser

import (
	"github.com/Linkinlog/MagLang/ast"
	"github.com/Linkinlog/MagLang/lexer"
	"testing"
)

func TestAskStatements(t *testing.T) {
	input := `
ask x = 5;
ask y = 10;
ask foo = 420;
`
	lex := lexer.New(input)
	parse := New(lex)

	program := parse.ParseProgram()
	checkParserErrors(t, parse)
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Wanted 3 statements, got %d",
			len(program.Statements),
		)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testAskStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testAskStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "ask" {
		t.Errorf("stmt.TokenLiteral not 'ask', got: %q", stmt.TokenLiteral())
		return false
	}

	askStmt, ok := stmt.(*ast.AskStatement)
	if !ok {
		t.Errorf("stmt not *ast.AskStatement, got: %T", stmt)
		return false
	}

	if askStmt.Name.Value != name {
		t.Errorf("askStmt.Name.Value not '%s', got: %s", name, askStmt.Name.Value)
		return false
	}
	if askStmt.Name.TokenLiteral() != name {
		t.Errorf("askStmt.Name.TokenLiteral() not '%s'. got: %s",
			name, askStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
giving 5;
giving 10;
giving 993322;
`
	lex := lexer.New(input)
	parse := New(lex)

	program := parse.ParseProgram()
	checkParserErrors(t, parse)
	if len(program.Statements) != 3 {
		t.Fatalf("Wanted 3 statements, got %d",
			len(program.Statements),
		)
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "giving" {
			t.Errorf("returnStmt.TokenLiteral not 'giving', got=%q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program doesnt have enough statements, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("exp not *ast.Identifier, got %T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s, got %s", "foobar", ident.TokenLiteral())
	}
}

func testReturnStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "giving" {
		t.Errorf("stmt.TokenLiteral not 'giving', got: %q", stmt.TokenLiteral())
		return false
	}

	_, ok := stmt.(*ast.ReturnStatement)
	if !ok {
		t.Errorf("stmt not *ast.ReturnStatement, got: %T", stmt)
		return false
	}

	return true
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has incorrent amount of statements. got %d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] not ast.ExpressionStatement, got %T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral, got %T",
			stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d, got %d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s, got %s", "5",
			literal.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errLen := 0
	errors := p.Errors()
	if errLen = len(errors); errLen == 0 {
		return
	}

	t.Errorf("parser has %d errors", errLen)
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
