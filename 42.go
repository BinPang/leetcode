package main

func main() {
	println("need 0", trap([]int{}))
	println("need 2", trap([]int{2, 0, 2}))
	println("need 6", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	m := 0
	max := 0
	for i, v := range height {
		if max < v {
			max = v
			m = i
		}
	}
	maxLevelNow := 0
	total := 0
	fakerTotal := 0
	for i := 0; i <= m; i++ {
		if maxLevelNow == 0 {
			maxLevelNow = height[i]
			continue
		}
		if height[i] < maxLevelNow {
			fakerTotal += maxLevelNow - height[i]
		} else if height[i] >= maxLevelNow {
			total += fakerTotal
			fakerTotal = 0
			maxLevelNow = height[i]
		}
	}
	l := len(height)
	maxLevelNow = 0
	fakerTotal = 0
	for i := l - 1; i >= m; i-- {
		if maxLevelNow == 0 {
			maxLevelNow = height[i]
			continue
		}
		if height[i] < maxLevelNow {
			fakerTotal += maxLevelNow - height[i]
		} else if height[i] >= maxLevelNow {
			total += fakerTotal
			fakerTotal = 0
			maxLevelNow = height[i]
		}
	}
	return total
}
