module nativelib.output;

import std.stdio;

import common;
import token;
import bindmap;
import evalstack;
import arrlist;

Token print(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    if(args[1].type == TypeEnum.block){
        for(int i=0; i<args[1].block.length; i++){
            write(args[1].block[i].outputStr());
            write("\n");
        }
    }else{
        writeln(args[1].outputStr);
    }
    return new Token(TypeEnum.nil);
}
