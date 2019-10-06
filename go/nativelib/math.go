package nativelib

import . "../core"

func Add(Es *EvalStack, ctx *BindMap) (*Token, error){
	var args = Es.Line[Es.LastStartPos() : Es.LastEndPos() + 1]

	var result Token
	if args[1].Tp == INTEGER {
		if args[2].Tp == INTEGER {
			result.Tp = INTEGER
			result.Val = args[1].Val.(int) + args[2].Val.(int)
			return &result, nil
		}else if args[2].Tp == DECIMAL {
			result.Tp = DECIMAL
			result.Val = float64(args[1].Val.(int)) + args[2].Val.(float64)
			return &result, nil
		}

	}else if args[1].Tp == DECIMAL {
		if args[2].Tp == INTEGER {
			result.Tp = DECIMAL
			result.Val = args[1].Val.(float64) + float64(args[2].Val.(int))
			return &result, nil
		}else if args[2].Tp == DECIMAL {
			result.Tp = DECIMAL
			result.Val = args[1].Val.(float64) + args[2].Val.(float64)
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}


func Sub(Es *EvalStack, ctx *BindMap) (*Token, error){
	var args = Es.Line[Es.LastStartPos() : Es.LastEndPos() + 1]

	var result Token
	if args[1].Tp == INTEGER {
		if args[2].Tp == INTEGER {
			result.Tp = INTEGER
			result.Val = args[1].Val.(int) - args[2].Val.(int)
			return &result, nil
		}else if args[2].Tp == DECIMAL {
			result.Tp = DECIMAL
			result.Val = float64(args[1].Val.(int)) - args[2].Val.(float64)
			return &result, nil
		}

	}else if args[1].Tp == DECIMAL {
		if args[2].Tp == INTEGER {
			result.Tp = DECIMAL
			result.Val = args[1].Val.(float64) - float64(args[2].Val.(int))
			return &result, nil
		}else if args[2].Tp == DECIMAL {
			result.Tp = DECIMAL
			result.Val = args[1].Val.(float64) - args[2].Val.(float64)
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}


func Mul(Es *EvalStack, ctx *BindMap) (*Token, error){
	var args = Es.Line[Es.LastStartPos() : Es.LastEndPos() + 1]

	var result Token
	if args[1].Tp == INTEGER {
		if args[2].Tp == INTEGER {
			result.Tp = INTEGER
			result.Val = args[1].Val.(int) * args[2].Val.(int)
			return &result, nil
		}else if args[2].Tp == DECIMAL {
			result.Tp = DECIMAL
			result.Val = float64(args[1].Val.(int)) * args[2].Val.(float64)
			return &result, nil
		}

	}else if args[1].Tp == DECIMAL {
		if args[2].Tp == INTEGER {
			result.Tp = DECIMAL
			result.Val = args[1].Val.(float64) * float64(args[2].Val.(int))
			return &result, nil
		}else if args[2].Tp == DECIMAL {
			result.Tp = DECIMAL
			result.Val = args[1].Val.(float64) * args[2].Val.(float64)
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Div(Es *EvalStack, ctx *BindMap) (*Token, error){
	var args = Es.Line[Es.LastStartPos() : Es.LastEndPos() + 1]

	var result Token
	if args[1].Tp == INTEGER {
		if args[2].Tp == INTEGER {
			result.Tp = DECIMAL
			result.Val = float64(args[1].Val.(int)) / float64(args[2].Val.(int))
			return &result, nil
		}else if args[2].Tp == DECIMAL {
			result.Tp = DECIMAL
			result.Val = float64(args[1].Val.(int)) / args[2].Val.(float64)
			return &result, nil
		}

	}else if args[1].Tp == DECIMAL {
		if args[2].Tp == INTEGER {
			result.Tp = DECIMAL
			result.Val = args[1].Val.(float64) / float64(args[2].Val.(int))
			return &result, nil
		}else if args[2].Tp == DECIMAL {
			result.Tp = DECIMAL
			result.Val = args[1].Val.(float64) / args[2].Val.(float64)
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}


func Mod(Es *EvalStack, ctx *BindMap) (*Token, error){
	var args = Es.Line[Es.LastStartPos() : Es.LastEndPos() + 1]

	var result Token
	if args[1].Tp == INTEGER && args[2].Tp == INTEGER {
		result.Tp = INTEGER
		result.Val = args[1].Val.(int) % args[2].Val.(int)
		return &result, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}