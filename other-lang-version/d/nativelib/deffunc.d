module nativelib.deffunc;

import common;
import token;
import bindmap;
import evalstack;
import func;
import arrlist;
import std.stdio;
Token defFunc(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token();
    if(args[1].type != TypeEnum.block || args[2].type != TypeEnum.block){
        result.type = TypeEnum.err;
        result.str = "Type Mismatch";
        return result;
    }
    for(int i=0; i<args[1].block.endIdx; i++){
        if(args[1].block.get(i).type != TypeEnum.word){
            result.type = TypeEnum.err;
            result.str = "Type Mismatch";
            return result;
        }
    }
    result.type = TypeEnum.func;
    result.func = new Func(args[1].block, args[2].block);
    result.func.ctx = new BindMap(ctx);
    return result;
}

