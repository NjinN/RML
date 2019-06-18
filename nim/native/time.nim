import times

proc getCpuTime*(args: var ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    result = newToken(TypeEnum.decimal)
    result.val.decimal = cpuTime()
    return result

proc gmt*(args: var ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    result = newToken(TypeEnum.decimal)
    result.val.decimal = epochTime()
    return result

