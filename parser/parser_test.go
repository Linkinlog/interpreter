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