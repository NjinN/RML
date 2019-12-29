package nativelib

import "math"
import "strconv"
// import "fmt"

import . "../core"

func To(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1].Tp == DATATYPE && args[2] != nil {
		if args[1].Int() == args[2].Tp {
			return args[2].CloneDeep(), nil
		}


		result.Tp = args[1].Int()

		switch args[1].Val {
		case ERR:
			result.Val = args[2].OutputStr()
			return &result, nil

		case LIT_WORD:
			result.Val = "'" + args[2].OutputStr()
			return &result, nil

		case GET_WORD:
			result.Val = ":" + args[2].OutputStr()
			return &result, nil

		case DATATYPE:
			result.Val = args[2].OutputStr() + "!"
			return &result, nil

		case LOGIC:
			result.Val = args[2].ToBool()
			return &result, nil

		case INTEGER:
			switch args[2].Tp {
			case LOGIC:
				if args[2].Val.(bool) {
					result.Val = 1
				}else{
					result.Val = 0
				}
				return &result, nil

			case DECIMAL:
				result.Val = int(math.Round(args[2].Float()))
				return &result, nil

			case CHAR:
				result.Val = int(args[2].Val.(byte))
				return &result, nil

			case STRING:
				i, err := strconv.Atoi(args[2].Str())
				if err != nil {
					result.Val = 0
				}else{
					result.Val = i
				}
				return &result, nil

			default:
				result.Val = 0
				return &result, nil
			}
			
		case DECIMAL:
			switch args[2].Tp {
			case LOGIC:
				if args[2].Val.(bool) {
					result.Val = 1.0
				}else{
					result.Val = 0.0
				}
				return &result, nil

			case INTEGER:
				result.Val = float64(args[2].Int())
				return &result, nil

			case CHAR:
				result.Val = float64(args[2].Val.(byte))
				return &result, nil

			case STRING:
				f, err := strconv.ParseFloat(args[2].Str(), 64)
				if err != nil {
					result.Val = 0.0
				}else{
					result.Val = f
				}
				return &result, nil

			default:
				result.Val = 0.0
				return &result, nil
			}
		

		case CHAR:
			switch args[2].Tp {
			case INTEGER:
				result.Val = byte(args[2].Int() % 256)
				return &result, nil

			case DECIMAL:
				result.Val = byte(args[2].Float())
				return &result, nil
				
			default:
				result.Val = byte(0x00)
				return &result, nil
			}

		case STRING:
			result.Val = args[2].OutputStr()
			return &result, nil

		case BLOCK, PAREN, PATH:
			switch args[2].Tp {
			case BLOCK, PAREN, PATH:
				result.Val = args[2].CloneDeep().Val
				return &result, nil

			default:
				result.Val = make([]*Token, 1)
				result.Tks()[0] = args[2].CloneDeep()
				return &result, nil
			}

		case WORD:
			result.Val = args[2].OutputStr()
			return &result, nil
			
		case SET_WORD:
			result.Val = args[2].OutputStr() + ":"
			return &result, nil

		case PUT_WORD:
			result.Val = args[2].OutputStr() + ":="
			return &result, nil

		default:

		}
	}


	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

