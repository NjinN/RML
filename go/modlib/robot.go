package modlib

import (
	"path/filepath"
	"strings"
	. "github.com/NjinN/RML/go/core"
	. "github.com/NjinN/RML/go/nativelib"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func ScrollMouse(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == INTEGER && args[2].Tp == STRING {
		robotgo.ScrollMouse(args[1].Int(), args[2].Str())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func MouseClick(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING && args[2].Tp == LOGIC {
		robotgo.MouseClick(args[1].Str(), args[2].ToBool())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func MoveMouse(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == INTEGER && args[2].Tp == INTEGER {
		robotgo.MoveMouse(args[1].Int(), args[2].Int())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func MoveMouseSmooth(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == INTEGER && args[2].Tp == INTEGER && args[3].Tp == DECIMAL && args[4].Tp == DECIMAL {
		robotgo.MoveMouseSmooth(args[1].Int(), args[2].Int(), args[3].Float(), args[4].Float())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}


func TypeStr(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {
		robotgo.TypeStr(args[1].Str())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func KeyTap(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token

	if args[1].Tp == STRING || args[1].Tp == BLOCK {
		result.Tp = LOGIC
		result.Val = true

		if args[1].Tp == STRING {
			robotgo.KeyTap(args[1].Str())
			return &result, nil
		}

		if args[1].Tp == BLOCK && args[1].List().Len() > 0 {
			checked := true
			key := ""
			arr := []string{}
			for idx, item := range args[1].Tks(){
				if item.Tp != STRING {
					checked = false
					break
				}else{
					if idx == 0 {
						key = item.Str()
					}else{
						arr = append(arr, item.Str())
					}
				}
			}
			if checked {
				robotgo.KeyTap(key, arr)
				return &result, nil
			}
		}

	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Ccopy(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {
		robotgo.WriteAll(args[1].Str())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func ReadCopy(es *EvalStack, ctx *BindMap) (*Token, error) {
	// var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	text, err := robotgo.ReadAll()
	if err == nil {
		result.Tp = STRING
		result.Val = text
	}else{
		result.Tp = NONE
		result.Val = ""
	}
	
	return &result, nil
}

func GetMousePos(es *EvalStack, ctx *BindMap) (*Token, error) {
	// var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	result.Tp = BLOCK
	result.Val = NewTks(2)

	x, y := robotgo.GetMousePos()
	
	result.List().Add(&Token{Tp: INTEGER, Val: x})
	result.List().Add(&Token{Tp: INTEGER, Val: y})
	
	return &result, nil
}


func GetPixelColor(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == INTEGER && args[2].Tp == INTEGER {
		color := robotgo.GetPixelColor(args[1].Int(), args[2].Int())
		
		result.Tp = STRING
		result.Val = color
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func CaptureScreenAll(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == FILE {
		cap := robotgo.CaptureScreen()
		
		filePath, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		saveBit := robotgo.SaveBitmap(cap, filePath)

		result.Tp = STRING
		result.Val = saveBit
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func CaptureScreen(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == INTEGER && args[2].Tp == INTEGER && args[3].Tp == INTEGER && args[4].Tp == INTEGER && args[5].Tp == FILE {
		cap := robotgo.CaptureScreen(args[1].Int(), args[2].Int(), args[3].Int(), args[4].Int())
		
		filePath, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		saveBit := robotgo.SaveBitmap(cap, filePath)

		result.Tp = STRING
		result.Val = saveBit
		return &result, nil
	
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func FindPic(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == FILE && args[2].Tp == FILE && args[3].Tp == DECIMAL {
		f1, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		f2, err := filepath.Abs(strings.ReplaceAll(args[2].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		bitMap := robotgo.OpenBitmap(f2)

		x, y := robotgo.FindPic(f1, bitMap, args[3].Float())

		result.Tp = BLOCK
		result.Val = NewTks(2)
		result.List().Add(&Token{Tp: INTEGER, Val: x})
		result.List().Add(&Token{Tp: INTEGER, Val: y})
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func MatchScreen(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == FILE && args[2].Tp == DECIMAL {
		f1, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		x, y := robotgo.FindPic(f1, nil, args[2].Float())

		result.Tp = BLOCK
		result.Val = NewTks(2)
		result.List().Add(&Token{Tp: INTEGER, Val: x})
		result.List().Add(&Token{Tp: INTEGER, Val: y})
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func EventHook(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING && (args[2].Tp == BLOCK || args[2].Tp == STRING) && args[3].Tp == BLOCK {
		checked := true
		arr := []string{}
		if args[2].Tp == BLOCK {
			for _, item := range args[2].Tks(){
				if item.Tp != STRING {
					checked = false
				}else{
					arr = append(arr, item.Str())
				}
			}
		}else if args[2].Tp == STRING {
			arr = append(arr, args[2].Str())
		}
		
		if checked {
			robotgo.EventHook(getHookType(args[1].Str()), arr, func(e hook.Event) {
				FORKS++
				go ForkEval(args[3].CloneDeep().Tks(), ctx, nil, false, nil, 1024)
				FORKS--
			})

			result.Tp = LOGIC
			result.Val = true
			return &result, nil
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func EventHooks(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == BLOCK {
		if args[1].List().Len() % 2 != 0 {
			return &Token{Tp:ERR, Val:"Type Mismatch"}, nil
		}

		idx := 0
		for idx < args[1].List().Len(){
			if args[1].List().Get(idx).List().Len() < 2 {
				return &Token{Tp:ERR, Val:"Type Mismatch"}, nil
			}
			for _, item := range args[1].List().Get(idx).Tks() {
				if item.Tp != STRING {
					return &Token{Tp:ERR, Val:"Type Mismatch"}, nil
				}
			}
			idx += 2
		}

		idx = 0
		
		for idx < args[1].List().Len(){
			arr := []string{}
			for i, item := range args[1].List().Get(idx).Tks() {
				if i > 0 {
					arr = append(arr, item.Str())
				}
				
			}

			code := args[1].List().Get(idx + 1).CloneDeep()
			robotgo.EventHook(getHookType(args[1].List().Get(idx).List().Get(0).Str()), arr, func(e hook.Event) {
				FORKS++
				go ForkEval(code.Tks(), ctx, nil, false, nil, 1024)
				FORKS--
			})
			idx += 2
		}
			
		
		result.Tp = LOGIC
		result.Val = true
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}


func getHookType(str string) uint8{
	switch str {
	case "KeyDown":
		return hook.KeyDown
	case "KeyHold":
		return hook.KeyHold
	case "KeyUp":
		return hook.KeyUp
	case "MouseUp":
		return hook.MouseUp
	case "MouseHold":
		return hook.MouseHold
	case "MouseDown":
		return hook.MouseDown
	case "MouseMove":
		return hook.MouseMove
	case "MouseDrag":
		return hook.MouseDrag
	case "MouseWheel":
		return hook.MouseWheel
	default:
		return hook.KeyDown
	}

}

func StartHook(es *EvalStack, ctx *BindMap) (*Token, error) {
	// var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
	
	result.Tp = LOGIC
	result.Val = true
	return &result, nil
}

func EndHook(es *EvalStack, ctx *BindMap) (*Token, error) {
	// var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	robotgo.EventEnd()
	
	result.Tp = LOGIC
	result.Val = true
	return &result, nil
}

func FindPids(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {
		fpids, err := robotgo.FindIds(args[1].Str())

		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		
		result.Tp = BLOCK
		result.Val = NewTks(8)
		for _, fpid := range fpids {
			result.List().Add(&Token{Tp: INTEGER, Val: int(fpid)})
		}
		
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func ActivePid(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == INTEGER {
		err := robotgo.ActivePID(int32(args[1].Int()))

		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		
		result.Tp = LOGIC
		result.Val = true
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func KillPid(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == INTEGER {
		err := robotgo.Kill(int32(args[1].Int()))

		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		
		result.Tp = LOGIC
		result.Val = true
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func ActiveProc(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {
		err := robotgo.ActiveName(args[1].Str())

		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		
		result.Tp = LOGIC
		result.Val = true
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func PidExists(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == INTEGER {
		isExist, err := robotgo.PidExists(int32(args[1].Int()))

		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		
		result.Tp = LOGIC
		result.Val = isExist
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func ShowAlert(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING && args[2].Tp == STRING {
		ibool := robotgo.ShowAlert(args[1].Str(), args[2].Str())

		result.Tp = INTEGER
		result.Val = ibool
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func GetTitle(es *EvalStack, ctx *BindMap) (*Token, error) {
	var result Token
	str := robotgo.GetTitle()
	result.Tp = STRING
	result.Val = str
	return &result, nil
}

func GetPid(es *EvalStack, ctx *BindMap) (*Token, error) {
	var result Token
	pid := robotgo.GetPID()
	result.Tp = INTEGER
	result.Val = int(pid)
	return &result, nil
}

func Is64Bit(es *EvalStack, ctx *BindMap) (*Token, error) {
	var result Token
	b := robotgo.Is64Bit()
	result.Tp = LOGIC
	result.Val = b
	return &result, nil
}








