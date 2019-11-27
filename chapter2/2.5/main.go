package main

import (
	"container/list"
	"math"
)

func sumList(l1 *list.List, l2 *list.List) int {
	shift, sum := 0, 0
	for pow, el1, el2 := 0, l1.Front(), l2.Front(); el1 != nil || el2 != nil; pow++ {
		v1, v2 := 0, 0
		if el1 != nil {
			v1, _ = el1.Value.(int)
			el1 = el1.Next()
		}
		if el2 != nil {
			v2, _ = el2.Value.(int)
			el2 = el2.Next()
		}
		v := v1 + v2 + shift
		shift = v / 10
		sum += v * int(math.Pow10(pow))
	}
	return sum
}
