package core

import "bytes"
import "fmt"

type TokenList struct {
	Room		uint
	EndIdx		uint
	Line		[]*Token
}

func NewTks(size int) *TokenList {
	var tks = &TokenList{}
	tks.Room = uint(size) + 1
	tks.EndIdx = 0
	tks.Line = make([]*Token, uint(size) + 1)
	return tks
}

func (tks *TokenList) Init() {
	tks.EndIdx = 0
	tks.Room = 9
	tks.Line = make([]*Token, uint(9))
}

func (tks *TokenList) List() []*Token{
	return tks.Line[0:tks.EndIdx]
}

func (tks *TokenList) Get(idx int) *Token{
	if idx < 0 || uint(idx) >= tks.EndIdx {
		return nil
	}
	return tks.Line[idx]
}

func (tks *TokenList) Size() int{
	return int(tks.Room)
}

func (tks *TokenList) Len() int{
	return int(tks.EndIdx)
}

func (tks *TokenList) Resize(size uint) {
	tks.Room = size
	temp := make([]*Token, size + 1)

	if tks.EndIdx > size {
		copy(temp, tks.Line[0 : size])
		tks.EndIdx = size
	}else{
		copy(temp, tks.Line[0 : tks.EndIdx])
	}
	tks.Line = temp
	
}

func (tks *TokenList) Add(t *Token) {
	if tks.Room == 0 || tks.EndIdx >= tks.Room - 1 {
		tks.Resize((tks.Room + 1) * 2)
	}
	tks.Line[tks.EndIdx] = t
	tks.EndIdx++
}

func (tks *TokenList) Pop() {
	if tks.EndIdx > 0 {
		tks.EndIdx--
	}
	if tks.EndIdx < tks.Room / 5 - 1 {
		tks.Resize(tks.Room / 5)
	}
}

func (tks *TokenList) Put(idx uint, t *Token) {
	for idx >= tks.Room - 1 {
		tks.Room = tks.Room * 2 + 1
	}
	if tks.Room > uint(len(tks.Line)) {
		tks.Resize(tks.Room)
	}

	tks.Line[idx] = t
	if idx >= tks.EndIdx {
		tks.EndIdx = idx + 1
	}
}

func (tks *TokenList) Insert(idx int, t *Token) {
	if tks.EndIdx >= tks.Room - 1 {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Token, tks.Room + 1)

	if tks.EndIdx >= uint(idx) {
		copy(temp, tks.Line[0:idx])
		temp[idx] = t
		copy(temp[idx+1:], tks.Line[idx : tks.EndIdx])
		tks.Line = temp
		tks.EndIdx++
	}else{
		tks.Put(uint(idx), t)
	}
}

func (tks *TokenList) InsertAll(idx int, list *TokenList) {
	if tks.EndIdx + list.EndIdx >= tks.Room - 1 {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Token, tks.Room + 1)

	if tks.EndIdx >= uint(idx) {
		copy(temp, tks.Line[0:idx])
		copy(temp[idx:], list.Line[0:list.EndIdx])
		copy(temp[uint(idx)+list.EndIdx:], tks.Line[idx : tks.EndIdx])
		tks.Line = temp
		tks.EndIdx += list.EndIdx
	}else{
		copy(temp, tks.Line)
		copy(temp[idx:], list.Line[0:list.EndIdx])
		tks.Line = temp
		tks.EndIdx = uint(idx) + list.EndIdx
	}
}

func (tks *TokenList) InsertArr(idx int, arr []*Token) {
	if tks.EndIdx + uint(len(arr)) >= tks.Room - 1 {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Token, tks.Room + 1)

	if tks.EndIdx > uint(idx) {
		copy(temp, tks.Line[0:idx])
		copy(temp[idx:], arr)
		copy(temp[idx+len(arr):], tks.Line[idx : tks.EndIdx])
		tks.Line = temp
		tks.EndIdx += uint(len(arr))
	}else{
		copy(temp, tks.Line)
		copy(temp[idx:], arr)
		tks.Line = temp
		tks.EndIdx = uint(idx) + uint(len(arr))
	}
}

func (tks *TokenList) First() *Token{
	if tks.EndIdx > 0 {
		return tks.Line[0]
	}
	return nil
}

func (tks *TokenList) Last() *Token{
	if tks.EndIdx > 0 {
		return tks.Line[tks.EndIdx - 1]
	}
	return nil
}

func (tks *TokenList) AddAll(list *TokenList) {
	for tks.Room < tks.EndIdx + list.EndIdx {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Token, tks.Room + 1)
	copy(temp, tks.Line[0:tks.EndIdx])
	copy(temp[tks.EndIdx:], list.Line[0:list.EndIdx])

	tks.EndIdx = tks.EndIdx + list.EndIdx
	tks.Line = temp
} 

