package nativelib

import . "../core"
import "fmt"
import "sync"

func forkEval(inp []*Token, ctx *BindMap, wg *sync.WaitGroup, wait bool, waiter *Token){
	var evalStack EvalStack
	evalStack.Init()
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

func forkEvalStr(inp string, ctx *BindMap, wg *sync.WaitGroup, wait bool, waiter *Token){
	var evalStack EvalStack
	evalStack.Init()
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

func Fork(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if args[1].Tp == BLOCK || args[1].Tp == STRING {
		

		if args[1].Tp == BLOCK {
			if args[2] != nil && args[2].Tp != NONE {
				go forkEval(args[1].Tks(), ctx, nil, false, args[2])
			}else{
				go forkEval(args[1].Tks(), ctx, nil, false, nil)
			} 
		}else if args[1].Tp == STRING {
			if args[2] != nil && args[2].Tp != NONE {
				go forkEvalStr(args[1].Str(), ctx, nil, false, args[2])
			}else{
				go forkEvalStr(args[1].Str(), ctx, nil, false, nil)
			}
		}
		return nil, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}

func Spawn(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	if args[1].Tp != BLOCK || args[2].Tp != LOGIC {
		return &Token{ERR, "Type Mismatch"}, nil
	}
	for _, item := range args[1].Tks() {
		if item.Tp != BLOCK && item.Tp != STRING {
			return &Token{ERR, "Type Mismatch"}, nil
		}
	}

	var wg sync.WaitGroup
	for _, item := range args[1].Tks(){
		if item.Tp == BLOCK {
			if args[2].ToBool(){
				wg.Add(1)
				go forkEval(item.Tks(), ctx, &wg, true, nil)
			}else{
				go forkEval(item.Tks(), ctx, nil, false, nil)
			}
		}else if item.Tp == STRING {
			if args[2].ToBool(){
				wg.Add(1)
				go forkEvalStr(item.Str(), ctx, &wg, true, nil)
			}else{
				go forkEvalStr(item.Str(), ctx, nil, false, nil)
			}
		}
	}

	if args[2].ToBool(){
		wg.Wait()
	}

	return nil, nil
}

