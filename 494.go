package main

func main() {
	println(findTargetSumWays([]int{0, 0}, 0))
	println(findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
	println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, S int) int {
	/*
		l := len(nums)
		if l == 0 {
			return 0
		}
		old := map[int]int{nums[0]: 1, -nums[0]: 1}
		if nums[0] == 0 {
			old[0] = 2
		}
		now := make(map[int]int)
		for i := 1; i < l; i++ {
			for k, v := range old {
				now[k+nums[i]] += v
				now[k-nums[i]] += v

			}
			old = now
			now = make(map[int]int, len(old))
		}
		return old[S]
	*/
	l := len(nums)
	if l == 0 {
		return 0
	}
	max := 1000
	if S > max || S < -max {
		return 0
	}
	old := make([]int, max*2+1)
	now := make([]int, max*2+1)
	if nums[0] == 0 {
		old[nums[0]+max] = 2
	} else {
		old[nums[0]+max] = 1
		old[-nums[0]+max] = 1
	}
	for i := 1; i < l; i++ {
		for sum := -max; sum <= max; sum++ {
			if old[sum+max] > 0 {
				now[sum+nums[i]+max] += old[sum+max]
				now[sum-nums[i]+max] += old[sum+max]
			}
			old[sum+max] = 0
		}
		old, now = now, old
	}
	return old[S+max]
}
