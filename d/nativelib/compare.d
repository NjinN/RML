module nativelib.compare;

import common;
import token;
import bindmap;
import evalstack;
import arrlist;

Token eq(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.logic);
    result.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.logic = true;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.logic = args[1].logic == args[2].logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].integer == args[2].integer;
                    return result;
                case TypeEnum.decimal:
                    result.logic = cast(double)args[1].integer == args[2].decimal;
                    return result;
                case TypeEnum.cchar:
                    result.logic = args[1].integer == cast(int)args[2].cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].decimal == cast(double)args[2].decimal;
                    return result;
                case TypeEnum.decimal:
                    result.logic = args[1].decimal == args[2].decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.logic = args[1].cchar == args[2].cchar;
                    return result;
                case TypeEnum.integer:
                    result.logic = cast(int)args[1].cchar == args[2].integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.logic = args[1].str == args[2].str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.logic = args[1].str == args[2].str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}


Token ne(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.logic);
    result.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.logic = false;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.logic = args[1].logic != args[2].logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].integer != args[2].integer;
                    return result;
                case TypeEnum.decimal:
                    result.logic = cast(double)args[1].integer != args[2].decimal;
                    return result;
                case TypeEnum.cchar:
                    result.logic = args[1].integer != cast(int)args[2].cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].decimal != cast(double)args[2].decimal;
                    return result;
                case TypeEnum.decimal:
                    result.logic = args[1].decimal != args[2].decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.logic = args[1].cchar != args[2].cchar;
                    return result;
                case TypeEnum.integer:
                    result.logic = cast(int)args[1].cchar != args[2].integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.logic = args[1].str != args[2].str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.logic = args[1].str != args[2].str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}


Token gt(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.logic);
    result.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.logic = false;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.logic = args[1].logic > args[2].logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].integer > args[2].integer;
                    return result;
                case TypeEnum.decimal:
                    result.logic = cast(double)args[1].integer > args[2].decimal;
                    return result;
                case TypeEnum.cchar:
                    result.logic = args[1].integer > cast(int)args[2].cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].decimal > cast(double)args[2].decimal;
                    return result;
                case TypeEnum.decimal:
                    result.logic = args[1].decimal > args[2].decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.logic = args[1].cchar > args[2].cchar;
                    return result;
                case TypeEnum.integer:
                    result.logic = cast(int)args[1].cchar > args[2].integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.logic = args[1].str > args[2].str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.logic = args[1].str > args[2].str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}


Token lt(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.logic);
    result.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.logic = false;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.logic = args[1].logic < args[2].logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].integer < args[2].integer;
                    return result;
                case TypeEnum.decimal:
                    result.logic = cast(double)args[1].integer < args[2].decimal;
                    return result;
                case TypeEnum.cchar:
                    result.logic = args[1].integer < cast(int)args[2].cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].decimal < cast(double)args[2].decimal;
                    return result;
                case TypeEnum.decimal:
                    result.logic = args[1].decimal < args[2].decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.logic = args[1].cchar < args[2].cchar;
                    return result;
                case TypeEnum.integer:
                    result.logic = cast(int)args[1].cchar < args[2].integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.logic = args[1].str < args[2].str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.logic = args[1].str < args[2].str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}


Token gteq(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.logic);
    result.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.logic = true;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.logic = args[1].logic >= args[2].logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].integer >= args[2].integer;
                    return result;
                case TypeEnum.decimal:
                    result.logic = cast(double)args[1].integer >= args[2].decimal;
                    return result;
                case TypeEnum.cchar:
                    result.logic = args[1].integer >= cast(int)args[2].cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].decimal >= cast(double)args[2].decimal;
                    return result;
                case TypeEnum.decimal:
                    result.logic = args[1].decimal >= args[2].decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.logic = args[1].cchar >= args[2].cchar;
                    return result;
                case TypeEnum.integer:
                    result.logic = cast(int)args[1].cchar >= args[2].integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.logic = args[1].str >= args[2].str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.logic = args[1].str >= args[2].str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}


Token lteq(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    Token result = new Token(TypeEnum.logic);
    result.logic = false;

    switch(args[1].type){
        case TypeEnum.none:
            if(args[2].type == TypeEnum.none){
                result.logic = true;
                return result;
            }
            break;
        case TypeEnum.logic:
            if(args[2].type == TypeEnum.logic){
                result.logic = args[1].logic <= args[2].logic;
                return result;
            }
            break;
        case TypeEnum.integer:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].integer <= args[2].integer;
                    return result;
                case TypeEnum.decimal:
                    result.logic = cast(double)args[1].integer <= args[2].decimal;
                    return result;
                case TypeEnum.cchar:
                    result.logic = args[1].integer <= cast(int)args[2].cchar;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.decimal:
            switch(args[2].type){
                case TypeEnum.integer:
                    result.logic = args[1].decimal <= cast(double)args[2].decimal;
                    return result;
                case TypeEnum.decimal:
                    result.logic = args[1].decimal <= args[2].decimal;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.cchar:
            switch(args[2].type){
                case TypeEnum.cchar:
                    result.logic = args[1].cchar <= args[2].cchar;
                    return result;
                case TypeEnum.integer:
                    result.logic = cast(int)args[1].cchar <= args[2].integer;
                    return result;
                default:
                    return result;
            }
            break;
        case TypeEnum.str:
            if(args[2].type == TypeEnum.str){
                result.logic = args[1].str <= args[2].str;
                return result;
            }
            break;
        case TypeEnum.word:
            if(args[2].type == TypeEnum.word){
                result.logic = args[1].str <= args[2].str;
                return result;
            }
            break;
        default:
            return result;
    }
    return result;
}



