package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	. "github.com/NjinN/RML/go/core"
	. "github.com/NjinN/RML/go/nativelib"
	. "github.com/NjinN/RML/go/oplib"
	. "github.com/NjinN/RML/go/moplib"
	. "github.com/NjinN/RML/go/modlib"

	. "github.com/NjinN/RML/go/extlib"

	"github.com/NjinN/RML/go/script"
)

func main() {

	/** 创建lib语境 **/
	var libCtx = BindMap{
		Table: make(map[string]*Token, 6),
		Tp:    SYS_CTX,
	}
	/** 初始化lib语境，加载原生函数、拓展函数 **/
	InitNative(&libCtx)
	InitOp(&libCtx)
	InitMop(&libCtx)
	InitMod(&libCtx)
	
	InitExt(&libCtx)

	var es = EvalStack{
		MainCtx: &libCtx,
	}

	/** 初始化执行栈，执行初始化脚本 **/
	es.Init()
	es.EvalStr(script.ZHScript, es.MainCtx)
	es.EvalStr(script.InitScript, es.MainCtx)

	/** 创建user语境 **/
	var userCtx = BindMap{
		Table:  make(map[string]*Token, 6),
		Father: &libCtx,
		Tp:     USR_CTX,
	}
	es.MainCtx = &userCtx

	/** 命令行参数不为空时，执行传入的脚本文件 **/
	if len(os.Args) > 1 {

		scriptPath, err := filepath.Abs(os.Args[1])
		if err != nil {
			fmt.Print("错误的文件路径！ \n")
		} else {
			fileData, err := ioutil.ReadFile(scriptPath)
			if err == nil {
				os.Chdir(GetParentDir(scriptPath))
				// fmt.Println(string(fileData))
				es.Init()

				t, err := es.EvalStr(string(fileData), es.MainCtx)
				if t != nil && t.Tp != NIL {
					fmt.Println(t.OutputStr())
				} else {
					fmt.Println("")
				}

				if err != nil {
					if err.Error() != "return" && err.Error() != "break" && err.Error() != "continue" {
						panic(err)
					}
				}
			} else {
				fmt.Print("读取文件失败 \n")
			}
		}

	} else {
		rand.Seed(time.Now().UnixNano())
		fmt.Println("如梦令 -- " + Cis[rand.Intn(len(Cis))])
		fmt.Println("RML no-version;\tGratitude to Carl!")
	}

	/** 获取控制台输入并执行 **/
	var reader = bufio.NewReader(os.Stdin)
	var inp string
	for {
		fmt.Print(">> ")

		temp, _ := reader.ReadString('\n')
		temp = strings.Replace(temp, "\r\n", "", -1)
		if temp == "" {
			continue
		}

		if temp[len(temp)-1] == '~' {
			inp += temp[0 : len(temp)-1]
			continue
		} else {
			if len(inp) > 0 {
				inp += temp
			} else {
				inp = temp
			}
		}

		inp = Trim(inp)
		es.Init()

		t, err := es.EvalStr(inp, es.MainCtx)
		if t != nil && t.Tp != NIL {
			fmt.Println(t.ToString())
		} else {
			fmt.Println("")
		}

		if err != nil && err.Error() != "return" && err.Error() != "break" && err.Error() != "continue" {
			panic(err)
		}

		inp = ""
	}

}
