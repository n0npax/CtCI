package main

import "container/list"

func contains_dupl(l *list.List) bool {
	occurence := map[interface{}]bool{}
	for element := l.Front(); element != nil; element = element.Next() {
		if occurence[element.Value] {
			return true
		}
		occurence[element.Value] = true
	}
	return false
}
