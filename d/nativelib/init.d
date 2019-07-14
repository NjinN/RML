module nativelib.init;

import token;
import native;
import bindmap;

import arrlist;
import nativelib.core;
import nativelib.math;
import nativelib.compare;
import nativelib.output;
import nativelib.deffunc;
import nativelib.control;
import nativelib.time;
import nativelib.series;

void initNative(BindMap ctx){
    Token quitToken = new Token(TypeEnum.native);
    quitToken.exec = new Native("quit", &quit, 1);
    ctx.put("quit", quitToken);
    ctx.put("q", quitToken);

    Token typeofToken = new Token(TypeEnum.native);
    typeofToken.exec = new Native("type?", &ttypeof, 2);
    ctx.put("type?", typeofToken);

    Token doToken = new Token(TypeEnum.native);
    doToken.exec = new Native("do", &ddo, 2);
    ctx.put("do", doToken);

    Token addToken = new Token(TypeEnum.native);
    addToken.exec = new Native("add", &add, 3);
    ctx.put("add", addToken);

    Token subToken = new Token(TypeEnum.native);
    subToken.exec = new Native("sub", &sub, 3);
    ctx.put("sub", subToken);

    Token mulToken = new Token(TypeEnum.native);
    mulToken.exec = new Native("mul", &mul, 3);
    ctx.put("mul", mulToken);

    Token divToken = new Token(TypeEnum.native);
    divToken.exec = new Native("div", &div, 3);
    ctx.put("div", divToken);

    Token modToken = new Token(TypeEnum.native);
    modToken.exec = new Native("mod", &mod, 3);
    ctx.put("mod", modToken);

    Token eqToken = new Token(TypeEnum.native);
    eqToken.exec = new Native("eq", &eq, 3);
    ctx.put("eq", eqToken);

    Token neToken = new Token(TypeEnum.native);
    neToken.exec = new Native("ne", &ne, 3);
    ctx.put("ne", neToken);

    Token gtToken = new Token(TypeEnum.native);
    gtToken.exec = new Native("gt", &gt, 3);
    ctx.put("gt", gtToken);

    Token ltToken = new Token(TypeEnum.native);
    ltToken.exec = new Native("lt", &lt, 3);
    ctx.put("lt", ltToken);

    Token gteqToken = new Token(TypeEnum.native);
    gteqToken.exec = new Native("gteq", &gteq, 3);
    ctx.put("gteq", gteqToken);

    Token lteqToken = new Token(TypeEnum.native);
    lteqToken.exec = new Native("lteq", &lteq, 3);
    ctx.put("lteq", lteqToken);

    Token printToken = new Token(TypeEnum.native);
    printToken.exec = new Native("print", &print, 2);
    ctx.put("print", printToken);

    Token funcToken = new Token(TypeEnum.native);
    funcToken.exec = new Native("func", &defFunc, 3);
    ctx.put("func", funcToken);

    Token iifToken = new Token(TypeEnum.native);
    iifToken.exec = new Native("if", &iif, 3);
    ctx.put("if", iifToken);

    Token eitherToken = new Token(TypeEnum.native);
    eitherToken.exec = new Native("either", &either, 4);
    ctx.put("either", eitherToken);

    Token loopToken = new Token(TypeEnum.native);
    loopToken.exec = new Native("loop", &loop, 3);
    ctx.put("loop", loopToken);

    Token repeatToken = new Token(TypeEnum.native);
    repeatToken.exec = new Native("repeat", &repeat, 4);
    repeatToken.exec.quoteList = new ArrList!int([0, 1, 1]);
    ctx.put("repeat", repeatToken);

    Token forToken = new Token(TypeEnum.native);
    forToken.exec = new Native("for", &ffor, 6);
    forToken.exec.quoteList = new ArrList!int([0, 1, 1, 1, 1]);
    ctx.put("for", forToken);

    Token whileToken = new Token(TypeEnum.native);
    whileToken.exec = new Native("while", &wwhile, 3);
    ctx.put("while", whileToken);

    Token breakToken = new Token(TypeEnum.native);
    breakToken.exec = new Native("break", &bbreak, 1);
    ctx.put("break", breakToken);

    Token continueToken = new Token(TypeEnum.native);
    continueToken.exec = new Native("continue", &ccontinue, 1);
    ctx.put("continue", continueToken);

    Token costToken = new Token(TypeEnum.native);
    costToken.exec = new Native("cost", &cost, 2);
    ctx.put("cost", costToken);

    Token lenToken = new Token(TypeEnum.native);
    lenToken.exec = new Native("len?", &len, 2);
    ctx.put("len?", lenToken);

    Token appendToken = new Token(TypeEnum.native);
    appendToken.exec = new Native("append", &append, 3);
    ctx.put("append", appendToken);
}



