package main

func main() {
	println("input:[3,6,1,5,1,1,1,1,10,3],need 3,result:", jump([]int{3, 6, 1, 5, 1, 1, 1, 1, 10, 3}))
	println("input:[2,3,1,1,4],need 2,result:", jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	dp := make([]int, l) //init dp[0] = 0@TODO not dp, only incr one by one
	jumpTimes := 1
	startIndex := 1
	endIndex := nums[0]

	for {
		if startIndex >= l {
			break
		}
		nextEnd := 0
		//println("**", startIndex, endIndex, l)
		j := startIndex
		for {
			if j >= l {
				break
			}
			dp[j] = jumpTimes
			if j+nums[j] > nextEnd {
				nextEnd = j + nums[j]
			}
			if j == endIndex {
				break
			}
			j ++
		}
		//print("dp:")
		//_printArray(dp)
		jumpTimes ++
		startIndex = endIndex + 1
		endIndex = nextEnd
	}
	return dp[l-1]
}

func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
