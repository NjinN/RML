package core

// import "fmt"

type Func struct {
	Args 		[]*Token
	Codes		[]*Token
	QuoteList	[]int
	Props 		[]*Token
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
