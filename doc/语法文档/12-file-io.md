
---
### 文件操作

RML中用于表示文件的Token以 `%` 开头，例如 `%c/rml/rml.exe`，若以 `/`结尾则表示是一个目录，否则是单个文件

```
now-dir
```
返回当前目录
</br></br>

```
cd %c/rml/
```
切换至指定目录
</br></br>

```
ls

ls/with %c/
```
列出当前目录中的所有文件，或使用 `/with` 修饰字指定目录
</br></br>

```
exist? %c/rml/rml.exe
```
判断文件是否存在，返回 `true` 或 `false`
</br></br>

```
abs-path %./test.txt
```
返回文件在系统中的绝对路径，仅当文件存在时可用，否则返回一个错误
</br></br>

```
rename %./test.txt %./try.txt
```
文件重命名，第一个参数是原文件名，第二个参数是新文件名，返回 `true` 或 `false`
</br></br>

```
remove %./test.txt
```
删除文件
</br></br>

```
make-dir %./rml/
```
创建目录
</br></br>

```
load %./test.txt
```
按字符串读取一个文件，并按照RML的语法格式将文本切割为一个方块。
</br></br>

```
read %./test.txt

read/bin %./test.txt
```
`read` 按字符串形式读取一个文件的，若需要按二进制读取，则使用 `/bin` 修饰字。
</br></br>

```
write %./test.txt "123"

write/append %./test.txt  #{5a4c}
```
`write` 向文件中写入数据，第一个参数时文件类型，第二个参数时字符串或二元类型。`write` 会直接覆盖原文件内容，若使用`/append` 修饰字则使用追加模式。


