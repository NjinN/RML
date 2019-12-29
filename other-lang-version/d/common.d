module common;

// import std.stdio;

T last(T)(T[] arr){
    if(arr.length == 0){
        T t;
        return t;
    }
    return arr[arr.length-1];
}


// void main(string[] args) {
//     int[] arr;
//     arr ~= 1;
//     arr ~= 2;
//     arr.length = 1;
//     arr ~= 99;
//     writeln(last(arr));
// }


