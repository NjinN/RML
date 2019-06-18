template `+`*[T](p: ptr T, off: int): ptr T =
    cast[ptr type(p[])](cast[ByteAddress](p) +% off * sizeof(p[]))
    
template `+=`*[T](p: ptr T, off: int) =
    p = p + off

template `-`*[T](p: ptr T, off: int): ptr T =
    cast[ptr type(p[])](cast[ByteAddress](p) -% off * sizeof(p[]))

template `-=`*[T](p: ptr T, off: int) =
    p = p - off

template `[]`*[T](p: ptr T, off: int): T =
    (p + off)[]

template `[]=`*[T](p: ptr T, off: int, val: T) =
    (p + off)[] = val

proc isNull*[T](t: ptr T):bool=
    return isNil(t)
    
proc isNull*[T](t: T):bool=
    var p: T
    return t == p
