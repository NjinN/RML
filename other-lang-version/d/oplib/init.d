module oplib.init;

import token;
import native;
import bindmap;

import nativelib.math;
import nativelib.compare;

void initOp(BindMap ctx){
    Token addToken = new Token(TypeEnum.op);
    addToken.exec = new Native("add", &add, 3);
    ctx.put("+", addToken);

    Token subToken = new Token(TypeEnum.op);
    subToken.exec = new Native("sub", &sub, 3);
    ctx.put("-", subToken);

    Token mulToken = new Token(TypeEnum.op);
    mulToken.exec = new Native("mul", &mul, 3);
    ctx.put("*", mulToken);

    Token divToken = new Token(TypeEnum.op);
    divToken.exec = new Native("div", &div, 3);
    ctx.put("/", divToken);

    Token eqToken = new Token(TypeEnum.op);
    eqToken.exec = new Native("=", &eq, 3);
    ctx.put("=", eqToken);

    Token neToken = new Token(TypeEnum.op);
    neToken.exec = new Native("<>", &ne, 3);
    ctx.put("<>", neToken);

    Token gtToken = new Token(TypeEnum.op);
    gtToken.exec = new Native(">", &gt, 3);
    ctx.put(">", gtToken);

    Token ltToken = new Token(TypeEnum.op);
    ltToken.exec = new Native("<", &lt, 3);
    ctx.put("<", ltToken);

    Token gteqToken = new Token(TypeEnum.op);
    gteqToken.exec = new Native(">=", &gteq, 3);
    ctx.put(">=", gteqToken);

    Token lteqToken = new Token(TypeEnum.op);
    lteqToken.exec = new Native("<=", &lteq, 3);
    ctx.put("<=", lteqToken);

}



