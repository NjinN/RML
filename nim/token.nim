import listType
import bindMap
import tSet
import types

export listType
export bindMap
export tSet
export types


type 
    TokenVal* {.union.} = object
        logic*: bool
        byte*: byte
        char*: char
        integer*: int32
        long*: int64
        decimal*: float64
        string*: cstring
        integerArr*: array[0..3, int32]
        longArr*: array[0..1, int64]
        floatArr*: array[0..1, float64]
        token*: ptr Token
        list*: ptr List[ptr Token]
        exec*: ptr Exec
        fc*: ptr Func

    Token* {.packed.}  = object
        tp*: TypeEnum
        val*: TokenVal

    Exec* = object
        string*: cstring
        explen*: uint16
        run*: proc(args:ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil, unit: ptr EvalUnit = nil):ptr Token 

    Func* = object 
        args*: ptr List[ptr Token]
        body*: ptr List[ptr Token]
        explen*: uint16

    EvalLine* = object 
        idx*: int
        line*: ptr List[ptr Token]
        father*: ptr EvalLine

    EvalUnit* = object
        mainCtx*: ptr BindMap[ptr Token]
        nowLine*: ptr EvalLine
        father*: ptr EvalUnit

proc newToken*():ptr Token=
    result = cast[ptr Token](alloc0(sizeof(Token)))
    return result

proc newToken*(tp: TypeEnum):ptr Token=
    result = cast[ptr Token](alloc0(sizeof(Token)))
    result.tp = tp
    return result


proc newExec*(s: cstring, f: proc(args:ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil, unit: ptr EvalUnit = nil):ptr Token, l: int):ptr Exec=
    result = cast[ptr Exec](alloc0(sizeof(Exec)))
    result.string = s
    result.run = f
    result.explen = l.uint16
    return result

proc newFunc*(args: ptr List[ptr Token], body: ptr List[ptr Token]):ptr Func=
    result = cast[ptr Func](alloc0(sizeof(Func)))
    result.args = args
    result.body = body
    result.explen = uint16(high(args) + 2)
    return result


proc toStr*(t: ptr Token):cstring=
    case t.tp
    of TypeEnum.null:
        return "null"
    of TypeEnum.none:
        return "none"
    of TypeEnum.err:
        return cstring("Error: " & $t.val.string)
    of TypeEnum.lit_word:
        return cstring("'" & $t.val.string)
    of TypeEnum.get_word:
        return cstring($t.val.string & ":")
    of TypeEnum.datatype:
        return t.val.string
    of TypeEnum.logic:
        return cstring($t.val.logic)
    of TypeEnum.integer:
        return cstring($t.val.integer)
    of TypeEnum.decimal:
        return cstring($t.val.decimal)
    of TypeEnum.char:
        return cstring($t.val.char)
    of TypeEnum.string:
        return cstring("\"" & $t.val.string & "\"")
    of TypeEnum.list:
        result = "[ "
        for i in 0..high(t.val.list):
            result = cstring($result & $toStr(t.val.list[i]) & " ")
        result = cstring($result & "]")
        return result
    of TypeEnum.paren:
        result = "( "
        for i in 0..high(t.val.list):
            result = cstring($result & $toStr(t.val.list[i]) & " ")
        result = cstring($result & ")")
        return result
    of TypeEnum.word:
        return t.val.string
    of TypeEnum.set_word:
        return cstring($t.val.string & ":")
    of TypeEnum.native:
        return t.val.exec.string
    of TypeEnum.function:
        var str = cstring("func [ ")
        for i in 0..high(t.val.fc.args):
            str = cstring($str & $t.val.fc.args[i].toStr & " ")
        str = cstring($str & "] [ ")
        for i in 0..high(t.val.fc.body):
            str = cstring($str & $t.val.fc.body[i].toStr & " ")
        str = cstring($str & "]")
        return str
        # return "function"
    of TypeEnum.op:
        return t.val.exec.string

