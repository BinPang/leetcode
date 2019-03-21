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
	dp := make([][l]bool, l)
	for i := 0; i < l; i++ {
		dp[i][i] = true
	}
	min := make([]int, l)
	min[0] = 0

	for i := 1; i < l; i++ {
		min[i] = min[i-1] + 1
		for j := 0; j < i; j++ {

		}
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
