module native;

import token;
import bindmap;
import evalstack;

class Native {
    string  str;
    uint    explen;
    Token   function(EvalStack stack, BindMap ctx) run;

    this(){}
    this(string name, Token   function(EvalStack stack, BindMap ctx) f, uint len){
        str = name;
        run = f;
        explen = len;
    }
}
