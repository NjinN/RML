package core

import "strings"
import "strconv"

func ToToken(s string) *Token{
	var result Token
	var str = Trim(s)

	if(strings.ToLower(str) == "true"){
		result.Tp = LOGIC
		result.Val = true
		return &result
	}

	if(strings.ToLower(str) == "false"){
		result.Tp = LOGIC
		result.Val = false
		return &result
	}

	if(len(str) == 4 && str[0 : 2] == "#'" && str[3] == '\''){
		result.Tp = CHAR
		result.Val = str[2]
		return &result
	}

	if(str[0] == '"'){
		result.Tp = STRING
		result.Val = str[1 : len(str)-1]
		return &result
	}

	if(str[0] == '['){
		result.Tp = BLOCK
		var endIdx int
		for endIdx=len(str)-1; endIdx>=0; endIdx-- {
			if(str[endIdx] == ']'){
				break
			}
		}
		result.Val = ToTokens(str[1 : endIdx])
		return &result
	}

	if(str[0] == '('){
		result.Tp = BLOCK
		var endIdx int
		for endIdx=len(str)-1; endIdx>=0; endIdx-- {
			if(str[endIdx] == ')'){
				break
			}
		}
		result.Val = ToTokens(str[1 : endIdx])
		return &result
	}

	if(IsNumberStr(str) == 0){
		result.Tp = INTEGER
		i, err := strconv.Atoi(str)

		if err != nil {
			panic(err)
		}else{
			result.Val = i
		}
		return &result
	}

	if(IsNumberStr(str) == 1){
		result.Tp = DECIMAL
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			panic(err)
		}else{
			result.Val = f
		}
		return &result
	}

	if(str[len(str)-1] == ':'){
		result.Tp = SET_WORD
		result.Val = str[0 : len(str)-1]
		return &result
	}

	if(str[0] == '\''){
		result.Tp = LIT_WORD
		result.Val = str[1 : len(str)]
		return &result
	}

	result.Tp = WORD
	result.Val = str
	return &result
}


func ToTokens(str string) []*Token{
	var result []*Token
	var strs = StrCut(str)

	for _, item := range strs {
		result = append(result, ToToken(item))
	}
	return result
}

