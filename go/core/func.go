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
		c.PutNow(item.Val.(string), es.Line[int(es.LastStartPos()) + i + 1])
	}
	for j:=0; j<len(f.Props); j+=2 {
		if f.Props[j+1] == nil {
			c.PutNow(f.Props[j].Val.(string), &Token{LOGIC, false})
		}else{
			c.PutNow(f.Props[j].Val.(string), &Token{NONE, "none"})
		}
	}
	return es.Eval(f.Codes, &c)

}

func (f Func) RunWithProps(es *EvalStack, ctx *BindMap, ps []*Token) (*Token, error){
	var fatherCtx = ctx
	if ps[1] != nil && ps[1].Val.(*BindMap) != nil {
		fatherCtx = ps[1].Val.(*BindMap)
	}

	var c = BindMap{
		Table: 	make(map[string]*Token, 8),
		Father: fatherCtx,
		Tp:		USR_CTX,
	}
	for i, item := range f.Args {
		c.PutNow(item.Val.(string), es.Line[int(es.LastStartPos()) + i + 1])
	}
	for j:=0; j<len(f.Props); j+=2 {
		if f.Props[j+1] == nil {
			var set = false
			for i:=2; i<len(ps); i++ {
				if ps[i].Val.(string) == f.Props[j].Val.(string) {
					c.PutNow(f.Props[j].Val.(string), &Token{LOGIC, true})
					set = true
					break
				}
			}
			if !set {
				c.PutNow(f.Props[j].Val.(string), &Token{LOGIC, false})
			}
		}else{
			var set = false
			for i:=1; i<len(ps); i++ {
				if ps[i].Val.(string) == f.Props[j].Val.(string) {
					c.PutNow(f.Props[j].Val.(string), es.Line[int(es.LastStartPos()) + len(f.Args) + i])
					set = true
					break
				}
			}
			if !set {
				c.PutNow(f.Props[j].Val.(string), &Token{NONE, "none"})
			}
		}
	}
	return es.Eval(f.Codes, &c)

}
