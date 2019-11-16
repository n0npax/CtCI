package main

func PalindromePermutations(s string) []string {
	m := runeCounter(s)
	oddCnt := 0
	for _, v := range m {
		if v%2 != 0 {
			oddCnt++
		}
	}
	if oddCnt > 1 {
		return []string{}
	}
	return permutations(s)
}

func permutations(s string) []string {
	ret := []string{}
	if len(s) == 1 {
		return []string{s}
	}
	for i := range s {
		rl := []rune(s)
		if rl[i] == rl[0] {
			continue
		}
		rl[0], rl[i] = rl[i], rl[0]
		sl := string(rl)
		ret = append(ret, permutations(sl[1:])...)
	}
	return ret
}

func runeCounter(s string) map[rune]int {
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}
	return m
}
