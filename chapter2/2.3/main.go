package main

import (
	"container/list"
)

func rmElFromMiddle(l *list.List) {

	if l == nil {
		return
	}
	front := l.Front()
	if front == nil {
		return
	}
	second := front.Next()
	if second == nil {
		return
	}
	third := second.Next()
	if third == nil {
		return
	}
	l.Remove(second)
}
