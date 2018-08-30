package main

func main() {
	println("input is:3, 2.need 3,result:", uniquePaths(3, 2))
	println("input is:7, 3.need 28,result:", uniquePaths(7, 3))
	println("input is:1,1.need 1,result:", uniquePaths(1, 1))
}

func uniquePaths(m int, n int) int {
	if m <= 1 || n <= 1 {
		return 1
	}
	path := make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				path[i][j] = 1
			} else if j == 0 {
				path[i][j] = 1
			} else {
				path[i][j] = path[i-1][j] + path[i][j-1]
			}
		}
	}
	return path[m-1][n-1]
}
