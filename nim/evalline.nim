import token
export token


type
    EvalLine* = object 
        idx*: int
        line*: ptr List[ptr Token]
        father*: ptr EvalLine

    EvalUnit* = object
        mainCtx*: ptr BindMap[ptr Token]
        nowLine*: ptr EvalLine

proc newEvalLine*(size: int = 8, father: ptr EvalLine = nil):ptr EvalLine=
    result = cast[ptr EvalLine](alloc0(sizeof(EvalLine)))
    result.idx = 0
    result.line = newList[ptr Token](size)
    result.father = father
    return result

proc freeEvalLine*(l: ptr EvalLine)=
    freeList(l.line)
    dealloc(l)


proc newEvalUnit*(cont: ptr BindMap[ptr Token]):ptr EvalUnit
proc freeEvalUnit*(u: ptr EvalUnit)
proc eval*(u: var ptr EvalUnit, inp: ptr List[ptr Token]):ptr Token

proc run*(f: var ptr Func; args: var ptr List[ptr Token], c: ptr BindMap[ptr Token]):ptr Token=
    var cont = newBindMap[ptr Token](16)
    cont.father = c
    for idx in 0..high(f.args):
        cont[f.args[idx].toStr] = args[idx + 1]
    var unit = newEvalUnit(cont)
    result = unit.eval(f.body)
    return result

proc eval*(l: var ptr EvalLine;c: ptr BindMap[ptr Token]):ptr Token=
    try:
        case l.line[0].tp
        of TypeEnum.set_word:
            if not isNil(l.line[1]):
                c[l.line[0].val.string] = l.line[1]
                return l.line[1]
            else:
                result = newToken(TypeEnum.err)
                result.val.string = "Illegal grammar!!!"
        of TypeEnum.native:
            return l.line[0].val.exec.run(l.line, c)
        of TypeEnum.op:
            return l.line[0].val.exec.run(l.line)
        of TypeEnum.function:
            return l.line[0].val.fc.run(l.line, c)
        else:
            return nil
    except:
        result = newToken(TypeEnum.err)
        result.val.string = "Illegal grammar!!!"

proc getFinalToken*(t: ptr Token, c: ptr BindMap[ptr Token]):ptr Token=
    result = nil

    case t.tp
    of TypeEnum.lit_word:
        result = newToken()
        result.tp = TypeEnum.word
        result.val.string = t.val.string[1..len(t.val.string)-1]
        return result
    of TypeEnum.word:
        var cont = c
        while isNil(result) and (not isNil(cont)):
            result = cont[t.val.string]
            cont = cont.father
        if isNil(result):
            result = newToken()
            result.tp = TypeEnum.none
            result.val.string = "none"
            return result
        elif result.explen == 1:
            case result.tp
            of TypeEnum.native:
                var temp = newList[ptr Token]()
                freeList(temp)
                result = result.val.exec.run(temp, c)
            of TypeEnum.function:
                var temp = newList[ptr Token]()
                result = result.val.fc.run(temp, c)
                freeList(temp)
            else:
                return result  
        return result
    of TypeEnum.paren:
        var unit = newEvalUnit(c)
        result = unit.eval(t.val.list)
        freeEvalUnit(unit)
        return result
    else:
        # echo("return origin")
        return t


# when isMainModule:
#     var line = newEvalLine()
#     var tk = new(Token)
#     tk.tp = TypeEnum.integer
#     #tk.val.integer = 123
#     tk.explen = 1
#     line.line[0] = tk
#     echo(repr(line))