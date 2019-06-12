import times

proc getCpuTime*(args: var seq[ref Token], cont: ref Context = nil):ref Token=
    result = newToken(TypeEnum.decimal, 1)
    result.val.decimal = cpuTime()
    return result

proc gmt*(args: var seq[ref Token], cont: ref Context = nil):ref Token=
    result = newToken(TypeEnum.decimal, 1)
    result.val.decimal = epochTime()
    return result

