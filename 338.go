package main

import "fmt"

func main() {
	println(fmt.Sprintf("%+v", countBits(5)))
}

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	} else if num == 1 {
		return []int{0, 1}
	}
	r := make([]int, num+1)
	r[0], r[1] = 0, 1
	for i := 2; i <= num; i++ {
		r[i] = r[i>>1] + i&1
	}
	return r
}
