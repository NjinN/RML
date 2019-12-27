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

insert*: func [serial item /at at /only] [
	if not at [at: 1]
	_insert serial item at only
]

insert: func [serial item /at at /only] [
	if not at [at: 1]
	_insert copy/deep serial item at only
]

append*: func [serial item /only] [
	_append serial item only
]

append: func [serial item /only] [
	_append copy/deep serial item only
]

take*: func [serial /at idx /part len /last] [
	if not at [at: 1]
	if not part [part: 1] 

	if last [
		at: len? serial
		part: 1
	]

	_take serial starIdx partLen true
]

take: func [serial /at at /part part /last] [
	if not at [at: 1]
	if not part [part: 1] 

	if last [
		at: len? serial
		part: 1
	]

	_take copy/deep serial at part false
]

replace*: func [serial old new /at at /amount amount] [
	if not at [at: 1]
	if not amount [amount: 1]
	if all [amount: -1]
	_replace serial old new at amount
]

replace: func [serial old new /at at /amount amount /all] [
	if not at [at: 1]
	if not amount [amount: 1]
	if all [amount: -1]
	_replace copy/deep serial old new at amount
]




`