package main

func main() {
	println("input:ccd,need:cc,result:", longestPalindrome("ccd"))
	return
	println("input:caba,need:aba,result:", longestPalindrome("caba"))
	println("input:cbbd,need:bb,result:", longestPalindrome("cbbd"))
	println("input:babad,need:bab,result:", longestPalindrome("babad"))
	println("input:abba,need:abba,result:", longestPalindrome("abba"))
	println("input:,need:,result:", longestPalindrome(""))
}

func longestPalindrome(s string) string {
	l := len(s)
	r := ""
	for i := 0; i < l; i++ {
		left := i
		right := i
		use := true
		for {
			if left < 1 || right >= l {
				break
			}
			if s[left-1] == s[right+1] {
				if len(s[left:right]) > len(r) {
					r = s[left:right+1]
				}
				left -= 1
				right += 1
				use = false
			} else {
				if (s[left-1] == s[i] || s[right+1] == s[i]) && use {
					sr := ""
					if s[left-1] == s[i] {
						left -= 1
						println(left, i, right)
						sr = s[left:i]
					} else {
						right += 1
						println(2)
						sr = s[i:right]
					}
					println(string(s[i]), left, right, s[left:right], sr)
					if len(sr) > len(r) {
						r = sr
					}
					use = false
				} else {
					break
				}
			}
		}
	}
	return r
}
