package main

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1
	sum := 0
	for {
		sum = numbers[start] + numbers[end]
		if sum == target {
			return []int{start+1, end+1}
		} else if sum > target {
			end--
		} else {
			start++
		}
	}
}
