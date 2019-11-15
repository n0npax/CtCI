package main

func uniqChar(s string) bool {
	m := map[rune]bool{}
	for _, c := range s {
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}
