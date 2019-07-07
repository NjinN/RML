module func;

import token;
import bindmap;
import evalstack;
import common;

class Func {
    Token[]     args;
    Token[]     code;

    this(){}
    this(Token[] a, Token[] c){
        args = a;
        code = c;
    }

    Token run(EvalStack stack, BindMap ctx){
        BindMap c = new BindMap();
        c.father = ctx;
        for(int i=0; i<args.length; i++){
            c.put(args[i].val.str, stack.line[last(stack.startPos) + i +1]);
        }
        return stack.eval(code, c);
    }
}