package core

import "bytes"
// import "fmt"

type TokenPair struct {
	Key 	*Token
	Val		*Token
}

type Rmap struct {
	Table 	map[string]TokenPair
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
	}
	buffer.WriteString("}")
	return buffer.String()
}

func (r *Rmap) Get(key *Token) *Token {
	var keyString = TypeToStr(key.Tp) + key.ToString()
	pair, ok := r.Table[keyString]
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
	r.Table[keyString] = pair
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
		result.Table[k] = entity
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
		result.Table[k] = entity
	}
	return &result
}

