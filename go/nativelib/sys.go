package nativelib

import (
	"bytes"
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"syscall"
	"unsafe"
	"strings"

	. "github.com/NjinN/RML/go/core"
)




func Quit(es *EvalStack, ctx *BindMap) (*Token, error) {
	os.Exit(0)
	return &Token{NIL, nil}, nil
}

func Clear(es *EvalStack, ctx *BindMap) (*Token, error) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return &Token{NIL, nil}, nil
}

func TypeOf(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	var result = Token{Tp: DATATYPE}
	if args[1] != nil {
		result.Val = args[1].Tp
	}

	return &result, nil
}

func Do(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	switch args[1].Tp {
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

func Reduce(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	switch args[1].Tp {
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

func Format(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token
	result.Tp = STRING
	result.Val = args[1].ToString()
	return &result, nil
}

func Copy(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result *Token
	if args[2] != nil && args[2].Tp == LOGIC && args[2].Val.(bool) {
		result = args[1].CloneDeep()
	} else {
		result = args[1].Clone()
	}
	return result, nil
}

func Pprint(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	if args[1].Tp == BLOCK && args[2].Tp == LOGIC {
		fmt.Print("[")
		for idx, item := range args[1].Tks() {
			if args[3].ToBool() {
				if idx == args[1].List().Len()-1 {
					fmt.Print(item.OutputStr())
				} else {
					fmt.Print(item.OutputStr() + " ")
				}
			} else {
				temp, err := item.GetVal(ctx, es)
				if err != nil {
					return nil, err
				}
				if idx == args[1].List().Len()-1 {
					fmt.Print(temp.OutputStr())
				} else {
					fmt.Print(temp.OutputStr() + " ")
				}

			}
		}
		if args[2].Val.(bool) {
			fmt.Print("]")
		} else {
			fmt.Println("]")
		}
	} else {
		if args[2].Val.(bool) {
			fmt.Print(args[1].OutputStr())
		} else {
			fmt.Println(args[1].OutputStr())
		}
	}
	// runtime.GC()
	return &Token{NIL, nil}, nil
}

//控制台输入准备 start
var getStdHandle, getConsoleMode, setConsoleMode *syscall.LazyProc

func WinInitConsole() {
	lib := syscall.NewLazyDLL("Kernel32.dll")
	getStdHandle = lib.NewProc("GetStdHandle")
	getConsoleMode = lib.NewProc("GetConsoleMode")
	setConsoleMode = lib.NewProc("SetConsoleMode")
}

func WinPrompt(msg string) string {
	handle, _, _ := getStdHandle.Call(uintptr(^uint(10) + 1))
	var mode uint
	getConsoleMode.Call(handle, uintptr(unsafe.Pointer(&mode)))
	newMode := mode & ^uint(4)
	setConsoleMode.Call(handle, uintptr(newMode))
	passwd, err := bufio.NewReader(os.Stdin).ReadString('\n')
	setConsoleMode.Call(handle, uintptr(mode))
	if err != nil {
		panic(err)
	}
	return strings.Replace(passwd, "\r\n", "", 1)
}

func IsWinOs() bool {
	os := runtime.GOOS
	return strings.Index(os, "win") >= 0 
}

//控制台输入准备 end

func Ask(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token
	if args[1].Tp == STRING {
		fmt.Print(args[1].Str())
		str := ""
		if args[2].ToBool() {
			if IsWinOs(){
				WinInitConsole()
				str = WinPrompt(args[1].Str())
			}else{
				return &Token{ERR, "OS don't support"}, nil
			}
		}else{
			reader := bufio.NewReader(os.Stdin)
			str, _ = reader.ReadString('\n')
			str = strings.Replace(str, "\r\n", "", 1)
		}
		result.Tp = STRING
		result.Val = str
		return &result, nil
	} 
	return &Token{ERR, "Type Mismatch"}, nil
}

func Let(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	if args[1].Tp == BLOCK {
		var orginSts = es.IsLocal
		es.IsLocal = true
		result, err := es.Eval(args[1].Tks(), ctx)
		es.IsLocal = orginSts
		return result, err
	}
	return &Token{ERR, "Type Mismatch"}, nil
}

func CallCmd(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	if args[1].Tp == STRING && args[2].Tp == LOGIC {
		cmd := exec.Command("cmd", "/c", args[1].Str())
		if args[3] != nil && args[3].Tp != NONE {
			var output bytes.Buffer
			cmd.Stdout = &output
			cmd.Run()
			args[3].Tp = BIN
			args[3].Val = append(make([]byte, 0), output.Bytes()...)
			return args[3], nil
		} else if args[2].ToBool() {
			cmd.Stdout = os.Stdout
			cmd.Start()
		} else {
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
		return &Token{NIL, nil}, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}

func HelpInfo(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	if args[1] == nil {
		fmt.Println("nil")
		return &Token{NIL, nil}, nil
	}
	temp, err := args[1].GetVal(ctx, es)
	if err != nil {
		return &Token{ERR, "Type Mismatch"}, nil
	}
	if temp.Tp == FUNC {
		fmt.Println(temp.Val.(Func).GetFuncInfo())
		return &Token{NIL, nil}, nil
	}
	fmt.Println("This is a " + TypeToStr(temp.Tp) + ", format to " + temp.ToString())
	return &Token{NIL, nil}, nil

}

func ThisRef(es *EvalStack, ctx *BindMap) (*Token, error) {
	var c = ctx
	for c.Tp != USR_CTX && c.Father != nil {
		c = c.Father
	}

	return &Token{OBJECT, c}, nil
}

func ThisPort(es *EvalStack, ctx *BindMap) (*Token, error) {
	var c = ctx
	for c.Tp != USR_CTX && c.Father != nil {
		c = c.Father
	}

	return &Token{PORT, c}, nil
}

func LibInfo(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
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
	return &Token{NIL, nil}, nil
}

func Rgc(es *EvalStack, ctx *BindMap) (*Token, error) {
	runtime.GC()
	return &Token{NIL, nil}, nil
}

func Unset(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == WORD {
		ctx.Unset(args[1].Str())
		return &Token{NIL, nil}, nil
	} else if args[1].Tp == PATH && args[1].List().Last().Tp == WORD {
		var holderPath = args[1].CloneDeep()
		holderPath.List().Pop()
		holder, err := holderPath.GetPathVal(ctx, es)
		if err != nil {
			return &Token{ERR, "Error Path"}, nil
		}
		if holder.Tp != OBJECT {
			return &Token{ERR, "Type Mismatch"}, nil
		}

		holder.Ctx().Unset(args[1].List().Last().Str())
		return &Token{NIL, nil}, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}

/********  collect实现开始  *********/

func keep(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	ctx.Get("__result__").List().Add(args[1].CloneDeep())
	return &Token{NIL, nil}, nil
}

func Collect(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
	var result = Token{BLOCK, NewTks(8)}
	c.PutNow("__result__", &result)
	c.PutNow("keep", &Token{
		NATIVE,
		Native{
			"keep",
			2,
			keep,
			nil,
		},
	})

	if args[1].Tp == BLOCK {
		temp, err := es.Eval(args[1].Tks(), &c)
		if (temp != nil && temp.Tp == ERR) || err != nil {
			return temp, err
		}
		return &result, nil
	} else if args[1].Tp == STRING {
		temp, err := es.EvalStr(args[1].Str(), &c)
		if (temp != nil && temp.Tp == ERR) || err != nil {
			return temp, err
		}
		return &result, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}

/********  collect实现结束  *********/
