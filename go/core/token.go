package core

import "strconv"
import "bytes"
import "fmt"
import "strings"
import "encoding/hex"

type Token struct {
	Tp 		int
	Val 	interface{}
}

func (t *Token) Str() string{
	return t.Val.(string)
}

func (t *Token) Tks() []*Token{
	return t.Val.([]*Token)
}

func (t *Token) Int() int{
	return t.Val.(int)
}

func (t *Token) Float() float64{
	return t.Val.(float64)
}

func (t *Token) Ctx() *BindMap{
	return t.Val.(*BindMap)
}

func (t *Token) ToString() string{
	if t == nil {
		return "nil"
	}
	switch t.Tp {
	case NIL:
		return "null"
	case NONE:
		return "none"
	case ERR:
		return "Err: " + t.Str()
	case LIT_WORD:
		return t.Str()
	case GET_WORD:
		return t.Str()
	case DATATYPE:
		return TypeToStr(t.Int())
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
		return "#'" + string(t.Val.(rune)) + "'"
	case STRING:
		temp := t.Str()
		temp = strings.ReplaceAll(temp, "^", "^^")
		temp = strings.ReplaceAll(temp, "\"", "^\"")
		for k, v := range(charToCaretMap) {
			temp = strings.ReplaceAll(temp, string(k), v)
		}
		return "\"" + temp + "\""
	case PROP:
		return "/" + t.Str()
	case FILE:
		return "%" + t.Str()
	case BIN:
		return "#{" + hex.EncodeToString(t.Val.([]byte)) + "}"
	case SET_WORD:
		return t.Str() + ":"
	case PUT_WORD:
		return t.Str() + ":="
	case PAREN:
		var buffer bytes.Buffer
		buffer.WriteString("(")
		for _, item := range t.Tks(){
			buffer.WriteString(item.ToString())
			buffer.WriteString(" ")
		}
		if len(buffer.Bytes()) > 1 {
			buffer.Bytes()[len(buffer.Bytes())-1] = ')'
		}else{
			buffer.WriteString(")")
		}
		return buffer.String()
	case BLOCK:
		var buffer bytes.Buffer
		buffer.WriteString("[")
		for _, item := range t.Tks(){
			buffer.WriteString(item.ToString())
			buffer.WriteString(" ")
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
			buffer.WriteString(k)
			buffer.WriteString(": ")
			buffer.WriteString(v.ToString())
			buffer.WriteString(" ")
		}
		if len(buffer.Bytes()) > 1 {
			buffer.Bytes()[len(buffer.Bytes())-1] = '}'
		}else{
			buffer.WriteString("}")
		}
		return buffer.String()
	case PATH:
		var buffer bytes.Buffer
		for _, item := range t.Tks(){
			buffer.WriteString(item.ToString())
			buffer.WriteString("/")
		}
		var temp = buffer.String()
		return temp[0:len(temp)-1]
	case NATIVE:
		return "native: " + t.Val.(Native).Str
	case OP:
		return "op: " + t.Val.(Native).Str
	case FUNC:
		var buffer bytes.Buffer
		buffer.WriteString("!func{[")
		for _, item := range t.Val.(Func).Args.List(){
			buffer.WriteString(item.ToString())
			buffer.WriteString(" ")
		}
		for i:=0; i<len(t.Val.(Func).Props.List()); i+=2 {
			buffer.WriteString(t.Val.(Func).Props.Get(i).ToString())
			buffer.WriteString(" ")
			if t.Val.(Func).Props.Get(i+1) != nil {
				buffer.WriteString(t.Val.(Func).Props.Get(i+1).ToString())
				buffer.WriteString(" ")
			}
		}
		if buffer.Bytes()[len(buffer.Bytes())-1] != '[' {
			buffer.Bytes()[len(buffer.Bytes())-1] = ']'
		}else{
			buffer.WriteString("]")
		}
		buffer.WriteString(" [")
		for _, item := range t.Val.(Func).Codes.List(){
			buffer.WriteString(item.ToString())
			buffer.WriteString(" ")
		}
		if buffer.Bytes()[len(buffer.Bytes())-1] != '[' {
			buffer.Bytes()[len(buffer.Bytes())-1] = ']'
		}else{
			buffer.WriteString("]")
		}
		buffer.WriteString("}")
		return buffer.String()
	default:
		return t.Str()
	
	}
}

func (t *Token) OutputStr() string{
	if(t.Tp == STRING){
		return t.Str()
	}else{
		return t.ToString()
	}
}

func (t *Token) Echo(){
	fmt.Println(t.ToString())
}

