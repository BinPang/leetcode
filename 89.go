package main

/**
[0]
[0,1]
[0,1,3,2]
[0,1,3,2,6,7,5,4]
[0,1,3,2,6,7,5,4,12,13,15,14,10,11,9,8]
[0,1,3,2,6,7,5,4,12,13,15,14,10,11,9,8,24,25,27,26,30,31,29,28,20,21,23,22,18,19,17,16]
[0,1,3,2,6,7,5,4,12,13,15,14,10,11,9,8,24,25,27,26,30,31,29,28,20,21,23,22,18,19,17,16,48,49,51,50,54,55,53,52,60,61,63,62,58,59,57,56,40,41,43,42,46,47,45,44,36,37,39,38,34,35,33,32]
 */
func main() {
	_printArray(grayCode(0))
	_printArray(grayCode(1))
	_printArray(grayCode(4))
	_printArray(grayCode(5))
	_printArray(grayCode(6))
}

func grayCode(n int) []int {
	if n < 0 {
		panic("not exist")
	}
	total := 1
	for i := 1; i <= n; i++ {
		total *= 2
	}
	r := make([]int, total)
	r[0] = 0
	end := 0
	total = 1
	for i := 1; i <= n; i++ {
		if i != 1 {
			total *= 2
		}
		firstBitOne := 1 << uint(i-1)
		for j := 0; j < total; j++ {
			r[end+j+1] = r[end-j] | firstBitOne
		}
		end += total
	}
	return r
}

func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
