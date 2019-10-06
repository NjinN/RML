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
}


func (es *EvalStack) Init(){
	// es.StartPos = make([]int, 8)
	// es.EndPos = make([]int, 8)
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

func (es *EvalStack) EvalStr(inpStr string, ctx *BindMap) (*Token, error){
	return es.Eval(ToTokens(inpStr), ctx)
}

func (es *EvalStack) Eval(inp []*Token, ctx *BindMap) (*Token, error){
	var result *Token

	if(len(inp) == 0){
		return result, nil
	}

	var startIdx = es.Idx
	var startDeep = len(es.EndPos)
	
	var i = 0
	for i < len(inp){
		var nowToken = inp[i]
		var nextToken *Token
		if(i < len(inp) - 1){
			nextToken = inp[i+1]
			if(nextToken.Tp != WORD){
				nextToken = nextToken.GetVal(ctx, es)
			}
		}

		if(nextToken != nil && nextToken.Tp == OP && (startDeep == 0 || es.Idx > es.EndPos[startDeep - 1])){
			if(len(es.StartPos) == 0 || es.Line[es.LastStartPos()].Tp != OP){
				es.StartPos = append(es.StartPos, es.Idx)
				es.Push(nextToken)
				es.Push(nowToken.GetVal(ctx, es))
				es.EndPos = append(es.EndPos, es.Idx)
			}else if(len(es.StartPos) == 0 || es.Line[es.LastStartPos()].Tp == OP){
				es.Push(nowToken.GetVal(ctx, es))
				es.EvalExp(ctx)
				es.Push(es.Line[es.Idx - 1])
				es.Line[es.Idx - 2] = nextToken
				es.StartPos = append(es.StartPos, es.Idx)
			}
			i += 1
		}else{
			if(len(es.QuoteList) > 0){
				if(es.QuoteList[0] > 0){
					nowToken = nowToken.GetVal(ctx, es)
				}
				es.QuoteList = es.QuoteList[1 : len(es.QuoteList)]
			}else{
				nowToken = nowToken.GetVal(ctx, es)
			}

			if(nowToken != nil && nowToken.Tp == ERR){
				return nowToken, nil
			}else if(nowToken != nil && nowToken.Tp == OP){
				if(es.Idx > startIdx){
					es.StartPos = append(es.StartPos, es.Idx - 1)
					es.Push(es.Line[es.Idx - 1])
					es.Line[es.Idx - 2] = nowToken
					es.EndPos = append(es.EndPos, es.Idx)
				}else{
					result.Tp = ERR
					result.Val = "Illegal grammar!!!"
					return result, errors.New("Illegal grammar!!!")
				}
			}else if(nowToken != nil && nowToken.Tp < SET_WORD){
				es.Push(nowToken)
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

				es.StartPos = append(es.StartPos, es.Idx)
				es.EndPos = append(es.EndPos, es.Idx + nowToken.Explen() - 1)
				es.Push(nowToken);
			}
		}

		for(len(es.EndPos) > startIdx && es.Idx == es.LastEndPos() + 1){
			es.EvalExp(ctx)
		}

		i += 1
	}
	result = es.Line[es.Idx - 1]
	es.Idx = startIdx
	return result, nil

}

func (es *EvalStack) EvalExp(ctx *BindMap){
	var temp *Token
	var err error

	switch es.Line[es.LastStartPos()].Tp {
	case SET_WORD:
		ctx.Put(es.Line[es.LastStartPos()].Val.(string), es.Line[es.LastEndPos()])
		temp = es.Line[es.LastEndPos()]
	case NATIVE, OP:
		temp, err = es.Line[es.LastStartPos()].Val.(Native).Exec(es, ctx)
	case FUNC:
		temp, err = es.Line[es.LastStartPos()].Val.(Func).Run(es, ctx)
	default:
		
	}

	es.Line[es.LastStartPos()] = temp
	es.Idx = es.LastStartPos() + 1
	es.StartPos = es.StartPos[0 : len(es.StartPos)-1]
	es.EndPos = es.EndPos[0 : len(es.EndPos)-1]

	if err != nil {

	}

}