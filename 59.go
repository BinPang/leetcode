package main

func main() {
	t0 := generateMatrix(5)
	_printArrayArray(t0)

	t1 := generateMatrix(4)
	_printArrayArray(t1)

}

func generateMatrix(n int) [][]int {
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, n)
	}
	index := 1
	startX, startY, endX, endY := 0, 0, n-1, n-1
	for {
		for i := startY; i <= endY; i++ {
			r[startX][i] = index
			index ++
		}
		for i := startX + 1; i <= endX; i++ {
			r[i][endY] = index
			index++
		}
		for i := endY - 1; i >= startY; i-- {
			r[endX][i] = index
			index++
		}
		for i := endX - 1; i > startX; i-- {
			r[i][startY] = index
			index++
		}
		startX++
		endX--
		startY++
		endY--
		if startX > endX {
			break
		}
	}
	return r
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
