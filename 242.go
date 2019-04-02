package main

func isAnagram(s string, t string) bool {
	l := len(s)
	if l != len(t) {
		return false
	}
	h := make(map[uint8]int, 26)
	for i := 0; i < l; i++ {
		h[s[i]]++
		h[t[i]]--
	}
	for _, v := range h {
		if v != 0 {
			return false
		}
	}
	return true
}
