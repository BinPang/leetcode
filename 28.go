package main

func main() {
	println(strStr("hello", "ll"))
	println(strStr("hello", "lo"))
	println(strStr("hello", "aa"))
}
//no kmp,because it's an easy algorithm, just compare
//besides i need to learn <introduction to algorithm again
func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	if l2 == 0 {
		return 0
	}
	for i := 0; i < l1-l2+1; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}
