package nativelib

import (
	"fmt"
	"sync"

	. "github.com/NjinN/RML/go/core"
)

func forkEval(inp []*Token, ctx *BindMap, wg *sync.WaitGroup, wait bool, waiter *Token, stackLen int) {
	var evalStack EvalStack
	evalStack.InitWithLen(stackLen)
	evalStack.MainCtx = ctx
	temp, err := evalStack.Eval(inp, ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
	if waiter != nil && temp != nil {
		waiter.Copy(temp)
	}
	if wg != nil && wait {
		wg.Done()
	}
}

func forkEvalStr(inp string, ctx *BindMap, wg *sync.WaitGroup, wait bool, waiter *Token, stackLen int) {
	var evalStack EvalStack
	evalStack.InitWithLen(stackLen)
	evalStack.MainCtx = ctx
	temp, err := evalStack.EvalStr(inp, ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
	if waiter != nil && temp != nil {
		waiter.Copy(temp)
	}
	if wg != nil && wait {
		wg.Done()
	}
}

func Fork(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == BLOCK || args[1].Tp == STRING && args[3].Tp == INTEGER && args[3].Int() > 0 {

		if args[1].Tp == BLOCK {
			if args[2] != nil && args[2].Tp != NONE {
				go forkEval(args[1].CloneDeep().Tks(), ctx, nil, false, args[2], args[3].Int())
			} else {
				go forkEval(args[1].CloneDeep().Tks(), ctx, nil, false, nil, args[3].Int())
			}
		} else if args[1].Tp == STRING {
			if args[2] != nil && args[2].Tp != NONE {
				go forkEvalStr(args[1].Str(), ctx, nil, false, args[2], args[3].Int())
			} else {
				go forkEvalStr(args[1].Str(), ctx, nil, false, nil, args[3].Int())
			}
		}
		return &Token{NIL, nil}, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}

func Spawn(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	if args[1].Tp != BLOCK || args[2].Tp != LOGIC || args[3].Tp != INTEGER || args[3].Int() <= 0 {
		return &Token{ERR, "Type Mismatch"}, nil
	}
	for _, item := range args[1].Tks() {
		if item.Tp != BLOCK && item.Tp != STRING {
			return &Token{ERR, "Type Mismatch"}, nil
		}
	}

	var wg sync.WaitGroup
	for _, item := range args[1].Tks() {
		if item.Tp == BLOCK {
			if args[2].ToBool() {
				wg.Add(1)
				go forkEval(item.CloneDeep().Tks(), ctx, &wg, true, nil, args[3].Int())
			} else {
				go forkEval(item.CloneDeep().Tks(), ctx, nil, false, nil, args[3].Int())
			}
		} else if item.Tp == STRING {
			if args[2].ToBool() {
				wg.Add(1)
				go forkEvalStr(item.Str(), ctx, &wg, true, nil, args[3].Int())
			} else {
				go forkEvalStr(item.Str(), ctx, nil, false, nil, args[3].Int())
			}
		}
	}

	if args[2].ToBool() {
		wg.Wait()
	}

	return &Token{NIL, nil}, nil
}
