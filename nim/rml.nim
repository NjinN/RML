import token
import evalunit
import strtool
import native/nativeinit
import op/opinit
import strutils

GC_disable()

var exectuor = newEvalUnit()
initNative(exectuor.mainCtx)
initOp(exectuor.mainCtx)

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
    if not (answer.tp == TypeEnum.string and $answer.val.string == ""):
        echo(">> " & $answer.toStr)
    