func EchoTokens(ts []*Token){
	var str = "[ "
	for _, item := range(ts){
		str += item.ToString() + " "
	}
	str += "]"
	fmt.Println(str)
}


func (t *Token) Copy(source *Token){
	t.Tp = source.Tp
	t.Val = source.Val
}

func (t *Token) Dup() *Token{
	return &Token{t.Tp, t.Val}
}

func (t *Token) Clone() *Token{
	var result = &Token{t.Tp, t.Val}
	switch t.Tp {
	case BLOCK, PAREN, PATH:
		result.Val = make([]*Token, 0)
		for _, item := range(t.Tks()){
			result.Val = append(result.Tks(), item.Dup())
		}
		return result
	case OBJECT:
		result.Val = &BindMap{make(map[string]*Token), t.Ctx().Father, t.Ctx().Tp}
		for k, v := range(t.Ctx().Table) {
			result.Ctx().Table[k] = v.Clone()
		}
		return result
	default:
		return result
	}
}

func (t *Token) CloneDeep() *Token{
	var result = &Token{t.Tp, t.Val}
	switch t.Tp {
	case BLOCK, PAREN, PATH:
		result.Val = make([]*Token, 0)
		for _, item := range(t.Tks()){
			result.Val = append(result.Tks(), item.CloneDeep())
		}
		return result
	case OBJECT:
		result.Val = &BindMap{make(map[string]*Token), t.Ctx().Father, t.Ctx().Tp}
		for k, v := range(t.Ctx().Table) {
			result.Ctx().Table[k] = v.CloneDeep()
		}
		return result
	default:
		return result
	}
}


func (t *Token) GetVal(ctx *BindMap, stack *EvalStack) (*Token, error){
	var result Token
	switch t.Tp {
	case WORD:
		return ctx.Get(t.Str()), nil
	case GET_WORD:
		return ctx.Get(t.Str()), nil
	case WRAP:
		return t.Val.(*Token), nil
	case LIT_WORD:
		result.Tp = WORD
		result.Val = t.Str()
		return &result, nil
	case PAREN:
		return stack.Eval(t.Tks(), ctx)
	case PATH:
		return t.GetPathVal(ctx, stack)
	default:
		return t, nil
	}
}

func (t Token) Explen() int{
	switch t.Tp{
	case SET_WORD:
		return 2
	case PUT_WORD:
		return 2
	case NATIVE:
		return t.Val.(Native).Explen
	case FUNC:
		return t.Val.(Func).Args.Len() + 1
	case OP:
		return 3
	case PATH:
		if t.IsGetPath() {
			return 1
		}else{
			return 2
		}
	default:
		return 1
	}
}

func (t *Token) IsSetPath() bool{
	if t.Tp != PATH || len(t.Tks()) <= 0 {
		return false
	}else{
		return t.Tks()[len(t.Tks())-1].Tp == SET_WORD 
	}
}

func (t *Token) IsGetPath() bool{
	if t.Tp != PATH || len(t.Tks()) <= 0 {
		return false
	}else{
		var lastTp = t.Tks()[len(t.Tks())-1].Tp
		return lastTp != SET_WORD 
	}
}

