import token
import totoken

proc initEvalStack*(stack: var EvalStack)=
    stack.idx = 0
    stack.startPos = initList[int](8)
    stack.endPos = initList[int](8)

proc clearEvalStack*(stack: var EvalStack)=
    stack.idx = 0
    stack.startPos.clear(0)
    stack.endPos.clear(0)

proc push*(stack: var EvalStack, t: Token)=
    stack.line[stack.idx] = t
    stack.idx += 1

proc eval*(stack: var EvalStack, inp: ptr List[Token], ctx: var BindMap[Token]):Token

proc getVal(t: Token, ctx: var BindMap[Token], stack: var EvalStack):Token=
    case t.tp
    of TypeEnum.word:
        result = ctx[t.toStr]
    of TypeEnum.lit_word:
        result.tp = TypeEnum.word
        result.val.string = t.val.string
    of TypeEnum.paren:
        result = stack.eval(t.val.list, ctx)
    else:
        result = t

proc run*(f: ptr Func, stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var c = initBindMap[Token]()
    c.father = addr(ctx)
    for i in 0..high(f.args):
        c[f.args[i].toStr] = stack.line[stack.startPos.last + i + 1]
    result = stack.eval(f.body, c)
    freeBindMap(c)


proc evalExp*(stack: var EvalStack, ctx: var BindMap[Token])=
    var temp: Token
    # var occ1 = getOccupiedMem()
    try:
        case stack.line[stack.startPos.last].tp
        of TypeEnum.set_word:
            ctx[stack.line[stack.startPos.last].val.string] = stack.line[stack.endPos.last]
            temp = stack.line[stack.endPos.last]
        of TypeEnum.native:
            temp = stack.line[stack.startPos.last].val.exec.run(stack, ctx)
        of TypeEnum.op:
            
            temp = stack.line[stack.startPos.last].val.exec.run(stack, ctx)
            
        of TypeEnum.function:
            temp = stack.line[stack.startPos.last].val.fc.run(stack, ctx)
        else:
            temp = temp
    except CatchableError:
        if getCurrentExceptionMsg() == "break":
            raise getCurrentException()
        elif getCurrentExceptionMsg() == "continue":
            raise getCurrentException()
    except:
        temp.tp = TypeEnum.err
        temp.val.string = "Illegal grammar!!!"
    finally:
        stack.line[stack.startPos.last] = temp
        stack.idx = stack.startPos.last + 1
        stack.startPos.pop()
        stack.endPos.pop()
        # var occ2 = getOccupiedMem()
        # echo (occ1 - occ2)  
        


proc eval*(stack: var EvalStack, inp: ptr List[Token], ctx: var BindMap[Token]):Token=
    result = initToken(TypeEnum.null)
    if high(inp) < 0:
        return result
    
    # for item in inp.each:
    #     print item

    var startIdx = stack.idx
    var startDeep = high(stack.endPos)

    var idx = 0
    while idx <= high(inp):
        var nowToken = inp[idx]
        var nextToken: Token
        if idx < high(inp):
            nextToken = getVal(inp[idx + 1], ctx, stack)

        if nextToken.tp == TypeEnum.op and (startDeep < 0 or stack.idx > stack.endPos[startDeep]) :
            if high(stack.startPos) < 0 or stack.line[stack.startPos.last].tp != TypeEnum.op:
                stack.startPos.add(stack.idx)
                stack.push(nextToken)
                stack.push(getVal(nowToken, ctx, stack))
                stack.endPos.add(stack.idx)
            elif high(stack.startPos) < 0 or stack.line[stack.startPos.last].tp == TypeEnum.op:
                stack.push(getVal(nowToken, ctx, stack))
                stack.evalExp(ctx)
                stack.push(stack.line[stack.idx - 1])
                stack.line[stack.idx - 2] = nextToken
                stack.startPos.add(stack.idx - 2)
                stack.endPos.add(stack.idx)
            idx += 1
        else:
            nowToken = getVal(nowToken, ctx, stack)
            if nowToken.tp == TypeEnum.err:
                return nowToken
            elif nowToken.tp == TypeEnum.op:
                stack.startPos.add(stack.idx - 1)
                stack.push(stack.line[stack.idx - 1])
                stack.line[stack.idx - 2] = nowToken
                stack.endPos.add(stack.idx)
            elif nowToken.tp < TypeEnum.set_word:
                stack.push(nowToken)
            else:
                stack.startPos.add(stack.idx)
                stack.endPos.add(stack.idx + nowToken.explen - 1)
                stack.push(nowToken)

        while high(stack.endPos) > startDeep and stack.idx == stack.endPos.last + 1:
            stack.evalExp(ctx)

        idx += 1
                
    result = stack.line[stack.idx - 1]
    stack.idx = startIdx





proc eval*(stack: var EvalStack, inpStr: string, ctx: var BindMap[Token]):Token=
    result = stack.eval(toTokens(inpStr), ctx)