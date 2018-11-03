package parser

import (
	"testing"

	"github.com/alexandreyc/go-calculator/lexer"
)

func TestParse(t *testing.T) {
	testExpression("42", "42", t)
	testExpression("-42", "(- 42)", t)
	testExpression("--42", "(- (- 42))", t)
	testExpression("42 + 42", "(42 + 42)", t)
	testExpression("42 * 42", "(42 * 42)", t)
	testExpression("42 / 42", "(42 / 42)", t)
	testExpression("42 + 42 * 42", "(42 + (42 * 42))", t)
	testExpression("42 + (42 * 42)", "(42 + (42 * 42))", t)
	testExpression("(42 + 42) * 42", "((42 + 42) * 42)", t)
}

func testExpression(exp string, expected string, t *testing.T) {
	parser := New(lexer.New(exp))
	parsed := parser.Parse().String()
	if parsed != expected {
		t.Fatalf("wrong parsing: expected=%q got=%q", expected, parsed)
	}
}
