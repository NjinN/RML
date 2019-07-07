module nativelib.output;

import std.stdio;

import common;
import typeenum;
import token;
import bindmap;
import evalstack;

Token print(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    if(args[1].type == TypeEnum.block){
        for(int i=0; i<args[1].val.block.length; i++){
            write(args[1].val.block[i].outputStr());
            write("\n");
        }
    }else{
        writeln(args[1].outputStr);
    }
    return new Token(TypeEnum.nil);
}
