import nativelib

proc initNative*(cont: ref Context)=
    var addToken = new Token
    addToken.tp = TypeEnum.native
    addToken.val.exec = nativelib.plus
    addToken.explen = 3
    cont.map[cstring("add")] = addToken

    var subToken = new Token
    subToken.tp = TypeEnum.native
    subToken.val.exec = nativelib.minus
    subToken.explen = 3
    cont.map[cstring("sub")] = subToken

    var mulToken = new Token
    mulToken.tp = TypeEnum.native
    mulToken.val.exec = nativelib.multiply
    mulToken.explen = 3
    cont.map[cstring("mul")] = mulToken

    var divToken = new Token
    divToken.tp = TypeEnum.native
    divToken.val.exec = nativelib.divide
    divToken.explen = 3
    cont.map[cstring("div")] = divToken

    var cpuTimeToken = new Token
    cpuTimeToken.tp = TypeEnum.native
    cpuTimeToken.val.exec = nativelib.getCpuTime
    cpuTimeToken.explen = 1
    cont.map[cstring("cputime")] = cpuTimeToken

    var gmtToken = new Token
    gmtToken.tp = TypeEnum.native
    gmtToken.val.exec = nativelib.gmt
    gmtToken.explen = 1
    cont.map[cstring("gmt")] = gmtToken

    var printToken = new Token
    printToken.tp = TypeEnum.native
    printToken.val.exec = nativelib.print
    printToken.explen = 2
    cont.map[cstring("print")] = printToken
