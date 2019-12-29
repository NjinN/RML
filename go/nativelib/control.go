package nativelib

import . "../core"
import "errors"
// import "fmt"

func Iif(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	var result Token
	if args[1].ToBool(){
		if args[2].Tp == BLOCK {
			return es.Eval(args[2].Val.([]*Token), ctx)
		}else if args[2].Tp == STRING {
			return es.EvalStr(args[2].Val.(string), ctx)
		}
	}else{
		return nil, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil

}

func Either(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	var result Token
	if args[1].ToBool(){
		if args[2].Tp == BLOCK {
			return es.Eval(args[2].Val.([]*Token), ctx)
		}else if args[2].Tp == STRING {
			return es.EvalStr(args[2].Val.(string), ctx)
		}
	}else{
		if args[3].Tp == BLOCK {
			return es.Eval(args[3].Val.([]*Token), ctx)
		}else if args[3].Tp == STRING {
			return es.EvalStr(args[3].Val.(string), ctx)
		}
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
			if rs != nil && rs.Tp == ERR {
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
		var countToken = args[2].Dup()
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
					if rs != nil && rs.Tp == ERR {
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
					if rs != nil && rs.Tp == ERR {
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
					if rs != nil && rs.Tp == ERR {
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
					if rs != nil && rs.Tp == ERR {
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
					if rs != nil && rs.Tp == ERR {
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
					if rs != nil && rs.Tp == ERR {
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
					if rs != nil && rs.Tp == ERR {
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
					if rs != nil && rs.Tp == ERR {
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
				if err.Error() == "continue" {
					continue
				}
				if err.Error() == "break" {
					break
				}
				return rs, err
			}
			if rs != nil && rs.Tp == ERR {
				return rs, err
			}
			b, err = es.Eval(args[1].Val.([]*Token), &c)
			if err != nil {
				return rs, err
			}
			if b.Tp == ERR {
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

func Fforeach(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	if args[3].Tp != BLOCK && args[3].Tp != STRING {
		var result = Token{ERR, "Type Mismatch"}
		return &result, nil
	}

	if args[1].Tp == WORD {
		if args[2].Tp == BLOCK || args[2].Tp == PAREN || args[2].Tp == PATH {
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX}
			for i:=0; i<len(args[2].Val.([]*Token)); i++ {
				c.PutNow(args[1].Val.(string), args[2].Val.([]*Token)[i])
				if args[3].Tp == BLOCK {
					temp, err := es.Eval(args[3].Val.([]*Token), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}

				}else if args[3].Tp == STRING {
					temp, err := es.EvalStr(args[3].Val.(string), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}
				}

			}
			return nil, nil

		}else if args[2].Tp == OBJECT {
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX}
			for k, v := range(args[2].Val.(*BindMap).Table){
				c.PutNow(args[1].Val.(string), &Token{BLOCK, []*Token{&Token{WORD, k}, v}})
				if args[3].Tp == BLOCK {
					temp, err := es.Eval(args[3].Val.([]*Token), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}

				}else if args[3].Tp == STRING {
					temp, err := es.EvalStr(args[3].Val.(string), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}
				}
			}
			return nil, nil
		}

	}else if args[1].Tp == BLOCK {
		for _, item := range(args[1].Val.([]*Token)) {
			if item.Tp != WORD {
				var result = Token{ERR, "Type Mismatch"}
				return &result, nil
			}
		}

		if args[2].Tp == BLOCK {
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX}
			for i:=0; i<len(args[2].Val.([]*Token)); i+=len(args[1].Val.([]*Token)){
				for j:=0; j<len(args[1].Val.([]*Token)); j++ {
					if i+j < len(args[2].Val.([]*Token)) {
						c.PutNow(args[1].Val.([]*Token)[j].Val.(string), args[2].Val.([]*Token)[i+j])
					}else{
						c.PutNow(args[1].Val.([]*Token)[j].Val.(string), &Token{NONE, "none"})
					}
				}
				temp, err := es.Eval(args[3].Val.([]*Token), &c)
				if err != nil {
					if err.Error() == "continue" {
						continue
					}
					if err.Error() == "break" {
						break
					}
					return temp, err
				}
				if temp != nil && temp.Tp == ERR {
					return temp, err
				}
			}
			
			return nil, nil
		}else if args[2].Tp == OBJECT {
			if len(args[1].Val.([]*Token)) < 2 || args[1].Val.([]*Token)[0].Tp != WORD || args[1].Val.([]*Token)[1].Tp != WORD {
				var result = Token{ERR, "Type Mismatch"}
				return &result, nil
			}
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX}
			for k, v := range(args[2].Val.(*BindMap).Table){
				c.PutNow(args[1].Val.([]*Token)[0].Val.(string), &Token{WORD, k})
				c.PutNow(args[1].Val.([]*Token)[1].Val.(string), v)
				temp, err := es.Eval(args[3].Val.([]*Token), &c)
				if err != nil {
					if err.Error() == "continue" {
						continue
					}
					if err.Error() == "break" {
						break
					}
					return temp, err
				}
				if temp != nil && temp.Tp == ERR {
					return temp, err
				}
			}
			return nil, nil
		}

	}



	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}

