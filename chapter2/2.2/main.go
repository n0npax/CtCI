package main

import (
	"container/list"
)

func kToLastElement(l *list.List, k int) interface{} {

	kElement := l.Front()
	if kElement == nil {
		// empty list returns nil
		return nil
	}
	for element := l.Front(); element != nil; element = element.Next() {
		if k < 0 {
			kElement = kElement.Next()
		} else {
			k--
		}
	}
	if k > 0 {
		// k grater than list length
		return nil
	}
	return kElement.Value

}
