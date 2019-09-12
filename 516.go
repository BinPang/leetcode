package main

func main() {
	println(longestPalindromeSubseq("aa"))
	println(longestPalindromeSubseq("bbbab"))
	println(longestPalindromeSubseq("a"))
	println(longestPalindromeSubseq("cbbd"))
}

func longestPalindromeSubseq(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
		dp[i][i] = 1
	}
	for i := 0; i+1 < l; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 1
		}
	}
	for i := 2; i < l; i++ {
		for j := 0; j+i < l; j++ {
			if s[j] == s[i+j] {
				dp[j][i+j] = dp[j+1][i+j-1] + 2
			} else {
				dp[j][i+j] = max(dp[j][i+j-1], dp[j+1][i+j])
			}
		}
	}
	return dp[0][l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
Given a string s, find the longest palindromic subsequence's length in s.
You may assume that the maximum length of s is 1000.

Example 1:
Input:

"bbbab"
Output:
4
One possible longest palindromic subsequence is "bbbb".
Example 2:
Input:

"cbbd"
Output:
2
One possible longest palindromic subsequence is "bb".
*/
