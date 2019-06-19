
proc eq*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    result = newToken(TypeEnum.logic)
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
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.integer.float == args[2].val.decimal)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer == ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.decimal:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float == args[2].val.decimal)
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.decimal == args[2].val.decimal)
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


proc ne*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    result = newToken(TypeEnum.logic)
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
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.integer.float != args[2].val.decimal)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer != ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.decimal:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float != args[2].val.decimal)
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.decimal != args[2].val.decimal)
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



proc lt*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    result = newToken(TypeEnum.logic)
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
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.integer.float < args[2].val.decimal)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer < ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.decimal:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float < args[2].val.decimal)
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.decimal < args[2].val.decimal)
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


proc gt*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    result = newToken(TypeEnum.logic)
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
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.integer.float > args[2].val.decimal)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer > ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.decimal:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float > args[2].val.decimal)
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.decimal > args[2].val.decimal)
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


proc lteq*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    result = newToken(TypeEnum.logic)
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
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.integer.float <= args[2].val.decimal)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer <= ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.decimal:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float <= args[2].val.decimal)
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.decimal <= args[2].val.decimal)
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


proc gteq*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    result = newToken(TypeEnum.logic)
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
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.integer.float >= args[2].val.decimal)
        of TypeEnum.char:
            result.val.logic = (args[1].val.integer >= ord(args[2].val.char))
        else:
            discard 0
    of TypeEnum.decimal:
        case args[2].tp
        of TypeEnum.integer:
            result.val.logic = (args[1].val.integer.float >= args[2].val.decimal)
        of TypeEnum.decimal:
            result.val.logic = (args[1].val.decimal >= args[2].val.decimal)
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