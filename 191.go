package main

import "strconv"

func main() {
	println(strconv.FormatInt(25, 2))
	println(hammingWeight(25))
}

func hammingWeight(num uint32) int {
	r := 0
	for num != 0 {
		if num&1 == 1 {
			r += 1
		}
		num = num >> 1
	}
	return r
}
