package main

func main() {
	println(totalNQueens(5))
}

func totalNQueens(n int) int {
	table := make([][]bool, n)
	for i := 0; i < n; i++ {
		table[i] = make([]bool, n)
	}
	r := 0
	startX := 0
	startY := 0
	needBack := false
	i := 0
	//debugN := 0
	for {
		if needBack {
			//find previous row
			startX--
			if startX < 0 {
				break //if n == 1
			}
			i = 0
			for {
				if table[startX][i] == true {
					break
				}
				if i == n-1 {
					break
				}
				i++
			}
			if startX == 0 && i == n-1 {
				break
			}
			table[startX][i] = false
			if i != n-1 {
				needBack = false
				startY = i + 1
			} else {
				needBack = true //dup notice
				continue        //debug got if n == 5
			}
		}

		table[startX][startY] = true
		//println("ok:", startX, startY, isOk(&table, startX, startY, n))
		//printTable(&table)
		if isOk(&table, startX, startY, n) {
			if startX == n-1 { //find it
				r += 1
				//println("gooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooool")
				table[startX][startY] = false
				needBack = true
			} else {
				startX++
				startY = 0
			}
		} else {
			table[startX][startY] = false
			if startY != n-1 {
				startY++
			} else {
				needBack = true
			}
		}
		//debugN += 1
		//if debugN > 200 {
		//	break
		//}
	}
	return r
}

func isOk(table *[][]bool, x, y, n int) bool {
	for i := 0; i < n; i++ {
		if i == y {
			continue
		}
		if (*table)[x][i] { //col
			return false
		}
	}
	for i := 0; i < n; i++ {
		if i == x {
			continue
		}
		if (*table)[i][y] { //row
			return false
		}
	}

	i := 1
	for { //left top
		if x-i < 0 || y-i < 0 {
			break
		}
		if (*table)[x-i][y-i] {
			return false
		}
		i++
	}
	i = 1
	for { //right bottom
		if x+i == n || y+i == n {
			break
		}
		if (*table)[x+i][y+i] {
			return false
		}
		i++
	}
	i = 1
	for { //left bottom
		if x-i < 0 || y+i == n {
			break
		}
		if (*table)[x-i][y+i] {
			return false
		}
		i++
	}
	i = 1
	for { //right top
		if x+i == n || y-i < 0 {
			break
		}
		if (*table)[x+i][y-i] {
			return false
		}
		i++
	}
	return true
}
