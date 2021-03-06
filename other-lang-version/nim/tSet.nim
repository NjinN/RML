import ptrMath
import listType

type
    SetBucket*[T] = object
        val*: T
        next*: ptr SetBucket[T]

    Set*[T] = object
        size*: uint
        len*: uint
        line*: ptr List[ptr SetBucket[T]]


proc hashCode(s: cstring):uint=
    var seed = 131.uint
    result = 0
    for i in 0..len(s)-1:
        result = (result * seed) + ord(s[i]).uint

proc newSetBucket*[T](t: T):ptr SetBucket[T]=
    result = cast[ptr SetBucket[T]](alloc0(sizeof(SetBucket[T])))
    result.val = t

proc freeSetBucket*[T](bt: ptr SetBucket[T])=
    if not isNil(bt):
        if not isNil(bt.next):
            freeSetBucket(bt.next)
        dealloc(bt)

proc newSet*[T](size: int = 16):ptr Set[T]=
    result = cast[ptr Set[T]](alloc0(sizeof(Set[T])))
    result.size = size.uint
    result.len = 0
    result.line = newList[ptr SetBucket[T]](size)

proc freeSet*[T](s: ptr Set[T])=
    if not isNil(s):
        for i in 0..high(s.line):
            freeSetBucket(s.line[i.int])
        freeList(s.line)
        dealloc(s)
   
proc upSize*[T](s: ptr Set[T], newSize: int)


proc add*[T](s: ptr Set[ptr T], t: ptr T):int{.discardable.}=
    var idx = hashCode(repr(cast[uint](t))) mod s.size
    var bt = s.line[idx.int]
    if not isNil(bt) and bt.val == t:
        result = 0
    elif isNil(bt):
        var bucket = newSetBucket[ptr T](t)
        s.line[idx.int] = bucket
        s.len += 1
        result = 1
    else:
        while not isNil(bt.next):
            bt = bt.next
        if bt.val == t:
            result = 0
        else:
            var bucket = newSetBucket[ptr T](t)
            bucket.next = s.line[idx.int]
            s.line[idx.int] = bucket
            s.len += 1
            result = 1
    if s.len > uint(int(s.size) / 4 * 3):
        s.upSize(s.size.int * 2)


proc add*[T](s: ptr Set[T], t: T):int{.discardable.}=
    var idx = hashCode(repr(t)) mod s.size
    var bt = s.line[idx.int]
    if not isNil(bt) and bt.val == t:
        result = 0
    elif isNil(bt):
        var bucket = newSetBucket[T](t)
        s.line[idx.int] = bucket
        s.len += 1
        result = 1
    else:
        while not isNil(bt.next):
            bt = bt.next
        if bt.val == t:
            result = 0
        else:
            var bucket = newSetBucket[T](t)
            bucket.next = s.line[idx.int]
            s.line[idx.int] = bucket
            s.len += 1
            result = 1
    if s.len > uint(int(s.size) / 4 * 3):
        s.upSize(s.size.int * 2)

proc has*[T](s: ptr Set[ptr T], t: ptr T):bool=
    var idx = hashCode(repr(cast[uint](t))) mod s.size
    var bt = s.line[idx.int]
    result = false
    while not isNil(bt):
        if bt.val == t:
            return true
        bt = bt.next

proc has*[T](s: ptr Set[T], t: T):bool=
    var idx = hashCode(repr(t)) mod s.size
    var bt = s.line[idx.int]
    result = false
    while not isNil(bt):
        if  bt.val == t:
            return true
        bt = bt.next


proc upSize*[T](s: ptr Set[T], newSize: int)=
    var oldLine = s.line
    s.size = newSize.uint
    s.len = 0
    s.line = newList[ptr SetBucket[T]](newSize)
    for i in 0..high(oldLine):
        var bt = oldLine[i]
        while not isNull(bt) and not isNull(bt.val):
            s.add(bt.val)
            bt = bt.next
    freeList(oldLine)  

proc del*[T](s: ptr Set[T], t: T):int{.discardable.}=
    result = 0
    var prev: ptr SetBucket[T]
    var idx = hashCode(repr(t)) mod s.size
    var bt = s.line[idx.int]
    if isNil(bt):
        return 0
    elif bt.val == t:
        s.line[idx.int] = bt.next
        s.len -= 1
        dealloc(bt)
        return 1
    else:
        while not isNil(bt.next):
            prev = bt
            bt = bt.next
            if isNil(bt):
                return 0
            elif bt.val == t:
                prev.next = bt.next
                s.len -= 1
                dealloc(bt)
                return 1

iterator each*[T](s: ptr Set[T]): T=
    for item in s.line.each:
        if not isNil(item):
            yield item.val
            var temp = item
            while not isNil(temp.next):
                temp = temp.next
                yield temp.val


when isMainModule:
    var set = newSet[int](16)

    # set.add(1)
    # set.add(2)
    # set.add(1)
    # echo(repr(set))
    # set.del(1)
    # echo set.has(2)
    # echo set.has(3)
    # set.upSize(1000000)
    # echo(GC_getStatistics())
    # set.freeSet()
    # echo(GC_getStatistics())

    var t = 10
    set.add(t)
    echo(set.has(t))

    set.add(2)
    set.add(3)
    for item in set.each:
        echo item






