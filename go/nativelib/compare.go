package nativelib

import . "github.com/NjinN/RML/go/core"

func Eq(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result = Token{Tp: LOGIC}

	switch args[1].Tp {
	case NONE:
		result.Val = false
		return &result, nil
	case LOGIC:
		if args[2].Tp == LOGIC {
			result.Val = args[1].Val.(bool) == args[2].Val.(bool)
			return &result, nil
		}
	case INTEGER:
		switch args[2].Tp {
		case INTEGER:
			result.Val = args[1].Int() == args[2].Int()
			return &result, nil
		case DECIMAL:
			result.Val = float64(args[1].Int()) == args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Int() == int(args[2].Val.(byte))
			return &result, nil
		default:
		}
	case DECIMAL:
		switch args[2].Tp {
		case INTEGER:
			result.Val = args[1].Float() == float64(args[2].Int())
			return &result, nil
		case DECIMAL:
			result.Val = args[1].Float() == args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Float() == float64(args[2].Val.(byte))
			return &result, nil
		default:
		}
	case CHAR:
		switch args[2].Tp {
		case INTEGER:
			result.Val = int(args[1].Val.(rune)) == args[2].Int()
			return &result, nil
		case DECIMAL:
			result.Val = float64(args[1].Val.(rune)) == args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Val.(rune) == args[2].Val.(rune)
			return &result, nil
		default:
		}
	case STRING:
		if args[2].Tp == STRING {
			result.Val = args[1].Str() == args[2].Str()
			return &result, nil
		}
	case WORD:
		if args[2].Tp == WORD {
			result.Val = args[1].Str() == args[2].Str()
			return &result, nil
		}
	case DATATYPE:
		if args[2].Tp == DATATYPE {
			result.Val = args[1].Uint8() == args[2].Uint8()
			return &result, nil
		}
	default:
		result.Tp = ERR
		result.Val = "Type Mismatch"
	}

	return &result, nil
}

func Gt(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result = Token{Tp: LOGIC}

	switch args[1].Tp {
	case NONE:
		result.Val = false
		return &result, nil
	case LOGIC:
		if args[2].Tp == LOGIC {
			if args[1].Val.(bool) == true && args[2].Val.(bool) == false {
				result.Val = true
			} else {
				result.Val = false
			}
			return &result, nil
		}
	case INTEGER:
		switch args[2].Tp {
		case INTEGER:
			result.Val = args[1].Int() > args[2].Int()
			return &result, nil
		case DECIMAL:
			result.Val = float64(args[1].Int()) > args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Int() > int(args[2].Val.(byte))
			return &result, nil
		default:
		}
	case DECIMAL:
		switch args[2].Tp {
		case INTEGER:
			result.Val = args[1].Float() > float64(args[2].Int())
			return &result, nil
		case DECIMAL:
			result.Val = args[1].Float() > args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Float() > float64(args[2].Val.(byte))
			return &result, nil
		default:
		}
	case CHAR:
		switch args[2].Tp {
		case INTEGER:
			result.Val = int(args[1].Val.(byte)) > args[2].Int()
			return &result, nil
		case DECIMAL:
			result.Val = float64(args[1].Val.(byte)) > args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Val.(byte) > args[2].Val.(byte)
			return &result, nil
		default:
		}
	case STRING:
		if args[2].Tp == STRING {
			result.Val = args[1].Str() > args[2].Str()
			return &result, nil
		}
	case WORD:
		if args[2].Tp == WORD {
			result.Val = args[1].Str() > args[2].Str()
			return &result, nil
		}
	default:
		result.Tp = LOGIC
		result.Val = "Type Mismatch"
	}

	return &result, nil
}

func Lt(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result = Token{Tp: LOGIC}

	switch args[1].Tp {
	case NONE:
		result.Val = false
		return &result, nil
	case LOGIC:
		if args[2].Tp == LOGIC {
			if args[1].Val.(bool) == false && args[2].Val.(bool) == true {
				result.Val = true
			} else {
				result.Val = false
			}
			return &result, nil
		}
	case INTEGER:
		switch args[2].Tp {
		case INTEGER:
			result.Val = args[1].Int() < args[2].Int()
			return &result, nil
		case DECIMAL:
			result.Val = float64(args[1].Int()) < args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Int() < int(args[2].Val.(byte))
			return &result, nil
		default:
		}
	case DECIMAL:
		switch args[2].Tp {
		case INTEGER:
			result.Val = args[1].Float() < float64(args[2].Int())
			return &result, nil
		case DECIMAL:
			result.Val = args[1].Float() < args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Float() < float64(args[2].Val.(byte))
			return &result, nil
		default:
		}
	case CHAR:
		switch args[2].Tp {
		case INTEGER:
			result.Val = int(args[1].Val.(byte)) < args[2].Int()
			return &result, nil
		case DECIMAL:
			result.Val = float64(args[1].Val.(byte)) < args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Val.(byte) < args[2].Val.(byte)
			return &result, nil
		default:
		}
	case STRING:
		if args[2].Tp == STRING {
			result.Val = args[1].Str() < args[2].Str()
			return &result, nil
		}
	case WORD:
		if args[2].Tp == WORD {
			result.Val = args[1].Str() < args[2].Str()
			return &result, nil
		}
	default:
		result.Tp = LOGIC
		result.Val = "Type Mismatch"
	}

	return &result, nil
}

