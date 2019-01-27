package main

func main() {
	println(numIslands([][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}))
	println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
	println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
}

func numIslands(grid [][]byte) int {
	l1 := len(grid)
	if l1 == 0 {
		return 0
	}
	l2 := len(grid[0])
	if l2 == 0 {
		return 0
	}
	tmp := make([][]bool, l1)
	for i := 0; i < l1; i++ {
		tmp[i] = make([]bool, l2)
	}
	r := 0
	for i, v := range grid {
		for i1, v1 := range v {
			if v1 == '0' {
				continue
			}
			if tmp[i][i1] {
				continue
			}
			_numIslands(&grid, &tmp, l1, l2, i, i1)
			r++
		}
	}
	return r
}

func _numIslands(grid *[][]byte, tmp *[][]bool, l1, l2, x, y int) {
	if y == l2 || x == l1 || y < 0 || x < 0 {
		return
	}
	if (*tmp)[x][y] {
		return
	}
	if (*grid)[x][y] == '1' {
		(*tmp)[x][y] = true
		_numIslands(grid, tmp, l1, l2, x, y+1)
		_numIslands(grid, tmp, l1, l2, x, y-1)
		_numIslands(grid, tmp, l1, l2, x+1, y)
		_numIslands(grid, tmp, l1, l2, x-1, y)
	}
}
