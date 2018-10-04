package main

func main() {
	t0 := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'X', 'X'},
		{'O', 'O', 'O', 'X'},
	}
	solve(t0)
}

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}
	for _, i := range []int{0, m - 1} {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				mark(board, m, n, i, j)
			}
		}
	}
	for _, j := range []int{0, n - 1} {
		for i := 0; i < m; i++ {
			if board[i][j] == 'O' {
				mark(board, m, n, i, j)
			}
		}
	}
	for i1, v1 := range board {
		for i2, v2 := range v1 {
			if v2 == 'O' {
				board[i1][i2] = 'X'
			} else if v2 == 'o' {
				board[i1][i2] = 'O'
			}
			//print(string(board[i1][i2]), ",")
		}
		//println("")
	}
}

func mark(board [][]byte, m, n, i, j int) {
	board[i][j] = 'o'
	//left
	if j > 0 && board[i][j-1] == 'O' {
		mark(board, m, n, i, j-1)
	}
	//right
	if j < n-1 && board[i][j+1] == 'O' {
		mark(board, m, n, i, j+1)
	}
	//top
	if i > 0 && board[i-1][j] == 'O' {
		mark(board, m, n, i-1, j)
	}
	//bottom
	if i < m-1 && board[i+1][j] == 'O' {
		mark(board, m, n, i+1, j)
	}
}
