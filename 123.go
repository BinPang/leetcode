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

	dp := make([]int, l)
	temp := 0

	for i := 1; i <= 2; i++ {
		balance := -prices[0]
		temp = dp[0]
		for j := 1; j < l; j++ {
			copy1 := dp[j]
			if dp[j-1] > balance+prices[j] {
				dp[j] = dp[j-1]
			} else {
				dp[j] = balance + prices[j]
			}
			println(j, "balance:", balance, "prices:", prices[j], "dp[j-1]:", dp[j-1], "dp now", dp[j])
			if balance <= temp-prices[j] {
				balance = temp - prices[j]
			}
			println(j, "balance:", balance, "temp:", temp, prices[j])
			temp = copy1
			//println("i:", i, ",j:", j, ",balance", balance, ",dp[j]:", dp[j])
		}
	}

	return dp[l-1]
}
