module nativelib.time;

import std.stdio;
import core.time;

import common;
import token;
import bindmap;
import evalstack;
import arrlist;

Token cost(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.nil);
    if(args[1].type == TypeEnum.block){
        auto before = MonoTime.currTime;
        Token temp = stack.eval(args[1].block, ctx);
        auto after = MonoTime.currTime;
        writeln("takes: ", (after - before).total!"msecs", " ms");
        return result;
    }
    result.type = TypeEnum.err;
    result.str = "Type Mismatch";
    return result;
}
