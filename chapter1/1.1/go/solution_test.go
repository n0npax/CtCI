package main

import (
	"testing"
)

func TestUniqChar(t *testing.T) {
	var testCases = []struct {
		s        string
		expected bool
	}{
		{"114", false},
		{"123", true},
		{"aac", false},
		{"abc", true},
		{"", true},
	}
	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			if uniqChar(tc.s) != tc.expected {
				t.Errorf("string %s didn't return %v", tc.s, tc.expected)
			}
		})
	}
}
