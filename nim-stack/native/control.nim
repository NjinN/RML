proc iff*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    if args[2].tp == TypeEnum.list:
        case args[1].tp
        of TypeEnum.logic:
            if args[1].val.logic:
                result = stack.eval(args[2].val.list, ctx)
                return result
            else:
                return initToken(TypeEnum.null)
        of TypeEnum.integer:
            if args[1].val.integer != 0:
                result = stack.eval(args[2].val.list, ctx)
                return result
            else:
                return initToken(TypeEnum.null)
        of TypeEnum.decimal:
            if args[1].val.decimal != 0.0:
                result = stack.eval(args[2].val.list, ctx)
                return result
            else:
                return initToken(TypeEnum.null)
        of TypeEnum.string:
            if args[1].val.string != "":
                result = stack.eval(args[2].val.list, ctx)
                return result
            else:
                return initToken(TypeEnum.null)
        of TypeEnum.none:
            return initToken(TypeEnum.null)
        else:
            result = stack.eval(args[2].val.list, ctx)
            return result
    else:
        result = initToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
        return result


proc either*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    if args[2].tp == TypeEnum.list and args[3].tp == TypeEnum.list:
        case args[1].tp
        of TypeEnum.logic:
            if args[1].val.logic:
                result = stack.eval(args[2].val.list, ctx)
            else:
                result = stack.eval(args[3].val.list, ctx)
        of TypeEnum.integer:
            if args[1].val.integer != 0:
                result = stack.eval(args[2].val.list, ctx)
            else:
                result = stack.eval(args[3].val.list, ctx)
        of TypeEnum.decimal:
            if args[1].val.decimal != 0.0:
                result = stack.eval(args[2].val.list, ctx)
            else:
                result = stack.eval(args[3].val.list, ctx)
        of TypeEnum.string:
            if args[1].val.string != "":
                result = stack.eval(args[2].val.list, ctx)
            else:
                result = stack.eval(args[3].val.list, ctx)
        of TypeEnum.none:
            result = stack.eval(args[3].val.list, ctx)
        else:
            result = stack.eval(args[2].val.list, ctx)
    else:
        result = initToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result


proc loop*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    if args[1].tp == TypeEnum.integer and args[2].tp == TypeEnum.list:
        # print args[1]
        # var s = readLine(stdin)
        for i in 1..args[1].val.integer:
            try:
                result = stack.eval(args[2].val.list, ctx)
            except:
                if getCurrentExceptionMsg() == "continue":
                    continue
                elif getCurrentExceptionMsg() == "break":
                    break
                else:
                    raise getCurrentException()

    else:
        result = initToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result

proc repeat*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    if args[1].tp == TypeEnum.word and args[2].tp == TypeEnum.integer and args[3].tp == TypeEnum.list:
        var c = initBindMap[Token](4)
        c.father = addr(ctx)
        var countToken = initToken(TypeEnum.integer)
        
        countToken.val.integer = 1
        c[args[1].val.string] = countToken 
        while c[args[1].val.string].val.integer <= args[2].val.integer:
            try:
                result = stack.eval(args[3].val.list, c)
            except:
                if getCurrentExceptionMsg() == "continue":
                    continue
                elif getCurrentExceptionMsg() == "break":
                    break
                else:
                    raise getCurrentException()
            finally:
                # echo(repr(unit.mainCtx[args[1].val.string]))
                var temp = c[args[1].val.string]
                temp.val.integer += 1
                c[args[1].val.string] = temp
        freeBindMap(c)
    else:
        result = initToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result


