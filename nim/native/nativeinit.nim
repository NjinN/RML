import nativelib

proc initNative*(cont: ref Context)=
    var addToken = newToken(TypeEnum.native, 3)
    addToken.val.exec = newExec("add", nativeLib.plus)
    cont.map[cstring("add")] = addToken

    var subToken = newToken(TypeEnum.native, 3)
    subToken.val.exec = newExec("sub", nativelib.minus)
    cont.map[cstring("sub")] = subToken

    var mulToken = newToken(TypeEnum.native, 3)
    mulToken.val.exec = newExec("mul", nativelib.multiply)
    cont.map[cstring("mul")] = mulToken

    var divToken = newToken(TypeEnum.native, 3)
    divToken.val.exec = newExec("div", nativelib.divide)
    cont.map[cstring("div")] = divToken

    var cpuTimeToken = newToken(TypeEnum.native, 1)
    cpuTimeToken.val.exec = newExec("cputime", nativelib.getCpuTime)
    cont.map[cstring("cputime")] = cpuTimeToken

    var gmtToken = newToken(TypeEnum.native, 1)
    gmtToken.val.exec = newExec("gmt", nativelib.gmt)
    cont.map[cstring("gmt")] = gmtToken

    var printToken = newToken(TypeEnum.native, 2)
    printToken.val.exec = newExec("print", nativelib.print)
    cont.map[cstring("print")] = printToken

    var ifToken = newToken(TypeEnum.native, 3)
    ifToken.val.exec = newExec("if", nativelib.iff)
    cont.map[cstring("if")] = ifToken

    var eitherToken = newToken(TypeEnum.native, 4)
    eitherToken.val.exec = newExec("either", nativelib.either)
    cont.map[cstring("either")] = eitherToken
  
    var funcToken = newToken(TypeEnum.native, 3)
    funcToken.val.exec = newExec("func", nativelib.fc)
    cont.map[cstring("func")] = funcToken
