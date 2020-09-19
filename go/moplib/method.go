package moplib

import (

	. "github.com/NjinN/RML/go/core"
)


func Oon(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == OBJECT && (args[2].Tp == WORD || args[2].Tp == STRING) {
		member := args[1].Ctx().GetNow(args[2].Str())
		if member.Tp == FUNC {
			expLen := member.Explen()
			startPos := es.LastStartPos()
			endPos := es.LastEndPos()
			
			es.StartPos.Add(startPos)
			es.EndPos.Pop()
			es.EndPos.Add(startPos + expLen)
			es.EndPos.Add(endPos)
			
			var f Token
			f.Tp = PATH
			f.Val = NewTks(4)
			f.List().Add(member)
			f.List().Add(args[1])
			
			return &f, nil

		}
		return member, nil
	}


	
	return &Token{ERR, "Type Mismatch"}, nil
}



