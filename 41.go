package main

func main() {

	println(firstMissingPositive([]int{1, 1}), "->", 2)
	//return
	println(firstMissingPositive([]int{3, 0, 1, -1}), "->", 2)
	println(firstMissingPositive([]int{1, 2, 0}), "->", 3)
	println(firstMissingPositive([]int{7, 8, 9, 11, 12}), "->", 1)
	println(firstMissingPositive([]int{}), "->", 1)
	println(firstMissingPositive([]int{1}), "->", 2)
}

func firstMissingPositive(nums []int) int {
	l := len(nums)
	j := 0
	k := 0
	for i := 0; i < l; i++ {
		if nums[i] == i+1 {
			continue
		}
		j = i
		k = 0
		//println("start")
		for {
			k, nums[j] = nums[j], k
			j = k-1
			//println("after", k, j)

			if j >= l || j < 0 || nums[j] ==k {
				break
			}
		}
		//println("end")
		//for j:=0;j<l ;j++  {
		//	print("loop:", nums[j], ",")
		//}
		//println("")
	}

	//for _,v := range nums {
	//	print(v, ",")
	//}
	//println("")

	for i, v := range nums {
		if i+1 != v {
			return i + 1
		}
	}
	return l + 1
}
