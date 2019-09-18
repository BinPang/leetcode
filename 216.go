package main

import "fmt"

func main() {
	println(fmt.Sprintf("%+v", combinationSum3(2, 18)))
	println(fmt.Sprintf("%+v", combinationSum3(3, 7)))
	println(fmt.Sprintf("%+v", combinationSum3(3, 9)))
}

func combinationSum3(k int, n int) [][]int {
	if k < 1 {
		return nil
	}
	return _combinationSum3(k, 1, n)
}

func _combinationSum3(k int, start int, n int) [][]int {
	println(k, start, n)
	if k == 1 {
		if n > 0 && start <= n && n < 10 {
			return [][]int{{n}}
		} else {
			return nil
		}
	}
	r := make([][]int, 0)
	for i := start; i <= 9; i++ {
		tmp := _combinationSum3(k-1, i+1, n-i)
		if tmp != nil {
			for _, v := range tmp {
				r = append(r, append([]int{i}, v...))
			}
		}
	}
	return r
}
