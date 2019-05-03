package main

/*
Consider the following matrix:

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
Given target = 5, return true.

Given target = 20, return false.
//find 15
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return _searchMatrix(&matrix, target, 0, 0, len(matrix), len(matrix[0]))
}

func _searchMatrix(matrix *[][]int, target int, startX, startY, endX, endY int) bool {
	pointX := startX
	pointY := startY
	for {
		if (*matrix)[pointX][pointY] < target {
			pointX++
			pointY++
		} else if (*matrix)[pointX][pointY] == target {
			return true
		} else {
			return _searchMatrix(matrix, target, pointX, 0, pointX-1, endY) ||
				_searchMatrix(matrix, target, startX, pointY-1, pointX-1, endY)
		}
	}
}
