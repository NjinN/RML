module nativelib.deffunc;

import common;
import typeenum;
import token;
import bindmap;
import evalstack;
import func;

Token defFunc(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token();
    if(args[1].type != TypeEnum.block || args[2].type != TypeEnum.block){
        result.type = TypeEnum.err;
        result.val.str = "Type Mismatch";
        return result;
    }
    for(int i=0; i<args[1].val.block.length; i++){
        if(args[1].val.block[i].type != TypeEnum.word){
            result.type = TypeEnum.err;
            result.val.str = "Type Mismatch";
            return result;
        }
    }
    result.type = TypeEnum.func;
    result.val.func = new Func(args[1].val.block, args[2].val.block);
    return result;
}