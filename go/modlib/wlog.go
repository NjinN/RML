package modlib

import (
	"strings"
	"os"

	. "github.com/NjinN/RML/go/core"

	. "github.com/dixonwille/wlog/v3"
)


var reader = strings.NewReader("\n")
var wlogUI = New(reader, os.Stdout, os.Stdout)
var colorWlogUI = AddColor(BrightBlue, Red, Blue, White, BrightWhite, Magenta, BrightGreen, Green, Yellow, wlogUI )
var safeWlogUI = AddConcurrent(colorWlogUI)

func WlogLog(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {
		safeWlogUI.Log(args[1].Str())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func WlogOutput(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {
		safeWlogUI.Output(args[1].Str())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func WlogSuccess(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {
		safeWlogUI.Success(args[1].Str())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func WlogInfo(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {
		safeWlogUI.Info(args[1].Str())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func WlogError(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {
		safeWlogUI.Error(args[1].Str())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func WlogWarn(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {
		safeWlogUI.Warn(args[1].Str())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func WlogRunning(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {
		safeWlogUI.Running(args[1].Str())

		result.Tp = NONE
		result.Val = ""
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func WlogAsk(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	var result Token
	if args[1].Tp == STRING {

		str, err := safeWlogUI.Ask(args[1].Str(), "")

		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		result.Tp = STRING
		result.Val = str
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}





