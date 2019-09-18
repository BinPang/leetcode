package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	index := 0
	l := len(s)
	r := 0
	for _, v := range g {
		for index < l {
			if v <= s[index] {
				r += 1
				index++
				break
			}
			index++
		}
		if index == l {
			break
		}
	}
	return r
}
