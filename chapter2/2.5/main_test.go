package main

import (
	"container/list"
	"testing"
)

func Test_rmElFromMiddle(t *testing.T) {

	tests := []struct {
		name     string
		numbers1 []int
		numbers2 []int
		want     int
	}{
		{"len 0 & 0", []int{}, []int{}, 0},
		{"nil & nil", nil, nil, 0},
		{"len 1 & 1", []int{7}, []int{7}, 14},
		{"len 4 & 4", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, 8642},
		{"len 4 & 2", []int{1, 2, 3, 4}, []int{3, 4}, 4364},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := list.New()
			for _, v := range tt.numbers1 {
				l1.PushBack(v)
			}
			l2 := list.New()
			for _, v := range tt.numbers2 {
				l2.PushBack(v)
			}
			got := sumList(l1, l2)

			if got != tt.want {
				t.Errorf("sumList(%v,%v) = %v, want %v", tt.numbers1, tt.numbers2, got, tt.want)
			}
		})
	}
}
