package main

//cost my two days,but need back to check
func main() {
	println("need false, result:", isMatch("", ""))
	println("need true, result:", isMatch("", ".*.*.*a*.*"))
	println("need true, result:", isMatch("", ""))
	println("need false, result:", isMatch("aa", "a"))
	println("need false, result:", isMatch("acd", "ab.*cb.*d"))
	println("need true, result:", isMatch("abcbd", "ab.*cb.*d"))
	println("need false, result:", isMatch("aa", "*"))
	println("need true, result:", isMatch("aa", "a*"))
	println("need true, result:", isMatch("ab", ".*"))
	println("need true, result:", isMatch("aab", "c*a*b"))
	println("need false, result:", isMatch("mississippi", "mis*is*p*."))

}

func isMatch(s string, p string) bool {
	//from end to before(.=46 *=42)
	s = " " + s
	p = " " + p
	l1 := len(s)
	l2 := len(p)
	//p can't start with *
	if l2 > 1 && p[1] == 42 { //* sign equal 42
		return false
	}
	dp := make([][]bool, l1)
	for i := 0; i < l1; i++ {
		dp[i] = make([]bool, l2)
	}
	dp[0][0] = true
	for i := 0; i < l1; i++ { // s, input string             like:abbaaa
		for j := 1; j < l2; j++ { // p, reg check one by one like:ab*a*
			if p[j] == 42 { //is *
				if dp[i][j-2] {
					dp[i][j] = true
				}
				if (s[i] == p[j-2] || p[j-2] == 46) && dp[i][j-2] {
					dp[i][j] = true
				}
				if (s[i] == p[j-1] || p[j-1] == 46) && i > 0 && dp[i-1][j] {
					dp[i][j] = true
				}
			} else {
				if (s[i] == p[j] || p[j] == 46) && i > 0 && dp[i-1][j-1] {
					dp[i][j] = true
				}
			}
		}
	}
	//for i:=0;i<l2;i++ {
	//	if i==0 {
	//		print(" ,")
	//	}
	//	print(string(p[i]), ",")
	//}
	//println("")
	//for i := 0; i < l1; i++ {
	//	for j := 0; j < l2; j++ {
	//		if j == 0 {
	//			print(string(s[i]),",")
	//
	//		}
	//		if dp[i][j] {
	//			print("1,")
	//		} else {
	//			print("0,")
	//		}
	//	}
	//	println("")
	//}
	return dp[l1-1][l2-1]
}
