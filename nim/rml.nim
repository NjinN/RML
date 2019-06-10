import evalunit
import strtool
import native/nativeinit
import op/opinit
import strutils

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
    if $answer.val.string != "":
        echo(">> " & $answer.val.string)
    