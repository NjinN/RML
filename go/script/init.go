package script

const InitScript = 
`


do: func [code /with with] [
	_do code with
]

执行为 术 [代码 /于 于] [
	_执行 代码 于
]

reduce: func [code /with with] [
	_reduce code with
]

收敛为 术 [代码 /于 于] [
	_收敛 代码 于
]

copy: func [source /deep] [
	_copy source deep
]

复制为  术 [源 /深] [
	_复制 源 深
] 

insert*: func [serial item /at at /only] [
	if not at [at: 1]
	_insert serial item at only
]

插入*为 术 [集合 单体 /于 于 /单独] [
	若 非 于 [于为 1]
	_插入 集合 单体 于 单独
]

insert: func [serial item /at at /only] [
	if not at [at: 1]
	_insert copy/deep serial item at only
]

插入为 术 [集合 单体 /于 于 /单独] [
	若 非 于 [于为 1]
	_插入 复制/深 集合 单体 于 单独
]

append*: func [serial item /only] [
	_append serial item only
]

添加*为 术 [集合 单体 /单独] [
	_添加 集合 单体 /单独
]

append: func [serial item /only] [
	_append copy/deep serial item only
]

添加为 术 [集合 单体 /单独] [
	_添加 复制/深 集合 单体 单独
]

take*: func [serial /at at /part part /last] [
	if not at [at: 1]
	if not part [part: 1] 

	if last [
		at: len? serial
		part: 1
	]

	_take serial at part true
]

取*为 术 [集合 /于 于 /部分 部分 /尾] [
	若 非 于 [于为 1]
	若 非 部分 [部分为 1]

	若 尾 [
		于为 	长? 集合
		部分为 	1
	]

	_取 集合 于 部分 真
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

取为 术 [集合 /于 于 /部分 部分 /尾] [
	若 非 于 [于为 1]
	若 非 部分 [部分为 1]

	若 尾 [
		于为 	长? 集合
		部分为 	1
	]

	_取 复制/深 集合 于 部分 假
]

replace*: func [serial old new /at at /amount amount /all] [
	if not at [at: 1]
	if not amount [amount: 1]
	if all [amount: -1]
	_replace serial old new at amount
]

替换*为 术 [集合 旧 新 /于 于 /数量 数量 /全部] [
	若 非 于 [于为 1]
	若 非 数量 [数量为 1]
	若 全部 [数量为 -1]
	_替换 集合 旧 新 于 数量
]

replace: func [serial old new /at at /amount amount /all] [
	if not at [at: 1]
	if not amount [amount: 1]
	if all [amount: -1]
	_replace copy/deep serial old new at amount
]

替换为 术 [集合 旧 新 /于 于 /数量 数量 /全部] [
	若 非 于 [于为 1]
	若 非 数量 [数量为 1]
	若 全部 [数量为 -1]
	_替换 复制/深 集合 旧 新 于 数量
]




`