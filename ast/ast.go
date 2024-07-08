package ast

import "Eclab/token"

// every node in AST has to implement this interface,
// which means that it has to provide a method that
// returns the literal value of the associated token
type Node interface {
	TokenLiteral() string
}

// Interface for Statement Nodes
type Statement interface {
	Node
	statementNode()
}

// Interface for Expression Nodes
type Expression interface {
	Node
	expressionNode()
}

// root node of every AST our parser provides
type Program struct {
	Statements []Statement
}

// return <expression>
type ReturnStatement struct {
	Token       token.Token // the return token
	ReturnValue Expression  // expression returned by return statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (rs *ReturnStatement) statementNode() {}

// returns literal for a given token
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // variable name
	Value Expression  // right side expression
}

func (ls *LetStatement) statementNode() {

}

// returns literal for a given token
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {

}

// returns token literal from a given identifier
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
