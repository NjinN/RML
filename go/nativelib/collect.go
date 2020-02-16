package nativelib

import (
	"strings"
	"sync"

	. "github.com/NjinN/RML/go/core"
)

// import "fmt"

func Length(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token

	if args[1].Tp == BLOCK || args[1].Tp == PAREN {
		result.Tp = INTEGER
		result.Val = args[1].List().Len()
		return &result, nil
	} else if args[1].Tp == STRING {
		result.Tp = INTEGER
		result.Val = len([]rune(args[1].Str()))
		return &result, nil
	} else if args[1].Tp == BIN {
		result.Tp = INTEGER
		result.Val = len(args[1].Val.([]byte))
		return &result, nil
	}else if args[1].Tp == OBJECT {
		result.Tp = INTEGER
		result.Val = len(args[1].Ctx().Table)
		return &result, nil
	}else if args[1].Tp == MAP {
		result.Tp = INTEGER
		result.Val = len(args[1].Map().Table)
		return &result, nil
	}


	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Append(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token
	if args[1].Tp == BLOCK || args[1].Tp == PAREN {
		if !args[3].ToBool() && (args[2].Tp == BLOCK || args[2].Tp == PAREN) {
			//逐个添加整个block
			args[1].List().AddAll(args[2].List().Clone())
		} else {
			//添加整个Token
			args[1].List().Add(args[2].Clone())
		}
		return args[1], nil
	} else if args[1].Tp == STRING {
		if args[2].Tp == STRING {
			if !args[3].ToBool() {
				args[1].Val = args[1].Str() + args[2].Str()
			} else {
				args[1].Val = args[1].Str() + `"` + args[2].Str() + `"`
			}
			return args[1], nil
		} else if args[2].Tp == CHAR {
			if !args[3].ToBool() {
				args[1].Val = args[1].Str() + string(args[2].Val.(byte))
			} else {
				args[1].Val = args[1].Str() + `'` + string(args[2].Val.(byte)) + `'`
			}
			return args[1], nil
		} else if args[2].Tp == BLOCK {
			if !args[3].ToBool() {
				for _, item := range args[2].Tks() {
					if item.Tp == STRING {
						args[1].Val = args[1].Str() + item.Str()
					} else {
						args[1].Val = args[1].Str() + item.ToString()
					}
				}
			} else {
				args[1].Val = args[1].Str() + args[2].ToString()
			}
			return args[1], nil

		} else {
			args[1].Val = args[1].Str() + args[2].ToString()
			return args[1], nil
		}
	} else if args[1].Tp == OBJECT {
		if args[2].Tp == BLOCK {
			for i := 0; i < args[2].List().Len()-1; i += 2 {
				args[1].Ctx().PutNow(args[2].Tks()[i].ToString(), args[2].Tks()[i+1])
			}
			return args[1], nil
		} else if args[2].Tp == OBJECT {
			for k, v := range args[2].Ctx().Table {
				args[1].Ctx().PutNow(k, v.Clone())
			}
			return args[1], nil
		}
	} else if args[1].Tp == FILE {
		if args[2].Tp == STRING || args[2].Tp == FILE {
			args[1].Val = args[1].Str() + args[2].Str()
			return args[1], nil
		}
		args[1].Val = args[1].Str() + args[2].ToString()
		return args[1], nil
	}else if args[1].Tp == URL {
		if args[2].Tp == STRING || args[2].Tp == URL {
			args[1].Val = args[1].Str() + args[2].Str()
			return args[1], nil
		}
		args[1].Val = args[1].Str() + args[2].ToString()
		return args[1], nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Insert(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token

	if args[3].Tp != INTEGER {
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return &result, nil
	} else if args[3].Int() <= 0 {
		result.Tp = ERR
		result.Val = "Bound Overflow"
		return &result, nil
	}

	var idx = args[3].Int() - 1

	if args[1].Tp == BLOCK || args[1].Tp == PAREN {
		if !args[4].ToBool() && (args[2].Tp == BLOCK || args[2].Tp == PAREN) {
			//逐个添加Block
			if idx <= args[1].List().Len() {
				args[1].List().InsertAll(idx, args[2].List())
				// temp := make([]*Token, len(args[1].Tks()) + len(args[2].Clone().Tks()))
				// for i := 0; i<idx; i++ {
				// 	temp[i] = args[1].Tks()[i]
				// }
				// for i:=0; i<len(args[2].Tks()); i++ {
				// 	temp[idx+i] = args[2].Clone().Tks()[i]
				// }
				// for i:=idx; i<len(args[1].Tks()); i++ {
				// 	temp[len(args[2].Clone().Tks()) + i] = args[1].Tks()[i]
				// }
				// args[1].Val = temp
				// return args[1].Clone(), nil
			} else {
				for args[1].List().Len() < idx {
					args[1].List().Add(&Token{NONE, "none"})
					// args[1].Val = append(args[1].Tks(), &Token{NONE, "none"})
				}
				args[1].List().InsertAll(idx, args[2].List())
				// args[1].Val = append(args[1].Tks(), args[2].Clone().Tks()...)
				return args[1], nil
			}
		} else {
			if idx <= args[1].List().Len() {
				args[1].List().Insert(idx, args[2].Clone())
				// temp := make([]*Token, len(args[1].Tks()) + 1)
				// for i := 0; i<idx; i++ {
				// 	temp[i] = args[1].Tks()[i]
				// }
				// temp[idx] = args[2].Clone()
				// for i:=idx; i<len(args[1].Tks()); i++ {
				// 	temp[i+1] = args[1].Tks()[i]
				// }
				// args[1].Val = temp
				return args[1], nil
			} else {
				for i := args[1].List().Len(); i < idx; i++ {
					args[1].List().Add(&Token{NONE, "none"})
					// args[1].Val = append(args[1].Tks(), &Token{NONE, "none"})
				}
				args[1].List().Add(args[2].Clone())
				// args[1].Val = append(args[1].Tks(), args[2].Clone())
				return args[1], nil
			}
		}
		return args[1].Clone(), nil
	} else if args[1].Tp == STRING {
		if args[2].Tp == STRING {
			if !args[4].ToBool() {
				if idx <= len(args[1].Str()) {
					args[1].Val = args[1].Str()[0:idx] + args[2].Str() + args[1].Str()[idx:]
				} else {
					for len(args[1].Str()) < idx {
						args[1].Val = args[1].Str() + " "
					}
					args[1].Val = args[1].Str() + args[2].Str()
				}

			} else {
				if idx <= len(args[1].Str()) {
					args[1].Val = args[1].Str()[0:idx] + `"` + args[2].Str() + `"` + args[1].Str()[idx:]
				} else {
					for len(args[1].Str()) < idx {
						args[1].Val = args[1].Str() + " "
					}
					args[1].Val = args[1].Str() + `"` + args[2].Str() + `"`
				}
			}
			return args[1], nil
		} else if args[2].Tp == CHAR {
			if !args[4].ToBool() {
				if idx <= len(args[1].Str()) {
					args[1].Val = args[1].Str()[0:idx] + string(args[2].Val.(byte)) + args[1].Str()[idx:]
				} else {
					for len(args[1].Str()) < idx {
						args[1].Val = args[1].Str() + " "
					}
					args[1].Val = args[1].Str() + string(args[2].Val.(byte))
				}

			} else {
				if idx <= len(args[1].Str()) {
					args[1].Val = args[1].Str()[0:idx] + `'` + string(args[2].Val.(byte)) + `'` + args[1].Str()[idx:]
				} else {
					for len(args[1].Str()) < idx {
						args[1].Val = args[1].Str() + " "
					}
					args[1].Val = args[1].Str() + `'` + string(args[2].Val.(byte)) + `'`
				}
			}
			return args[1], nil
		} else {
			args[1].Val = args[1].Str() + args[2].ToString()
			return args[1], nil
		}
	} else if args[1].Tp == OBJECT {
		if args[2].Tp == BLOCK {
			for i := 0; i < args[2].List().Len()-1; i += 2 {
				args[1].Ctx().PutNow(args[2].Tks()[i].ToString(), args[2].Tks()[i+1])
			}
			return args[1], nil
		} else if args[2].Tp == OBJECT {
			for k, v := range args[2].Ctx().Table {
				args[1].Ctx().PutNow(k, v.Clone())
			}
			return args[1], nil
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Take(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token

	if args[1].Tp == BLOCK && args[2].Tp == INTEGER && args[3].Tp == INTEGER {

		if args[3].Int() > 1 {
			//取多个Token
			result.Tp = BLOCK
			var starIdx = args[2].Int() - 1
			var endIdx = starIdx + args[3].Int()

			if starIdx < 0 {
				starIdx = 0
			}
			if endIdx > args[1].List().Len() {
				endIdx = args[1].List().Len()
			}

			if starIdx < endIdx {
				result.Val = args[1].Tks()[starIdx:endIdx]
			} else {
				endIdx = starIdx
			}

			// if args[4].Tp == LOGIC && args[4].Val.(bool){
			// 	args[1].Val = append(args[1].Tks()[0:starIdx], args[1].Tks()[endIdx:]...)
			// }
			result.Val = args[1].List().Take(starIdx, endIdx)
			return &result, nil

		} else if args[3].Int() == 1 {
			//取单个Token
			var idx = args[2].Int() - 1
			if idx < 0 || idx >= args[1].List().Len() {
				result.Tp = NONE
				result.Val = "none"
			} else {
				result.Tp = args[1].List().Get(idx).Tp
				result.Val = args[1].List().Get(idx).Val
			}

			if args[4].Tp == LOGIC && args[4].Val.(bool) {
				args[1].Val = append(args[1].Tks()[0:idx], args[1].Tks()[idx+1:]...)
			}
			return &result, nil
		}
	} else if args[1].Tp == STRING && args[2].Tp == INTEGER && args[3].Tp == INTEGER {
		if args[3].Int() > 1 {
			result.Tp = STRING
			result.Val = ""
			var starIdx = args[2].Int() - 1
			var endIdx = starIdx + args[3].Int()

			if starIdx < 0 {
				starIdx = 0
			}
			if endIdx > len(args[1].Str()) {
				endIdx = len(args[1].Str())
			}
			if starIdx < endIdx {
				result.Val = args[1].Str()[starIdx:endIdx]
			}

			args[1].Val = args[1].Str()[0:starIdx] + args[1].Str()[endIdx:]

			return &result, nil

		} else if args[3].Int() == 1 {
			var idx = args[2].Int() - 1
			runes := []rune(args[1].Str())
			if idx < 0 || idx >= len(runes) {
				result.Tp = NONE
				result.Val = "none"
			} else {
				result.Tp = CHAR
				result.Val = runes[idx]
			}

			args[1].Val = args[1].Str()[0:idx] + args[1].Str()[idx+1:]

			return &result, nil
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Replace(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token

	if args[4].Tp != INTEGER || args[5].Tp != INTEGER {
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return &result, nil
	}
	var at = args[4].Int() - 1
	var amount = args[5].Int()

	if args[1].Tp == STRING || args[1].Tp == URL {
		var old = ""
		if args[2].Tp == STRING {
			old = args[2].Str()
		} else {
			old = args[2].ToString()
		}
		var new = ""
		if args[3].Tp == STRING {
			new = args[3].Str()
		} else {
			new = args[3].ToString()
		}

		args[1].Val = args[1].Str()[0:at] + strings.Replace(args[1].Str()[at:], old, new, amount)
		return args[1], nil
	} else if args[1].Tp == BLOCK || args[1].Tp == PAREN {
		if amount < 0 {
			amount = int(^uint(0) >> 1) //取有符号int最大值
		}

		if args[2].Tp == BLOCK || args[2].Tp == PAREN {
			// for i := at; i < len(args[1].Tks()) - len(args[2].Tks()); i++ {
			// 	var match = true
			// 	for j := 0; j < len(args[2].Tks()); j++{
			// 		if args[2].Tks()[j].Tp != args[1].Tks()[i+j].Tp || args[2].Tks()[j].Val != args[1].Tks()[i+j].Val {
			// 			match = false
			// 			break
			// 		}
			// 	}
			// 	if match && amount > 0 {

			// 		var temp = make([]*Token, 0)
			// 		if args[3].Tp == BLOCK || args[3].Tp == PAREN {
			// 			for n := 0; n < i; n++ {
			// 				temp = append(temp, args[1].Tks()[n])
			// 			}
			// 			for n := 0; n < len(args[3].Tks()); n++ {
			// 				temp = append(temp, args[3].Tks()[n])
			// 			}
			// 			for n := i + len(args[2].Tks()); n < len(args[1].Tks()); n++ {
			// 				temp = append(temp, args[1].Tks()[n])
			// 			}
			// 		}else{
			// 			for n := 0; n < i; n++ {
			// 				temp = append(temp, args[1].Tks()[n])
			// 			}

			// 			temp = append(temp, args[3])

			// 			for n := i + len(args[2].Tks()); n < len(args[1].Tks()); n++ {
			// 				temp = append(temp, args[1].Tks()[n])
			// 			}
			// 		}

			// 		args[1].Val = temp
			// 	}
			// }
			if args[3].Tp == BLOCK || args[3].Tp == PAREN {
				args[1].List().ReplacePart(args[2].List(), args[3].List(), at, amount)
			} else {
				args[1].List().ReplacePartToOne(args[2].List(), args[3], at, amount)
			}

			return args[1], nil

		} else {
			// for i := at; i < len(args[1].Tks()) && amount > 0; i++ {
			// 	var item = args[1].Tks()[i]
			// 	if item.Tp == args[2].Tp && item.Val == args[2].Val {
			// 		args[1] = args[3].Clone()
			// 		amount--
			// 	}
			// }
			if args[3].Tp == BLOCK || args[3].Tp == PAREN {
				args[1].List().ReplaceOneToPart(args[2], args[3].List(), at, amount)
			} else {
				args[1].List().Replace(args[2], args[3], at, amount)
			}
			return args[1], nil
		}

	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Gget(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token

	if args[1].Tp == OBJECT {
		if args[2].Tp == WORD || args[2].Tp == STRING {
			v := args[1].Ctx().GetNow(args[2].Str())
			return v, nil

		}
	} else if args[1].Tp == MAP {
		return args[1].Map().Get(args[2]), nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Pput(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token

	if args[1].Tp == OBJECT {
		if args[2].Tp == WORD || args[2].Tp == STRING {
			args[1].Ctx().PutNow(args[2].Str(), args[3])
			return args[1], nil
		}
	} else if args[1].Tp == MAP {
		args[1].Map().Put(args[2], args[3])
		return args[1], nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}


func Fforeach(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	if args[3].Tp != BLOCK && args[3].Tp != STRING {
		var result = Token{ERR, "Type Mismatch"}
		return &result, nil
	}

	if args[1].Tp == WORD {
		if args[2].Tp == BLOCK || args[2].Tp == PAREN || args[2].Tp == PATH {
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for i := 0; i < args[2].List().Len(); i++ {
				c.PutNow(args[1].Str(), args[2].Tks()[i])
				if args[3].Tp == BLOCK {
					temp, err := es.Eval(args[3].Tks(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}

				} else if args[3].Tp == STRING {
					temp, err := es.EvalStr(args[3].Str(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}
				}

			}
			return &Token{NIL, nil}, nil

		} else if args[2].Tp == OBJECT {
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for k, v := range args[2].Ctx().Table {
				var blk = NewTks(4)
				blk.AddArr([]*Token{&Token{SET_WORD, k}, v})
				c.PutNow(args[1].Str(), &Token{BLOCK, blk})
				if args[3].Tp == BLOCK {
					temp, err := es.Eval(args[3].Tks(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}

				} else if args[3].Tp == STRING {
					temp, err := es.EvalStr(args[3].Str(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}
				}
			}
			return &Token{NIL, nil}, nil
		} else if args[2].Tp == MAP {
			if len(args[2].Table()) == 0 {
				return &Token{NIL, nil}, nil
			}
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for _, v := range args[2].Table() {
				var blk = NewTks(4)
				blk.AddArr([]*Token{v.Key.CloneDeep(), v.Val})
				c.PutNow(args[1].Str(), &Token{BLOCK, blk})
				if args[3].Tp == BLOCK {
					temp, err := es.Eval(args[3].Tks(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}

				} else if args[3].Tp == STRING {
					temp, err := es.EvalStr(args[3].Str(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}
				}
			}
			return &Token{NIL, nil}, nil
		}

	} else if args[1].Tp == BLOCK {
		for _, item := range args[1].Tks() {
			if item.Tp != WORD {
				var result = Token{ERR, "Type Mismatch"}
				return &result, nil
			}
		}

		if args[2].Tp == BLOCK {
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for i := 0; i < args[2].List().Len(); i += args[1].List().Len() {
				for j := 0; j < args[1].List().Len(); j++ {
					if i+j < args[2].List().Len() {
						c.PutNow(args[1].Tks()[j].Str(), args[2].Tks()[i+j])
					} else {
						c.PutNow(args[1].Tks()[j].Str(), &Token{NONE, "none"})
					}
				}
				temp, err := es.Eval(args[3].Tks(), &c)
				if err != nil {
					if err.Error() == "continue" {
						continue
					}
					if err.Error() == "break" {
						break
					}
					return temp, err
				}
				if temp != nil && temp.Tp == ERR {
					return temp, err
				}
			}

			return &Token{NIL, nil}, nil
		} else if args[2].Tp == OBJECT {
			if args[1].List().Len() < 2 || args[1].Tks()[0].Tp != WORD || args[1].Tks()[1].Tp != WORD {
				var result = Token{ERR, "Type Mismatch"}
				return &result, nil
			}
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for k, v := range args[2].Ctx().Table {
				c.PutNow(args[1].Tks()[0].Str(), &Token{WORD, k})
				c.PutNow(args[1].Tks()[1].Str(), v)
				temp, err := es.Eval(args[3].Tks(), &c)
				if err != nil {
					if err.Error() == "continue" {
						continue
					}
					if err.Error() == "break" {
						break
					}
					return temp, err
				}
				if temp != nil && temp.Tp == ERR {
					return temp, err
				}
			}
			return &Token{NIL, nil}, nil
		} else if args[2].Tp == MAP {
			if args[1].List().Len() < 2 || args[1].Tks()[0].Tp != WORD || args[1].Tks()[1].Tp != WORD {
				var result = Token{ERR, "Type Mismatch"}
				return &result, nil
			}
			if len(args[2].Table()) == 0 {
				return &Token{NIL, nil}, nil
			}
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for _, v := range args[2].Table() {
				c.PutNow(args[1].Tks()[0].Str(), v.Key.CloneDeep())
				c.PutNow(args[1].Tks()[1].Str(), v.Val)
				temp, err := es.Eval(args[3].Tks(), &c)
				if err != nil {
					if err.Error() == "continue" {
						continue
					}
					if err.Error() == "break" {
						break
					}
					return temp, err
				}
				if temp != nil && temp.Tp == ERR {
					return temp, err
				}
			}

			return &Token{NIL, nil}, nil
		}

	}

	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}
