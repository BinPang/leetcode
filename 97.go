package main

func main() {
	println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	println(isInterleave("", "aab", "aab"))
	println(isInterleave("aabcc", "dbbcad", "aadbbcbcacd"))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)
	if l1+l2 != l3 {
		return false
	}

	dp := make([][]bool, l2+1)
	for i := 0; i <= l2; i++ {
		dp[i] = make([]bool, l1+1)
	}
	dp[0][0] = true
	for i := 1; i <= l1; i++ {
		if s1[i-1] == s3[i-1] {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= l2; i++ {
		if s2[i-1] == s3[i-1] {
			dp[i][0] = true
		} else {
			break
		}
	}
	for i := 1; i <= l2; i++ {
		for j := 1; j <= l1; j++ {
			if s2[i-1] == s3[i+j-1] && dp[i-1][j] {
				dp[i][j] = true
			} else if s1[j-1] == s3[i+j-1] && dp[i][j-1] {
				dp[i][j] = true
			}
		}
	}
	//for _, v := range dp {
	//	for _, v1 := range v {
	//		if v1 {
	//			print("T,")
	//		} else {
	//			print("F,")
	//		}
	//
	//	}
	//	println("")
	//}

	return dp[l2][l1]
}
