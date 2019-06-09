import ../native/nativelib

proc initOp*(cont: ref Context)=
    var addToken = new Token
    addToken.tp = TypeEnum.op
    addToken.val.exec = nativelib.plus
    addToken.explen = 3
    cont.map[cstring("+")] = addToken

    var subToken = new Token
    subToken.tp = TypeEnum.op
    subToken.val.exec = nativelib.minus
    subToken.explen = 3
    cont.map[cstring("-")] = subToken

    var mulToken = new Token
    mulToken.tp = TypeEnum.op
    mulToken.val.exec = nativelib.multiply
    mulToken.explen = 3
    cont.map[cstring("*")] = mulToken

    var divToken = new Token
    divToken.tp = TypeEnum.op
    divToken.val.exec = nativelib.divide
    divToken.explen = 3
    cont.map[cstring("/")] = divToken

