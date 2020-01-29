package core

import "errors"
// import "fmt"

type EvalStack struct {
	StartPos 		*IntList
	EndPos 			*IntList
	Line			[]*Token
	Idx 			int
	MainCtx 		*BindMap
	QuoteList		[]int
	IsLocal 		bool
}


func (es *EvalStack) Init(){
	es.StartPos = NewIntList(8)
	es.EndPos = NewIntList(8)
	es.Line = make([]*Token, 1024*1024)
	es.Idx = 0
	es.QuoteList = make([]int, 0)
	es.IsLocal = false
}

func (es *EvalStack) InitWithLen(len int){
	es.StartPos = NewIntList(8)
	es.EndPos = NewIntList(8)
	es.Line = make([]*Token, len)
	es.Idx = 0
	es.QuoteList = make([]int, 0)
	es.IsLocal = false
}

func (es *EvalStack) Push(t *Token){
	es.Line[es.Idx] = t
	es.Idx += 1
}

func (es *EvalStack) LastStartPos() int{
	if(es.StartPos.Len() <= 0){
		return -999999
	}
	return es.StartPos.Last()
}

func (es *EvalStack) LastEndPos() int{
	if(es.EndPos.Len() <= 0){
		return -1
	}
	return es.EndPos.Last()
}

func (es *EvalStack) EvalStr(inpStr string, ctx *BindMap, args ...int) (*Token, error){
	return es.Eval(ToTokens(inpStr, ctx, es), ctx, args...)
}
  

func (es *EvalStack) Eval(inp []*Token, ctx *BindMap, args ...int) (*Token, error){
	var result *Token
	var resultBlk = NewTks(8)

	if(len(inp) == 0){
		return result, nil
	}

	// fmt.Println("------  start eval -------")
	// for _, item := range inp {
	// 	fmt.Println(item.OutputStr())
	// }
	// fmt.Println("------  end eval -------")

	var startIdx = es.Idx
	var startDeep = es.EndPos.Len()
	
	var i = 0
	for i < len(inp){
		var nowToken = inp[i]
		var nextToken *Token
		
		var skip = false
		if nowToken.Tp == GET_WORD || nowToken.Tp == WRAP {
			skip = true
		}

		var nextSkip = false
		if(i < len(inp) - 1){
			nextToken = inp[i+1]
			if nextToken.Tp == GET_WORD || nextToken.Tp == WRAP {
				nextSkip = true
			}
			if(nextToken.Tp == WORD){
				temp, err := nextToken.GetVal(ctx, es)
				if err != nil {
					return nextToken, nil
				}
				nextToken = temp
			}
		}


		if(nextToken != nil && nextToken.Tp == OP && (startDeep == 0 || es.Idx >= es.EndPos.Get(startDeep - 1)) && !nextSkip){
			if(es.StartPos.Len() == 0 || es.Line[es.LastStartPos()].Tp != OP){
				es.StartPos.Add(es.Idx)
				es.Push(nextToken)
				temp, err := nowToken.GetVal(ctx, es)
				if err != nil {
					return temp, err
				}
				es.EndPos.Add(es.Idx + 2)
				es.Push(temp)
			}else if(es.StartPos.Len() == 0 || es.Line[es.LastStartPos()].Tp == OP){
				temp, err := nowToken.GetVal(ctx, es)
				if err != nil {
					return temp, err
				}
				es.Push(temp)
				es.EvalExp(ctx)
				es.Push(es.Line[es.Idx - 1])
				es.Line[es.Idx - 2] = nextToken
				es.StartPos.Add(es.Idx - 2)
				es.EndPos.Add(es.Idx + 1)
			}
			i++
		}else{
			if len(es.QuoteList) > 0 {
				
				if(es.QuoteList[0] > 0){
					temp, err := nowToken.GetVal(ctx, es)
					if err != nil {
						return temp, err
					}
					nowToken = temp
				}

				if len(es.QuoteList) > 0 { //todo I don't know why
					es.QuoteList = es.QuoteList[1 :]
				}
			}else{
				var temp *Token
				var err error
				if !(nowToken != nil && nowToken.Tp == PATH && nowToken.IsSetPath()){
					temp, err = nowToken.GetVal(ctx, es)
					nowToken = temp
				}
				if err != nil {
					return temp, err
				}
			}
			if nowToken == nil {
				continue
			}else if nowToken.Tp == ERR{
				return nowToken, nil
			}else if nowToken.Tp == OP && !skip && es.Idx >= 1{
				if(es.Idx > startIdx){
					es.StartPos.Add(es.Idx - 1)
					es.EndPos.Add(es.Idx + 2) 
					es.Push(es.Line[es.Idx - 1])
					es.Line[es.Idx - 2] = nowToken
				}else{
					result.Tp = ERR
					result.Val = "Illegal grammar!!!"
					return result, errors.New("Illegal grammar!!!")
				}
			}else if nowToken.Tp < SET_WORD{
				es.Push(nowToken)
				if len(args) > 0 && args[0] == 1 && es.StartPos.Len() == startDeep {
					resultBlk.Add(nowToken.Clone())
				}
			}else if nowToken.Tp == PATH {
				if nowToken.Tks()[0] != nil && nowToken.Tks()[0].Tp == FUNC {
					es.StartPos.Add(es.Idx)
					es.EndPos.Add(es.Idx + nowToken.GetPathExpLen())
					es.Push(nowToken)
				}else if nowToken.IsSetPath() {
					es.StartPos.Add(es.Idx)
					es.EndPos.Add(es.Idx + 2)
					es.Push(nowToken)
					i++
					continue
				}else{
					es.Push(nowToken)
				}	
			}else{
				if(nowToken.Tp == NATIVE){
					if(len(nowToken.Val.(Native).QuoteList) > 0){
						es.QuoteList = append(es.QuoteList, nowToken.Val.(Native).QuoteList...)
					}
				}else if(nowToken.Tp == FUNC){
					if(len(nowToken.Val.(Func).QuoteList) > 0){
						es.QuoteList = append(es.QuoteList, nowToken.Val.(Func).QuoteList...)
					}
				}

				if !skip {
					es.StartPos.Add(es.Idx) 
					es.EndPos.Add(es.Idx + nowToken.Explen()) 
				}
				es.Push(nowToken)
			}
		}
 
		for(es.EndPos.Len() > startDeep && es.Idx == es.LastEndPos()){
			temp, err := es.EvalExp(ctx)
			if err != nil {
				return temp, err
			}
			if temp != nil && temp.Tp == ERR {
				getStackErrInfo(es, inp, i, temp)
				es.Line[es.Idx - 1] = temp
				break
			}
			if len(args) > 0 && args[0] == 1 {
				resultBlk.Add(temp)
			}
		}

		i++
	}
	result = es.Line[es.Idx - 1]
	es.Idx = startIdx

	if len(args) > 0 && args[0] == 1 && result.Tp != ERR {
		return &Token{BLOCK, resultBlk}, nil
	}

	return result, nil

}

