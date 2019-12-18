package nativelib

import . "../core"

func fibonacci(n int) int{
	if n < 2 {
		return n
	}else{
		return fibonacci(n-1)+fibonacci(n-2)
	}
}

func Fib(Es *EvalStack, ctx *BindMap) (*Token, error){
	var args = Es.Line[Es.LastStartPos() : Es.LastEndPos() + 1]

	var result Token
	if args[1].Tp == INTEGER {
		result.Tp = INTEGER
		result.Val = fibonacci(args[1].Val.(int)) 
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

