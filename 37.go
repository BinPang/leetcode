package main

func main() {
	t0 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(t0)
	for _, v1 := range t0 {
		for _, v2 := range v1 {
			print(string(v2), ",")
		}
		println("")
	}
}

func solveSudoku(board [][]byte) {
	_solveSudoku(board)
}

func _solveSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//println("start,i:", i, ",j:", j)
			if board[i][j] == '.' {
				for k := 1; k <= 9; k++ {
					board[i][j] = byte(k)+48//byte 48 == int 0, int to byte
					//println("deal,i:", i, ",j:", j, ",k:", k, board[i][j])
					if isValid(board, i, j) && _solveSudoku(board) {
						return true
					}
					board[i][j] = '.'
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, x, y int) bool {
	for i := 0; i < 9; i++ {
		if i != y && board[x][i] == board[x][y] {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if i != x && board[i][y] == board[x][y] {
			return false
		}
	}
	for i := 3 * (x / 3); i < 3*(x/3)+3; i++ {
		for j := 3 * (y / 3); j < 3*(y/3)+3; j++ {
			if x != i && y != j && board[i][j] == board[x][y] {
				return false
			}
		}
	}
	return true
}
