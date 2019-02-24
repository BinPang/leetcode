package main

import "fmt"

/**
Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
 */
// aabbaa
// [
// ["a","a","b","b","a","a"],
// ["a","a","b","b","aa"],
// ["a","a","bb","a","a"],
// ["a","a","bb","aa"],
// ["a","abba","a"],
// ["aa","b","b","a","a"],
// ["aa","b","b","aa"],
// ["aa","bb","a","a"],
// ["aa","bb","aa"],
// ["aabbaa"]
// ]
func main() {
	v := partition("aabbaa")
	println(fmt.Sprintf("%+v", v))

	v1 := partition("aabbaaaa")
	println(fmt.Sprintf("%+v", v1))
}

func partition(s string) [][]string {
	if s == "" {
		return nil
	}
	r := make([][]string, 0)
	for i := 0; i < len(s); i++ {
		s1 := s[:i+1]
		if palindrome(s1) {
			tmp := partition(s[i+1:])
			for _, v := range tmp {
				r = append(r, append([]string{s1}, v...))
			}
			if s1 == s {
				r = append(r, []string{s1})
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