proc print*(t: ptr Token)=
    echo(toStr(t))


proc outputStr*(t: ptr Token):cstring=
    case t.tp
    of TypeEnum.null:
        return "null"
    of TypeEnum.none:
        return "none"
    of TypeEnum.err:
        return cstring("Error: " & $t.val.string)
    of TypeEnum.lit_word:
        return cstring("'" & $t.val.string)
    of TypeEnum.get_word:
        return cstring($t.val.string & ":")
    of TypeEnum.datatype:
        return t.val.string
    of TypeEnum.logic:
        return cstring($t.val.logic)
    of TypeEnum.integer:
        return cstring($t.val.integer)
    of TypeEnum.decimal:
        return cstring($t.val.decimal)
    of TypeEnum.char:
        return cstring($t.val.char)
    of TypeEnum.string:
        return t.val.string
    of TypeEnum.list:
        result = "[ "
        for i in 0..high(t.val.list):
            result = cstring($result & $toStr(t.val.list[i]) & " ")
        result = cstring($result & "]")
        return result
    of TypeEnum.paren:
        result = "( "
        for i in 0..high(t.val.list):
            result = cstring($result & $toStr(t.val.list[i]) & " ")
        result = cstring($result & ")")
        return result
    of TypeEnum.word:
        return t.val.string
    of TypeEnum.set_word:
        return cstring($t.val.string & ":")
    of TypeEnum.native:
        return t.val.exec.string
    of TypeEnum.function:
        var str = cstring("func [ ")
        for i in 0..high(t.val.fc.args):
            str = cstring($str & $t.val.fc.args[i].toStr & " ")
        str = cstring($str & "] [ ")
        for i in 0..high(t.val.fc.body):
            str = cstring($str & $t.val.fc.body[i].toStr & " ")
        str = cstring($str & "]")
        return str
        # return "function"
    of TypeEnum.op:
        return t.val.exec.string

proc repr*(t: ptr Token):string=
    result = "type = " & $t.tp & "\n"
    result = result & "val = " & $t.toStr & "\n"
    return result

proc explen*(t: ptr Token):int=
    case t.tp
    of TypeEnum.null:
        return 1
    of TypeEnum.none:
        return 1
    of TypeEnum.err:
        return 1
    of TypeEnum.lit_word:
        return 1
    of TypeEnum.get_word:
        return 1
    of TypeEnum.datatype:
        return 1
    of TypeEnum.logic:
        return 1
    of TypeEnum.integer:
        return 1
    of TypeEnum.decimal:
        return 1
    of TypeEnum.char:
        return 1
    of TypeEnum.string:
        return 1
    of TypeEnum.list:
        return 1
    of TypeEnum.paren:
        return 1
    of TypeEnum.word:
        return 1
    of TypeEnum.set_word:
        return 2
    of TypeEnum.native:
        return t.val.exec.explen.int
    of TypeEnum.function:
        return t.val.fc.explen.int
    of TypeEnum.op:
        return t.val.exec.explen.int



when isMainModule:
    # var token = new(Token)
    # token.tp = TypeEnum.integer
    # token.val.integer = 123
    # token.explen = 1

    # echo(repr(token))

    # var cont = newContext(2)

    # cont.map[cstring("123")] = token

    # var cont: Context
    # cont.map = newTable[cstring, ref Token](2)
    # echo(isNil(cont.father))

    # cont.map[cstring("1")] = token
    # echo(cont.map[cstring("1")].val.string)

    # for i in 5..10:
    #     cont.map[cstring($i)] = token
    
    # echo(len(cont.map))

    echo(sizeof Token)

    # proc fc(arg: var seq[ref Token]):ref Token=
    #     return nil

    # var token = new(Token)
    # token.tp = TypeEnum.function
    # token.val.exec = fc
    # var temp = newSeq[ref Token]()
    # discard token.val.exec(temp)
   




    


    


