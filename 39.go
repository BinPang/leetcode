package main

import "sort"

func main() {
	t0 := combinationSum([]int{2, 3, 6, 7}, 7)
	_printArrayArray(t0)
	println("")

	t1 := combinationSum([]int{2, 3, 5}, 8)
	_printArrayArray(t1)
	println("")

	t3 := combinationSum([]int{5, 6, 9, 18}, 18)
	_printArrayArray(t3)
	println("")

	t2 := combinationSum([]int{7, 3}, 17)
	_printArrayArray(t2)
	println("")
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return _combinationSum(candidates, target, 0, len(candidates))
}

func _combinationSum(candidates []int, target int, start, end int) [][]int {
	r := make([][]int, 0)
	for i := start; i < end; i++ {
		if candidates[i] >= target {
			if candidates[i] == target {
				r = append(r, []int{candidates[i]})
			}
			break
		}
		tmp := _combinationSum(candidates, target-candidates[i], i, end)
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
