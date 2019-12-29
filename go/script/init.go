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

print: func [n /inline /only] [
	_print n inline only
]

打印为 术 [甲 /行内 /单独] [
	_打印 甲 行内 单独
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


to-lit-word: func [a] [to lit-word! a]
转成原字为 术 [甲] [转成 原字类型 甲]

to-get-word: func [a] [to get-word! a]
转成取字为 术 [甲] [转成 取字类型 甲]

to-datatype: func [a] [to datatype! a]
转成类型为 术 [甲] [转成 类型类型 甲]

to-logic: func [a] [to logic! a]
转成逻辑为 术 [甲] [转成 逻辑类型 甲]

to-integer: func [a] [to integer! a]
转成整数为 术 [甲] [转成 整数类型 甲]

to-decimal: func [a] [to decimal! a]
转成小数为 术 [甲] [转成 小数类型 甲]

to-char: func [a] [to char! a]
转成字符为 术 [甲] [转成 字符类型 甲]

to-string: func [a] [to string! a]
转成字符串为 术 [甲] [转成 字符串类型 甲]

to-paren: func [a] [to paren! a]
转成圆块为 术 [甲] [转成 圆块类型 甲]

to-block: func [a] [to block! a]
转成方块为 术 [甲] [转成 方块类型 甲]

to-word: func [a] [to word! a]
转成单字为 术 [甲] [转成 单字类型 甲]

to-set-word: func [a] [to set-word! a]
转成设字为 术 [甲] [转成 设字类型 甲]

to-put-word: func [a] [to put-word! a]
转成置字为 术 [甲] [转成 置字类型 甲]

to-path: func [a] [to path! a]
转成路径为 术 [甲] [转成 路径类型 甲]


read: func [target /bin /string] [
	type:= string!
	if bin [type: bin!]
	if string [type: string!]
	_read target type
]

读取为 术 [目标 /二元 /字符串] [
	类型设为 字符串类型
	若 二元 [类型为 二元类型]
	若 字符串 [类型为 字符串类型]
	_读取 目标 类型
]

file?: func [path] [
	either (take/last to-string path) = #'/' [false] [true]
]

dir?: func [path] [
	either (take/last to-string path) = #'/' [true] [false]
]

文件?为 术 [路径] [
	是非 (取/尾 转成字符串 路径) 等于 #'/' [假] [真]
]

文件夹?为 术 [路径] [
	是非 (取/尾 转成字符串 路径) 等于 #'/' [真] [假]
]

ls: func [/dir dir] [
	_ls dir
]

列出目录为 术 [/目录 目录] [
	_列出目录 目录
]

write: func [path data /append] [
	_write path data append
]

写出为 术 [路径 数据 /添加] [
	_写出 路径 数据 添加
]



`