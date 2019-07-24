module strtool;
import std.stdio;
import std.uni;
import std.conv;



string trim(string s){
    if(s == ""){
        return "";
    }

    int startIdx = 0;
    for(int i=0; i<s.length; i++){
        if(!isWhite(s[i])){
            startIdx = i;
            break;
        }
    }
    int endIdx = 0;
    for(int j=(cast(int)s.length)-1; j>=0; j--){
        if(!isWhite(s[j])){
            endIdx = j;
            break;
        }
    }
    return s[startIdx..endIdx+1];
}

string[] strCut(string s){
    string[] result;
    string str = trim(s);
    if(str == ""){
        return result;
    }

    int startIdx = -1;
    bool isParen = false;
    bool isStr = false;
    bool isBlock = false;
    uint pFloor = 0;
    uint bFloor = 0;
    char nowChar;
    bool isInnerStr =false;

    for(int nowIdx=0; nowIdx<str.length; nowIdx++){
        nowChar = str[nowIdx];
        if(nowIdx == str.length-1){
            if(startIdx < 0 && !isWhite(nowChar)){
                result ~= text(nowChar);
                break;
            }
            if(startIdx >= 0){
                if(isWhite(nowChar)){
                    result ~= s[startIdx..nowIdx];
                }else{
                    if(!isStr && !isParen && !isBlock){
                        result ~= s[startIdx..nowIdx+1];
                        break;
                    }
                }
            }
        }

        if(startIdx < 0 && !isWhite(nowChar)){
            if(nowChar == '"'){
                isStr = true;
            }else if(nowChar == '('){
                isParen = true;
                pFloor = 1;
            }else if(nowChar == '['){
                isBlock = true;
                bFloor = 1;
            }
            startIdx = nowIdx;
            continue;
        } 

        if(startIdx >= 0 && isWhite(nowChar) && !isStr && !isParen && !isBlock){
            result ~= str[startIdx..nowIdx];
            startIdx = -1;
            continue;
        }

        if(startIdx >= 0 && isStr){
            if(nowChar == '"' && !(str[nowIdx-1..nowIdx+1] == "^\"")){
                result ~= str[startIdx..nowIdx+1];
                isStr = false;
                startIdx = -1;
                continue;
            }
        }

        if(startIdx >=0 && isParen){
            if(isInnerStr){
                if(nowChar == '"' && !(str[nowIdx-1..nowIdx+1] == "^\"")){
                    isInnerStr = false;
                }
            }else{
                if(nowChar == '"' && !(str[nowIdx-1..nowIdx+1] == "^\"")){
                    isInnerStr = true;
                }else if(nowChar == '('){
                    pFloor += 1;
                }else if(nowChar == ')'){
                    pFloor -= 1;
                }
                if(pFloor == 0){
                    result ~= str[startIdx..nowIdx+1];
                    isParen = false;
                    startIdx = -1;
                    continue;
                }
            }
        }

        if(startIdx >=0 && isBlock){
            if(isInnerStr){
                if(nowChar == '"' && !(str[nowIdx-1..nowIdx+1] == "^\"")){
                    isInnerStr = false;
                }
            }else{
                if(nowChar == '"' && !(str[nowIdx-1..nowIdx+1] == "^\"")){
                    isInnerStr = true;
                }else if(nowChar == '['){
                    bFloor += 1;
                }else if(nowChar == ']'){
                    bFloor -= 1;
                }
                if(bFloor == 0){
                    result ~= str[startIdx..nowIdx+1];
                    isBlock = false;
                    startIdx = -1;
                    continue;
                }
            }
        }
    }
    return result;
}


bool isNumber(char c){
    if(cast(int)c >= 48 && cast(int)c <= 57){
        return true;
    }
    return false;
}

int isNumberStr(string s){
    if(s.length == 0){
        return -1;
    }
    if((s[0] != '-' && !isNumber(s[0])) || s == "-") {
        return -1;
    }

    int dot = 0;
    for(int idx=1; idx<s.length; idx++){
        if(!isNumber(s[idx]) && s[idx] != '.'){
            return -1;
        }
        if(s[idx] == '.'){
            dot += 1;
        }
    }
    return dot;
}


// void main(string[] args) {
//     string[] strs = strCut("   123 \"this is a string  with space   ^\"  ([ 1 2 3 ] and tranChar) \"  ([ 123 456 \"anthor ^\" str \" 987 ] 456) ");
//     // string[] strs = strCut(" 1 \"this is a string\"");
//     for(int i=0; i<strs.length; i++){
//         writeln(strs[i]);
//     }
// }


