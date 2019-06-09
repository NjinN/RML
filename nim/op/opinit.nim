import ../native/nativelib

proc initOp*(cont: ref Context)=
    var addToken = new Token
    addToken.tp = TypeEnum.op
    addToken.explen = 3
    addToken.val.exec = newExec("+", nativelib.plus)
    cont.map[cstring("+")] = addToken

    var subToken = new Token
    subToken.tp = TypeEnum.op
    subToken.explen = 3
    subToken.val.exec = newExec("-", nativelib.minus)
    cont.map[cstring("-")] = subToken

    var mulToken = new Token
    mulToken.tp = TypeEnum.op
    mulToken.explen = 3
    mulToken.val.exec = newExec("*", nativelib.multiply)
    cont.map[cstring("*")] = mulToken

    var divToken = new Token
    divToken.tp = TypeEnum.op
    divToken.explen = 3
    divToken.val.exec = newExec("/", nativelib.divide)
    cont.map[cstring("/")] = divToken

