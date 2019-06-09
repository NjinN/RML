import evalline
import totoken

type
    EvalUnit* = object
        mainCtx*: ref Context
        nowLine*: ref EvalLine

proc newEvalUnit*():ref EvalUnit=
    result = new(EvalUnit)
    result.mainCtx = newContext()
    result.nowLine = newEvalLine(10, nil)
    return result

proc eval*(u: var ref EvalUnit, s: string):string=
    var inp = toTokens(s)
    if(len(inp) == 0):
        return

    var temp = new(Token)
    var idx = 0

    while idx < len(inp):
        var nowToken = inp[idx]
        var nextToken: ref Token
        if idx < len(inp)-1:
            nextToken = getVal(inp[idx+1], u.mainCtx)
        if not isNil(nextToken) and nextToken.tp == TypeEnum.op:
            if isNil(u.nowLine.line[0]) or (u.nowLine.line[0].tp != TypeEnum.op):
                var newLine = newEvalLine(3, u.nowLine)
                newLine.idx = 2
                newLine.line[0] = nextToken
                newLine.line[1] = getVal(nowToken, u.mainCtx)
                u.nowLine = newLine
                idx += 1
            else:
                var newLine = newEvalLine(3, u.nowLine.father)
                newLine.idx = 1
                newLine.line[0] = nextToken
                u.nowLine.father = newLine
                u.nowLine.line[u.nowLine.idx] = getVal(nowToken, u.mainCtx)
                u.nowLine.idx += 1
                idx += 1
        else:
            if nowToken.tp == TypeEnum.word:
                nowToken = getVal(nowToken, u.mainCtx)
            if nowToken.tp == TypeEnum.op:
                if isNil(u.nowLine.line[0]):
                    return "Error: illegal grammar"
                if u.nowLine.idx > 0:
                    u.nowLine.idx -= 1
                var newLine = newEvalLine(3, u.nowLine)
                newLine.idx = 2
                newLine.line[0] = nowToken
                newLine.line[1] = u.nowLine.line[u.nowLine.idx]
                u.nowLine = newLine
            elif nowToken.tp < TypeEnum.set_word:
                u.nowLine.line[u.nowLine.idx] = getVal(nowToken, u.mainCtx)
                if not isNil(u.nowLine.father):
                    u.nowLine.idx += 1
            else:
                var newLine = newEvalLine(nowToken.explen.int, u.nowLine)
                newLine.idx = 1
                newLine.line[0] = nowToken
                u.nowLine = newLine
            
        while not isNil(u.nowLine.line[0]) and (u.nowLine.idx.uint16 == u.nowLine.line[0].explen):
            temp = u.nowLine.eval(u.mainCtx)
            if not isNil(temp) and temp.tp == TypeEnum.err:
                result = $temp.val.string & "\n"
                var i = 0
                while i < len(u.nowLine.line) and (not isNil(u.nowLine.line[i])):
                    result = result & $u.nowLine.line[i].toStr & " "
                    i += 1
            if not isNil(u.nowLine.father):
                u.nowLine = u.nowLine.father
                if not isNil(temp):
                    u.nowLine.line[u.nowLine.idx] = temp
                    if not isNil(u.nowLine.father):
                        u.nowLine.idx += 1
            else:
                break
        idx += 1
    if isNil(u.nowLine.line[0]) or isNil(temp):
        return ""
    elif isNil(u.nowLine.father) and (u.nowLine.line[0].explen == 1):
        return $u.nowLine.line[0].toStr
    else:
        u.nowLine.idx = 0
        u.nowLine.father = nil
        result = "Error: incomplete expression \n-- Near: "
        var i = 0
        while i < len(u.nowLine.line) and (not isNil(u.nowLine.line[i])):
            result = result & $u.nowLine.line[i].toStr & " "
            i += 1
        return result
        

            


when isMainModule:
    var unit = newEvalUnit()

    discard unit.eval("i: 123 j: 456")
    # echo(isNil(unit.mainCtx.map[cstring("i")]))
    for item in unit.mainCtx.map.keys:
        echo(item)


