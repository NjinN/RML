import token
import strtool
import strutils

proc toTokens*(s: string):seq[ref Token]

proc toToken*(s: string):ref Token=
    result = new Token
    var str = trim(s)
    if str == "":
        result.tp = TypeEnum.none
        result.val.string = "none"
        result.explen = 1
        return result

    if str[0] == '"':
        result.tp = TypeEnum.string
        result.val.string = str[1..len(str)-2]
        result.explen = 1
        return result

    if str[0] == '[':
        result.tp = TypeEnum.list 
        var endIdx = 0
        for idx in countdown(len(str)-1, 1):
            if str[idx] == ']':
                endIdx = idx-1
                break
        result.val.list = toTokens(str[1..endIdx])
        result.explen = 1
        return result

    if isNumberStr(str) == 0:
        result.tp = TypeEnum.integer
        result.val.integer = parseInt(str).int32
        result.explen = 1
        return result

    if isNumberStr(str) == 1:
        result.tp = TypeEnum.float
        result.val.float = parseFloat(str)
        result.explen = 1
        return result

    if str[len(str)-1] == ':':
        result.tp = TypeEnum.set_word
        result.val.string = str[0..len(str)-2]
        result.explen = 2
        return result

    if $str[0] == "'":
        result.tp = TypeEnum.lit_word
        result.val.string = str[1..len(str)-1]
        result.explen = 1
        return result

    
    result.tp = TypeEnum.word
    result.val.string = str
    result.explen = 1 

    return result


proc toTokens*(s: string):seq[ref Token]=
    result = newSeq[ref Token]()
    var strs = strCut(s)

    for str in strs:
        result.add(toToken(str))
    return result



when isMainModule:
    var t = toTokens("  123 \"this is a string  with space   ^\"  [1 2 3] and tranChar \"  [ [1 2 3] 123 456 \"anthor ^\" str \" 987 ] 456  -1.23 'word ")
    # var t = toTokens("[1 2 3")
    for item in t:
        print(item)




