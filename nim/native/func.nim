proc fc*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    if args[1].tp != TypeEnum.list or args[2].tp != TypeEnum.list:
        result = newToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
        return result
    for idx in 0..high(args[1].val.list):
        if args[1].val.list[idx].tp != TypeEnum.word:
            result = newToken(TypeEnum.err)
            result.val.string = "Type Mismatch"
            return result
    
    result = newToken(TypeEnum.function)
    result.val.fc = newFunc(args[1].val.list, args[2].val.list)
    return result


    


