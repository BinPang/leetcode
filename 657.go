package main

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, v := range moves {
		if v == 'L' {
			x--
		} else if v == 'R' {
			x++
		} else if v == 'U' {
			y++
		} else if v == 'D' {
			y--
		}
	}
	return x == 0 && y == 0
}
