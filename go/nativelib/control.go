package nativelib

import (
	"errors"
	"sync"

	. "github.com/NjinN/RML/go/core"
)

// import "fmt"

func Iif(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	var result Token
	if args[1].ToBool() {
		if args[2].Tp == BLOCK {
			es.Eval(args[2].Tks(), ctx)
		} else if args[2].Tp == STRING {
			es.EvalStr(args[2].Str(), ctx)
		}
		return &Token{LOGIC, true}, nil
	} else {
		return &Token{NIL, nil}, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil

}

func Either(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	var result Token
	if args[1].ToBool() {
		if args[2].Tp == BLOCK {
			return es.Eval(args[2].Tks(), ctx)
		} else if args[2].Tp == STRING {
			return es.EvalStr(args[2].Str(), ctx)
		}
	} else {
		if args[3].Tp == BLOCK {
			return es.Eval(args[3].Tks(), ctx)
		} else if args[3].Tp == STRING {
			return es.EvalStr(args[3].Str(), ctx)
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Loop(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == INTEGER && args[2].Tp == BLOCK {
		var rs *Token
		var err error
		for i := 0; i < args[1].Int(); i++ {
			rs, err = es.Eval(args[2].Tks(), ctx)
			if err != nil {
				if err.Error() == "continue" {
					continue
				}
				if err.Error() == "break" {
					break
				}
				return rs, err
			}
		}

		return rs, nil

	}

	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}

func Repeat(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	// args[2].Echo()

	if args[1].Tp == WORD && args[2].Tp == INTEGER && args[3].Tp == BLOCK {
		var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
		var countToken = Token{INTEGER, 1}

		c.PutNow(args[1].Str(), &countToken)
		var rs *Token
		var err error
		for countToken.Int() <= args[2].Int() {
			rs, err = es.Eval(args[3].Tks(), &c)
			countToken.Val = countToken.Int() + 1
			if err != nil {
				if err.Error() == "continue" {
					continue
				}
				if err.Error() == "break" {
					break
				}
				return rs, err
			}
			if rs != nil && rs.Tp == ERR {
				return rs, err
			}
		}
		return &Token{NIL, nil}, nil

	}

	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}

func Ffor(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == WORD && args[5].Tp == BLOCK && (args[2].Tp == INTEGER || args[2].Tp == DECIMAL) && (args[3].Tp == INTEGER || args[3].Tp == DECIMAL) && (args[4].Tp == INTEGER || args[4].Tp == DECIMAL) {
		var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
		var countToken = args[2].Dup()
		c.PutNow(args[1].Str(), countToken)
		var rs *Token
		var err error

		if args[2].Tp == INTEGER && args[4].Tp == INTEGER {
			if args[3].Tp == INTEGER {
				for countToken.Int() <= args[3].Int() {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Int() + args[4].Int()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			} else {
				for countToken.Int() <= int(args[3].Float()) {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Int() + args[4].Int()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}

		} else if args[2].Tp == INTEGER && args[4].Tp == DECIMAL {
			countToken.Tp = DECIMAL
			countToken.Val = float64(countToken.Int())
			if args[3].Tp == INTEGER {
				for countToken.Float() <= float64(args[3].Int()) {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + args[4].Float()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			} else {
				for countToken.Float() <= args[3].Float() {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + args[4].Float()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}
		} else if args[2].Tp == DECIMAL && args[4].Tp == INTEGER {
			if args[3].Tp == INTEGER {
				for countToken.Float() <= float64(args[3].Int()) {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + float64(args[4].Int())
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			} else {
				for countToken.Float() <= args[3].Float() {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + float64(args[4].Int())
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}
		} else if args[2].Tp == DECIMAL && args[4].Tp == DECIMAL {
			if args[3].Tp == INTEGER {
				for countToken.Float() <= float64(args[3].Int()) {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + args[4].Float()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			} else {
				for countToken.Float() <= args[3].Float() {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + args[4].Float()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}
		}

		return &Token{NIL, nil}, nil
	}

	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}

func Wwhile(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == BLOCK && args[2].Tp == BLOCK {
		var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
		b, e1 := es.Eval(args[1].Tks(), &c)
		if e1 != nil {
			return nil, e1
		}
		var rs *Token
		var err error
		for b.Val.(bool) {
			rs, err = es.Eval(args[2].Tks(), &c)
			if err != nil {
				if err.Error() == "continue" {
					continue
				}
				if err.Error() == "break" {
					break
				}
				return rs, err
			}
			if rs != nil && rs.Tp == ERR {
				return rs, err
			}
			b, err = es.Eval(args[1].Tks(), &c)
			if err != nil {
				return rs, err
			}
			if b.Tp == ERR {
				return rs, err
			}
		}
		return &Token{NIL, nil}, nil
	}
	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}

func Bbreak(es *EvalStack, ctx *BindMap) (*Token, error) {
	return nil, errors.New("break")
}

func Ccontinue(es *EvalStack, ctx *BindMap) (*Token, error) {
	return nil, errors.New("continue")
}

func Rreturn(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	return args[1], errors.New("return")
}

func Until(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp ==  BLOCK {
		for {
			rst, err := es.Eval(args[1].Tks(), ctx)

			if err != nil {
				if err.Error() == "break" {
					return &Token{NONE, ""}, nil
				}
				if err.Error() == "continue" {
					continue
				}
				return rst, err
			}

			if rst != nil && rst.ToBool() {
				return rst, err
			}
		}
	}

	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}


func Ttry(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == BLOCK && args[2].Tp == BLOCK {
		temp, err := es.Eval(args[1].Tks(), ctx)
		if (temp != nil && temp.Tp == ERR) || err != nil {
			if err != nil {
				temp.Val = err.Error()
			}

			var c = BindMap{make(map[string]*Token, 4), ctx, TMP_CTX, sync.RWMutex{}}
			temp.Tp = STRING
			c.PutNow("e", temp)

			return es.Eval(args[2].Tks(), &c)
		}

		return temp, err
	}

	return &Token{ERR, "Type Mismatch"}, nil
}

func Cause(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == STRING {
		return &Token{ERR, args[1].Str()}, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}
