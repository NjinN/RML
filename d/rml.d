module rml;

import std.stdio;
import std.uni;
import std.conv;

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
    evalStack.mainCtx = userCtx;

    char[] inp;
    while(true){

        write(">> ");
        
        char[] temp;
        stdin.readln(temp);
        inp ~= temp;
        if(inp[inp.length-2] == '~'){
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


