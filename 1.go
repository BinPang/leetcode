package main

func twoSum(nums []int, target int) []int {
	tmp := map[int]int{}
	for key, value := range nums {
		if exist, ok := tmp[target-value]; ok {
			return []int{key, exist}
		}
		tmp[value] = key
	}
	panic(target)
}