package main

func main() {
	t0 := plusOne([]int{1, 2, 3})
	_printArray(t0)

	t1 := plusOne([]int{9, 9})
	_printArray(t1)

	t2 := plusOne([]int{0})
	_printArray(t2)

}

func plusOne(digits []int) []int {
	plus := 0
	tmp := 0
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if i == l-1 {
			tmp = digits[i] + 1
		} else {
			tmp = digits[i] + plus
		}
		if tmp >= 10 {
			plus = 1
			digits[i] = tmp - 10
		} else {
			digits[i] = tmp
			plus = 0
		}
	}
	if plus == 0 {
		return digits
	} else {
		r := make([]int, len(digits)+1)
		r = append([]int{1}, digits...)
		return r
	}
}

func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
