package main

import (
	"container/list"
)

func partitionByX(l *list.List, k int) *list.List {
	nl := list.New()
	for element := l.Front(); element != nil; element, _ = element.Next(), l.Remove(element) {
		v, ok := element.Value.(int)
		if !ok {
			panic("list expected to contain integers")
		}
		if v <= k {
			nl.PushFront(v)
		} else {
			nl.PushBack(v)
		}
	}
	return nl
}
