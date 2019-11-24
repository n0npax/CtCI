package main

import (
	"container/list"
	"testing"
)

func Test_kToLastElement(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		k       int
		want    interface{}
	}{
		{"empty k0", []int{}, 0, nil},
		{"empty k1", []int{}, 1, nil},
		{"len6 k2", []int{7, 9, 3, 4, 5, 6}, 2, 4},
		{"len6 k0", []int{7, 9, 3, 4, 5, 6}, 0, 6},
		{"len6 k5", []int{7, 9, 3, 4, 5, 6}, 5, 7},
		{"len6 k8", []int{7, 9, 3, 4, 5, 6}, 8, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := list.New()
			for _, v := range tt.numbers {
				l.PushBack(v)
			}
			if got := kToLastElement(l, tt.k); got != tt.want {
				t.Errorf("kToLastElement(%v, %v) = %v, want %v", tt.numbers, tt.k, got, tt.want)
			}
		})
	}
}
