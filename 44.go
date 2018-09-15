package main

func main() {
	println("input:,need:true,return:", isMatch("adceb", "*a*b"))
	println("input:,need:false,return:", isMatch("acdcb", "a*c?b"))
}

func isMatch(s string, p string) bool {
	s = " " + s
	p = " " + p
	l1 := len(s)
	l2 := len(p)
	dp := make([][]bool, l2)
	for i := 0; i < l2; i++ {
		dp[i] = make([]bool, l1)
	}
	dp[0][0] = true
	//first column
	for i := 1; i < l2; i++ {
		if p[i] == '*' {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if p[j] == 42 { //*
				if dp[j-1][i] || dp[j][i-1] || dp[j-1][i-1] {
					dp[j][i] = true
				}
			} else if p[j] == 63 { //?
				dp[j][i] = dp[j-1][i-1]
			} else { //normal
				if p[j] == s[i] && dp[j-1][i-1] {
					dp[j][i] = true
				}
			}
		}
	}

	return dp[l2-1][l1-1]
}
