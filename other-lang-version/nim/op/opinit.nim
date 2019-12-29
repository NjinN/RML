import ../native/nativeLib

proc initOp*(cont: ptr BindMap[ptr Token])=
    var addToken = newToken(TypeEnum.op)
    addToken.val.exec = newExec("+", nativelib.plus, 3)
    cont["+"] = addToken

    var subToken = newToken(TypeEnum.op)
    subToken.val.exec = newExec("-", nativelib.minus, 3)
    cont["-"] = subToken

    var mulToken = newToken(TypeEnum.op)
    mulToken.val.exec = newExec("*", nativelib.multiply, 3)
    cont["*"] = mulToken

    var divToken = newToken(TypeEnum.op)
    divToken.val.exec = newExec("/", nativelib.divide, 3)
    cont["/"] = divToken

    var eqToken = newToken(TypeEnum.op)
    eqToken.val.exec = newExec("=", nativelib.eq, 3)
    cont["="] = eqToken

    var neToken = newToken(TypeEnum.op)
    neToken.val.exec = newExec("<>", nativelib.ne, 3)
    cont["<>"] = neToken

    var ltToken = newToken(TypeEnum.op)
    ltToken.val.exec = newExec("<", nativelib.lt, 3)
    cont["<"] = ltToken

    var gtToken = newToken(TypeEnum.op)
    gtToken.val.exec = newExec(">", nativelib.gt, 3)
    cont[">"] = gtToken

    var lteqToken = newToken(TypeEnum.op)
    lteqToken.val.exec = newExec("<=", nativelib.lteq, 3)
    cont["<="] = lteqToken

    var gteqToken = newToken(TypeEnum.op)
    gteqToken.val.exec = newExec(">=", nativelib.gteq, 3)
    cont[">="] = gteqToken

