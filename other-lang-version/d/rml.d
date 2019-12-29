module rml;

import std.stdio;
import std.uni;
import std.conv;
import std.path;
import std.file;

import token;
import bindmap;
import evalstack;
import strtool;

import nativelib.init;
import oplib.init;


void main(string[] args) {

    BindMap libCtx = new BindMap();
    initNative(libCtx);
    initOp(libCtx);
    
    BindMap userCtx = new BindMap();
    userCtx.father = libCtx;

    EvalStack evalStack = new EvalStack();
    evalStack.libCtx = libCtx;
    evalStack.mainCtx = userCtx;

    if(args.length > 1){
        string scriptPath = absolutePath(args[1], dirName(args[0]));
        if(exists(scriptPath)){
            try{
                string scriptText = readText(scriptPath);
                // writeln(scriptText);
                evalStack.init();
                Token answer = evalStack.eval(scriptText, userCtx);
                if(answer && answer.type != TypeEnum.nil){
                    writeln("== ", answer.toStr(), "\n");
                }else{
                    writeln("");
                }
            }catch(Exception e){
                writeln("读取文件失败！");
            }

        }else{
            writeln("系统找不到指定的文件！");
        }
    
    }

    char[] inp;
    while(true){
        write(">> ");
        
        char[] temp;
        stdin.readln(temp);
        inp ~= temp;
        if(inp.length >= 2 && inp[inp.length-2] == '~'){
            inp.length = inp.length - 2;
            continue;
        }
        string inpStr = trim(toLower(text(inp)));

        if(inpStr == ""){
            continue;
        }
        
        evalStack.init();
        Token answer = evalStack.eval(inpStr, userCtx);
        if(answer && answer.type != TypeEnum.nil){
            writeln("== ", answer.toStr(), "\n");
        }else{
            writeln("");
        }
        inp.length = 0;
    }

}


