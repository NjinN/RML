package core

import "unicode"
import "strings"
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

	var runes = []rune(str)

	var startIdx = -1
	var isParen = false
	var isStr = false
	var isBlock = false
	var isObject = false
	var isFile = false
	var pFloor = 0
	var bFloor = 0
	var oFloor = 0
	var nowRune rune
	var isInnerStr = false

	for nowIdx := 0; nowIdx < len(runes); nowIdx++ {
		nowRune= runes[nowIdx]

		if(nowIdx == len(runes)-1){
			if(startIdx < 0 && !unicode.IsSpace(runes[nowIdx])){
				result = append(result, string(nowRune))
				break
			}

			if(startIdx >= 0){
				if(unicode.IsSpace(nowRune)){
					result = append(result, string(runes[startIdx : nowIdx]))
					break
				}else{
					if(!isStr && !isParen && !isBlock){
						result = append(result, string(runes[startIdx : nowIdx+1]))
						break
					}
				}
			}
		}

		if(startIdx < 0 && !unicode.IsSpace(nowRune)){
			if(nowRune == rune('"')){
				isStr = true
			}else if(nowRune == rune('(')){
				isParen = true
				pFloor = 1
			}else if(nowRune == rune('[')){
				isBlock = true
				bFloor = 1
			}else if(nowRune == rune('{')){
				isObject = true
				oFloor = 1
			}else if (nowRune == rune('%')){
				isFile = true
			}else if(nowRune == rune(';')){
				for nowRune != rune('\n') && nowIdx < len(runes)-1 {
					nowIdx++
					nowRune = runes[nowIdx]
				}
				continue
			}
			startIdx = nowIdx
			continue
		}

		if startIdx >= 0 && runes[startIdx] == rune('!') && nowRune == rune('{') {
			isObject = true
			oFloor = 1
			continue
		}

		if(startIdx >= 0 && unicode.IsSpace(nowRune) && !isStr && !isParen && !isBlock && !isObject && !isFile){
			result = append(result, string(runes[startIdx : nowIdx]))
			startIdx = -1
			continue
		}

		if(startIdx >= 0 && isStr){
			if nowRune == rune('^') {
				nowIdx++
				continue
			}
			if nowRune == rune('"'){
				result = append(result, string(runes[startIdx : nowIdx+1]))
				isStr = false
				startIdx = -1
				continue
			}
			continue
		}

		if(startIdx >= 0 && isParen){
			if(isInnerStr){
				if nowRune == rune('^') {
					nowIdx++
					continue
				}
				if nowRune == rune('"'){
					isInnerStr = false
				}
			}else{
				if nowRune == '"' && !(string(runes[nowIdx-1 : nowIdx+1]) == "^\"") {
					isInnerStr = true
				}else if nowRune == rune('('){
					pFloor += 1
				}else if nowRune == rune(')'){
					pFloor -= 1
				}
				if(pFloor == 0){
					result = append(result, string(runes[startIdx : nowIdx+1]))
					isParen = false
					startIdx = -1
					continue
				}
			}
			continue
		}

		if(startIdx >= 0 && isBlock){
			if(isInnerStr){
				if nowRune == rune('^') {
					nowIdx++
					continue
				}
				if nowRune == rune('"'){
					isInnerStr = false
				}
			}else{
				if nowRune == '"' && !(string(runes[nowIdx-1 : nowIdx+1]) == "^\""){
					isInnerStr = true
				}else if nowRune == rune('['){
					bFloor += 1
				}else if nowRune == rune(']'){
					bFloor -= 1
				}
				if(bFloor == 0){
					result = append(result, string(runes[startIdx : nowIdx+1]))
					isBlock = false
					startIdx = -1
					continue
				}
			}
			continue
		}

		if(startIdx >= 0 && isFile){
			if(isInnerStr){
				if nowRune == rune('^') {
					nowIdx++
					continue
				}
				if nowRune == rune('"'){
					isInnerStr = false
				}
			}else{
				if nowRune == '"' && !(string(runes[nowIdx-1 : nowIdx+1]) == "^\""){
					isInnerStr = true
				}else if unicode.IsSpace(nowRune){
					result = append(result, string(runes[startIdx : nowIdx+1]))
					isFile = false
					startIdx = -1
					continue
				}
			}
			continue
		}

		if(startIdx >= 0 && isObject){
			if(isInnerStr){
				if nowRune == rune('^') {
					nowIdx++
					continue
				}
				if nowRune == rune('"'){
					isInnerStr = false
				}
			}else{
				if nowRune == rune('"') && !(string(runes[nowIdx-1 : nowIdx+1]) == "^\"") {
					isInnerStr = true
				}else if nowRune == rune('{'){
					oFloor += 1
				}else if nowRune == rune('}'){
					oFloor -= 1
				}
				if(oFloor == 0){
					result = append(result, string(runes[startIdx : nowIdx+1]))
					isObject = false
					startIdx = -1
					continue
				}
			}
			continue
		}

		if startIdx >= 0 && nowRune == rune('/') && nowIdx < len(runes) && !isStr && !isBlock && !isParen && !isObject && !isFile {
			for startIdx >= 0 && nowRune == rune('/') && nowIdx < len(runes)-1 {
				nowIdx++
				nowIdx += len(GetOneToken(runes, nowIdx)) - 1
				if(nowIdx < len(runes)-1){
					nowRune = runes[nowIdx + 1]
				}
			}
			result = append(result, string(runes[startIdx : nowIdx+1]))
			startIdx = -1
			continue
		}
		

	}
	return result
}

