package main

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	println(exist(board, "ABCCED"))

	board1 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	println(exist(board1, "SEE"))

	board2 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	println(exist(board2, "ABCB"))

	println(exist([][]byte{{'A'}}, "A"))

}

func exist(board [][]byte, word string) bool {
	l := len(word)
	if l == 0 {
		return false
	}
	l1 := len(board)
	if l1 == 0 {
		return false
	}
	l2 := len(board[0])
	if l2 == 0 {
		return false
	}
	for j := 0; j < l2; j++ {
		for i := 0; i < l1; i++ {
			//println("start to go", i, j)
			if gogo(board, &word, 0, i, j, l1, l2) {
				return true
			}
		}
	}
	return false
}

func gogo(board [][]byte, word *string, l int, x, y int, l1, l2 int) bool {
	//println(x, y, l)
	if l == len(*word) {
		return true
	}
	if x == l1 || y == l2 || x < 0 || y < 0 {
		return false
	}
	if board[x][y] == (*word)[l] {
		if y <= l2-1 { //right
			tmp := board[x][y]
			board[x][y] = 0
			if gogo(board, word, l+1, x, y+1, l1, l2) {
				return true
			} else {
				board[x][y] = tmp
			}
		}
		if y >= 0 { //left
			tmp := board[x][y]
			board[x][y] = 0
			if gogo(board, word, l+1, x, y-1, l1, l2) {
				return true
			} else {
				board[x][y] = tmp
			}

		}
		if x <= l1-1 { //bottom
			tmp := board[x][y]
			board[x][y] = 0
			if gogo(board, word, l+1, x+1, y, l1, l2) {
				return true
			} else {
				board[x][y] = tmp
			}
		}
		if x >= 0 { //top
			tmp := board[x][y]
			board[x][y] = 0
			if gogo(board, word, l+1, x-1, y, l1, l2) {
				return true
			} else {
				board[x][y] = tmp
			}
		}
	}
	return false
}
