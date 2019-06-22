import ../token
export token
import ../evalUnit

include "logic.nim"
include "time.nim"
include "control.nim"
include "func.nim"

proc plus*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil, unit: ptr EvalUnit = nil):ptr Token=
    result = newToken(TypeEnum.err)
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
    result.val.string = "Type Mismatch"


proc minus*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil, unit: ptr EvalUnit = nil):ptr Token=
    result = newToken(TypeEnum.err)
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
    result.val.string = "Type Mismatch"
    
    
proc multiply*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil, unit: ptr EvalUnit = nil):ptr Token=
    result = newToken(TypeEnum.err)
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
    result.val.string = "Type Mismatch"


proc divide*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil, unit: ptr EvalUnit = nil):ptr Token=
    result = newToken(TypeEnum.err)
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
    result.val.string = "Type Mismatch"


proc print*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil, unit: ptr EvalUnit = nil):ptr Token=
    if args[1].tp == TypeEnum.list:
        for i in 0..len(args[1].val.list)-1:
            write(stdout, args[1].val.list[i].toStr)
        write(stdout, "\n")
        flushFile(stdout)
    else:
        echo(outputStr(args[1]))
    return nil

proc ttypeof*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil, unit: ptr EvalUnit = nil):ptr Token=
    if not isNull(args[1]):
        result = newToken(TypeEnum.dataType)
        result.val.string = cstring($args[1].tp & "!")
    else:
        result = newToken(TypeEnum.null)
        result.val.string = "null" & "!"