func (t *Token) GetPathVal(ctx *BindMap, stack *EvalStack) (*Token, error){
	result, err := t.Tks()[0].GetVal(ctx, stack)
	if err != nil {
		return nil, err
	}
	
	var curCtx = ctx
	for idx := 1; idx < len(t.Tks()); idx++ {
		if result.Tp == OBJECT {
			curCtx = result.Ctx()
		}
		key := t.Tks()[idx]
		if key.Tp == PAREN || key.Tp == GET_WORD {
			key, err = key.GetVal(ctx, stack)
		}

		if err != nil {
			return nil, err
		}
		if result.Tp == BLOCK || result.Tp == PAREN {
			if key.Tp == INTEGER {
				if key.Int() > 0 && key.Int() - 1 < len(result.Tks()) {
					result = result.Tks()[key.Int()-1]
					continue
				}else{
					return &Token{ERR, "Error path!"}, nil
				}
			}else if key.Tp == WORD || key.Tp == STRING {
				var found = false
				for idx := 0; idx < len(result.Tks()) - 1; idx++ {
					if (result.Tks()[idx].Tp == WORD || result.Tks()[idx].Tp == SET_WORD || result.Tks()[idx].Tp == STRING) && 
							result.Tks()[idx].Str() == key.Str(){
						result = result.Tks()[idx+1]
						found = true
						break
					}
				}
				if found {
					continue
				}
				result = &Token{NONE, nil}
				continue
			}
			return &Token{ERR, "Error path!"}, nil
		}else if result.Tp == OBJECT {
			if key.Tp == WORD || key.Tp == STRING {
				var found bool
				result, found = result.Ctx().Table[key.ToString()]
				if idx == len(t.Tks())-1 {
					if !found {
						return &Token{NONE, "none"}, nil
					}
					if result.Tp == FUNC {
						temp := Token{PATH, make([]*Token, 0, 8)}
						temp.Val = append(temp.Tks(), result)
						temp.Val = append(temp.Tks(), &Token{OBJECT, curCtx})
						return &temp, nil
					}
				}

				continue
			}
			return &Token{ERR, "Error path!"}, nil
		}else if result.Tp == FUNC {
			temp := Token{PATH, make([]*Token, 0, 8)}
			temp.Val = append(temp.Tks(), result)
			temp.Val = append(temp.Tks(), &Token{OBJECT, curCtx})
			for i:=idx; i<len(t.Tks()); i++ {
				temp.Val = append(temp.Tks(), t.Tks()[i])
			}
			return &temp, nil
		}else if result.Tp == STRING && key.Tp == INTEGER {
			runes := []rune(result.Str())
			if key.Int() > 0 && key.Int() <= len(runes) {
				result = &Token{CHAR, runes[key.Int() - 1]}
				continue
			}

		}
		return &Token{ERR, "Error path!"}, nil
	}

	return result, nil
}

func (t *Token)SetPathVal(val *Token, ctx *BindMap, stack *EvalStack) (*Token, error){
	var holderPath = t.Dup()
	holderPath.Val = holderPath.Tks()[0: len(holderPath.Tks())-1]
	holder, err := holderPath.GetPathVal(ctx, stack)
	if err != nil {
		return nil, err
	}

	var key = t.Tks()[len(t.Tks())-1].Str()

	if holder != nil {
		if holder.Tp == BLOCK || holder.Tp == PAREN {
			if IsNumberStr(key) == 0 {
				idx, err := strconv.Atoi(key)
				if err != nil {
					panic(err)
				}
				if idx > 0 && idx <= len(holder.Tks()){
					holder.Tks()[idx-1] = val.Clone()
					return holder, nil
				}
			} else {
				for i:=0; i<len(holder.Tks())-1; i+=2{
					if holder.Tks()[i].OutputStr() == key {
						holder.Tks()[i+1] = val.Clone()
						return holder, nil
					}
				}
			}

			return &Token{ERR, "Error path!"}, nil
		}else if holder.Tp == OBJECT {
			holder.Ctx().Table[key] = val
			return holder, nil
		}else if holder.Tp == STRING {
			if IsNumberStr(key) == 0 {
				idx, err := strconv.Atoi(key)
				if err != nil {
					panic(err)
				}
				runes := []rune(holder.Val.(string))
				length := len(runes)
				if idx > 0 && idx <= length {
					idx--
					if val.Tp == STRING {
						holder.Val = string(runes[0:idx]) + val.Str() + string(runes[idx+1:length])
						return holder, nil
					}else if val.Tp == CHAR {
						var temp = runes[0:idx]
						temp = append(temp, val.Val.(rune))
						temp = append(temp, runes[idx+1:]...)
						holder.Val = string(temp)
						return holder, nil
					}
				}
			}

		}else{
			return &Token{ERR, "Error path!"}, nil
		}

	}
	return &Token{ERR, "Error path!"}, nil
}


func (t *Token)GetPathExpLen() int{
	var f = t.Tks()[0]
	if f.Tp != FUNC {
		return 1
	}

	var length = f.Val.(Func).Args.Len() + 1

	for i:=2; i<len(t.Tks()); i++ {
		for j:=0; j<f.Val.(Func).Props.Len(); j+=2 {
			if t.Tks()[i].Str() == f.Val.(Func).Props.Get(j).Str() && f.Val.(Func).Props.Get(j +1) != nil {
				length++
			}
		}
	}

	return length

}


func (t *Token)ToBool() bool {
	if t == nil {
		return false
	}
	switch t.Tp {
	case LOGIC:
		return t.Val.(bool)
	case INTEGER:
		return t.Int() != 0
	case DECIMAL:
		return t.Float() != 0.0
	case CHAR:
		return t.Val.(byte) != 0
	case STRING:
		return t.Str() != ""
	case BLOCK, PAREN, PATH:
		return len(t.Tks()) > 0
	case OBJECT:
		return t.Ctx() != nil
	default:
		return false
	}
}