proc ffor*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    if args[1].tp == TypeEnum.word and args[5].tp == TypeEnum.list and (args[2].tp == TypeEnum.integer or args[2].tp == TypeEnum.decimal) and (args[3].tp == TypeEnum.integer or args[3].tp == TypeEnum.decimal) and (args[4].tp == TypeEnum.integer or args[4].tp == TypeEnum.decimal) :
        var c = initBindMap[Token](4)
        c.father = addr(ctx)
        if(args[2].tp == TypeEnum.integer and args[3].tp == TypeEnum.integer and args[4].tp == TypeEnum.integer):
            var count = initToken(TypeEnum.integer)
            
            count.val.integer = args[2].val.integer
            c[args[1].val.string] = count
            while c[args[1].val.string].val.integer <= args[3].val.integer:
                try:
                    result = stack.eval(args[5].val.list, c)
                except:
                    if getCurrentExceptionMsg() == "continue":
                        continue
                    elif getCurrentExceptionMsg() == "break":
                        break
                    else:
                        raise getCurrentException()
                finally:
                    var temp = c[args[1].val.string]
                    temp.val.integer += args[4].val.integer
                    c[args[1].val.string] = temp
            
        else:
            var count = initToken(TypeEnum.decimal)
            
            if args[2].tp == TypeEnum.integer:
                count.val.decimal = float64(args[2].val.integer)
            else:
                count.val.decimal = args[2].val.decimal
            c[args[1].val.string] = count
            
            var temp: Token
            if args[3].tp == TypeEnum.integer:
                while c[args[1].val.string].val.decimal <= args[3].val.integer.float64:
                    try:
                        result = stack.eval(args[5].val.list, c)
                    except:
                        if getCurrentExceptionMsg() == "continue":
                            continue
                        elif getCurrentExceptionMsg() == "break":
                            break
                        else:
                            raise getCurrentException()
                    finally:
                        if args[4].tp == TypeEnum.integer:
                            temp = c[args[1].val.string]
                            temp.val.decimal += args[4].val.integer.float64
                            c[args[1].val.string] = temp
                        else:
                            temp = c[args[1].val.string]
                            temp.val.decimal += args[4].val.decimal
                            c[args[1].val.string] = temp
            else:
                while c[args[1].val.string].val.decimal <= args[3].val.decimal:
                    try:
                        result = stack.eval(args[5].val.list, ctx)
                    except:
                        if getCurrentExceptionMsg() == "continue":
                            continue
                        elif getCurrentExceptionMsg() == "break":
                            break
                        else:
                            raise getCurrentException()
                    finally:
                        if args[4].tp == TypeEnum.integer:
                            temp = c[args[1].val.string]
                            temp.val.decimal += args[4].val.integer.float64
                            c[args[1].val.string] = temp
                        else:
                            temp = c[args[1].val.string]
                            temp.val.decimal += args[4].val.decimal
                            c[args[1].val.string] = temp
            
        freeBindMap(c)
    else:
        result = initToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result


proc wwhile*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    var args = addr(stack.line[stack.startPos.last])
    if args[1].tp == TypeEnum.list and args[2].tp == TypeEnum.list:
        var condCtx = initBindMap[Token](4)
        condCtx.father = addr(ctx)
        var bodyCtx = initBindMap[Token](4)
        bodyCtx.father = addr(ctx)
        var b = stack.eval(args[1].val.list, condCtx)
        
        while b.tp == TypeEnum.logic and b.val.logic:
            try:
                result = stack.eval(args[2].val.list, bodyCtx)
            except:
                if getCurrentExceptionMsg() == "continue":
                    continue
                elif getCurrentExceptionMsg() == "break":
                    break
                else:
                    raise getCurrentException()
            finally:
                b = stack.eval(args[1].val.list, condCtx)

                if not (b.tp == TypeEnum.logic):
                    result = initToken(TypeEnum.err)
                    result.val.string = "Bad Logic Expression "
                    return result
        freeBindMap(condCtx)
        freeBindMap(bodyCtx)
    else:
        result = initToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result

proc bbreak*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    raise newException(CatchableError, "break")

proc ccontinue*(stack: var EvalStack, ctx: var BindMap[Token]):Token=
    raise newException(CatchableError, "continue")