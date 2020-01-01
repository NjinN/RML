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
		if(args[1].Tks()[i].Tp != WORD && args[1].Tks()[i].Tp != PROP && args[1].Tks()[i].Tp != STRING){
			result.Tp = ERR
			result.Val = "Type Mismatch"
			return &result, nil
		}
	}

	result.Tp = FUNC
	
	var a = NewTks(8)
	var p = NewTks(8)
	var desc 	[]string

	if len(args[1].Tks()) > 0 && args[1].Tks()[0].Tp == STRING {
		desc = append(desc, "self", args[1].Tks()[0].Str())
	}

	for i:=0; i < len(args[1].Tks()); i++ {
		
		if args[1].Tks()[i].Tp == WORD {
			a.Add(args[1].Tks()[i])
			if i + 1 < len(args[1].Tks()) && args[1].Tks()[i+1].Tp == STRING {
				desc = append(desc, args[1].Tks()[i].Str(), args[1].Tks()[i+1].Str())
			}
		}else if args[1].Tks()[i].Tp == PROP {
			if i == len(args[1].Tks())-1 || args[1].Tks()[i+1].Tp != WORD {
				p.Add(args[1].Tks()[i])
				p.Add(nil)
				if i + 1 < len(args[1].Tks()) && args[1].Tks()[i+1].Tp == STRING {
					desc = append(desc, args[1].Tks()[i].Str(), args[1].Tks()[i+1].Str())
				}
			}else{
				p.Add(args[1].Tks()[i])
				p.Add(args[1].Tks()[i+1])
				if i + 2 < len(args[1].Tks()) && args[1].Tks()[i+2].Tp == STRING {
					desc = append(desc, args[1].Tks()[i+1].Str(), args[1].Tks()[i+2].Str())
				}
				i++
			}
		}
	}

	result.Val = Func{
		Args: 	a,
		// Codes: 	args[2].Tks()[0:],
		Codes: &TokenList{uint(len(args[2].Tks())), uint(len(args[2].Tks())), args[2].Tks()},
		Props: 	p,
		Desc: 	desc,
	}

	return &result, nil
}

