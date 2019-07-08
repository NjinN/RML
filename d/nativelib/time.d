module nativelib.time;

import std.stdio;
import core.time;

import common;
import typeenum;
import token;
import bindmap;
import evalstack;

Token cost(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.nil);
    if(args[1].type == TypeEnum.block){
        auto before = MonoTime.currTime;
        Token temp = stack.eval(args[1].val.block, ctx);
        auto after = MonoTime.currTime;
        writeln(after - before);
        return result;
    }
    result.type = TypeEnum.err;
    result.val.str = "Type Mismatch";
    return result;
}
