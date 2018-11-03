package eval

import (
	"github.com/alexandreyc/go-calculator/ast"
	"github.com/alexandreyc/go-calculator/token"
)

// Result ...
type Result struct {
	Value     int64
	IsDefined bool
}

// Eval ...
func Eval(node ast.Node) Result {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return Result{Value: node.Value, IsDefined: true}
	case *ast.PrefixOperator:
		right := Eval(node.Right)
		switch node.Token.Type {
		case token.MINUS:
			return Result{Value: -right.Value, IsDefined: true}
		}
	case *ast.InfixOperator:
		left, right := Eval(node.Left), Eval(node.Right)
		switch node.Token.Type {
		case token.PLUS:
			return Result{Value: left.Value + right.Value, IsDefined: true}
		case token.MINUS:
			return Result{Value: left.Value - right.Value, IsDefined: true}
		case token.ASTERISK:
			return Result{Value: left.Value * right.Value, IsDefined: true}
		case token.SLASH:
			return Result{Value: left.Value / right.Value, IsDefined: true}
		}
	}
	return Result{IsDefined: false}
}
