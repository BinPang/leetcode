package main

type Point struct {
	X int
	Y int
}

func main() {
	maxPoints([]Point{
		{X:1,Y:1},
		{X:3,Y:2},
		{X:5,Y:3},
		{X:4,Y:1},
		{X:2,Y:3},
		{X:1,Y:4},
	})
}

func maxPoints(points []Point) int {
	l := len(points)
	column := map[int]int{}
	row := map[int]int{}
	normal := map[int]map[int]map[int]int{}
	var x, y, gcd int
	var x1, y1 int
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if points[i].X == points[j].X {
				column[points[i].X] += 1
			} else if points[i].Y == points[j].Y{
				row[points[i].Y] += 1
			} else {
				x = points[i].X - points[j].X
				y = points[i].Y - points[j].Y
				if x < 0 && y < 0 {
					x = -x
					y = -y
				} else if x > 0 && y < 0 {
					x = -x
					y = -y
				}
				gcd = GCD(x, y)
				x1 = x%gcd
				y1 = y%gcd
				if normal[x1] == nil {
					normal[x1] = map[int]map[int]int{}
				}
				if normal[x1][y1] == nil {
					normal[x1][y1] = map[int]int{}
				}
				println(x, y, gcd)
			}
		}
	}

	return 0
}

func GCD(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}
