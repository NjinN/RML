
---
### 网络连接

RML中使用 `port!` 各种网络连接。</br>
本质上`port!`是对Go语言的`Listener`和`Conn`接口的封装，其在RML中底层与对象一致。</br>
下例打开一个`tcp`端口

```
open tcp://:8384
```
由于没有指定ip，所以被视为本地端口，底层创建一个`Listener`</br>
若指定了ip，则在底层创建一个`Conn`

```
open tcp://11.12.12.14:8384
```
`open` 会返回一个 `port!` 类型的Token，可以使用 `print` 查看，其中`Listener` 显示如下

```
{addr: ":8384" awake: none on-close: none conn: none 
listening: false port: none is-host: true protocol: "tcp"}
```
注意其中的 `port` 对应一个`none`值，请不要修改它。实际上其底层就是一个 `Listener`(当然，这种做法并不优雅，暂时的实现方式)。`is-host`为`true`表示这是一个 `Listener`。`awake`用来定义接收到新连接时的操作。`on-close`用来定义关闭时的操作。

`Conn`类型的`port!`显示如下

```
{local-addr: "192.168.1.3:59679" read-timeout: 0 write-timeout: 0 
out-buffersize: 4096 is-host: false in-buffer: none on-close: none 
protocol: "tcp" listening: false port: none host-addr: "www.baidu.com:80" 
remote-addr: "14.215.177.39:80" out-buffer: none in-buffer-size: 4096 
awake: none awake-ts: 0}
```
`is-host`为`false`表示这是一个 `Conn`。`awake`用来定义接收到数据时的操作。`on-close`用来定义关闭时的操作。
</br></br>
使用`wait`函数，等待`port!`接收新连接或新数据。等待结束，新连接存放在`conn`属性中，新数据存放在`in-buffer`中。
</br></br>
对于`port!`同样使用`read`和`write`实现接收和发送数据。

```
p: open tcp://11.12.13.14:80

read p

write p "hello"

```

为了简化操作，对于 `port!` 提供了`listen`函数来进行监听。`listen`会循环调用`read`，当收到新连接或新数据时，执行`awake`绑定的方块。示例如下

```
server: open tcp://:8384

server/awake: [

    conn/awake: copy [
        write this-port "HTTP/1.1 200 OK^M^/Content-Type:text/html; charset=utf8^M^/Content-Length:14^M^/^M^/<h1>HELLO</h1>"      
    ]
    conn/read-timeout: 30
    conn/on-close: [print "close"]
    fork [listen conn]
]

print "start listen"
fork [listen server]

```

本网站目前就采用这种方式实现，当然，具体的实现里还需要对请求信息进行处理。
</br></br>
关闭`port!`使用`close`函数

```
close server
```
若`on-close`绑定了方块，则在关闭时，会执行该方块







