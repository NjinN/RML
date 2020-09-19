package nativelib

import (
	"fmt"
	"time"

	. "github.com/NjinN/RML/go/core"
)

func Cost(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token

	if args[1].Tp != BLOCK {
		result.Tp = ERR
		result.Val = "Type Mismatch"
		return &result, nil
	}

	var start = time.Now()
	es.Eval(args[1].Tks(), ctx)
	var end = time.Now()
	fmt.Printf("cost time: %s\n", end.Sub(start))

	return &result, nil
}

func Nnow(es *EvalStack, ctx *BindMap) (*Token, error) {
	var result Token
	result.Tp = TIME
	var tc = TimeClock{}

	var t = time.Now()
	
	tc.Date = DateToDays(t.Year(), int(t.Month()), t.Day())

	tc.Second = t.Hour() * 60 * 60 + t.Minute() * 60 + t.Second()
	
	tc.FloatSecond = float64(t.Nanosecond()) / float64(Exponent(10, 9))
	result.Val = &tc

	return &result, nil
}


func Ssleep(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	// var result Token

	if args[1].Tp == INTEGER {
		
		time.Sleep(time.Duration(args[1].Int()) * time.Second)
		return &Token{NONE, ""}, nil

	}else if args[1].Tp == DECIMAL {
		time.Sleep(time.Duration(args[1].Float() * 1000000) * time.Microsecond)
		return &Token{NONE, ""}, nil
	}


	return &Token{ERR, "Type Mismatch"}, nil
}


func Ttimer(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	var result Token
	var timer Timer
	var tm = 0.0

	if args[2].Tp == BLOCK && (args[1].Tp == INTEGER || args[1].Tp == DECIMAL) {
		timer.Code = args[2].CloneDeep().List()
		if args[1].Tp == INTEGER {
			tm = float64(args[1].Int())
		}else {
			tm = args[1].Float()
		}

		timer.Time = tm

		result.Tp = TIMER
		result.Val = &timer
		return &result, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}


func Sstart(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	// var result Token

	if args[1].Tp == TIMER {

		if args[1].Timer().Ticker != nil {
			args[1].Timer().Ticker.Stop()
		}

		go func() {
			ticker := time.NewTicker(time.Second * time.Duration(args[1].Timer().Time))
			args[1].Timer().Ticker = ticker
			for {
				select {
				case <-ticker.C:
					ForkEval(args[1].Timer().Code.CloneDeep().List(), ctx, nil, false, nil, 1024 * 10)
				}
			}
		}()
		
		return &Token{LOGIC, true}, nil

	}


	return &Token{ERR, "Type Mismatch"}, nil
}


func Sstop(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
	// var result Token

	if args[1].Tp == TIMER {
		ticker := args[1].Timer().Ticker
		ticker.Stop()
		
		return &Token{LOGIC, true}, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}



