proc fc*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    if args[1].tp != TypeEnum.list or args[2].tp != TypeEnum.list:
        result = initToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
        return result
    for idx in 0..high(args[1].val.list):
        if args[1].val.list[idx].tp != TypeEnum.word:
            result = initToken(TypeEnum.err)
            result.val.string = "Type Mismatch"
            return result
    
    result = initToken(TypeEnum.function)
    result.val.fc = newFunc(args[1].val.list, args[2].val.list)
    return result


    


