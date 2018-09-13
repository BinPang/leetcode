package main

func main() {
	println("input:[7,1,5,3,6,4],need 7,result:", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	println("input:[7,1,5,3,20,4],need 21,result:", maxProfit([]int{7, 1, 5, 3, 20, 4}))
	println("input:[1,2,3,4,5],need 4,result:", maxProfit([]int{1, 2, 3, 4, 5}))
	println("input:[7,6,4,3,1],need 0,result:", maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	for i := l - 1; i >= 1; i-- {
		prices[i] = prices[i] - prices[i-1]
	}
	prices[0] = 0

	r := 0
	for _, v := range prices {
		if v > 0 {
			r += v
		}
	}
	return r
}
