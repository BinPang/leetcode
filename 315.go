package main

import (
	"fmt"
	"sort"
)

func main() {
	countSmaller([]int{5, 2, 6, 1})
}

func countSmaller(nums []int) []int {
	l := len(nums)
	n := make([]int, l)
	s := make(map[int]int)
	for e := range nums {
		n[e] = nums[e]
	}
	sort.Ints(n)
	for i := l - 1; i >= 0; i-- {
		if s[n[i]] == 0 {
			s[n[i]] = i
		}
	}

	println(fmt.Sprintf("%+v,%+v", nums, n))
	return nil
}

/*
You are given an integer array nums and you have to return a new counts array.
The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].

Example:

Input: [5,2,6,1]
Output: [2,1,1,0]
Explanation:
To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.

*/
