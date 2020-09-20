package nativelib

import (
	"strconv"
	"strings"
	"bytes"
	"sync"
	// "fmt"
	"unicode"

	. "github.com/NjinN/RML/go/core"
)

func Jjson(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	return &Token{STRING, ToJsonStr(args[1])}, nil
}

func Dejson(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	if args[1].Tp == STRING{
		return DecodeJson(args[1].Str()), nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}



func ToJsonStr(t *Token) string {
	switch t.Tp {
	case NIL:
		return "null"
	
	case LOGIC:
		return strconv.FormatBool(t.Val.(bool))

	case INTEGER:
		return strconv.Itoa(t.Int())

	case DECIMAL:
		var result = strconv.FormatFloat(t.Float(), 'f', -1, 64)
		if strings.IndexByte(result, '.') < 0 {
			result += ".0"
		}
		return result

	case CHAR:
		return "\"" + string(t.Val.(rune)) + "\""

	case STRING, FILE, URL:
		return "\"" + t.Str() + "\""

	case TIME:
		v := t.Time()

		if v.Date == 0 {
			if v.FloatSecond > 0{
				return "\"" + SecsToTimeStr(v.Second) + strconv.FormatFloat(v.FloatSecond, 'f', -1, 64)[1:] + "\""
			}else{
				return "\"" + SecsToTimeStr(v.Second) + "\""
			}
			
		}else{
			if v.FloatSecond > 0 {
				return "\"" + DaysToDate(v.Date) + "+" + SecsToTimeStr(v.Second) + strconv.FormatFloat(v.FloatSecond, 'f', -1, 64)[1:] + "\""
			}else{
				return "\"" + DaysToDate(v.Date) + "+" + SecsToTimeStr(v.Second) + "\""
			}
			
		}
		
	case BLOCK:
		var buffer bytes.Buffer
		buffer.WriteString("[")
		for _, item := range t.Tks(){
			buffer.WriteString(ToJsonStr(item))
			buffer.WriteString(",")
		}
		if len(buffer.Bytes()) > 1 {
			buffer.Bytes()[len(buffer.Bytes())-1] = ']'
		}else{
			buffer.WriteString("]")
		}
		return buffer.String()

	case OBJECT:
		var buffer bytes.Buffer
		buffer.WriteString("{")
		for k, v := range t.Ctx().Table {
			buffer.WriteString("\"" + k + "\"")
			buffer.WriteString(":")
			buffer.WriteString(v.ToString())
			buffer.WriteString(",")
		}
		if len(buffer.Bytes()) > 1 {
			buffer.Bytes()[len(buffer.Bytes())-1] = '}'
		}else{
			buffer.WriteString("}")
		}
		return buffer.String()

	case MAP:
		var buffer bytes.Buffer
		buffer.WriteString("{")
		for _, v := range t.Map().Table {
			buffer.WriteString("\"")
			buffer.WriteString(v.Key.ToString())
			buffer.WriteString("\":")
			buffer.WriteString(ToJsonStr(v.Val))
			buffer.WriteString(",")
		}
		if len(buffer.Bytes()) > 1 {
			buffer.Bytes()[len(buffer.Bytes())-1] = '}'
		}else{
			buffer.WriteString("}")
		}
		
		return buffer.String()


	default:
		return "null"
	}

}


func DecodeJson(s string) *Token {
	var result Token
	// fmt.Println(s)
	var str = Trim(s)

	if str == "null" {
		return &Token{NONE, ""}
	}

	if strings.ToLower(str) == "true" {
		result.Tp = LOGIC
		result.Val = true
		return &result
	}

	if strings.ToLower(str) == "false" {
		result.Tp = LOGIC
		result.Val = false
		return &result
	}

	if str[0] == '"' {
		result.Tp = STRING
		result.Val = str[1:len(str)-1]
		return &result
	}

	if IsNumberStr(str) == 0 {
		result.Tp = INTEGER
		i, err := strconv.Atoi(str)

		if err != nil {
			return &Token{ERR, err.Error()}
		}else{
			result.Val = i
		}
		return &result
	}

	if IsNumberStr(str) == 1 {
		result.Tp = DECIMAL
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return &Token{ERR, err.Error()}
		}else{
			result.Val = f
		}
		return &result
	}

	if str[0] == '[' {
		result.Tp = BLOCK
		result.Val = NewTks(8)

		arr := JsonStrCut(str[1:len(str)-1])
		for _, item := range arr {
			result.List().Add(DecodeJson(item))
		}

		return &result
	}

	if str[0] == '{' {
		result.Tp = OBJECT
		var m = BindMap{make(map[string]*Token, 8), nil, USR_CTX, sync.RWMutex{}}

		arr := JsonStrCut(str[1:len(str)-1])

		for _, item := range arr {
			if strings.Index(item, ":") < 0 {
				return &Token{ERR, "Invalid JSON of " + str}
			}

			formated := strings.Replace(item, ":", ";;;;;;;;;;", 1)

			slices := strings.Split(formated, ";;;;;;;;;;")

			if len(slices) != 2 {
				return &Token{ERR, "Invalid JSON"}
			}

			m.PutNow(slices[0][1:len(slices[0])-1], DecodeJson(slices[1]))
		}
		
		
		result.Val = &m
		return &result
	}



	return &result
}


