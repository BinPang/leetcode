package main

/**
Given an integer matrix, find the length of the longest increasing path.

From each cell, you can either move to four directions: left, right, up or down.
You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).

Example 1:

Input: nums =
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
Output: 4
Explanation: The longest increasing path is [1, 2, 6, 9].
Example 2:

Input: nums =
[
  [3,4,5],
  [3,2,6],
  [2,2,1]
]
Output: 4
Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
*/

func main() {
	d1 := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	println(longestIncreasingPath(d1))

	d2 := [][]int{
		{9, 9, 4, 2},
		{6, 6, 8, 7},
		{3, 2, 1, 5},
	}
	println(longestIncreasingPath(d2))
}

func longestIncreasingPath(matrix [][]int) int {
	r := 0
	l := len(matrix)
	if l == 0 {
		return 0
	}
	l1 := len(matrix[0])
	if l1 == 0 {
		return 0
	}
	tmp := make([][]int, l)
	for i := 0; i < l; i++ {
		tmp[i] = make([]int, l1)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l1; j++ {
			_longestIncreasingPath(&matrix, &tmp, l-1, l1-1, i, j)
			r = max(r, tmp[i][j])
		}
	}
	return r
}

func _longestIncreasingPath(matrix *[][]int, tmp *[][]int, l, l1, x, y int) {
	if (*tmp)[x][y] != 0 {
		return
	}
	tmpR := 1
	//right, bottom, left, top
	if y < l1 && (*matrix)[x][y] < (*matrix)[x][y+1] {
		if (*tmp)[x][y+1] == 0 {
			_longestIncreasingPath(matrix, tmp, l, l1, x, y+1)
		}
		tmpR = max(tmpR, (*tmp)[x][y+1]+1)
	}
	if x < l && (*matrix)[x][y] < (*matrix)[x+1][y] {
		if (*tmp)[x+1][y] == 0 {
			_longestIncreasingPath(matrix, tmp, l, l1, x+1, y)
		}
		tmpR = max(tmpR, (*tmp)[x+1][y]+1)
	}
	if y > 0 && (*matrix)[x][y] < (*matrix)[x][y-1] {
		if (*tmp)[x][y-1] == 0 {
			_longestIncreasingPath(matrix, tmp, l, l1, x, y-1)
		}
		tmpR = max(tmpR, (*tmp)[x][y-1]+1)
	}
	if x > 0 && (*matrix)[x][y] < (*matrix)[x-1][y] {
		if (*tmp)[x-1][y] == 0 {
			_longestIncreasingPath(matrix, tmp, l, l1, x-1, y)
		}
		tmpR = max(tmpR, (*tmp)[x-1][y]+1)
	}
	(*tmp)[x][y] = tmpR
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
