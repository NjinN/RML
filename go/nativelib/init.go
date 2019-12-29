package nativelib

import . "../core"



func InitNative(ctx *BindMap){
	/*******  sys  *******/

	var quitToken = Token{
		NATIVE,
		Native{
			"quite",
			1,
			Quit,
			nil,
		},
	}
	ctx.PutNow("quit", &quitToken)
	ctx.PutNow("q", &quitToken)

	var typeOfToken = Token{
		NATIVE,
		Native{
			"type?",
			2,
			TypeOf,
			nil,
		},
	}
	ctx.PutNow("type?", &typeOfToken)

	var toToken = Token{
		NATIVE,
		Native{
			"to",
			3,
			To,
			nil,
		},
	}
	ctx.PutNow("to", &toToken)

	var doToken = Token{
		NATIVE,
		Native{
			"_do",
			3,
			Do,
			nil,
		},
	}
	ctx.PutNow("_do", &doToken)

	var reduceToken = Token{
		NATIVE,
		Native{
			"_reduce",
			3,
			Reduce,
			nil,
		},
	}
	ctx.PutNow("_reduce", &reduceToken)

	var formatToken = Token{
		NATIVE,
		Native{
			"format",
			2,
			Format,
			nil,
		},
	}
	ctx.PutNow("format", &formatToken)

	var copyToken = Token{
		NATIVE,
		Native{
			"_copy",
			3,
			Copy,
			nil,
		},
	}
	ctx.PutNow("_copy", &copyToken)

	var printToken = Token{
		NATIVE,
		Native{
			"_print",
			4,
			Pprint,
			nil,
		},
	}
	ctx.PutNow("_print", &printToken)

	var letToken = Token{
		NATIVE,
		Native{
			"let",
			2,
			Let,
			nil,
		},
	}
	ctx.PutNow("let", &letToken)

	var loadToken = Token{
		NATIVE,
		Native{
			"load",
			2,
			Load,
			nil,
		},
	}
	ctx.PutNow("load", &loadToken)

	var readToken = Token{
		NATIVE,
		Native{
			"_read",
			3,
			Read,
			nil,
		},
	}
	ctx.PutNow("_read", &readToken)


	/*******  math  *******/

	var addToken = Token{
		NATIVE,
		Native{
			"add",
			3,
			Add,
			nil,
		},
	}
	ctx.PutNow("add", &addToken)

	var subToken = Token{
		NATIVE,
		Native{
			"sub",
			3,
			Sub,
			nil,
		},
	}
	ctx.PutNow("sub", &subToken)

	var mulToken = Token{
		NATIVE,
		Native{
			"mul",
			3,
			Mul,
			nil,
		},
	}
	ctx.PutNow("mul", &mulToken)

	var divToken = Token{
		NATIVE,
		Native{
			"div",
			3,
			Div,
			nil,
		},
	}
	ctx.PutNow("div", &divToken)

	var modToken = Token{
		NATIVE,
		Native{
			"mod",
			3,
			Mod,
			nil,
		},
	}
	ctx.PutNow("mod", &modToken)

	var addSetToken = Token{
		NATIVE,
		Native{
			"addset",
			3,
			AddSet,
			nil,
		},
	}
	ctx.PutNow("addset", &addSetToken)

	var subSetToken = Token{
		NATIVE,
		Native{
			"subset",
			3,
			SubSet,
			nil,
		},
	}
	ctx.PutNow("subset", &subSetToken)

	var mulSetToken = Token{
		NATIVE,
		Native{
			"mulSet",
			3,
			MulSet,
			nil,
		},
	}
	ctx.PutNow("mulset", &mulSetToken)

	var divSetToken = Token{
		NATIVE,
		Native{
			"divset",
			3,
			DivSet,
			nil,
		},
	}
	ctx.PutNow("divset", &divSetToken)

	var modSetToken = Token{
		NATIVE,
		Native{
			"modset",
			3,
			ModSet,
			nil,
		},
	}
	ctx.PutNow("modset", &modSetToken)

	var swapToken = Token{
		NATIVE,
		Native{
			"swap",
			3,
			Swap,
			nil,
		},
	}
	ctx.PutNow("swap", &swapToken)


	/*******  logic  *******/
	
	var notToken = Token{
		NATIVE,
		Native{
			"not",
			2,
			Not,
			nil,
		},
	}
	ctx.PutNow("not", &notToken)


	/*******  control  *******/

	var ifToken = Token{
		NATIVE,
		Native{
			"if",
			3,
			Iif,
			nil,
		},
	}
	ctx.PutNow("if", &ifToken)

	var eitherToken = Token{
		NATIVE,
		Native{
			"either",
			4,
			Either,
			nil,
		},
	}
	ctx.PutNow("either", &eitherToken)

	var loopToken = Token{
		NATIVE,
		Native{
			"loop",
			3,
			Loop,
			nil,
		},
	}
	ctx.PutNow("loop", &loopToken)

	var repeatToken = Token{
		NATIVE,
		Native{
			"repeat",
			4,
			Repeat,
			[]int{0, 1, 1},
		},
	}
	ctx.PutNow("repeat", &repeatToken)

	var forToken = Token{
		NATIVE,
		Native{
			"for",
			6,
			Ffor,
			[]int{0, 1, 1, 1, 1},
		},
	}
	ctx.PutNow("for", &forToken)

	var whileToken = Token{
		NATIVE,
		Native{
			"while",
			3,
			Wwhile,
			nil,
		},
	}
	ctx.PutNow("while", &whileToken)

	var breakToken = Token{
		NATIVE,
		Native{
			"break",
			1,
			Bbreak,
			nil,
		},
	}
	ctx.PutNow("break", &breakToken)

	var continueToken = Token{
		NATIVE,
		Native{
			"continue",
			1,
			Ccontinue,
			nil,
		},
	}
	ctx.PutNow("continue", &continueToken)

	var returnToken = Token{
		NATIVE,
		Native{
			"return",
			2,
			Rreturn,
			nil,
		},
	}
	ctx.PutNow("return", &returnToken)

	var foreachToken = Token{
		NATIVE,
		Native{
			"foreach",
			4,
			Fforeach,
			[]int{0, 1, 1},
		},
	}
	ctx.PutNow("foreach", &foreachToken)

	/*******  deffunc  *******/

	var defFuncToken = Token{
		NATIVE,
		Native{
			"func",
			3,
			DefFunc,
			nil,
		},
	}
	ctx.PutNow("func", &defFuncToken)


	/*******  compare  *******/

	var eqToken = Token{
		NATIVE,
		Native{
			"eq",
			3,
			Eq,
			nil,
		},
	}
	ctx.PutNow("eq", &eqToken)

	var gtToken = Token{
		NATIVE,
		Native{
			"gt",
			3,
			Gt,
			nil,
		},
	}
	ctx.PutNow("gt", &gtToken)

	var ltToken = Token{
		NATIVE,
		Native{
			"lt",
			3,
			Lt,
			nil,
		},
	}
	ctx.PutNow("lt", &ltToken)

	var geToken = Token{
		NATIVE,
		Native{
			"ge",
			3,
			Ge,
			nil,
		},
	}
	ctx.PutNow("ge", &geToken)

	var leToken = Token{
		NATIVE,
		Native{
			"le",
			3,
			Le,
			nil,
		},
	}
	ctx.PutNow("le", &leToken)


	/*******  time  *******/

	var costToken = Token{
		NATIVE,
		Native{
			"cost",
			2,
			Cost,
			nil,
		},
	}
	ctx.PutNow("cost", &costToken)


	/*******  collect  *******/

	var lenToken = Token{
		NATIVE,
		Native{
			"len?",
			2,
			Length,
			nil,
		},
	}
	ctx.PutNow("len?", &lenToken)

	var insertToken = Token{
		NATIVE,
		Native{
			"_insert",
			5,
			Insert,
			nil,
		},
	}
	ctx.PutNow("_insert", &insertToken)

	var appendToken = Token{
		NATIVE,
		Native{
			"_append",
			4,
			Append,
			nil,
		},
	}
	ctx.PutNow("_append", &appendToken)

	var takeToken = Token{
		NATIVE,
		Native{
			"_take",
			5,
			Take,
			nil,
		},
	}
	ctx.PutNow("_take", &takeToken)

	var replaceToken = Token{
		NATIVE,
		Native{
			"_replace",
			6,
			Replace,
			nil,
		},
	}
	ctx.PutNow("_replace", &replaceToken)


}