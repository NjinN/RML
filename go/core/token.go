package core

import "strconv"
import "bytes"
import "fmt"
import "strings"

type Token struct {
	Tp 		int
	Val 	interface{}
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
		return "Err: " + t.Val.(string)
	case LIT_WORD:
		return t.Val.(string)
	case GET_WORD:
		return t.Val.(string)
	case DATATYPE:
		return t.Val.(string)
	case LOGIC:
		return strconv.FormatBool(t.Val.(bool))
	case INTEGER:
		return strconv.Itoa(t.Val.(int))
	case DECIMAL:
		var result = strconv.FormatFloat(t.Val.(float64), 'f', -1, 64)
		if strings.IndexByte(result, '.') < 0 {
			result += ".0"
		}
		return result
	case CHAR:
		return "#'" + string(t.Val.(uint8)) + "'"
	case STRING:
		temp := t.Val.(string)
		temp = strings.ReplaceAll(temp, "^", "^^")
		temp = strings.ReplaceAll(temp, "\"", "^\"")
		for k, v := range(charToCaretMap) {
			temp = strings.ReplaceAll(temp, string(k), v)
		}
		return "\"" + temp + "\""
	case PROP:
		return "/" + t.Val.(string)
	case SET_WORD:
		return t.Val.(string) + ":"
	case PUT_WORD:
		return t.Val.(string) + ":="
	case PAREN:
		var buffer bytes.Buffer
		buffer.WriteString("(")
		for _, item := range t.Val.([]*Token){
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
		for _, item := range t.Val.([]*Token){
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
		for k, v := range t.Val.(*BindMap).Table {
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
		for _, item := range t.Val.([]*Token){
			buffer.WriteString(item.ToString())
			buffer.WriteString("/")
		}
		var temp = buffer.String()
		return temp[0:len(temp)-1]
	case NATIVE:
		return "native: " + t.Val.(Native).Str
	case FUNC:
		var buffer bytes.Buffer
		buffer.WriteString("func [")
		for _, item := range t.Val.(Func).Args{
			buffer.WriteString(item.ToString())
			buffer.WriteString(" ")
		}
		for i:=0; i<len(t.Val.(Func).Props); i+=2 {
			buffer.WriteString(t.Val.(Func).Props[i].ToString())
			buffer.WriteString(" ")
			if t.Val.(Func).Props[i+1] != nil {
				buffer.WriteString(t.Val.(Func).Props[i+1].ToString())
				buffer.WriteString(" ")
			}
		}
		if buffer.Bytes()[len(buffer.Bytes())-1] != '[' {
			buffer.Bytes()[len(buffer.Bytes())-1] = ']'
		}else{
			buffer.WriteString("]")
		}
		buffer.WriteString(" [")
		for _, item := range t.Val.(Func).Codes{
			buffer.WriteString(item.ToString())
			buffer.WriteString(" ")
		}
		if buffer.Bytes()[len(buffer.Bytes())-1] != '[' {
			buffer.Bytes()[len(buffer.Bytes())-1] = ']'
		}else{
			buffer.WriteString("]")
		}
		return buffer.String()
	default:
		return t.Val.(string)
	
	}
}

func (t *Token) OutputStr() string{
	if(t.Tp == STRING){
		return t.Val.(string)
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

func (t *Token) Clone() *Token{
	var result = Token{t.Tp, t.Val}
	return &result
}

func (t *Token) CloneDeep() *Token{
	var result = &Token{t.Tp, t.Val}
	switch t.Tp {
	case BLOCK, PAREN, PATH:
		result.Val = make([]*Token, 0)
		for _, item := range(t.Val.([]*Token)){
			result.Val = append(result.Val.([]*Token), item.CloneDeep())
		}
		return result
	case OBJECT:
		result.Val = &BindMap{make(map[string]*Token), t.Val.(*BindMap).Father, t.Val.(*BindMap).Tp}
		for k, v := range(t.Val.(*BindMap).Table) {
			result.Val.(*BindMap).Table[k] = v.CloneDeep()
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
		return ctx.Get(t.Val.(string)), nil
	case GET_WORD:
		return ctx.Get(t.Val.(string)), nil
	case LIT_WORD:
		result.Tp = WORD
		result.Val = t.Val.(string)
		return &result, nil
	case PAREN:
		return stack.Eval(t.Val.([]*Token), ctx)
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
		return len(t.Val.(Func).Args) + 1
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
	if t.Tp != PATH || len(t.Val.([]*Token)) <= 0 {
		return false
	}else{
		return t.Val.([]*Token)[len(t.Val.([]*Token))-1].Tp == SET_WORD 
	}
}

func (t *Token) IsGetPath() bool{
	if t.Tp != PATH || len(t.Val.([]*Token)) <= 0 {
		return false
	}else{
		var lastTp = t.Val.([]*Token)[len(t.Val.([]*Token))-1].Tp
		return lastTp != SET_WORD 
	}
}

func (t *Token) GetPathVal(ctx *BindMap, stack *EvalStack) (*Token, error){
	result, err := t.Val.([]*Token)[0].GetVal(ctx, stack)
	if err != nil {
		return nil, err
	}
	var curCtx = ctx
	for idx := 1; idx < len(t.Val.([]*Token)); idx++ {
		if result.Tp == OBJECT {
			curCtx = result.Val.(*BindMap)
		}
		key := t.Val.([]*Token)[idx]
		if key.Tp == PAREN || key.Tp == GET_WORD {
			key, err = key.GetVal(ctx, stack)
		}

		if err != nil {
			return nil, err
		}
		if result.Tp == BLOCK || result.Tp == PAREN {
			if key.Tp == INTEGER {
				if key.Val.(int) > 0 && key.Val.(int) - 1 < len(result.Val.([]*Token)) {
					result = result.Val.([]*Token)[key.Val.(int)-1]
					continue
				}else{
					return &Token{ERR, "Error path!"}, nil
				}
			}else if key.Tp == WORD || key.Tp == STRING {
				var found = false
				for idx := 0; idx < len(result.Val.([]*Token)) - 1; idx++ {
					if (result.Val.([]*Token)[idx].Tp == WORD || result.Val.([]*Token)[idx].Tp == SET_WORD || result.Val.([]*Token)[idx].Tp == STRING) && 
							result.Val.([]*Token)[idx].Val.(string) == key.Val.(string){
						result = result.Val.([]*Token)[idx+1]
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
				result, found = result.Val.(*BindMap).Table[key.ToString()]
				if idx == len(t.Val.([]*Token))-1 {
					if !found {
						return &Token{NONE, "none"}, nil
					}
					if result.Tp == FUNC {
						temp := Token{PATH, make([]*Token, 0, 8)}
						temp.Val = append(temp.Val.([]*Token), result)
						temp.Val = append(temp.Val.([]*Token), &Token{OBJECT, curCtx})
						return &temp, nil
					}
				}

				continue
			}
			return &Token{ERR, "Error path!"}, nil
		}else if result.Tp == FUNC {
			temp := Token{PATH, make([]*Token, 0, 8)}
			temp.Val = append(temp.Val.([]*Token), result)
			temp.Val = append(temp.Val.([]*Token), &Token{OBJECT, curCtx})
			for i:=idx; i<len(t.Val.([]*Token)); i++ {
				temp.Val = append(temp.Val.([]*Token), t.Val.([]*Token)[i])
			}
			return &temp, nil
		}
		return &Token{ERR, "Error path!"}, nil
	}

	return result, nil
}

func (t *Token)SetPathVal(val *Token, ctx *BindMap, stack *EvalStack) (*Token, error){
	var holderPath = t.Clone()
	holderPath.Val = holderPath.Val.([]*Token)[0: len(holderPath.Val.([]*Token))-1]
	holder, err := holderPath.GetPathVal(ctx, stack)
	if err != nil {
		return nil, err
	}

	var key = t.Val.([]*Token)[len(t.Val.([]*Token))-1].Val.(string)

	if holder != nil {
		if holder.Tp == BLOCK || holder.Tp == PAREN {
			if IsNumberStr(key) == 0 {
				idx, err := strconv.Atoi(key)
				if err != nil {
					panic(err)
				}
				if idx > 0 && idx < len(holder.Val.([]*Token)){
					holder.Val.([]*Token)[idx-1] = val
					return val, nil
				}
			} 

			return &Token{ERR, "Error path!"}, nil
		}else if holder.Tp == OBJECT {
			holder.Val.(*BindMap).Table[key] = val
			return val, nil
		}else{
			return &Token{ERR, "Error path!"}, nil
		}

	}
	return &Token{ERR, "Error path!"}, nil
}


func (t *Token)GetPathExpLen() int{
	var f = t.Val.([]*Token)[0]
	if f.Tp != FUNC {
		return 1
	}

	var length = len(f.Val.(Func).Args) + 1

	for i:=2; i<len(t.Val.([]*Token)); i++ {
		for j:=0; j<len(f.Val.(Func).Props); j+=2 {
			if t.Val.([]*Token)[i].Val.(string) == f.Val.(Func).Props[j].Val.(string) && f.Val.(Func).Props[j +1] != nil {
				length++
			}
		}
	}

	return length

}


