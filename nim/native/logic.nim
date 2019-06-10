
proc eq*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.logic, 1)
    result.val.logic = false
    case args[1].tp
    of TypeEnum.none:
        if args[2].tp == TypeEnum.none:
            result.val.logic = true
    of TypeEnum.logic:
        if args[2].tp == TypeEnum.logic:
            result.val.logic = (args[1].val.logic == args[2].val.logic)
    of TypeEnum.integer:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer == args[2].val.integer)
        of TypeEnum.float:
            result.val.logic = (args[1].val.integer.float == args[2].val.float)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer == ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.float:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float == args[2].val.float)
        of TypeEnum.float:
            result.val.logic = (args[1].val.float == args[2].val.float)
        else:
            discard 0
    of TypeEnum.char:
        case args[2].tp
        of TypeEnum.char:
            result.val.logic = (args[1].val.char == args[2].val.char)
        of TypeEnum.integer:
            result.val.logic = (ord(args[1].val.char) == args[2].val.integer)
        else:
            discard 0
    of TypeEnum.string:
        if args[2].tp == TypeEnum.string:
            result.val.logic = (args[1].val.string == args[2].val.string)
    of TypeEnum.word:
        if args[2].tp == TypeEnum.word:
            result.val.logic = (args[1].val.string == args[2].val.string)
    else: 
        discard 0


proc ne*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.logic, 1)
    result.val.logic = true
    case args[1].tp
    of TypeEnum.none:
        if args[2].tp == TypeEnum.none:
            result.val.logic = false
    of TypeEnum.logic:
        if args[2].tp == TypeEnum.logic:
            result.val.logic = (args[1].val.logic != args[2].val.logic)
    of TypeEnum.integer:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer != args[2].val.integer)
        of TypeEnum.float:
            result.val.logic = (args[1].val.integer.float != args[2].val.float)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer != ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.float:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float != args[2].val.float)
        of TypeEnum.float:
            result.val.logic = (args[1].val.float != args[2].val.float)
        else:
            discard 0
    of TypeEnum.char:
        case args[2].tp
        of TypeEnum.char:
            result.val.logic = (args[1].val.char != args[2].val.char)
        of TypeEnum.integer:
            result.val.logic = (ord(args[1].val.char) != args[2].val.integer)
        else:
            discard 0
    of TypeEnum.string:
        if args[2].tp == TypeEnum.string:
            result.val.logic = (args[1].val.string != args[2].val.string)
    of TypeEnum.word:
        if args[2].tp == TypeEnum.word:
            result.val.logic = (args[1].val.string != args[2].val.string)
    else: 
        discard 0



proc lt*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.logic, 1)
    result.val.logic = false
    case args[1].tp
    of TypeEnum.none:
        if args[2].tp == TypeEnum.none:
            result.val.logic = false
    of TypeEnum.logic:
        if args[2].tp == TypeEnum.logic:
            result.val.logic = (args[1].val.logic < args[2].val.logic)
    of TypeEnum.integer:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer < args[2].val.integer)
        of TypeEnum.float:
            result.val.logic = (args[1].val.integer.float < args[2].val.float)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer < ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.float:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float < args[2].val.float)
        of TypeEnum.float:
            result.val.logic = (args[1].val.float < args[2].val.float)
        else:
            discard 0
    of TypeEnum.char:
        case args[2].tp
        of TypeEnum.char:
            result.val.logic = (args[1].val.char < args[2].val.char)
        of TypeEnum.integer:
            result.val.logic = (ord(args[1].val.char) < args[2].val.integer)
        else:
            discard 0
    of TypeEnum.string:
        if args[2].tp == TypeEnum.string:
            result.val.logic = (args[1].val.string < args[2].val.string)
    of TypeEnum.word:
        if args[2].tp == TypeEnum.word:
            result.val.logic = (args[1].val.string < args[2].val.string)
    else: 
        discard 0


