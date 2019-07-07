module totoken;

import typeenum;
import token;
import strtool;
import std.string;
import std.conv;

Token toToken(string s){
    string str = trim(s);
    Token result = new Token;

    if(toLower(str) == "none"){
        result.type = TypeEnum.none;
        result.val.str = "none";
        return result;
    }

    if(toLower(str) == "true"){
        result.type = TypeEnum.logic;
        result.val.logic = true;
        return result;
    }

    if(toLower(str) == "false"){
        result.type = TypeEnum.logic;
        result.val.logic = false;
        return result;
    }

    if(str.length == 4 && str[0..2] == "#'" && str[3] == '\''){
        result.type = TypeEnum.cchar;
        result.val.cchar = str[2];
        return result;
    }

    if(str[0] == '"'){
        result.type = TypeEnum.str;
        result.val.str = str[1..str.length-1];
        return result;
    }

    if(str[0] == '['){
        result.type = TypeEnum.block;
        int endIdx = 0;
        for(int idx=str.length-1; idx>=0; idx--){
            if(str[idx] == ']'){
                endIdx = idx;
                break;
            }
        }
        result.val.block = toTokens(str[1..endIdx]);
        return result;
    }

    if(str[0] == '('){
        result.type = TypeEnum.paren;
        int endIdx = 0;
        for(int idx=str.length-1; idx>=0; idx--){
            if(str[idx] == ')'){
                endIdx = idx;
                break;
            }
        }
        result.val.block = toTokens(str[1..endIdx]);
        return result;
    }

    if(isNumberStr(str) == 0){
        result.type = TypeEnum.integer;
        result.val.integer = parse!int(str);
        return result;
    }

    if(isNumberStr(str) == 1){
        result.type = TypeEnum.decimal;
        result.val.decimal = parse!double(str);
        return result;
    }

    if(str[str.length-1] == ':'){
        result.type = TypeEnum.set_word;
        result.val.str = str[0..str.length-1];
        return result;
    }

    if(str[0] == '\''){
        result.type = TypeEnum.lit_word;
        result.val.str = str[1..str.length];
        return result;
    }

    result.type = TypeEnum.word;
    result.val.str = str;
    return result;
}

Token[] toTokens(string str){
    Token[] result;
    string[] strs = strCut(str);

    for(int i=0; i<strs.length; i++){
        result ~= toToken(strs[i]);
    }
    return result;
}



