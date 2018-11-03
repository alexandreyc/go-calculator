package ast

import (
	"fmt"

	"github.com/alexandreyc/go-calculator/token"
)

// Node ...
type Node interface {
	TokenLiteral() string
	String() string
}

// Operator ...
type Operator interface {
	Node
	operatorNode()
}

// Operand ...
type Operand interface {
	Node
	operandNode()
}

// IntegerLiteral ...
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// TokenLiteral ...
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) operandNode() {}

// PrefixOperator ...
type PrefixOperator struct {
	Token    token.Token
	Operator string
	Right    Node
}

// TokenLiteral ...
func (po *PrefixOperator) TokenLiteral() string {
	return po.Token.Literal
}

func (po *PrefixOperator) String() string {
	return fmt.Sprintf("(%s %s)", po.Operator, po.Right.String())
}

func (po *PrefixOperator) operatorNode() {}

// InfixOperator ...
type InfixOperator struct {
	Token    token.Token
	Left     Node
	Operator string
	Right    Node
}

// TokenLiteral ...
func (io *InfixOperator) TokenLiteral() string {
	return io.Token.Literal
}

func (io *InfixOperator) String() string {
	return fmt.Sprintf("(%s %s %s)", io.Left.String(), io.Operator, io.Right.String())
}

func (io *InfixOperator) operatorNode() {}
