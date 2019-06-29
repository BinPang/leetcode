package main

func main() {
	println("input:[1,2,4,2,5,7,2,4,9,0,9],need 17,result:", maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9}))
	println("input:[3,3],need 0,result:", maxProfit([]int{3, 3}))
	println("input:[1,2,4,2,5,7,2,4,9,0],need 13,result:", maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
	println("input:[3,3,5,0,0,3,1,4],need 6,result:", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	println("input:[7,6,4,3,1],need 0,result:", maxProfit([]int{7, 6, 4, 3, 1}))
	println("input:[3,3,5,0,0,3,6,0,4,9],need 15,result:", maxProfit([]int{3, 3, 5, 0, 0, 3, 6, 0, 4, 9}))
	println("input:[1,2,3,4,5],need 4,result:", maxProfit([]int{1, 2, 3, 4, 5}))
}
/*
step0: TODO time O(kn^2) space O(kn)
//第k次交易在第n天
dp[k][i] = max{
	dp[k][i-1] //第i天没有交易，比如前一天的价格是5，现在的价格为0
	p[i] - p[j] + dp[k-1][j] j=> [0, i-1]
}

step1: TODO time O(kn) space O(kn)
dp[2][1] = max {
	dp[2][0]
	p[1] - p[0] + dp[1][0]
}
dp[2][2] = max {
	dp[2][1]
	p[2] - p[0] + dp[1][0]
	p[2] - p[1] + dp[1][1]
}
dp[2][3] = max {
	dp[2][2]
	p[3] - p[0] + dp[1][0]
	p[3] - p[1] + dp[1][1]
	p[3] - p[2] + dp[1][2]
}

dp[2][4] = max {
	dp[2][2]
	p[4] - p[0] + dp[1][0]
	p[4] - p[1] + dp[1][1]
	p[4] - p[2] + dp[1][2]
	p[4] - p[3] + dp[1][3]
}

//see dp[2][4]
dp[2][4] = max {
	dp[2][2]
	p[4] - {p[0] + dp[1][0]} //dp[2][0] dp[2][1] dp[2][2] dp[2][3]
	p[4] - {p[1] + dp[1][1]} //dp[2][1] dp[2][2] dp[2][3]
	p[4] - {p[2] + dp[1][2]} //dp[2][2] dp[2][3]
	p[4] - {p[3] + dp[1][3]} //dp[2][3]
}
因此我们可以把前面dp[k-1][i] - p[i]的最大值记录下来
为啥能想到这个:需要把这个每步的表达式写下来，我那时没有写表达式，就按照第一步的方法
总是想不通，通过dp[k][i]具体的值（for循环展开）就明白了

next TODO
step2: TODO space O(kn)->O(k)

 */

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
	tmp := 0
	for k := 1; k <= K; k++ {
		for i := 1; i < l; i++ {
			if i == 1 {
				tmp = dp[k-1][i-1] - prices[i-1]
			} else {
				tmp = max(tmp, dp[k-1][i-1]-prices[i-1])
			}
			dp[k][i] = max(dp[k][i-1], prices[i]+tmp)
		}
	}
	//for k := 0; k <= K; k++ {
	//	for i := 0; i < l; i++ {
	//		print(dp[k][i], ",")
	//	}
	//	println("")
	//}
	return dp[K][l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}