func (es *EvalStack) EvalExp(ctx *BindMap) (*Token, error){
	var temp *Token
	var err error
	var isReturn = false
	// fmt.Println( es.Line[es.LastStartPos()].OutputStr())
	// for i := es.LastStartPos(); i <= es.LastEndPos(); i++{
	// 	fmt.Println(es.Line[i].OutputStr())
	// }
	
	var startPos = es.LastStartPos()
	var endPos = es.LastEndPos()-1
	var startToken = es.Line[startPos]

	// if startToken == nil {
	// 	EchoTokens(es.Line[es.LastStartPos():es.LastEndPos()+1])
	// }

	switch startToken.Tp {
	case SET_WORD:
		if es.IsLocal {
			ctx.PutLocal(startToken.Str(), es.Line[endPos])
		}else{
			ctx.Put(startToken.Str(), es.Line[endPos])
		}
		temp = es.Line[endPos]
	case PUT_WORD:
		ctx.PutLocal(startToken.Str(), es.Line[endPos])
		temp = es.Line[endPos]
	case PATH:
		if startToken.Tks()[0].Tp == FUNC {
			temp, err = startToken.Tks()[0].Val.(Func).RunWithProps(es, ctx, startToken.Tks())
		}else{
			temp, err = startToken.SetPathVal(es.Line[endPos], ctx, es)
		}
	case NATIVE, OP:
		temp, err = startToken.Val.(Native).Exec(es, ctx)
	case FUNC:
		temp, err = startToken.Val.(Func).Run(es, ctx)
	default:
		
	}

	if temp != nil && temp.Tp == ERR {
		return temp, err
	}

	if err != nil {
		if err.Error() == "return"{
			isReturn = true
			if startToken.Tp == FUNC || (startToken.Tp == PATH && startToken.Tks()[0].Tp == FUNC) {
				es.Line[startPos] = temp
				es.Idx = startPos + 1
				es.StartPos.Pop() 
				es.EndPos.Pop() 
				return temp, nil
			}
		}
		
	}

	if(!isReturn){
		es.Line[startPos] = temp
		es.Idx = startPos + 1
	}	

	es.StartPos.Pop() 
	es.EndPos.Pop() 

	
	return temp, err
}


func getStackErrInfo(es *EvalStack, inp []*Token, idx int, t *Token) *Token{
	t.Val = t.Str() + "\nNear: "

	for i := 5; i >= 0 ; i-- {
		if idx - i > 0 {
			t.Val = t.Str() + inp[idx-i].ToString() + "   "
		}
	}

	t.Val = t.Str() + "\nCall-by: "

	for i := 0; i < 5 && es.StartPos.Len() - i > 0; i++ {
		t.Val = t.Str() + es.Line[es.StartPos.Get(es.StartPos.Len()-i-1)].ToString() + "   "
	}

	return t
}
