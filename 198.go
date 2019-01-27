package main

func main() {
	println(rob([]int{8, 1, 1, 1, 1, 1, 8}))
	println(rob([]int{1, 2, 3, 1}))
	println(rob([]int{2, 7, 9, 3, 1}))
	println(rob([]int{2, 1, 1, 2}))
	println(rob([]int{8, 1, 1, 1, 1, 1, 8}))
	println(rob([]int{2, 1}))
}

//easiest dp
//dp[i] = max(dp[i-1], dp[i-2]+a[i])
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	} else if l == 2 {
		return max(nums[0], nums[1])
	}
	if nums[1] < nums[0] {
		nums[1] = nums[0]
	}
	for i := range nums {
		if i == 0 || i == 1 {
			continue
		}
		nums[i] = max(nums[i-1], nums[i-2]+nums[i])
	}
	return nums[l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
