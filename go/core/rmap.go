package core

import "bytes"
import "sync"
// import "fmt"

type TokenPair struct {
	Key 	*Token
	Val		*Token
}

type Rmap struct {
	Table 	map[string]TokenPair
	Lock 	sync.RWMutex
}

func (r *Rmap) ToString() string{
	var buffer bytes.Buffer
	buffer.WriteString("!map{")
	for _, v := range r.Table {
		buffer.WriteString("[")
		buffer.WriteString(v.Key.ToString())
		buffer.WriteString(" ")
		buffer.WriteString(v.Val.ToString())
		buffer.WriteString("]")
		buffer.WriteString(" ")
	}
	if len(buffer.Bytes()) == 5 {
		buffer.WriteString("}")
	}else{
		buffer.Bytes()[len(buffer.Bytes())-1] = '}'
	}
	
	return buffer.String()
}

func (r *Rmap) Get(key *Token) *Token {
	var keyString = TypeToStr(key.Tp) + key.ToString()
	r.Lock.RLock()
	pair, ok := r.Table[keyString]
	r.Lock.RUnlock()
	if ok {
		return pair.Val
	}else{
		return &Token{NONE, "none"}
	}
}

func (r *Rmap) Put(key *Token, val *Token) {
	var keyString = TypeToStr(key.Tp) + key.ToString()
	var pair TokenPair
	pair.Key = key.CloneDeep()
	pair.Val = val.Clone()
	r.Lock.Lock()
	r.Table[keyString] = pair
	r.Lock.Unlock()
}

func (r *Rmap) Delete(key *Token) {
	var keyString = TypeToStr(key.Tp) + key.ToString()
	delete(r.Table, keyString)
}


func (r *Rmap) Clone() *Rmap {
	var result Rmap
	result.Table = make(map[string]TokenPair, 8)

	for k, v := range r.Table {
		var entity TokenPair
		entity.Key = v.Key.Clone()
		entity.Val = v.Val.Clone()
		r.Lock.Lock()
		result.Table[k] = entity
		r.Lock.Unlock()
	}
	return &result
}

func (r *Rmap) CloneDeep() *Rmap {
	var result Rmap
	result.Table = make(map[string]TokenPair, 8)

	for k, v := range r.Table {
		var entity TokenPair
		entity.Key = v.Key.CloneDeep()
		entity.Val = v.Val.CloneDeep()
		r.Lock.Lock()
		result.Table[k] = entity
		r.Lock.Unlock()
	}
	return &result
}

