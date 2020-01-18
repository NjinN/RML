package nativelib

import . "../core"
import "fmt"
import "strconv"

const MAX_PARSE_DEEP = 500

type Rule struct {
	minTimes 	int
	maxTimes 	int
	ruleStr		string
	ruleBlk		*Token
	code		*Token
	isEnd		bool
	isSkip		bool
	model		string
}

func (r *Rule) init() {
	r.minTimes = -1
	r.maxTimes = -1
	r.ruleStr = ""
	r.ruleBlk = nil
	r.code = nil
	r.isEnd = false
	r.isSkip = false
	r.model = ""
}

func (r *Rule) isRuleComplete() bool {
	return ((r.minTimes >= 0 && r.maxTimes >= r.minTimes) || r.model > "") && (r.ruleStr > "" || (r.ruleBlk != nil && r.ruleBlk.Tp == BLOCK) || r.isEnd || r.isSkip)
}

func (r *Rule) isEmpty() bool {
	return r.minTimes == - 1 && r.maxTimes == -1 && r.ruleStr == "" && r.ruleBlk == nil && r.code == nil
}

func (r *Rule) completeRuleRange(){
	if r.minTimes < 0 {
		r.minTimes = 1
	}
	if r.maxTimes < r.minTimes {
		r.maxTimes = r.minTimes
	}
}

func (r *Rule) Echo(){
	fmt.Println("Rule:")
	fmt.Println("\tminTimes is " + strconv.FormatInt(int64(r.minTimes), 10));
	fmt.Println("\tmaxTimes is " + strconv.FormatInt(int64(r.maxTimes), 10));
	fmt.Println("\truleStr is " + r.ruleStr);
	fmt.Println("\truleBlk is " + r.ruleBlk.ToString());
	fmt.Println("\tcode is " + r.code.ToString());
	fmt.Println("\tisEnd is " + strconv.FormatBool(r.isEnd));
	fmt.Println("\tisSkip is " + strconv.FormatBool(r.isSkip));
	fmt.Println("\tmodel is " + r.model);
}

func (r *Rule) match(str string, nowIdx *int, startDeep *int, es *EvalStack, ctx *BindMap) (bool, *Token){
	var matchTimes = 0
	var mch bool
	var rst *Token
	var err error
	// r.Echo()
	// fmt.Println("nowIdx is " + strconv.FormatInt(int64(*nowIdx), 10));
	if r.isEnd {
		if *nowIdx >= len(str){
			return true, nil
		}else{
			return false, nil
		}
	}else if r.model > "" {
		if r.model == "thru" {
			r.model = ""
			r.minTimes = 1
			r.maxTimes = 1
			for *nowIdx <= len(str){
				mch, rst = r.match(str, nowIdx, startDeep, es, ctx)
				if mch {
					r.init()
					return true, nil
				}
				*nowIdx++
			}
			if !mch {
				return false, nil
			}
		}else if r.model == "to" {
			r.model = ""
			r.minTimes = 1
			r.maxTimes = 1
			var tempIdx = *nowIdx
			for tempIdx <= len(str){
				*nowIdx = tempIdx
				mch, rst = r.match(str, &tempIdx, startDeep, es, ctx)
				if mch {
					r.init()
					return true, nil
				}
				tempIdx++
			}
			if !mch {
				return false, nil
			}
		}
	
	}else if r.ruleStr > "" {
		for matchTimes < r.maxTimes && *nowIdx < len(str) {
			if *nowIdx + len(r.ruleStr) > len(str){
				if matchTimes >= r.minTimes {
					return true, rst
				}else{
					return false, rst
				}
			}else{

				if str[*nowIdx : *nowIdx+len(r.ruleStr)] == r.ruleStr {
					
					if r.code != nil {
						rst, err = es.Eval(r.code.Tks(), ctx)
						if err != nil {
							return false, &Token{ERR, "Error when eval parsing code"}
						}
					}
			
					matchTimes++
					*nowIdx += len(r.ruleStr)
				}else{
					if matchTimes >= r.minTimes {
						return true, rst
					}else{
						return false, rst
					}
				}
			}
		}
		if matchTimes >= r.minTimes {
			return true, rst
		}else{
			return false, rst
		}

	}else if r.ruleBlk != nil && r.ruleBlk.Tp == BLOCK {
		for matchTimes < r.maxTimes && *nowIdx < len(str) {
			mch, rst = matchRuleBlk(str, r.ruleBlk, nowIdx, startDeep, es, ctx)
			if rst != nil && rst.Tp == ERR {
				return false, rst
			}
			if mch {
				if r.code != nil {
					rst, err = es.Eval(r.code.Tks(), ctx)
					if err != nil {
						return false, &Token{ERR, "Error when eval parsing code"}
					}
				}
				matchTimes++
			}
		}
		if matchTimes >= r.minTimes {
			return true, rst
		}else{
			return false, rst
		}

	}
	return false, &Token{ERR, "Error parsing rule"}
}

