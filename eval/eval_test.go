package eval

import (
	"testing"

	"github.com/alexandreyc/go-calculator/lexer"
	"github.com/alexandreyc/go-calculator/parser"
)

func TestEval(t *testing.T) {
	testEval("42", 42, t)
	testEval("-42", -42, t)
	testEval("-(42)", -42, t)
	testEval("1+2", 3, t)
	testEval("(1+2)", 3, t)
	testEval("2*2", 4, t)
	testEval("(2*2)", 4, t)
	testEval("10/5", 2, t)
	testEval("3+4*4", 19, t)
	testEval("(3+4)*4", 28, t)
	testEval("--5", 5, t)
}

func testEval(input string, expected int64, t *testing.T) {
	lexer := lexer.New(input)
	parser := parser.New(lexer)
	output := Eval(parser.Parse())
	if !output.IsDefined {
		t.Fatalf("wrong eval: value for %q is undefined", input)
	}
	if output.Value != expected {
		t.Fatalf("wrong eval: expected=%d got=%d", expected, output.Value)
	}
}
