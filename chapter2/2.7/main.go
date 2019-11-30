package main

import (
	"container/list"
)

func findIntersection(l1 *list.List, l2 *list.List) *list.List {
	rl := list.New()
	for el1, el2 := l1.Back(), l2.Back(); el1 != nil && el2 != nil && el1.Value == el2.Value; el1, el2 = el1.Prev(), el2.Prev() {
		rl.PushFront(el1.Value)
	}
	return rl
}
