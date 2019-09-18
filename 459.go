package main

func repeatedSubstringPattern(s string) bool {
	l := len(s)
	if l&1 == 1 {
		return false
	}
	l2 := l >> 2
	for i := 0; i < l2; i++ {
		if s[i] != s[l2+i] {
			return false
		}
	}
	return true
}
