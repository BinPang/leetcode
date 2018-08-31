package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	start1 := 1
	start2 := 2
	for i := 1; i <= n; i++ {
		if i <= 2 {
			continue
		}
		start2 = start1 + start2
		start1 = start2 - start1
	}
	return start2
}
