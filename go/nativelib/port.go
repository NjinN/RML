package nativelib

import (
	"database/sql"
	"net"
	"strings"
	"sync"
	"time"

	. "github.com/NjinN/RML/go/core"

	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

func Oopen(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == URL {
		var temp = strings.Split(args[1].Str(), "://")
		if len(temp) < 2 {
			return &Token{ERR, "Error addr"}, nil
		}

		var protocol = strings.ToLower(temp[0])
		var addr = temp[1]
		temp = strings.Split(addr, "/")

		if protocol == "mysql" {
			db, err := sql.Open(protocol, addr)
			if err != nil {
				return &Token{ERR, err.Error()}, nil
			}
			return newMysqlPort(db, ctx), nil
		}

		if addr[0] == ':' || temp[0] == "127.0.0.1" || temp[0] == "localhost" {
			listener, err := net.Listen(protocol, addr)
			if err != nil {
				return &Token{ERR, err.Error()}, nil
			}

			return newListenerPort(listener, protocol, addr, ctx, es), nil

		} else {
			if strings.IndexByte(temp[0], ':') < 0 {
				addr = strings.Replace(addr, temp[0], temp[0]+":80", 1)
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
	var p = BindMap{make(map[string]*Token, 8), ctx, USR_CTX, sync.RWMutex{}}

	p.PutNow("port", &Token{NONE, listener})
	p.PutNow("is-host", &Token{LOGIC, true})
	p.PutNow("protocol", &Token{STRING, protocol})
	p.PutNow("addr", &Token{STRING, addr})
	// p.PutNow("sub-ports", &Token{BLOCK, NewTks(8)})
	p.PutNow("awake", &Token{NONE, "none"})
	p.PutNow("on-close", &Token{NONE, "none"})
	p.PutNow("conn", &Token{NONE, "none"})
	p.PutNow("listening", &Token{LOGIC, false})

	return &Token{PORT, &p}
}

func newConnPort(conn net.Conn, protocol string, addr string, ctx *BindMap) *Token {
	var p = BindMap{make(map[string]*Token, 8), ctx, USR_CTX, sync.RWMutex{}}

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
	p.PutNow("on-close", &Token{NONE, "none"})
	p.PutNow("listening", &Token{LOGIC, false})
	p.PutNow("awake-ts", &Token{INTEGER, 0})

	return &Token{PORT, &p}
}

func listenListener(listener net.Listener, p *BindMap, es *EvalStack) {
	p.GetNow("listening").Val = true
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		var subConn = newConnPort(conn, p.GetNow("protocol").Str(), p.GetNow("addr").Str(), p.Father)

		// p.GetNow("sub-ports").List().Add(subConn)
		p.PutNow("conn", subConn)

		var callback = p.GetNow("awake")
		if callback.Tp == BLOCK && callback.List().Len() > 0 {
			temp, err := es.Eval(callback.Tks(), p)
			if err != nil {
				fmt.Println(err.Error())
			}
			if temp != nil && temp.Tp == ERR {
				fmt.Println(temp.Str())
			}
		}

		if p.GetNow("listening").Tp != LOGIC || p.GetNow("listening").Val.(bool) != true {
			break
		}
	}
}

func listenConn(conn net.Conn, p *BindMap, es *EvalStack) {
	var bufferSizeToken = p.GetNow("in-buffer-size")
	var bufferSize int
	if bufferSizeToken.Tp == INTEGER && bufferSizeToken.Int() > 0 {
		bufferSize = bufferSizeToken.Int()
	} else {
		bufferSize = 4096
	}
	p.GetNow("in-buffer").Tp = BIN
	p.GetNow("in-buffer").Val = make([]byte, 0)
	p.GetNow("listening").Val = true
	var buffer = make([]byte, bufferSize)



	for {

		if p.GetNow("read-timeout").Int() > 0{
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(p.GetNow("read-timeout").Int())))
		}else{
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(200)))
		}

		n, err := conn.Read(buffer)
		// fmt.Println("conn awake")
		if err != nil {
			if err.Error() == "EOF" {
				if p.GetNow("read-timeout").Int() > 0 {
					if int(time.Now().Unix())-p.GetNow("awake-ts").Int() > p.GetNow("read-timeout").Int() {
						conn.Close()
						// fmt.Println("Conn is closed")
						var closeCode = p.GetNow("on-close")
						if closeCode != nil && closeCode.Tp == BLOCK && closeCode.List().Len() > 0 {
							temp, err := es.Eval(closeCode.Tks(), p)
							if err != nil {
								fmt.Println(err.Error())
							}
							if temp != nil && temp.Tp == ERR {
								fmt.Println(temp.Str())
							}
						}

						break
					}
				}

				// time.Sleep(time.Duration(200) * time.Millisecond)
				continue
			} else if strings.Contains(err.Error(), "use of closed network connection") || strings.Contains(err.Error(), "wsarecv") {
				conn.Close()
				// fmt.Println("Conn is closed")
				var closeCode = p.GetNow("on-close")
				if closeCode != nil && closeCode.Tp == BLOCK && closeCode.List().Len() > 0 {
					temp, err := es.Eval(closeCode.Tks(), p)
					if err != nil {
						fmt.Println(err.Error())
					}
					if temp != nil && temp.Tp == ERR {
						fmt.Println(temp.Str())
					}
				}

				break
			}
			// fmt.Println(err.Error())
		}

		if n > 0 {
			p.GetNow("in-buffer").Val = buffer[0:n]
			p.GetNow("awake-ts").Val = int(time.Now().Unix())
		} else {
			continue
		}

		var callback = p.GetNow("awake")
		if callback.Tp == BLOCK && callback.List().Len() > 0 {
			temp, err := es.Eval(callback.Tks(), p)
			if err != nil {
				fmt.Println(err.Error())
			}
			if temp != nil && temp.Tp == ERR {
				if strings.Contains(temp.Str(), "use of closed network connection") || strings.Contains(temp.Str(), "wsarecv") {
					conn.Close()
					// fmt.Println("Conn is closed")
					var closeCode = p.GetNow("on-close")
					if closeCode != nil && closeCode.Tp == BLOCK && closeCode.List().Len() > 0 {
						temp, err := es.Eval(closeCode.Tks(), p)
						if err != nil {
							fmt.Println(err.Error())
						}
						if temp != nil && temp.Tp == ERR {
							fmt.Println(temp.Str())
						}
					}
					break
				}
				fmt.Println(temp.Str())
			}
		}

		if p.GetNow("listening").Tp != LOGIC || p.GetNow("listening").Val.(bool) != true {
			break
		}
	}
}

