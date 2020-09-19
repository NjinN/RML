package core


type Mop struct {
	Str 	string
	Explen	int
	Exec 	func(stack *EvalStack, ctx *BindMap) (*Token, error)
	QuoteList []int
}





