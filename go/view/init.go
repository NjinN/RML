package view

import (

	. "github.com/NjinN/RML/go/core"
)

func InitView(ctx *BindMap) {

	var formToken = Token{
		NATIVE,
		Native{
			"form",
			3,
			Fform,
			nil,
		},
	}
	ctx.PutNow("form", &formToken)

	var showToken = Token{
		NATIVE,
		Native{
			"show",
			2,
			Sshow,
			nil,
		},
	}
	ctx.PutNow("show", &showToken)

}
