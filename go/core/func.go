package core

// import "fmt"

type Func struct {
	Args 		*TokenList
	Codes		*TokenList
	QuoteList	[]int
	Props 		*TokenList
	Desc		[]string
}

func (f Func) Run(es *EvalStack, ctx *BindMap) (*Token, error){
	var c = BindMap{
		Table: 	make(map[string]*Token, 8),
		Father: ctx,
		Tp:		USR_CTX,
	}
	for i, item := range f.Args.List() {
		c.PutNow(item.Str(), es.Line[int(es.LastStartPos()) + i + 1])
	}
	for j:=0; j<f.Props.Len(); j+=2 {
		if f.Props.Get(j+1) == nil {
			c.PutNow(f.Props.Get(j).Str(), &Token{LOGIC, false})
		}else{
			c.PutNow(f.Props.Get(j).Str(), &Token{NONE, "none"})
		}
	}
	return es.Eval(f.Codes.List(), &c)

}

func (f Func) RunWithProps(es *EvalStack, ctx *BindMap, ps []*Token) (*Token, error){
	var fatherCtx = ctx
	if ps[1] != nil && ps[1].Ctx() != nil {
		fatherCtx = ps[1].Ctx()
	}

	var startPos = int(es.LastStartPos())

	var c = BindMap{
		Table: 	make(map[string]*Token, 8),
		Father: fatherCtx,
		Tp:		USR_CTX,
	}
	for i, item := range f.Args.List() {
		c.PutNow(item.Str(), es.Line[startPos + i + 1])
	}

	var nowArgIdx =  startPos + f.Args.Len() + 1

	for i:=2; i<len(ps);i++ {
		for j:=0; j<f.Props.Len(); j+=2 {
			if ps[i].Str() == f.Props.Get(j).Str() {
				if f.Props.Get(j+1) == nil {
					c.PutNow(f.Props.Get(j).Str(), &Token{LOGIC, true})
				}else{
					c.PutNow(f.Props.Get(j+1).Str(), es.Line[nowArgIdx])
					nowArgIdx++
				}
				break
			}
		}
	} 

	for j:=0; j<f.Props.Len(); j+=2 {
		_, ok := c.Table[f.Props.Get(j).Str()]
		if !ok {
			if f.Props.Get(j+1) == nil {
				c.PutNow(f.Props.Get(j).Str(), &Token{LOGIC, false})
			}else{
				c.PutNow(f.Props.Get(j+1).Str(), &Token{NONE, ""})
			}
		}
	}
	
	return es.Eval(f.Codes.List(), &c)

}

func (f Func) GetFuncInfo() string{
	var result = "FUNC: \n"
	result += "  desc:    "
	if len(f.Desc) >= 2 {
		result += f.Desc[1]
	}
	result += "\n\n"
	result +="  args:    "
	for i:=0; i<f.Args.Len(); i++ {
		result += f.Args.Get(i).Str() + "\t\t"
		for j:=0; j<len(f.Desc); j++ {
			if f.Desc[j] == f.Args.Get(i).Str() {
				result += f.Desc[j+1]
			}
		}
		result += "\n           "
	}
	result += "\n"
	result += "  props:   "
	for i:=0; i<f.Props.Len(); i+=2 {
		result += "/" + f.Props.Get(i).Str()
		if i+1 < f.Props.Len() && f.Props.Get(i+1) != nil {
			result += "  " + f.Props.Get(i+1).Str()
		}
		result += "\t"
		for j:=0; j<len(f.Desc); j++ {
			if f.Desc[j] == f.Props.Get(i).Str() {
				result += f.Desc[j+1]
			}
		}
		result += "\n           "
	}
	return result
}
