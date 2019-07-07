module nativelib.math;

import common;
import typeenum;
import token;
import bindmap;
import evalstack;

Token add(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token();

    if(args[1].type == TypeEnum.integer){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.integer;
            result.val.integer = args[1].val.integer + args[2].val.integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.val.decimal = cast(double)args[1].val.integer + args[2].val.decimal;
            return result;
        }
    }else if(args[1].type == TypeEnum.decimal){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.decimal;
            result.val.decimal = args[1].val.decimal + cast(double)args[2].val.integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.val.decimal = args[1].val.decimal + args[2].val.decimal;
            return result;
        }
    }
    result.type = TypeEnum.err;
    result.val.str = "Type Mismatch";
    return result;
}

Token sub(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token();

    if(args[1].type == TypeEnum.integer){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.integer;
            result.val.integer = args[1].val.integer - args[2].val.integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.val.decimal = cast(double)args[1].val.integer - args[2].val.decimal;
            return result;
        }
    }else if(args[1].type == TypeEnum.decimal){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.decimal;
            result.val.decimal = args[1].val.decimal - cast(double)args[2].val.integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.val.decimal = args[1].val.decimal - args[2].val.decimal;
            return result;
        }
    }
    result.type = TypeEnum.err;
    result.val.str = "Type Mismatch";
    return result;
}

Token mul(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token();

    if(args[1].type == TypeEnum.integer){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.integer;
            result.val.integer = args[1].val.integer * args[2].val.integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.val.decimal = cast(double)args[1].val.integer * args[2].val.decimal;
            return result;
        }
    }else if(args[1].type == TypeEnum.decimal){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.decimal;
            result.val.decimal = args[1].val.decimal * cast(double)args[2].val.integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.val.decimal = args[1].val.decimal * args[2].val.decimal;
            return result;
        }
    }
    result.type = TypeEnum.err;
    result.val.str = "Type Mismatch";
    return result;
}

Token div(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token();

    if(args[1].type == TypeEnum.integer){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.decimal;
            result.val.decimal = cast(double)args[1].val.integer * cast(double)args[2].val.integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.val.decimal = cast(double)args[1].val.integer * args[2].val.decimal;
            return result;
        }
    }else if(args[1].type == TypeEnum.decimal){
        if(args[2].type == TypeEnum.integer){
            result.type = TypeEnum.decimal;
            result.val.decimal = args[1].val.decimal * cast(double)args[2].val.integer;
            return result;
        }else if(args[2].type == TypeEnum.decimal){
            result.type = TypeEnum.decimal;
            result.val.decimal = args[1].val.decimal * args[2].val.decimal;
            return result;
        }
    }
    result.type = TypeEnum.err;
    result.val.str = "Type Mismatch";
    return result;
}
