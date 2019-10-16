package oplib

import . "../core"
import . "../nativelib"

func InitOp(ctx *BindMap){

	var addToken = Token{
		OP,
		Native{
			"+",
			3,
			Add,
			nil,
		},
	}
	ctx.PutNow("+", &addToken)

	var subToken = Token{
		OP,
		Native{
			"-",
			3,
			Sub,
			nil,
		},
	}
	ctx.PutNow("-", &subToken)

	var mulToken = Token{
		OP,
		Native{
			"*",
			3,
			Mul,
			nil,
		},
	}
	ctx.PutNow("*", &mulToken)

	var divToken = Token{
		OP,
		Native{
			"/",
			3,
			Div,
			nil,
		},
	}
	ctx.PutNow("/", &divToken)

	var eqToken = Token{
		OP,
		Native{
			"=",
			3,
			Eq,
			nil,
		},
	}
	ctx.PutNow("=", &eqToken)

	var gtToken = Token{
		OP,
		Native{
			">",
			3,
			Gt,
			nil,
		},
	}
	ctx.PutNow(">", &gtToken)

	var ltToken = Token{
		OP,
		Native{
			"<",
			3,
			Lt,
			nil,
		},
	}
	ctx.PutNow("<", &ltToken)

	var geToken = Token{
		OP,
		Native{
			">=",
			3,
			Ge,
			nil,
		},
	}
	ctx.PutNow(">=", &geToken)

	var leToken = Token{
		OP,
		Native{
			"<=",
			3,
			Le,
			nil,
		},
	}
	ctx.PutNow("<=", &leToken)
}