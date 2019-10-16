package core

// import "fmt"

type Func struct {
	Args 		[]*Token
	Codes		[]*Token
	QuoteList	[]int
	Ctx 		BindMap
}

func (f Func) Run(stack *EvalStack, ctx *BindMap) (*Token, error){
	var c = BindMap{
		Table: make(map[string]*Token, 8),
		Father: stack.MainCtx,
	}
	for i, item := range f.Args {
		c.PutNow(item.Val.(string), stack.Line[int(stack.LastStartPos()) + i + 1])
	}
	return stack.Eval(f.Codes, &c)

}
