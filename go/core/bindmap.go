package core

import "runtime"

const (
	SYS_CTX = iota
	USR_CTX
	TMP_CTX
)
	


type BindMap struct{
	Table 	map[string]*Token
	Father 	*BindMap
	Tp 		int
}


func (bm *BindMap) Get(key string) *Token{
	var ctx = bm
	var prev = ctx

	var tk *Token
	var ok bool
	if(ctx.Table != nil){
		tk, ok = ctx.Table[key]
		if(ok){
			return tk
		}
	}

	for !ok && ctx.Father != nil {
		prev = ctx
		ctx = ctx.Father

		if(ctx.Table != nil){
			tk, ok = ctx.Table[key]
		}
	}

	if tk != nil {
		if ctx.Father == nil {
			prev.Table[key] = tk
		}
		return tk
	}else{
		return &Token{Tp: NONE}
	}
}


func (bm *BindMap)PutNow(key string, val *Token){
	bm.Table[key] = val
}


func (bm *BindMap)Put(key string, val *Token){
	var ctx = bm
	var inserted = false
	var ok = false

	if(ctx.Table != nil){
		_, ok = ctx.Table[key]
	}

	if(ok){
		bm.Table[key] = val.Clone()
		inserted = true
	}else{
		for !inserted && !ok && ctx.Father != nil {
			if(ctx.Table != nil){
				_, ok = ctx.Table[key]
			}
			if(ok){
				ctx.Table[key] = val.Clone()
				inserted = true
				break
			}
			ctx = ctx.Father
		}
	}
	if(!inserted){
		bm.PutLocal(key, val)
	}
}


func (bm *BindMap)PutLocal(key string, val *Token){
	var ctx = bm

	for ctx.Tp != USR_CTX && ctx.Father != nil {
		ctx = ctx.Father
	}

	ctx.Table[key] = val.Dup()
}


func (bm *BindMap)Unset(key string){
	var ctx = bm
	var ok = false

	if(ctx.Table != nil){
		_, ok = ctx.Table[key]
	}

	if(ok){
		delete(ctx.Table, key)
		runtime.GC()
	}else{
		for !ok && ctx.Father != nil {
			if(ctx.Table != nil){
				_, ok = ctx.Table[key]
			}
			if(ok){
				delete(ctx.Table, key)
				runtime.GC()
				break
			}
			ctx = ctx.Father
		}
	}
	
}
