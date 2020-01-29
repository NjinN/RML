
---
### 条件分支

主流编程语言中，条件分支有指定的语法要求，而在RML中，则统一通过函数的形式提供</br>
目前主要实现了 `if` 和 `either` 函数</br>
其中 `if` 函数需要两个参数，第一个参数可以是任意类型的Token，其对应的逻辑值参照文末附表。第二个参数为方块类型，作为代码块供解释器执行。当第一个参数的值被认为是 `true` 时，执行第二个参数中的代码块，若为 `false` 则跳过。

```
if true [print "123"]		;显示 123
if 0 > 1 [print "123"]		;无输出
if 0 [print "123"]			;无输出, 0被视为逻辑假
```

`either` 函数需要三个参数，相较 `if` 多了一个逻辑假时执行的代码块，相当于主流语言的 `if ... else ...` 

```
either 0 > 1 [print 0] [print 1] 	;显示 1

```
由于RML中函数的参数都是定长的，所以无法实现连续的 `if ... else if ... else ... ` 语法。 </br>


##### 逻辑对照表 

类型 | flase值 
- | :-: 
none! 		| 必为false
logic!		| false 
integer! 	| 0
decimal!	| 0.0
char!		| 码值为0
string!		| ""
block!		| []
paren!		| ()
map!		| !map{}
object!		| 底层为nil

其他类型均视为 `false`





















