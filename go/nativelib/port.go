package nativelib

import . "../core"
import "strings"
import "net"

import "fmt"



func Oopen(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	
	if args[1].Tp == URL {
		var temp = strings.Split(args[1].Str(), "://")
		if len(temp) < 2 {
			return &Token{ERR, "Error addr"}, nil
		}

		var protocol = strings.ToLower(temp[0])
		var addr = temp[1]
		temp = strings.Split(addr, "/")

		

		if addr[0] == ':' || temp[0] == "127.0.0.1" || temp[0] == "localhost" {
			listener, err:= net.Listen(protocol, addr)
			if err != nil {
				return &Token{ERR, err.Error()}, nil
			}

			return newListenerPort(listener, protocol, addr, ctx, es), nil

		}else{
			if strings.IndexByte(temp[0], ':') < 0 {
				addr = strings.Replace(addr, temp[0], temp[0] + ":80", 1)
			}
			conn, err := net.Dial(protocol, addr)
			if err != nil {
				return &Token{ERR, err.Error()}, nil
			}

			return newConnPort(conn, protocol, addr, ctx), nil
		}

		
	}
	return &Token{ERR, "Type Mismatch"}, nil
}


func newListenerPort(listener net.Listener, protocol string, addr string, ctx *BindMap, es *EvalStack) *Token {
	var p = BindMap{make(map[string]*Token, 8), ctx, USR_CTX}

	p.PutNow("port", &Token{NONE, listener})
	p.PutNow("is-host", &Token{LOGIC, true})
	p.PutNow("protocol", &Token{STRING, protocol})
	p.PutNow("addr", &Token{STRING, addr})
	// p.PutNow("sub-ports", &Token{BLOCK, NewTks(8)})
	p.PutNow("awake", &Token{NONE, "none"})
	p.PutNow("conn", &Token{NONE, "none"})
	p.PutNow("listening", &Token{LOGIC, false})

	return &Token{PORT, &p}
}


func newConnPort(conn net.Conn, protocol string, addr string, ctx *BindMap) *Token {
	var p = BindMap{make(map[string]*Token, 8), ctx, USR_CTX}

	p.PutNow("port", &Token{NONE, conn})
	p.PutNow("is-host", &Token{LOGIC, false})
	p.PutNow("protocol", &Token{STRING, protocol})
	p.PutNow("host-addr", &Token{STRING, addr})
	p.PutNow("local-addr", &Token{STRING, conn.LocalAddr().String()})
	p.PutNow("remote-addr", &Token{STRING, conn.RemoteAddr().String()})
	p.PutNow("read-timeout", &Token{INTEGER, 0})
	p.PutNow("write-timeout", &Token{INTEGER, 0})
	p.PutNow("in-buffer", &Token{NONE, "none"})
	p.PutNow("out-buffer", &Token{NONE, "none"})
	p.PutNow("in-buffer-size", &Token{INTEGER, 4096})
	p.PutNow("out-buffersize", &Token{INTEGER, 4096})
	p.PutNow("awake", &Token{NONE, "none"})
	p.PutNow("listening", &Token{LOGIC, false})

	return &Token{PORT, &p}
}


func listenListener(listener net.Listener, p *BindMap, es *EvalStack){
	p.Table["listening"].Val = true
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		var subConn = newConnPort(conn, p.Table["protocol"].Str(), p.Table["addr"].Str(), p.Father)

		// p.Table["sub-ports"].List().Add(subConn)
		p.Table["conn"] = subConn

		var callback = p.Table["awake"]
		if callback.Tp == BLOCK && callback.List().Len() > 0 {
			temp, err := es.Eval(callback.Tks(), p)
			if err != nil {
				fmt.Println(err.Error())
			}
			if temp != nil && temp.Tp == ERR {
				fmt.Println(temp.Str())
			}
		}

		if p.Get("listening").Tp != LOGIC || p.Get("listening").Val.(bool) != true {
			break
		}
	}
}

func listenConn(conn net.Conn, p *BindMap, es *EvalStack){
	var bufferSizeToken = p.Get("in-buffer-size")
	var bufferSize int
	if bufferSizeToken.Tp == INTEGER && bufferSizeToken.Int() > 0 {
		bufferSize = bufferSizeToken.Int()
	}else{
		bufferSize = 4096
	}
	p.Table["listening"].Val = true
	for {
		p.Table["in-buffer"] = &Token{NONE, ""}
		var buffer = make([]byte, bufferSize)

		n, err := conn.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" || strings.Contains(err.Error(), "use of closed network connection") {
				conn.Close()
				// fmt.Println("Conn is closed")
				break
			}
			fmt.Println(err.Error())
		}

		if n > 0 {
			p.Table["in-buffer"] = &Token{BIN, buffer[0:n]}
		}

		var callback = p.Table["awake"]
		if callback.Tp == BLOCK && callback.List().Len() > 0 {
			temp, err := es.Eval(callback.Tks(), p)
			if err != nil {
				fmt.Println(err.Error())
			}
			if temp != nil && temp.Tp == ERR {
				if strings.Contains(temp.Str(), "use of closed network connection") {
					conn.Close()
					// fmt.Println("Conn is closed")
					break
				}
				fmt.Println(temp.Str())
			}
		}

		if p.Table["listening"].Tp != LOGIC || p.Table["listening"].Val.(bool) != true {
			break
		}
	}
}


