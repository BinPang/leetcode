package main

func main() {
	println(searchMatrix([][]int{
		{1},
	}, 1))

	println(searchMatrix([][]int{
		{1},
	}, 0))

	println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}, 11))

	println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}, 13))

	println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}, 3))

}

func searchMatrix(matrix [][]int, target int) bool {
	l1 := len(matrix)
	if l1 == 0 {
		return false
	}
	l2 := len(matrix[0])
	if l2 == 0 {
		return false
	}
	start := 0
	end := l1
	findLine := 0
	for {
		//println("loop0:", start, ",", end)
		if start+1 == end {
			findLine = start
			break
		}
		findLine = (start + end) / 2
		if matrix[findLine][0] > target {
			end = findLine
		} else if matrix[findLine][0] < target {
			start = findLine
		} else {
			return true
		}
	}
	//println("find line:", findLine)
	start = 0
	end = l2
	middle := 0
	for {
		//println("loop1:", start, ",", end)
		if start+1 == end {
			if matrix[findLine][start] == target {
				return true
			}
			break
		}
		middle = (start + end) / 2
		if matrix[findLine][middle] == target {
			return true
		} else if matrix[findLine][middle] < target {
			start = middle
		} else {
			end = middle
		}
	}
	return false
}
