package main

import "sort"

func main() {
	t5 := combinationSum2([]int{1, 2, 3}, 6)
	_printArrayArray(t5)
	println("")
	return

	t4 := combinationSum2([]int{1, 1, 3, 3, 4, 4}, 10)
	_printArrayArray(t4)
	println("")
	return

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

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	return _combinationSum2(candidates, 0, len(candidates), target)
}

func _combinationSum2(candidates []int, start, l, target int) [][]int {
	println(start, l, target)
	if start >= l {
		return nil
	}
	if target <= 0 {
		return nil
	}
	if candidates[start] > target {
		return nil
	} else if candidates[start] == target {
		return [][]int{{target}}
	}
	sum := 0
	r := make([][]int, 0)
	for i := start+1; i < l; i++ {
		sum += candidates[i]
		tmp := _combinationSum2(candidates, i, l, target-sum)
		println(i+1, target-sum)
		if tmp != nil {
			for _, v1 := range tmp {
				r = append(r, append([]int{candidates[i]}, v1...))
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
