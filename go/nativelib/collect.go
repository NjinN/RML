package nativelib

import "strings"

import . "../core"
// import "fmt"

func Length(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1].Tp == BLOCK || args[1].Tp == PAREN {
		result.Tp = INTEGER
		result.Val = len(args[1].Val.([]*Token))
		return &result, nil
	}else if args[1].Tp == STRING {
		result.Tp = INTEGER
		result.Val = len(args[1].Val.(string))
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
			args[1].Val = append(args[1].Val.([]*Token), args[2].Clone().Val.([]*Token)...)
		}else{
			args[1].Val = append(args[1].Val.([]*Token), args[2].Clone())
		}
		return args[1].Clone(), nil
	}else if args[1].Tp == STRING {
		if args[2].Tp == STRING {
			if !args[3].ToBool(){
				args[1].Val = args[1].Val.(string) + args[2].Val.(string)
			}else{
				args[1].Val = args[1].Val.(string) + `"` + args[2].Val.(string) + `"`
			}
			return args[1].Clone(), nil
		}else if args[2].Tp == CHAR {
			if !args[3].ToBool(){
				args[1].Val = args[1].Val.(string) + string(args[2].Val.(byte))
			}else{
				args[1].Val = args[1].Val.(string) + `'` + string(args[2].Val.(byte)) + `'`
			}
			return args[1].Clone(), nil
		}else{
			args[1].Val = args[1].Val.(string) + args[2].ToString()
			return args[1].Clone(), nil
		}
	}else if args[1].Tp == OBJECT {
		if args[2].Tp == BLOCK {
			for i := 0; i < len(args[2].Val.([]*Token)) - 1; i+=2 {
				args[1].Val.(*BindMap).Table[args[2].Val.([]*Token)[i].ToString()] = args[2].Val.([]*Token)[i+1]
			}
			return args[1].Clone(), nil 
		}else if args[2].Tp == OBJECT {
			for k, v := range(args[2].Val.(*BindMap).Table){
				args[1].Val.(*BindMap).Table[k] = v.Clone()
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
	}else if args[3].Val.(int) <= 0 {
		result.Tp = ERR
		result.Val = "Bound Overflow"
		return &result, nil
	}

	var idx = args[3].Val.(int) - 1

	if args[1].Tp == BLOCK || args[1].Tp == PAREN {
		if !args[4].ToBool() && (args[2].Tp == BLOCK || args[2].Tp == PAREN) {
			if idx <= len(args[1].Val.([]*Token)){
				temp := make([]*Token, len(args[1].Val.([]*Token)) + len(args[2].Clone().Val.([]*Token)))
				for i := 0; i<idx; i++ {
					temp[i] = args[1].Val.([]*Token)[i]
				}
				for i:=0; i<len(args[2].Val.([]*Token)); i++ {
					temp[idx+i] = args[2].Clone().Val.([]*Token)[i]
				}
				for i:=idx; i<len(args[1].Val.([]*Token)); i++ {
					temp[len(args[2].Clone().Val.([]*Token)) + i] = args[1].Val.([]*Token)[i]
				}
				args[1].Val = temp
				return args[1].Clone(), nil
			}else{
				for len(args[1].Val.([]*Token)) < idx {
					args[1].Val = append(args[1].Val.([]*Token), &Token{NONE, "none"})
				}
				args[1].Val = append(args[1].Val.([]*Token), args[2].Clone().Val.([]*Token)...)
				return args[1].Clone(), nil
			}
		}else{
			if idx <= len(args[1].Val.([]*Token)){
				temp := make([]*Token, len(args[1].Val.([]*Token)) + 1)
				for i := 0; i<idx; i++ {
					temp[i] = args[1].Val.([]*Token)[i]
				}
				temp[idx] = args[2].Clone()
				for i:=idx; i<len(args[1].Val.([]*Token)); i++ {
					temp[i+1] = args[1].Val.([]*Token)[i]
				}
				args[1].Val = temp
				return args[1].Clone(), nil
			}else{
				for i:=len(args[1].Val.([]*Token)); i<idx; i++ {
					args[1].Val = append(args[1].Val.([]*Token), &Token{NONE, "none"})
				}
				args[1].Val = append(args[1].Val.([]*Token), args[2].Clone())
				return args[1].Clone(), nil
			}
		}
		return args[1].Clone(), nil
	}else if args[1].Tp == STRING {
		if args[2].Tp == STRING {
			if !args[4].ToBool(){
				if idx <= len(args[1].Val.(string)) {
					args[1].Val = args[1].Val.(string)[0:idx] + args[2].Val.(string) + args[1].Val.(string)[idx:]
				}else{
					for len(args[1].Val.(string)) < idx {
						args[1].Val = args[1].Val.(string) + " "
					}
					args[1].Val = args[1].Val.(string) + args[2].Val.(string)
				}
				
			}else{
				if idx <= len(args[1].Val.(string)) {
					args[1].Val = args[1].Val.(string)[0:idx] + `"` + args[2].Val.(string) + `"` + args[1].Val.(string)[idx:]
				}else{
					for len(args[1].Val.(string)) < idx {
						args[1].Val = args[1].Val.(string) + " "
					}
					args[1].Val = args[1].Val.(string) + `"` + args[2].Val.(string) + `"`
				}
			}
			return args[1].Clone(), nil
		}else if args[2].Tp == CHAR {
			if !args[4].ToBool(){
				if idx <= len(args[1].Val.(string)) {
					args[1].Val = args[1].Val.(string)[0:idx] + string(args[2].Val.(byte)) + args[1].Val.(string)[idx:]
				}else{
					for len(args[1].Val.(string)) < idx {
						args[1].Val = args[1].Val.(string) + " "
					}
					args[1].Val = args[1].Val.(string) + string(args[2].Val.(byte))
				}
				
			}else{
				if idx <= len(args[1].Val.(string)) {
					args[1].Val = args[1].Val.(string)[0:idx] + `'` + string(args[2].Val.(byte)) + `'` + args[1].Val.(string)[idx:]
				}else{
					for len(args[1].Val.(string)) < idx {
						args[1].Val = args[1].Val.(string) + " "
					}
					args[1].Val = args[1].Val.(string) + `'` + string(args[2].Val.(byte)) + `'`
				}
			}
			return args[1].Clone(), nil
		}else{
			args[1].Val = args[1].Val.(string) + args[2].ToString()
			return args[1].Clone(), nil
		}
	}else if args[1].Tp == OBJECT {
		if args[2].Tp == BLOCK {
			for i := 0; i < len(args[2].Val.([]*Token)) - 1; i+=2 {
				args[1].Val.(*BindMap).Table[args[2].Val.([]*Token)[i].ToString()] = args[2].Val.([]*Token)[i+1]
			}
			return args[1].Clone(), nil 
		}else if args[2].Tp == OBJECT {
			for k, v := range(args[2].Val.(*BindMap).Table){
				args[1].Val.(*BindMap).Table[k] = v.Clone()
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
		if args[3].Val.(int) > 1 {
			result.Tp = BLOCK
			result.Val = make([]*Token, 0  )
			var starIdx = args[2].Val.(int) - 1
			var endIdx = starIdx + args[3].Val.(int)

			if starIdx < 0 {
				starIdx = 0
			}
			if endIdx > len(args[1].Val.([]*Token)){
				endIdx = len(args[1].Val.([]*Token))
			}

			if starIdx < endIdx {
				result.Val = args[1].Val.([]*Token)[starIdx:endIdx]
			}

			if args[4].Tp == LOGIC && args[4].Val.(bool){
				args[1].Val = append(args[1].Val.([]*Token)[0:starIdx], args[1].Val.([]*Token)[endIdx:]...)
			}
			return &result, nil

		}else if args[3].Val.(int) == 1 {
			var idx = args[2].Val.(int) - 1
			if idx < 0 || idx >= len(args[1].Val.([]*Token)){
				result.Tp = NONE
				result.Val = "none"
			}else{
				result.Tp = args[1].Val.([]*Token)[idx].Tp
				result.Val = args[1].Val.([]*Token)[idx].Val
			}

			if args[4].Tp == LOGIC && args[4].Val.(bool){
				args[1].Val = append(args[1].Val.([]*Token)[0:idx], args[1].Val.([]*Token)[idx+1:]...)
			}
			return &result, nil
		}
	}else if args[1].Tp == STRING && args[2].Tp == INTEGER && args[3].Tp == INTEGER {
		if args[3].Val.(int) > 1 {
			result.Tp = STRING
			result.Val = ""
			var starIdx = args[2].Val.(int) - 1
			var endIdx = starIdx + args[3].Val.(int)
			
			if starIdx < 0 {
				starIdx = 0
			}
			if endIdx > len(args[1].Val.(string)){
				endIdx = len(args[1].Val.(string))
			}
			if starIdx < endIdx {
				result.Val = args[1].Val.(string)[starIdx:endIdx]
			}
			
			if args[4].Tp == LOGIC && args[4].Val.(bool){
				args[1].Val = args[1].Val.(string)[0:starIdx] + args[1].Val.(string)[endIdx:]
			}
			return &result, nil

		}else if args[3].Val.(int) == 1 {
			var idx = args[2].Val.(int) - 1
			if idx < 0 || idx >= len(args[1].Val.(string)){
				result.Tp = NONE
				result.Val = "none"
			}else{
				result.Tp = CHAR
				result.Val = uint8(args[1].Val.(string)[idx])
			}

			if args[4].Tp == LOGIC && args[4].Val.(bool){
				args[1].Val = args[1].Val.(string)[0:idx] + args[1].Val.(string)[idx+1:]
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
	var at = args[4].Val.(int) - 1
	var amount = args[5].Val.(int)

	if args[1].Tp == STRING {
		var old = ""
		if args[2].Tp == STRING {
			old = args[2].Val.(string)
		}else{
			old = args[2].ToString()
		}
		var new = ""
		if args[3].Tp == STRING {
			new = args[3].Val.(string)
		}else{
			new = args[3].ToString()
		}
		
		args[1].Val = args[1].Val.(string)[0:at] + strings.Replace(args[1].Val.(string)[at:], old, new, amount)
		return args[1], nil
	}else if args[1].Tp == BLOCK || args[1].Tp == PAREN {
		if amount < 0 {
			amount = int(^uint(0) >> 1) //取有符号int最大值
		}

		if args[2].Tp == BLOCK || args[2].Tp == PAREN {
			for i := at; i < len(args[1].Val.([]*Token)) - len(args[2].Val.([]*Token)); i++ {
				var match = true
				for j := 0; j < len(args[2].Val.([]*Token)); j++{
					if args[2].Val.([]*Token)[j].Tp != args[1].Val.([]*Token)[i+j].Tp || args[2].Val.([]*Token)[j].Val != args[1].Val.([]*Token)[i+j].Val {
						match = false
						break
					}
				}
				if match && amount > 0 {

					var temp = make([]*Token, 0)
					if args[3].Tp == BLOCK || args[3].Tp == PAREN {
						for n := 0; n < i; n++ {
							temp = append(temp, args[1].Val.([]*Token)[n])
						}
						for n := 0; n < len(args[3].Val.([]*Token)); n++ {
							temp = append(temp, args[3].Val.([]*Token)[n])
						}
						for n := i + len(args[2].Val.([]*Token)); n < len(args[1].Val.([]*Token)); n++ {
							temp = append(temp, args[1].Val.([]*Token)[n])
						}
					}else{
						for n := 0; n < i; n++ {
							temp = append(temp, args[1].Val.([]*Token)[n])
						}
						
						temp = append(temp, args[3])
						
						for n := i + len(args[2].Val.([]*Token)); n < len(args[1].Val.([]*Token)); n++ {
							temp = append(temp, args[1].Val.([]*Token)[n])
						}
					}
					
					args[1].Val = temp
				}
			}
			return args[1], nil

		}else{
			for i := at; i < len(args[1].Val.([]*Token)) && amount > 0; i++ {
				var item = args[1].Val.([]*Token)[i]
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

