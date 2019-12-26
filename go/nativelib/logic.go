package nativelib

import . "../core"

func Not(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	return &Token{LOGIC, !args[1].ToBool()}, nil
}

func And(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if args[1].ToBool() && args[2].ToBool(){
		return &Token{LOGIC, true}, nil
	}else{
		return &Token{LOGIC, false}, nil
	}
}

func Or(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if args[1].ToBool() || args[2].ToBool(){
		return &Token{LOGIC, true}, nil
	}else{
		return &Token{LOGIC, false}, nil
	}
}

