package nativelib

import . "../core"
import "strings"
import "io/ioutil"
import "path/filepath"

func Load(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1] != nil && args[1].Tp == FILE {
		var path = args[1].Str()
		if strings.IndexByte(path, '"') >=0 {
			path = strings.ReplaceAll(path, `"`, ``)
		}
		filePath, err := filepath.Abs(path)
		if err != nil {
			return &Token{ERR, "Error File Path!"}, nil
		}
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			return &Token{ERR, "Error when reading the file"}, nil
		}

		result.Tp = BLOCK
		result.Val = ToTokens(string(fileData), ctx, es)
		
		return &result, nil

	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}


func Read(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1].Tp == FILE {
		var path = args[1].Str()
		if strings.IndexByte(path, '"') >=0 {
			path = strings.ReplaceAll(path, `"`, ``)
		}
		filePath, err := filepath.Abs(path)
		if err != nil {
			return &Token{ERR, "Error File Path!"}, nil
		}
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			return &Token{ERR, "Error when reading the file"}, nil
		}

		if args[2].Tp == DATATYPE {
			if args[2].Int() == BIN {
				result.Tp = BIN
				result.Val = fileData
				return &result, nil
			}else if args[2].Int() == STRING {
				result.Tp = STRING
				result.Val = string(fileData)
				return &result, nil
			}

		}


		result.Tp = STRING
		result.Val = string(fileData)
		
		return &result, nil

	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

