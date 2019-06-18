import nativeLib

proc initNative*(cont: ptr BindMap[ptr Token])=
    var addToken = newToken(TypeEnum.native)
    addToken.val.exec = newExec("add", nativeLib.plus, 3)
    cont["add"] = addToken

    var subToken = newToken(TypeEnum.native)
    subToken.val.exec = newExec("sub", nativelib.minus, 3)
    cont["sub"] = subToken

    var mulToken = newToken(TypeEnum.native)
    mulToken.val.exec = newExec("mul", nativelib.multiply, 3)
    cont["mul"] = mulToken

    var divToken = newToken(TypeEnum.native)
    divToken.val.exec = newExec("div", nativelib.divide, 3)
    cont["div"] = divToken

    var cpuTimeToken = newToken(TypeEnum.native)
    cpuTimeToken.val.exec = newExec("cputime", nativelib.getCpuTime, 1)
    cont["cputime"] = cpuTimeToken

    var gmtToken = newToken(TypeEnum.native)
    gmtToken.val.exec = newExec("gmt", nativelib.gmt, 1)
    cont["gmt"] = gmtToken

    var printToken = newToken(TypeEnum.native)
    printToken.val.exec = newExec("print", nativelib.print, 2)
    cont["print"] = printToken

    var ifToken = newToken(TypeEnum.native)
    ifToken.val.exec = newExec("if", nativelib.iff, 3)
    cont["if"] = ifToken

    var eitherToken = newToken(TypeEnum.native)
    eitherToken.val.exec = newExec("either", nativelib.either, 4)
    cont["either"] = eitherToken
  
    var funcToken = newToken(TypeEnum.native)
    funcToken.val.exec = newExec("func", nativelib.fc, 3)
    cont["func"] = funcToken
