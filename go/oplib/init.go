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

}