proc gt*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.logic, 1)
    result.val.logic = false
    case args[1].tp
    of TypeEnum.none:
        if args[2].tp == TypeEnum.none:
            result.val.logic = false
    of TypeEnum.logic:
        if args[2].tp == TypeEnum.logic:
            result.val.logic = (args[1].val.logic > args[2].val.logic)
    of TypeEnum.integer:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer > args[2].val.integer)
        of TypeEnum.float:
            result.val.logic = (args[1].val.integer.float > args[2].val.float)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer > ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.float:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float > args[2].val.float)
        of TypeEnum.float:
            result.val.logic = (args[1].val.float > args[2].val.float)
        else:
            discard 0
    of TypeEnum.char:
        case args[2].tp
        of TypeEnum.char:
            result.val.logic = (args[1].val.char > args[2].val.char)
        of TypeEnum.integer:
            result.val.logic = (ord(args[1].val.char) > args[2].val.integer)
        else:
            discard 0
    of TypeEnum.string:
        if args[2].tp == TypeEnum.string:
            result.val.logic = (args[1].val.string > args[2].val.string)
    of TypeEnum.word:
        if args[2].tp == TypeEnum.word:
            result.val.logic = (args[1].val.string > args[2].val.string)
    else: 
        discard 0


proc lteq*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.logic, 1)
    result.val.logic = false
    case args[1].tp
    of TypeEnum.none:
        if args[2].tp == TypeEnum.none:
            result.val.logic = true
    of TypeEnum.logic:
        if args[2].tp == TypeEnum.logic:
            result.val.logic = (args[1].val.logic <= args[2].val.logic)
    of TypeEnum.integer:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer <= args[2].val.integer)
        of TypeEnum.float:
            result.val.logic = (args[1].val.integer.float <= args[2].val.float)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer <= ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.float:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float <= args[2].val.float)
        of TypeEnum.float:
            result.val.logic = (args[1].val.float <= args[2].val.float)
        else:
            discard 0
    of TypeEnum.char:
        case args[2].tp
        of TypeEnum.char:
            result.val.logic = (args[1].val.char <= args[2].val.char)
        of TypeEnum.integer:
            result.val.logic = (ord(args[1].val.char) <= args[2].val.integer)
        else:
            discard 0
    of TypeEnum.string:
        if args[2].tp == TypeEnum.string:
            result.val.logic = (args[1].val.string <= args[2].val.string)
    of TypeEnum.word:
        if args[2].tp == TypeEnum.word:
            result.val.logic = (args[1].val.string <= args[2].val.string)
    else: 
        discard 0


proc gteq*(args: var seq[ref Token]):ref Token=
    result = newToken(TypeEnum.logic, 1)
    result.val.logic = false
    case args[1].tp
    of TypeEnum.none:
        if args[2].tp == TypeEnum.none:
            result.val.logic = true
    of TypeEnum.logic:
        if args[2].tp == TypeEnum.logic:
            result.val.logic = (args[1].val.logic >= args[2].val.logic)
    of TypeEnum.integer:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer >= args[2].val.integer)
        of TypeEnum.float:
            result.val.logic = (args[1].val.integer.float >= args[2].val.float)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer >= ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.float:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float >= args[2].val.float)
        of TypeEnum.float:
            result.val.logic = (args[1].val.float >= args[2].val.float)
        else:
            discard 0
    of TypeEnum.char:
        case args[2].tp
        of TypeEnum.char:
            result.val.logic = (args[1].val.char >= args[2].val.char)
        of TypeEnum.integer:
            result.val.logic = (ord(args[1].val.char) >= args[2].val.integer)
        else:
            discard 0
    of TypeEnum.string:
        if args[2].tp == TypeEnum.string:
            result.val.logic = (args[1].val.string >= args[2].val.string)
    of TypeEnum.word:
        if args[2].tp == TypeEnum.word:
            result.val.logic = (args[1].val.string >= args[2].val.string)
    else: 
        discard 0