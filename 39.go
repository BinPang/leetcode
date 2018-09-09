package main

import "sort"

func main() {
	t3 := combinationSum([]int{5, 6, 9, 18}, 18)
	_printArrayArray(t3)
	println("")

	t0 := combinationSum([]int{2, 3, 6, 7}, 7)
	_printArrayArray(t0)
	println("")

	t2 := combinationSum([]int{7, 3}, 17)
	_printArrayArray(t2)
	println("")

	t1 := combinationSum([]int{2, 3, 5}, 8)
	_printArrayArray(t1)
	println("")

}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	l := len(candidates)
	stackSum := make([]int, 0)
	stackPoint := make([]int, 0)
	p := 0
	pc := 0
	loopEmpty := false
	r := make([][]int, 0)

	//breakLoopDebug := 0
	for {
		if p == len(stackSum) {
			stackSum = append(stackSum, 0)
			stackPoint = append(stackPoint, 0)
		}
		if p == 0 {
			stackSum[p] = candidates[pc]
			stackPoint[p] = pc
		} else {
			stackSum[p] = stackSum[p-1] + candidates[pc]
			stackPoint[p] = pc
		}
		if stackSum[p] >= target {
			//for j := 0; j <= p; j++ {
			//	if j == 0 {
			//		println("j:", j, ",item:", stackSum[j], ",point:", stackPoint[j])
			//	} else {
			//		println("j:", j, ",item:", stackSum[j]-stackSum[j-1], ",point:", stackPoint[j])
			//	}
			//}
			if stackSum[p] == target {
				//got it
				//println("got it", p)
				tmp := make([]int, p+1)
				for j := 0; j <= p; j++ {
					tmp[j] = candidates[stackPoint[j]]
				}
				r = append(r, tmp)
			}
			//println("start:p,pc,loopEmpty:", p, pc, loopEmpty)
			for {
				p--
				if p < 0 {
					p = 0
					loopEmpty = true
					break
				}
				if stackPoint[p]+1 < l {
					pc = stackPoint[p] + 1
					break
				}
			}
			//println("end:p,pc,loopEmpty:", p, pc, loopEmpty)
			if loopEmpty {
				break
			}
		} else {
			p++
		}
		//breakLoopDebug ++
		//if breakLoopDebug > 20 {
		//	println("debug loop reach 20")
		//	break
		//}
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
