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

proc eval*(u: var ref EvalUnit, s: string):ref Token=
    result = new Token
    result.tp = TypeEnum.string
    result.explen = 1
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
                    result.val.string = cstring("Error: illegal grammar")
                    return result
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
                result.val.string = cstring($temp.val.string & "\n-->Near: ")
                var i = 0
                if idx >= 3:
                    for i in 0..2:
                        result.val.string = cstring($result.val.string & $inp[idx-i].toStr & " ")
                else:
                    result.val.string = cstring($result.val.string & $inp[idx].toStr)
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
        result.val.string = ""
    elif isNil(u.nowLine.father) and (u.nowLine.line[0].explen == 1):
        result.val.string = u.nowLine.line[0].toStr
    else:
        u.nowLine.idx = 0
        u.nowLine.father = nil
        result.val.string = "Error: incomplete expression \n-- Near: "
        var i = 0
        if idx >= 3:
            for i in 0..2:
                result.val.string = cstring($result.val.string & $inp[idx-i].toStr & " ")
        else:
            result.val.string = cstring($result.val.string & $inp[idx].toStr)
    return result
        

            


when isMainModule:
    var unit = newEvalUnit()

    discard unit.eval("i: 123 j: 456")
    # echo(isNil(unit.mainCtx.map[cstring("i")]))
    for item in unit.mainCtx.map.keys:
        echo(item)


