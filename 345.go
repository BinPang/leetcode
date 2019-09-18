package main

func main() {
	println(reverseVowels("a"))
	println(reverseVowels("ai"))
	println(reverseVowels("race car"))
	println(reverseVowels("hello"))
	println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	r := []byte(s)
	start := 0
	end := l - 1
	for {
		if start >= end {
			return string(r)
		}
		if s[start] == 'a' || s[start] == 'e' || s[start] == 'i' || s[start] == 'o' || s[start] == 'u' || s[start] == 'A' || s[start] == 'E' || s[start] == 'I' || s[start] == 'O' || s[start] == 'U' {
			for {
				if end == start {
					return string(r)
				}
				if s[end] == 'a' || s[end] == 'e' || s[end] == 'i' || s[end] == 'o' || s[end] == 'u' || s[end] == 'A' || s[end] == 'E' || s[end] == 'I' || s[end] == 'O' || s[end] == 'U' {
					r[start], r[end] = r[end], r[start]
					end--
					break
				}
				end--
			}
		}
		start++
	}
}
