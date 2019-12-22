package main

import "fmt"
import . "./core"
import . "./nativelib"
import . "./oplib"
import . "./extlib"
import "os"
import "bufio"
import "strings"
import "io/ioutil"
import "path/filepath"

func main() {
	// fmt.Println(ToTokens("b/:a")[0].Val.([]*Token)[1].OutputStr())

	var libCtx = BindMap{
		Table: 	make(map[string]*Token, 6),
		Tp:		SYS_CTX,
	}
	InitNative(&libCtx)
	InitOp(&libCtx)
	InitExt(&libCtx)

	var userCtx = BindMap{
		Table:  make(map[string]*Token, 6),
		Father: &libCtx,
		Tp:		USR_CTX,
	}

	var Es = EvalStack{
		MainCtx: &userCtx,
	}

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

	}

	var reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")

		inp, _ := reader.ReadString('\n')
		inp = strings.Replace(inp, "\r\n", "", -1)
		inp = Trim(inp)
		inp = strings.ToLower(inp)
		if inp == "" {
			continue
		}

		Es.Init()

		t, err := Es.EvalStr(inp, Es.MainCtx)
		if t != nil && t.Tp != NIL {
			fmt.Println(t.OutputStr())
		} else {
			fmt.Println("")
		}

		if err != nil {
			panic(err)
		}

	}

}
