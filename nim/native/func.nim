proc fc*(args: var seq[ref Token], cont: ref Context = nil):ref Token=
    if args[1].tp != TypeEnum.list or args[2].tp != TypeEnum.list:
        result = newToken(TypeEnum.err, 1)
        result.val.string = cstring("Type Mismatch")
        return result
    for idx in 0..len(args[1].val.list)-1:
        if args[1].val.list[idx].tp != TypeEnum.word:
            result = newToken(TypeEnum.err, 1)
            result.val.string = cstring("Type Mismatch")
            return result
    
    result = newToken(TypeEnum.function, len(args[1].val.list) + 1)
    result.val.fc = newFunc(args[1].val.list, args[2].val.list)
    return result


    


