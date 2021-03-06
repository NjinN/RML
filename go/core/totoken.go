package core

import "strings"
import "strconv"
import "encoding/hex"
import "sync"
import "regexp"
import "errors"
// import "fmt"

func ToToken(s string, ctx *BindMap, es *EvalStack) *Token{
	var result Token
	var str = Trim(s)

	var matched bool
	var err error

	// if str[0] != '"' && !strings.Contains(str, "://") {
	// 	str = strings.ToLower(str)
	// }

	// fmt.Println(s)

	if(strings.ToLower(str) == "none"){
		result.Tp = NONE
		result.Val = ""
		return &result
	}

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

	if(str[0] == ':'){
		result.Tp = GET_WORD
		result.Val = strings.ToLower(str[1 :])
		return &result
	}
	
	if(str[len(str)-1] == ':' && strings.IndexByte(str, '/') < 0){
		result.Tp = SET_WORD
		result.Val = strings.ToLower(str[0 : len(str)-1])
		return &result
	}

	if EndWith(str, ":="){
		result.Tp = PUT_WORD
		result.Val = strings.ToLower(str[0 : len(str)-2])
		return &result
	}

	if EndWith(str, "设为"){
		result.Tp = PUT_WORD
		result.Val = strings.ToLower(str[0 : len(str)-6])
		return &result
	}

	if EndWith(str, "为") && strings.IndexByte(str, '/') < 0 {
		result.Tp = SET_WORD
		result.Val = strings.ToLower(str[0 : len(str)-3])
		return &result
	}

	if(str[len(str)-1] == '!'){
		result.Tp = DATATYPE
		result.Val = StrToType(strings.ToLower(str))
		return &result
	}

	if(str[0] == '%' && len(str) > 1){
		result.Tp = FILE
		result.Val = str[1:]
		return &result
	}

	if len(str) > 3 && StartWith(str, "#{") && str[len(str)-1] == '}' {
		bin, err := hex.DecodeString(str[2:len(str)-1])
		if err != nil {
			result.Tp = ERR
			result.Val = "Error bin format of " + str
		}else{
			result.Tp = BIN
			result.Val = bin
		}
		return &result
	}

	if(str[0] == '/' && str != "/" && str != "/="){
		result.Tp = PROP
		result.Val = strings.ToLower(str[1 :])
		return &result
	}

	if(str[0] == '"'){
		result.Tp = STRING
		temp := str[1 : len(str)-1]
		for k, v := range(caretToCharMap) {
			temp = strings.ReplaceAll(temp, string(k), v)
		}
		temp = strings.ReplaceAll(temp, "^^", "^")
		temp = strings.ReplaceAll(temp, `^"`, `"`)
		result.Val = temp
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
		result.Val = NewTks(8)
		result.Val.(*TokenList).AddArr(ToTokens(str[1 : endIdx], ctx, es))
		return &result
	}

	if(str[0] == '('){
		result.Tp = PAREN
		var endIdx int
		for endIdx=len(str)-1; endIdx>=0; endIdx-- {
			if(str[endIdx] == ')'){
				break
			}
		}
		result.Val = NewTks(8)
		result.Val.(*TokenList).AddArr(ToTokens(str[1 : endIdx], ctx, es))
		return &result
	}

	if(str[0] == '{'){
		result.Tp = OBJECT
		var endIdx int
		for endIdx=len(str)-1; endIdx>=0; endIdx-- {
			if(str[endIdx] == '}'){
				break
			}
		}
		var blk = ToTokens(str[1 : endIdx], ctx, es)
		
		result.Val = &BindMap{nil, nil, USR_CTX, nil, blk}
		es.TempResult = &result
		return &result
	}

	if(len([]rune(str)) == 4 && str[0 : 2] == "#'" && str[3] == '\''){
		result.Tp = CHAR
		result.Val = []rune(str)[2]
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

	if(IsNumberStr(str) >= 2){
		var strs = strings.Split(str, "..")
		if len(strs) != 2 || IsNumberStr(strs[0]) < 0 || IsNumberStr(strs[0]) > 1 || IsNumberStr(strs[1]) < 0 || IsNumberStr(strs[1]) > 1 {
			result.Tp = ERR
			result.Val = "Error format of " + str
			return &result
		}

		result.Tp = RANGE
		var sToken, eToken Token
		var err error
		if IsNumberStr(strs[0]) == 0 && IsNumberStr(strs[1]) == 0 {
			sToken.Tp = INTEGER
			sToken.Val, err = strconv.Atoi(strs[0])
			if err != nil {
				result.Tp = ERR
				result.Val = "Error format of " + str
				return &result
			}
			eToken.Tp = INTEGER
			eToken.Val, err = strconv.Atoi(strs[1])
			if err != nil {
				result.Tp = ERR
				result.Val = "Error format of " + str
				return &result
			}
		}else{
			sToken.Tp = DECIMAL
			sToken.Val, err = strconv.ParseFloat(strs[0], 64)
			if err != nil {
				result.Tp = ERR
				result.Val = "Error format of " + str
				return &result
			}
			eToken.Tp = DECIMAL
			eToken.Val, err = strconv.ParseFloat(strs[1], 64)
			if err != nil {
				result.Tp = ERR
				result.Val = "Error format of " + str
				return &result
			}
		}

		result.Val = NewTks(4)
		result.List().Add(&sToken)
		result.List().Add(&eToken)
		return &result
	}

	matched, err = regexp.MatchString("^[0-9]{1,9}x[0-9]{1,9}$", str)
	if err != nil {
		result.Tp = ERR
		result.Val = err.Error()
		return &result
	}
	if matched {
		var sToken, eToken Token

		pair := strings.Split(str, "x")

		sToken.Tp = INTEGER
		sToken.Val, err = strconv.Atoi(pair[0])
		if err != nil {
			result.Tp = ERR
			result.Val = "Error format of " + str
			return &result
		}
		eToken.Tp = INTEGER
		eToken.Val, err = strconv.Atoi(pair[1])
		if err != nil {
			result.Tp = ERR
			result.Val = "Error format of " + str
			return &result
		}

		result.Tp = PAIR
		result.Val = NewTks(4)
		result.List().Add(&sToken)
		result.List().Add(&eToken)
		return &result
	}



	if(str[0] == '\''){
		result.Tp = LIT_WORD
		result.Val = strings.ToLower(str[1 : ])
		return &result
	}

	if(strings.Contains(str, "://")) {
		result.Tp = URL
		result.Val =str
		return &result
	}

	if(strings.IndexByte(str, '/') >= 0 && str != "/" && str != "/="){
		result.Tp = PATH
		result.Val = PathToTokens(str, ctx, es)
		return &result
	}

	if str[0] == '!' && strings.IndexByte(str, '{') > 1 && str[len(str)-1] == '}' {
		result.Tp = WRAP
		var startIdx = strings.IndexByte(str, '{')
		var typeStr = strings.ToLower(str[1:startIdx])
		var bodyStr = str[startIdx+1:len(str)-1]
		var bodyBlock = ToTokens(Trim(bodyStr), ctx, es)

		if typeStr == "func" {
			if len(bodyBlock) >= 2 {
				bodyBlock = append([]*Token{&Token{WORD, "func"}}, bodyBlock...)
				temp, err := es.Eval(bodyBlock, ctx)
				if err != nil {
					panic(err)
				}
				result.Val = temp
				return &result
			} 
		}else if typeStr == "map" {
			var m Rmap 
			m.Lazy = bodyBlock
			result.Tp = MAP
			result.Val = &m
			return &result
		
		}else if typeStr == "timer" {
			if len(bodyBlock)  != 2 {
				return &Token{ERR, "Error format of " + str}
			}

			time := 0.0
			if bodyBlock[0].Tp == INTEGER {
				time = float64(bodyBlock[0].Int())
			}else if bodyBlock[0].Tp == DECIMAL {
				time = bodyBlock[0].Float()
			}else {
				return &Token{ERR, "Error format of " + str}
			}

			if bodyBlock[1].Tp != BLOCK {
				return &Token{ERR, "Error format of " + str}
			}

			var timer Timer
			timer.Time = time
			timer.Code = bodyBlock[1].List()

			result.Tp = TIMER
			result.Val = &timer
			return &result
		}

		result.Tp = ERR
		result.Val = "Error format of" + str
		return &result
	}

	/*************  parse time format start  *************/

	//only date
	matched, err = regexp.MatchString("^-?[0-9]{4}-[0-9]{2}-[0-9]{2}$", str)

	if err != nil {
		result.Tp = ERR
		result.Val = err.Error()
		return &result
	}

	if matched {
		result.Tp = TIME
		timeClock := TimeClock{}
		timeClock.Second = 0
		timeClock.FloatSecond = 0

		if str[0] == '-' {
			str = str[1:]
			timeClock.Negative = true
		}
		
		days := DateStrToDays(str)
		if days == 0 {
			result.Tp = ERR
			result.Val = "Error format of " + str
			return &result
		}

		timeClock.Date = days

		result.Val = &timeClock
		return &result
	}

	//only time
	matched, err = regexp.MatchString("^-?[0-9]{2}:[0-9]{2}:[0-9]{2}(\\.[0-9]{1,8})?$", str)

	if err != nil {
		result.Tp = ERR
		result.Val = err.Error()
		return &result
	}

	if matched {
		result.Tp = TIME
		timeClock := TimeClock{}
		timeClock.Date = 0
		timeClock.FloatSecond = 0

		if str[0] == '-' {
			str = str[1:]
			timeClock.Negative = true
		}

		secSlice := strings.Split(str, ".")
		secs := TimeStrToSecs(secSlice[0])

		if secs < 0 {
			result.Tp = ERR
			result.Val = "Error format of " + str
			return &result
		}

		timeClock.Second = secs

		if len(secSlice) > 1 {
			floatSecs, err := strconv.ParseFloat("0." + secSlice[1], 64)
			if err != nil {
				result.Tp = ERR
				result.Val = "Error format of " + str
				return &result
			}
			timeClock.FloatSecond = floatSecs
		}
		
		result.Val = &timeClock
		return &result
	}

	//date and time
	matched, err = regexp.MatchString("^-?[0-9]{4}-[0-9]{2}-[0-9]{2}\\+[0-9]{2}:[0-9]{2}:[0-9]{2}(\\.[0-9]{1,8})?$", str)
	if err != nil {
		result.Tp = ERR
		result.Val = err.Error()
		return &result
	}

	if matched {

		result.Tp = TIME
		timeClock := TimeClock{}
		timeClock.FloatSecond = 0

		if str[0] == '-' {
			str = str[1:]
			timeClock.Negative = true
		}

		timeSlice := strings.Split(str, "+")

		days := DateStrToDays(timeSlice[0])
		if days <= 0 {
			
			result.Tp = ERR
			result.Val = "Error format of " + str
			return &result
		}

		timeClock.Date = days

		secSlice := strings.Split(timeSlice[1], ".")

		secs := TimeStrToSecs(secSlice[0])
		if secs < 0 {
			result.Tp = ERR
			result.Val = "Error format of " + str
			return &result
		}

		timeClock.Second = secs

		if len(secSlice) > 1 {
			floatSecs, err := strconv.ParseFloat("0." + secSlice[1], 64)
			if err != nil {
				result.Tp = ERR
				result.Val = "Error format of " + str
				return &result
			}
			timeClock.FloatSecond = floatSecs
		}

		result.Val = &timeClock
		return &result
	}

	/*************  parse time format end  *************/

	result.Tp = WORD
	result.Val = str
	return &result
}

func ToTokens(str string, ctx *BindMap, es *EvalStack) []*Token{
	var result []*Token
	var strs = StrCut(str)
	for _, item := range strs {
		result = append(result, ToToken(item, ctx, es))
	}
	return result
}

func PathToTokens(str string, ctx *BindMap, es *EvalStack) *TokenList{
	var result = NewTks(8)
	var strs = PathCut(str)
	for _, item := range strs {
		result.Add(ToToken(item, ctx, es))
	}
	return result
}


func MakeObject(blk []*Token, ctx *BindMap, es *EvalStack) *BindMap {
	var c = BindMap{make(map[string]*Token, 8), ctx, USR_CTX, &sync.RWMutex{}, nil}
	var orginSts = es.IsLocal
	es.IsLocal = true
	es.Eval(blk, &c)
	es.IsLocal = orginSts
	return &c
}


func MakeRmap(bodyBlock []*Token, ctx *BindMap, es *EvalStack) (*Rmap, error) {
	var m Rmap
	m.Table = make(map[string]TokenPair, 8)
	m.Lock = &sync.RWMutex{}
	for _, item := range bodyBlock {
		

		if item.Tp != BLOCK {
			return nil, errors.New("Error format") 
		}
		blk, err := es.Eval(item.Tks(), ctx, 1)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		if blk.List().Len() != 2 {
			return nil, errors.New("Error format") 
		}

		if item.List().Len() == 2 {
			var pair TokenPair
			pair.Key = item.Tks()[0]
			pair.Val = blk.Tks()[1]
			
			var keyString = TypeToStr(item.Tks()[0].Tp) + item.Tks()[0].ToString()
			m.Table[keyString] = pair
		}else{
			var pair TokenPair
			pair.Key = blk.Tks()[0]
			pair.Val = blk.Tks()[1]
			
			var keyString = TypeToStr(blk.Tks()[0].Tp) + blk.Tks()[0].ToString()
			m.Table[keyString] = pair
		}
		
	
	}
	return &m, nil
}












