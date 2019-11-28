package main

import (
	"container/list"
	"testing"
)

func Test_isPalindrome(t *testing.T) {

	tests := []struct {
		name    string
		numbers []int
		want    bool
	}{
		{"len 0 palindrome", []int{}, true},
		{"len 1 palindrome", []int{1}, true},
		{"len 4 palindrome", []int{1, 2, 2, 1}, true},
		{"len 4 not palindrome", []int{1, 2, 3, 4}, false},
		{"len 5 palindrome", []int{1, 2, 3, 2, 1}, true},
		{"len 5 not palindrome", []int{1, 2, 3, 3, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := list.New()
			for _, v := range tt.numbers {
				l.PushBack(v)
			}

			got := isPalindrome(l)

			if got != tt.want {
				t.Errorf("isPalindrome(%v) = %v, want %v", tt.numbers, got, tt.want)
			}
		})
	}
}
