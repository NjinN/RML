# import evalline
import totoken

include "evalLine.nim"
 

proc newEvalUnit*():ptr EvalUnit=
    result = cast[ptr EvalUnit](alloc0(sizeof(EvalUnit)))
    result.mainCtx = newBindMap[ptr Token](16)
    result.nowLine = newEvalLine(8, nil)
    return result

proc newEvalUnit*(cont: ptr BindMap[ptr Token]):ptr EvalUnit=
    result = cast[ptr EvalUnit](alloc0(sizeof(EvalUnit)))
    result.mainCtx = newBindMap[ptr Token](16)
    result.mainCtx.father = cont
    result.nowLine = newEvalLine(8, nil)
    return result

proc freeEvalUnit*(u: ptr EvalUnit)=
    freeBindMap(u.mainCtx)
    freeEvalLine(u.nowLine)
    dealloc(u)

proc eval*(u: var ptr EvalUnit, inp: ptr List[ptr Token]):ptr Token=
    # echo("start new unit")
    # for i in 0..len(inp)-1:
    #     print inp[i]
    result = newToken(TypeEnum.string)
    if(high(inp) < 0):
        result.val.string = ""
        return result
    var temp = newToken()

    var idx = 0

    while idx <= high(inp):
        # echo idx
        var nowToken = inp[idx]
        # print nowToken
        var nextToken: ptr Token
        if idx < high(inp):
            # print inp[idx+1]
            nextToken = getFinalToken(inp[idx+1], u.mainCtx)
        if not isNil(nextToken) and nextToken.tp == TypeEnum.op:
            if isNil(u.nowLine.line[0]) or (u.nowLine.line[0].tp != TypeEnum.op):
                var newLine = newEvalLine(3, u.nowLine)
                newLine.idx = 2
                newLine.line[0] = nextToken
                newLine.line[1] = getFinalToken(nowToken, u.mainCtx)
                u.nowLine = newLine
                idx += 1
            else:
                var newLine = newEvalLine(3, u.nowLine.father)
                newLine.idx = 1
                newLine.line[0] = nextToken
                u.nowLine.father = newLine
                u.nowLine.line[u.nowLine.idx] = getFinalToken(nowToken, u.mainCtx)
                u.nowLine.idx += 1
                idx += 1
        else:
            if nowToken.tp == TypeEnum.word:
                nowToken = getFinalToken(nowToken, u.mainCtx)
            if nowToken.tp == TypeEnum.op:
                if isNil(u.nowLine.line[0]):
                    result.tp = TypeEnum.err
                    result.val.string = "Illegal grammar!!!"
                if u.nowLine.idx > 0:
                    u.nowLine.idx -= 1
                var newLine = newEvalLine(3, u.nowLine)
                newLine.idx = 2
                newLine.line[0] = nowToken
                newLine.line[1] = u.nowLine.line[u.nowLine.idx]
                u.nowLine = newLine
            elif nowToken.tp < TypeEnum.set_word:
                u.nowLine.line[u.nowLine.idx] = getFinalToken(nowToken, u.mainCtx)
                if not isNil(u.nowLine.father):
                    u.nowLine.idx += 1
            else:
                var newLine = newEvalLine(nowToken.explen.int, u.nowLine)
                newLine.idx = 1
                newLine.line[0] = nowToken
                u.nowLine = newLine
            
        while not isNil(u.nowLine.line[0]) and (u.nowLine.idx == u.nowLine.line[0].explen):
            
            # for i in 0..u.nowLine.idx-1:
            #     write(stdout, $u.nowLine.line[i].toStr & " ")
            # write(stdout, "\n")
            # flushFile(stdout)
            # var s = readLine(stdin)

            temp = u.nowLine.eval(u.mainCtx)
            if not isNil(temp) and temp.tp == TypeEnum.err:
                result.val.string = $temp.val.string & "\n-->Near: "
                if idx >= 3:
                    for i in countdown(2, 0):
                        result.val.string = result.val.string & inp[idx-i].toStr & " "
                else:
                    result.val.string = result.val.string & inp[idx].toStr
            if not isNil(u.nowLine.father):
                var oldLine = u.nowLine
                u.nowLine = u.nowLine.father
                freeEvalLine(oldLine)
                u.nowLine.line[u.nowLine.idx] = temp
                if not isNil(u.nowLine.father):
                    u.nowLine.idx += 1
            else:
                break
        idx += 1
    if isNil(u.nowLine.line[0]) or isNil(temp):
        result.val.string = ""
    elif isNil(u.nowLine.father):
        result = u.nowLine.line[0]
    else:
        u.nowLine = newEvalLine(10, nil)
        u.nowLine.idx = 0
        result.tp = TypeEnum.err
        result.val.string = "Incomplete expression \n-> Near: "
        if idx >= 3:
            for i in countdown(3, 1):
                result.val.string = result.val.string & inp[idx-i].toStr & " "
        else:
            result.val.string = result.val.string & inp[idx-1].toStr
    return result
        

proc eval*(u: var ptr EvalUnit, s: string):ptr Token=
    return u.eval(toTokens(s))
            


when isMainModule:
    var unit = newEvalUnit()

    print unit.eval("i: 123 j: 456")
    # echo(isNil(unit.mainCtx.map[cstring("i")]))
    echo(repr(unit.mainCtx))