func waitListener(listener net.Listener, p *BindMap, es *EvalStack) *Token {

	conn, err := listener.Accept()
	if err != nil {
		return &Token{ERR, err.Error()}
	}

	var subConn = newConnPort(conn, p.GetNow("protocol").Str(), p.GetNow("addr").Str(), p.Father)

	// p.GetNow("sub-ports").List().Add(subConn)
	p.PutNow("conn", subConn)

	return subConn
}

func waitConn(conn net.Conn, p *BindMap, es *EvalStack) *Token {
	var bufferSizeToken = p.GetNow("in-buffer-size")
	var bufferSize int
	if bufferSizeToken.Tp == INTEGER && bufferSizeToken.Int() > 0 {
		bufferSize = bufferSizeToken.Int()
	} else {
		bufferSize = 4096
	}

	var none = &Token{NONE, ""}

	p.PutNow("in-buffer", none)
	var buffer = make([]byte, bufferSize)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err.Error())
	}

	if n > 0 {
		p.PutNow("in-buffer", &Token{BIN, buffer[0:n]})
		return &Token{BIN, buffer[0:n]}
	} else {
		return none
	}

}

func ReadPort(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == PORT && args[2].Tp == DATATYPE {
		var isHost = args[1].Ctx().GetNow("is-host")
		if isHost == nil || isHost.Tp != LOGIC || isHost.Val.(bool) {
			return &Token{ERR, "Target is not a conn"}, nil
		}

		if args[2].Uint8() != STRING && args[2].Uint8() != BIN {
			return &Token{ERR, "Error output type"}, nil
		}

		var inBuffer = args[1].Ctx().GetNow("in-buffer")
		if inBuffer.Tp == NONE {
			return inBuffer, nil
		}

		if args[2].Uint8() == BIN {
			return inBuffer, nil
		} else if args[2].Uint8() == STRING {
			return &Token{STRING, string(inBuffer.Val.([]byte))}, nil
		}

	}

	return &Token{ERR, "Type Mismatch"}, nil
}

