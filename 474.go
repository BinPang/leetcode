package main

func main() {
	println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
	println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 3, 5))
}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	count0, count1 := 0, 0
	for _, v := range strs {
		count0, count1 = 0, 0
		for _, v1 := range v {
			if v1 == '0' {
				count0++
			} else {
				count1++
			}
		}
		if count0 > m || count1 > n {
			continue
		}
		for i := m; i >= count0; i-- {
			for j := n; j >= count1; j-- {
				dp[i][j] = max(dp[i][j], dp[i-count0][j-count1]+1)
			}
		}
		//println(fmt.Sprintf("%+v", dp))
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
