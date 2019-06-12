import ../token
export token
import ../evalunit

include "logic.nim"
include "time.nim"
include "control.nim"
include "func.nim"

proc plus*(args: var seq[ref Token], cont: ref Context = nil):ref Token=
    result = newToken(TypeEnum.err, 1)
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.integer
            result.val.integer = args[1].val.integer + args[2].val.integer
            return result
        elif args[2].tp == TypeEnum.decimal:
            result.tp = TypeEnum.decimal
            result.val.decimal = float(args[1].val.integer) + args[2].val.decimal
            return result
    elif args[1].tp == TypeEnum.decimal:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.decimal
            result.val.decimal = args[1].val.decimal + float(args[2].val.integer)
            return result
        elif args[2].tp == TypeEnum.decimal:
            result.tp = TypeEnum.decimal
            result.val.decimal = args[1].val.decimal + args[2].val.decimal
            return result
    result.val.string = cstring("Type Mismatch")


proc minus*(args: var seq[ref Token], cont: ref Context = nil):ref Token=
    result = newToken(TypeEnum.err, 1)
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.integer
            result.val.integer = args[1].val.integer - args[2].val.integer
            return result
        elif args[2].tp == TypeEnum.decimal:
            result.tp = TypeEnum.decimal
            result.val.decimal = float(args[1].val.integer) - args[2].val.decimal
            return result
    elif args[1].tp == TypeEnum.decimal:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.decimal
            result.val.decimal = args[1].val.decimal - float(args[2].val.integer)
            return result
        elif args[2].tp == TypeEnum.decimal:
            result.tp = TypeEnum.decimal
            result.val.decimal = args[1].val.decimal - args[2].val.decimal
            return result
    result.val.string = cstring("Type Mismatch")
    
    
proc multiply*(args: var seq[ref Token], cont: ref Context = nil):ref Token=
    result = newToken(TypeEnum.err, 1)
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.integer
            result.val.integer = args[1].val.integer * args[2].val.integer
            return result
        elif args[2].tp == TypeEnum.decimal:
            result.tp = TypeEnum.decimal
            result.val.decimal = float(args[1].val.integer) * args[2].val.decimal
            return result
    elif args[1].tp == TypeEnum.decimal:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.decimal
            result.val.decimal = args[1].val.decimal * float(args[2].val.integer)
            return result
        elif args[2].tp == TypeEnum.decimal:
            result.tp = TypeEnum.decimal
            result.val.decimal = args[1].val.decimal * args[2].val.decimal
            return result
    result.val.string = cstring("Type Mismatch")


proc divide*(args: var seq[ref Token], cont: ref Context = nil):ref Token=
    result = newToken(TypeEnum.err, 1)
    if args[1].tp == TypeEnum.integer:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.decimal
            result.val.decimal = args[1].val.integer / args[2].val.integer
            return result
        elif args[2].tp == TypeEnum.decimal:
            result.tp = TypeEnum.decimal
            result.val.decimal = float(args[1].val.integer) / args[2].val.decimal
            return result
    elif args[1].tp == TypeEnum.decimal:
        if args[2].tp == TypeEnum.integer:
            result.tp = TypeEnum.decimal
            result.val.decimal = args[1].val.decimal / float(args[2].val.integer)
            return result
        elif args[2].tp == TypeEnum.decimal:
            result.tp = TypeEnum.decimal
            result.val.decimal = args[1].val.decimal / args[2].val.decimal
            return result
    result.val.string = cstring("Type Mismatch")


proc print*(args: var seq[ref Token], cont: ref Context = nil):ref Token=
    if args[1].tp == TypeEnum.list:
        for i in 0..len(args[1].val.list)-1:
            write(stdout, $args[1].val.list[i].toStr)
        write(stdout, "\n")
        flushFile(stdout)
    else:
        echo(outputStr(args[1]))
    return nil