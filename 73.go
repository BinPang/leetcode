package main

func setZeroes(matrix [][]int)  {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])
	if n == 0 {
		return
	}
	firstRowHaveZero := false
	firstColumnHaveZero := false
	i, j := 0, 0
	for i = 0; i < n; i++ {
		if matrix[0][i] == 0 {
			firstRowHaveZero = true
		}
	}
	for i = 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColumnHaveZero = true
		}
	}
	for i = 1; i < m; i++ {
		for j = 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i = 1; i < m; i++ {
		for j = 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0{
				matrix[i][j] = 0
			}
		}
	}
	if firstRowHaveZero {
		for i = 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if firstColumnHaveZero {
		for i = 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
