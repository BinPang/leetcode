package main

func main() {
	println("input:[1,2,4,2,5,7,2,4,9,0],need 13,result:", maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
	return
	println("input:[3,3,5,0,0,3,1,4],need 6,result:", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	println("input:[1,2,3,4,5],need 4,result:", maxProfit([]int{1, 2, 3, 4, 5}))
	println("input:[7,6,4,3,1],need 0,result:", maxProfit([]int{7, 6, 4, 3, 1}))
	println("input:[3,3,5,0,0,3,6,0,4,9],need 15,result:", maxProfit([]int{3, 3, 5, 0, 0, 3, 6, 0, 4, 9}))
}

//no no no ,I copy from other, I don't know why.TODO I am the stupidest programmer in the world
//TODO I will solve the problem, event though
func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	K := 2
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, l)
	}
	for k := 1; k <= K; k++ {
		//第i天无交易，最大收益就是为dp[k][i-1]
		//第i天有交易，最大收益为dp[k-1][i]
		//第i天有交易，最大收益为prices[i]-prices[j] + dp[k-1][j-1]，j为[0,i-1]
		for i := 1; i < l; i++ {
			tmp := max(dp[k][i-1], dp[k-1][i])
			for j := 0; j <= i-1; j++ {
				if j == 0 {
					tmp = max(tmp, prices[i]-prices[j])
				} else {
					tmp = max(tmp, prices[i]-prices[j]+dp[k-1][j-1])
				}
			}
			dp[k][i] = tmp
		}
	}
	for k := 0; k <= K; k++ {
		for i := 0; i < l; i++ {
			print(dp[k][i], ",")
		}
		println("")
	}
	return dp[K][l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
