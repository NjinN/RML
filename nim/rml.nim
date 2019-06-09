import evalunit
import strtool
import native/nativeinit
import op/opinit

var exectuor = newEvalUnit()
initNative(exectuor.mainCtx)
initOp(exectuor.mainCtx)

while true:
    write(stdout, ">> ")
    flushFile(stdout)
    var inputStr = readLine(stdin)
    inputStr = trim(inputStr)
    if(inputStr == "quit" or inputStr == "q" or inputStr == "Quit" or inputStr == "QUIT" or inputStr == "Q"):
        quit(0)

    if(inputStr == ""):
        continue

    var answer = exectuor.eval(inputStr)
    if answer != "":
        echo(">> " & answer)
    