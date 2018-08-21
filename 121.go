package main

func main() {
	println("need 5,result:", maxProfit([]int{7,1,5,3,6,4}))
	println("need 0,result:", maxProfit([]int{7,6,4,3,1}))
	println("need 5,result:", maxProfit([]int{2,5,1,6,3}))
	println("need 5,result:", maxProfit([]int{1,2,5,6,3}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min := 0
	max := 0
	lastMax := 0
	for i, v := range prices {
		if i == 0 {
			min = v
			max = v
			continue
		}
		if v > max {
			max = v
		}
		if v < min {
			nowMax := max - min
			if nowMax > lastMax {
				lastMax = nowMax
			}
			max = v
			min = v
		}
	}
	nowMax := max - min
	if nowMax > lastMax {
		lastMax = nowMax
	}
	return lastMax
}
