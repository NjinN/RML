package moplib

import (
	. "github.com/NjinN/RML/go/core"
)

func InitMop(ctx *BindMap) {

	var elifToken = Token{
		MOP,
		Mop{
			"elif",
			4,
			Elif,
			nil,
		},
	}
	ctx.PutNow("elif", &elifToken)

	var elseToken = Token{
		MOP,
		Mop{
			"else",
			3,
			Eelse,
			nil,
		},
	}
	ctx.PutNow("else", &elseToken)

	var onToken = Token{
		MOP,
		Mop{
			"on",
			3,
			Oon,
			[]int{0, 1},
		},
	}
	ctx.PutNow("on", &onToken)



}
