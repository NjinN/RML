import strutils

proc trim*(s: string):string=
    if s == "":
        return ""
    var startIdx = 0
    for i in 0..len(s)-1:
        if not isSpaceAscii(s[i]):
            startIdx = i
            break
    var endIdx = len(s)-1
    for i in countdown(len(s)-1, 0):
        if not isSpaceAscii(s[i]):
            endIdx = i
            break
    return s[startIdx..endIdx]


proc strCut*(s: string):seq[string]=
    result = @[]
    var str = trim(s)
    if str == "":
        return result
    
    var startIdx = -1
    var isParen = false
    var isStr = false
    var isBlock = false
    var pFloor = 0
    var bFloor = 0
    var nowChar: char

    for nowIdx in 0..len(str)-1:
        nowChar = str[nowIdx]
        if nowIdx == len(str)-1:
            if startIdx < 0 and not isSpaceAscii(nowChar):
                result.add($nowChar)
                break
            if startIdx >= 0:
                if isSpaceAscii(nowChar):
                    result.add(str[startIdx..(nowIdx-1)])
                else:
                    if not isStr and (not isBlock):
                        result.add(str[startIdx..nowIdx])
                        break
        
        if startIdx < 0 and not isSpaceAscii(nowChar):
            if nowChar == '"':
                isStr = true
            if nowChar == '(':
                isParen = true
                pFloor = 1
            if nowChar == '[':
                isBlock = true
                bFloor = 1
            startIdx = nowIdx
            continue

        if startIdx >= 0 and isSpaceAscii(nowChar) and not isStr and not isParen and not isBlock:
            result.add(str[startIdx..nowIdx-1])
            startIdx = -1
            continue

        if startIdx >= 0 and isStr:
            if nowChar == '"' and not (str[nowIdx-1..nowIdx] == "^\""):
                result.add(str[startIdx..nowIdx])
                isStr = false
                startIdx = -1
            continue
        
        if startIdx >= 0 and isParen:
            if nowChar == '(':
                pFloor += 1
            if nowChar == ')':
                pFloor -= 1
            if pFloor == 0:
                result.add(str[startIdx..nowIdx])
                isParen = false
                startIdx = -1
        
        if startIdx >= 0 and isBlock:
            if isStr:
                if nowChar == '"' and not (str[nowIdx-1..nowIdx] == "^\""):
                    isStr = false
            else:
                if nowChar == '[':
                    bFloor += 1
                if nowChar == ']':
                    bFloor -= 1
                if bFloor == 0:
                    result.add(str[startIdx..nowIdx])
                    isBlock = false
                    startIdx = -1
    return result





proc isNumberStr*(s: string):int=
    if len(s) == 0:
        return -1
    if s[0] != '-' and not isDigit($s[0]):
        return -1
    
    var dot = 0
    for idx in 0..len(s)-1:
        if not isDigit($s[idx]) and s[idx] != '.':
            return -1
        if s[idx] == '.':
            dot += 1
    
    return dot



when isMainModule:
    var strs = strCut("   123 \"this is a string  with space   ^\"  ([ 1 2 3 ] and tranChar) \"  ([ 123 456 \"anthor ^\" str \" 987 ] 456) ")

    for s in strs:
        echo(s)



