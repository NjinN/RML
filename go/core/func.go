package core

// import "fmt"

type Func struct {
	Args 		[]*Token
	Codes		[]*Token
	QuoteList	[]int
	Props 		[]*Token
	Desc		[]string
}

func (f Func) Run(es *EvalStack, ctx *BindMap) (*Token, error){
	var c = BindMap{
		Table: 	make(map[string]*Token, 8),
		Father: ctx,
		Tp:		USR_CTX,
	}
	for i, item := range f.Args {
		c.PutNow(item.Str(), es.Line[int(es.LastStartPos()) + i + 1])
	}
	for j:=0; j<len(f.Props); j+=2 {
		if f.Props[j+1] == nil {
			c.PutNow(f.Props[j].Str(), &Token{LOGIC, false})
		}else{
			c.PutNow(f.Props[j].Str(), &Token{NONE, "none"})
		}
	}
	return es.Eval(f.Codes, &c)

}

func (f Func) RunWithProps(es *EvalStack, ctx *BindMap, ps []*Token) (*Token, error){
	var fatherCtx = ctx
	if ps[1] != nil && ps[1].Ctx() != nil {
		fatherCtx = ps[1].Ctx()
	}

	var c = BindMap{
		Table: 	make(map[string]*Token, 8),
		Father: fatherCtx,
		Tp:		USR_CTX,
	}
	for i, item := range f.Args {
		c.PutNow(item.Str(), es.Line[int(es.LastStartPos()) + i + 1])
	}
	for j:=0; j<len(f.Props); j+=2 {
		if f.Props[j+1] == nil {
			var set = false
			for i:=2; i<len(ps); i++ {
				if ps[i].Str() == f.Props[j].Str() {
					c.PutNow(f.Props[j].Str(), &Token{LOGIC, true})
					set = true
					break
				}
			}
			if !set {
				c.PutNow(f.Props[j].Str(), &Token{LOGIC, false})
			}
		}else{
			var set = false
			for i:=2; i<len(ps); i++ {
				if ps[i].Str() == f.Props[j].Str() {
					c.PutNow(f.Props[j+1].Str(), es.Line[int(es.LastStartPos()) + len(f.Args) + i - 1])
					set = true
					break
				}
			}
			if !set {
				c.PutNow(f.Props[j+1].Str(), &Token{NONE, "none"})
			}
		}
	}
	return es.Eval(f.Codes, &c)

}

func (f Func) GetFuncInfo() string{
	var result = "FUNC: \n"
	result += "  desc:    "
	if len(f.Desc) >= 2 {
		result += f.Desc[1]
	}
	result += "\n\n"
	result +="  args:    "
	for i:=0; i<len(f.Args); i++ {
		result += f.Args[i].Str() + "\t\t"
		for j:=0; j<len(f.Desc); j++ {
			if f.Desc[j] == f.Args[i].Str() {
				result += f.Desc[j+1]
			}
		}
		result += "\n           "
	}
	result += "\n"
	result += "  props:   "
	for i:=0; i<len(f.Props); i+=2 {
		result += "/" + f.Props[i].Str()
		if i+1 < len(f.Props) && f.Props[i+1] != nil {
			result += "  " + f.Props[i+1].Str()
		}
		result += "\t"
		for j:=0; j<len(f.Desc); j++ {
			if f.Desc[j] == f.Props[i].Str() {
				result += f.Desc[j+1]
			}
		}
		result += "\n           "
	}
	return result
}
