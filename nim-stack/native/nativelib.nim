import ../token
export token
import ../evalStack

include "logic.nim"
include "time.nim"
include "control.nim"
include "func.nim"

proc plus*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    result = initToken(TypeEnum.decimal)
 
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
    result.tp = TypeEnum.err
    result.val.string = "Type Mismatch"
   
    


proc minus*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    result = initToken(TypeEnum.decimal)
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
    result.tp = TypeEnum.err
    result.val.string = "Type Mismatch"
    
    
proc multiply*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    result = initToken(TypeEnum.err)
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


proc divide*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    result = initToken(TypeEnum.err)
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


proc print*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    if args[1].tp == TypeEnum.list:
        for i in 0..len(args[1].val.list)-1:
            write(stdout, args[1].val.list[i].toStr)
        write(stdout, "\n")
        flushFile(stdout)
    else:
        echo(outputStr(args[1]))
    return initToken(TypeEnum.null)

proc ttypeof*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    if not isNull(args[1]):
        result = initToken(TypeEnum.dataType)
        result.val.string = cstring($args[1].tp & "!")
    else:
        result = initToken(TypeEnum.null)
        result.val.string = "null" & "!"