func GetOneToken(rs []rune, startIdx int) []rune{
	
	if rs[startIdx] == rune('"') {
		return GetSubStr(rs, startIdx)
	}else if rs[startIdx] == rune('(') {
		return GetSubParen(rs, startIdx)
	}else{
		return getSubOne(rs, startIdx)
	}

}

func getSubOne(rs []rune, startIdx int) []rune {
	for idx := startIdx+1; idx < len(rs); idx++ {
		if rs[idx] == rune('/') || unicode.IsSpace(rs[idx]) {
			return rs[startIdx:idx]
		}
	}
	return rs[startIdx:]
}

func GetSubStr(rs []rune, startIdx int) []rune{
	for idx := startIdx+1; idx < len(rs); idx++ {
		if rs[idx] == rune('"') && string(rs[idx-1:idx+1]) != "^\"" {
			return rs[startIdx:idx+1]
		}
	}
	return rs[startIdx:]
}

func GetSubParen(rs []rune, startIdx int) []rune{
	var pFloor = 1;
	for idx := startIdx+1; idx < len(rs); idx++ {
		if rs[idx] == rune('(') {
			pFloor++
		}else if rs[idx] == rune('"') {
			idx += len(GetSubStr(rs, idx)) + 1
		}else if rs[idx] == rune(')') {
			pFloor--
		}
		if pFloor == 0 {
			return rs[startIdx:idx+1]
		}
	}
	return rs[startIdx:]
}

func PathCut(s string) []string{
	var result []string
	var rs = []rune(s)
	for idx:=0; idx<len(rs); idx++ {
		var temp = GetOneToken(rs, idx)
		result = append(result, string(temp))
		idx += len(temp)
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

func StartWith(source string, target string) bool{
	if len(source) == 0 {
		return false
	}
	if len(target) == 0 {
		return true
	}
	if len(target) > len(source){
		return false
	}
	return source[0:len(target)] == target
}

func EndWith(source string, target string) bool{
	if len(source) == 0 {
		return false
	}
	if len(target) == 0 {
		return true
	}
	if len(target) > len(source){
		return false
	}
	return source[len(source) - len(target):] == target
}



func GetParentDir(dir string) string {

	var end = strings.LastIndexByte(dir, '/')
	if end == 0 {
		return "/"
	}else if end < 0 {
		return "./"
	}else{
		return dir[0:end]
	}

}


var charToCaretMap = map[int]string {
	0x00: "^@",
	0x01: "^A",
	0x02: "^B",
	0x03: "^C",
	0x04: "^D",
	0x05: "^E",
	0x06: "^F",
	0x07: "^G",
	0x08: "^H",
	0x09: "^-",
	0x0A: "^/",
	0x0B: "^K",
	0x0C: "^L",
	0x0D: "^M",
	0x0E: "^N",
	0x0F: "^O",
	0x10: "^P",
	0x11: "^Q",
	0x12: "^R",
	0x13: "^S",
	0x14: "^T",
	0x15: "^U",
	0x16: "^V",
	0x17: "^W",
	0x18: "^X",
	0x19: "^Y",
	0x1A: "^Z",
	0x1B: "^[",
	0x1C: "^\\",
	0x1D: "^]",
	0x1E: "^!",
	0x1F: "^_",
	0x7F: "^~",
}

var caretToCharMap = map[string]string {
	"^@": string(0x00),
	"^A": string(0x01),
	"^B": string(0x02),
	"^C": string(0x03),
	"^D": string(0x04),
	"^E": string(0x05),
	"^F": string(0x06),
	"^G": string(0x07),
	"^H": string(0x08),
	"^I": string(0x09),
	"^-": string(0x09),
	"^(tab)": string(0x09),
	"^J": string(0x0A),
	"^/": string(0x0A),
	"^K": string(0x0B),
	"^L": string(0x0C),
	"^M": string(0x0D),
	"^N": string(0x0E),
	"^O": string(0x0F),
	"^P": string(0x10),
	"^Q": string(0x11),
	"^R": string(0x12),
	"^S": string(0x13),
	"^T": string(0x14),
	"^U": string(0x15),
	"^V": string(0x16),
	"^W": string(0x17),
	"^X": string(0x18),
	"^Y": string(0x19),
	"^Z": string(0x1A),
	"^[": string(0x1B),
	"^\\": string(0x1C),
	"^]": string(0x1D),
	"^!": string(0x1E),
	"^_": string(0x1F),
	"^~": string(0x7F),
}

