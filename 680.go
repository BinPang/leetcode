package main

func validPalindrome(s string) bool {
	l := len(s)
	if l <= 1 {
		return true
	}
	start := 0
	end := l - 1
	for {
		if start >= end {
			return true
		}
		if s[start] == s[end] {
			start++
			end--
			continue
		}
		return _validPalindrome(&s, start, end-1) || _validPalindrome(&s, start+1, end)
	}
}
func _validPalindrome(s *string, start, end int) bool {
	for {
		if start >= end {
			return true
		}
		if (*s)[start] != (*s)[end] {
			return false
		}
		start++
		end--
	}
}
