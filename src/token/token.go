package token

type TokenType string

type Token struct {
	Token   TokenType
	Literal string
}
