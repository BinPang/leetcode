package main

func twoSum(nums []int, target int) []int {
	//tmp := map[int]int{}
	tmp := make(map[int]int, len(nums))
	for key, value := range nums {
		if location, ok := tmp[target-value]; ok {
			return []int{location, key}
		}
		tmp[value] = key
	}
	panic(target)
}