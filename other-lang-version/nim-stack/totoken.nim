import token
import strtool
import strutils

proc toTokens*(s: string):ptr List[Token]

proc toToken*(s: string):Token=
    var str = trim(s)

    if str == "none":
        result.tp = TypeEnum.none
        result.val.string = "none"
        return result

    if toLowerAscii(str) == "true":
        result.tp = TypeEnum.logic
        result.val.logic = true
        return result

    if toLowerAscii(str) == "false":
        result.tp = TypeEnum.logic
        result.val.logic = false
        return result
    
    if  len(str) == 4 and (str[0..1] == "#'") and (str[3] == '\''):
        result.tp = TypeEnum.char
        result.val.char = str[2]
        return result

    if str[0] == '"':
        result.tp = TypeEnum.string
        result.val.string = cstring(str[1..len(str)-2])
        return result

    if str[0] == '[':
        result.tp = TypeEnum.list 
        var endIdx = 0
        for idx in countdown(len(str)-1, 1):
            if str[idx] == ']':
                endIdx = idx-1
                break
        result.val.list = toTokens(str[1..endIdx])
        return result
    
    if str[0] == '(':
        result.tp = TypeEnum.paren 
        var endIdx = 0
        for idx in countdown(len(str)-1, 1):
            if str[idx] == ')':
                endIdx = idx-1
                break
        result.val.list = toTokens(str[1..endIdx])
        return result

    if isNumberStr(str) == 0:
        result.tp = TypeEnum.integer
        result.val.integer = parseInt(str).int32
        return result

    if isNumberStr(str) == 1:
        result.tp = TypeEnum.decimal
        result.val.decimal = parseFloat(str)
        return result

    if str[len(str)-1] == ':':
        result.tp = TypeEnum.set_word
        result.val.string = cstring(str[0..len(str)-2])
        return result

    if $str[0] == "'":
        result.tp = TypeEnum.lit_word
        result.val.string = cstring(str[1..len(str)-1])
        return result

    
    result.tp = TypeEnum.word
    result.val.string = cstring(str) 
    return result


proc toTokens*(s: string):ptr List[Token]=
    result = newList[Token]()
    var strs = strCut(s)

    for str in strs:
        result.add(toToken(str))
    return result



when isMainModule:
    var t = toTokens("  123 \"this is a string  with space   ^\"  [1 2 3] and tranChar \"  [ [1 2 3] 123 456 \"anthor ^\" str \" 987 ] 456  -1.23 'word ")
    # var t = toTokens("[1 2 3")
    for i in 0..high(t):
        print(t[i])




