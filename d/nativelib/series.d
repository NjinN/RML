module nativelib.series;

import common;
import token;
import bindmap;
import evalstack;
import arrlist;

Token len(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.integer);
    if(args[1].type == TypeEnum.block){
        result.integer = cast(int)args[1].block.len;
        return result;
    }else if(args[1].type == TypeEnum.str){
        result.integer = cast(int)args[1].str.length;
        return result;
    }
    result.type = TypeEnum.err;
    result.str = "Type Mismatch";
    return result;
}

Token append(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    if(args[1].type == TypeEnum.block){
        args[1].block.add(args[2].dup);
        return args[1];
    }else if(args[1].type == TypeEnum.str){
        args[1].str ~= args[2].outputStr;
        return args[1];
    }
    Token result = new Token(TypeEnum.err);
    result.str = "Type Mismatch";
    return result;
}


