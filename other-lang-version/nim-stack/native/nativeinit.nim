import nativeLib

proc initNative*(cont: var BindMap[Token])=
    var addToken = initToken(TypeEnum.native)
    addToken.val.exec = newExec("add", nativeLib.plus, 3)
    cont["add"] = addToken

    var subToken = initToken(TypeEnum.native)
    subToken.val.exec = newExec("sub", nativelib.minus, 3)
    cont["sub"] = subToken

    var mulToken = initToken(TypeEnum.native)
    mulToken.val.exec = newExec("mul", nativelib.multiply, 3)
    cont["mul"] = mulToken

    var divToken = initToken(TypeEnum.native)
    divToken.val.exec = newExec("div", nativelib.divide, 3)
    cont["div"] = divToken

    var cpuTimeToken = initToken(TypeEnum.native)
    cpuTimeToken.val.exec = newExec("cputime", nativelib.getCpuTime, 1)
    cont["cputime"] = cpuTimeToken

    var gmtToken = initToken(TypeEnum.native)
    gmtToken.val.exec = newExec("gmt", nativelib.gmt, 1)
    cont["gmt"] = gmtToken

    var costToken = initToken(TypeEnum.native)
    costToken.val.exec = newExec("cost", nativelib.cost, 2)
    cont["cost"] = costToken

    var printToken = initToken(TypeEnum.native)
    printToken.val.exec = newExec("print", nativelib.print, 2)
    cont["print"] = printToken

    var typeOfToken = initToken(TypeEnum.native)
    typeOfToken.val.exec = newExec("type?", nativelib.ttypeof, 2)
    cont["type?"] = typeOfToken

    var ifToken = initToken(TypeEnum.native)
    ifToken.val.exec = newExec("if", nativelib.iff, 3)
    cont["if"] = ifToken

    var eitherToken = initToken(TypeEnum.native)
    eitherToken.val.exec = newExec("either", nativelib.either, 4)
    cont["either"] = eitherToken
  
    var funcToken = initToken(TypeEnum.native)
    funcToken.val.exec = newExec("func", nativelib.fc, 3)
    cont["func"] = funcToken

    var loopToken = initToken(TypeEnum.native)
    loopToken.val.exec = newExec("loop", nativelib.loop, 3)
    cont["loop"] = loopToken

    var repeatToken = initToken(TypeEnum.native)
    repeatToken.val.exec = newExec("repeat", nativelib.repeat, 4)
    cont["repeat"] = repeatToken

    var forToken = initToken(TypeEnum.native)
    forToken.val.exec = newExec("for", nativelib.ffor, 6)
    cont["for"] = forToken

    var whileToken = initToken(TypeEnum.native)
    whileToken.val.exec = newExec("while", nativelib.wwhile, 3)
    cont["while"] = whileToken

    var breakToken = initToken(TypeEnum.native)
    breakToken.val.exec = newExec("break", nativelib.bbreak, 1)
    cont["break"] = breakToken

    var continueToken = initToken(TypeEnum.native)
    continueToken.val.exec = newExec("continue", nativelib.ccontinue, 1)
    cont["continue"] = continueToken
