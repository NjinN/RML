package core

import "fmt"

type IntList struct {
	Room		uint
	EndIdx		uint
	Line		[]int
}

func NewIntList(size int) *IntList {
	var ils = IntList{}
	ils.Room = uint(size) + 1
	ils.EndIdx = 0
	ils.Line = make([]int, uint(size) + 1)
	return &ils
}

func (ils *IntList) Init() {
	ils.EndIdx = 0
	ils.Room = 9
	ils.Line = make([]int, uint(9))
}

func (ils *IntList) List() []int{
	return ils.Line[0:ils.EndIdx]
}

func (ils *IntList) Get(idx int) int{
	if idx < 0 {
		return -999999
	}

	return ils.Line[idx]
}

func (ils *IntList) Size() int{
	return int(ils.Room)
}

func (ils *IntList) Len() int{
	return int(ils.EndIdx)
}

func (ils *IntList) Resize(size uint) {
	ils.Room = size
	temp := make([]int, size + 1)

	if ils.EndIdx > size {
		copy(temp, ils.Line[0 : size])
		ils.EndIdx = size
	}else{
		copy(temp, ils.Line[0 : ils.EndIdx])
	}
	ils.Line = temp
	
}

func (ils *IntList) Add(i int) {
	if ils.Room == 0 || ils.EndIdx >= ils.Room - 1 {
		ils.Resize((ils.Room + 1) * 2)
	}
	ils.Line[ils.EndIdx] = i
	ils.EndIdx++
}

func (ils *IntList) Pop() {
	if ils.EndIdx > 0 {
		ils.EndIdx--
	}
	if ils.EndIdx < ils.Room / 5 - 1 {
		ils.Resize(ils.Room / 5)
	}
}

func (ils *IntList) Put(idx uint, i int) {
	for idx >= ils.Room - 1 {
		ils.Room = ils.Room * 2 + 1
	}
	if ils.Room > uint(len(ils.Line)) {
		ils.Resize(ils.Room)
	}

	ils.Line[idx] = i
	if idx >= ils.EndIdx {
		ils.EndIdx = idx + 1
	}
}

func (ils *IntList) Insert(idx int, i int) {
	if ils.EndIdx >= ils.Room - 1 {
		ils.Room = ils.Room * 2 + 1
	}
	temp := make([]int, ils.Room + 1)

	if ils.EndIdx >= uint(idx) {
		copy(temp, ils.Line[0:idx])
		temp[idx] = i
		copy(temp[idx+1:], ils.Line[idx : ils.EndIdx])
		ils.Line = temp
		ils.EndIdx++
	}else{
		ils.Put(uint(idx), i)
	}
}

func (ils *IntList) InsertAll(idx int, list *IntList) {
	if ils.EndIdx + list.EndIdx >= ils.Room - 1 {
		ils.Room = ils.Room * 2 + 1
	}
	temp := make([]int, ils.Room + 1)

	if ils.EndIdx >= uint(idx) {
		copy(temp, ils.Line[0:idx])
		copy(temp[idx:], list.Line[0:list.EndIdx])
		copy(temp[uint(idx)+list.EndIdx:], ils.Line[idx : ils.EndIdx])
		ils.Line = temp
		ils.EndIdx += list.EndIdx
	}else{
		copy(temp, ils.Line)
		copy(temp[idx:], ils.Line[0:ils.EndIdx])
		ils.Line = temp
		ils.EndIdx = uint(idx) + list.EndIdx
	}
}

func (ils *IntList) InsertArr(idx int, arr []int) {
	if ils.EndIdx + uint(len(arr)) >= ils.Room - 1 {
		ils.Room = ils.Room * 2 + 1
	}
	temp := make([]int, ils.Room + 1)

	if ils.EndIdx > uint(idx) {
		copy(temp, ils.Line[0:idx])
		copy(temp[idx:], arr)
		copy(temp[idx+len(arr):], ils.Line[idx : ils.EndIdx])
		ils.Line = temp
		ils.EndIdx += uint(len(arr))
	}else{
		copy(temp, ils.Line)
		copy(temp[idx:], arr)
		ils.Line = temp
		ils.EndIdx = uint(idx) + uint(len(arr))
	}
}

func (ils *IntList) First() int{
	return ils.Line[0]
}

func (ils *IntList) Last() int{
	return ils.Line[ils.EndIdx - 1]
}

func (ils *IntList) AddAll(list *IntList) {
	for ils.Room < ils.EndIdx + list.EndIdx {
		ils.Room = ils.Room * 2 + 1
	}
	temp := make([]int, ils.Room + 1)
	copy(temp, ils.Line[0:ils.EndIdx])
	copy(temp[ils.EndIdx:], list.Line[0:list.EndIdx])

	ils.EndIdx = ils.EndIdx + list.EndIdx
	ils.Line = temp
} 

func (ils *IntList) AddArr(arr []int) {
	for ils.Room < ils.EndIdx + uint(len(arr)) {
		ils.Room = ils.Room * 2 + 1
	}
	temp := make([]int, ils.Room + 1)
	copy(temp, ils.Line[0:ils.EndIdx])
	copy(temp[ils.EndIdx:], arr)

	ils.EndIdx = ils.EndIdx + uint(len(arr))
	ils.Line = temp
}

func (ils *IntList) PopFirst() {
	if ils.EndIdx <= 0 {
		return
	}
	temp := make([]int, ils.Room)
	copy(temp, ils.Line[1:])
	ils.Line = temp
	ils.EndIdx--
	if ils.EndIdx < ils.Room / 5 - 1 {
		ils.Resize(ils.Room / 5)
	}
}

func (ils *IntList) Echo() {
	fmt.Println(ils.List())
}

func (ils *IntList) Clone() IntList{
	var result IntList
	result.Room = ils.Room
	result.EndIdx = ils.EndIdx
	result.Line = make([]int, result.Room + 1)
	copy(result.Line, ils.Line)
	return result
}

