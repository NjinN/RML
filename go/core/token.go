package core

import "strconv"
import "bytes"
import "fmt"

type Token struct {
	Tp 		int
	Val 	interface{}
}

func (t Token) ToString() string{
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
		return strconv.FormatFloat(t.Val.(float64), 'f', -1, 64)
	case CHAR:
		return string(t.Val.(int))
	case STRING:
		return "\"" + t.Val.(string) + "\""
	case SET_WORD:
		return t.Val.(string)
	case PAREN:
		var buffer bytes.Buffer
		buffer.WriteString("( ")
		for _, item := range t.Val.([]*Token){
			buffer.WriteString(item.ToString())
			buffer.WriteString(" ")
		}
		buffer.WriteString(")")
		return buffer.String()
	case BLOCK:
		var buffer bytes.Buffer
		buffer.WriteString("[ ")
		for _, item := range t.Val.([]*Token){
			buffer.WriteString(item.ToString())
			buffer.WriteString(" ")
		}
		buffer.WriteString("]")
		return buffer.String()
	case NATIVE:
		return "native: " + t.Val.(Native).Str
	case FUNC:
		var buffer bytes.Buffer
		buffer.WriteString("func [")
		for _, item := range t.Val.(Func).Args{
			buffer.WriteString(item.ToString())
			buffer.WriteString(" ")
		}
		buffer.WriteString("] [")
		for _, item := range t.Val.(Func).Codes{
			buffer.WriteString(item.ToString())
			buffer.WriteString(" ")
		}
		buffer.WriteString("]")
		return buffer.String()
	default:
		return t.Val.(string)
	
	}
}

func (t *Token) OutputStr() string{
	if(t.Tp == STRING){
		var str = t.ToString()
		return str[1: len(str)-1]
	}else{
		return t.ToString()
	}
}

func (t Token) Echo(){
	fmt.Println(t.OutputStr())
}


func (t *Token) Copy(source *Token){
	t.Tp = source.Tp
	t.Val = source.Val
}

func (t *Token) Clone() Token{
	var result = Token{t.Tp, t.Val}
	return result
}


func (t *Token) GetVal(ctx *BindMap, stack *EvalStack) (*Token, error){
	var result Token
	switch t.Tp {
	case WORD:
		return ctx.Get(t.Val.(string)), nil
	case LIT_WORD:
		result.Tp = WORD
		result.Val = t.Val.(string)
		return &result, nil
	case PAREN:
		rs, err := stack.Eval(t.Val.([]*Token), ctx)
		return rs, err
	default:
		return t, nil
	}
}

func (t Token) Explen() int{
	switch t.Tp{
	case SET_WORD:
		return 2
	case NATIVE:
		return t.Val.(Native).Explen
	case FUNC:
		return len(t.Val.(Func).Args) + 1
	case OP:
		return 3
	default:
		return 1
	}
}


