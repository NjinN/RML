package nativelib

import . "../core"
import "os"
import "fmt"

func Quit(es *EvalStack, ctx *BindMap) (*Token, error){
	os.Exit(0)
	return nil, nil
}

func TypeOf(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	var result = Token{Tp: DATATYPE}
	if args[1] != nil {
		result.Val = string(TypeStr(args[1].Tp))
	}

	return &result, nil
}

func Do(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	switch args[1].Tp{
	case BLOCK:
		return es.Eval(args[1].Val.([]*Token), ctx)
	case STRING:
		return es.EvalStr(args[1].Val.(string), ctx)
	default:
		var result *Token
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return result, nil
	}
}

func Reduce(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	switch args[1].Tp{
	case BLOCK:
		return es.Eval(args[1].Val.([]*Token), ctx, 1)
	case STRING:
		return es.EvalStr(args[1].Val.(string), ctx, 1)
	default:
		var result *Token
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return result, nil
	}
}

func Copy(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	var result = Token{args[1].Tp, args[1].Val}
	return &result, nil
}

func Pprint(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	if args[1].Tp == BLOCK {
		for _, item := range args[2].Val.([]*Token){
			fmt.Println(item.OutputStr())
		}
	}else{
		fmt.Println(args[1].OutputStr())
	}
	return nil, nil
}

func Let(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	if args[1].Tp == BLOCK {
		var orginSts = es.IsLocal
		es.IsLocal = true
		result, err := es.Eval(args[1].Val.([]*Token), ctx)
		es.IsLocal = orginSts
		return result, err
	}
	return &Token{ERR, "Type Mismatch"}, nil
}

