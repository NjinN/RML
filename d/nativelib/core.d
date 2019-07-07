module nativelib.core;

import std.conv;

import common;
import typeenum;
import token;
import bindmap;
import evalstack;

Token ttypeof(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.datatype);
    if(args[1]){
        result.val.str = text(args[1].type) ~ "!";
    }else{
        result.val.str = "nil!";
    }
    return result;
}


