package main

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	println("need 6,return:", maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	sum := make([][]int, m)

	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
	}
	r := 0
	v := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				v = 1
			} else {
				v = 0
			}
			if i == 0 && j == 0 {
				sum[i][j] = v
				r = v
			} else if j == 0 { //first column
				if v == 1 {
					sum[i][j] = sum[i-1][j] + 1
				}
			} else if i == 0 {
				if v == 1 {
					sum[i][j] = sum[i][j-1] + 1
				}
			} else {
				if v == 1 {
					if sum[i-1][j-1] > 0 && sum[i][j-1] > 0 && sum[i-1][j] > 0 {
						sum[i][j] = sum[i-1][j-1] + 3
					} else {
						sum[i][j] = 1
					}
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			println(i, j, ":", sum[i][j])
			if r < sum[i][j] {
				r = sum[i][j]
			}
		}
	}
	return r
}
