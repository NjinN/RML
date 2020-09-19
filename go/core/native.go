package core



type Native struct {
	Str 	string
	Explen	int
	Exec 	func(stack *EvalStack, ctx *BindMap) (*Token, error)
 	QuoteList []int
}



func (nv Native) RunWithCtx(es *EvalStack, ctx *BindMap, ps []*Token) (*Token, error){
	var fatherCtx = ctx
	if ps[1] != nil && ps[1].Ctx() != nil {
		fatherCtx = ps[1].Ctx()
	}

	return nv.Exec(es, fatherCtx)

}
