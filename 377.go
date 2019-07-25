package main

import (
	"sort"
)

func main() {
	println(combinationSum4([]int{1, 2, 3}, 32))
	return
	println(combinationSum4([]int{1, 1, 1}, 2))
	println(combinationSum4([]int{}, 4))
	println(combinationSum4([]int{1, 2, 3}, 4))
	println(combinationSum4([]int{1, 1, 1, 1, 1, 1, 1, 1}, 4))
}


func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)

	return 0
}

/*
Given an integer array with all positive numbers and no duplicates,
find the number of possible combinations that add up to a positive integer target.

Example:

nums = [1, 2, 3]
target = 4

The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)

Note that different sequences are counted as different combinations.

Therefore the output is 7.


Follow up:
What if negative numbers are allowed in the given array?
How does it change the problem?
What limitation we need to add to the question to allow negative numbers?

Credits:
Special thanks to @pbrother for adding this problem and creating all test cases.
*/
