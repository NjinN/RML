import ../token
export token

import times

proc plus*(args: var seq[ref Token]):ref Token=
    result = new Token
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.integer
            result.val.integer = args[1].val.integer + args[2].val.integer
            result.explen = 1
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = float(args[1].val.integer) + args[2].val.float
            result.explen = 1
            return result
    elif args[1].tp == TypeEnum.float:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float + float(args[2].val.integer)
            result.explen = 1
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float + args[2].val.float
            result.explen = 1
            return result
    result.tp = TypeEnum.err
    result.val.string = cstring("Type Mismatch")
    result.explen = 1


proc minus*(args: var seq[ref Token]):ref Token=
    result = new Token
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.integer
            result.val.integer = args[1].val.integer - args[2].val.integer
            result.explen = 1
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = float(args[1].val.integer) - args[2].val.float
            result.explen = 1
            return result
    elif args[1].tp == TypeEnum.float:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float - float(args[2].val.integer)
            result.explen = 1
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float - args[2].val.float
            result.explen = 1
            return result
    result.tp = TypeEnum.err
    result.val.string = cstring("Type Mismatch")
    result.explen = 1
    
    
proc multiply*(args: var seq[ref Token]):ref Token=
    result = new Token
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.integer
            result.val.integer = args[1].val.integer * args[2].val.integer
            result.explen = 1
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = float(args[1].val.integer) * args[2].val.float
            result.explen = 1
            return result
    elif args[1].tp == TypeEnum.float:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float * float(args[2].val.integer)
            result.explen = 1
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float * args[2].val.float
            result.explen = 1
            return result
    result.tp = TypeEnum.err
    result.val.string = cstring("Type Mismatch")
    result.explen = 1


proc divide*(args: var seq[ref Token]):ref Token=
    result = new Token
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.integer / args[2].val.integer
            result.explen = 1
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = float(args[1].val.integer) / args[2].val.float
            result.explen = 1
            return result
    elif args[1].tp == TypeEnum.float:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float / float(args[2].val.integer)
            result.explen = 1
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float / args[2].val.float
            result.explen = 1
            return result
    result.tp = TypeEnum.err
    result.val.string = cstring("Type Mismatch")
    result.explen = 1

proc getCpuTime*(args: var seq[ref Token]):ref Token=
    result = new Token
    result.tp = TypeEnum.float
    result.val.float = cpuTime()
    result.explen = 1
    return result

proc gmt*(args: var seq[ref Token]):ref Token=
    result = new Token
    result.tp = TypeEnum.float
    result.val.float = epochTime()
    result.explen = 1
    return result

proc print*(args: var seq[ref Token]):ref Token=
    echo(outputStr(args[1]))
    return nil