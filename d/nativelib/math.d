module nativelib.math;

import common;
import token;
import bindmap;
import evalstack;
import arrlist;

Token add(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token();

    if(args[1].type == TypeEnum.integer){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.integer;
            result.integer = args[1].integer + args[2].integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.decimal = cast(double)args[1].integer + args[2].decimal;
            return result;
        }
    }else if(args[1].type == TypeEnum.decimal){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.decimal;
            result.decimal = args[1].decimal + cast(double)args[2].integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.decimal = args[1].decimal + args[2].decimal;
            return result;
        }
    }
    result.type = TypeEnum.err;
    result.str = "Type Mismatch";
    return result;
}

Token sub(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token();

    if(args[1].type == TypeEnum.integer){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.integer;
            result.integer = args[1].integer - args[2].integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.decimal = cast(double)args[1].integer - args[2].decimal;
            return result;
        }
    }else if(args[1].type == TypeEnum.decimal){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.decimal;
            result.decimal = args[1].decimal - cast(double)args[2].integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.decimal = args[1].decimal - args[2].decimal;
            return result;
        }
    }
    result.type = TypeEnum.err;
    result.str = "Type Mismatch";
    return result;
}

Token mul(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token();

    if(args[1].type == TypeEnum.integer){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.integer;
            result.integer = args[1].integer * args[2].integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.decimal = cast(double)args[1].integer * args[2].decimal;
            return result;
        }
    }else if(args[1].type == TypeEnum.decimal){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.decimal;
            result.decimal = args[1].decimal * cast(double)args[2].integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.decimal = args[1].decimal * args[2].decimal;
            return result;
        }
    }
    result.type = TypeEnum.err;
    result.str = "Type Mismatch";
    return result;
}

Token div(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token();

    if(args[1].type == TypeEnum.integer){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.decimal;
            result.decimal = cast(double)args[1].integer / cast(double)args[2].integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.decimal = cast(double)args[1].integer / args[2].decimal;
            return result;
        }
    }else if(args[1].type == TypeEnum.decimal){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.decimal;
            result.decimal = args[1].decimal / cast(double)args[2].integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.decimal = args[1].decimal * args[2].decimal;
            return result;
        }
    }
    result.type = TypeEnum.err;
    result.str = "Type Mismatch";
    return result;
}

Token mod(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    if(args[1].type == TypeEnum.integer && args[2].type == TypeEnum.integer){
        Token result = new Token(TypeEnum.integer);
        result.integer = args[1].integer % args[2].integer;
        return result;
    }
    return new Token(TypeEnum.err, "Type Mismatch");
}
