package main

import "fmt"
import . "./core"
import . "./nativelib"
import . "./oplib"
import "os"
import "bufio"
import "strings"

func main(){

	var libCtx = BindMap{
		Table: make(map[string]*Token, 6),
	}
	InitNative(&libCtx)
	InitOp(&libCtx)

	var userCtx = BindMap{
		Table: make(map[string]*Token, 6),
		Father: &libCtx,
	}

	var Es = EvalStack{
		MainCtx: &userCtx,
	}

	var reader = bufio.NewReader(os.Stdin)

	for{
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
		if t != nil && t.Tp != NIL{
			fmt.Println(t.OutputStr())
		}else{
			fmt.Println("")
		}

		if err != nil {
			panic(err)
		}

	}

}



