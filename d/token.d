module token;
import typeenum;
import bindmap;
import evalstack;
import native;
import func;
import strtool;
import totoken;

import std.stdio;
import std.conv;


alias TK = Token;
union TokenVal {
    bool        logic;
    byte        bt;
    char        cchar;
    int         integer;
    long        bigint;
    double      decimal;
    string      str;
    int[4]      integerArr;
    long[2]     bigintArr;
    double[2]   decimalArr;
    TK          tk;
    TK[]        block;
    Native      exec;
    Func        func;
}


class Token {
    TypeEnum    type;
    TokenVal    val;
    
    this(){}
    this(TypeEnum tp){
        type = tp;
    }

    string toStr(){
        switch(type){
            case TypeEnum.nil:
                return "nil";
            case TypeEnum.none:
                return "none";
            case TypeEnum.err:
                return "Error: " ~ val.str;
            case TypeEnum.lit_word:
                return "'" ~ val.str;
            case TypeEnum.get_word:
                return ":" ~ val.str;
            case TypeEnum.datatype:
                return val.str;
            case TypeEnum.logic:
                return text(val.logic);
            case TypeEnum.integer:
                return text(val.integer);
            case TypeEnum.decimal:
                return text(val.decimal);
            case TypeEnum.cchar:
                return text(val.cchar);
            case TypeEnum.str:
                return "\"" ~ val.str ~ "\"";
            case TypeEnum.block:
                string str = "[ ";
                for(int i=0; i < val.block.length; i++){
                    str = str ~ val.block[i].toStr() ~ " ";
                }
                str = str ~ "]";
                return str;
            case TypeEnum.paren:
                string str = "( ";
                for(int i=0; i < val.block.length; i++){
                    str = str ~ val.block[i].toStr() ~ " ";
                }
                str = str ~ ")";
                return str;
            case TypeEnum.word:
                return val.str;
            case TypeEnum.set_word:
                return val.str ~ ":";
            case TypeEnum.native:
                return "native: " ~ val.exec.str;
            case TypeEnum.func:
                string str = "func [ ";
                for(int i=0; i< val.func.args.length; i++){
                    str = str ~ val.func.args[i].toStr ~ " ";
                }
                str = str ~ "] [ ";
                for(int i=0; i< val.func.code.length; i++){
                    str = str ~ val.func.code[i].toStr ~ " ";
                }
                str = str ~ "]";
                return str;
            case TypeEnum.op:
                return "op: " ~ val.exec.str;
            default:
                return val.str;
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
                return val.exec.explen;
            case TypeEnum.func:
                return val.func.args.length + 1;
            case TypeEnum.op:
                return 3;
            default:
                return 1;
        }
    }

    Token getVal(BindMap ctx, EvalStack stack){
        Token result = new Token();
        switch(this.type){
            case TypeEnum.word:
                result = ctx.get(val.str);
                return result;
            case TypeEnum.lit_word:
                result.type = TypeEnum.word;
                result.val.str = val.str;
                return result;
            case TypeEnum.paren:
                result = stack.eval(val.block, ctx);
                return result;
            default:
                return this;
        }

    }
}




// void main(string[] args) {

//     Token tk = new Token(TypeEnum.str);
//     tk.val.str = "Hello world";
//     writeln(tk.type);
//     tk.echo();

//     // writeln(trim(" 123  "));
//     // Token tks = toToken("123");
//     Token[] tks = toTokens("   123 \"this is a string  with space   ^\"  ([ 1 2 3 ] and tranChar) \"  ([ 123 456 \"anthor ^\" str \" 987 ] 456) ");
//     for(int i=0; i<tks.length; i++){
//         tks[i].echo();
//     }
// }