func matchRuleBlk(str string, blk *Token, nowIdx *int, startDeep *int, es *EvalStack, ctx *BindMap) (bool, *Token){
	*startDeep++
	if *startDeep > MAX_PARSE_DEEP {
		return false, &Token{ERR, "Parse too Deep"}
	}
	var rst *Token
	var mch bool
	var rule Rule
	rule.init()
	if isOrRules(blk){
		var rules = splitOrRules(blk)
		var tempIdx int
		for _, item := range rules.Tks() {
			tempIdx = *nowIdx
			mch, rst = matchRuleBlk(str, item, &tempIdx, startDeep, es, ctx)
			if rst != nil && rst.Tp == ERR {
				*startDeep--
				return false, rst
			}
			if mch {
				*nowIdx = tempIdx
				*startDeep--
				return true, nil
			}
		}
		*startDeep--
		return false, nil

	}else{
		var blkIdx = 0
		for blkIdx < blk.List().Len() {
			var nowRule = blk.Tks()[blkIdx]
			// nowRule.Echo()
			if nowRule.Tp == WORD {
				// nowRule.Echo()
				if nowRule.Str() == "copy" {
					if blkIdx < blk.List().Len() - 1 && blk.Tks()[blkIdx + 1].Tp == WORD {
						var startIdx = *nowIdx
						var word = blk.Tks()[blkIdx + 1].Str()
						blkIdx++
						getNextRule(&rule, blk, &blkIdx)
						if rule.isRuleComplete() {
							mch, rst := rule.match(str, nowIdx, startDeep, es, ctx)
							if !mch || (rst != nil && rst.Tp == ERR){
								if rule.ruleBlk != nil {
									*startDeep--
								}
								return false, rst
							}else{
								copy(str, startIdx, *nowIdx, word, ctx)
							}
							rule.init()
						}else{
							return false, &Token{ERR, "Error parsing rule"}
						}
					}else{
						return false, &Token{ERR, "Error parsing rule"}
					}
					// blkIdx++
				}else if nowRule.Str() == "end" {
					rule.isEnd = true
					rule.completeRuleRange()
					blkIdx++
				}else if nowRule.Str() == "skip" {
					rule.isSkip = true
					rule.completeRuleRange()
					blkIdx++
				}

			}
			
			if !rule.isRuleComplete(){
				getNextRule(&rule, blk, &blkIdx)
			}
			
			if rule.isRuleComplete() {
				// rule.Echo()
				if blkIdx < blk.List().Len() - 1 && blk.Tks()[blkIdx + 1].Tp == PAREN {
					rule.code = blk.Tks()[blkIdx + 1]
					blkIdx++
				}

				mch, rst := rule.match(str, nowIdx, startDeep, es, ctx)
				if !mch || (rst != nil && rst.Tp == ERR){
					if rule.ruleBlk != nil {
						*startDeep--
					}
					return false, rst
				}
				rule.init()
			}else{
				return false, &Token{ERR, "Error parsing rule"}
			}

			// blkIdx++
		}

		if *startDeep == 1 && (*nowIdx != len(str) || !rule.isEmpty()){
			*startDeep--
			return false, nil
		}else{
			*startDeep--
			return true, nil
		}
	}
	
}

func getNextRule(rule *Rule, blk *Token, blkIdx *int){
	for *blkIdx < len(blk.Tks()) {
		var nowRule = blk.Tks()[*blkIdx]
		if nowRule.Tp == INTEGER {
			if rule.minTimes < 0 {
				rule.minTimes = nowRule.Int()
			}else{
				rule.maxTimes = nowRule.Int()
			}
		}else if nowRule.Tp == RANGE {
			if nowRule.Tks()[0].Tp != INTEGER {
				return
			}
			rule.minTimes = nowRule.Tks()[0].Int()
			rule.maxTimes = nowRule.Tks()[1].Int()
		}else if nowRule.Tp == STRING{
			rule.ruleStr = nowRule.Str()
			rule.completeRuleRange()
			*blkIdx++
			return
		}else if nowRule.Tp == BLOCK {
			rule.ruleBlk = nowRule
			rule.completeRuleRange()
			*blkIdx++
			return
		}else if nowRule.Tp == WORD {
			if nowRule.Str() == "thru" {
				rule.model = "thru"
			}else if nowRule.Str() == "to" {
				rule.model = "to"
			}

		}

		*blkIdx++
	}

}


func isOrRules(blk *Token) bool{
	for idx, item := range blk.Tks(){
		if item.Tp == WORD && (item.Str() == "|" || item.Str() == "or") && idx > 0 && idx < len(blk.Tks())-1 {
			return true
		}
	}
	return false
}

func splitOrRules(blk *Token) *Token{
	var result Token
	result.Tp = BLOCK
	result.Val = NewTks(8)

	var startIdx = 0
	for idx, item := range blk.Tks() {
		if item.Tp == WORD && (item.Str() == "|" || item.Str() == "or") {
			var temp Token
			temp.Tp = BLOCK
			temp.Val = NewTks(8)
			temp.List().AddArr(blk.Tks()[startIdx:idx])

			result.List().Add(&temp)
			startIdx = idx + 1
		}
		
		if idx == len(blk.Tks()){
			var temp Token
			temp.Tp = BLOCK
			temp.Val = NewTks(8)
			temp.List().AddArr(blk.Tks()[startIdx:])

			result.List().Add(&temp)
		}
	}
	return &result
}


func Parse(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	
	if args[1].Tp == STRING && args[2].Tp == BLOCK {
		var nowIdx = 0
		var startDeep = 0
		mch, rst := matchRuleBlk(args[1].Str(), args[2], &nowIdx, &startDeep, es, ctx)

		if rst != nil && rst.Tp == ERR {
			return rst, nil
		}

		return &Token{LOGIC, mch}, nil
	}

	
	return &Token{ERR, "Type Mismatch"}, nil
}






func copy(str string, startIdx int, endIdx int, word string, ctx *BindMap){
	ctx.PutLocal(word, &Token{STRING, str[startIdx : endIdx]})
}



