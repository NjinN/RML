proc freeToken*(t: ptr Token)

proc markMap*(map: ptr BindMap[ptr Token], mark: ptr Set[ptr Token], clean: ptr Set[ptr Token])=
    for item in map.each:
        mark.add(item)
        clean.del(item)
    for child in map.child.each:
        markMap(child, mark, clean)

proc gc*(mark: var ptr Set[ptr Token], stack: ptr EvalStack, map: ptr BindMap[ptr Token], inpSet: ptr Set[ptr List[ptr Token]], debug: cstring = "")=
    if mark.len.int < (int(mark.size.int / 4 * 3) - 4):
        return
    # echo("gc start ")
    # echo(debug)
    # echo(len(u.nowLine.line))
    # for item in u.nowLine.line.each:
    #     print item
    var newSet = newSet[ptr Token](mark.len.int)
    # echo("startCleanSize: ", s.len)
    for i in inpSet.each:
        for j in i.each:
            # print j
            newSet.add(j)
            mark.del(j)

    for i in 0..stack.idx - 1:
        newSet.add(stack.line[i])
        mark.del(stack.line[i])

    markMap(map, newSet, mark)

    # for i in mark.each:
    #     print i
 
    echo("newSize: ", newSet.len)
    echo("cleanSize:", mark.len)
    
    for t in mark.each:
        # print t
        if t.tp == TypeEnum.integer or t.tp == TypeEnum.string:
            freeToken(t)
    mark = newSet
    if mark.len.int > int(mark.size.int / 2):
        mark.upSize(mark.size.int * 2)

    # echo("gc end")






