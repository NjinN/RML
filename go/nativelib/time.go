package nativelib

import . "../core"
import "time"
import "fmt"


func Cost(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if(args[1].Tp != BLOCK){
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return &result, nil
	}

	var start = time.Now()
	es.Eval(args[1].Val.([]*Token), ctx)
	var end = time.Now()
	fmt.Printf("cost time: %s\n", end.Sub(start))

	return &result, nil
}

