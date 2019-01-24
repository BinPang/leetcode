package main

import (
	"sort"
	"strconv"
)

func main() {
	println(largestNumber([]int{0, 0}))
	println(largestNumber([]int{3, 30, 34, 5, 9}))
}

type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return strconv.Itoa(c[i])+strconv.Itoa(c[j]) > strconv.Itoa(c[j])+strconv.Itoa(c[i])
}

func largestNumber(nums []int) string {
	nn := make(IntSlice, len(nums))
	for i, v := range nums {
		nn[i] = v
	}
	sort.Sort(nn)
	if nn[0] == 0 {
		return "0"
	}
	r := ""
	for _, v := range nn {
		r += strconv.Itoa(v)
	}
	return r
}
