package core

import "runtime"
import "sync"
import "bytes"
import "fmt"

const (
	SYS_CTX = iota
	USR_CTX
	TMP_CTX
)
	


type BindMap struct{
	Table 	map[string]*Token
	Father 	*BindMap
	Tp 		int
	Lock 	sync.RWMutex
}

func (bm *BindMap)GetNow(key string) *Token{
	var tk *Token
	var ok bool
	if FORKS > 1 {
		bm.Lock.RLock()
		tk, ok = bm.Table[key]
		bm.Lock.RUnlock()
	}else{
		tk, ok = bm.Table[key]
	}
	
	if ok {
		return tk
	}else{
		return &Token{NONE, ""}
	}
}

func (bm *BindMap) Get(key string) *Token{
	
	var ctx = bm
	var prev = ctx

	var tk *Token
	var ok bool
	if(ctx.Table != nil){
		if FORKS > 1 {
			bm.Lock.RLock()
			tk, ok = ctx.Table[key]
			bm.Lock.RUnlock()
		}else{
			tk, ok = ctx.Table[key]
		}
		if(ok){
			return tk
		}
	}

	for !ok && ctx.Father != nil {
		prev = ctx
		ctx = ctx.Father
		
		if(ctx.Table != nil){
			if FORKS > 1 {
				ctx.Lock.RLock()
				tk, ok = ctx.Table[key]
				ctx.Lock.RUnlock()
			}else{
				tk, ok = ctx.Table[key]
			}
		}
		
	}

	if tk != nil {
		if ctx.Father == nil {
			if FORKS > 1 {
				prev.Lock.Lock()
				prev.Table[key] = tk
				prev.Lock.Unlock()
			}else{
				prev.Table[key] = tk
			}
		}
		return tk
	}else{
		return &Token{Tp: NONE}
	}
}


func (bm *BindMap)PutNow(key string, val *Token){
	if FORKS > 1 {
		bm.Lock.Lock()
		bm.Table[key] = val
		bm.Lock.Unlock()
	}else{
		bm.Table[key] = val
	}
}


func (bm *BindMap)Put(key string, val *Token){

	var ctx = bm
	var inserted = false
	var ok = false

	if(ctx.Table != nil){
		if FORKS > 1 {
			ctx.Lock.RLock()
			_, ok = ctx.Table[key]
			ctx.Lock.RUnlock()
		}else{
			_, ok = ctx.Table[key]
		}
	}

	if(ok){
		if FORKS > 1 {
			bm.Lock.Lock()
			bm.Table[key] = val.Clone()
			bm.Lock.Unlock()
		}else{
			bm.Table[key] = val.Clone()
		}
		
		inserted = true
	}else{
		for !inserted && !ok && ctx.Father != nil {
			if(ctx.Table != nil){
				if FORKS > 1 {
					ctx.Lock.RLock()
					_, ok = ctx.Table[key]
					ctx.Lock.RUnlock()
				}else{
					_, ok = ctx.Table[key]
				}
			}
			if(ok){
				if FORKS > 1 {
					ctx.Lock.Lock()
					ctx.Table[key] = val.Clone()
					ctx.Lock.Unlock()
				}else{
					ctx.Table[key] = val.Clone()
				}
				
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
	if FORKS > 1 {
		ctx.Lock.Lock()
		ctx.Table[key] = val.Dup()
		ctx.Lock.Unlock()
	}else{
		ctx.Table[key] = val.Dup()
	}
	
}


func (bm *BindMap)Unset(key string){

	var ctx = bm
	var ok = false

	if(ctx.Table != nil){
		if FORKS > 1 {
			ctx.Lock.RLock()
			_, ok = ctx.Table[key]
			ctx.Lock.RUnlock()
		}else{
			_, ok = ctx.Table[key]
		}
		
	}

	if(ok){
		if FORKS > 1 {
			ctx.Lock.Lock()
			delete(ctx.Table, key)
			ctx.Lock.Unlock()
		}else{
			delete(ctx.Table, key)
		}
		
		runtime.GC()
	}else{
		for !ok && ctx.Father != nil {
			if(ctx.Table != nil){
				if FORKS > 1 {
					ctx.Lock.RLock()
					_, ok = ctx.Table[key]
					ctx.Lock.RUnlock()
				}else{
					_, ok = ctx.Table[key]
				}
				
			}
			if(ok){
				if FORKS > 1 {
					ctx.Lock.Lock()
					delete(ctx.Table, key)
					ctx.Lock.RUnlock()
				}else{
					delete(ctx.Table, key)
				}
				
				runtime.GC()
				break
			}
			ctx = ctx.Father
		}
	}
	
}

func (bm *BindMap) Echo() {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	for k, v := range bm.Table {
		buffer.WriteString(k)
		buffer.WriteString(": ")
		buffer.WriteString(v.ToString())
		buffer.WriteString(" ")
	}
	if len(buffer.Bytes()) > 1 {
		buffer.Bytes()[len(buffer.Bytes())-1] = '}'
	}else{
		buffer.WriteString("}")
	}
	fmt.Println(buffer.String())
}
