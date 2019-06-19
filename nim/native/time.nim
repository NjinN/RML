import times

proc getCpuTime*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    result = newToken(TypeEnum.decimal)
    result.val.decimal = cpuTime()
    return result

proc gmt*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    result = newToken(TypeEnum.decimal)
    result.val.decimal = epochTime()
    return result

proc cost*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    if args[1].tp == TypeEnum.list:
        var t = cputime()
        var unit = newEvalUnit(cont)
        discard unit.eval(args[1].val.list)
        result = newToken(TypeEnum.decimal)
        result.val.decimal = cpuTime() - t
        freeEvalUnit(unit)
    else:
        result = newToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result
