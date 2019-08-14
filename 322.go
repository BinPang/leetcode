package main

import "fmt"

func main() {
	println(coinChange([]int{2, 5, 9}, 18))
	println(coinChange([]int{1, 2, 5}, 11))
	println(coinChange([]int{2}, 3))
}

func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	} else if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for _, v := range coins {
		if 0 < v && v <= amount {
			dp[v] = 1
		}
	}
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if i <= v || v <= 0 {
				continue
			}
			if dp[i-v] > 0 {
				if dp[i] == 0 {
					dp[i] = dp[i-v] + 1
				} else {
					if dp[i] > dp[i-v]+1 {
						dp[i] = dp[i-v] + 1
					}
				}
			}
			//println(fmt.Sprintf("%d,%d,%+v", i, v, dp))
		}
	}
	println(fmt.Sprintf("%+v", dp))
	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}
