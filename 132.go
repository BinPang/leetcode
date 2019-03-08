package main

//import "fmt"

func main() {
	println(minCut("aaaab"))
	println(minCut("aabbaaaa"))
}

func minCut(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}

	start := 0
	length := 0
	for {
		for {
			if palindrome(s[start : start+length]) {
				dp[start][start+length] = 0
			} else {
				
			}
		}
		length++
		start = 0
	}

	return 0
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
