package view

import (
	// "fmt"

	. "github.com/NjinN/RML/go/core"
	_ "github.com/NjinN/RML/go/view/winRes"

	"github.com/ying32/govcl/vcl"
)


func init(){
	vcl.Application.Initialize()
	// vcl.Application.SetMainFormOnTaskBar(true)
}

func Fform(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	if args[1].Tp == OBJECT && args[2].Tp == BLOCK {
		mft := NewRForm(args[1], args[2])
		mft.View().Raw.(*RForm).Update()
		return mft, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}




func Sshow(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos():es.LastEndPos()]

	if args[1].Tp == VIEW {
		if args[1].View().Tp == "Form" {
			args[1].View().Raw.(*RForm).Show()
			vcl.Application.Run()
		}

		return &Token{NIL, nil}, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}


func MakeViewGroup(t *Token, f *Token) *TokenList {
	if t.Tp != BLOCK {
		return nil
	}

	var result *TokenList


	return result
}


