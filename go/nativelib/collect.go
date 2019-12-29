package nativelib

import "strings"

import . "../core"
// import "fmt"

func Length(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1].Tp == BLOCK || args[1].Tp == PAREN {
		result.Tp = INTEGER
		result.Val = len(args[1].Tks())
		return &result, nil
	}else if args[1].Tp == STRING {
		result.Tp = INTEGER
		result.Val = len([]rune(args[1].Str()))
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Append(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token
	if args[1].Tp == BLOCK || args[1].Tp == PAREN {
		if !args[3].ToBool() && (args[2].Tp == BLOCK || args[2].Tp == PAREN) {
			args[1].Val = append(args[1].Tks(), args[2].Clone().Tks()...)
		}else{
			args[1].Val = append(args[1].Tks(), args[2].Clone())
		}
		return args[1].Clone(), nil
	}else if args[1].Tp == STRING {
		if args[2].Tp == STRING {
			if !args[3].ToBool(){
				args[1].Val = args[1].Str() + args[2].Str()
			}else{
				args[1].Val = args[1].Str() + `"` + args[2].Str() + `"`
			}
			return args[1].Clone(), nil
		}else if args[2].Tp == CHAR {
			if !args[3].ToBool(){
				args[1].Val = args[1].Str() + string(args[2].Val.(byte))
			}else{
				args[1].Val = args[1].Str() + `'` + string(args[2].Val.(byte)) + `'`
			}
			return args[1].Clone(), nil
		}else{
			args[1].Val = args[1].Str() + args[2].ToString()
			return args[1].Clone(), nil
		}
	}else if args[1].Tp == OBJECT {
		if args[2].Tp == BLOCK {
			for i := 0; i < len(args[2].Tks()) - 1; i+=2 {
				args[1].Ctx().Table[args[2].Tks()[i].ToString()] = args[2].Tks()[i+1]
			}
			return args[1].Clone(), nil 
		}else if args[2].Tp == OBJECT {
			for k, v := range(args[2].Ctx().Table){
				args[1].Ctx().Table[k] = v.Clone()
			}
			return args[1].Clone(), nil
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}


func Insert(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token
	
	if args[3].Tp != INTEGER {
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return &result, nil
	}else if args[3].Int() <= 0 {
		result.Tp = ERR
		result.Val = "Bound Overflow"
		return &result, nil
	}

	var idx = args[3].Int() - 1

	if args[1].Tp == BLOCK || args[1].Tp == PAREN {
		if !args[4].ToBool() && (args[2].Tp == BLOCK || args[2].Tp == PAREN) {
			if idx <= len(args[1].Tks()){
				temp := make([]*Token, len(args[1].Tks()) + len(args[2].Clone().Tks()))
				for i := 0; i<idx; i++ {
					temp[i] = args[1].Tks()[i]
				}
				for i:=0; i<len(args[2].Tks()); i++ {
					temp[idx+i] = args[2].Clone().Tks()[i]
				}
				for i:=idx; i<len(args[1].Tks()); i++ {
					temp[len(args[2].Clone().Tks()) + i] = args[1].Tks()[i]
				}
				args[1].Val = temp
				return args[1].Clone(), nil
			}else{
				for len(args[1].Tks()) < idx {
					args[1].Val = append(args[1].Tks(), &Token{NONE, "none"})
				}
				args[1].Val = append(args[1].Tks(), args[2].Clone().Tks()...)
				return args[1].Clone(), nil
			}
		}else{
			if idx <= len(args[1].Tks()){
				temp := make([]*Token, len(args[1].Tks()) + 1)
				for i := 0; i<idx; i++ {
					temp[i] = args[1].Tks()[i]
				}
				temp[idx] = args[2].Clone()
				for i:=idx; i<len(args[1].Tks()); i++ {
					temp[i+1] = args[1].Tks()[i]
				}
				args[1].Val = temp
				return args[1].Clone(), nil
			}else{
				for i:=len(args[1].Tks()); i<idx; i++ {
					args[1].Val = append(args[1].Tks(), &Token{NONE, "none"})
				}
				args[1].Val = append(args[1].Tks(), args[2].Clone())
				return args[1].Clone(), nil
			}
		}
		return args[1].Clone(), nil
	}else if args[1].Tp == STRING {
		if args[2].Tp == STRING {
			if !args[4].ToBool(){
				if idx <= len(args[1].Str()) {
					args[1].Val = args[1].Str()[0:idx] + args[2].Str() + args[1].Str()[idx:]
				}else{
					for len(args[1].Str()) < idx {
						args[1].Val = args[1].Str() + " "
					}
					args[1].Val = args[1].Str() + args[2].Str()
				}
				
			}else{
				if idx <= len(args[1].Str()) {
					args[1].Val = args[1].Str()[0:idx] + `"` + args[2].Str() + `"` + args[1].Str()[idx:]
				}else{
					for len(args[1].Str()) < idx {
						args[1].Val = args[1].Str() + " "
					}
					args[1].Val = args[1].Str() + `"` + args[2].Str() + `"`
				}
			}
			return args[1].Clone(), nil
		}else if args[2].Tp == CHAR {
			if !args[4].ToBool(){
				if idx <= len(args[1].Str()) {
					args[1].Val = args[1].Str()[0:idx] + string(args[2].Val.(byte)) + args[1].Str()[idx:]
				}else{
					for len(args[1].Str()) < idx {
						args[1].Val = args[1].Str() + " "
					}
					args[1].Val = args[1].Str() + string(args[2].Val.(byte))
				}
				
			}else{
				if idx <= len(args[1].Str()) {
					args[1].Val = args[1].Str()[0:idx] + `'` + string(args[2].Val.(byte)) + `'` + args[1].Str()[idx:]
				}else{
					for len(args[1].Str()) < idx {
						args[1].Val = args[1].Str() + " "
					}
					args[1].Val = args[1].Str() + `'` + string(args[2].Val.(byte)) + `'`
				}
			}
			return args[1].Clone(), nil
		}else{
			args[1].Val = args[1].Str() + args[2].ToString()
			return args[1].Clone(), nil
		}
	}else if args[1].Tp == OBJECT {
		if args[2].Tp == BLOCK {
			for i := 0; i < len(args[2].Tks()) - 1; i+=2 {
				args[1].Ctx().Table[args[2].Tks()[i].ToString()] = args[2].Tks()[i+1]
			}
			return args[1].Clone(), nil 
		}else if args[2].Tp == OBJECT {
			for k, v := range(args[2].Ctx().Table){
				args[1].Ctx().Table[k] = v.Clone()
			}
			return args[1].Clone(), nil
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Take(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1].Tp == BLOCK && args[2].Tp == INTEGER && args[3].Tp == INTEGER {
		if args[3].Int() > 1 {
			result.Tp = BLOCK
			result.Val = make([]*Token, 0  )
			var starIdx = args[2].Int() - 1
			var endIdx = starIdx + args[3].Int()

			if starIdx < 0 {
				starIdx = 0
			}
			if endIdx > len(args[1].Tks()){
				endIdx = len(args[1].Tks())
			}

			if starIdx < endIdx {
				result.Val = args[1].Tks()[starIdx:endIdx]
			}

			if args[4].Tp == LOGIC && args[4].Val.(bool){
				args[1].Val = append(args[1].Tks()[0:starIdx], args[1].Tks()[endIdx:]...)
			}
			return &result, nil

		}else if args[3].Int() == 1 {
			var idx = args[2].Int() - 1
			if idx < 0 || idx >= len(args[1].Tks()){
				result.Tp = NONE
				result.Val = "none"
			}else{
				result.Tp = args[1].Tks()[idx].Tp
				result.Val = args[1].Tks()[idx].Val
			}

			if args[4].Tp == LOGIC && args[4].Val.(bool){
				args[1].Val = append(args[1].Tks()[0:idx], args[1].Tks()[idx+1:]...)
			}
			return &result, nil
		}
	}else if args[1].Tp == STRING && args[2].Tp == INTEGER && args[3].Tp == INTEGER {
		if args[3].Int() > 1 {
			result.Tp = STRING
			result.Val = ""
			var starIdx = args[2].Int() - 1
			var endIdx = starIdx + args[3].Int()
			
			if starIdx < 0 {
				starIdx = 0
			}
			if endIdx > len(args[1].Str()){
				endIdx = len(args[1].Str())
			}
			if starIdx < endIdx {
				result.Val = args[1].Str()[starIdx:endIdx]
			}
			
			if args[4].Tp == LOGIC && args[4].Val.(bool){
				args[1].Val = args[1].Str()[0:starIdx] + args[1].Str()[endIdx:]
			}
			return &result, nil

		}else if args[3].Int() == 1 {
			var idx = args[2].Int() - 1
			if idx < 0 || idx >= len(args[1].Str()){
				result.Tp = NONE
				result.Val = "none"
			}else{
				result.Tp = CHAR
				result.Val = uint8(args[1].Str()[idx])
			}

			if args[4].Tp == LOGIC && args[4].Val.(bool){
				args[1].Val = args[1].Str()[0:idx] + args[1].Str()[idx+1:]
			}
			return &result, nil
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}


func Replace(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[4].Tp != INTEGER || args[5].Tp != INTEGER {
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return &result, nil
	}
	var at = args[4].Int() - 1
	var amount = args[5].Int()

	if args[1].Tp == STRING {
		var old = ""
		if args[2].Tp == STRING {
			old = args[2].Str()
		}else{
			old = args[2].ToString()
		}
		var new = ""
		if args[3].Tp == STRING {
			new = args[3].Str()
		}else{
			new = args[3].ToString()
		}
		
		args[1].Val = args[1].Str()[0:at] + strings.Replace(args[1].Str()[at:], old, new, amount)
		return args[1], nil
	}else if args[1].Tp == BLOCK || args[1].Tp == PAREN {
		if amount < 0 {
			amount = int(^uint(0) >> 1) //取有符号int最大值
		}

		if args[2].Tp == BLOCK || args[2].Tp == PAREN {
			for i := at; i < len(args[1].Tks()) - len(args[2].Tks()); i++ {
				var match = true
				for j := 0; j < len(args[2].Tks()); j++{
					if args[2].Tks()[j].Tp != args[1].Tks()[i+j].Tp || args[2].Tks()[j].Val != args[1].Tks()[i+j].Val {
						match = false
						break
					}
				}
				if match && amount > 0 {

					var temp = make([]*Token, 0)
					if args[3].Tp == BLOCK || args[3].Tp == PAREN {
						for n := 0; n < i; n++ {
							temp = append(temp, args[1].Tks()[n])
						}
						for n := 0; n < len(args[3].Tks()); n++ {
							temp = append(temp, args[3].Tks()[n])
						}
						for n := i + len(args[2].Tks()); n < len(args[1].Tks()); n++ {
							temp = append(temp, args[1].Tks()[n])
						}
					}else{
						for n := 0; n < i; n++ {
							temp = append(temp, args[1].Tks()[n])
						}
						
						temp = append(temp, args[3])
						
						for n := i + len(args[2].Tks()); n < len(args[1].Tks()); n++ {
							temp = append(temp, args[1].Tks()[n])
						}
					}
					
					args[1].Val = temp
				}
			}
			return args[1], nil

		}else{
			for i := at; i < len(args[1].Tks()) && amount > 0; i++ {
				var item = args[1].Tks()[i]
				if item.Tp == args[2].Tp && item.Val == args[2].Val {
					args[1] = args[3].Clone()
					amount--
				}
			}
			return args[1], nil
		}

	}



	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

