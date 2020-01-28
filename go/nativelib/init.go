package nativelib

import . "github.com/NjinN/RML/go/core"

func InitNative(ctx *BindMap) {
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

	var clearToken = Token{
		NATIVE,
		Native{
			"clear",
			1,
			Clear,
			nil,
		},
	}
	ctx.PutNow("clear", &clearToken)

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

	var cmdToken = Token{
		NATIVE,
		Native{
			"_cmd",
			4,
			CallCmd,
			nil,
		},
	}
	ctx.PutNow("_cmd", &cmdToken)

	var helpToken = Token{
		NATIVE,
		Native{
			"help",
			2,
			HelpInfo,
			[]int{0},
		},
	}
	ctx.PutNow("help", &helpToken)
	ctx.PutNow("?", &helpToken)

	var thisToken = Token{
		NATIVE,
		Native{
			"this",
			1,
			ThisRef,
			nil,
		},
	}
	ctx.PutNow("this", &thisToken)

	var thisPortToken = Token{
		NATIVE,
		Native{
			"this-port",
			1,
			ThisPort,
			nil,
		},
	}
	ctx.PutNow("this-port", &thisPortToken)

	var libInfoToken = Token{
		NATIVE,
		Native{
			"_lib?",
			2,
			LibInfo,
			nil,
		},
	}
	ctx.PutNow("_lib?", &libInfoToken)

	var gcToken = Token{
		NATIVE,
		Native{
			"gc",
			1,
			Rgc,
			nil,
		},
	}
	ctx.PutNow("gc", &gcToken)

	var unsetToken = Token{
		NATIVE,
		Native{
			"unset",
			2,
			Unset,
			[]int{0},
		},
	}
	ctx.PutNow("unset", &unsetToken)

	var collectToken = Token{
		NATIVE,
		Native{
			"collect",
			2,
			Collect,
			nil,
		},
	}
	ctx.PutNow("collect", &collectToken)

	/*******  file  *******/

	var nowDirToken = Token{
		NATIVE,
		Native{
			"now-dir",
			1,
			NowDir,
			nil,
		},
	}
	ctx.PutNow("now-dir", &nowDirToken)

	var absPathToken = Token{
		NATIVE,
		Native{
			"abs-path",
			2,
			AbsFilePath,
			nil,
		},
	}
	ctx.PutNow("abs-path", &absPathToken)

	var chDirToken = Token{
		NATIVE,
		Native{
			"cd",
			2,
			ChangeDir,
			nil,
		},
	}
	ctx.PutNow("cd", &chDirToken)

	var lsToken = Token{
		NATIVE,
		Native{
			"_ls",
			2,
			LsDir,
			nil,
		},
	}
	ctx.PutNow("_ls", &lsToken)

	var renameToken = Token{
		NATIVE,
		Native{
			"rename",
			3,
			RenameFile,
			nil,
		},
	}
	ctx.PutNow("rename", &renameToken)

	var removeToken = Token{
		NATIVE,
		Native{
			"remove",
			2,
			RemoveFile,
			nil,
		},
	}
	ctx.PutNow("remove", &removeToken)

	var makeDirToken = Token{
		NATIVE,
		Native{
			"make-dir",
			2,
			makeDir,
			nil,
		},
	}
	ctx.PutNow("make-dir", &makeDirToken)

	var existToken = Token{
		NATIVE,
		Native{
			"exist?",
			2,
			FileExist,
			nil,
		},
	}
	ctx.PutNow("exist?", &existToken)

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

	var readFileToken = Token{
		NATIVE,
		Native{
			"_readfile",
			3,
			ReadFile,
			nil,
		},
	}
	ctx.PutNow("_readfile", &readFileToken)

	var writeFileToken = Token{
		NATIVE,
		Native{
			"_writefile",
			4,
			WriteFile,
			nil,
		},
	}
	ctx.PutNow("_writefile", &writeFileToken)

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

	var untilToken = Token{
		NATIVE,
		Native{
			"until",
			2,
			Until,
			nil,
		},
	}
	ctx.PutNow("until", &untilToken)

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

	var tryToken = Token{
		NATIVE,
		Native{
			"try",
			3,
			Ttry,
			nil,
		},
	}
	ctx.PutNow("try", &tryToken)

	var causeToken = Token{
		NATIVE,
		Native{
			"cause",
			2,
			Cause,
			nil,
		},
	}
	ctx.PutNow("cause", &causeToken)

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
			4,
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

	var getToken = Token{
		NATIVE,
		Native{
			"get",
			3,
			Gget,
			nil,
		},
	}
	ctx.PutNow("get", &getToken)

	var putToken = Token{
		NATIVE,
		Native{
			"put",
			4,
			Pput,
			nil,
		},
	}
	ctx.PutNow("put", &putToken)

	/*******  fork  *******/
	var forkToken = Token{
		NATIVE,
		Native{
			"_fork",
			4,
			Fork,
			nil,
		},
	}
	ctx.PutNow("_fork", &forkToken)

	var spawnToken = Token{
		NATIVE,
		Native{
			"_spawn",
			4,
			Spawn,
			nil,
		},
	}
	ctx.PutNow("_spawn", &spawnToken)

	/*******  parse  *******/

	var parseToken = Token{
		NATIVE,
		Native{
			"parse",
			3,
			Parse,
			nil,
		},
	}
	ctx.PutNow("parse", &parseToken)

	/*******  port  *******/

	var openToken = Token{
		NATIVE,
		Native{
			"open",
			2,
			Oopen,
			nil,
		},
	}
	ctx.PutNow("open", &openToken)

	var readPortToken = Token{
		NATIVE,
		Native{
			"_readport",
			3,
			ReadPort,
			nil,
		},
	}
	ctx.PutNow("_readport", &readPortToken)

	var writePortToken = Token{
		NATIVE,
		Native{
			"_writeport",
			4,
			WritePort,
			nil,
		},
	}
	ctx.PutNow("_writeport", &writePortToken)

	var waitToken = Token{
		NATIVE,
		Native{
			"wait",
			2,
			Wait,
			nil,
		},
	}
	ctx.PutNow("wait", &waitToken)

	var listenToken = Token{
		NATIVE,
		Native{
			"listen",
			2,
			Listen,
			nil,
		},
	}
	ctx.PutNow("listen", &listenToken)

	var closeToken = Token{
		NATIVE,
		Native{
			"close",
			2,
			Close,
			nil,
		},
	}
	ctx.PutNow("close", &closeToken)


	/*******  net  *******/
	var readUrlToken = Token{
		NATIVE,
		Native{
			"_readurl",
			3,
			ReadUrl,
			nil,
		},
	}
	ctx.PutNow("_readurl", &readUrlToken)

}