func WritePort(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == PORT && (args[2].Tp == STRING || args[2].Tp == BIN || args[2].Tp == BLOCK) {
		var protocol = args[1].Ctx().GetNow("protocol")
		if protocol != nil && protocol.Str() == "mysql" {
			return writeMysql(args[1], args[2], args[3].ToBool())
		}

		var isHost = args[1].Ctx().GetNow("is-host")
		if isHost == nil || isHost.Tp != LOGIC || isHost.Val.(bool) {
			return &Token{ERR, "Target is not a conn"}, nil
		}
		var outBuffer []byte
		if args[2].Tp == BIN {
			outBuffer = args[2].Val.([]byte)
		} else if args[2].Tp == STRING {
			outBuffer = []byte(args[2].Str())
		} else {
			return &Token{ERR, "Type Mismatch"}, nil
		}

		n, err := args[1].Ctx().GetNow("port").Val.(net.Conn).Write(outBuffer)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		return &Token{INTEGER, n}, nil
	}

	return &Token{ERR, "Type Mismatch"}, nil
}

func Wait(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == PORT {
		var isHost = args[1].Ctx().GetNow("is-host")
		if isHost.Tp == LOGIC {
			if isHost.Val.(bool) {
				return waitListener(args[1].Ctx().GetNow("port").Val.(net.Listener), args[1].Ctx(), es), nil
			} else {
				return waitConn(args[1].Ctx().GetNow("port").Val.(net.Conn), args[1].Ctx(), es), nil
			}

		}
	}
	return &Token{ERR, "Type Mismatch"}, nil
}

func Listen(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == PORT {
		var isHost = args[1].Ctx().GetNow("is-host")
		if isHost.Tp == LOGIC {
			if isHost.Val.(bool) {
				listenListener(args[1].Ctx().GetNow("port").Val.(net.Listener), args[1].Ctx(), es)
				return &Token{NIL, nil}, nil
			} else {
				listenConn(args[1].Ctx().GetNow("port").Val.(net.Conn), args[1].Ctx(), es)
				return &Token{NIL, nil}, nil
			}
		}
	}
	return &Token{ERR, "Type Mismatch"}, nil
}

func Close(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == PORT {
		var isHost = args[1].Ctx().GetNow("is-host")
		if isHost.Tp == LOGIC {

			if args[1].Ctx().GetNow("protocol").Str() == "mysql" {
				err := args[1].Ctx().GetNow("port").Val.(*sql.DB).Close()
				if err != nil {
					return &Token{ERR, err.Error()}, nil
				}
			} else {
				if isHost.Val.(bool) {
					err := args[1].Ctx().GetNow("port").Val.(net.Listener).Close()
					if err != nil {
						return &Token{ERR, err.Error()}, nil
					}
				} else {
					err := args[1].Ctx().GetNow("port").Val.(net.Conn).Close()
					if err != nil {
						return &Token{ERR, err.Error()}, nil
					}
				}
			}

			var closeCode = args[1].Ctx().GetNow("on-close")
			if closeCode != nil && closeCode.Tp == BLOCK && closeCode.List().Len() > 0 {
				return es.Eval(closeCode.Tks(), ctx)
			}
			return &Token{NIL, nil}, nil
		}
	}
	return &Token{ERR, "Type Mismatch"}, nil
}

func newMysqlPort(db *sql.DB, ctx *BindMap) *Token {
	var p = BindMap{make(map[string]*Token, 8), ctx, USR_CTX, sync.RWMutex{}}

	p.PutNow("port", &Token{NONE, db})
	p.PutNow("is-host", &Token{LOGIC, false})
	p.PutNow("protocol", &Token{STRING, "mysql"})
	p.PutNow("on-close", &Token{NONE, "none"})

	return &Token{PORT, &p}
}

