package nativelib

import . "../core"
import "os"
import "fmt"

func Quit(Es *EvalStack, ctx *BindMap) (*Token, error){
	os.Exit(0)
	return nil, nil
}

func TypeOf(Es *EvalStack, ctx *BindMap) (*Token, error){
	var args = Es.Line[Es.LastStartPos() : Es.LastEndPos() + 1]

	var result = Token{Tp: DATATYPE}
	if args[1] != nil {
		result.Val = string(TypeStr(args[1].Tp))
	}

	return &result, nil
}

func Do(Es *EvalStack, ctx *BindMap) (*Token, error){
	var args = Es.Line[Es.LastStartPos() : Es.LastEndPos() + 1]

	switch args[1].Tp{
	case BLOCK:
		return Es.Eval(args[1].Val.([]*Token), ctx)
	case STRING:
		return Es.EvalStr(args[1].Val.(string), ctx)
	default:
		var result *Token
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return result, nil
	}
}

func Copy(Es *EvalStack, ctx *BindMap) (*Token, error){
	var args = Es.Line[Es.LastStartPos() : Es.LastEndPos() + 1]

	var result = Token{args[1].Tp, args[1].Val}
	return &result, nil
}

func Pprint(Es *EvalStack, ctx *BindMap) (*Token, error){
	var args = Es.Line[Es.LastStartPos() : Es.LastEndPos() + 1]

	if args[1].Tp == BLOCK {
		for _, item := range args[2].Val.([]*Token){
			fmt.Println(item.OutputStr())
		}
	}else{
		fmt.Println(args[1].OutputStr())
	}
	return nil, nil
}