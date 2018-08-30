package main

func main() {

}
//no test case, the same with 62.go 63.go
func minPathSum(grid [][]int) int {
	for i, v := range grid {
		for j := range v {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] += grid[i][j-1]
				} else {
					grid[i][j] += grid[i-1][j]
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
