package nativelib

import . "../core"
import "os"
import "os/exec"
import "fmt"

func Quit(es *EvalStack, ctx *BindMap) (*Token, error){
	os.Exit(0)
	return nil, nil
}

func Clear(es *EvalStack, ctx *BindMap) (*Token, error){
	cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
	return nil, nil
}

func TypeOf(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	var result = Token{Tp: DATATYPE}
	if args[1] != nil {
		result.Val = args[1].Tp
	}

	return &result, nil
}

func Do(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	switch args[1].Tp{
	case BLOCK:
		if args[2].Tp == OBJECT {
			return es.Eval(args[1].Tks(), args[2].Ctx())
		}
		return es.Eval(args[1].Tks(), ctx)
	case STRING:
		if args[2].Tp == OBJECT {
			return es.EvalStr(args[1].Str(), args[2].Ctx())
		}
		return es.EvalStr(args[1].Str(), ctx)
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
		if args[2].Tp == OBJECT {
			return es.Eval(args[1].Tks(), args[2].Ctx(), 1)
		}
		return es.Eval(args[1].Tks(), ctx, 1)
	case STRING:
		if args[2].Tp == OBJECT {
			return es.EvalStr(args[1].Str(), args[2].Ctx(), 1)
		}
		return es.EvalStr(args[1].Str(), ctx, 1)
	default:
		var result *Token
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return result, nil
	}
}

func Format(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token
	result.Tp = STRING
	result.Val = args[1].ToString()
	return &result, nil
}

func Copy(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result *Token
	if args[2] != nil && args[2].Tp == LOGIC && args[2].Val.(bool) {
		result = args[1].CloneDeep()
	}else{
		result = args[1].Clone()
	}
	return result, nil
}

func Pprint(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	if args[1].Tp == BLOCK && args[2].Tp == LOGIC {
		fmt.Print("[")
		for idx, item := range args[1].Tks(){
			if args[3].ToBool(){
				if idx == len(args[1].Tks()) - 1 {
					fmt.Print(item.OutputStr())
				}else{
					fmt.Print(item.OutputStr() + " ")
				}
			}else{
				temp, err := item.GetVal(ctx, es)
				if err != nil {
					return nil, err
				}
				if idx == len(args[1].Tks()) - 1 {
					fmt.Print(temp.OutputStr())
				}else{
					fmt.Print(temp.OutputStr() + " ")
				}
				
			}
		}
		if args[2].Val.(bool){
			fmt.Print("]")
		}else{
			fmt.Println("]")
		}
	}else{
		if args[2].Val.(bool){
			fmt.Print(args[1].OutputStr())
		}else{
			fmt.Println(args[1].OutputStr())
		}
	}
	return nil, nil
}

func Let(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	if args[1].Tp == BLOCK {
		var orginSts = es.IsLocal
		es.IsLocal = true
		result, err := es.Eval(args[1].Tks(), ctx)
		es.IsLocal = orginSts
		return result, err
	}
	return &Token{ERR, "Type Mismatch"}, nil
}



