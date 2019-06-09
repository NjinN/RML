import tables
import types

export types
export tables

type 
    TokenVal* {.union.} = object
        byte*: byte
        char*: char
        integer*: int32
        long*: int64
        float*: float64
        string*: cstring
        integerArr*: array[0..3, int32]
        longArr*: array[0..1, int64]
        floatArr*: array[0..1, float64]
        token*: ref Token
        list*: seq[ref Token]
        exec*: proc(args: var seq[ref Token]):ref Token  

    Context* = object
        map*: TableRef[cstring, ref Token]
        father*: ref Context

    Token*  = object
        tp*: TypeEnum
        val*: TokenVal
        explen*: uint16
        context*: ref Context


proc newContext*(size = 32):ref Context=
    result = new(Context)
    result.map = newTable[cstring, ref Token](size)
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
    of TypeEnum.integer:
        return cstring($t.val.integer)
    of TypeEnum.float:
        return cstring($t.val.float)
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
    of TypeEnum.word:
        return t.val.string
    of TypeEnum.set_word:
        return cstring($t.val.string & ":")
    of TypeEnum.native:
        return cstring("native")
    of TypeEnum.function:
        return cstring("function")
    of TypeEnum.op:
        return t.val.string

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
    of TypeEnum.integer:
        return $t.val.integer
    of TypeEnum.float:
        return $t.val.float
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
    of TypeEnum.word:
        return $t.val.string
    of TypeEnum.set_word:
        return $t.val.string & ":"
    of TypeEnum.native:
        return "native"
    of TypeEnum.function:
        return "function"
    of TypeEnum.op:
        return $t.val.string

proc repr*(t: ref Token):string=
    result = "type = " & $t.tp & "\n"
    result = result & "val = " & $t.toStr & "\n"
    result = result & "explen = " & $t.explen & "\n"
    result = result & repr(t.context) & "\n"
    return result

proc getVal*(t: ref Token, c: ref Context):ref Token=
    result = new(Token)
    result.tp = TypeEnum.none
    result.val.string = "none"
    result.explen = 1
    case t.tp
    of TypeEnum.lit_word:
        result.tp = TypeEnum.word
        result.val.string = cstring(($t.val.string)[1..len(t.val.string)-1])
        result.explen = 1
        return result
    of TypeEnum.word:
        var cont = c
        while result.tp == TypeEnum.none and (not isNil(cont)):
            result = cont.map.getOrDefault(t.val.string, result)
            cont = cont.father
        if result.tp >= TypeEnum.native and result.explen == 1:
            var temp = newSeq[ref Token]()
            result = result.val.exec(temp)
        return result
    else:
        return t



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
   




    


    


