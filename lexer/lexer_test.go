package lexer

import (
	"testing"

	"github.com/alexandreyc/go-calculator/token"
)

type test struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	// Test 1
	input := ""
	tests := []test{
		{token.EOF, ""},
	}
	testInput(t, input, tests)

	// Test 2
	input = "0"
	tests = []test{
		{token.INT, "0"},
	}
	testInput(t, input, tests)

	// Test 3
	input = "42"
	tests = []test{
		{token.INT, "42"},
	}
	testInput(t, input, tests)

	// Test 4
	input = " 42 \t"
	tests = []test{
		{token.INT, "42"},
	}
	testInput(t, input, tests)

	// Test 5
	input = "-42"
	tests = []test{
		{token.MINUS, "-"},
		{token.INT, "42"},
	}
	testInput(t, input, tests)

	// Test 6
	input = "2+40"
	tests = []test{
		{token.INT, "2"},
		{token.PLUS, "+"},
		{token.INT, "40"},
	}
	testInput(t, input, tests)

	// Test 7
	input = "(3*3)"
	tests = []test{
		{token.LPAREN, "("},
		{token.INT, "3"},
		{token.ASTERISK, "*"},
		{token.INT, "3"},
		{token.RPAREN, ")"},
	}
	testInput(t, input, tests)

	// Test 8
	input = "3 - (43/10)"
	tests = []test{
		{token.INT, "3"},
		{token.MINUS, "-"},
		{token.LPAREN, "("},
		{token.INT, "43"},
		{token.SLASH, "/"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
	}
	testInput(t, input, tests)

	// Test 9
	input = "3 & 3"
	tests = []test{
		{token.INT, "3"},
		{token.ILLEGAL, "&"},
		{token.INT, "3"},
	}
	testInput(t, input, tests)
}

func testInput(t *testing.T, input string, tests []test) {
	lexer := New(input)

	for _, tt := range tests {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("wrong token type: expected=%q actual=%q", tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("wrong token literal: expected=%q actual=%q", tt.expectedLiteral, tok.Literal)
		}
	}
}
