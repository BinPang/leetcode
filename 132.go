package main

//import "fmt"

func main() {
	println(minCut("leet"))
	println(minCut("aaaab"))
	println(minCut("aabbaaaa"))
}

//some one is clever
func minCut(s string) int {
	l := len(s)
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}
	cut := make([]int, l)
	min := 0

	for i := 0; i < l; i++ {
		min = i
		for j := 0; j <= i; j++ {
			if s[j] == s[i] && (j+1 > i-1 || dp[j+1][i-1]) {
				dp[j][i] = true
				if j == 0 {
					min = 0
				} else {
					if min > cut[j-1]+1 {
						min = cut[j-1] + 1
					}
				}
			}
		}
		cut[i] = min
	}

	return cut[l-1]
}
