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

append*: func [collect item] [
	_append collect item
]

append: func [collect item] [
	_append copy/deep collect item
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