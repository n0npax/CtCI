package main

import "reflect"

func isPermutation(s1, s2 string) bool {
	m1 := runeCounter(s1)
	m2 := runeCounter(s2)
	return reflect.DeepEqual(m1, m2)
}

func runeCounter(s string) map[rune]int {
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}
	return m
}
