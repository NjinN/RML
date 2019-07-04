import token
import strTool
import native/nativeInit
import op/opInit
import strutils
import evalStack


var execStack*: EvalStack
initEvalStack(execStack)
var libCtx* = initBindMap[Token](128)
initNative(libCtx)
initOp(libCtx)
var userCtx* = initBindMap[Token](128)
userCtx.father = addr(libCtx)

for i in 0..10:
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
    
    clearEvalStack(execStack)
    var answer = execStack.eval(inputStr, userCtx)
    if not (answer.tp == TypeEnum.null):
        echo(">> " & $answer.toStr)

    