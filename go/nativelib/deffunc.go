package nativelib

import . "../core"
// import "fmt"

func DefFunc(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if(args[1].Tp != BLOCK || args[2].Tp != BLOCK){
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return &result, nil
	}

	for i:=0; i < len(args[1].Tks()); i++ {
		if(args[1].Tks()[i].Tp != WORD && args[1].Tks()[i].Tp != PROP){
			result.Tp = ERR
			result.Val = "Type Mismatch"
			return &result, nil
		}
	}

	result.Tp = FUNC
	
	var a 	[]*Token
	var p	[]*Token

	for i:=0; i < len(args[1].Tks()); i++ {
		
		if args[1].Tks()[i].Tp == WORD {
			a = append(a, args[1].Tks()[i])
		}else{
			if i == len(args[1].Tks())-1 || args[1].Tks()[i+1].Tp != WORD {
				p = append(p, args[1].Tks()[i])
				p = append(p, nil)
			}else{
				p = append(p, args[1].Tks()[i])
				p = append(p, args[1].Tks()[i+1])
				i++
			}
		}
	}

	result.Val = Func{
		Args: 	a,
		Codes: 	args[2].Tks()[0:],
		Props: 	p,
	}

	return &result, nil
}

