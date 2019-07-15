package main

func hammingDistance(x int, y int) int {
	r := 0
	t := x ^ y
	for t != 0 {
		if t&1 == 1 {
			r += 1
		}
		t >>= 1
	}
	return r
}
