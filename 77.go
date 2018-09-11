package main

func main() {
	t0 := combine(4, 2)
	_printArrayArray(t0)
}

func combine(n int, k int) [][]int {
	return _combine(1, n, k)
}

func _combine(start, end int, k int) [][]int {
	r := make([][]int, 0)
	if k == 0 {
		return nil
	} else if k == 1 {
		for i := start; i <= end; i++ {
			r = append(r, append([]int{i}))
		}
	} else {
		for i := end - k + 1; i >= start; i-- {
			tmp := _combine(i+1, end, k-1)
			if tmp != nil {
				for _, v := range tmp {
					r = append(r, append([]int{i}, v...))
				}
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
