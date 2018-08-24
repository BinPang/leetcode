package main

func main() {
	println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
	println(removeDuplicates([]int{0,0,0, 0, 0, 0}))
	println(removeDuplicates([]int{0}))
	println(removeDuplicates([]int{}))
}

func removeDuplicates(nums []int) int {
	if len(nums) <=1 {
		return len(nums)
	}
	start := 0
	index := 0
	for i, v := range nums{
		if i == 0 {
			start = nums[i]
			index = i+1
			continue
		}
		if v != start {
			nums[index] =v
			index +=1
			start = v
		}
	}
	//for i:=0;i<index;i++ {
	//	println(i, ",", nums[i])
	//}
	//println(index)
	return index
}
