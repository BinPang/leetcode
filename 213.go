package main

func main() {
	println(rob([]int{1, 3, 1, 3, 100}))
}

func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	} else if l == 2 {
		return max(nums[0], nums[1])
	} else if l == 3 {
		return max(max(nums[0], nums[1]), nums[2])
	}
	r := 0
	//[0 ~ l-2], [1, l-1]
	prevPrev, prev := nums[0], max(nums[0], nums[1])
	tmp := 0
	for i := 2; i <= l-2; i++ {
		tmp = prevPrev + nums[i]
		if tmp > prev {
			prevPrev = prev
			prev = tmp
		} else {
			prevPrev = prev
		}
	}
	r = prev

	prevPrev, prev = nums[1], max(nums[1], nums[2])
	tmp = 0
	for i := 3; i <= l-1; i++ {
		tmp = prevPrev + nums[i]
		println(i, tmp, prevPrev, prev)
		if tmp > prev {
			prevPrev = prev
			prev = tmp
		} else {
			prevPrev = prev
		}
	}

	return max(r, prev)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
