package main

func main() {
	println("input:[1,2,4,2,5,7,2,4,9,0],need 13,result:", maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
	return
	println("input:[3,3,5,0,0,3,1,4],need 6,result:", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	println("input:[1,2,3,4,5],need 4,result:", maxProfit([]int{1, 2, 3, 4, 5}))
	println("input:[7,6,4,3,1],need 0,result:", maxProfit([]int{7, 6, 4, 3, 1}))
	println("input:[3,3,5,0,0,3,6,0,4,9],need 15,result:", maxProfit([]int{3, 3, 5, 0, 0, 3, 6, 0, 4, 9}))
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
	for i, v := range prices {
		println(i, ",", v)
	}

	max1 := 0
	max2 := 0
	tmp := 0
	i := 0
	j := 0
	for {
		if i >= l {
			break
		}
		if prices[i] < 0 {
			i++
			continue
		}
		j = i
		tmp = 0
		for {
			if j >= l || prices[j] < 0 {
				if tmp >= max2 {
					max1 = max2
					max2 = tmp
				} else if tmp > max1 {
					max1 = tmp
				}
				break
			}
			tmp += prices[j]
			j++
		}
		i = j + 1

	}
	return max1 + max2
}
