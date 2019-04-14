package main

//maybe math trick
//toooooooooooooooo difficult for me to solve, give up. just copy from discuss
func singleNumber(nums []int) int {
	ones := 0
	twos := 0

	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}
