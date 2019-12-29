package nativelib

import . "../core"


func Eq(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
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
	default:
		result.Tp = LOGIC
		result.Val = "Type Mismatch"
	}
	
	return &result, nil
}


func Gt(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result = Token{Tp: LOGIC}

	switch args[1].Tp {
	case NONE:
		result.Val = false
		return &result, nil
	case LOGIC:
		if args[2].Tp == LOGIC {
			if(args[1].Val.(bool) == true && args[2].Val.(bool) == false){
				result.Val = true
			}else{
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


func Lt(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result = Token{Tp: LOGIC}

	switch args[1].Tp {
	case NONE:
		result.Val = false
		return &result, nil
	case LOGIC:
		if args[2].Tp == LOGIC {
			if(args[1].Val.(bool) == false && args[2].Val.(bool) == true){
				result.Val = true
			}else{
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


func Ge(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result = Token{Tp: LOGIC}

	switch args[1].Tp {
	case NONE:
		result.Val = false
		return &result, nil
	case LOGIC:
		if args[2].Tp == LOGIC {
			if((args[1].Val.(bool) == true && args[2].Val.(bool) == false) || (args[1].Val.(bool) == args[2].Val.(bool)) ){
				result.Val = true
			}else{
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


func Le(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result = Token{Tp: LOGIC}

	switch args[1].Tp {
	case NONE:
		result.Val = false
		return &result, nil
	case LOGIC:
		if args[2].Tp == LOGIC {
			if((args[1].Val.(bool) == false && args[2].Val.(bool) == true) || (args[1].Val.(bool) == args[2].Val.(bool)) ){
				result.Val = true
			}else{
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