func waitListener(listener net.Listener, p *BindMap, es *EvalStack) *Token{

	
	conn, err := listener.Accept()
	if err != nil {
		return &Token{ERR, err.Error()}
	}

	var subConn = newConnPort(conn, p.Table["protocol"].Str(), p.Table["addr"].Str(), p.Father)

	// p.Table["sub-ports"].List().Add(subConn)
	p.Table["conn"] = subConn

	return subConn
}

func waitConn(conn net.Conn, p *BindMap, es *EvalStack) *Token{
	var bufferSizeToken = p.Get("in-buffer-size")
	var bufferSize int
	if bufferSizeToken.Tp == INTEGER && bufferSizeToken.Int() > 0 {
		bufferSize = bufferSizeToken.Int()
	}else{
		bufferSize = 4096
	}

	var none = &Token{NONE, ""}

	p.Table["in-buffer"] = none
	var buffer = make([]byte, bufferSize)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err.Error())
	}

	if n > 0 {
		p.Table["in-buffer"] = &Token{BIN, buffer[0:n]}
		return &Token{BIN, buffer[0:n]}
	}else{
		return none
	}
	
}

func ReadPort(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	
	if args[1].Tp == PORT && args[2].Tp == DATATYPE {
		var isHost = args[1].Ctx().Get("is-host")
		if isHost == nil || isHost.Tp != LOGIC || isHost.Val.(bool) {
			return &Token{ERR, "Target is not a conn"}, nil
		} 

		if args[2].Int() != STRING && args[2].Int() != BIN {
			return &Token{ERR, "Error output type"}, nil
		}

		var inBuffer = args[1].Ctx().Get("in-buffer")
		if inBuffer.Tp == NONE {
			return inBuffer, nil
		}

		if args[2].Int() == BIN {
			return inBuffer, nil
		}else if args[2].Int() == STRING {
			return &Token{STRING, string(inBuffer.Val.([]byte))}, nil
		}

	}

	return &Token{ERR, "Type Mismatch"}, nil
}

func WritePort(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if args[1].Tp == PORT && (args[2].Tp == STRING || args[2].Tp == BIN) {
		var isHost = args[1].Ctx().Get("is-host")
		if isHost == nil || isHost.Tp != LOGIC || isHost.Val.(bool) {
			return &Token{ERR, "Target is not a conn"}, nil
		} 
		var outBuffer []byte
		if args[2].Tp == BIN {
			outBuffer = args[2].Val.([]byte)
		}
		if args[2].Tp == STRING {
			outBuffer = []byte(args[2].Str())
		}

		n, err := args[1].Ctx().Get("port").Val.(net.Conn).Write(outBuffer)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		return &Token{INTEGER, n}, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}



func Wait(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if args[1].Tp == PORT {
		var isHost = args[1].Ctx().Get("is-host")
		if isHost.Tp == LOGIC {
			if isHost.Val.(bool) {
				return waitListener(args[1].Ctx().Get("port").Val.(net.Listener), args[1].Ctx(), es), nil
			}else{
				return waitConn(args[1].Ctx().Get("port").Val.(net.Conn), args[1].Ctx(), es), nil
			}

		}
	}
	return &Token{ERR, "Type Mismatch"}, nil
}


func Listen(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if args[1].Tp == PORT {
		var isHost = args[1].Ctx().Get("is-host")
		if isHost.Tp == LOGIC {
			if isHost.Val.(bool) {
				listenListener(args[1].Ctx().Get("port").Val.(net.Listener), args[1].Ctx(), es)
				return nil, nil
			}else{
				listenConn(args[1].Ctx().Get("port").Val.(net.Conn), args[1].Ctx(), es)
				return nil, nil
			}
		}
	}
	return &Token{ERR, "Type Mismatch"}, nil
}

func Close(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if args[1].Tp == PORT {
		var isHost = args[1].Ctx().Get("is-host")
		if isHost.Tp == LOGIC {
			if isHost.Val.(bool) {
				err := args[1].Ctx().Get("port").Val.(net.Listener).Close()
				if err != nil {
					return  &Token{ERR, err.Error()}, nil
				}
				return nil, nil
			}else{
				err := args[1].Ctx().Get("port").Val.(net.Conn).Close()
				if err != nil {
					return  &Token{ERR, err.Error()}, nil
				}
				return nil, nil
			}
		}
	}
	return &Token{ERR, "Type Mismatch"}, nil
}
