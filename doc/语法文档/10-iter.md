
---
### 遍历集合
RML中使用 `foreach` 函数遍历集合

```
foreach item [1 2 3] [print item] 	;显示 1 ~ 3
``` 
迭代变量可以为方块

```
foreach [a b] [1 2 3] [
	print a 		
	print b
]				;依次显示 1 2 3 none
```
若最后一次迭代时集合元素不足，会补上 `none` 值
</br>
</br>
`foreach` 可以用在对象上

```
foreach item {a: 123} [print item]		;显示 [a: 123]

foreach [k v] {a: 123} [
	print k		;显示 a
	print v		;显示 123
]
```
</br>
同样`foreach` 可以用在哈希表上

```
foreach item !map{[1 "abc"]} [print item]		;显示 [1 "abc"]

foreach [k v] !map{[1 "abc"]} [
	print k		;显示 1
	print v		;显示 abc
]
```

















