package script

const InitScript = 
`


do: func [code /with with] [
	_do code with
]

reduce: func [code /with with] [
	_reduce code with
]

copy: func [source /deep] [
	_copy source deep
]

insert*: func [collect item /at at /only] [
	idx:= 1
	if at [idx: at]
	_insert collect item idx only
]

insert: func [collect item /at at /only] [
	idx:= 1
	if at [idx: at]
	_insert copy/deep collect item idx only
]

append*: func [collect item /only] [
	_append collect item only
]

append: func [collect item /only] [
	_append copy/deep collect item only
]

take*: func [serial /at idx /part len /last] [
	starIdx:= 1
	partLen:= 1
	
	if idx [starIdx: idx]
	if len [partLen: len]

	if last [
		starIdx: len? serial
		partLen: 1
	]

	_take serial starIdx partLen true
]

take: func [serial /at idx /part len /last] [
	starIdx:= 1
	partLen:= 1
	
	if idx [starIdx: idx]
	if len [partLen: len]

	if last [
		starIdx: len? serial
		partLen: 1
	]

	_take serial starIdx partLen false
]



`