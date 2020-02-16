package nativelib

import (
	"math"
	"strconv"

	// "encoding/hex"
	// "fmt"

	. "github.com/NjinN/RML/go/core"
)

func To(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token

	if args[1].Tp == DATATYPE && args[2] != nil {
		if args[1].Uint8() == args[2].Tp {
			return args[2].CloneDeep(), nil
		}

		result.Tp = args[1].Uint8()

		switch args[1].Uint8() {
		case ERR:
			result.Val = args[2].OutputStr()
			return &result, nil

		case LIT_WORD:
			result.Val = args[2].OutputStr()
			return &result, nil

		case GET_WORD:
			result.Val = args[2].OutputStr()
			return &result, nil

		case DATATYPE:
			result.Val = args[2].Tp
			return &result, nil

		case LOGIC:
			result.Val = args[2].ToBool()
			return &result, nil

		case INTEGER:
			switch args[2].Tp {
			case LOGIC:
				if args[2].Val.(bool) {
					result.Val = 1
				} else {
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
				} else {
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
				} else {
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
				} else {
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
			if args[2].Tp == BIN {
				result.Val = string(args[2].Val.([]byte))
				return &result, nil
			}

			result.Val = args[2].OutputStr()
			return &result, nil

		case BLOCK, PAREN, PATH:
			switch args[2].Tp {
			case BLOCK, PAREN, PATH, RANGE:
				result.Val = args[2].CloneDeep().Val
				return &result, nil

			default:
				result.Val = NewTks(8)
				result.List().Add(args[2].CloneDeep())
				return &result, nil
			}

		case WORD:
			result.Val = args[2].OutputStr()
			return &result, nil

		case SET_WORD:
			result.Val = args[2].OutputStr()
			return &result, nil

		case PUT_WORD:
			result.Val = args[2].OutputStr()
			return &result, nil

		case FILE:
			if args[2].Tp == STRING {
				result.Val = args[2].Val
				return &result, nil
			}

		case URL:
			if args[2].Tp == STRING {
				result.Val = args[2].Val
				return &result, nil
			}

		case BIN:
			if args[2].Tp == STRING {
				result.Val = []byte(args[2].Str())
				return &result, nil
			}

		case RANGE:
			if args[2].Tp == BLOCK || args[2].Tp == PAREN || args[3].Tp == PATH {
				if args[2].List().Len() >= 2 {
					if (args[2].Tks()[0].Tp == INTEGER || args[2].Tks()[0].Tp == DECIMAL) && (args[2].Tks()[1].Tp == INTEGER || args[2].Tks()[1].Tp == DECIMAL) {
						result.Val = NewTks(4)
						result.List().AddArr(args[2].Tks()[0:2])
						return &result, nil
					}
				}
			}

		case WRAP:
			result.Val = args[2]
			return &result, nil

		default:

		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}
