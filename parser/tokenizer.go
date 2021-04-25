package parser

import (
	"bufio"
	"os"
	"text/scanner"
	"unicode"
)

func tokenizeString(s *scanner.Scanner) (TokenType, string) {
	str := "'"
	for {
		if string(s.Peek()) == "'" {
			str += string(s.Next())
			if string(s.Peek()) == "'" { // '' escapes the ' character
				str += string(s.Next())
			} else {
				break
			}
		} else {
			if s.Peek() == scanner.EOF {
				return UNDEFINED, str
			}
			str += string(s.Next())
		}
	}
	return STRING, str
}

func tokenizeColon(s *scanner.Scanner) (TokenType, string) {
	if string(s.Peek()) == "-" {
		s.Next()
		return COLON_DASH, ":-"
	} else {
		return COLON, ":"
	}
}

func tokenizeIdentifier(ident string) (TokenType, string) {
	// We borrow functionality from scanner.Scanner to scan identifiers,
	// so this function just determines which identifier has been scanned.
	chars := []rune(ident)
	if len(chars) == 0 || (len(chars) > 0 && !unicode.IsLetter(chars[0])) {
		return UNDEFINED, ident
	} else if ident == "Schemes" {
		return SCHEMES, ident
	} else if ident == "Facts" {
		return FACTS, ident
	} else if ident == "Rules" {
		return RULES, ident
	} else if ident == "Queries" {
		return QUERIES, ident
	} else {
		return ID, ident
	}
}

func getToken(tokenType TokenType, val string, pos scanner.Position) Token {
	return Token{tokenType, val, pos.Line, pos.Column}
}

type DatalogTokenizer struct {
	scanner scanner.Scanner
	buffer  [2]Token
}

func (dt *DatalogTokenizer) Init(file *os.File) {
	dt.scanner.Init(bufio.NewReader(file))
	dt.scanner.Mode = scanner.ScanIdents | scanner.SkipComments
}

func (dt *DatalogTokenizer) updateBuffer(t Token) Token {
	dt.buffer[0] = dt.buffer[1]
	dt.buffer[1] = t
	return t
}

func (dt DatalogTokenizer) Prev() Token {
	return dt.buffer[0]
}

func (dt DatalogTokenizer) Current() Token {
	return dt.buffer[1]
}

func (dt *DatalogTokenizer) Next() Token {
	if dt.scanner.Peek() != scanner.EOF {
		pos := dt.scanner.Pos()
		dt.scanner.Scan()
		text := dt.scanner.TokenText()
		switch text {
		case "'":
			tokenType, value := tokenizeString(&dt.scanner)
			return dt.updateBuffer(getToken(tokenType, value, pos))
		case ":":
			tokenType, value := tokenizeColon(&dt.scanner)
			return dt.updateBuffer(getToken(tokenType, value, pos))
		case ",":
			return dt.updateBuffer(getToken(COMMA, text, pos))
		case ".":
			return dt.updateBuffer(getToken(PERIOD, text, pos))
		case "?":
			return dt.updateBuffer(getToken(Q_MARK, text, pos))
		case "(":
			return dt.updateBuffer(getToken(LEFT_PAREN, text, pos))
		case ")":
			return dt.updateBuffer(getToken(RIGHT_PAREN, text, pos))
		case "*":
			return dt.updateBuffer(getToken(MULTIPLY, text, pos))
		case "+":
			return dt.updateBuffer(getToken(ADD, text, pos))
		default:
			if text != "" {
				tokenType, value := tokenizeIdentifier(text)
				return dt.updateBuffer(getToken(tokenType, value, pos))
			}
		}
	}
	return dt.updateBuffer(getToken(EOFT, "", dt.scanner.Pos()))
}
