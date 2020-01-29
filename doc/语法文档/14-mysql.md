
---
### Mysql连接

由于本网站的需求，RML中直接实现了Mysql连接与操作，其本质是一个使用`mysql`协议的`port!`

##### 连接mysql数据库
```
db: open mysql://root:root@(127.0.0.1:3306)/test
```
执行sql语句使用`write`函数

```
res: write db "SELECT * FROM table"
```
返回结果类似于

```
[[1 "张三"] [2 "李四"]]
```
可以使用方块的操作方式对返回结果进行操作。</br>
如果希望返回结果带上列名，可以使用`/name`修饰字

```
res: write/name db "SELECT id, name FROM table"
```
则返回结果为

```
[[id: 1 name: "张三"] [id: 2 name: "李四"]]

res/1/name		;显示 张三
```
</br>
支持使用sql预编译

```
res: write db [
	"SELECT * FROM table WHERE id=? AND name=?"
	1
	"张三"
] 
```
同样，mysql连接使用完毕后，注意及时关闭连接

```
close db
```
