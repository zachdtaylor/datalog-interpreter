package main

import (
	"fmt"
)

type TokenType int

const (
	UNDEFINED TokenType = iota
	PERIOD
	Q_MARK
	LEFT_PAREN
	RIGHT_PAREN
	COLON
	COLON_DASH
	MULTIPLY
	ADD
	SCHEMES
	FACTS
	RULES
	QUERIES
	ID
	STRING
	COMMENT
	COMMA
	EOFT
)

func (t TokenType) String() string {
	switch t {
	case COMMA:
		return "COMMA"
	case PERIOD:
		return "PERIOD"
	case Q_MARK:
		return "Q_MARK"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case COLON:
		return "COLON"
	case COLON_DASH:
		return "COLON_DASH"
	case MULTIPLY:
		return "MULTIPLY"
	case ADD:
		return "ADD"
	case SCHEMES:
		return "SCHEMES"
	case FACTS:
		return "FACTS"
	case RULES:
		return "RULES"
	case QUERIES:
		return "QUERIES"
	case ID:
		return "ID"
	case STRING:
		return "STRING"
	case COMMENT:
		return "COMMENT"
	case UNDEFINED:
		return "UNDEFINED"
	case EOFT:
		return "EOFT"
	default:
		return "UNKNOWN"
	}
}

type Token struct {
	tokenType  TokenType
	value      string
	lineNumber int
	column     int
}

func (t Token) String() string {
	return fmt.Sprintf("(%s,\"%s\",%d,%d)", t.tokenType, t.value, t.lineNumber, t.column)
}
