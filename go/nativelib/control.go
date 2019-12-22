package nativelib

import . "../core"
import "errors"

func Iif(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	var result Token
	switch args[1].Tp {
	case LOGIC:
		if args[1].Val.(bool) {
			return es.Eval(args[2].Val.([]*Token), ctx)
		}
	case INTEGER:
		if args[1].Val.(int) != 0 {
			return es.Eval(args[2].Val.([]*Token), ctx)
		}
	case DECIMAL:
		if args[1].Val.(float64) != 0.0 {
			return es.Eval(args[2].Val.([]*Token), ctx)
		}
	case STRING:
		if args[1].Val.(string) != "" {
			return es.Eval(args[2].Val.([]*Token), ctx)
		}
	case NONE:
		
	default:
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return &result, nil
	}

	return nil, nil
}

func Either(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	var result Token
	switch args[1].Tp {
	case LOGIC:
		if args[1].Val.(bool) {
			return es.Eval(args[2].Val.([]*Token), ctx)
		}else{
			return es.Eval(args[3].Val.([]*Token), ctx)
		}
	case INTEGER:
		if args[1].Val.(int) != 0 {
			return es.Eval(args[2].Val.([]*Token), ctx)
		}else{
			return es.Eval(args[3].Val.([]*Token), ctx)
		}
	case DECIMAL:
		if args[1].Val.(float64) != 0.0 {
			return es.Eval(args[2].Val.([]*Token), ctx)
		}else{
			return es.Eval(args[3].Val.([]*Token), ctx)
		}
	case STRING:
		if args[1].Val.(string) != "" {
			return es.Eval(args[2].Val.([]*Token), ctx)
		}else{
			return es.Eval(args[3].Val.([]*Token), ctx)
		}
	case NONE:
		return es.Eval(args[3].Val.([]*Token), ctx)
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Loop(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if(args[1].Tp == INTEGER && args[2].Tp == BLOCK){
		var rs *Token
		var err error
		for i := 0; i < args[1].Val.(int); i++ {
			rs, err = es.Eval(args[2].Val.([]*Token), ctx) 
			if err != nil {
				if err.Error() == "continue" {
					continue
				}
				if err.Error() == "break" {
					break
				}
				return rs, err
			}
		}

		return rs, nil

	}

	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}


func Repeat(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	// args[2].Echo()

	if(args[1].Tp == WORD && args[2].Tp == INTEGER && args[3].Tp == BLOCK){
		 var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX}
		 var countToken = Token{INTEGER, 1}
		 
		 c.PutNow(args[1].Val.(string), &countToken)
		 var rs *Token
		 var err error
		 for countToken.Val.(int) <= args[2].Val.(int) {
			rs, err = es.Eval(args[3].Val.([]*Token), &c)
			countToken.Val = countToken.Val.(int) + 1
			if err != nil {
				if err.Error() == "continue" {
					continue
				}
				if err.Error() == "break" {
					break
				}
				return rs, err
			}
		 }
		 return nil, nil

	}

	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}


func Ffor(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if(args[1].Tp == WORD && args[5].Tp == BLOCK && (args[2].Tp == INTEGER || args[2].Tp == DECIMAL) && (args[3].Tp == INTEGER || args[3].Tp == DECIMAL) && (args[4].Tp == INTEGER || args[4].Tp == DECIMAL)){
		var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX}
		var countToken = args[2].Clone()
		c.PutNow(args[1].Val.(string), countToken)
		var rs *Token
		var err error

		if(args[2].Tp == INTEGER && args[4].Tp == INTEGER){
			if args[3].Tp == INTEGER {
				for countToken.Val.(int) <= args[3].Val.(int) {
					rs, err = es.Eval(args[5].Val.([]*Token), &c)
					countToken.Val = countToken.Val.(int) + args[4].Val.(int)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
				}
			}else{
				for countToken.Val.(int) <= int(args[3].Val.(float64)) {
					rs, err = es.Eval(args[5].Val.([]*Token), &c)
					countToken.Val = countToken.Val.(int) + args[4].Val.(int)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
				}
			}

		}else if(args[2].Tp == INTEGER && args[4].Tp == DECIMAL) {
			countToken.Tp = DECIMAL
			countToken.Val = float64(countToken.Val.(int))
			if args[3].Tp == INTEGER {
				for countToken.Val.(float64) <= float64(args[3].Val.(int)) {
					rs, err = es.Eval(args[5].Val.([]*Token), &c)
					countToken.Val = countToken.Val.(float64) + args[4].Val.(float64)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
				}
			}else{
				for countToken.Val.(float64) <= args[3].Val.(float64) {
					rs, err = es.Eval(args[5].Val.([]*Token), &c)
					countToken.Val = countToken.Val.(float64) + args[4].Val.(float64)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
				}
			}
		}else if(args[2].Tp == DECIMAL && args[4].Tp == INTEGER) {
			if args[3].Tp == INTEGER {
				for countToken.Val.(float64) <= float64(args[3].Val.(int)) {
					rs, err = es.Eval(args[5].Val.([]*Token), &c)
					countToken.Val = countToken.Val.(float64) + float64(args[4].Val.(int))
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
				}
			}else{
				for countToken.Val.(float64) <= args[3].Val.(float64) {
					rs, err = es.Eval(args[5].Val.([]*Token), &c)
					countToken.Val = countToken.Val.(float64) + float64(args[4].Val.(int))
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
				}
			}
		}else if(args[2].Tp == DECIMAL && args[4].Tp == DECIMAL) {
			if args[3].Tp == INTEGER {
				for countToken.Val.(float64) <= float64(args[3].Val.(int)) {
					rs, err = es.Eval(args[5].Val.([]*Token), &c)
					countToken.Val = countToken.Val.(float64) + args[4].Val.(float64)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
				}
			}else{
				for countToken.Val.(float64) <= args[3].Val.(float64) {
					rs, err = es.Eval(args[5].Val.([]*Token), &c)
					countToken.Val = countToken.Val.(float64) + args[4].Val.(float64)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
				}
			}
		}

		return nil, nil
	}

	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}

func Wwhile(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if(args[1].Tp == BLOCK && args[2].Tp == BLOCK){
		var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX}
		b, e1 := es.Eval(args[1].Val.([]*Token), &c)
		if e1 != nil {
			return nil, e1
		}
		var rs *Token
		var err error
		for b.Val.(bool) {
			rs, err = es.Eval(args[2].Val.([]*Token), &c)
			if err != nil {
				return rs, err
			}
			b, err = es.Eval(args[1].Val.([]*Token), &c)
			if err != nil {
				return rs, err
			}
		}
	}
	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}


func Bbreak(es *EvalStack, ctx *BindMap) (*Token, error){
	return nil, errors.New("break")
}

func Ccontinue(es *EvalStack, ctx *BindMap) (*Token, error){
	return nil, errors.New("continue")
}

func Rreturn(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	return args[1], errors.New("return")
}
