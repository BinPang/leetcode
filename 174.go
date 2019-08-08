package main

func main() {
	t0 := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	println(calculateMinimumHP(t0))
}

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 0
	}
	n := len(dungeon[0])
	if n == 0 {
		return 0
	}
	/*
		-2 (K)	-3	3
		-5	-10	1
		10	30	-5 (P)
	*/
	tmp := 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i < m-1 && j < n-1 {
				tmp = dungeon[i][j] + dungeon[i][j+1]
				if dungeon[i][j]+dungeon[i+1][j] > dungeon[i][j]+dungeon[i][j+1] {
					tmp = dungeon[i][j] + dungeon[i+1][j]
				}
				if tmp >= 0 {
					dungeon[i][j] = 0
				} else {
					dungeon[i][j] = tmp
				}
			} else if i < m-1 { //j == n-1
				if dungeon[i][j]+dungeon[i+1][j] >= 0 {
					dungeon[i][j] = 0
				} else {
					dungeon[i][j] += dungeon[i+1][j]
				}
			} else if j < n-1 { //i == m-1
				if dungeon[i][j]+dungeon[i][j+1] >= 0 {
					dungeon[i][j] = 0
				} else {
					dungeon[i][j] += dungeon[i][j+1]
				}
			} else {
				if dungeon[i][j] > 0 {
					dungeon[i][j] = 0
				}
			}
		}
	}
	if dungeon[0][0] == 0 {
		return 1
	} else {
		return 0 - dungeon[0][0] + 1
	}
}
