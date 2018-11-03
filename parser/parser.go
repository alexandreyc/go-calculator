package parser

import (
	"fmt"
	"strconv"

	"github.com/alexandreyc/go-calculator/ast"
	"github.com/alexandreyc/go-calculator/lexer"
	"github.com/alexandreyc/go-calculator/token"
)

const (
	_ int = iota
	LOWEST
	SUM     // +
	PRODUCT // *
	PREFIX  // -X or !X
)

var precedences = map[token.TokenType]int{
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

type (
	prefixParseFn func() ast.Node
	infixParseFn  func(ast.Node) ast.Node
)

// Parser ...
type Parser struct {
	lexer          *lexer.Lexer
	errors         []string
	curToken       token.Token
	peekToken      token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

// New ...
func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer:  lexer,
		errors: []string{},
	}

	parser.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	parser.prefixParseFns[token.INT] = parser.parseIntegerLiteral
	parser.prefixParseFns[token.MINUS] = parser.parseMinus
	parser.prefixParseFns[token.LPAREN] = parser.parseGrouped

	parser.infixParseFns = make(map[token.TokenType]infixParseFn)
	parser.infixParseFns[token.PLUS] = parser.parseInfixOperator
	parser.infixParseFns[token.ASTERISK] = parser.parseInfixOperator
	parser.infixParseFns[token.SLASH] = parser.parseInfixOperator

	parser.nextToken()
	parser.nextToken()

	return parser
}

// Errors ...
func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) expectPeek(t token.TokenType) bool {
	if parser.peekTokenIs(t) {
		parser.nextToken()
		return true
	}

	parser.peekError(t)
	return false
}

func (parser *Parser) peekTokenIs(t token.TokenType) bool {
	return parser.peekToken.Type == t
}

func (parser *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, parser.peekToken.Type)
	parser.errors = append(parser.errors, msg)
}

func (parser *Parser) curPrecedence() int {
	if p, ok := precedences[parser.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (parser *Parser) peekPrecedence() int {
	if p, ok := precedences[parser.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

// Parse ...
func (parser *Parser) Parse() ast.Node {
	return parser.parse(LOWEST)
}

func (parser *Parser) parse(precedence int) ast.Node {
	prefix := parser.prefixParseFns[parser.curToken.Type]
	if prefix == nil {
		msg := fmt.Sprintf("could not find a prefix parsing function for: %s", parser.curToken.Type)
		parser.errors = append(parser.errors, msg)
		return nil
	}

	left := prefix()

	for !parser.peekTokenIs(token.EOF) && precedence < parser.peekPrecedence() {
		infix := parser.infixParseFns[parser.peekToken.Type]
		if infix == nil {
			return left
		}

		parser.nextToken()
		left = infix(left)
	}

	return left
}

func (parser *Parser) parseIntegerLiteral() ast.Node {
	il := &ast.IntegerLiteral{Token: parser.curToken}

	value, err := strconv.ParseInt(parser.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", parser.curToken.Literal)
		parser.errors = append(parser.errors, msg)
		return nil
	}

	il.Value = value
	return il
}

func (parser *Parser) parseMinus() ast.Node {
	po := &ast.PrefixOperator{
		Token:    parser.curToken,
		Operator: parser.curToken.Literal,
	}

	parser.nextToken()

	po.Right = parser.parse(PREFIX)

	return po
}

func (parser *Parser) parseGrouped() ast.Node {
	parser.nextToken()

	exp := parser.parse(LOWEST)

	if !parser.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (parser *Parser) parseInfixOperator(left ast.Node) ast.Node {
	io := &ast.InfixOperator{
		Token:    parser.curToken,
		Operator: parser.curToken.Literal,
		Left:     left,
	}

	precedence := parser.curPrecedence()
	parser.nextToken()
	io.Right = parser.parse(precedence)

	return io
}
