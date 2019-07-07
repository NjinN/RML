module bindmap;
import std.stdio;
import typeenum;
import token;


class BindMap {
    Token[string]  map;
    BindMap     father;

    this(){}
    this(BindMap f){
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
            if(tk.type != TypeEnum.nil && tk.type != TypeEnum.none && prev){
                prev.map[key] = tk;
            }
        }
        return tk;
    }

    void put(string key, TK val){
        this.map[key] = val;
    }

}

