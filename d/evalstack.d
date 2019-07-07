module evalstack;

import std.stdio;

import typeenum;
import token;
import bindmap;
import totoken;
import common;

class EvalStack {
    uint[]              startPos;
    uint[]              endPos;
    Token[1024 * 1024]  line;
    uint                idx;

    void init(){
        startPos.length = 0;
        endPos.length = 0;
        idx = 0;
    }

    void push(Token t){
        line[idx] = t;
        idx += 1;
    }

    Token eval(string inpStr, BindMap ctx){
        return eval(toTokens(inpStr), ctx);
    }

    Token eval(Token[] inp, BindMap ctx){
        Token result = new Token();
        if(inp.length == 0){
            return result;
        }

        uint startIdx = idx;
        uint startDeep = endPos.length;

        uint i = 0;
        while(i < inp.length){
            Token nowToken = inp[i];
            Token nextToken;
            if(i < inp.length-1){
                nextToken = inp[i + 1].getVal(ctx, this);
            }
            
            if(nextToken && nextToken.type == TypeEnum.op && (startDeep == 0 || idx > endPos[startDeep - 1])){
                if(startPos.length == 0 || line[last(startPos)].type != TypeEnum.op){
                    startPos ~= idx;
                    push(nextToken);
                    push(nowToken.getVal(ctx, this));
                    endPos ~= idx;
                }else if(startPos.length == 0 || line[last(startPos)].type == TypeEnum.op){
                    push(nowToken.getVal(ctx, this));
                    evalExp(ctx);
                    push(line[idx - 1]);
                    line[idx - 2] = nextToken;
                    startPos ~= idx - 2;
                    endPos ~= idx;
                }
                i += 1;
            }else{
                nowToken = nowToken.getVal(ctx, this);
                if(nowToken && nowToken.type == TypeEnum.err){
                    return nowToken;
                }else if(nowToken && nowToken.type == TypeEnum.op){
                    if(idx > startIdx){
                        startPos ~= idx - 1;
                        push(line[idx - 1]);
                        line[idx - 2] = nowToken;
                        endPos ~= idx;
                    }else{
                        result.type = TypeEnum.err;
                        result.val.str = "Illegal grammar!!!";
                        return result;
                    }
                }else if(nowToken && nowToken.type < TypeEnum.set_word){
                    push(nowToken);
                }else{
                    startPos ~= idx;
                    endPos ~= idx + nowToken.explen() - 1;
                    push(nowToken);
                    // writeln("last endPos is ", last(endPos));
                    // writeln("last startPos is ", last(startPos));
                }
            }
         
            while(endPos.length > startDeep && idx == last(endPos) + 1){
                evalExp(ctx);
            }

            i += 1;
        }

        result = line[idx - 1];
        idx = startIdx;
        return result;
    }


    void evalExp(BindMap ctx){
        Token temp;
        try{
            switch(line[last(startPos)].type){
                case TypeEnum.set_word:
                    ctx.put(line[last(startPos)].val.str, line[last(endPos)]);
                    temp = line[last(endPos)];
                    break;
                case TypeEnum.native:
                    temp = line[last(startPos)].val.exec.run(this, ctx);
                    break;
                case TypeEnum.op:
                    temp = line[last(startPos)].val.exec.run(this, ctx);
                    break;
                case TypeEnum.func:
                    temp = line[last(startPos)].val.func.run(this, ctx);
                    break;
                default:
                    temp = new Token();
            }
        }catch(Exception e){
            if(e.msg == "break" || e.msg == "continue"){
                throw e;
            }else{
                temp = new Token(TypeEnum.err);
                temp.val.str = "Illegal grammar!!!";
            }
        }finally{
            line[last(startPos)] = temp;
            idx = last(startPos) + 1;
            startPos.length = startPos.length - 1;
            endPos.length = endPos.length - 1;
        }

    }


}


