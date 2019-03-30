package main

func main() {
	println(majorityElement([]int{6, 5, 5}))
}

func majorityElement(nums []int) int {
	l := len(nums)
	count := 1
	for i := 1; i < l; i++ {
		if nums[i] == nums[0] {
			count++
		} else {
			count--
			if count < 0 {
				nums[0] = nums[i]
				count = 1
			}
		}
	}

	return nums[0]
}
