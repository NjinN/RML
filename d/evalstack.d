module evalstack;

import std.stdio;

import token;
import bindmap;
import totoken;
import common;
import arrlist;

class EvalStack {
    ArrList!uint        startPos;
    ArrList!uint        endPos;
    ArrList!int         quoteList;
    Token[1024 * 1024]  line;
    uint                idx;
    BindMap             mainCtx;
    BindMap             libCtx;

    this(){
        startPos = new ArrList!uint;
        endPos = new ArrList!uint;
        quoteList = new ArrList!int;
    }

    void init(){
        startPos.clear;
        endPos.clear;
        quoteList.clear;
        idx = 0;
    }

    void push(Token t){
        line[idx] = t;
        idx += 1;
    }

    Token eval(string inpStr, BindMap ctx){
        return eval(toTokens(inpStr), ctx);
    }

    Token eval(ArrList!Token inp, BindMap ctx){
        Token result = new Token();
        if(inp.len == 0){
            return result;
        }

        uint startIdx = idx;
        uint startDeep = endPos.len;

        uint i = 0;
        while(i < inp.len){
            Token nowToken = inp.get(i);
            Token nextToken;
            if(i < inp.len-1){
                nextToken = inp.get(i + 1);
                if(nextToken.type == TypeEnum.word){
                    nextToken = nextToken.getVal(ctx, this);
                }
            }
            
            if(nextToken && nextToken.type == TypeEnum.op && (startDeep == 0 || idx > endPos.get(startDeep - 1))){
                if(startPos.len == 0 || line[startPos.last].type != TypeEnum.op){
                    startPos.add(idx);
                    push(nextToken);
                    push(nowToken.getVal(ctx, this));
                    endPos.add(idx);
                }else if(startPos.len == 0 || line[startPos.last].type == TypeEnum.op){
                    push(nowToken.getVal(ctx, this));
                    evalExp(ctx);
                    push(line[idx - 1]);
                    line[idx - 2] = nextToken;
                    startPos.add(idx - 2);
                    endPos.add(idx);
                }
                i += 1;
            }else{
                if(quoteList.len > 0){
                    if(quoteList.get(0)){
                        nowToken = nowToken.getVal(ctx, this);
                    }
                    quoteList.popFirst();
                }else{
                    nowToken = nowToken.getVal(ctx, this);
                }
                
                if(nowToken && nowToken.type == TypeEnum.err){
                    return nowToken;
                }else if(nowToken && nowToken.type == TypeEnum.op){
                    if(idx > startIdx){
                        startPos.add(idx - 1);
                        push(line[idx - 1]);
                        line[idx - 2] = nowToken;
                        endPos.add(idx);
                    }else{
                        result.type = TypeEnum.err;
                        result.str = "Illegal grammar!!!";
                        return result;
                    }
                }else if(nowToken && nowToken.type < TypeEnum.set_word){
                    push(nowToken);
                }else{
                    if(nowToken.type == TypeEnum.native){
                        if(nowToken.exec.quoteList && nowToken.exec.quoteList.len > 0){
                            quoteList.addAll(nowToken.exec.quoteList);
                        }
                    }else if(nowToken.type == TypeEnum.func){
                        if(nowToken.func.quoteList && nowToken.func.quoteList.len > 0){
                            quoteList.addAll(nowToken.func.quoteList);
                        }
                    }

                    startPos.add(idx);
                    endPos.add(idx + nowToken.explen() - 1);
                    push(nowToken);
                }
            }
            while(endPos.len > startDeep && idx == endPos.last + 1){
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
        bool isReturn = false;
        try{
            switch(line[startPos.last].type){
                case TypeEnum.set_word:
                    ctx.put(line[startPos.last].str, line[endPos.last]);
                    temp = line[endPos.last];
                    break;
                case TypeEnum.native:
                    temp = line[startPos.last].exec.run(this, ctx);
                    break;
                case TypeEnum.op:
                    temp = line[startPos.last].exec.run(this, ctx);
                    break;
                case TypeEnum.func:
                    temp = line[startPos.last].func.run(this, ctx);
                    break;
                default:
                    temp = new Token();
            }
        }catch(Exception e){
            if(e.msg == "break" || e.msg == "continue"){
                throw e;
            }else if(e.msg == "return"){
                isReturn = true;
                if(line[startPos.last].type == TypeEnum.func){
                    line[startPos.last] = line[idx - 1];
                    idx = startPos.last + 1;
                    startPos.pop;
                    endPos.pop;
                }else{
                    startPos.pop;
                    endPos.pop;
                    throw e;
                }
            }else{
                temp = new Token(TypeEnum.err);
                temp.str = "Illegal grammar!!!";
            }
        }finally{
            if(!isReturn){
                line[startPos.last] = temp;
                idx = startPos.last + 1;
                startPos.pop;
                endPos.pop;
            }
        }

    }    
}
