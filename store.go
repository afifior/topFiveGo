package main

import (
	"fmt"
	"strings"
	"sync"
)

type chunksmap struct {
	dictionary map[string]int
	hp         *maxheap
	mu         sync.RWMutex
}

func (mp *chunksmap) addString(text string) {
	chunks := strings.Split(text, " ")
	mp.mu.Lock()
	defer mp.mu.Unlock()
	for _, chunk := range chunks {
		if _, ok := mp.dictionary[chunk]; ok {
			mp.dictionary[chunk] = 1 + mp.dictionary[chunk]
		} else {
			mp.dictionary[chunk] = 1
		}
		he := HeapElement{index: mp.dictionary[chunk], value: chunk}
		mp.hp.insert(&he)
	}
}

func (mp *chunksmap) getTopMembers() string {
	mp.mu.RLock()
	defer mp.mu.RUnlock()
	var retstr = ""
	for _, val := range mp.hp.heapArray {
		retstr += fmt.Sprintf("%s -> %d\n", val.value, val.index)
	}
	return retstr
}
