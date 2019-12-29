package nativelib

import . "../core"
import "strings"
import "os"
import "io/ioutil"
import "path/filepath"
// import "fmt"

func NowDir(es *EvalStack, ctx *BindMap) (*Token, error){
	var result Token
	result.Tp = FILE
	result.Val, _ = os.Getwd()
	result.Val = result.Val.(string) + "/"
	return &result, nil
}

func ChangeDir(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	
	filePath, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
	if err != nil {
		return &Token{ERR, err.Error()}, nil
	}
	e := os.Chdir(filePath)
	if e != nil {
		return &Token{ERR, err.Error()}, nil
	}
	return &Token{LOGIC, true}, nil
}

func LsDir(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1] != nil && args[1].Tp == FILE {
		path, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		dir, err := os.Open(path)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		fileNames, err := dir.Readdirnames(-1)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		result.Tp = BLOCK
		result.Val = make([]*Token, len(fileNames))
		for i := 0; i < len(fileNames); i++ {
			fi, _ := os.Stat(path + "/" + fileNames[i]) 
			if fi.IsDir(){
				result.Tks()[i] = &Token{FILE, fileNames[i] + "/"}
			}else{
				result.Tks()[i] = &Token{FILE, fileNames[i]}
			}
		}
		return &result, nil

	}else{
		nowDir, _ := os.Getwd()
		dir, _ := os.Open(nowDir)
		fileNames, err := dir.Readdirnames(-1)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		result.Tp = BLOCK
		result.Val = make([]*Token, len(fileNames))
		for i := 0; i < len(fileNames); i++ {
			fi, _ := os.Stat(fileNames[i]) 
			if fi.IsDir(){
				result.Tks()[i] = &Token{FILE, fileNames[i] + "/"}
			}else{
				result.Tks()[i] = &Token{FILE, fileNames[i]}
			}
			
		}
		return &result, nil

	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func RenameFile(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1].Tp == FILE && args[2].Tp == FILE {
		oldPath, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		newPath, err := filepath.Abs(strings.ReplaceAll(args[2].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		
		err = os.Rename(oldPath, newPath)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		return &Token{LOGIC, true}, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func RemoveFile(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1] != nil && args[1].Tp == FILE {
		path, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		err = os.RemoveAll(path)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		return &Token{LOGIC, true}, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func makeDir(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1] != nil && args[1].Tp == FILE {
		path, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		err = os.MkdirAll(path, os.ModeDir|os.ModePerm)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		return &Token{LOGIC, true}, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

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
			return &Token{ERR, err.Error()}, nil
		}
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}

		result.Tp = BLOCK
		result.Val = ToTokens(string(fileData), ctx, es)
		
		return &result, nil

	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func FileExist(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1].Tp == FILE {
		var path = args[1].Str()
		if strings.IndexByte(path, '"') >=0 {
			path = strings.ReplaceAll(path, `"`, ``)
		}
		filePath, err := filepath.Abs(path)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		_, e := os.Stat(filePath)
		if os.IsNotExist(e) {
			return &Token{LOGIC, false}, nil
		}else{
			return &Token{LOGIC, true}, nil
		}
	}


	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}


func ReadFile(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1].Tp == FILE {
		var path = args[1].Str()
		if strings.IndexByte(path, '"') >=0 {
			path = strings.ReplaceAll(path, `"`, ``)
		}
		filePath, err := filepath.Abs(path)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
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


func WriteFile(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	var result Token

	if args[1].Tp == FILE && (args[2].Tp == STRING || args[2].Tp == BIN) {
		path, err := filepath.Abs(strings.ReplaceAll(args[1].Str(), `"`, ``))
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		var exist = false
		
		_, err = os.Stat(path)
		if err != nil {
			if os.IsNotExist(err){
				exist = false
			}else{
				return &Token{ERR, err.Error()}, nil
			}
		}else{
			exist = true
		}
		
		var data []byte
		if args[3].ToBool() && exist {
			data, err = ioutil.ReadFile(path)
			if err != nil {
				return &Token{ERR, err.Error()}, nil
			}

			if args[2].Tp == STRING {
				data = append(data, []byte(args[2].Str())...)
			}else if args[2].Tp == BIN {
				data = append(data, args[2].Val.([]byte)...)
			}

		}else{
			if args[2].Tp == STRING {
				data = []byte(args[2].Str())
			}else if args[2].Tp == BIN {
				data = args[2].Val.([]byte)
			}
		}

		err = ioutil.WriteFile(path, data, os.ModePerm)
		if err != nil {
			return &Token{ERR, err.Error()}, nil
		}
		return &Token{LOGIC, true}, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}
