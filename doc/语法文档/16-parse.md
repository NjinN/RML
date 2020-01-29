
---
### Parse方言

`parse` 方言时RML中用于进行文本处理的神器，充分掌握`parse`的用法，可以无限RML的功能。
</br>
RML中`parse`用于对字符串按给定规则进行匹配，并可在匹配过程中进行各种操作
</br></br>
首先来看一下`parse	`中`Rule`的定义

```
type Rule struct {
	minTimes int
	maxTimes int
	ruleStr  string
	ruleBlk  *Token
	code     *Token
	...	
	...

```
下面看一个小例子

```
parse "aaa" [1 3 "a"]
```
`parse`函数需要两个参数，第一个参数时要匹配的字符串，第二个参数时用于匹配的规则，格式上是方块类型。
</br></br>
`parse`函数启动时会创建空的`rule`，接着遍历规则方块。首先遇到一个整数`1`，此时，由于`rule`为空，所以给`rule`的`minTimes`设置值为`1`。接着遇到`3`，此时`minTimes`已有值，所以给`maxTimes`设置值为`3`。最后遇到字符串`"a"`，将`rule`的`ruleStr`设置值为`"a"`。此时`rule`完整，规则可表述为匹配子字符串`"a"`最少1次，最多3次。`parse`会根据规则，顺着参数一的字符串进行匹配。此例中字符串符合规则，返回`true`
</br></br>
`parse`函数支持使用特殊的表示范围的格式`range!`，例如

```
parse "aaa" [1..3 "a"]
```
此例与上例功能相同。
</br></br>
`parse`遇到字符串或方块时，会检查`rule`的`minTimes`和`maxTimes`，若`minTimes`和`maxTimes`都为空，则都设置为1，使`rule`完整。若`minTimes`不为空，而`maxTimes`为空，则设置`maxTimes`等于`minTimes`

```
parse "aaa" ["a"]
```
这个例子中`rule`的最小匹配次数和最大匹配次数都为1，返回`false`

```
parse "aaa" [2 "a"]
```
这个例子中`rule`的最小匹配次数和最大匹配次数都为2，返回`false`
</br></br>
`parse`中可以指定每次匹配时要进行的操作，例如

```
parse "ababab" [1..3 "ab" (print "match")]
```
在给定的匹配子字符串后再加一个圆块`( )`，则会在每次匹配成功时执行一次圆块中的代码。此例中匹配了3次子串，所以控制台打印了3次`"match"`
</br></br>
除了用字符串作为匹配规则，还可以使用方块指定子规则，例如

```
parse "abaab" [1..3 [0..1 "a" 0..1 "b"]]
```
这里定义了子规则，包含0到1个`"a"`和0到1个`"b"`，然后匹配这个规则1到3次，最终返回`true`。
</br></br>
`parse`中可以使用变量来实现动态的匹配规则，例如我们希望匹配这样一个规则的字符串

```
"aabb"
"aaabbb"		;a在前b在后，且出现的次数相同
```
则我们可以使用这种方式实现

```
n: 0
parse "aaabbb" [0..999 "a" (n += 1) n "b"]
```
给`n`一个初始值`0`，每次匹配到`"a"`都会使`n`自增1，然后用`n`限定`"b"`出现的次数。
</br>
当然这里还有更简便的方法，`parse`中支持匹配规则嵌套，但目前限制最大层数为500。所以我们可以这样写

```
rule: ["a" 0..1 rule "b"]

parse "aaabbb" rule 		;返回 true

parse "aabbb" rule			;返回 false

```
</br>
除了给定一种匹配规则外，`parse`中还可以同时给出匹配的多种可能

```
parse "aabbccabc" [0..999 ["a" | "b" | "c"]]
```
使用`|`可以表示这是多种可能并行的匹配规则，`parse`会依次进行匹配，只要匹配到一种规则就认为匹配成功，只有全都匹配不上是才判断为匹配失败。（注意，`|` 与规则间要有间隔）


</br></br>
#####关键字
作为RML的方言，`parse`中定义了一下关键字，用于实现一些特定操作，主要有

```
parse "aaabbb" [skip]			;skip是通配符，会跳过一个字符

parse "aaabbb" [to "b"]			;to会匹配字符串直到匹配"b"，并返回匹配"b"前的位置

parse "aaabbb" [thru "b"]		;thru会匹配字符串直到匹配"b"，并返回匹配"b"后的位置

parse "aaabbb" [some "a"]		;some等同于 1..int的最大值

parse "aaabbb" [any "a"]		;any等同于 0..int的最大值

parse "aaabbb" [opt "a"]		;opt等同于 0..1

parse "aaabbb" [end]			;end代表字符串的末尾

parse "aaabbb" [not "a"]		;对规则取反，此逻辑功能尚在测试中，感兴趣可自行探索

```

另外`parse`中一个非常重要的功能是在匹配过程中复制匹配到的子串，此时可以使用关键字`copy`

```
parse "123456" [to "2" copy str thru "5" (print str)]
```
执行此句，首先会向后查找直到`"2"`，此时游标指向`"2"`，`copy`会记住此时游标的位置，然后向后查找一个规则并进行匹配。当后一个规则匹配成功时，`copy`会拷贝前后游标间的子串。因此，这个例子中，控制台显示 `2345`，然后由于给的规则不能完全匹配字符串，最终返回一个 `false`
</br></br>
目前`parse`的主要规则就是这些，通过组合这些规则可以完成复杂逻辑的实现，下面展示本网站服务器使用的将动态页面代码转换成RML的`collect`代码的规则

```
res: copy ""

parse copy inp [
	opt [
		copy start-str to "<?" ( if start-str [append* res " keep " append* res format start-str append* res " " ] )
		|
		copy start-str to end ( if start-str [append* res " keep " append* res format start-str append* res " " ] )
	]
	
	some [
		thru "<?" copy code to "?>"
		(append* res code)
		[
			thru "?>" copy str to "<?"
			(append* res " keep " append* res format str)
			|
			thru "?>" copy str to end
			(append* res " keep " append* res format str)  
			|
			thru "?>" end         
		]
	]
]

res: append "collect [ " res
append* res "]"
```
















