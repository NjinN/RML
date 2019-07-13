module totoken;

import token;
import strtool;
import std.string;
import std.conv;
import arrlist;

Token toToken(string s){
    string str = trim(s);
    Token result = new Token;

    if(toLower(str) == "none"){
        result.type = TypeEnum.none;
        result.str = "none";
        return result;
    }

    if(toLower(str) == "true"){
        result.type = TypeEnum.logic;
        result.logic = true;
        return result;
    }

    if(toLower(str) == "false"){
        result.type = TypeEnum.logic;
        result.logic = false;
        return result;
    }

    if(str.length == 4 && str[0..2] == "#'" && str[3] == '\''){
        result.type = TypeEnum.cchar;
        result.cchar = str[2];
        return result;
    }

    if(str[0] == '"'){
        result.type = TypeEnum.str;
        result.str = str[1..str.length-1];
        return result;
    }

    if(str[0] == '['){
        result.type = TypeEnum.block;
        int endIdx = 0;
        for(int idx=cast(int)(str.length-1); idx>=0; idx--){
            if(str[idx] == ']'){
                endIdx = idx;
                break;
            }
        }
        result.block = toTokens(str[1..endIdx]);
        return result;
    }

    if(str[0] == '('){
        result.type = TypeEnum.paren;
        int endIdx = 0;
        for(int idx=cast(int)(str.length-1); idx>=0; idx--){
            if(str[idx] == ')'){
                endIdx = idx;
                break;
            }
        }
        result.block = toTokens(str[1..endIdx]);
        return result;
    }

    if(isNumberStr(str) == 0){
        result.type = TypeEnum.integer;
        result.integer = parse!int(str);
        return result;
    }

    if(isNumberStr(str) == 1){
        result.type = TypeEnum.decimal;
        result.decimal = parse!double(str);
        return result;
    }

    if(str[str.length-1] == ':'){
        result.type = TypeEnum.set_word;
        result.str = str[0..str.length-1];
        return result;
    }

    if(str[0] == '\''){
        result.type = TypeEnum.lit_word;
        result.str = str[1..str.length];
        return result;
    }

    result = new Token(TypeEnum.word);
    result.word.name = str;
    return result;
}

ArrList!Token toTokens(string str){
    auto result = new ArrList!Token(8);
    string[] strs = strCut(str);

    for(int i=0; i<strs.length; i++){
        result.add(toToken(strs[i]));
    }
    return result;
}



