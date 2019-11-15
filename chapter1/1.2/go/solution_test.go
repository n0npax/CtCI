package main

import (
	"fmt"
	"testing"
)

func TestIsPermutation(t *testing.T) {
	var testCases = []struct {
		s1, s2   string
		expected bool
	}{
		{"114", "123", false},
		{"123", "321", true},
		{"aac", "llama", false},
		{"abc", "cba", true},
		{"", "", true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s-%s", tc.s1, tc.s2), func(t *testing.T) {
			if isPermutation(tc.s1, tc.s2) != tc.expected {
				t.Errorf("IsPermutation(%v, %v) should return %v", tc.s1, tc.s2, tc.expected)
			}
		})
	}
}
