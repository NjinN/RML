proc iff*(args: var seq[ref Token], cont: ref Context = nil):ref Token=
    if args[2].tp == TypeEnum.list:
        case args[1].tp
        of TypeEnum.logic:
            if args[1].val.logic:
                var unit = newEvalUnit(cont)
                result = unit.eval(args[2].val.list)
                return result
            else:
                return nil
        of TypeEnum.integer:
            if args[1].val.integer != 0:
                var unit = newEvalUnit(cont)
                result = unit.eval(args[2].val.list)
                return result
            else:
                return nil
        of TypeEnum.decimal:
            if args[1].val.decimal != 0.0:
                var unit = newEvalUnit(cont)
                result = unit.eval(args[2].val.list)
                return result
            else:
                return nil
        of TypeEnum.string:
            if args[1].val.string != cstring(""):
                var unit = newEvalUnit(cont)
                result = unit.eval(args[2].val.list)
                return result
            else:
                return nil
        of TypeEnum.none:
            return nil
        else:
            var unit = newEvalUnit(cont)
            result = unit.eval(args[2].val.list)
            return result
    else:
        result = newToken(TypeEnum.err, 1)
        result.val.string = "Type Mismatch"
        return result


proc either*(args: var seq[ref Token], cont: ref Context = nil):ref Token=
    if args[2].tp == TypeEnum.list and args[3].tp == TypeEnum.list:
        var unit = newEvalUnit(cont)
        case args[1].tp
        of TypeEnum.logic:
            if args[1].val.logic:
                result = unit.eval(args[2].val.list)
            else:
                result = unit.eval(args[3].val.list)
        of TypeEnum.integer:
            if args[1].val.integer != 0:
                result = unit.eval(args[2].val.list)
            else:
                result = unit.eval(args[3].val.list)
        of TypeEnum.decimal:
            if args[1].val.decimal != 0.0:
                result = unit.eval(args[2].val.list)
            else:
                result = unit.eval(args[3].val.list)
        of TypeEnum.string:
            if args[1].val.string != cstring(""):
                result = unit.eval(args[2].val.list)
            else:
                result = unit.eval(args[3].val.list)
        of TypeEnum.none:
            result = unit.eval(args[3].val.list)
        else:
            result = unit.eval(args[2].val.list)
    else:
        result = newToken(TypeEnum.err, 1)
        result.val.string = "Type Mismatch"
    return result