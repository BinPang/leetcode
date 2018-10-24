package main

func main() {
	println(numDistinct("babgbag", ""))
	println(numDistinct("", "bag"))
	println(numDistinct("babgbag", "bag"))
	println(numDistinct("rabbbit", "rabbit"))
}

func numDistinct(s string, t string) int {
	s = " " + s
	t = " " + t
	ls := len(s)
	lt := len(t)
	dp := make([][]int,lt)
	for i := range dp  {
		dp[i] = make([]int, ls)
	}
	for i := ls-1; i >= 0; i-- {
		dp[0][i] = 1
	}
	for i := lt-1; i > 0; i-- {
		dp[i][0] = 0
	}
	for i := 1;i<lt;i++ {
		for j := 1; j < ls; j++ {
			if t[i] == s[j] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
			//print(dp[i][j], ", ")
		}
		//println("")
	}
	return dp[lt-1][ls-1]
}
