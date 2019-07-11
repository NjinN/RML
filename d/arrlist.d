module arrlist;
import std.stdio;

class ArrList(T){
    uint    size;
    uint    endIdx;
    T[]     line;

    this(){}
    this(uint sz){
        size = sz;
        endIdx = 0;
        line.length = sz;
    }
    this(T[] arr){
        size = cast(uint)(arr.length + 1);
        endIdx = cast(uint)arr.length;
        line = arr;
        line.length = size;
    }
    
    void clear(){
        endIdx = 0;
    }
    uint room(){
        return size;
    }
    uint len(){
        return endIdx;
    }

    void resize(uint ns){
        size = ns;
        line.length = ns;
        if(ns < endIdx + 1){
            endIdx = ns - 1;
        }
    }

    void add(T v){
        if(size == 0 || endIdx >= size - 1){
            resize((size + 1) * 2);
        }
        line[endIdx] = v;
        endIdx += 1;
    }
    void pop(){
        if(endIdx > 0){
            endIdx -= 1;
        }
    }
    T get(uint idx){
        if(idx > endIdx){
            T rs;
            return rs;
        }
        return line[idx];
    }
    void put(uint idx, T v){
        while(idx >= size - 1){
            size *= 2;
        }
        if(size > line.length){
            line.length = size;
        }
        line[idx] = v;
        if(idx > endIdx){
            endIdx = idx;
        }
    }
    void insert(uint idx, T v){
        if(endIdx >= size - 1){
            resize(size * 2);
        }
        
        if(endIdx > idx){
            line[(idx+1)..(endIdx+1)] = line[idx..endIdx];
            line[idx] = v;
            endIdx += 1;
        }else{
            put(idx, v);
        }
    }
    T first(){
        if(endIdx>=0){
            return line[0];
        }
        T rs;
        return rs;
    }
    T last(){
        if(endIdx>0){
            return line[endIdx-1];
        }
        T rs;
        return rs;
    }
    void addAll(ArrList!T al){
        if(al.len > 0){
            for(int i=0; i< al.endIdx; i++){
                add(al.line[i]);
            }
        }
    }

    void addArr(T[] arr){
        foreach(item;arr){
            add(item);
        }
    }
    void popFirst(){
        if(endIdx > 0){
            line[0..endIdx-1] = line[1..endIdx].dup;
            endIdx -= 1;
        }
    }
    void echo(){
        writeln(line[0..endIdx]);
    }
}


// void main(string[] args) {
//     ArrList!int al = new ArrList!int(10);
//     for(int i=0; i<5000; i++){
//         // writeln(i);
//         al.add(i);
//     }
//     for(int i=0; i<5000; i++){
//         al.pop;
//     }
//     al.addArr([1, 2, 3]);
//     writeln(al.room);
//     writeln(al.len);
//     for(int i=0; i<al.len; i++){
//         writeln(al.get(i));
//     }
// }


