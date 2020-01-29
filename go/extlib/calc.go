package nativelib

import . "github.com/NjinN/RML/go/core"

func fibonacci(n int) int {
	if n < 2 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func Fib(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == INTEGER {
		result.Tp = INTEGER
		result.Val = fibonacci(args[1].Int())
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}
