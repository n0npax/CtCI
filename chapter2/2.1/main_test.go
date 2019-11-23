package main

import (
	"container/list"
	"testing"
)

func Test_contains_dupl(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want bool
	}{
		{"empty", []int{}, false},
		{"on item", []int{}, false},
		{"no dupl", []int{1, 2, 3, 4, 5, 6}, false},
		{"single dupl", []int{1, 2, 3, 4, 5, 1}, true},
		{"multiple dupl", []int{1, 2, 3, 2, 3}, true},
		{"just dupl", []int{1, 1, 1, 1, 1, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := list.New()
			for _, v := range tt.args {
				l.PushFront(v)
			}
			if got := contains_dupl(l); got != tt.want {
				t.Errorf("contains_dupl() = %v, want %v", got, tt.want)
			}
		})
	}
}
