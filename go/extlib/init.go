package nativelib

import . "../core"

func InitExt(ctx *BindMap){
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