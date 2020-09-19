package moplib

import (

	. "github.com/NjinN/RML/go/core"
)



func Elif(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[3].Tp == BLOCK || args[3].Tp == STRING {
		if args[1].Tp == NIL {
			if !args[2].ToBool(){
				return &Token{NIL, ""}, nil
			}

			if args[3].Tp == BLOCK {
				return es.Eval(args[3].Tks(), ctx)
			}else if args[3].Tp == STRING {
				return es.EvalStr(args[3].Str(), ctx)
			}

		}else if args[1].ToBool(){
			return &Token{LOGIC, true}, nil
		}

		return &Token{NIL, ""}, nil
	}


	
	return &Token{ERR, "Type Mismatch"}, nil
}

func Eelse(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[2].Tp == BLOCK || args[2].Tp == STRING {
		if args[1].Tp == NIL {

			if args[2].Tp == BLOCK {
				return es.Eval(args[2].Tks(), ctx)
			}else if args[2].Tp == STRING {
				return es.EvalStr(args[2].Str(), ctx)
			}

		}else if args[1].ToBool(){
			return &Token{LOGIC, true}, nil
		}

		return &Token{NIL, ""}, nil
	}


	
	return &Token{ERR, "Type Mismatch"}, nil
}

