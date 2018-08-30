package main

func main() {
	t:= [][]int{
		{0,0,0},
		{0,1,0},
		{0,0,0},
	}
	println("need 2,return:", uniquePathsWithObstacles(t))
	t1 := [][]int{
		{1},
	}
	println("need 0,return:", uniquePathsWithObstacles(t1))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	path := make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i==0 && j== 0 {
				if obstacleGrid[i][j] == 0 {
					path[i][j] = 1
				}
				continue
			}
			if i == 0 {
				if path[i][j-1] == 1 && obstacleGrid[i][j] == 0 {
					path[i][j] = 1
				}
			} else if j == 0 {
				if path[i-1][j] == 1 && obstacleGrid[i][j] == 0 {
					path[i][j] = 1
				}
			} else {
				if obstacleGrid[i][j] == 0 {
					path[i][j] = path[i-1][j] + path[i][j-1]
				}
			}
		}
	}
	return path[m-1][n-1]
}