package core

import "unicode"
// import "fmt"

func Trim(s string) string{
	var str = []rune(s)
	var startIdx = 0
	var endIdx = len(str) - 1

	for startIdx < len(s) {
		if(!unicode.IsSpace(str[startIdx])){
			break
		}
		startIdx++
	}

	for endIdx > 0 {
		if(!unicode.IsSpace(str[endIdx])){
			break
		}
		endIdx--
	}

	if(endIdx < startIdx){
		return ""
	}else{
		return string(str[startIdx : endIdx+1])
	}
}

func IsWhite(ch byte) bool{
	return unicode.IsSpace(rune(ch))
}


func StrCut(s string) []string{
	// fmt.Println(s)
	var result []string

	var str = Trim(s)

	if(str == ""){
		return result
	}

	var startIdx = -1
	var isParen = false
	var isStr = false
	var isBlock = false
	var pFloor = 0
	var bFloor = 0
	var nowChar byte
	var isInnerStr = false

	for nowIdx := 0; nowIdx < len(str); nowIdx++ {
		nowChar = str[nowIdx]

		if(nowIdx == len(str)-1){
			if(startIdx < 0 && !IsWhite(nowChar)){
				result = append(result, string(nowChar))
				break
			}

			if(startIdx >= 0){
				if(IsWhite(nowChar)){
					result = append(result, str[startIdx : nowIdx])
					break
				}else{
					if(!isStr && !isParen && !isBlock){
						result = append(result, str[startIdx : nowIdx+1])
						break
					}
				}
			}
		}

		if(startIdx < 0 && !IsWhite(nowChar)){
			if(nowChar == '"'){
				isStr = true
			}else if(nowChar == '('){
				isParen = true
				pFloor = 1
			}else if(nowChar == '['){
				isBlock = true
				bFloor = 1
			}
			startIdx = nowIdx
			continue;
		}

		if(startIdx >= 0 && IsWhite(nowChar) && !isStr && !isParen && !isBlock){
			result = append(result, str[startIdx : nowIdx])
			startIdx = -1
			continue
		}

		if(startIdx >= 0 && isStr){
			if(nowChar == '"' && !(str[nowIdx-1 : nowIdx+1] == "^\"")){
				result = append(result, str[startIdx : nowIdx+1])
				isStr = false
				startIdx = -1
				continue
			}
		}

		if(startIdx >= 0 && isParen){
			if(isInnerStr){
				if(nowChar == '"' && !(str[nowIdx-1 : nowIdx+1] == "^\"")){
					isInnerStr = false
				}
			}else{
				if(nowChar == '"' && !(str[nowIdx-1 : nowIdx+1] == "^\"")){
					isInnerStr = true
				}else if(nowChar == '('){
					pFloor += 1
				}else if(nowChar == ')'){
					pFloor -= 1
				}
				if(pFloor == 0){
					result = append(result, str[startIdx : nowIdx+1])
					isParen = false
					startIdx = -1
					continue
				}
			}
		}

		if(startIdx >= 0 && isBlock){
			if(isInnerStr){
				if(nowChar == '"' && !(str[nowIdx-1 : nowIdx+1] == "^\"")){
					isInnerStr = false
				}
			}else{
				if(nowChar == '"' && !(str[nowIdx-1 : nowIdx+1] == "^\"")){
					isInnerStr = true
				}else if(nowChar == '['){
					bFloor += 1
				}else if(nowChar == ']'){
					bFloor -= 1
				}
				if(bFloor == 0){
					result = append(result, str[startIdx : nowIdx+1])
					isBlock = false
					startIdx = -1
					continue
				}
			}
		}
	}
	return result
}


func IsNumber(c byte) bool{
	if(c >= 48 && c <= 57){
		return true
	}
	return false
}

func IsNumberStr(s string) int{
	if(len(s) == 0){
		return -1
	}
	if(s[0] != '-' && !IsNumber(s[0]) || s== "-"){
		return -1
	}

	var dot = 0
	for idx:=1; idx<len(s); idx++ {
		if(!IsNumber(s[idx]) && s[idx] != '.'){
			return -1
		}
		if(s[idx] == '.'){
			dot += 1
		}
	}
	return dot
}




