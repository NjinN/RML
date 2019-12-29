import ptrMath
export ptrMath

type
    List*[T] = object
        size*: uint
        endIdx*: int
        line*: ptr T

proc newList*[T](size: int = 16):ptr List[T]=
    var list = cast[ptr List[T]](alloc0(sizeof(List[T])))
    list.size = size.uint
    list.line = cast[ptr T](alloc0(size * sizeof(T)))
    list.endIdx = -1
    return list

proc freeList*[T](l: var List[T])=
    dealloc(l.line)

proc freeList*[T](l: ptr List[T])=
    dealloc(l.line)
    dealloc(l)

proc initList*[T](size: int = 16):List[T]=
    result.size = size.uint
    result.line = cast[ptr T](alloc0(size * sizeof(T)))
    result.endIdx = -1

proc clear*[T](l: var List[T])=
    zeroMem(l.line, sizeof(T) * (l.endIdx + 1))
    l.endIdx = -1

proc clear*[T](l: ptr List[T])=
    zeroMem(l.line, sizeof(T) * (l.endIdx + 1))
    l.endIdx = -1

proc clear*[T](l: var List[T], startIdx: int)=
    if startIdx <= l.endIdx:
        zeroMem(addr(l.line[startIdx]), sizeof(T) * (l.endIdx - startIdx + 1))
        l.endIdx = startIdx - 1

proc clear*[T](l: ptr List[T], startIdx: int)=
    if startIdx <= l.endIdx:
        zeroMem(l.line[startIdx], sizeof(T) * (l.endIdx - startIdx + 1))
        l.endIdx = startIdx - 1

proc upSize*[T](l: var List[T], newSize: uint)
proc upSize*[T](l: ptr List[T], newSize: uint)

proc `[]=`*[T](l: var List[T], idx: int, t: T)=
    if idx >= int(l.size-1):
        while l.size <= idx.uint:
            l.size *= 2
        l.upSize(l.size)
    l.line[idx] = t
    if idx > l.endIdx:
        l.endIdx = idx

proc `[]=`*[T](l: ptr List[T], idx: int, t: T)=
    if idx >= int(l.size-1):
        while l.size <= idx.uint:
            l.size *= 2
        l.upSize(l.size)
    l.line[idx] = t
    if idx > l.endIdx:
        l.endIdx = idx

proc `[]`*[T](l: List[T], idx: int):T=
    if idx < 0 or idx > l.size.int - 1:
        var e = new(OSError)
        e.msg = "Out of bound"
        raise e
    return l.line[idx]

proc `[]`*[T](l: ptr List[T], idx: int):T=
    if idx < 0 or idx > l.size.int - 1:
        var e = new(OSError)
        e.msg = "Out of bound"
        raise e
    return l.line[idx]

proc len*[T](l: List[T]):int=
    return l.size.int

proc len*[T](l: ptr List[T]):int=
    return l.size.int

proc high*[T](l: List[T]):int=
    return l.endIdx

proc high*[T](l: ptr List[T]):int=
    return l.endIdx

proc upSize*[T](l: var List[T], newSize: uint)=
    if newSize > l.size:
        l.line = cast[ptr T](realloc(l.line, newSize.int * sizeof(T)))
        l.size = newSize

proc upSize*[T](l: ptr List[T], newSize: uint)=
    if newSize > l.size:
        l.line = cast[ptr T](realloc(l.line, newSize.int * sizeof(T)))
        l.size = newSize

proc add*[T](l: var List[T], t: T)=
    if l.endIdx >= int(l.size - 1):
        l.upSize(l.size * 2)
    l[l.endIdx + 1] = t

proc add*[T](l: ptr List[T], t: T)=
    if l.endIdx >= int(l.size - 1):
        l.upSize(l.size * 2)
    l[l.endIdx + 1] = t

proc insert*[T](l: var List[T], t: T, idx: int)=
    if l.endIdx >= int(l.size - 1):
        l.upSize(l.size * 2)
    if idx >= int(l.size - 1):
        while idx >= int(l.size - 1):
            l.size = l.size * 2
        l.upSize(l.size)

    if l.endIdx >= idx:
        for i in countdown(high(l) + 1, idx + 1):
            copyMem(l.line + i, l.line + i - 1, sizeof(T))
        l.endIdx += 1
    else:
        l.endIdx = idx
    if isNull(t):
        zeroMem(l.line + idx, sizeof(T))
    else:
        copyMem(l.line + idx, cast[ptr T](unsafeAddr(t)), sizeof(T))

proc insert*[T](l: ptr List[T], t: T, idx: int)=
    if l.endIdx >= int(l.size - 1):
        l.upSize(l.size * 2)
    if idx >= int(l.size - 1):
        while idx >= int(l.size - 1):
            l.size = l.size * 2
        l.upSize(l.size)

    if l.endIdx >= idx:
        for i in countdown(high(l) + 1, idx + 1):
            copyMem(l.line + i, l.line + i - 1, sizeof(T))
        l.endIdx += 1
    else:
        l.endIdx = idx
    if isNull(t):
        zeroMem(l.line + idx, sizeof(T))
    else:
        copyMem(l.line + idx, cast[ptr T](unsafeAddr(t)), sizeof(T))


proc pop*[T](l: var List[T])=
    if l.endIdx >= 0:
        zeroMem(l.line + high(l), sizeof(T))
    l.endIdx -= 1   

proc pop*[T](l: ptr List[T])=
    if l.endIdx >= 0:
        zeroMem(l.line + high(l), sizeof(T))
    l.endIdx -= 1

proc drop*[T](l: var List[T], idx: int)=
    if idx == l.endIdx:
        pop(l)
    elif idx < l.endIdx:
        for i in idx..l.endIdx-1:
            l[i]=l[i + 1]
        pop(l)

proc drop*[T](l: ptr List[T], idx: int)=
    if idx == l.endIdx:
        pop(l)
    elif idx < l.endIdx:
        for i in idx..l.endIdx-1:
            l[i]=l[i + 1]
        pop(l)

proc first*[T](l: List[T]): T=
    return l.line[0]

proc first*[T](l: ptr List[T]): T=
    return l.line[0]

proc last*[T](l: List[T]): T=
    return l.line[l.endIdx]

proc last*[T](l: ptr List[T]): T=
    return l.line[l.endIdx]

iterator each*[T](l: List[T]): T=
    for i in 0..l.endIdx:
        yield l[i]

iterator each*[T](l: ptr List[T]): T=
    for i in 0..l.endIdx:
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
    echo(l[99])
    l.insert(199, 100)
    # l.pop()
    echo(l.first)
    echo(l.last)
    # echo GC_getStatistics()
    # freeList(l)
    # echo GC_getStatistics()