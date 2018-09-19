package main

func main() {
	_printArray(getRow(0))
	return

	_printArray(getRow(3))
	_printArray(getRow(4))
	_printArray(getRow(9))
	_printArray(getRow(10))
}

func getRow(rowIndex int) []int {
	r := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		r[i] = 1
	}
	num := 0
	for i := 2; i <= rowIndex; i++ {
		num = i / 2
		if i%2 == 0 {
			r[num] = r[num-1] + r[num-1]
			num--
		}
		for j := num; j >= 1; j-- {
			r[j] = r[j] + r[j-1]
		}
	}
	start := 0
	end := rowIndex
	for {
		if start >= end {
			break
		}
		r[end] = r[start]
		start++
		end--
	}

	return r
}

func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
