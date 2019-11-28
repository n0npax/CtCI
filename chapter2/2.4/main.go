package main

import (
	"container/list"
)

func partitionByX(l *list.List, k int) *list.List {
	nl := list.New()
	for element := l.Front(); element != nil; element, _ = element.Next(), l.Remove(element) {
		v, _ := element.Value.(int)
		if v <= k {
			nl.PushFront(v)
		} else {
			nl.PushBack(v)
		}
	}
	return nl
}