func JsonStrCut(s string) []string{
	// fmt.Println(s)
	var result []string

	var str = Trim(s)

	if(str == ""){
		return result
	}

	var runes = []rune(str)

	var startIdx = -1
	var isStr = false
	var isBlock = false
	var isObject = false
	var bFloor = 0
	var oFloor = 0
	var nowRune rune
	var isInnerStr = false

	for nowIdx := 0; nowIdx < len(runes); nowIdx++ {
		nowRune= runes[nowIdx]

		if(nowIdx == len(runes)-1){
			if(startIdx < 0 && !(runes[nowIdx] == rune(','))){
				result = append(result, string(nowRune))
				break
			}

			if(startIdx >= 0){
				if(runes[nowIdx] == rune(',')){
					result = append(result, string(runes[startIdx : nowIdx]))
					break
				}else{
					
					result = append(result, string(runes[startIdx : nowIdx+1]))
					break
					
				}
			}
		}

		if(startIdx < 0 && !unicode.IsSpace(runes[nowIdx]) && !(runes[nowIdx] == rune(','))){
			if(nowRune == rune('"')){
				isStr = true
			}else if(nowRune == rune('[')){
				isBlock = true
				bFloor = 1
			}else if(nowRune == rune('{')){
				isObject = true
				oFloor = 1
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



		if(startIdx >= 0 && runes[nowIdx] == rune(',') && !isStr && !isBlock && !isObject){
			result = append(result, string(runes[startIdx : nowIdx]))
			startIdx = -1
			continue
		}

		if(startIdx >= 0 && isStr){
			if nowRune == rune('\\') {
				nowIdx++
				continue
			}
			if nowRune == rune('"'){
				if nowIdx < len(runes)-1 && runes[nowIdx+1] == rune(':') {
					nowIdx++
					continue
				}

				result = append(result, string(runes[startIdx : nowIdx+1]))
				isStr = false
				startIdx = -1
				continue
			}
			continue
		}

		if(startIdx >= 0 && isBlock){
			if(isInnerStr){
				if nowRune == rune('\\') {
					nowIdx++
					continue
				}
				if nowRune == rune('"'){
					isInnerStr = false
				}
			}else{
				if nowRune == '"' && !(string(runes[nowIdx-1 : nowIdx+1]) == "\\\""){
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


		if(startIdx >= 0 && isObject){
			if(isInnerStr){
				if nowRune == rune('\\') {
					nowIdx++
					continue
				}
				if nowRune == rune('"'){
					isInnerStr = false
				}
			}else{
				if nowRune == rune('"') && !(string(runes[nowIdx-1 : nowIdx+1]) == "\\\"") {
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

		

	}
	return result
}



	





