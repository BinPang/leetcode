package main

import "sort"

func main() {
	t6 := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	_printArrayArray(t6)
	println("")

	t5 := combinationSum2([]int{1, 2, 3}, 6)
	_printArrayArray(t5)
	println("")

	t4 := combinationSum2([]int{1, 1, 3, 3, 4, 4}, 10)
	_printArrayArray(t4)
	println("")

	t3 := combinationSum2([]int{1, 1, 1, 1, 1}, 9)
	_printArrayArray(t3)
	println("")

	//							0  1  2  3  4  5  6  7  8  9  10 11 12 13
	t1 := combinationSum2([]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 5, 6, 7, 10}, 14)
	_printArrayArray(t1)
	println("")

	t0 := combinationSum2([]int{1, 1, 2, 5, 6, 7, 10}, 8)
	_printArrayArray(t0)
	println("")

	t2 := combinationSum2([]int{2, 5, 2, 1, 2}, 5)
	_printArrayArray(t2)
	println("")
}

//1 1 2 5 6 7 10@TODO stupid method, this method is coding by human's activity
// @TODO hard to coding. use backtracking
/**
1 1 2 5 no
1 1 5 6 no
1 1 6 yes
1 2 5 yes
1 5 6 no
1 6 7 no
1 7 yes
2 5 6 no
2 6 yes
5 6 no
6 7 no
7 10 no
10 no
 */
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return _combinationSum2(candidates, target, 0, len(candidates))
}

func _combinationSum2(candidates []int, target int, start, end int) [][]int {
	r := make([][]int, 0)
	for i := start; i < end; i++ {
		if i != start && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] >= target {
			if candidates[i] == target {
				r = append(r, []int{candidates[i]})
			}
			break
		}
		tmp := _combinationSum2(candidates, target-candidates[i], i+1, end)
		if len(tmp) > 0 {
			for _, v := range tmp {
				r = append(r, append([]int{candidates[i]}, v...))
			}
		}
	}
	return r
}

func _printArrayArray(t [][]int) {
	for _, v := range t {
		_printArray(v)
	}
}
func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
