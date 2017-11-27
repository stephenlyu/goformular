// Code generated by goyacc - DO NOT EDIT.

package easylang

import __yyfmt__ "fmt"

var _context = newContext()

type yySymType struct {
	yys          int
	value        float64
	str          string
	expr         expression
	descriptions []string
	arguments    []expression
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault   = 57374
	yyEofCode   = 57344
	AND         = 57362
	COLONEQUAL  = 57352
	COMMA       = 57355
	DIVIDE      = 57372
	DOLLAR      = 57359
	DOT         = 57357
	EQ          = 57363
	EQUALS      = 57350
	GE          = 57366
	GT          = 57365
	ID          = 57346
	LE          = 57368
	LPAREN      = 57353
	LT          = 57367
	MINUS       = 57369
	NE          = 57364
	NOT         = 57360
	NUM         = 57347
	OR          = 57361
	PARAMEQUAL  = 57351
	PLUS        = 57370
	POUND       = 57358
	RPAREN      = 57354
	SEMI        = 57356
	STRING      = 57348
	STRING_EXPR = 57349
	TIMES       = 57371
	UNARY       = 57373
	yyErrCode   = 57345

	yyMaxDepth = 200
	yyTabOfs   = -45
)

var (
	yyPrec = map[int]int{
		OR:     0,
		AND:    1,
		EQ:     2,
		NE:     2,
		GT:     3,
		GE:     3,
		LT:     3,
		LE:     3,
		MINUS:  4,
		PLUS:   4,
		TIMES:  5,
		DIVIDE: 5,
		UNARY:  6,
	}

	yyXLAT = map[int]int{
		57369: 0,  // MINUS (53x)
		57355: 1,  // COMMA (43x)
		57356: 2,  // SEMI (39x)
		57361: 3,  // OR (37x)
		57354: 4,  // RPAREN (35x)
		57346: 5,  // ID (34x)
		57353: 6,  // LPAREN (32x)
		57347: 7,  // NUM (32x)
		57362: 8,  // AND (31x)
		57363: 9,  // EQ (29x)
		57364: 10, // NE (29x)
		57360: 11, // NOT (29x)
		57348: 12, // STRING (29x)
		57366: 13, // GE (27x)
		57365: 14, // GT (27x)
		57368: 15, // LE (27x)
		57367: 16, // LT (27x)
		57370: 17, // PLUS (24x)
		57384: 18, // postfix_expression (21x)
		57385: 19, // primary_expression (21x)
		57390: 20, // unary_expression (21x)
		57372: 21, // DIVIDE (19x)
		57371: 22, // TIMES (19x)
		57383: 23, // multiplicative_expression (17x)
		57375: 24, // additive_expression (15x)
		57386: 25, // relational_expression (11x)
		57344: 26, // $end (10x)
		57377: 27, // equality_expression (9x)
		57382: 28, // logical_and_expression (8x)
		57378: 29, // expression (7x)
		57381: 30, // graph_description_list (3x)
		57358: 31, // POUND (3x)
		57389: 32, // statement_suffix (3x)
		57357: 33, // DOT (2x)
		57380: 34, // graph_description (2x)
		57387: 35, // statement (2x)
		57376: 36, // argument_expression_list (1x)
		57352: 37, // COLONEQUAL (1x)
		57350: 38, // EQUALS (1x)
		57379: 39, // formula (1x)
		57351: 40, // PARAMEQUAL (1x)
		57388: 41, // statement_list (1x)
		57374: 42, // $default (0x)
		57359: 43, // DOLLAR (0x)
		57345: 44, // error (0x)
		57349: 45, // STRING_EXPR (0x)
		57373: 46, // UNARY (0x)
	}

	yySymNames = []string{
		"MINUS",
		"COMMA",
		"SEMI",
		"OR",
		"RPAREN",
		"ID",
		"LPAREN",
		"NUM",
		"AND",
		"EQ",
		"NE",
		"NOT",
		"STRING",
		"GE",
		"GT",
		"LE",
		"LT",
		"PLUS",
		"postfix_expression",
		"primary_expression",
		"unary_expression",
		"DIVIDE",
		"TIMES",
		"multiplicative_expression",
		"additive_expression",
		"relational_expression",
		"$end",
		"equality_expression",
		"logical_and_expression",
		"expression",
		"graph_description_list",
		"POUND",
		"statement_suffix",
		"DOT",
		"graph_description",
		"statement",
		"argument_expression_list",
		"COLONEQUAL",
		"EQUALS",
		"formula",
		"PARAMEQUAL",
		"statement_list",
		"$default",
		"DOLLAR",
		"error",
		"STRING_EXPR",
		"UNARY",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {39, 1},
		2:  {41, 1},
		3:  {41, 2},
		4:  {35, 4},
		5:  {35, 4},
		6:  {35, 10},
		7:  {35, 2},
		8:  {32, 1},
		9:  {32, 2},
		10: {30, 2},
		11: {30, 3},
		12: {34, 1},
		13: {19, 1},
		14: {19, 1},
		15: {19, 1},
		16: {19, 3},
		17: {18, 1},
		18: {18, 4},
		19: {18, 3},
		20: {18, 3},
		21: {18, 5},
		22: {36, 1},
		23: {36, 3},
		24: {20, 1},
		25: {20, 2},
		26: {20, 2},
		27: {23, 1},
		28: {23, 3},
		29: {23, 3},
		30: {24, 1},
		31: {24, 3},
		32: {24, 3},
		33: {25, 1},
		34: {25, 3},
		35: {25, 3},
		36: {25, 3},
		37: {25, 3},
		38: {27, 1},
		39: {27, 3},
		40: {27, 3},
		41: {28, 1},
		42: {28, 3},
		43: {29, 1},
		44: {29, 3},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [85][]uint16{
		// 0
		{57, 5: 49, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 60, 61, 27: 62, 63, 50, 35: 48, 39: 46, 41: 47},
		{26: 45},
		{57, 5: 49, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 60, 61, 44, 62, 63, 50, 35: 129},
		{43, 5: 43, 43, 43, 11: 43, 43, 26: 43},
		{32, 32, 32, 32, 6: 87, 8: 32, 32, 32, 13: 32, 32, 32, 32, 32, 21: 32, 32, 31: 89, 33: 88, 37: 115, 114, 40: 116},
		// 5
		{1: 108, 106, 96, 30: 107, 32: 105},
		{31, 31, 31, 31, 31, 8: 31, 31, 31, 13: 31, 31, 31, 31, 31, 21: 31, 31},
		{30, 30, 30, 30, 30, 8: 30, 30, 30, 13: 30, 30, 30, 30, 30, 21: 30, 30},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 60, 61, 27: 62, 63, 103},
		{28, 28, 28, 28, 28, 8: 28, 28, 28, 13: 28, 28, 28, 28, 28, 21: 28, 28},
		// 10
		{21, 21, 21, 21, 21, 8: 21, 21, 21, 13: 21, 21, 21, 21, 21, 21: 21, 21},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 102},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 101},
		{18, 18, 18, 18, 18, 8: 18, 18, 18, 13: 18, 18, 18, 18, 18, 21: 18, 18},
		{15, 15, 15, 15, 15, 8: 15, 15, 15, 13: 15, 15, 15, 15, 15, 21: 79, 78},
		// 15
		{76, 12, 12, 12, 12, 8: 12, 12, 12, 13: 12, 12, 12, 12, 75},
		{1: 7, 7, 7, 7, 8: 7, 7, 7, 13: 73, 71, 72, 70},
		{1: 4, 4, 4, 4, 8: 4, 67, 68},
		{1: 2, 2, 2, 2, 8: 64},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 60, 61, 27: 66},
		// 20
		{32, 32, 32, 32, 32, 6: 87, 8: 32, 32, 32, 13: 32, 32, 32, 32, 32, 21: 32, 32, 31: 89, 33: 88},
		{1: 3, 3, 3, 3, 8: 3, 67, 68},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 60, 86},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 60, 69},
		{1: 5, 5, 5, 5, 8: 5, 5, 5, 13: 73, 71, 72, 70},
		// 25
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 85},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 84},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 83},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 74},
		{76, 8, 8, 8, 8, 8: 8, 8, 8, 13: 8, 8, 8, 8, 75},
		// 30
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 82},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 77},
		{13, 13, 13, 13, 13, 8: 13, 13, 13, 13: 13, 13, 13, 13, 13, 21: 79, 78},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 81},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 80},
		// 35
		{16, 16, 16, 16, 16, 8: 16, 16, 16, 13: 16, 16, 16, 16, 16, 21: 16, 16},
		{17, 17, 17, 17, 17, 8: 17, 17, 17, 13: 17, 17, 17, 17, 17, 21: 17, 17},
		{14, 14, 14, 14, 14, 8: 14, 14, 14, 13: 14, 14, 14, 14, 14, 21: 79, 78},
		{76, 9, 9, 9, 9, 8: 9, 9, 9, 13: 9, 9, 9, 9, 75},
		{76, 10, 10, 10, 10, 8: 10, 10, 10, 13: 10, 10, 10, 10, 75},
		// 40
		{76, 11, 11, 11, 11, 8: 11, 11, 11, 13: 11, 11, 11, 11, 75},
		{1: 6, 6, 6, 6, 8: 6, 6, 6, 13: 73, 71, 72, 70},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 60, 61, 27: 62, 63, 95, 36: 94},
		{5: 91},
		{5: 90},
		// 45
		{25, 25, 25, 25, 25, 8: 25, 25, 25, 13: 25, 25, 25, 25, 25, 21: 25, 25},
		{26, 26, 26, 26, 26, 8: 26, 26, 26, 13: 26, 26, 26, 26, 26, 21: 26, 26, 31: 92},
		{5: 93},
		{24, 24, 24, 24, 24, 8: 24, 24, 24, 13: 24, 24, 24, 24, 24, 21: 24, 24},
		{1: 99, 4: 98},
		// 50
		{1: 23, 3: 96, 23},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 60, 61, 27: 62, 97},
		{1: 1, 1, 1, 1, 8: 64},
		{27, 27, 27, 27, 27, 8: 27, 27, 27, 13: 27, 27, 27, 27, 27, 21: 27, 27},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 60, 61, 27: 62, 63, 100},
		// 55
		{1: 22, 3: 96, 22},
		{19, 19, 19, 19, 19, 8: 19, 19, 19, 13: 19, 19, 19, 19, 19, 21: 19, 19},
		{20, 20, 20, 20, 20, 8: 20, 20, 20, 13: 20, 20, 20, 20, 20, 21: 20, 20},
		{3: 96, 104},
		{29, 29, 29, 29, 29, 8: 29, 29, 29, 13: 29, 29, 29, 29, 29, 21: 29, 29},
		// 60
		{38, 5: 38, 38, 38, 11: 38, 38, 26: 38},
		{37, 5: 37, 37, 37, 11: 37, 37, 26: 37},
		{1: 112, 111},
		{5: 110, 34: 109},
		{1: 35, 35},
		// 65
		{1: 33, 33},
		{36, 5: 36, 36, 36, 11: 36, 36, 26: 36},
		{5: 110, 34: 113},
		{1: 34, 34},
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 60, 61, 27: 62, 63, 127},
		// 70
		{57, 5: 65, 53, 51, 11: 56, 52, 18: 55, 54, 58, 23: 59, 60, 61, 27: 62, 63, 125},
		{6: 117},
		{7: 118},
		{1: 119},
		{7: 120},
		// 75
		{1: 121},
		{7: 122},
		{4: 123},
		{2: 124},
		{39, 5: 39, 39, 39, 11: 39, 39, 26: 39},
		// 80
		{1: 108, 106, 96, 30: 107, 32: 126},
		{40, 5: 40, 40, 40, 11: 40, 40, 26: 40},
		{1: 108, 106, 96, 30: 107, 32: 128},
		{41, 5: 41, 41, 41, 11: 41, 41, 26: 41},
		{42, 5: 42, 42, 42, 11: 42, 42, 26: 42},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("'%c'", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 44

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 4:
		{
			if _, ok := yyS[yypt-1].expr.(*stringexpr); ok {
				lexer, _ := yylex.(*yylexer)
				_context.addError(GeneralError(lexer.lineno, lexer.column, "string can't be right value"))
			}
			yyVAL.expr = AssignmentExpression(_context, yyS[yypt-3].str, yyS[yypt-1].expr, false)
			_context.addOutput(yyVAL.expr.VarName(), yyS[yypt-0].descriptions, 0, 0)
		}
	case 5:
		{
			if _, ok := yyS[yypt-1].expr.(*stringexpr); ok {
				lexer, _ := yylex.(*yylexer)
				_context.addError(GeneralError(lexer.lineno, lexer.column, "string can't be right value"))
			}
			yyVAL.expr = AssignmentExpression(_context, yyS[yypt-3].str, yyS[yypt-1].expr, false)
			_context.addNotOutputVar(yyVAL.expr.VarName(), yyS[yypt-0].descriptions, 0, 0)
		}
	case 6:
		{
			yyVAL.expr = ParamExpression(_context, yyS[yypt-9].str, yyS[yypt-6].value, yyS[yypt-4].value, yyS[yypt-2].value)
		}
	case 7:
		{
			if _, ok := yyS[yypt-1].expr.(*stringexpr); ok {
				lexer, _ := yylex.(*yylexer)
				_context.addError(GeneralError(lexer.lineno, lexer.column, "string can't be right value"))
			}
			varName := _context.newAnonymousVarName()
			yyVAL.expr = AssignmentExpression(_context, varName, yyS[yypt-1].expr, true)
			_context.addOutput(varName, yyS[yypt-0].descriptions, 0, 0)
		}
	case 9:
		{
			yyVAL.descriptions = yyS[yypt-1].descriptions
		}
	case 10:
		{
			if !IsValidDescription(yyS[yypt-0].str) {
				lexer, _ := yylex.(*yylexer)
				_context.addError(BadGraphDescError(lexer.lineno, lexer.column, yyS[yypt-0].str))
			}
			yyVAL.descriptions = append(yyVAL.descriptions, yyS[yypt-0].str)
		}
	case 11:
		{
			if !IsValidDescription(yyS[yypt-0].str) {
				lexer, _ := yylex.(*yylexer)
				_context.addError(BadGraphDescError(lexer.lineno, lexer.column, yyS[yypt-0].str))
			}
			yyVAL.descriptions = append(yyS[yypt-2].descriptions, yyS[yypt-0].str)
		}
	case 12:
		{
			yyVAL.str = yyS[yypt-0].str
		}
	case 13:
		{
			expr := _context.defined(yyS[yypt-0].str)
			if expr == nil {
				expr = _context.definedParam(yyS[yypt-0].str)
			}
			if expr != nil {
			} else if funcName, ok := noArgFuncMap[yyS[yypt-0].str]; ok {
				expr = FunctionExpression(_context, funcName, nil)
			} else {
				lexer, _ := yylex.(*yylexer)
				_context.addError(UndefinedVarError(lexer.lineno, lexer.column, yyS[yypt-0].str))
				expr = ErrorExpression(_context, yyS[yypt-0].str)
			}
			yyVAL.expr = expr
		}
	case 14:
		{
			yyVAL.expr = ConstantExpression(_context, yyS[yypt-0].value)
		}
	case 15:
		{
			yyVAL.expr = StringExpression(_context, yyS[yypt-0].str[1:len(yyS[yypt-0].str)-1])
		}
	case 16:
		{
			yyVAL.expr = yyS[yypt-1].expr
		}
	case 17:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 18:
		{
			if _, ok := funcMap[yyS[yypt-3].str]; !ok {
				lexer, _ := yylex.(*yylexer)
				_context.addError(UndefinedFunctionError(lexer.lineno, lexer.column, yyS[yypt-3].str))
				yyVAL.expr = ErrorExpression(_context, yyS[yypt-3].str)
			} else {
				yyVAL.expr = FunctionExpression(_context, yyS[yypt-3].str, yyS[yypt-1].arguments)
			}
		}
	case 19:
		{
			if !_context.isReferenceSupport(yyS[yypt-2].str, yyS[yypt-0].str) {
				lexer, _ := yylex.(*yylexer)
				_context.addError(GeneralError(lexer.lineno, lexer.column, __yyfmt__.Sprintf("%s.%s not supported", yyS[yypt-2].str, yyS[yypt-0].str)))
				yyVAL.expr = ErrorExpression(_context, yyS[yypt-0].str)
			} else {
				yyVAL.expr = ReferenceExpression(_context, yyS[yypt-2].str, yyS[yypt-0].str)
			}
		}
	case 20:
		{
			period := translatePeriod(yyS[yypt-0].str)
			if funcName, ok := noArgFuncMap[yyS[yypt-2].str]; ok && _context.isPeriodSupport(period) {
				yyVAL.expr = CrossFunctionExpression(_context, funcName, "", period)
			} else {
				lexer, _ := yylex.(*yylexer)
				_context.addError(GeneralError(lexer.lineno, lexer.column, __yyfmt__.Sprintf("%s#%s not supported", yyS[yypt-2].str, yyS[yypt-0].str)))
				yyVAL.expr = ErrorExpression(_context, yyS[yypt-0].str)
			}
		}
	case 21:
		{
			period := translatePeriod(yyS[yypt-0].str)
			if !_context.isReferenceSupport(yyS[yypt-4].str, yyS[yypt-2].str) || !_context.isPeriodSupport(period) {
				lexer, _ := yylex.(*yylexer)
				_context.addError(GeneralError(lexer.lineno, lexer.column, __yyfmt__.Sprintf("%s.%s#%s not supported", yyS[yypt-4].str, yyS[yypt-2].str, yyS[yypt-0].str)))
				yyVAL.expr = ErrorExpression(_context, yyS[yypt-2].str)
			} else {
				yyVAL.expr = CrossReferenceExpression(_context, yyS[yypt-4].str, yyS[yypt-2].str, "", period)
			}
		}
	case 22:
		{
			yyVAL.arguments = []expression{yyS[yypt-0].expr}
		}
	case 23:
		{
			yyVAL.arguments = append(yyVAL.arguments, yyS[yypt-0].expr)
		}
	case 24:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 25:
		{
			yyVAL.expr = UnaryExpression(_context, yyS[yypt-1].str, yyS[yypt-0].expr)
		}
	case 26:
		{
			yyVAL.expr = UnaryExpression(_context, yyS[yypt-1].str, yyS[yypt-0].expr)
		}
	case 27:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 28:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 29:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 30:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 31:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 32:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 33:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 34:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 35:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 36:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 37:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 38:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 39:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 40:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 41:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 42:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}
	case 43:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 44:
		{
			yyVAL.expr = BinaryExpression(_context, yyS[yypt-1].str, yyS[yypt-2].expr, yyS[yypt-0].expr)
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
