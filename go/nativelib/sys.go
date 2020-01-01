package nativelib

import . "../core"
import "os"
import "os/exec"
import "bytes"
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
		return &Token{ERR, "Type Mismatch"}, nil
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
				if idx == args[1].List().Len() - 1 {
					fmt.Print(item.OutputStr())
				}else{
					fmt.Print(item.OutputStr() + " ")
				}
			}else{
				temp, err := item.GetVal(ctx, es)
				if err != nil {
					return nil, err
				}
				if idx == args[1].List().Len() - 1 {
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

func CallCmd(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	if args[1].Tp == STRING && args[2].Tp == LOGIC {
		cmd := exec.Command("cmd", "/c", args[1].Str())
		if args[3] != nil && args[3].Tp != NONE {
			var output bytes.Buffer
			cmd.Stdout = &output
			cmd.Run()
			args[3].Tp = BIN
			args[3].Val = append(make([]byte, 0), output.Bytes()...)
			return args[3], nil
		}else if args[2].ToBool() {
			cmd.Stdout = os.Stdout
			cmd.Start()
		}else{
			cmd.Stdout = os.Stdout
			cmd.Run()
		}	
		return nil, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}

func HelpInfo(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	if args[1] == nil {
		fmt.Println("nil")
		return nil, nil
	}
	temp, err:= args[1].GetVal(ctx, es)
	if err != nil {
		return &Token{ERR, "Type Mismatch"}, nil
	}
	if temp.Tp == FUNC {
		fmt.Println(temp.Val.(Func).GetFuncInfo())
		return nil, nil
	}
	fmt.Println("This is a " + TypeToStr(temp.Tp) + ", format to " + temp.ToString())
	return nil, nil

}

func ThisRef(es *EvalStack, ctx *BindMap) (*Token, error){
	var c = ctx
	for c.Tp != USR_CTX && c.Father != nil {
		c = c.Father
	}

	return &Token{OBJECT, c}, nil
}

func LibInfo(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var c = ctx
	for c.Father != nil {
		c = c.Father
	}

	result := "Lib:"
	for k, v := range c.Table {
		result += "  " + k + "\t"
		if len([]rune(k)) < 4 {
			result += "\t"
		}
		str := v.ToString()
		runes := []rune(str)
		if len(runes) > 50 {
			str = string(runes[0:50]) + "..."
		}
		result += str + "\n"
	}

	fmt.Print(result)
	if args[1] != nil && args[1].Tp != NONE {
		args[1].Tp = STRING
		args[1].Val = result
	}
	return nil, nil
}

