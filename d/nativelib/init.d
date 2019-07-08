module nativelib.init;

import typeenum;
import token;
import native;
import bindmap;

import nativelib.core;
import nativelib.math;
import nativelib.compare;
import nativelib.output;
import nativelib.deffunc;
import nativelib.control;
import nativelib.time;

void initNative(BindMap ctx){
    Token typeofToken = new Token(TypeEnum.native);
    typeofToken.val.exec = new Native("type?", &ttypeof, 2);
    ctx.put("type?", typeofToken);

    Token addToken = new Token(TypeEnum.native);
    addToken.val.exec = new Native("add", &add, 3);
    ctx.put("add", addToken);

    Token subToken = new Token(TypeEnum.native);
    subToken.val.exec = new Native("sub", &sub, 3);
    ctx.put("sub", subToken);

    Token mulToken = new Token(TypeEnum.native);
    mulToken.val.exec = new Native("mul", &mul, 3);
    ctx.put("mul", mulToken);

    Token divToken = new Token(TypeEnum.native);
    divToken.val.exec = new Native("div", &div, 3);
    ctx.put("div", divToken);

    Token eqToken = new Token(TypeEnum.native);
    eqToken.val.exec = new Native("eq", &eq, 3);
    ctx.put("eq", eqToken);

    Token neToken = new Token(TypeEnum.native);
    neToken.val.exec = new Native("ne", &ne, 3);
    ctx.put("ne", neToken);

    Token gtToken = new Token(TypeEnum.native);
    gtToken.val.exec = new Native("gt", &gt, 3);
    ctx.put("gt", gtToken);

    Token ltToken = new Token(TypeEnum.native);
    ltToken.val.exec = new Native("lt", &lt, 3);
    ctx.put("lt", ltToken);

    Token gteqToken = new Token(TypeEnum.native);
    gteqToken.val.exec = new Native("gteq", &gteq, 3);
    ctx.put("gteq", gteqToken);

    Token lteqToken = new Token(TypeEnum.native);
    lteqToken.val.exec = new Native("lteq", &lteq, 3);
    ctx.put("lteq", lteqToken);

    Token printToken = new Token(TypeEnum.native);
    printToken.val.exec = new Native("print", &print, 2);
    ctx.put("print", printToken);

    Token funcToken = new Token(TypeEnum.native);
    funcToken.val.exec = new Native("func", &defFunc, 3);
    ctx.put("func", funcToken);

    Token iifToken = new Token(TypeEnum.native);
    iifToken.val.exec = new Native("if", &iif, 3);
    ctx.put("if", iifToken);

    Token eitherToken = new Token(TypeEnum.native);
    eitherToken.val.exec = new Native("either", &either, 4);
    ctx.put("either", eitherToken);

    Token loopToken = new Token(TypeEnum.native);
    loopToken.val.exec = new Native("loop", &loop, 3);
    ctx.put("loop", loopToken);

    Token repeatToken = new Token(TypeEnum.native);
    repeatToken.val.exec = new Native("repeat", &repeat, 4);
    ctx.put("repeat", repeatToken);

    Token forToken = new Token(TypeEnum.native);
    forToken.val.exec = new Native("for", &ffor, 6);
    ctx.put("for", forToken);

    Token whileToken = new Token(TypeEnum.native);
    whileToken.val.exec = new Native("while", &wwhile, 3);
    ctx.put("while", whileToken);

    Token breakToken = new Token(TypeEnum.native);
    breakToken.val.exec = new Native("break", &bbreak, 1);
    ctx.put("break", breakToken);

    Token continueToken = new Token(TypeEnum.native);
    continueToken.val.exec = new Native("continue", &ccontinue, 1);
    ctx.put("continue", continueToken);

    Token costToken = new Token(TypeEnum.native);
    costToken.val.exec = new Native("cost", &cost, 2);
    ctx.put("cost", costToken);
}



