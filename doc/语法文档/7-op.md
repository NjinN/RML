
---
### 中缀运算符
RML中提供了以下中缀运算符

功能 | 符号 
- | :-: 
加法 		| +
减法		| - 
乘法 		| *
除以		| /
模		| %
加等		| +=
减等		| -=
乘等		| *=
除等		| /=
模等		| %=
交换		| ><
等于		| =
大于		| >
小于		| <
大于等于		| >=
小于等于		| <=
逻辑与		| and
逻辑或		| or		


中缀运算符在RML中运算优先级高于一般表达式，会被优先求值，但中缀运算符间平级，按照从左到右的顺序执行。

```
1 + 2 * 3		;返回 9
```
优先级可通过圆块 `( )` 控制

```
1 + (2 * 3) 	;返回 7
```

















