import token
export token

import tables


type
    EvalLine* = object 
        idx*: int
        line*: seq[ref Token]
        father*: ref EvalLine

proc newEvalLine*(size: int = 10, father: ref EvalLine = nil):ref EvalLine=
    result = new(EvalLine)
    result.idx = 0
    result.line = newSeq[ref Token](size)
    result.father = father
    return result

proc eval*(l: var ref EvalLine;c: ref Context):ref Token=
    case l.line[0].tp
    of TypeEnum.set_word:
        c.map[l.line[0].val.string] = l.line[1]
        return l.line[1]
    of TypeEnum.native:
        echo("start eval")
        return l.line[0].val.exec.run(l.line)
    of TypeEnum.op:
        return l.line[0].val.exec.run(l.line)
    else:
        return nil


when isMainModule:
    var line = newEvalLine()
    var tk = new(Token)
    tk.tp = TypeEnum.integer
    #tk.val.integer = 123
    tk.explen = 1
    line.line[0] = tk
    echo(repr(line))