func (tks *TokenList) AddArr(arr []*Token) {
	for tks.Room < tks.EndIdx + uint(len(arr)) {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Token, tks.Room + 1)
	copy(temp, tks.Line[0:tks.EndIdx])
	copy(temp[tks.EndIdx:], arr)

	tks.EndIdx = tks.EndIdx + uint(len(arr))
	tks.Line = temp
}

func (tks *TokenList) PopFirst() {
	if tks.EndIdx <= 0 {
		return
	}
	temp := make([]*Token, tks.Room)
	copy(temp, tks.Line[1:])
	tks.Line = temp
	tks.EndIdx--
	if tks.EndIdx < tks.Room / 5 - 1 {
		tks.Resize(tks.Room / 5)
	}
}

func (tks *TokenList) ToString() string{
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i:=0; i<int(tks.EndIdx); i++ {
		buffer.WriteString(tks.Line[i].ToString())
		buffer.WriteString(" ")
	}
	if len(buffer.Bytes()) > 1 {
		buffer.Bytes()[len(buffer.Bytes())-1] = ']'
	}else{
		buffer.WriteString("]")
	}
	return buffer.String()
}

func (tks *TokenList) Echo() {
	fmt.Println(tks.ToString())
}

func (tks *TokenList) Clone() *TokenList{
	var result TokenList
	result.Room = tks.Room
	result.EndIdx = tks.EndIdx
	result.Line = make([]*Token, result.Room + 1)
	copy(result.Line, tks.Line)
	return &result
}

func (tks *TokenList) CloneDeep() *TokenList{
	var result TokenList
	result.Room = tks.Room
	result.EndIdx = tks.EndIdx
	result.Line = make([]*Token, result.Room + 1)
	for i:=0; i< int(tks.EndIdx); i++ {
		result.Line[i] = tks.Line[i].CloneDeep()
	}
	return &result
}


func (tks *TokenList) Take(startIdx int, endIdx int) *TokenList{
	var result = NewTks(8)
	if startIdx < 0 {
		startIdx = 0
	}
	if endIdx > int(tks.EndIdx) {
		endIdx = int(tks.EndIdx)
	}
	result.AddArr(tks.Line[startIdx:endIdx])
	
	var part = endIdx - startIdx
	temp := make([]*Token, tks.Room + 1)
	copy(temp, tks.Line[0:startIdx])
	copy(temp[startIdx:], tks.Line[endIdx:])
	tks.Line = temp
	tks.EndIdx -= uint(part)
	if tks.EndIdx < tks.Room / 5 - 1 {
		tks.Resize(tks.Room / 5)
	}

	return result
}

func (tks *TokenList) Replace(old *Token, new *Token, at int, amount int){
	if amount < 0 {
		amount = int(^uint(0) >> 1) //取有符号int最大值
	}

	for i:=at; i<tks.Len() && amount>0; i++ {
		if tks.Line[i].Tp == old.Tp && tks.Line[i].Val == old.Val && i >= 0 {
			tks.Line[i].Copy(new)
			amount--
		}
	}
}

func (tks *TokenList) ReplacePart(old *TokenList, new *TokenList, at int, amount int){
	if amount < 0 {
		amount = int(^uint(0) >> 1) //取有符号int最大值
	}

	for i:=at; i<tks.Len() && amount>0; i++ {
		var eq = true
		for j:=0; j<old.Len(); j++ {
			if tks.Line[i+j].Tp != old.Line[j].Tp || tks.Line[i+j].Val != old.Line[j].Val {
				eq = false
			}
		}
		if eq {
			tks.Take(i, i + old.Len())
			tks.InsertAll(i, new)
			amount--
		}
	}
}

func (tks *TokenList) ReplacePartToOne(old *TokenList, new *Token, at int, amount int){
	if amount < 0 {
		amount = int(^uint(0) >> 1) //取有符号int最大值
	}

	for i:=at; i<tks.Len() && amount>0; i++ {
		var eq = true
		for j:=0; j<old.Len(); j++ {
			if tks.Line[i+j].Tp != old.Line[j].Tp || tks.Line[i+j].Val != old.Line[j].Val {
				eq = false
			}
		}
		if eq {
			tks.Take(i, i + old.Len())
			tks.Insert(i, new)
			amount--
		}
	}
}

func (tks *TokenList) ReplaceOneToPart(old *Token, new *TokenList, at int, amount int){
	if amount < 0 {
		amount = int(^uint(0) >> 1) //取有符号int最大值
	}

	for i:=at; i<tks.Len() && amount>0; i++ {
		if tks.Line[i].Tp == old.Tp && tks.Line[i].Val == old.Val {
			tks.Take(i, i + 1)
			tks.InsertAll(i, new)
			amount--
		}
	}
}

