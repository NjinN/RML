package core

import "errors"
// import "fmt"

type EvalStack struct {
	StartPos 		[]int
	EndPos 			[]int
	Line			[]*Token
	Idx 			int
	MainCtx 		*BindMap
	QuoteList		[]int
	IsLocal 		bool
}


func (es *EvalStack) Init(){
	es.StartPos = make([]int, 0)
	es.EndPos = make([]int, 0)
	es.Line = make([]*Token, 1024*1024)
	es.Idx = 0
}

func (es *EvalStack) Push(t *Token){
	es.Line[es.Idx] = t
	es.Idx += 1
}

func (es *EvalStack) LastStartPos() int{
	if(len(es.StartPos) <= 0){
		return -1
	}
	return es.StartPos[len(es.StartPos) - 1]
}

func (es *EvalStack) LastEndPos() int{
	if(len(es.EndPos) <= 0){
		return -1
	}
	return es.EndPos[len(es.EndPos) - 1]
}

func (es *EvalStack) EvalStr(inpStr string, ctx *BindMap, args ...int) (*Token, error){
	return es.Eval(ToTokens(inpStr, ctx, es), ctx, args...)
}
  

func (es *EvalStack) Eval(inp []*Token, ctx *BindMap, args ...int) (*Token, error){
	var result *Token
	var resultBlk []*Token

	if(len(inp) == 0){
		return result, nil
	}


	// fmt.Println("------  start eval -------")
	// for _, item := range inp {
	// 	fmt.Println(item.OutputStr())
	// }
	// fmt.Println("------  end eval -------")

	var startIdx = es.Idx
	var startDeep = len(es.EndPos)
	
	var i = 0
	for i < len(inp){
		var nowToken = inp[i]
		var nextToken *Token
		
		var skip = false
		if nowToken.Tp == GET_WORD {
			skip = true
		}

		var nextSkip = false
		if(i < len(inp) - 1){
			nextToken = inp[i+1]
			if nextToken.Tp == GET_WORD {
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


		if(nextToken != nil && nextToken.Tp == OP && (startDeep == 0 || es.Idx > es.EndPos[startDeep - 1]) && !nextSkip){
			if(len(es.StartPos) == 0 || es.Line[es.LastStartPos()].Tp != OP){
				es.StartPos = append(es.StartPos, es.Idx)
				es.Push(nextToken)
				temp, err := nowToken.GetVal(ctx, es)
				if err != nil {
					return temp, err
				}
				es.EndPos = append(es.EndPos, es.Idx + 1)
				es.Push(temp)
			}else if(len(es.StartPos) == 0 || es.Line[es.LastStartPos()].Tp == OP){
				temp, err := nowToken.GetVal(ctx, es)
				if err != nil {
					return temp, err
				}
				es.Push(temp)
				es.EvalExp(ctx)
				es.Push(es.Line[es.Idx - 1])
				es.Line[es.Idx - 2] = nextToken
				es.StartPos = append(es.StartPos, es.Idx - 2)
				es.EndPos = append(es.EndPos, es.Idx)
			}
			i += 1
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
			
			if(nowToken != nil && nowToken.Tp == ERR){
				return nowToken, nil
			}else if(nowToken != nil && nowToken.Tp == OP && !skip && es.Idx >= 1){
				if(es.Idx > startIdx){
					es.StartPos = append(es.StartPos, es.Idx - 1)
					es.EndPos = append(es.EndPos, es.Idx + 1)
					es.Push(es.Line[es.Idx - 1])
					es.Line[es.Idx - 2] = nowToken
				}else{
					result.Tp = ERR
					result.Val = "Illegal grammar!!!"
					return result, errors.New("Illegal grammar!!!")
				}
			}else if(nowToken != nil && nowToken.Tp < SET_WORD){
				es.Push(nowToken)
				if len(args) > 0 && args[0] == 1 && (len(es.StartPos) == 0 || es.Line[es.LastStartPos()].Tp != OP) {
					resultBlk = append(resultBlk, nowToken)
				}
			}else if nowToken.Tp == PATH && nowToken.Val.([]*Token)[0] != nil && nowToken.Val.([]*Token)[0].Tp == FUNC {
				es.StartPos = append(es.StartPos, es.Idx)
				es.EndPos = append(es.EndPos, es.Idx + nowToken.GetPathExpLen() - 1)
				es.Push(nowToken);
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
					es.StartPos = append(es.StartPos, es.Idx)
					es.EndPos = append(es.EndPos, es.Idx + nowToken.Explen() - 1)
				}
				
				es.Push(nowToken);
			}
		}
		
		for(len(es.EndPos) > startDeep && es.Idx == es.LastEndPos() + 1){
			temp, err := es.EvalExp(ctx)
			if err != nil {
				return temp, err
			}
			if len(args) > 0 && args[0] == 1 {
				resultBlk = append(resultBlk, temp)
			}
		}

		i += 1
	}
	result = es.Line[es.Idx - 1]
	es.Idx = startIdx

	if len(args) > 0 && args[0] == 1 {
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

	var startToken = es.Line[es.LastStartPos()]

	switch startToken.Tp {
	case SET_WORD:
		if es.IsLocal {
			ctx.PutLocal(es.Line[es.LastStartPos()].Val.(string), es.Line[es.LastEndPos()])
		}else{
			ctx.Put(es.Line[es.LastStartPos()].Val.(string), es.Line[es.LastEndPos()])
		}
		temp = es.Line[es.LastEndPos()]
	case PUT_WORD:
		ctx.PutLocal(es.Line[es.LastStartPos()].Val.(string), es.Line[es.LastEndPos()])
		temp = es.Line[es.LastEndPos()]
	case PATH:
		if startToken.Val.([]*Token)[0].Tp == FUNC {
			temp, err = startToken.Val.([]*Token)[0].Val.(Func).RunWithProps(es, ctx, startToken.Val.([]*Token))
		}else{
			startToken.SetPathVal(es.Line[es.LastEndPos()], ctx, es)
			temp = es.Line[es.LastEndPos()]
		}
	case NATIVE, OP:
		temp, err = es.Line[es.LastStartPos()].Val.(Native).Exec(es, ctx)
	case FUNC:
		temp, err = es.Line[es.LastStartPos()].Val.(Func).Run(es, ctx)
	default:
		
	}



	if err != nil {
		if err.Error() == "return"{
			isReturn = true
			if startToken.Tp == FUNC {
				es.Line[es.LastStartPos()] = temp
				es.Idx = es.LastStartPos() + 1
				es.StartPos = es.StartPos[0 : len(es.StartPos)-1]
				es.EndPos = es.EndPos[0 : len(es.EndPos)-1]
				return temp, nil
			}
		}
		
	}

	if(!isReturn){
		es.Line[es.LastStartPos()] = temp
		es.Idx = es.LastStartPos() + 1
	}	

	es.StartPos = es.StartPos[0 : len(es.StartPos)-1]
	es.EndPos = es.EndPos[0 : len(es.EndPos)-1]
	
	return temp, err
}