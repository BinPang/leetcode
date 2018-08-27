package main

func main() {
	println("input:cbbd,need:bb,result:", longestPalindrome("cbbd"))
	println("input:cbbc,need:bb,result:", longestPalindrome("cbbc"))

	println("input:caba,need:aba,result:", longestPalindrome("caba"))
	println("input:ccd,need:cc,result:", longestPalindrome("ccd"))
	println("input:babad,need:bab,result:", longestPalindrome("babad"))
	println("input:abba,need:abba,result:", longestPalindrome("abba"))
	println("input:,need:,result:", longestPalindrome(""))
}

func longestPalindrome(s string) string {
	l := len(s)
	r := ""
	for i := 0; i < l; i++ {
		left := i-1
		right := i+1
		tmp := ""
		//case one, item in middle
		for {
			if left < 0 || right >= l {
				tmp = s[left+1 : right]
				break
			}
			if s[left] != s[right] {
				tmp = s[left+1 : right]
				break
			}
			left--
			right++
		}
		if len(r) < len(tmp) {
			r = tmp
		}
		//case two, item and next item's left&right, we don't use item's before
		left = i
		right = i + 1
		tmp = ""
		for {
			if left < 0 || right >= l {
				tmp = s[left+1:right]
				break
			}
			if s[left] != s[right] {
				tmp = s[left+1:right]
				break
			}
			left--
			right++
		}
		if len(r) < len(tmp) {
			r = tmp
		}
	}
	return r
}
