package main

func distributeCandies(candies []int) int {
	l := len(candies)
	tmp := make(map[int]bool, l)
	for _, v := range candies {
		if _, ok := tmp[v]; !ok {
			tmp[v] = true
		}
	}
	if len(tmp) > l>>1 {
		return l >> 1
	} else {
		return len(tmp)
	}
}
