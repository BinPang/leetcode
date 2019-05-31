package main

func main() {
	println(maxProduct([]int{2, -5, -2, -4, 3}))
	println(maxProduct([]int{3, -1, 4}))
	println(maxProduct([]int{-2, 0, -1}))
	println(maxProduct([]int{2, 3, -2, 4}))
	println(maxProduct([]int{0, 3, 2, 4}))
	println(maxProduct([]int{-2, 3, -2, 4}))
	println(maxProduct([]int{-2, 3, 0, 0, 0, -2, 4, -2, 0, 0, 3, 4}))
}

func maxProduct(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	r := nums[0]
	preZero := false
	if nums[0] == 0 {
		preZero = true
	}
	maxPre := nums[0]
	minPre := nums[0]
	tmp := 0
	for i := 1; i < l; i++ {
		if nums[i] == 0 {
			if r < 0 {
				r = 0
			}
			preZero = true
		} else {
			if preZero {
				maxPre = nums[i]
				minPre = nums[i]
			} else {
				if nums[i] > 0 {
					maxPre = pmax(maxPre*nums[i], nums[i])
					minPre = pmin(minPre*nums[i], nums[i])
				} else {
					tmp = pmax(minPre*nums[i], nums[i])
					minPre = pmin(maxPre*nums[i], nums[i])
					maxPre = tmp
				}
			}
			preZero = false
		}
		if r < maxPre {
			r = maxPre
		}
	}

	return r
}

func pmax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func pmin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