func Ge(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result = Token{Tp: LOGIC}

	switch args[1].Tp {
	case NONE:
		result.Val = false
		return &result, nil
	case LOGIC:
		if args[2].Tp == LOGIC {
			if (args[1].Val.(bool) == true && args[2].Val.(bool) == false) || (args[1].Val.(bool) == args[2].Val.(bool)) {
				result.Val = true
			} else {
				result.Val = false
			}
			return &result, nil
		}
	case INTEGER:
		switch args[2].Tp {
		case INTEGER:
			result.Val = args[1].Int() >= args[2].Int()
			return &result, nil
		case DECIMAL:
			result.Val = float64(args[1].Int()) >= args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Int() >= int(args[2].Val.(byte))
			return &result, nil
		default:
		}
	case DECIMAL:
		switch args[2].Tp {
		case INTEGER:
			result.Val = args[1].Float() >= float64(args[2].Int())
			return &result, nil
		case DECIMAL:
			result.Val = args[1].Float() >= args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Float() >= float64(args[2].Val.(byte))
			return &result, nil
		default:
		}
	case CHAR:
		switch args[2].Tp {
		case INTEGER:
			result.Val = int(args[1].Val.(byte)) >= args[2].Int()
			return &result, nil
		case DECIMAL:
			result.Val = float64(args[1].Val.(byte)) >= args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Val.(byte) >= args[2].Val.(byte)
			return &result, nil
		default:
		}
	case STRING:
		if args[2].Tp == STRING {
			result.Val = args[1].Str() >= args[2].Str()
			return &result, nil
		}
	case WORD:
		if args[2].Tp == WORD {
			result.Val = args[1].Str() >= args[2].Str()
			return &result, nil
		}
	default:
		result.Tp = LOGIC
		result.Val = "Type Mismatch"
	}

	return &result, nil
}

func Le(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result = Token{Tp: LOGIC}

	switch args[1].Tp {
	case NONE:
		result.Val = false
		return &result, nil
	case LOGIC:
		if args[2].Tp == LOGIC {
			if (args[1].Val.(bool) == false && args[2].Val.(bool) == true) || (args[1].Val.(bool) == args[2].Val.(bool)) {
				result.Val = true
			} else {
				result.Val = false
			}
			return &result, nil
		}
	case INTEGER:
		switch args[2].Tp {
		case INTEGER:
			result.Val = args[1].Int() <= args[2].Int()
			return &result, nil
		case DECIMAL:
			result.Val = float64(args[1].Int()) <= args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Int() <= int(args[2].Val.(byte))
			return &result, nil
		default:
		}
	case DECIMAL:
		switch args[2].Tp {
		case INTEGER:
			result.Val = args[1].Float() <= float64(args[2].Int())
			return &result, nil
		case DECIMAL:
			result.Val = args[1].Float() <= args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Float() <= float64(args[2].Val.(byte))
			return &result, nil
		default:
		}
	case CHAR:
		switch args[2].Tp {
		case INTEGER:
			result.Val = int(args[1].Val.(byte)) <= args[2].Int()
			return &result, nil
		case DECIMAL:
			result.Val = float64(args[1].Val.(byte)) <= args[2].Float()
			return &result, nil
		case CHAR:
			result.Val = args[1].Val.(byte) <= args[2].Val.(byte)
			return &result, nil
		default:
		}
	case STRING:
		if args[2].Tp == STRING {
			result.Val = args[1].Str() <= args[2].Str()
			return &result, nil
		}
	case WORD:
		if args[2].Tp == WORD {
			result.Val = args[1].Str() <= args[2].Str()
			return &result, nil
		}
	default:
		result.Tp = LOGIC
		result.Val = "Type Mismatch"
	}

	return &result, nil
}

func Equals(t1 *Token, t2 *Token) bool {
	if t1.Tp != t2.Tp {
		return false
	}

	switch t1.Tp {
	case NONE:
		return true
	case LIT_WORD, GET_WORD, DATATYPE, STRING, FILE, URL, WORD, SET_WORD, PUT_WORD:
		return t1.Str() == t2.Str()
	case ERR, LOGIC, INTEGER, DECIMAL, CHAR, BIN, WRAP, OP, NATIVE, FUNC:
		return t1.Val == t2.Val
	case RANGE, PAREN, BLOCK, PROP, PATH:
		if t1.List().Len() != t2.List().Len() {
			return false
		}
		for idx, item := range t1.Tks() {
			if !Equals(item, t2.Tks()[idx]) {
				return false
			}
		}
		return true
	case OBJECT:
		if len(t1.Ctx().Table) != len(t2.Ctx().Table) {
			return false
		}
		for k, v := range t1.Ctx().Table {
			if !Equals(v, t2.Ctx().GetNow(k)) {
				return false
			}
		}
		return true
	default:
		return false

	}

}