func writeMysql(port *Token, arg *Token, colName bool) (*Token, error) {
	var db = port.Ctx().GetNow("port").Val.(*sql.DB)

	if arg.Tp == STRING {
		var sqlStr = Trim(arg.Str())
		var sqlStrSlice = StrCut(sqlStr)
		var sqlType = ""
		if len(sqlStrSlice) > 0 {
			sqlType = strings.ToUpper(sqlStrSlice[0])
		}

		switch sqlType {
		case "SELECT":
			rows, err := db.Query(sqlStr)
			if err != nil {
				return &Token{ERR, err.Error()}, nil
			}
			return rowsPacker(rows, colName), nil

		default:
			rst, err := db.Exec(sqlStr)
			if err != nil {
				return &Token{ERR, err.Error()}, nil
			}
			affected, err := rst.RowsAffected()
			if err != nil {
				return &Token{ERR, err.Error()}, nil
			}
			return &Token{INTEGER, int(affected)}, nil
		}

	} else if arg.Tp == BLOCK && arg.List().Len() > 0 {
		var sqlStr = Trim(arg.Tks()[0].Str())
		var args []interface{}
		var sqlStrSlice = StrCut(sqlStr)
		var sqlType = ""
		if len(sqlStrSlice) > 0 {
			sqlType = strings.ToUpper(sqlStrSlice[0])
		}

		for idx := 1; idx < arg.List().Len(); idx++ {
			args = append(args, arg.Tks()[idx].Val)
		}

		switch sqlType {
		case "SELECT":
			rows, err := db.Query(sqlStr, args...)
			if err != nil {
				return &Token{ERR, err.Error()}, nil
			}
			return rowsPacker(rows, colName), nil

		default:
			rst, err := db.Exec(sqlStr, args...)
			if err != nil {
				return &Token{ERR, err.Error()}, nil
			}
			affected, err := rst.RowsAffected()
			if err != nil {
				return &Token{ERR, err.Error()}, nil
			}
			return &Token{INTEGER, int(affected)}, nil
		}

	}

	return &Token{ERR, "Type Mismatch"}, nil
}

func rowsPacker(rows *sql.Rows, colName bool) *Token {
	defer rows.Close()
	var cols, err = rows.Columns()
	if err != nil {
		return &Token{ERR, err.Error()}
	}

	var result = &Token{BLOCK, NewTks(8)}
	for rows.Next() {
		var row = make([]interface{}, len(cols))
		var rowRef = make([]interface{}, len(cols))
		for idx, _ := range row {
			rowRef[idx] = &row[idx]
		}
		err := rows.Scan(rowRef...)
		if err != nil {
			return &Token{ERR, err.Error()}
		}

		var rst = &Token{BLOCK, NewTks(8)}
		for idx, item := range row {
			if colName {
				rst.List().Add(&Token{SET_WORD, cols[idx]})
			}
			switch reflect.TypeOf(item) {
			case reflect.TypeOf(int(0)):
				rst.List().Add(&Token{INTEGER, item.(int)})
			case reflect.TypeOf(int8(0)):
				rst.List().Add(&Token{INTEGER, int(item.(int8))})
			case reflect.TypeOf(int16(0)):
				rst.List().Add(&Token{INTEGER, int(item.(int16))})
			case reflect.TypeOf(int32(0)):
				rst.List().Add(&Token{INTEGER, int(item.(int32))})
			case reflect.TypeOf(int64(0)):
				rst.List().Add(&Token{INTEGER, int(item.(int64))})
			case reflect.TypeOf(uint(0)):
				rst.List().Add(&Token{INTEGER, int(item.(uint))})
			case reflect.TypeOf(uint8(0)):
				rst.List().Add(&Token{INTEGER, int(item.(uint8))})
			case reflect.TypeOf(uint16(0)):
				rst.List().Add(&Token{INTEGER, int(item.(uint16))})
			case reflect.TypeOf(uint32(0)):
				rst.List().Add(&Token{INTEGER, int(item.(uint32))})
			case reflect.TypeOf(uint64(0)):
				rst.List().Add(&Token{INTEGER, int(item.(uint64))})
			case reflect.TypeOf(false):
				rst.List().Add(&Token{LOGIC, item.(bool)})
			case reflect.TypeOf(byte('0')):
				rst.List().Add(&Token{BIN, item.(byte)})
			case reflect.TypeOf(rune('0')):
				rst.List().Add(&Token{CHAR, item.(rune)})
			case reflect.TypeOf(time.Now()):
				rst.List().Add(&Token{TIME, ParseTimeStr(item.(time.Time).Format("2006-01-02+15:04:05"))})
			case reflect.TypeOf(float64(0.0)):
				rst.List().Add(&Token{DECIMAL, item.(float64)})
			case reflect.TypeOf(float32(0.0)):
				rst.List().Add(&Token{DECIMAL, float64(item.(float32))})
			default:
				rst.List().Add(&Token{STRING, string(item.([]byte))})
			}
		}
		result.List().Add(rst)
	}

	return result
}
