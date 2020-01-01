package core

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

func (tks *TokenList) Clear() {
	tks.EndIdx = 0
	tks.Room = 9
	tks.Line = make([]*Token, uint(9))
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
		copy(temp, tks.Line[0 : size])
	}
	tks.Line = temp
	
}

func (tks *TokenList) Add(t *Token) {
	if tks.Room == 0 || tks.EndIdx >= tks.Room - 1 {
		tks.Resize(tks.Room + 1 * 2)
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
	if idx > tks.EndIdx {
		tks.EndIdx = idx + 1
	}
}

func (tks *TokenList) Insert(idx uint, t *Token) {
	if tks.EndIdx >= tks.Room - 1 {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Token, tks.Room)

	if tks.EndIdx > idx {
		copy(temp, tks.Line[0:idx])
		temp[idx] = t
		copy(temp, tks.Line[idx+1 : tks.EndIdx])
		tks.Line = temp
		tks.EndIdx++
	}else{
		tks.Put(idx, t)
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
	temp := make([]*Token, tks.Room)
	copy(temp, tks.Line[0:tks.EndIdx])
	copy(temp[tks.EndIdx:], list.Line[0:list.EndIdx])

	tks.EndIdx = tks.EndIdx + list.EndIdx
	tks.Line = temp
} 

func (tks *TokenList) AddArr(arr []*Token) {
	for tks.Room < tks.EndIdx + uint(len(arr)) {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Token, tks.Room)
	copy(temp, tks.Line[0:tks.EndIdx])
	copy(temp[tks.EndIdx:], arr)

	tks.EndIdx = tks.EndIdx + uint(len(arr)
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
	for i:=0; i<tks.EndIdx; i++ {
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

func (tks *TokenList) Copy() {
	var result TokenList
	result.Room = tks.Room
	result.EndIdx = tks.EndIdx
	result.Line = make([]*Token, result.Room)
	copy(result.Line, tks.Line)
}

func (tks *TokenList) CopyDeep() {
	var result TokenList
	result.Room = tks.Room
	result.EndIdx = tks.EndIdx
	result.Line = make([]*Token, result.Room)
	for i:=0; i<tks.EndIdx; i++ {
		result.Line[i] = tks.Line[i].CopyDeep()
	}
}


