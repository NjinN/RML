
---
### 并行

RML封装了Go中的`goroutine`，从而很便捷的实现并行功能。
</br>
目前主要终于开启`goroutine`的函数有`fork`和`spawn`

```
fork [loop 100 [print "hello"]]
```
`fork`需要一个方块指定协程要执行的代码块。

```
res: ""
fork/result [1 + 2] res 
```
可以用 `/result` 指定协程执行代码块的返回值的接受者。当然，完全可以在代码块中进行绑定，而不使用此种方式。

```
fork/len [loop 100 [print "hello"]] 1024
```
`fork` 还提供一个`/len`修饰字指定协程的栈长，默认为1024。
</br>
</br>

```
spawn [ 
	[loop 100 [print 1]]
	[loop 100 [print 2]]
] 
```
使用`spawn`可以一次性开启多个协程。

```
spawn/wait [ 
	[loop 100 [print 1]]
	[loop 100 [print 2]]
] 
```
`spawn`提供一个修饰字`/wait`用于等待所有新起的协程执行完毕，程序才继续往下进行。

```
spawn/len/wait [ 
	[loop 100 [print 1]]
	[loop 100 [print 2]]
] 1024
```
同样`spawn` 提供一个`/len`修饰字指定协程的栈长，默认为1024。

