package main

import "fmt"

func main() {
	println(maxPoints([][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}))

	println(maxPoints([][]int{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
	}))
}

func maxPoints(points [][]int) int {
	l := len(points)
	if l == 0 {
		return 0
	}
	dup := make(map[int]map[int]int, 0)
	for _, v := range points {
		if dup[v[0]] == nil {
			dup[v[0]] = make(map[int]int, 0)
		}
		dup[v[0]][v[1]] += 1
	}
	column := make(map[int]int, 0)
	row := make(map[int]int, 0)
	normal := make(map[int]map[int]map[int]map[int]int, 0)
	r := 0
	var x, y, gcd int
	var a, b int
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				column[points[i][0]] += 1
				if r < column[points[i][0]] {
					r = column[points[i][0]]
				}
				row[points[i][1]] += 1
				if r < row[points[i][1]] {
					r = row[points[i][1]]
				}
			} else if points[i][0] == points[j][0] {
				column[points[i][0]] += 1
				if r < column[points[i][0]] {
					r = column[points[i][0]]
				}
			} else if points[i][1] == points[j][1] {
				row[points[i][1]] += 1
				if r < row[points[i][1]] {
					r = row[points[i][1]]
				}
			} else {
				x = points[i][0] - points[j][0]
				y = points[i][1] - points[j][1]
				if x < 0 {
					x = -x
					y = -y
				}
				gcd = GCD(x, y)
				//slope = y/x
				x = x / gcd
				y = y / gcd
				if normal[y] == nil {
					normal[y] = make(map[int]map[int]map[int]int)
				}
				if normal[y][x] == nil {
					normal[y][x] = make(map[int]map[int]int)
				}
				//y = ax+b;==>
				a = points[j][1]*(points[j][0]-points[i][0]) - points[j][0]*(points[j][1]-points[i][1])
				b = points[j][0] - points[i][0]
				if a == 0 {
					b = 0
				} else {
					if a < 0 {
						a = -a
						b = -b
					}
					gcd = GCD(a, b)
					a = a / gcd
					b = b / gcd
				}
				if normal[y][x][a] == nil {
					normal[y][x][a] = make(map[int]int)
				}
				normal[y][x][a][b] += 1
				if r < normal[y][x][a][b] {
					r = normal[y][x][a][b]
				}
				println(fmt.Sprintf("%+v, %+v ,%d,%d,%d,%d", points[i], points[j], x/y, a, b, r))
			}
		}
	}
	//println(fmt.Sprintf("%+v, %+v, %+v", column, row, normal))
	return r + 1
}

func GCD(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}
