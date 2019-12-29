import times

proc getCpuTime*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    result = initToken(TypeEnum.decimal)
    result.val.decimal = cpuTime()
    return result

proc gmt*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    result = initToken(TypeEnum.decimal)
    result.val.decimal = epochTime()
    return result

proc cost*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    if args[1].tp == TypeEnum.list:
        var t = cputime()
        discard stack.eval(args[1].val.list, ctx)
        result = initToken(TypeEnum.decimal)
        result.val.decimal = cpuTime() - t
    else:
        result = initToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result
