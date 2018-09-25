package main

func main() {
	println("8176,need 2,result:", numDecodings("8176"))
	println("301,need 0,result:", numDecodings("301"))
	println("17,need 2,result:", numDecodings("17"))
	println("27,need 1,result:", numDecodings("27"))
	println("226,need 3,result:", numDecodings("226"))
}

func numDecodings(s string) int {
	l := len(s)
	if l == 0 || s[0] == '0' {
		return 0
	}
	if l == 0 || s[0] == '0' {
		return 0
	} else if l == 1 {
		return 1
	}
	dp := make([]int, l)
	dp[0], dp[1] = 1, 1
	if s[1] == '0' && s[0] > '2' {
		return 0
	}
	if s[0] == '1' || s[0] == '2' {
		if s[1] == '0' {
			dp[1] = 1
		} else if s[0] == '1' || (s[0] == '2' && s[1] < '7') {
			dp[1] = 2
		}
	}

	for i := 2; i < l; i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dp[i] = dp[i-2]
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7') {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[l-1]
}
