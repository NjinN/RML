import ptrMath
export ptrMath

type
    List*[T] = object
        size*: uint
        line*: ptr T

proc newList*[T](size: int = 16):ptr List[T]=
    var list = cast[ptr List[T]](alloc0(sizeof(List[T])))
    list.size = size.uint
    list.line = cast[ptr T](alloc0(size * sizeof(T)))
    return list

proc freeList*[T](l: ptr List[T])=
    dealloc(l.line)
    dealloc(l)

proc `[]=`*[T](l: ptr List[T], idx: int, t: T)=
    if idx >= int(l.size-1):
        while l.size <= idx.uint:
            l.size *= 2
        l.upSize(l.size)
    l.line[idx] = t

proc `[]`*[T](l: ptr List[T], idx: int):T=
    if idx < 0 or idx > l.size.int - 1:
        var e = new(OSError)
        e.msg = "Out of bound"
        raise e
    return l.line[idx]

proc len*[T](l: ptr List[T]):int=
    return l.size.int

proc high*[T](l: ptr List[T]):int=
    var endIdx = int(l.size - 1)
    for i in countdown(endIdx, 0):
        if not isNull(l[i]):
            return i
    return -1

proc upSize*[T](l: ptr List[T], newSize: uint)=
    if newSize > l.size:
        l.line = cast[ptr T](realloc(l.line, newSize.int * sizeof(T)))
        l.size = newSize

proc add*[T](l: ptr List[T], t: T)=
    if high(l) >= int(l.size - 1):
        l.upSize(l.size * 2)
    l[high(l) + 1] = t

proc insert*[T](l: ptr List[T], idx: int, t: T)=
    if high(l) >= int(l.size - 1):
        l.upSize(l.size * 2)
    if idx >= int(l.size - 1):
        while idx >= int(l.size - 1):
            l.size = l.size * 2
        l.upSize(l.size)

    if high(l) > idx:
        for i in countdown(high(l) + 1, idx + 1):
            copyMem(l.line + i, l.line + i - 1, sizeof(T))
    if isNull(t):
        zeroMem(l.line + idx, sizeof(T))
    else:
        copyMem(l.line + idx, cast[ptr T](unsafeAddr(t)), sizeof(T))

proc pop*[T](l: ptr List[T])=
    if high(l) >= 0:
        zeroMem(l.line + high(l), sizeof(T))

proc drop*[T](l: ptr List[T], idx: int)=
    if idx == high(l):
        pop(l)
    elif idx < high(l):
        for i in idx..high(l)-1:
            l[i]=l[i + 1]
        pop(l)

iterator each*[T](l: ptr List[T]): T=
    for i in 0..high(l):
        yield l[i]


when isMainModule:
    # echo GC_getStatistics()
    var l = newList[int](100)
    l[0]=0
    l[1]=1
    l[2]=2
    l[3]=3

    for item in l.each:
        echo(item)

    # echo GC_getStatistics()
    # echo(l[99])
    # l.insert(199, 100)
    # l.pop()
    # echo(high(l))
    # echo GC_getStatistics()
    # freeList(l)
    # echo GC_getStatistics()