package main

import (
	"container/list"
)

func isPalindrome(l *list.List) bool {
	for lPointer, rPointer := l.Front(), l.Back(); lPointer != nil && rPointer != nil && rPointer != lPointer && lPointer.Next() != rPointer; lPointer, rPointer = lPointer.Next(), rPointer.Prev() {
		if lPointer.Value != rPointer.Value {
			return false
		}
	}
	return true
}
