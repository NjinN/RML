package script

const InitScript = 
`


do: func [code /with with] [
	_do code with
]

执行为 术 [
	"执行一个代码块或符合代码格式的字符串" 
	代码 	"方块或字符串"
	/于 于	"带参，指定执行代码的语境，类型为对象类型"
	] [
	_执行 代码 于
]

reduce: func [code /with with] [
	_reduce code with
]

收敛为 术 [
	"执行一个代码块或符合代码格式的字符串，将每个表达式的返回值置于一个方块中，最终返回该方块"
	代码	"方块或字符串"
	/于 于	"带参，指定执行代码的语境，类型为对象类型"
	] [
	_收敛 代码 于
]

copy: func [source /deep] [
	_copy source deep
]

复制为  术 [
	"复制一个Token"
	源 		"被复制的Token，任意类型"
	/深		"无参，深拷贝一个Token"
	] [
	_复制 源 深
]

print: func [n /inline /only] [
	_print n inline only
]

打印为 术 [
	"在控制台打印一个Token"
	甲 		"被打印的Token，任意类型"
	/行内 	"无参，打印且不换行"
	/单独	"无参，在接受一个方块时，视为单个Token"
	] [
	_打印 甲 行内 单独
]

insert*: func [serial item /at at /only] [
	if not at [at: 1]
	_insert serial item at only
]

插入*为 术 [
	"向集合类型中插入值，可用于方块、字符串，会改变传入的集合"
	集合 	"被插入的集合，方块或字符串"
	单体 	"要插入的值"
	/于 于 	"带参，指定插入的位置，整数类型"
	/单独	"无参，在接受一个方块时，视为单个集合"
	] [
	若 非 于 [于为 1]
	_插入 集合 单体 于 单独
]

insert: func [serial item /at at /only] [
	if not at [at: 1]
	_insert copy/deep serial item at only
]

插入为 术 [
	"向集合类型中插入值，可用于方块、字符串，不改变传入的集合"
	集合 	"被插入的集合，方块或字符串"
	单体 	"要插入的值"
	/于 于	"带参，指定插入的位置，整数类型" 
	/单独	"无参，在接受一个方块时，视为单个集合"
	] [
	若 非 于 [于为 1]
	_插入 复制/深 集合 单体 于 单独
]

append*: func [serial item /only] [
	_append serial item only
]

添加*为 术 [
	"在集合类型末尾添加值，可用于方块、字符串，会改变传入的集合"
	集合 	"要添加值的集合，方块或字符串"
	单体 	"要添加的值"
	/单独	"无参，在接受一个方块时，视为单个集合"
	] [
	_添加 集合 单体 /单独
]

append: func [serial item /only] [
	_append copy/deep serial item only
]

添加为 术 [
	"在集合类型末尾添加值，可用于方块、字符串，不会改变传入的集合"
	集合 	"要添加值的集合，方块或字符串"
	单体 	"要添加的值"
	/单独	"无参，在接受一个方块时，视为单个集合"
	] [
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

取*为 术 [
	"从集合中取出值，可用于方块、字符串，会改变传入的集合"
	集合 		"要取值的目标集合"
	/于 于 		"带参，指定取出值的位置，整数类型"
	/部分 部分	"带参，指定取出的值的长度，整数类型"
	/尾			"无参，取出集合最后一个值，高优先级"
	] [
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

取为 术 [
	"从集合中取出值，可用于方块、字符串，不会改变传入的集合"
	集合 		"要取值的目标集合"
	/于 于 		"带参，指定取出值的位置，整数类型"
	/部分 部分	"带参，指定取出的值的长度，整数类型" 
	/尾			"无参，取出集合最后一个值，高优先级"
	] [
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

替换*为 术 [
	"替换集合中的值，可用于方块、字符串，会改变传入的集合"
	集合 		"要替换值的目标集合"
	旧 			"要替换的旧值"
	新 			"要替换的新值"
	/于 于 		"带参，指定替换的起始位置，整数类型"
	/数量 数量 	"带参，指定替换的数量，整数类型，小于0时代表替换全部"
	/全部		"无参，替换所有匹配到的值，高优先级"
	] [
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

替换为 术 [
	"替换集合中的值，可用于方块、字符串，不会改变传入的集合"
	集合 		"要替换值的目标集合"
	旧 			"要替换的旧值"
	新 			"要替换的新值"
	/于 于 		"带参，指定替换的起始位置，整数类型"
	/数量 数量	"带参，指定替换的数量，整数类型，小于0时代表替换全部" 
	/全部		"无参，替换所有匹配到的值，高优先级"
	] [
	若 非 于 [于为 1]
	若 非 数量 [数量为 1]
	若 全部 [数量为 -1]
	_替换 复制/深 集合 旧 新 于 数量
]


to-lit-word: func [a] [to lit-word! a]
转成原字为 术 ["将Token转为原字" 甲] [转成 原字类型 甲]

to-get-word: func [a] [to get-word! a]
转成取字为 术 ["将Token转为取字" 甲] [转成 取字类型 甲]

to-datatype: func [a] [to datatype! a]
转成类型为 术 ["将Token转为类型" 甲] [转成 类型类型 甲]

to-logic: func [a] [to logic! a]
转成逻辑为 术 ["将Token转为逻辑值" 甲] [转成 逻辑类型 甲]

to-integer: func [a] [to integer! a]
转成整数为 术 ["将Token转为整数" 甲] [转成 整数类型 甲]

to-decimal: func [a] [to decimal! a]
转成小数为 术 ["将Token小数" 甲] [转成 小数类型 甲]

to-char: func [a] [to char! a]
转成字符为 术 ["将Token字符" 甲] [转成 字符类型 甲]

to-string: func [a] [to string! a]
转成字符串为 术 ["将Token字符串" 甲] [转成 字符串类型 甲]

to-paren: func [a] [to paren! a]
转成圆块为 术 ["将Token圆块" 甲] [转成 圆块类型 甲]

to-block: func [a] [to block! a]
转成方块为 术 ["将Token方块" 甲] [转成 方块类型 甲]

to-word: func [a] [to word! a]
转成单字为 术 ["将Token单字" 甲] [转成 单字类型 甲]

to-set-word: func [a] [to set-word! a]
转成设字为 术 ["将Token设字" 甲] [转成 设字类型 甲]

to-put-word: func [a] [to put-word! a]
转成置字为 术 ["将Token置字" 甲] [转成 置字类型 甲]

to-path: func [a] [to path! a]
转成路径为 术 ["将Token路径" 甲] [转成 路径类型 甲]


read: func [target /bin /string] [
	type:= string!
	if bin [type: bin!]
	if string [type: string!]
	_read target type
]

读取为 术 [
	"读取文件"
	目标 		"要读取的目标文件"
	/二元 		"无参，指定读取结果为二元类型"		
	/字符串		"无参，指定读取结果为字符串类型"	
	] [
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

文件?为 术 [
	"判断文件路径是否为单个文件"
	路径
	] [
	是非 (取/尾 转成字符串 路径) 等于 #'/' [假] [真]
]

文件夹?为 术 [
	"判断文件路径是否为文件夹"
	路径
	] [
	是非 (取/尾 转成字符串 路径) 等于 #'/' [真] [假]
]

ls: func [/dir dir] [
	_ls dir
]

列出目录为 术 [
	"列出文件目录下的所有文件和文件夹"
	/目录 目录	"带参，指定目录"
	] [
	_列出目录 目录
]

write: func [path data /append] [
	_write path data append
]

写出为 术 [
	"将数据写出到文件中"
	路径 	"要写出数据的文件"
	数据 	"要写出的数据，接受字符串和二元类型"
	/添加	"无参，在文件的结尾添加数据而不是覆盖"
	] [
	_写出 路径 数据 添加
]

cmd: func [c /no-wait /output output] [
	_cmd c no-wait output
]

命令为 术 [
	"执行cmd命令"
	令 			"要执行的cmd命令， 字符串"
	/不等待 	"无参，不等待命令执行结束"
	/输出 输出	"带参，将命令执行结果输出到指定Token"
	] [
	_命令 令 不等待 输出
]

fork: func [code /result result] [
	_fork code result
]

分支为 术 [
	"启动一个线程执行代码"
	代码 		"要执行的代码，方块或字符串"
	/结果 结果	"带参，指定保存代码执行结果的Token"
	] [
	_分支 代码 结果
]

spawn: func [codes /wait] [
	_spawn codes wait
]

并行为 术 [
	"同时启动多个线程执行代码，可以选择是否等待线程执行完毕"
	代码集 	"要执行的代码，方块类型，内部包含的每个方块或字符串会启动一个线程执行"
	/等待	"无参，等待所有线程执行完毕"
	] [
	_并行 代码集 等待
]




`