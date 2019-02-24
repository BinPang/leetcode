package main

import "strings"

func main() {
	println(isScramble("aa", "ba"))
	return
	println(isScramble("aa", "ab"))
	return
	println(isScramble("great", "rgeat"))
	println(isScramble("abcde", "caebd"))
}

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	index := strings.Index(s2, string(s1[0]))
	//println(s1, s2, index)
	if index == -1 {
		return false
	}
	if index == 0 {
		return isScramble(s1[index+1:], s2[index+1:])
	} else {
		return isScramble(s1[1:index], s2[:index-1]) && isScramble(s1[index+1:], s2[index+1:])
	}
}
