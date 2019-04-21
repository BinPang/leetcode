package main

func main() {
	println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	var x1, y1 int
	var x2, y2 int
	find := false
	r := (C-A)*(D-B) + (G-E)*(H-F)

	if A <= E && E <= C {
		x1 = E
		if E <= C && C <= G {
			x2 = C
		} else {
			x2 = G
		}
		if F <= B && B <= H {
			y1 = B
			if B <= H && H <= D {
				y2 = H
			} else {
				y2 = D
			}
			find = true
		} else if B <= F && F <= D {
			y1 = F
			if F <= D && D <= H {
				y2 = D
			} else {
				y2 = H
			}
			find = true
		}
	} else if E <= A && A <= G {
		x1 = A
		if A <= G && G <= C {
			x2 = G
		} else {
			x2 = C
		}
		if B <= F && F <= D {
			y1 = F
			if F <= D && D <= H {
				y2 = D
			} else {
				y2 = H
			}
			find = true
		} else if F <= B && B <= H {
			y1 = B
			if B <= H && H <= D {
				y2 = H
			} else {
				y2 = D
			}
			find = true
		}
	}
	//println(x1, y1, x2, y2)
	if find {
		return r - (x2-x1)*(y2-y1)
	}
	return r
}
