module token;
public import typeenum;
import bindmap;
import evalstack;
import native;
import func;

import std.stdio;
import std.conv;


alias TK = Token;



class Token {
    TypeEnum    type;
    union {
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
        Token       tk;
        Token[]     block;
        Native      exec;
        Func        func;
    }
    
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
                for(int i=0; i < block.length; i++){
                    str = str ~ block[i].toStr() ~ " ";
                }
                str = str ~ "]";
                return str;
            case TypeEnum.paren:
                string str = "( ";
                for(int i=0; i < block.length; i++){
                    str = str ~ block[i].toStr() ~ " ";
                }
                str = str ~ ")";
                return str;
            case TypeEnum.word:
                return str;
            case TypeEnum.set_word:
                return str ~ ":";
            case TypeEnum.native:
                return "native: " ~ exec.str;
            case TypeEnum.func:
                string str = "func [ ";
                for(int i=0; i< func.args.length; i++){
                    str = str ~ func.args[i].toStr ~ " ";
                }
                str = str ~ "] [ ";
                for(int i=0; i< func.code.length; i++){
                    str = str ~ func.code[i].toStr ~ " ";
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
                return cast(uint)(func.args.length + 1);
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
                result = ctx.get(str);
                return result;
            case TypeEnum.lit_word:
                result.type = TypeEnum.word;
                result.str = str;
                return result;
            case TypeEnum.paren:
                result = stack.eval(block, ctx);
                return result;
            default:
                return this;
        }

    }
}




// void main(string[] args) {

//     Token tk = new Token(TypeEnum.str);
//     tk.str = "Hello world";
//     writeln(tk.type);
//     tk.echo();

//     // writeln(trim(" 123  "));
//     // Token tks = toToken("123");
//     Token[] tks = toTokens("   123 \"this is a string  with space   ^\"  ([ 1 2 3 ] and tranChar) \"  ([ 123 456 \"anthor ^\" str \" 987 ] 456) ");
//     for(int i=0; i<tks.length; i++){
//         tks[i].echo();
//     }
// }