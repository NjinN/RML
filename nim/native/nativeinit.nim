import nativelib

proc initNative*(cont: ref Context)=
    var addToken = new Token
    addToken.tp = TypeEnum.native
    addToken.explen = 3
    addToken.val.exec = newExec("add", nativeLib.plus)
    cont.map[cstring("add")] = addToken

    var subToken = new Token
    subToken.tp = TypeEnum.native
    subToken.explen = 3
    subToken.val.exec = newExec("sub", nativelib.minus)
    cont.map[cstring("sub")] = subToken

    var mulToken = new Token
    mulToken.tp = TypeEnum.native
    mulToken.explen = 3
    mulToken.val.exec = newExec("mul", nativelib.multiply)
    cont.map[cstring("mul")] = mulToken

    var divToken = new Token
    divToken.tp = TypeEnum.native
    divToken.explen = 3
    divToken.val.exec = newExec("div", nativelib.divide)
    cont.map[cstring("div")] = divToken

    var cpuTimeToken = new Token
    cpuTimeToken.tp = TypeEnum.native
    cpuTimeToken.explen = 1
    cpuTimeToken.val.exec = newExec("cputime", nativelib.getCpuTime)
    cont.map[cstring("cputime")] = cpuTimeToken

    var gmtToken = new Token
    gmtToken.tp = TypeEnum.native
    gmtToken.explen = 1
    gmtToken.val.exec = newExec("gmt", nativelib.gmt)
    cont.map[cstring("gmt")] = gmtToken

    var printToken = new Token
    printToken.tp = TypeEnum.native
    printToken.explen = 2
    printToken.val.exec = newExec("print", nativelib.print)
    cont.map[cstring("print")] = printToken
  
