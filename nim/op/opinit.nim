import ../native/nativelib

proc initOp*(cont: ref Context)=
    var addToken = newToken(TypeEnum.op, 3)
    addToken.val.exec = newExec("+", nativelib.plus)
    cont.map[cstring("+")] = addToken

    var subToken = newToken(TypeEnum.op, 3)
    subToken.val.exec = newExec("-", nativelib.minus)
    cont.map[cstring("-")] = subToken

    var mulToken = newToken(TypeEnum.op, 3)
    mulToken.val.exec = newExec("*", nativelib.multiply)
    cont.map[cstring("*")] = mulToken

    var divToken = newToken(TypeEnum.op, 3)
    divToken.val.exec = newExec("/", nativelib.divide)
    cont.map[cstring("/")] = divToken

    var eqToken = newToken(TypeEnum.op, 3)
    eqToken.val.exec = newExec("=", nativelib.eq)
    cont.map[cstring("=")] = eqToken

    var neToken = newToken(TypeEnum.op, 3)
    neToken.val.exec = newExec("<>", nativelib.ne)
    cont.map[cstring("<>")] = neToken

    var ltToken = newToken(TypeEnum.op, 3)
    ltToken.val.exec = newExec("=", nativelib.lt)
    cont.map[cstring("<")] = ltToken

    var gtToken = newToken(TypeEnum.op, 3)
    gtToken.val.exec = newExec(">", nativelib.gt)
    cont.map[cstring(">")] = gtToken

    var lteqToken = newToken(TypeEnum.op, 3)
    lteqToken.val.exec = newExec("<=", nativelib.lteq)
    cont.map[cstring("<=")] = lteqToken

    var gteqToken = newToken(TypeEnum.op, 3)
    gteqToken.val.exec = newExec(">=", nativelib.gteq)
    cont.map[cstring(">=")] = gteqToken

