import times

proc getCpuTime*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.float, 1)
    result.val.float = cpuTime()
    return result

proc gmt*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.float, 1)
    result.val.float = epochTime()
    return result

