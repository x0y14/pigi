package tokenize

import "strings"

type Token struct {
	Kind TokenKind
	Raw  string
	Pos  *Position

	IsDoubleQuote bool
	NumFloat      float64
	NumInt        int

	Next *Token
}

func NewToken(kind TokenKind, raw string, pos *Position) *Token {
	return &Token{
		Kind: kind,
		Raw:  raw,
		Pos:  pos,
		Next: nil,
	}
}

func NewEof(cur *Token, pos *Position) *Token {
	tok := NewToken(TKEof, "", pos)
	cur.Next = tok
	return tok
}

func NewOpSymbol(cur *Token, raw string, pos *Position) *Token {
	var tok *Token
	switch raw {
	case "(":
		tok = NewToken(TKLrb, raw, pos)
	case ")":
		tok = NewToken(TKRrb, raw, pos)
	case "[":
		tok = NewToken(TKLsb, raw, pos)
	case "]":
		tok = NewToken(TKRsb, raw, pos)
	case "{":
		tok = NewToken(TKLcb, raw, pos)
	case "}":
		tok = NewToken(TKRcb, raw, pos)
	case ".":
		tok = NewToken(TKDot, raw, pos)
	case ",":
		tok = NewToken(TKComma, raw, pos)
	case ":":
		tok = NewToken(TKColon, raw, pos)
	case ";":
		tok = NewToken(TKSemi, raw, pos)
	case "@":
		tok = NewToken(TKAt, raw, pos)

	case "+":
		tok = NewToken(TKAdd, raw, pos)
	case "-":
		tok = NewToken(TKSub, raw, pos)
	case "*":
		tok = NewToken(TKMul, raw, pos)
	case "/":
		tok = NewToken(TKDiv, raw, pos)
	case "%":
		tok = NewToken(TKMod, raw, pos)
	case "**":
		tok = NewToken(TKExp, raw, pos)

	case "==":
		tok = NewToken(TKEq, raw, pos)
	case "!=":
		tok = NewToken(TKNe, raw, pos)
	case ">":
		tok = NewToken(TKGt, raw, pos)
	case "<":
		tok = NewToken(TKLt, raw, pos)
	case ">=":
		tok = NewToken(TKGe, raw, pos)
	case "<=":
		tok = NewToken(TKLe, raw, pos)

	case "=":
		tok = NewToken(TKAssign, raw, pos)
	case "+=":
		tok = NewToken(TKAddAssign, raw, pos)
	case "-=":
		tok = NewToken(TKSubAssign, raw, pos)
	case "*=":
		tok = NewToken(TKMulAssign, raw, pos)
	case "/=":
		tok = NewToken(TKDivAssign, raw, pos)
	case "%=":
		tok = NewToken(TKModAssign, raw, pos)
	case "**=":
		tok = NewToken(TKExpAssign, raw, pos)

	case "&&":
		tok = NewToken(TKAnd, raw, pos)
	case "||":
		tok = NewToken(TKOr, raw, pos)
	case "!":
		tok = NewToken(TKNot, raw, pos)

	case "..":
		tok = NewToken(TKInclusiveRange, raw, pos)
	case "...":
		tok = NewToken(TKExclusiveRange, raw, pos)

	default:
		tok = &Token{Kind: TKIllegal, Pos: pos}
	}
	cur.Next = tok
	return tok
}

func NewIdent(cur *Token, raw string, pos *Position) *Token {
	var tok *Token
	switch strings.ToLower(raw) {
	case "begin":
		tok = NewToken(TKBegin, raw, pos)
	case "class":
		tok = NewToken(TKClass, raw, pos)
	case "ensure":
		tok = NewToken(TKEnsure, raw, pos)
	case "nil":
		tok = NewToken(TKNil, raw, pos)
	case "self":
		tok = NewToken(TKSelf, raw, pos)
	case "when":
		tok = NewToken(TKWhen, raw, pos)
	case "end":
		tok = NewToken(TKEnd, raw, pos)
	case "def":
		tok = NewToken(TKDef, raw, pos)
	case "false":
		tok = NewToken(TKFalse, raw, pos)
	case "not":
		tok = NewToken(TKNot, raw, pos)
	case "super":
		tok = NewToken(TKSuper, raw, pos)
	case "while":
		tok = NewToken(TKWhile, raw, pos)
	case "alias":
		tok = NewToken(TKAlias, raw, pos)
	case "defined":
		tok = NewToken(TKDefined, raw, pos)
	case "for":
		tok = NewToken(TKFor, raw, pos)
	case "or":
		tok = NewToken(TKOr, raw, pos)
	case "then":
		tok = NewToken(TKThen, raw, pos)
	case "yield":
		tok = NewToken(TKYield, raw, pos)
	case "and":
		tok = NewToken(TKAnd, raw, pos)
	case "do":
		tok = NewToken(TKDo, raw, pos)
	case "if":
		tok = NewToken(TKIf, raw, pos)
	case "redo":
		tok = NewToken(TKRedo, raw, pos)
	case "true":
		tok = NewToken(TKTrue, raw, pos)
	case "line":
		tok = NewToken(TKLine, raw, pos)
	case "else":
		tok = NewToken(TKElse, raw, pos)
	case "in":
		tok = NewToken(TKIn, raw, pos)
	case "rescue":
		tok = NewToken(TKRescue, raw, pos)
	case "undef":
		tok = NewToken(TKUndef, raw, pos)
	case "file":
		tok = NewToken(TKFile, raw, pos)
	case "break":
		tok = NewToken(TKBreak, raw, pos)
	case "elsif":
		tok = NewToken(TKElsif, raw, pos)
	case "module":
		tok = NewToken(TKModule, raw, pos)
	case "retry":
		tok = NewToken(TKRetry, raw, pos)
	case "unless":
		tok = NewToken(TKUnless, raw, pos)
	case "encoding":
		tok = NewToken(TKEncoding, raw, pos)
	case "case":
		tok = NewToken(TKCase, raw, pos)
	case "next":
		tok = NewToken(TKNext, raw, pos)
	case "return":
		tok = NewToken(TKReturn, raw, pos)
	case "until":
		tok = NewToken(TKUntil, raw, pos)
	default:
		tok = NewToken(TKIdent, raw, pos)
	}
	cur.Next = tok
	return tok
}

func NewString(cur *Token, raw string, pos *Position, isDq bool) *Token {
	tok := &Token{
		Kind:          TKString,
		Raw:           raw,
		Pos:           pos,
		IsDoubleQuote: isDq,
		Next:          nil,
	}
	cur.Next = tok
	return tok
}

func NewFloat(cur *Token, raw string, pos *Position, num float64) *Token {
	tok := &Token{
		Kind:          TKFloat,
		Raw:           raw,
		Pos:           pos,
		IsDoubleQuote: false,
		NumFloat:      num,
		NumInt:        0,
		Next:          nil,
	}
	cur.Next = tok
	return tok
}

func NewInt(cur *Token, raw string, pos *Position, num int) *Token {
	tok := &Token{
		Kind:          TKInt,
		Raw:           raw,
		Pos:           pos,
		IsDoubleQuote: false,
		NumFloat:      0,
		NumInt:        num,
		Next:          nil,
	}
	cur.Next = tok
	return tok
}

func NewNL(cur *Token, raw string, pos *Position) *Token {
	tok := NewToken(TKNewLine, raw, pos)
	cur.Next = tok
	return tok
}

func NewComment(cur *Token, raw string, pos *Position) *Token {
	tok := NewToken(TKComment, raw, pos)
	cur.Next = tok
	return tok
}
