module func;

import token;
import bindmap;
import evalstack;
import common;
import arrlist;

class Func {
    ArrList!Token   args;
    ArrList!Token   code;
    ArrList!int     quoteList;

    this(){}
    this(ArrList!Token a, ArrList!Token c){
        args = a;
        code = c;
    }

    Token run(EvalStack stack, BindMap ctx){
        BindMap c = new BindMap(stack.mainCtx);
        for(int i=0; i<args.endIdx; i++){
            c.putNow(args.get(i).word.name, stack.line[stack.startPos.last + i + 1]);
        }
        return stack.eval(code, c);
    }
}