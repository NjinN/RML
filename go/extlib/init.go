package nativelib

import . "github.com/NjinN/RML/go/core"

func InitExt(ctx *BindMap) {
	var fibToken = Token{
		NATIVE,
		Native{
			"fib",
			2,
			Fib,
			nil,
		},
	}
	ctx.PutNow("fib", &fibToken)

}
