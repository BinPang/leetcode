package main

func main() {
	minDistance("abc", "aaabd")
}

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	dp := make([][]int, l1+2)
	for i := range dp {
		dp[i] = make([]int, l2+2)
	}
	for i := 1; i < l2+2; i++ {
		dp[1][i] = i - 1
	}
	for i := 1; i < l1+2; i++ {
		dp[i][1] = i - 1
	}
	min:=0
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if word1[i] == word2[j] {
				dp[i+2][j+2]=dp[i+1][j+1]
			} else {
				min = dp[i+1][j+1]
				if min > dp[i+2][j+1] {
					min = dp[i+2][j+1]
				}
				if min >  dp[i+1][j+2] {
					min = dp[i+1][j+2]
				}
				dp[i+2][j+2] = min +1
			}
		}
	}
	//_printArrayArray(dp)
	return dp[l1+1][l2+1]
}

func _printArrayArray(t [][]int) {
	for _, v := range t {
		_printArray(v)
	}
}
func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
