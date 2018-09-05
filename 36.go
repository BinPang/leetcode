package main

func main() {
	input := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	println("result:", isValidSudoku(input))
}

func isValidSudoku(board [][]byte) bool {
	a := make([]int, 9)
	item := byte(99)
	//row
	for i := 0; i < 9; i++ {
		_newArray(a)
		for j := 0; j < 9; j++ {
			item = board[i][j]
			if item == 46 {
				continue
			}
			if a[item-49] != 0 {
				return false
			} else {
				a[item-49] = 1
			}
		}
	}
	//column
	for i := 0; i < 9; i++ {
		_newArray(a)
		for j := 0; j < 9; j++ {
			item = board[j][i]
			if item == 46 {
				continue
			}
			if a[item-49] != 0 {
				return false
			} else {
				a[item-49] = 1
			}
		}
	}
	//3*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			_newArray(a)
			for i1 := 0; i1 < 3; i1++ {
				for j1 := 0; j1 < 3; j1++ {
					item = board[i*3+i1][j*3+j1]
					if item == 46 {
						continue
					}
					if a[item-49] != 0 {
						return false
					} else {
						a[item-49] = 1
					}
				}
			}
		}
	}

	return true
}
func _newArray(a []int) {
	for i := range a {
		a[i] = 0
	}
}
