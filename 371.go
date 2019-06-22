package main

import "strconv"

func main() {
	println(strconv.FormatInt(10, 2))
	println(strconv.FormatInt(-10, 2))

	t := -5
	println(t <<1)
	println(t >>1)
	println(t >>2)
	println(t >>3)
	println(t >>65)
}

func getSum(a int, b int) int {

	return 0
}

/*
Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

Example 1:

Input: a = 1, b = 2
Output: 3
Example 2:

Input: a = -2, b = 3
Output: 1
 */