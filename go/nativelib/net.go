package nativelib

import (
	"io/ioutil"
	"net/http"

	. "github.com/NjinN/RML/go/core"
)

// import "fmt"


func ReadUrl(es *EvalStack, ctx *BindMap) (*Token, error) {
	var args = es.Line[es.LastStartPos() : es.LastEndPos()]

	if args[1].Tp == URL {
		resp, err := http.Get(args[1].Str())
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		defer resp.Body.Close()
    	body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		if args[2].Tp == DATATYPE {
			if args[2].Uint8() == STRING {
				return &Token{STRING, string(body)}, nil
			} 
		}


		return &Token{BIN, body}, nil
	}


	return &Token{ERR, "Type Mismatch"}, nil
}

// func WriteUrl(es *EvalStack, ctx *BindMap) (*Token, error) {
// 	var args = es.Line[es.LastStartPos() : es.LastEndPos()]
// 	var result Token

// 	if args[1].Tp == FILE && (args[2].Tp == STRING || args[2].Tp == BIN) {
// 		path, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
// 		if err != nil {
// 			return &Token{ERR, err.Error()}, nil
// 		}
// 		var exist = false

// 		_, err = os.Stat(path)
// 		if err != nil {
// 			if os.IsNotExist(err) {
// 				exist = false
// 			} else {
// 				return &Token{ERR, err.Error()}, nil
// 			}
// 		} else {
// 			exist = true
// 		}

// 		var data []byte
// 		if args[3].ToBool() && exist {
// 			data, err = ioutil.ReadFile(path)
// 			if err != nil {
// 				return &Token{ERR, err.Error()}, nil
// 			}

// 			if args[2].Tp == STRING {
// 				data = append(data, []byte(args[2].Str())...)
// 			} else if args[2].Tp == BIN {
// 				data = append(data, args[2].Val.([]byte)...)
// 			}

// 		} else {
// 			if args[2].Tp == STRING {
// 				data = []byte(args[2].Str())
// 			} else if args[2].Tp == BIN {
// 				data = args[2].Val.([]byte)
// 			}
// 		}

// 		err = ioutil.WriteFile(path, data, os.ModePerm)
// 		if err != nil {
// 			return &Token{ERR, err.Error()}, nil
// 		}
// 		return &Token{LOGIC, true}, nil
// 	}

// 	result.Tp = ERR
// 	result.Val = "Type Mismatch"
// 	return &result, nil
// }
