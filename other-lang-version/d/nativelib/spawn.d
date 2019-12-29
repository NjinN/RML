module nativelib.spawn;

import std.stdio;
import std.concurrency;
import core.thread;

import token;
import bindmap;
import evalstack;
import arrlist;



Token sspawn(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    if(args[1].type == TypeEnum.block && args[2].type == TypeEnum.block){
        for(int i=0; i<args[1].block.len; i++){
            if(args[1].block.get(i).type != TypeEnum.word){
                return new Token(TypeEnum.err, "Type Mismatch");
            }
        }
        for(int i=0; i<args[2].block.len; i++){
            if(args[2].block.get(i).type != TypeEnum.block){
                return new Token(TypeEnum.err, "Type Mismatch");
            }
        }
        for(int j=0; j<args[2].block.len; j++){
            EvalStack tempStack = new EvalStack();
            tempStack.libCtx = stack.libCtx;
            BindMap tempMainCtx = new BindMap(stack.libCtx);
            for(int x=0; x<args[1].block.len; x++){
                tempMainCtx.putNow(args[1].block.get(x).str, args[1].block.get(x).getVal(ctx, stack));
            }
            tempStack.mainCtx = tempMainCtx;
            auto tid = spawn(&spawnEval);
            tid.send(thisTid, cast(shared)tempStack);
            tid.send(thisTid, cast(shared)args[2].block.get(j).block);
        }
        thread_joinAll();
        return null;
    }
    return new Token(TypeEnum.err, "Type Mismatch");
}


void spawnEval(){
    auto stackMsg = receiveOnly!(Tid, shared EvalStack)();
    auto inpMsg = receiveOnly!(Tid, shared ArrList!Token)();

    EvalStack subStack = cast(EvalStack)stackMsg[1];
    BindMap subMainCtx = cast(BindMap)subStack.mainCtx;
    ArrList!Token subInp = cast(ArrList!Token)inpMsg[1];
    
    subStack.eval(subInp, subMainCtx);
}


Token mult(EvalStack stack, BindMap ctx){
    Token *args = &stack.line[stack.startPos.last];
    if(args[1].type == TypeEnum.block && args[2].type == TypeEnum.block){
        for(int i=0; i<args[1].block.len; i++){
            if(args[1].block.get(i).type != TypeEnum.word){
                return new Token(TypeEnum.err, "Type Mismatch");
            }
        }
        for(int i=0; i<args[2].block.len; i++){
            if(args[2].block.get(i).type != TypeEnum.block){
                return new Token(TypeEnum.err, "Type Mismatch");
            }
        }
        for(int j=0; j<args[2].block.len; j++){
            (){
                ArrList!Token inp = args[2].block.get(j).block;
                EvalStack tempStack = new EvalStack();
                tempStack.libCtx = stack.libCtx;
                BindMap tempMainCtx = new BindMap(stack.libCtx);
                for(int x=0; x<args[1].block.len; x++){
                    tempMainCtx.putNow(args[1].block.get(x).str, args[1].block.get(x).getVal(ctx, stack));
                }
                tempStack.mainCtx = tempMainCtx;
                new Thread((){ tempStack.eval(inp, tempMainCtx);}).start();
            }();
            
        }
        thread_joinAll();
        return null;
    }
    return new Token(TypeEnum.err, "Type Mismatch");
}


