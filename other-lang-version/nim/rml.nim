import token
import evalUnit
import strTool
import native/nativeInit
import op/opInit
import strutils

var markSet = newSet[ptr Token](1024)
var nowUnit: EvalUnit

var libCtx = newBindMap[ptr Token](128)
initNative(libCtx)
initOp(libCtx)
var exectuor = newEvalUnit(libCtx)

GC_disable()

while true:
    write(stdout, ">> ")
    flushFile(stdout)
    var inputStr = readLine(stdin)
    inputStr = trim(inputStr)
    if(toLowerAscii(inputStr) == "quit" or toLowerAscii(inputStr) == "q"):
        quit(0)

    if(inputStr == ""):
        continue

    var answer = exectuor.eval(inputStr)
    if not (answer.tp == TypeEnum.string and answer.val.string == ""):
        echo(">> " & $answer.toStr)
    