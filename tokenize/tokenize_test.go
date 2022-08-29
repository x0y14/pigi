package tokenize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewL1Position(lp int) *Position {
	return NewPosition(1, lp, lp)
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		expect *Token
	}{
		{
			"wild plus",
			"1+1",
			&Token{
				Kind:   TKInt,
				Pos:    NewPosition(1, 0, 0),
				NumInt: 1,
				Raw:    "1",
				Next: &Token{
					Kind:   TKAdd,
					Pos:    NewPosition(1, 1, 1),
					NumInt: 0,
					Raw:    "+",
					Next: &Token{
						Kind:   TKInt,
						Pos:    NewPosition(1, 2, 2),
						NumInt: 1,
						Raw:    "1",
						Next: &Token{
							Kind:   TKEof,
							Pos:    NewPosition(1, 3, 3),
							NumInt: 0,
							Raw:    "",
							Next:   nil,
						},
					},
				},
			},
		},
		{
			"wild string",
			"\"string\"",
			&Token{
				Kind:          TKString,
				Pos:           NewL1Position(0),
				NumInt:        0,
				Raw:           "string",
				IsDoubleQuote: true,
				Next: &Token{
					Kind:   TKEof,
					Pos:    NewL1Position(len("\"string\"")),
					NumInt: 0,
					Raw:    "",
					Next:   nil,
				},
			},
		},
		{
			"wild comp symbol",
			"/=",
			&Token{
				Kind: TKDivAssign,
				Pos:  NewL1Position(0),
				Raw:  "/=",
				Next: &Token{
					Kind: TKEof,
					Pos:  NewL1Position(len("/=")),
					Raw:  "",
					Next: nil,
				},
			},
		},
		{
			"wild single symbol",
			"+",
			&Token{
				Kind: TKAdd,
				Pos:  NewL1Position(0),
				Raw:  "+",
				Next: &Token{
					Kind: TKEof,
					Pos:  NewL1Position(1),
					Raw:  "",
					Next: nil,
				},
			},
		},
		{
			"wild ident",
			"ident",
			&Token{
				Kind: TKIdent,
				Pos:  NewL1Position(0),
				Raw:  "ident",
				Next: &Token{
					Kind: TKEof,
					Pos:  NewL1Position(len("ident")),
					Raw:  "",
					Next: nil,
				},
			},
		},
		{
			"welcome",
			`3.times do
  print 'Welcome '
end`,
			&Token{
				Kind:          TKInt,
				Raw:           "3",
				Pos:           NewPosition(1, 0, 0),
				IsDoubleQuote: false,
				NumFloat:      0,
				NumInt:        3,
				Next: &Token{
					Kind:          TKDot,
					Raw:           ".",
					Pos:           NewPosition(1, 1, 1),
					IsDoubleQuote: false,
					NumFloat:      0,
					NumInt:        0,
					Next: &Token{
						Kind:          TKIdent,
						Raw:           "times",
						Pos:           NewPosition(1, 2, 2),
						IsDoubleQuote: false,
						NumFloat:      0,
						NumInt:        0,
						Next: &Token{
							Kind:          TKWhite,
							Raw:           " ",
							Pos:           NewPosition(1, len("3.times"), len("3.times")),
							IsDoubleQuote: false,
							NumFloat:      0,
							NumInt:        0,
							Next: &Token{
								Kind:          TKDo,
								Raw:           "do",
								Pos:           NewPosition(1, len("3.times "), len("3.times ")),
								IsDoubleQuote: false,
								NumFloat:      0,
								NumInt:        0,
								Next: &Token{
									Kind:          TKNewLine,
									Raw:           "\n",
									Pos:           NewPosition(1, len("3.times do"), len("3.times do")),
									IsDoubleQuote: false,
									NumFloat:      0,
									NumInt:        0,
									Next: &Token{
										Kind:          TKWhite,
										Raw:           "  ",
										Pos:           NewPosition(2, 0, len("3.times do\n")),
										IsDoubleQuote: false,
										NumFloat:      0,
										NumInt:        0,
										Next: &Token{
											Kind:          TKIdent,
											Raw:           "print",
											Pos:           NewPosition(2, 2, len("3.times do\n  ")),
											IsDoubleQuote: false,
											NumFloat:      0,
											NumInt:        0,
											Next: &Token{
												Kind:          TKWhite,
												Raw:           " ",
												Pos:           NewPosition(2, len("  print"), len("3.times do\n  print")),
												IsDoubleQuote: false,
												NumFloat:      0,
												NumInt:        0,
												Next: &Token{
													Kind:          TKString,
													Raw:           "Welcome ",
													Pos:           NewPosition(2, len("  print "), len("3.times do\n  print ")),
													IsDoubleQuote: false,
													NumFloat:      0,
													NumInt:        0,
													Next: &Token{
														Kind:          TKNewLine,
														Raw:           "\n",
														Pos:           NewPosition(2, len("  print 'Welcome '"), len("3.times do\n  print 'Welcome '")),
														IsDoubleQuote: false,
														NumFloat:      0,
														NumInt:        0,
														Next: &Token{
															Kind:          TKEnd,
															Raw:           "end",
															Pos:           NewPosition(3, 0, len("3.times do\n  print 'Welcome '\n")),
															IsDoubleQuote: false,
															NumFloat:      0,
															NumInt:        0,
															Next: &Token{
																Kind:          TKEof,
																Raw:           "",
																Pos:           NewPosition(3, len("end"), len("3.times do\n  print 'Welcome '\nend")),
																IsDoubleQuote: false,
																NumFloat:      0,
																NumInt:        0,
																Next:          nil,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok, err := Tokenize(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expect, tok)
		})
	}
}
