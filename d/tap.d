module tap;

import token;

class Tap {
    string  str;
    Token   val;

    this(){}
    this(string s, Token v){
        str = s;
        val = v;
    }
}
