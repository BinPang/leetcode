package main

import (
	"fmt"
	"strconv"
)

func main() {
	println(fmt.Sprintf("%+v", summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})))
	println(fmt.Sprintf("%+v", summaryRanges([]int{0, 2, 3, 4, 6, 9})))
	println(fmt.Sprintf("%+v", summaryRanges([]int{0, 1, 2, 3, 4, 6, 7, 9})))
}

/**
Input:  [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: 2,3,4 form a continuous range; 8,9 form a continuous range.
*/
func summaryRanges(nums []int) []string {
	l := len(nums)
	if l == 0 {
		return []string{}
	}
	r := make([]string, 0)
	start := nums[0]
	end := nums[0]
	index := 1
	for {
		if index == l {
			if start == end {
				r = append(r, strconv.Itoa(start))
			} else {
				r = append(r, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}
			break
		}
		if nums[index] == end+1 {
			end++
		} else {
			//find item
			if start == end {
				r = append(r, strconv.Itoa(start))
			} else {
				r = append(r, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}

			start = nums[index]
			end = start
		}
		index++
	}

	return r
}
