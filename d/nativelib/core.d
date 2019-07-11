module nativelib.core;

import std.conv;
import core.stdc.stdlib;

import common;
import token;
import bindmap;
import evalstack;
import arrlist;

Token ttypeof(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.datatype);
    if(args[1]){
        result.str = text(args[1].type) ~ "!";
    }else{
        result.str = "nil!";
    }
    return result;
}

Token quit(EvalStack stack, BindMap ctx){
    exit(0);
    return null;
}
