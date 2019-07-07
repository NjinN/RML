module nativelib.compare;

import common;
import typeenum;
import token;
import bindmap;
import evalstack;

Token eq(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.logic);
    result.val.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.val.logic = true;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.val.logic = args[1].val.logic == args[2].val.logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.integer == args[2].val.integer;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = cast(double)args[1].val.integer == args[2].val.decimal;
                    return result;
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.integer == cast(int)args[2].val.cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.decimal == cast(double)args[2].val.decimal;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = args[1].val.decimal == args[2].val.decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.cchar == args[2].val.cchar;
                    return result;
                case TypeEnum.integer:
                    result.val.logic = cast(int)args[1].val.cchar == args[2].val.integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.val.logic = args[1].val.str == args[2].val.str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.val.logic = args[1].val.str == args[2].val.str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}


Token ne(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.logic);
    result.val.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.val.logic = false;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.val.logic = args[1].val.logic != args[2].val.logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.integer != args[2].val.integer;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = cast(double)args[1].val.integer != args[2].val.decimal;
                    return result;
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.integer != cast(int)args[2].val.cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.decimal != cast(double)args[2].val.decimal;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = args[1].val.decimal != args[2].val.decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.cchar != args[2].val.cchar;
                    return result;
                case TypeEnum.integer:
                    result.val.logic = cast(int)args[1].val.cchar != args[2].val.integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.val.logic = args[1].val.str != args[2].val.str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.val.logic = args[1].val.str != args[2].val.str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}


Token gt(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.logic);
    result.val.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.val.logic = false;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.val.logic = args[1].val.logic > args[2].val.logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.integer > args[2].val.integer;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = cast(double)args[1].val.integer > args[2].val.decimal;
                    return result;
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.integer > cast(int)args[2].val.cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.decimal > cast(double)args[2].val.decimal;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = args[1].val.decimal > args[2].val.decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.cchar > args[2].val.cchar;
                    return result;
                case TypeEnum.integer:
                    result.val.logic = cast(int)args[1].val.cchar > args[2].val.integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.val.logic = args[1].val.str > args[2].val.str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.val.logic = args[1].val.str > args[2].val.str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}


Token lt(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.logic);
    result.val.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.val.logic = false;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.val.logic = args[1].val.logic < args[2].val.logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.integer < args[2].val.integer;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = cast(double)args[1].val.integer < args[2].val.decimal;
                    return result;
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.integer < cast(int)args[2].val.cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.decimal < cast(double)args[2].val.decimal;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = args[1].val.decimal < args[2].val.decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.cchar < args[2].val.cchar;
                    return result;
                case TypeEnum.integer:
                    result.val.logic = cast(int)args[1].val.cchar < args[2].val.integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.val.logic = args[1].val.str < args[2].val.str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.val.logic = args[1].val.str < args[2].val.str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}


Token gteq(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.logic);
    result.val.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.val.logic = true;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.val.logic = args[1].val.logic >= args[2].val.logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.integer >= args[2].val.integer;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = cast(double)args[1].val.integer >= args[2].val.decimal;
                    return result;
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.integer >= cast(int)args[2].val.cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.decimal >= cast(double)args[2].val.decimal;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = args[1].val.decimal >= args[2].val.decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.cchar >= args[2].val.cchar;
                    return result;
                case TypeEnum.integer:
                    result.val.logic = cast(int)args[1].val.cchar >= args[2].val.integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.val.logic = args[1].val.str >= args[2].val.str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.val.logic = args[1].val.str >= args[2].val.str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}


Token lteq(EvalStack stack, BindMap ctx){
    Token[] args = stack.line[last(stack.startPos)..(last(stack.endPos) + 1)];
    Token result = new Token(TypeEnum.logic);
    result.val.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.val.logic = true;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.val.logic = args[1].val.logic <= args[2].val.logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.integer <= args[2].val.integer;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = cast(double)args[1].val.integer <= args[2].val.decimal;
                    return result;
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.integer <= cast(int)args[2].val.cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.val.logic = args[1].val.decimal <= cast(double)args[2].val.decimal;
                    return result;
                case TypeEnum.decimal:
                    result.val.logic = args[1].val.decimal <= args[2].val.decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.val.logic = args[1].val.cchar <= args[2].val.cchar;
                    return result;
                case TypeEnum.integer:
                    result.val.logic = cast(int)args[1].val.cchar <= args[2].val.integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.val.logic = args[1].val.str <= args[2].val.str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.val.logic = args[1].val.str <= args[2].val.str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}



