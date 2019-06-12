import token
export token

import tables


type
    EvalLine* = object 
        idx*: int
        line*: seq[ref Token]
        father*: ref EvalLine

    EvalUnit* = object
        mainCtx*: ref Context
        nowLine*: ref EvalLine

proc newEvalLine*(size: int = 10, father: ref EvalLine = nil):ref EvalLine=
    result = new(EvalLine)
    result.idx = 0
    result.line = newSeq[ref Token](size)
    result.father = father
    return result


proc newEvalUnit*(cont: ref Context):ref EvalUnit
proc eval*(u: var ref EvalUnit, inp: seq[ref Token]):ref Token

proc run*(f: var ref Func; args: var seq[ref Token], c: ref Context):ref Token=
    var cont = newContext(16)
    cont.father = c
    for idx in 0..len(f.args)-1:
        cont.map[f.args[idx].toStr] = args[idx + 1]
    var unit = newEvalUnit(cont)
    result = unit.eval(f.body)
    return result

proc eval*(l: var ref EvalLine;c: ref Context):ref Token=
    try:
        case l.line[0].tp
        of TypeEnum.set_word:
            if not isNil(l.line[1]):
                c.map[l.line[0].val.string] = l.line[1]
                return l.line[1]
            else:
                result = newToken(TypeEnum.err, 1)
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
        result = newToken(TypeEnum.err, 1)
        result.val.string = "Illegal grammar!!!"

proc getFinalToken*(t: ref Token, c: ref Context):ref Token=
    result = new(Token)
    result.tp = TypeEnum.none
    result.val.string = "none"
    result.explen = 1
    # print t
    # echo(repr(t))
    # echo(t.tp)
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
        if result.explen == 1:
            case result.tp
            of TypeEnum.native:
                var temp = newSeq[ref Token]()
                result = result.val.exec.run(temp, c)
            of TypeEnum.function:
                var temp = newSeq[ref Token]()
                result = result.val.fc.run(temp, c)
            else:
                return result  
        return result
    of TypeEnum.paren:
        var unit = newEvalUnit(c)
        result = unit.eval(t.val.list)
        return result
    else:
        # echo("return origin")
        return t


when isMainModule:
    var line = newEvalLine()
    var tk = new(Token)
    tk.tp = TypeEnum.integer
    #tk.val.integer = 123
    tk.explen = 1
    line.line[0] = tk
    echo(repr(line))