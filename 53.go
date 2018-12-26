package main

func main() {
	println("input:[-2,1,-3,4,-1,2,1,-5,4] need 6,return:", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	println("input:[-2,1,-3,4,-1,2,1,-5,8] need 9,return:", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 8}))
	println("input:[-2,10,-3,4,-1,2,1,-5,1] need 13,return:", maxSubArray([]int{-2, 10, -3, 4, -1, 2, 1, -5, 1}))
	println("input:[-2,-3,-4] need -2,return:", maxSubArray([]int{-2, -3, -4}))
	println("input:[-4,-3,-2] need -2,return:", maxSubArray([]int{-4, -3, -2}))

}

//second.
// first有2次循环，而第一次算的dp没用，可以在第一次循环中计算出来
//
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	beforeDpItem := nums[0]
	var nowDpItem int//not init
	r := nums[0]
	//这个时候不管第i-1项，只看第i项
	//第i项的最大值=a[i]+(第i-1项目 > 0 ? 第i-1项目 : 0)
	for i := 1; i < l; i++ {
		if beforeDpItem > 0 {
			nowDpItem = nums[i] + beforeDpItem
		} else {
			nowDpItem = nums[i]
		}
		if nowDpItem > r {
			r = nowDpItem
		}
		beforeDpItem = nowDpItem
	}
	return r
}

//first
/*func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	dp := make([]int, l)
	dp[0] = nums[0]

	//这个时候不管第i-1项，只看第i项

	//第i项的最大值=a[i]+(第i-1项目 > 0 ? 第i-1项目 : 0)
	for i := 1; i < l; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
	}
	r := dp[0]
	for i := 1; i < l; i++ {
		if r < dp[i] {
			r = dp[i]
		}
	}
	return r
}*/
