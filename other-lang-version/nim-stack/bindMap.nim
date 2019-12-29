import listType
import strutils
# import times

type
    HashBucket*[T] = object
        key*: cstring
        val*: T
        next*: ptr HashBucket[T]

    BindMap*[T] = object
        size*: uint
        len*: uint
        line*: List[ptr HashBucket[T]]
        father*: ptr BindMap[T]


proc hashCode*(s: cstring):uint=
    var seed = 131.uint
    result = 0
    for i in 0..len(s)-1:
        result = (result * seed) + ord(s[i]).uint

proc newHashBucket*[T]():ptr HashBucket[T]=
    result = cast[ptr HashBucket[T]](alloc0(sizeof(HashBucket[T])))

proc freeHashBucket*[T](bt: ptr HashBucket[T])=
    if not isNil(bt):
        if not isNil(bt.next):
            freeHashBucket(bt.next)
        dealloc(bt)
        


proc newBindMap*[T](size: int = 16):ptr BindMap[T]=
    result = cast[ptr BindMap[T]](alloc0(sizeof(BindMap[T])))
    result.size = size.uint
    result.len = 0
    result.line = initList[ptr HashBucket[T]](size)

proc initBindMap*[T](size: int = 16):BindMap[T]=
    result.size = size.uint
    result.len = 0
    result.line = initList[ptr HashBucket[T]](size)

proc freeBindMap*[T](m: var BindMap[T])=
    if not isNull(m):
        for item in m.line.each:
            freeHashBucket(item)
        freeList(m.line)

proc freeBindMap*[T](m: ptr BindMap[T])=
    if not isNil(m):
        for item in m.line.each:
            freeHashBucket(item)
        freeList(m.line)
        dealloc(m)

proc upSize*[T](m: var BindMap[T], newSize: int)
proc upSize*[T](m: ptr BindMap[T], newSize: int)

proc put*[T](m: var BindMap[T], k: cstring, v: T)=
    var idx = hashCode(k) mod m.size
    var bt = m.line[idx.int]
    if isNull(bt):
        var newBt = newHashBucket[T]()
        newBt.key = k
        newBt.val = v
        m.line[idx.int] = newBt
        m.len += 1.uint
    else:
        while (not isNil(bt.next)) and (bt.key != k):
            bt = bt.next
        if bt.key == k:
            bt.val = v
        else:
            var newBt = newHashBucket[T]()
            newBt.key = k
            newBt.val = v
            bt.next = newBt
            m.len += 1.uint
    if m.len > uint(m.size.int / 4 * 3 ):
        m.upSize(2 * m.size.int)

proc put*[T](m: ptr BindMap[T], k: cstring, v: T)=
    var idx = hashCode(k) mod m.size
    var bt = m.line[idx.int]
    if isNull(bt):
        var newBt = newHashBucket[T]()
        newBt.key = k
        newBt.val = v
        m.line[idx.int] = newBt
        m.len += 1.uint
    else:
        while (not isNil(bt.next)) and (bt.key != k):
            bt = bt.next
        if bt.key == k:
            bt.val = v
        else:
            var newBt = newHashBucket[T]()
            newBt.key = k
            newBt.val = v
            bt.next = newBt
            m.len += 1.uint
    if m.len > uint(m.size.int / 4 * 3 ):
        m.upSize(2 * m.size.int)

proc get*[T](m: ptr BindMap[T], k: cstring):T

proc get*[T](m: var BindMap[T], k: cstring):T=
    try:
        var idx = hashCode(k) mod m.size
        var bt = m.line[idx.int]
        if not isNull(bt):
            while not isNull(bt.key) and bt.key != k and not isNil(bt.next):
                bt = bt.next
            if bt.key == k:
                result = bt.val
                return result
        if isNull(result) and not isNil(m.father):
            result = m.father.get(k)
            if not isNull(result):
                m[k] = result
    except:
        var rs: T
        return rs

proc get*[T](m: ptr BindMap[T], k: cstring):T=
    try:
        var idx = hashCode(k) mod m.size
        var bt = m.line[idx.int]
        if not isNull(bt):
            while not isNull(bt.key) and bt.key != k and not isNil(bt.next):
                bt = bt.next
            if bt.key == k:
                result = bt.val
                return result
        if isNull(result) and not isNil(m.father):
            result = m.father.get(k)
            if not isNull(result):
                m[k] = result
    except:
        var rs: T
        return rs

proc `[]=`*[T](m: var BindMap[T], k: cstring, v: T)=
    m.put(k, v)

proc `[]=`*[T](m: ptr BindMap[T], k: cstring, v: T)=
    m.put(k, v)

proc `[]`*[T](m: var BindMap[T], k: cstring):T=
    result = m.get(k)

proc `[]`*[T](m: var ptr BindMap[T], k: cstring):T=
    result = m.get(k)

proc upSize*[T](m: var BindMap[T], newSize: int)=
    var oldLine = m.line
    m.size = newSize.uint
    m.len = 0
    m.line = initList[ptr HashBucket[T]](newSize)
    for i in 0..high(oldLine):
        var bt = oldLine[i]
        while not isNull(bt) and not isNull(bt.key):
            m[bt.key] = bt.val
            if isNil(bt.next):
                break
            bt = bt.next
    for item in oldLine.each:
        freeHashBucket(item)
    freeList(addr(oldLine))

proc upSize*[T](m: ptr BindMap[T], newSize: int)=
    var oldLine = m.line
    m.size = newSize.uint
    m.len = 0
    m.line = initList[ptr HashBucket[T]](newSize)
    for i in 0..high(oldLine):
        var bt = oldLine[i]
        while not isNull(bt) and not isNull(bt.key):
            m[bt.key] = bt.val
            if isNil(bt.next):
                break
            bt = bt.next
    for item in oldLine.each:
        freeHashBucket(item)
    freeList(addr(oldLine))

proc size*[T](m: BindMap[T]):int=
    return m.size.int  

proc size*[T](m: ptr BindMap[T]):int=
    return m.size.int  

proc len*[T](m: BindMap[T]):int=
    return m.len.int 

proc len*[T](m: ptr BindMap[T]):int=
    return m.len.int  

iterator each*[T](m: BindMap[T]):T=
    for item in m.line.each:
        if not isNull(item):
            yield item.val
            var temp = item
            while not isNil(temp.next):
                temp = temp.next
                yield temp.val

iterator each*[T](m: ptr BindMap[T]):T=
    for item in m.line.each:
        if not isNull(item):
            yield item.val
            var temp = item
            while not isNil(temp.next):
                temp = temp.next[0]
                yield temp.val


when isMainModule:
    # echo(hashCode("Hello World"))
    # echo(GC_getStatistics())
    # var map = initBindMap[int](1)
    # map["123"] = 666

    # echo(map["123"])
    # echo(GC_getStatistics())
    # echo(size(map))
    # # echo(map["123"])
    # # freeBindMap(map)
    # echo(GC_getStatistics())

    # for i in 0..10:
    #     map[cstring($i)]=i

    # for item in map.each:
    #     echo item

    var map = initBindMap[int](100)
    # echo(GC_getStatistics())
    echo getFreeMem()
    for i in 0..100000:
        map[cstring($i)] = i
    # echo(GC_getStatistics())
    echo getFreeMem()
    freeBindMap(map)
    # deallocHeap(false)
    # echo(GC_getStatistics())
    echo getFreeMem()