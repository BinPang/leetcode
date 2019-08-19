package main

func islandPerimeter(grid [][]int) int {
	r := 0
	l1 := len(grid)
	if l1 == 0 {
		return 0
	}
	l2 := len(grid[0])
	if l2 == 0 {
		return 0
	}
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if grid[i][j] == 0 {
				continue
			}
			//up, right, bottom, left
			if i == 0 || grid[i-1][j] == 0 {
				r++
			}
			if j == l2-1 || grid[i][j+1] == 0 {
				r++
			}
			if i == l1-1 || grid[i+1][j] == 0 {
				r++
			}
			if j == 0 || grid[i][j-1] == 0 {
				r++
			}
		}
	}
	return r
}
