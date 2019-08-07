package main

func longestPalindrome(s string) int {
	tmp := make([]int, 52)
	r := 0
	for _, v := range s {
		if v <= 'Z' {
			tmp[v-'A'] += 1
		} else {
			tmp[v-'a'+26] += 1
		}
	}
	plus := false
	for _, v := range tmp {
		if v&1 == 1 {
			r += v - 1
			plus = true
		} else {
			r += v
		}
	}
	if plus {
		r += 1
	}
	return r
}
