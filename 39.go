package main

import "sort"

func main() {
	t0 := combinationSum([]int{2,3,6,7}, 7)
	_printArrayArray(t0)
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	l := len(candidates)
	stack := make([]int, l)
	num := make([]int, l)
	s:=0
	n := 0
	for {
		if n == 0 {
			num[n] = candidates[s]
		} else {
			n+=1
			num[n] = num[n-1]+candidates[s]
		}
		if num[n] > target {
			
		} else if num[n] == target {

		} else {

		}
	}
}


func _printArrayArray(t [][]int) {
	for _, v := range t{
		_printArray(v)
	}
}
func _printArray(t []int)  {
	for _, v := range t{
		print(v, ",")
	}
	println("")
}
