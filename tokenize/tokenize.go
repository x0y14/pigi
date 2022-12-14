package tokenize

import (
	"log"
	"strconv"
	"unicode"
)

var userInput []rune

// Pos
var lno int
var lat int
var wat int

var singleOpSymbols []string
var compositeOpSymbols []string

func init() {
	singleOpSymbols = []string{
		"(", ")", "[", "]", "{", "}",
		".", ",", ":", ";", "@",
		"+", "-", "*", "/", "%",
		">", "<",
		"=", "!",
	}
	compositeOpSymbols = []string{
		"**",
		"==", "!=", ">=", "<=",
		"+=", "-=", "*=", "/=", "%=", "**=",
		"&&", "||",
		"...", "..", // ...を..より優先するため先に記述する
	}
}

func startWith(q string) bool {
	qRunes := []rune(q)
	for i := 0; i < len(qRunes); i++ {
		if len(userInput) <= wat+i || userInput[wat+i] != qRunes[i] {
			return false
		}
	}
	return true
}

func isIdentRune(r rune) bool {
	return ('a' <= r && r <= 'z') ||
		('A' <= r && r <= 'Z') ||
		('0' <= r && r <= '9') ||
		('_' == r)
}

func isNotEof() bool {
	return wat < len(userInput)
}

func consumeComment() string {
	var s string
	for isNotEof() {
		if userInput[wat] == '\n' {
			break
		}
		s += string(userInput[wat])
		lat++
		wat++
	}
	return s
}

func consumeIdent() string {
	var s string
	for isNotEof() {
		if !isIdentRune(userInput[wat]) {
			break
		}
		s += string(userInput[wat])
		lat++
		wat++
	}
	return s
}

func consumeString() (string, bool) {
	var s string
	isDq := true
	if userInput[wat] == '\'' {
		isDq = false
	}
	// "/'
	lat++
	wat++

	for isNotEof() {
		if isDq && userInput[wat] == '"' {
			break
		} else if !isDq && userInput[wat] == '\'' {
			break
		}

		if userInput[wat] == '\\' &&
			(userInput[wat+1] == '\'' || userInput[wat+1] == '"') {
			if isDq {
				s += "\\\""
			} else {
				s += "\\'"
			}
			lat += 2
			wat += 2
			continue
		}

		s += string(userInput[wat])
		lat++
		wat++
	}
	// "/'
	lat++
	wat++

	return s, isDq
}

func consumeNumber() (string, bool) {
	isFloat := false
	var s string
	for isNotEof() {
		if unicode.IsDigit(userInput[wat]) {
			s += string(userInput[wat])
			lat++
			wat++
			continue
		} else if userInput[wat] == '.' {
			if len(userInput) <= wat+1 || !unicode.IsDigit(userInput[wat+1]) {
				break
			}
			s += string(userInput[wat])
			lat++
			wat++
			isFloat = true
			continue
		} else {
			break
		}
	}

	return s, isFloat
}

func consumeWhite() string {
	var s string
	for isNotEof() {
		if userInput[wat] == ' ' || userInput[wat] == '\t' {
			s += string(userInput[wat])
			lat++
			wat++
		} else {
			break
		}
	}
	return s
}

func Tokenize(input string) (*Token, error) {
	// init
	userInput = []rune(input)
	lno = 1
	lat = 0
	wat = 0
	var head Token
	cur := &head

inputLoop:
	for isNotEof() {
		// white
		if userInput[wat] == ' ' || userInput[wat] == '\t' {
			pos := NewPosition(lno, lat, wat)
			s := consumeWhite()
			cur = NewWhite(cur, s, pos)
			continue
		}
		// newline
		if userInput[wat] == '\n' {
			cur = NewNL(cur, "\n", NewPosition(lno, lat, wat))
			lno++
			lat = 0
			wat++
			continue
		}
		// comment
		if userInput[wat] == '#' {
			pos := NewPosition(lno, lat, wat)
			comment := consumeComment()
			cur = NewComment(cur, comment, pos)
			continue
		}
		// op symbols
		for _, r := range append(compositeOpSymbols, singleOpSymbols...) {
			if startWith(r) {
				cur = NewOpSymbol(cur, r, NewPosition(lno, lat, wat))
				lat += len(r)
				wat += len(r)
				continue inputLoop
			}
		}
		// ident
		if isIdentRune(userInput[wat]) && !unicode.IsDigit(userInput[wat]) {
			pos := NewPosition(lno, lat, wat)
			id := consumeIdent()
			cur = NewIdent(cur, id, pos)
			continue
		}

		// string
		if userInput[wat] == '\'' || userInput[wat] == '"' {
			pos := NewPosition(lno, lat, wat)
			str, isDq := consumeString()
			cur = NewString(cur, str, pos, isDq)
			continue
		}

		// number
		if unicode.IsDigit(userInput[wat]) {
			pos := NewPosition(lno, lat, wat)
			numStr, isFloat := consumeNumber()
			if isFloat {
				n, err := strconv.ParseFloat(numStr, 64)
				if err != nil {
					return nil, err
				}
				cur = NewFloat(cur, numStr, pos, n)
				continue
			} else {
				n, err := strconv.ParseInt(numStr, 10, 0)
				if err != nil {
					return nil, err
				}
				cur = NewInt(cur, numStr, pos, int(n))
				continue
			}
		}

		log.Fatalf("[%d:%d] unexpected character: %s", lno, lat, string(userInput[wat]))
	}

	cur = NewEof(cur, NewPosition(lno, lat, wat))
	return head.Next, nil
}
