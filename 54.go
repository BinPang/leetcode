package main

func main() {
	o1 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	n1 := spiralOrder(o1)
	_printArray(n1)

	o2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	n2 := spiralOrder(o2)
	_printArray(n2)

	o3:= [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	n3 := spiralOrder(o3)
	_printArray(n3)
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	} else if m == 1 {
		return matrix[0]
	}
	n := len(matrix[0])
	size := m * n
	r := make([]int, size)
	index := 0
	startX, startY, endX, endY := 0, 0, m-1, n-1
	for {
		for i := startY; i <= endY; i++ {
			if index >= size {
				break
			}
			r[index] = matrix[startX][i]
			index ++
		}
		for i := startX + 1; i <= endX; i++ {
			if index >= size {
				break
			}
			r[index] = matrix[i][endY]
			index++
		}
		for i := endY - 1; i >= startY; i-- {
			if index >= size {
				break
			}
			r[index] = matrix[endX][i]
			index++
		}
		for i := endX - 1; i > startX; i-- {
			if index >= size {
				break
			}
			r[index] = matrix[i][startY]
			index++
		}
		startX++
		endX--
		startY++
		endY--
		if index >= size {
			break
		}
	}
	return r
}

func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
