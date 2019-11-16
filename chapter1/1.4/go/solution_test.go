package main

import (
	"reflect"
	"testing"
)

func TestIsPermutation(t *testing.T) {
	var testCases = []struct {
		s        string
		expected []string
	}{
		{"", []string{""}},
		{"abc", nil},
		{"aab", []string{"aab", "aba", "baa"}},
	}
	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			palindromes := PalindromePermutations(tc.s)
			if reflect.DeepEqual(palindromes, tc.expected) {
				t.Errorf("PalindromePermutations(%v) should return %[2]v(%[2]T) but got %[3]v(%[3]T)", tc.s, tc.expected, palindromes)
			}
		})
	}
}
