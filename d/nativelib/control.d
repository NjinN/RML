module nativelib.control;

import common;
import token;
import bindmap;
import evalstack;
import arrlist;

Token iif(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.nil);

    if(args[2].type == TypeEnum.block){
        switch(args[1].type){
            case TypeEnum.logic:
                if(args[1].logic){
                    result = stack.eval(args[2].block, ctx);
                }
                return result;
            case TypeEnum.integer:
                if(args[1].integer != 0){
                    result = stack.eval(args[2].block, ctx);
                }
                return result;
            case TypeEnum.decimal:
                if(args[1].decimal != 0.0){
                    result = stack.eval(args[2].block, ctx);
                }
                return result;

            case TypeEnum.str:
                if(args[1].str != ""){
                    result = stack.eval(args[2].block, ctx);
                }
                return result;
            case TypeEnum.none:
                return result;
            default:
                return result;
        }

    }else{
        result.type = TypeEnum.err;
        result.str = "Type Mismatch";
        return result;
    }

}


Token either(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.nil);

    if(args[2].type == TypeEnum.block && args[3].type == TypeEnum.block){
        switch(args[1].type){
            case TypeEnum.logic:
                if(args[1].logic){
                    result = stack.eval(args[2].block, ctx);
                }else{
                    result = stack.eval(args[3].block, ctx);
                }
                return result;
            case TypeEnum.integer:
                if(args[1].integer != 0){
                    result = stack.eval(args[2].block, ctx);
                }else{
                    result = stack.eval(args[3].block, ctx);
                }
                return result;
            case TypeEnum.decimal:
                if(args[1].decimal != 0.0){
                    result = stack.eval(args[2].block, ctx);
                }else{
                    result = stack.eval(args[3].block, ctx);
                }
                return result;
            case TypeEnum.str:
                if(args[1].str != ""){
                    result = stack.eval(args[2].block, ctx);
                }else{
                    result = stack.eval(args[3].block, ctx);
                }
                return result;
            case TypeEnum.none:
                result = stack.eval(args[3].block, ctx);
                return result;
            default:
                result = stack.eval(args[3].block, ctx);
                return result;
        }

    }else{
        result.type = TypeEnum.err;
        result.str = "Type Mismatch";
        return result;
    }

}


Token loop(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.err);

    if(args[1].type == TypeEnum.integer && args[2].type == TypeEnum.block){
        for(int i=1; i <= args[1].integer; i++){
            try{
                result = stack.eval(args[2].block, ctx);
            }catch(Exception e){
                if(e.msg == "continue"){
                    continue;
                }else if(e.msg == "break"){
                    break;
                }else{
                    throw new Exception("Runtime Error!");
                }
            }
        }

    }else{
        result.type = TypeEnum.err;
        result.str = "Type Mismatch";
    }
    return result;
}


Token repeat(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.err);

    if(args[1].type == TypeEnum.word && args[2].type == TypeEnum.integer && args[3].type == TypeEnum.block){
        BindMap c = new BindMap(ctx);
        Token countToken = new Token(TypeEnum.integer);
        countToken.integer = 1;
        args[1].word.val = countToken;
        c.putNow(args[1].word.name, countToken);
        while(countToken.integer <= args[2].integer){
            try{
                result = stack.eval(args[3].block, c);
            }catch(Exception e){
                if(e.msg == "continue"){
                    continue;
                }else if(e.msg == "break"){
                    break;
                }else{
                    throw new Exception("Runtime Error!");
                }
            }finally{
                countToken.integer += 1;
            }
        }
    }else{
        result.str = "Type Mismatch";
    }
    return result;
}


Token ffor(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.err);

    if(args[1].type == TypeEnum.word && args[5].type == TypeEnum.block && (args[2].type == TypeEnum.integer || args[2].type == TypeEnum.decimal) && (args[3].type == TypeEnum.integer || args[3].type == TypeEnum.decimal) && (args[4].type == TypeEnum.integer || args[4].type == TypeEnum.decimal)){
        BindMap c = new BindMap(ctx);

        if(args[2].type == TypeEnum.integer && args[3].type == TypeEnum.integer && args[4].type == TypeEnum.integer){
            Token countToken = new Token(TypeEnum.integer);
            countToken.integer = args[2].integer;
            c.putNow(args[1].word.name, countToken);

            while(countToken.integer <= args[3].integer){
                try{
                    result = stack.eval(args[5].block, c);
                }catch(Exception e){
                    if(e.msg == "continue"){
                        continue;
                    }else if(e.msg == "break"){
                        break;
                    }else{
                        throw new Exception("Runtime Error!");
                    }
                }finally{
                    countToken.integer += args[4].integer;
                }
            }
        }else{
            Token countToken = new Token(TypeEnum.decimal);
            if(args[2].type == TypeEnum.integer){
                countToken.decimal = cast(double)args[2].decimal;
            }else{
                countToken.decimal = args[2].decimal;
            }
            c.putNow(args[1].word.name, countToken);
            Token temp;
            if(args[3].type == TypeEnum.integer){
                while(countToken.decimal <= cast(double)args[3].integer){
                    try{
                        result = stack.eval(args[5].block, c);
                    }catch(Exception e){
                        if(e.msg == "continue"){
                            continue;
                        }else if(e.msg == "break"){
                            break;
                        }else{
                            throw new Exception("Runtime Error!");
                        }
                    }finally{
                        if(args[4].type == TypeEnum.integer){
                            countToken.decimal += cast(double)args[4].integer;
                        }else{
                            countToken.decimal += args[4].decimal;
                        }
                    }
                }
            }else{
                while(countToken.decimal <= args[3].decimal){
                    try{
                        result = stack.eval(args[5].block, c);
                    }catch(Exception e){
                        if(e.msg == "continue"){
                            continue;
                        }else if(e.msg == "break"){
                            break;
                        }else{
                            throw new Exception("Runtime Error!");
                        }
                    }finally{
                        if(args[4].type == TypeEnum.integer){
                            countToken.decimal += cast(double)args[4].integer;
                        }else{
                            countToken.decimal += args[4].decimal;
                        }
                    }
                }
            }

        }

    }else{
        result.str = "Type Mismatch";
    }
    return result;
}


Token wwhile(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.none);
    if(args[1].type == TypeEnum.block && args[2].type == TypeEnum.block){
        BindMap c = new BindMap(ctx);
        Token b = stack.eval(args[1].block, c);
        while(b.type == TypeEnum.logic && b.logic){
            if(result.type == TypeEnum.err){
                return result;
            }
            try{
                result = stack.eval(args[2].block, c);
            }catch(Exception e){
                if(e.msg == "continue"){
                    continue;
                }else if(e.msg == "break"){
                    break;
                }else{
                    throw new Exception("Runtime Error!");
                }
            }finally{
                b = stack.eval(args[1].block, c);
                if(b.type != TypeEnum.logic){
                    result.type = TypeEnum.err;
                    result.str = "Bad Logic Expression!";
                }
            }
        }
    }else{
        result.type = TypeEnum.err;
        result.str = "Type Mismatch";
    }
    return result;
}

Token bbreak(EvalStack stack, BindMap ctx){
    throw new Exception("break");
    return new Token(TypeEnum.none);
}

Token ccontinue(EvalStack stack, BindMap ctx){
    throw new Exception("continue");
    return new Token(TypeEnum.none);
}