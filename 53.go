package main

func main() {
	println("input:[-2,1,-3,4,-1,2,1,-5,4] need 6,return:", maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	println("input:[-2,1,-3,4,-1,2,1,-5,8] need 9,return:", maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,8}))
	println("input:[-2,10,-3,4,-1,2,1,-5,1] need 13,return:", maxSubArray([]int{-2,10,-3,4,-1,2,1,-5,1}))
	println("input:[-2,-3,-4] need -2,return:", maxSubArray([]int{-2,-3,-4}))
	println("input:[-4,-3,-2] need -2,return:", maxSubArray([]int{-4,-3,-2}))

}

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		panic("need at least one item")
	}
	max := nums[0]
	tmp := 0
	for _ , v := range nums {
		if tmp+v > 0 {
			tmp+=v
			if max < tmp {
				max = tmp
			}
			continue
		} else {
			tmp = 0
		}
		if v > max {
			max = v
		}
	}
	return max
}
