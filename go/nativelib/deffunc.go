package nativelib

import . "../core"


func DefFunc(Es *EvalStack, ctx *BindMap) (*Token, error){
	var args = Es.Line[Es.LastStartPos() : Es.LastEndPos() + 1]
	var result Token

	if(args[1].Tp != BLOCK || args[2].Tp != BLOCK){
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return &result, nil
	}

	for i:=1; i < len(args[1].Val.([]*Token)); i++ {
		if(args[1].Val.([]*Token)[i].Tp != WORD){
			result.Tp = ERR
			result.Val = "Type Mismatch"
			return &result, nil
		}
	}

	result.Tp = FUNC
	result.Val = Func{
		Args: args[1].Val.([]*Token),
		Codes: args[2].Val.([]*Token),
		Ctx: BindMap{make(map[string]*Token, 8), ctx},
	}
	return &result, nil
}

