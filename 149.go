package main

func main() {
	println(maxPoints([][]int{
		{3, 1},
		{12, 3},
		{3, 1},
		{-6, -1},
	}))
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
	} else if l == 1 {
		return 1
	}
	r := 0

	var t, s int
	var u, v int
	var same, vertical, horizontal, lineMax int
	var gcd int
	tmp := map[int]map[int]map[int]map[int]int{}
	for i := 0; i < l; i++ {
		same = 0
		vertical = 0
		horizontal = 0
		lineMax = 0
		tmp = map[int]map[int]map[int]map[int]int{}
		for j := i + 1; j < l; j++ {
			t = points[j][1] - points[i][1]
			s = points[j][0] - points[i][0]
			if t == 0 && s == 0 {
				same++
				continue
			} else if t == 0 {
				vertical++
				continue
			} else if s == 0 {
				horizontal++
				continue
			}
			if t < 0 {
				t = -t
				s = -s
			}
			gcd = GCD(t, s)
			t = t / gcd
			s = s / gcd
			//
			u = s*points[j][1] - t*points[j][0]
			v = s
			if u == 0 {
				s = 0
			} else {
				gcd = GCD(u, v)
				u = u / gcd
				v = v / gcd
				if u < 0 {
					u = -u
					v = -v
				}
			}
			if _, ok := tmp[t]; !ok {
				tmp[t] = make(map[int]map[int]map[int]int)
			}
			if _, ok := tmp[t][s]; !ok {
				tmp[t][s] = make(map[int]map[int]int)
			}
			if _, ok := tmp[t][s][u]; !ok {
				tmp[t][s][u] = make(map[int]int)
			}
			tmp[t][s][u][v] += 1
			if lineMax < tmp[t][s][u][v] {
				lineMax = tmp[t][s][u][v]
			}
		}
		r = max(r, vertical+same+1)
		r = max(r, horizontal+same+1)
		r = max(r, lineMax+same+1)
	}

	return r
}

func GCD(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
