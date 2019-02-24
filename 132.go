package main

func main() {
	println(minCut("aab"))
	println(minCut("aabbaaaa"))
}

func minCut(s string) int {
	r := len(s) - 1
	for i := 0; i < len(s); i++ {
		s1 := s[:i+1]
		if palindrome(s1) {
			if s1 == s {
				return 0
			}
			tmp := minCut(s[i+1:])
			if tmp+1 < r {
				r = tmp + 1
			}
		}
	}
	return r
}

func palindrome(s1 string) bool {
	start := 0
	end := len(s1) - 1
	for {
		if start > end {
			return true
		}
		if s1[start] != s1[end] {
			return false
		}
		start++
		end--
	}
}
