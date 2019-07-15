package main

func main() {
	t := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	println(searchMatrix(t, 0))
	println(searchMatrix(t, 18))
	println(searchMatrix(t, 20))
	println(searchMatrix(t, 30))
	println(searchMatrix(t, 31))
}

func searchMatrix(matrix [][]int, target int) bool {
	l := len(matrix)
	if l == 0 {
		return false
	}
	l1 := len(matrix[0])
	startX, startY := 0, l1-1
	for {
		if startX == l || startY < 0 {
			return false
		}
		if matrix[startX][startY] == target {
			return true
		} else if matrix[startX][startY] < target {
			startX++
		} else {
			startY--
		}
	}
}
