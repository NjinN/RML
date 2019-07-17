module token;
public import typeenum;
import bindmap;
import evalstack;
import native;
import func;
import arrlist;

import std.stdio;
import std.conv;



class Token {
    TypeEnum    type;
    union {
        bool            logic;
        byte            bt;
        char            cchar;
        int             integer;
        long            bigint;
        double          decimal;
        string          str;
        int[4]          integerArr;
        long[2]         bigintArr;
        double[2]       decimalArr;
        Token           tk;
        ArrList!Token   block;
        Native          exec;
        Func            func;
        byte[16]        all;
    }
    
    this(){}
    this(TypeEnum tp){
        type = tp;
    }
    this(TypeEnum tp, string s){
        type = tp;
        str = s;
    }

    string toStr(){
        switch(type){
            case TypeEnum.nil:
                return "nil";
            case TypeEnum.none:
                return "none";
            case TypeEnum.err:
                return "Error: " ~ str;
            case TypeEnum.lit_word:
                return "'" ~ str;
            case TypeEnum.get_word:
                return ":" ~ str;
            case TypeEnum.datatype:
                return str;
            case TypeEnum.logic:
                return text(logic);
            case TypeEnum.integer:
                return text(integer);
            case TypeEnum.decimal:
                return text(decimal);
            case TypeEnum.cchar:
                return text(cchar);
            case TypeEnum.str:
                return "\"" ~ str ~ "\"";
            case TypeEnum.block:
                string str = "[ ";
                for(int i=0; i < block.len; i++){
                    str = str ~ block.get(i).toStr() ~ " ";
                }
                str = str ~ "]";
                return str;
            case TypeEnum.paren:
                string str = "( ";
                for(int i=0; i < block.len; i++){
                    str = str ~ block.get(i).toStr() ~ " ";
                }
                str = str ~ ")";
                return str;
            case TypeEnum.set_word:
                return str ~ ":";
            case TypeEnum.native:
                return "native: " ~ exec.str;
            case TypeEnum.func:
                string str = "func [ ";
                for(int i=0; i< func.args.len; i++){
                    str = str ~ func.args.get(i).toStr ~ " ";
                }
                str = str ~ "] [ ";
                for(int i=0; i< func.codes.len; i++){
                    str = str ~ func.codes.get(i).toStr ~ " ";
                }
                str = str ~ "]";
                return str;
            case TypeEnum.op:
                return "op: " ~ exec.str;
            default:
                return str;
        }
    }

    string outputStr(){
        string str = this.toStr();
        if(this.type == TypeEnum.str){
            str = str[1..str.length-1];
        }
        return str;
    }

    void echo(){
        writeln(this.outputStr());
    }

    uint explen(){
        switch(this.type){
            case TypeEnum.set_word:
                return 2;
            case TypeEnum.native:
                return exec.explen;
            case TypeEnum.func:
                return cast(uint)(func.args.len + 1);
            case TypeEnum.op:
                return 3;
            default:
                return 1;
        }
    }

    Token getVal(BindMap ctx, EvalStack stack){
        Token result;
        switch(this.type){
            case TypeEnum.word:
                result = ctx.map.get(str, null);
                if(result){
                    return result;
                }else{
                    result = ctx.get(str);
                    return result;
                }  
            case TypeEnum.lit_word:
                result = new Token(TypeEnum.word);
                result.str = str;
                return result;
            case TypeEnum.paren:
                result = stack.eval(block, ctx);
                return result;
            default:
                return this;
        }

    }

    Token dup(){
        Token result = new Token();
        result.type = type;
        result.all = all;
        return result;
    }

    void copy(Token val){
        type = val.type;
        all = val.all;
    }
}



// void main(string[] args) {

    // Token tk = new Token(TypeEnum.integer);
    // tk.integer = 99;
    // Token tk2 = tk.dup;
    // tk2.integer = 100;

    // writeln(tk2.integer);

    // Token tk = new Token(TypeEnum.str);
    // tk.str = "Hello world";
    // writeln(tk.type);
    // tk.echo();

    // // writeln(trim(" 123  "));
    // // Token tks = toToken("123");
    // Token[] tks = toTokens("   123 \"this is a string  with space   ^\"  ([ 1 2 3 ] and tranChar) \"  ([ 123 456 \"anthor ^\" str \" 987 ] 456) ");
    // for(int i=0; i<tks.length; i++){
    //     tks[i].echo();
    // }
// }
