package main

func main() {
	println(countBattleships([][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}))
}

func countBattleships(board [][]byte) int {
	r := 0
	l := len(board)
	if l == 0 {
		return 0
	}
	l1 := len(board[0])
	if l1 == 0 {
		return 0
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l1; j++ {
			if board[i][j] == 'X' {
				if (j > 0 && board[i][j-1] == 'X') || (i > 0 && board[i-1][j] == 'X') {
				} else {
					r += 1
				}
			}
		}
	}
	return r
}
