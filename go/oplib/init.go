package oplib

import (
	. "github.com/NjinN/RML/go/core"
	. "github.com/NjinN/RML/go/nativelib"
)

func InitOp(ctx *BindMap) {

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

	var modToken = Token{
		OP,
		Native{
			"%",
			3,
			Mod,
			nil,
		},
	}
	ctx.PutNow("%", &modToken)

	var addSetToken = Token{
		OP,
		Native{
			"+=",
			3,
			AddSet,
			nil,
		},
	}
	ctx.PutNow("+=", &addSetToken)

	var subSetToken = Token{
		OP,
		Native{
			"-=",
			3,
			SubSet,
			nil,
		},
	}
	ctx.PutNow("-=", &subSetToken)

	var mulSetToken = Token{
		OP,
		Native{
			"*=",
			3,
			MulSet,
			nil,
		},
	}
	ctx.PutNow("*=", &mulSetToken)

	var divSetToken = Token{
		OP,
		Native{
			"/=",
			3,
			DivSet,
			nil,
		},
	}
	ctx.PutNow("/=", &divSetToken)

	var modSetToken = Token{
		OP,
		Native{
			"%=",
			3,
			ModSet,
			nil,
		},
	}
	ctx.PutNow("%=", &modSetToken)

	var swapToken = Token{
		OP,
		Native{
			"><",
			3,
			Swap,
			nil,
		},
	}
	ctx.PutNow("><", &swapToken)

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

	/*******  logic  *******/

	var andToken = Token{
		OP,
		Native{
			"and",
			3,
			And,
			nil,
		},
	}
	ctx.PutNow("and", &andToken)

	var orToken = Token{
		OP,
		Native{
			"or",
			3,
			Or,
			nil,
		},
	}
	ctx.PutNow("or", &orToken)

}
