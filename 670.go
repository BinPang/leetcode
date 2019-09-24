package main

func main() {
	println(maximumSwap(2736))
	println(maximumSwap(9973))
	println(maximumSwap(23736))
	println(maximumSwap(73736))
	println(maximumSwap(5))
}

func maximumSwap(num int) int {
	if num == 0 {
		return 0
	}
	lowToHigh := make([]int, 0)
	for num > 0 {
		lowToHigh = append(lowToHigh, num%10)
		num = num / 10
	}
	//[3 7 9 9] [3 7 9 7 9] [6 3 7 2]
	l := len(lowToHigh)
	i := l - 1
	j := 0
	max := 0
	for i > 0 {
		j = i - 1
		max = lowToHigh[i]
		for j >= 0 {
			if max < lowToHigh[j] {
				max = lowToHigh[j]
			}
			j--
		}
		if lowToHigh[i] == max {
			i--
		} else {
			for j = 0; j < i; j++ {
				if lowToHigh[j] == max {
					lowToHigh[j], lowToHigh[i] = lowToHigh[i], lowToHigh[j]
					break
				}
			}
			break
		}
	}
	num = 0
	start := 1
	for _, v := range lowToHigh {
		num += v * start
		start *= 10
	}

	return num
}
