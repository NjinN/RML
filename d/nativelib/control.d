module nativelib.control;

import common;
import typeenum;
import token;
import bindmap;
import evalstack;

Token iif(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.nil);

    if(args[2].type == TypeEnum.block){
        switch(args[1].type){
            case TypeEnum.logic:
                if(args[1].val.logic){
                    result = stack.eval(args[2].val.block, ctx);
                }
                return result;
            case TypeEnum.integer:
                if(args[1].val.integer != 0){
                    result = stack.eval(args[2].val.block, ctx);
                }
                return result;
            case TypeEnum.decimal:
                if(args[1].val.decimal != 0.0){
                    result = stack.eval(args[2].val.block, ctx);
                }
                return result;

            case TypeEnum.str:
                if(args[1].val.str != ""){
                    result = stack.eval(args[2].val.block, ctx);
                }
                return result;
            case TypeEnum.none:
                return result;
            default:
                return result;
        }

    }else{
        result.type = TypeEnum.err;
        result.val.str = "Type Mismatch";
        return result;
    }

}


Token either(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.nil);

    if(args[2].type == TypeEnum.block && args[3].type == TypeEnum.block){
        switch(args[1].type){
            case TypeEnum.logic:
                if(args[1].val.logic){
                    result = stack.eval(args[2].val.block, ctx);
                }else{
                    result = stack.eval(args[3].val.block, ctx);
                }
                return result;
            case TypeEnum.integer:
                if(args[1].val.integer != 0){
                    result = stack.eval(args[2].val.block, ctx);
                }else{
                    result = stack.eval(args[3].val.block, ctx);
                }
                return result;
            case TypeEnum.decimal:
                if(args[1].val.decimal != 0.0){
                    result = stack.eval(args[2].val.block, ctx);
                }else{
                    result = stack.eval(args[3].val.block, ctx);
                }
                return result;
            case TypeEnum.str:
                if(args[1].val.str != ""){
                    result = stack.eval(args[2].val.block, ctx);
                }else{
                    result = stack.eval(args[3].val.block, ctx);
                }
                return result;
            case TypeEnum.none:
                result = stack.eval(args[3].val.block, ctx);
                return result;
            default:
                result = stack.eval(args[3].val.block, ctx);
                return result;
        }

    }else{
        result.type = TypeEnum.err;
        result.val.str = "Type Mismatch";
        return result;
    }

}


Token loop(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.err);

    if(args[1].type == TypeEnum.integer && args[2].type == TypeEnum.block){
        for(int i=1; i <= args[1].val.integer; i++){
            try{
                result = stack.eval(args[2].val.block, ctx);
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
        result.val.str = "Type Mismatch";
    }
    return result;
}

Token repeat(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.err);

    if(args[1].type == TypeEnum.word && args[2].type == TypeEnum.integer && args[3].type == TypeEnum.block){
        BindMap c = new BindMap(ctx);
        Token countToken = new Token(TypeEnum.integer);
        countToken.val.integer = 1;
        c.putNow(args[1].val.str, countToken);
        while(c.get(args[1].val.str).val.integer <= args[2].val.integer){
            try{
                result = stack.eval(args[3].val.block, c);
            }catch(Exception e){
                if(e.msg == "continue"){
                    continue;
                }else if(e.msg == "break"){
                    break;
                }else{
                    throw new Exception("Runtime Error!");
                }
            }finally{
                Token temp = c.get(args[1].val.str);
                temp.val.integer += 1;
                c.putNow(args[1].val.str, temp);
            }
        }
    }else{
        result.val.str = "Type Mismatch";
    }
    return result;
}


Token ffor(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.err);

    if(args[1].type == TypeEnum.word && args[5].type == TypeEnum.block && (args[2].type == TypeEnum.integer || args[2].type == TypeEnum.decimal) && (args[3].type == TypeEnum.integer || args[3].type == TypeEnum.decimal) && (args[4].type == TypeEnum.integer || args[4].type == TypeEnum.decimal)){
        BindMap c = new BindMap(ctx);

        if(args[2].type == TypeEnum.integer && args[3].type == TypeEnum.integer && args[4].type == TypeEnum.integer){
            Token countToken = new Token(TypeEnum.integer);
            countToken.val.integer = args[2].val.integer;
            c.putNow(args[1].val.str, countToken);

            while(c.get(args[1].val.str).val.integer <= args[3].val.integer){
                try{
                    result = stack.eval(args[5].val.block, c);
                }catch(Exception e){
                    if(e.msg == "continue"){
                        continue;
                    }else if(e.msg == "break"){
                        break;
                    }else{
                        throw new Exception("Runtime Error!");
                    }
                }finally{
                    Token temp = c.get(args[1].val.str);
                    temp.val.integer += args[4].val.integer;
                    c.putNow(args[1].val.str, temp);
                }
            }
        }else{
            Token countToken = new Token(TypeEnum.decimal);
            if(args[2].type == TypeEnum.integer){
                countToken.val.decimal = cast(double)args[2].val.decimal;
            }else{
                countToken.val.decimal = args[2].val.decimal;
            }
            c.putNow(args[1].val.str, countToken);
            Token temp;
            if(args[3].type == TypeEnum.integer){
                while(c.get(args[1].val.str).val.decimal <= cast(double)args[3].val.integer){
                    try{
                        result = stack.eval(args[5].val.block, c);
                    }catch(Exception e){
                        if(e.msg == "continue"){
                            continue;
                        }else if(e.msg == "break"){
                            break;
                        }else{
                            throw new Exception("Runtime Error!");
                        }
                    }finally{
                        temp = c.get(args[1].val.str);
                        if(args[4].type == TypeEnum.integer){
                            temp.val.decimal += cast(double)args[4].val.integer;
                        }else{
                            temp.val.decimal += args[4].val.decimal;
                        }
                        c.putNow(args[1].val.str, temp);
                    }
                }
            }else{
                while(c.get(args[1].val.str).val.decimal <= args[3].val.decimal){
                    try{
                        result = stack.eval(args[5].val.block, c);
                    }catch(Exception e){
                        if(e.msg == "continue"){
                            continue;
                        }else if(e.msg == "break"){
                            break;
                        }else{
                            throw new Exception("Runtime Error!");
                        }
                    }finally{
                        temp = c.get(args[1].val.str);
                        if(args[4].type == TypeEnum.integer){
                            temp.val.decimal += cast(double)args[4].val.integer;
                        }else{
                            temp.val.decimal += args[4].val.decimal;
                        }
                        c.putNow(args[1].val.str, temp);
                    }
                }
            }

        }

    }else{
        result.val.str = "Type Mismatch";
    }
    return result;
}


Token wwhile(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.none);
    if(args[1].type == TypeEnum.block && args[2].type == TypeEnum.block){
        BindMap c = new BindMap(ctx);
        Token b = stack.eval(args[1].val.block, c);
        while(b.type == TypeEnum.logic && b.val.logic){
            if(result.type == TypeEnum.err){
                return result;
            }
            try{
                result = stack.eval(args[2].val.block, c);
            }catch(Exception e){
                if(e.msg == "continue"){
                    continue;
                }else if(e.msg == "break"){
                    break;
                }else{
                    throw new Exception("Runtime Error!");
                }
            }finally{
                b = stack.eval(args[1].val.block, c);
                if(b.type != TypeEnum.logic){
                    result.type = TypeEnum.err;
                    result.val.str = "Bad Logic Expression!";
                }
            }
        }
    }else{
        result.type = TypeEnum.err;
        result.val.str = "Type Mismatch";
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