import ../token
export token

include "logic.nim"
include "time.nim"

proc plus*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.err, 1)
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.integer
            result.val.integer = args[1].val.integer + args[2].val.integer
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = float(args[1].val.integer) + args[2].val.float
            return result
    elif args[1].tp == TypeEnum.float:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float + float(args[2].val.integer)
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float + args[2].val.float
            return result
    result.val.string = cstring("Error: Type Mismatch")


proc minus*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.err, 1)
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.integer
            result.val.integer = args[1].val.integer - args[2].val.integer
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = float(args[1].val.integer) - args[2].val.float
            return result
    elif args[1].tp == TypeEnum.float:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float - float(args[2].val.integer)
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float - args[2].val.float
            return result
    result.val.string = cstring("Error: Type Mismatch")
    
    
proc multiply*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.err, 1)
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.integer
            result.val.integer = args[1].val.integer * args[2].val.integer
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = float(args[1].val.integer) * args[2].val.float
            return result
    elif args[1].tp == TypeEnum.float:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float * float(args[2].val.integer)
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float * args[2].val.float
            return result
    result.val.string = cstring("Error: Type Mismatch")


proc divide*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.err, 1)
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.integer / args[2].val.integer
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = float(args[1].val.integer) / args[2].val.float
            return result
    elif args[1].tp == TypeEnum.float:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float / float(args[2].val.integer)
            return result
        elif args[2].tp == TypeEnum.float:
            result.tp = TypeEnum.float
            result.val.float = args[1].val.float / args[2].val.float
            return result
    result.val.string = cstring("Error: Type Mismatch")


proc print*(args: var seq[ref Token]):ref Token=
    echo(outputStr(args[1]))
    return nil