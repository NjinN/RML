import ../native/nativeLib

proc initOp*(cont: var BindMap[Token])=
    var addToken = initToken(TypeEnum.op)
    addToken.val.exec = newExec("+", nativelib.plus, 3)
    cont["+"] = addToken

    var subToken = initToken(TypeEnum.op)
    subToken.val.exec = newExec("-", nativelib.minus, 3)
    cont["-"] = subToken

    var mulToken = initToken(TypeEnum.op)
    mulToken.val.exec = newExec("*", nativelib.multiply, 3)
    cont["*"] = mulToken

    var divToken = initToken(TypeEnum.op)
    divToken.val.exec = newExec("/", nativelib.divide, 3)
    cont["/"] = divToken

    var eqToken = initToken(TypeEnum.op)
    eqToken.val.exec = newExec("=", nativelib.eq, 3)
    cont["="] = eqToken

    var neToken = initToken(TypeEnum.op)
    neToken.val.exec = newExec("<>", nativelib.ne, 3)
    cont["<>"] = neToken

    var ltToken = initToken(TypeEnum.op)
    ltToken.val.exec = newExec("<", nativelib.lt, 3)
    cont["<"] = ltToken

    var gtToken = initToken(TypeEnum.op)
    gtToken.val.exec = newExec(">", nativelib.gt, 3)
    cont[">"] = gtToken

    var lteqToken = initToken(TypeEnum.op)
    lteqToken.val.exec = newExec("<=", nativelib.lteq, 3)
    cont["<="] = lteqToken

    var gteqToken = initToken(TypeEnum.op)
    gteqToken.val.exec = newExec(">=", nativelib.gteq, 3)
    cont[">="] = gteqToken

