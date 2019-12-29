package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "io/ioutil"
import "path/filepath"
import "math/rand"
import "time"

import . "./core"
import . "./nativelib"
import . "./oplib"
import . "./extlib"
import "./script"


func main() {

	/** 创建lib语境 **/
	var libCtx = BindMap{
		Table: 	make(map[string]*Token, 6),
		Tp:		SYS_CTX,
	}
	/** 初始化lib语境，加载原生函数、拓展函数 **/
	InitNative(&libCtx)
	InitOp(&libCtx)
	InitExt(&libCtx)

	var Es = EvalStack{
		MainCtx: &libCtx,
	}

	/** 初始化执行栈，执行初始化脚本 **/
	Es.Init()
	Es.EvalStr(script.ZHScript, Es.MainCtx)
	Es.EvalStr(script.InitScript, Es.MainCtx)


	/** 创建user语境 **/
	var userCtx = BindMap{
		Table:  make(map[string]*Token, 6),
		Father: &libCtx,
		Tp:		USR_CTX,
	}
	Es.MainCtx = &userCtx

	/** 命令行参数不为空时，执行传入的脚本文件 **/
	if len(os.Args) > 1 {

		scriptPath, err := filepath.Abs(os.Args[1])
		if err != nil {
			fmt.Print("错误的文件路径！ \n")
		} else {
			fileData, err := ioutil.ReadFile(scriptPath)
			if err == nil {
				// fmt.Println(string(fileData))
				Es.Init()

				t, err := Es.EvalStr(string(fileData), Es.MainCtx)
				if t != nil && t.Tp != NIL {
					fmt.Println(t.OutputStr())
				} else {
					fmt.Println("")
				}

				if err != nil {
					panic(err)
				}
			} else {
				fmt.Print("读取文件失败 \n")
			}
		}

	}else{
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
		temp = strings.ToLower(temp)
		if temp == "" {
			continue
		}
		
		if temp[len(temp) - 1] == '~' {
			inp += temp[0:len(temp)-1]
			continue
		}else{
			if len(inp) > 0 {
				inp += temp
			}else{
				inp = temp
			}
		}
		
		inp = Trim(inp)
		Es.Init()

		t, err := Es.EvalStr(inp, Es.MainCtx)
		if t != nil && t.Tp != NIL {
			fmt.Println(t.ToString())
		} else {
			fmt.Println("")
		}

		if err != nil {
			panic(err)
		}

		inp = ""
	}

}
