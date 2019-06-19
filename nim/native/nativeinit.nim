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

    var costToken = newToken(TypeEnum.native)
    costToken.val.exec = newExec("cost", nativelib.cost, 2)
    cont["cost"] = costToken

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

    var loopToken = newToken(TypeEnum.native)
    loopToken.val.exec = newExec("loop", nativelib.loop, 3)
    cont["loop"] = loopToken

    var repeatToken = newToken(TypeEnum.native)
    repeatToken.val.exec = newExec("repeat", nativelib.repeat, 4)
    cont["repeat"] = repeatToken

    var forToken = newToken(TypeEnum.native)
    forToken.val.exec = newExec("for", nativelib.ffor, 6)
    cont["for"] = forToken

    var whileToken = newToken(TypeEnum.native)
    whileToken.val.exec = newExec("while", nativelib.wwhile, 3)
    cont["while"] = whileToken
