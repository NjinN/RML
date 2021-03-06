module bindmap;
import std.stdio;

import token;


class BindMap {
    Token[string]  map;
    BindMap     father;

    this(){}
    this(BindMap f){
        map = f.map.dup;
        father = f;
    }

    Token get(string key){
        BindMap ctx = this;
        BindMap prev;
        
        Token nilTk = new Token(TypeEnum.none);
        Token tk = new Token(TypeEnum.none);
        if(ctx.map && ctx.map.length > 0){
            tk = ctx.map.get(key, nilTk);
        }
        if(tk.type != TypeEnum.nil && tk.type != TypeEnum.none){
            return tk;
        }else{
            while((tk.type == TypeEnum.nil || tk.type == TypeEnum.none) && ctx.father){
                prev = ctx;
                ctx = ctx.father;
                if(ctx.map && ctx.map.length > 0){
                    tk = ctx.map.get(key, nilTk);
                } 
            }
            if(tk.type != TypeEnum.nil && tk.type != TypeEnum.none && prev && !ctx.father){
                prev.map[key] = tk;
            }
        }
        return tk;
    }

    void putNow(string key, Token val){
        map[key] = val;
    }

    void put(string key, Token val){
        BindMap ctx = this;
        Token tk = null;
        bool inserted = false;
        if(ctx.map && ctx.map.length > 0){
            tk = ctx.map.get(key, null);
        }
        if(tk){
            map[key].copy(val);
            inserted = true;
        }else{
            while(!tk && ctx.father){
                ctx = ctx.father;
                if(ctx.map && ctx.map.length > 0){
                    tk = ctx.map.get(key, null);
                } 
                if(tk){
                    ctx.map[key].copy(val);
                    inserted = true;
                    break;
                }
            }
        }
        if(!inserted){
            map[key] = val.dup;
        }
    }

}
