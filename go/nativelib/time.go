package nativelib

import (
	"fmt"
	"time"

	. "github.com/NjinN/RML/go/core"
)

func Cost(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token

	if args[1].Tp != BLOCK {
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return &result, nil
	}

	var start = time.Now()
	es.Eval(args[1].Tks(), ctx)
	var end = time.Now()
	fmt.Printf("cost time: %s\n", end.Sub(start))

	return &result, nil
}
