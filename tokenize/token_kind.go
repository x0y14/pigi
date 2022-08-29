package tokenize

// 参考: https://www.tutorialspoint.com/ruby/ruby_operators.htm

type TokenKind int

const (
	TKIllegal TokenKind = iota
	TKEof

	TKNewLine
	TKWhite
	TKComment

	TKIdent
	TKString
	TKInt
	TKFloat

	symbolsBegin
	TKLrb   // (
	TKRrb   // )
	TKLsb   // [
	TKRsb   // ]
	TKLcb   // {
	TKRcb   // }
	TKDot   // .
	TKComma // ,
	TKColon // :
	TKSemi  // ;
	TKAt    // @
	symbolsEnd

	operatorsBegin
	TKAdd // +
	TKSub // -
	TKMul // *
	TKDiv // /
	TKMod // %
	TKExp // **

	TKEq // ==
	TKNe // !=
	TKGt // >
	TKLt // <
	TKGe // >=
	TKLe // <=
	// <=>
	// ===
	// .eq?
	// .equal?

	TKAssign    // =
	TKAddAssign // +=
	TKSubAssign // -=
	TKMulAssign // *=
	TKDivAssign // /=
	TKModAssign // %=
	TKExpAssign // **=

	// &
	// |
	// ^
	// ~
	// <<
	// >>

	TKAnd // &&
	TKOr  // ||
	TKNot // !

	// ?:

	TKInclusiveRange // ..
	TKExclusiveRange // ...
	operatorsEnd

	reservedIdentBegin // 予約後
	TKBegin
	TKClass
	TKEnsure
	TKNil
	TKSelf
	TKWhen
	TKEnd
	TKDef
	TKFalse
	TKSuper
	TKWhile
	TKAlias
	TKDefined
	TKFor
	TKThen
	TKYield
	TKDo
	TKIf
	TKRedo
	TKTrue
	TKLine // __LINE__
	TKElse
	TKIn
	TKRescue
	TKUndef
	TKFile // __FILE__
	TKBreak
	TKElsif
	TKModule
	TKRetry
	TKUnless
	TKEncoding // __ENCODING__
	TKCase
	TKNext
	TKReturn
	TKUntil
	reservedIdentEnd
)
