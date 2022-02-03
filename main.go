package main

import (
	"flag"
	"fmt"
)

var cm chunksmap

func main() {
	flag.Parse()
	cm := new(chunksmap)
	cm.dictionary = make(map[string]int)
	mh := maxheap{heapArray: []HeapElement{}, values: make(map[string]int, 5), size: 0, maxsize: 5}
	cm.hp = &mh

	cm.addString("is ball ball ball ball eggs eggs pool pool wild daily last")
	cm.addString("is is is is is")
	fmt.Printf("%s", cm.getTopMembers())
}
