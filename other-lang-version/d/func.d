module func;

import token;
import bindmap;
import evalstack;
import common;
import arrlist;

class Func {
    ArrList!Token   args;
    ArrList!Token   codes;
    ArrList!int     quoteList;
    BindMap         ctx;   

    this(){}
    this(ArrList!Token a, ArrList!Token c){
        args = a;
        codes = c;
    }

    Token run(EvalStack stack, BindMap ctx){
        BindMap c = new BindMap(stack.mainCtx);
        for(int i=0; i<args.endIdx; i++){
            c.putNow(args.get(i).str, stack.line[stack.startPos.last + i + 1]);
        }
        return stack.eval(codes, c);
    }
}
