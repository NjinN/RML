import tables
import types

export types
export tables

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
        token*: ref Token
        list*: seq[ref Token]
        exec*: ref Exec
        fc*: ref Func

    Context* = object
        map*: TableRef[cstring, ref Token]
        father*: ref Context

    Token*  = object
        tp*: TypeEnum
        val*: TokenVal
        explen*: uint16
        context*: ref Context

    Exec* = object
        string*: cstring
        run*: proc(args: var seq[ref Token], cont: ref Context = nil):ref Token 

    Func* = object 
        args*: seq[ref Token]
        body*: seq[ref Token]

proc newToken*(tp: TypeEnum, l: int):ref Token=
    result = new Token
    result.tp = tp
    result.explen = l.uint16
    return result

proc newContext*(size = 32):ref Context=
    result = new Context 
    result.map = newTable[cstring, ref Token](size)
    return result 

proc newExec*(s: string, f: proc(args: var seq[ref Token], cont: ref Context = nil):ref Token):ref Exec=
    result = new Exec
    result.string = cstring(s)
    result.run = f
    return result

proc newFunc*(args: seq[ref Token], body: seq[ref Token]):ref Func=
    result = new Func
    result.args = args
    result.body = body
    return result


proc toStr*(t: ref Token):cstring=
    case t.tp
    of TypeEnum.none:
        return cstring("none")
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
        return "\"" & $t.val.string & "\""
    of TypeEnum.list:
        result = "[ "
        for item in t.val.list:
            result = $result & $toStr(item) & " "
        result = $result & "]"
        return result
    of TypeEnum.paren:
        result = "( "
        for item in t.val.list:
            result = $result & $toStr(item) & " "
        result = $result & ")"
        return result
    of TypeEnum.word:
        return t.val.string
    of TypeEnum.set_word:
        return cstring($t.val.string & ":")
    of TypeEnum.native:
        return t.val.exec.string
    of TypeEnum.function:
        var str = "func [ "
        for i in 0..len(t.val.fc.args)-1:
            str = str & $t.val.fc.args[i].toStr & " "
        str = str & "] [ "
        for i in 0..len(t.val.fc.body)-1:
            str = str & $t.val.fc.body[i].toStr & " "
        str = str & "]"
        return cstring(str)
        # return cstring("function")
    of TypeEnum.op:
        return t.val.exec.string

proc print*(t: ref Token)=
    echo(toStr(t))


proc outputStr*(t: ref Token):string=
    case t.tp
    of TypeEnum.none:
        return "none"
    of TypeEnum.err:
        return "Error: " & $t.val.string
    of TypeEnum.lit_word:
        return "'" & $t.val.string
    of TypeEnum.get_word:
        return $t.val.string & ":"
    of TypeEnum.datatype:
        return $t.val.string
    of TypeEnum.logic:
        return $t.val.logic
    of TypeEnum.integer:
        return $t.val.integer
    of TypeEnum.decimal:
        return $t.val.decimal
    of TypeEnum.char:
        return $t.val.char
    of TypeEnum.string:
        return $t.val.string 
    of TypeEnum.list:
        result = "[ "
        for item in t.val.list:
            result = result & outputStr(item) & " "
        result = result & "]"
        return result
    of TypeEnum.paren:
        result = "( "
        for item in t.val.list:
            result = result & outputStr(item) & " "
        result = result & ")"
        return result
    of TypeEnum.word:
        return $t.val.string
    of TypeEnum.set_word:
        return $t.val.string & ":"
    of TypeEnum.native:
        return $t.val.exec.string
    of TypeEnum.function:
        return "function"
    of TypeEnum.op:
        return $t.val.exec.string

proc repr*(t: ref Token):string=
    result = "type = " & $t.tp & "\n"
    result = result & "val = " & $t.toStr & "\n"
    result = result & "explen = " & $t.explen & "\n"
    result = result & repr(t.context) & "\n"
    return result




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

    proc fc(arg: var seq[ref Token]):ref Token=
        return nil

    var token = new(Token)
    token.tp = TypeEnum.function
    token.val.exec = fc
    var temp = newSeq[ref Token]()
    discard token.val.exec(temp)
   




    


    


