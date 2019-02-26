package main

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	println("need 6,return:", maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	r := 0
	tmp := make([]int, len(matrix[0]))
	for _, v := range matrix {
		for ii, vv := range v {
			if vv == '1' {
				tmp[ii] += 1
			} else {
				tmp[ii] = 0
			}
		}
		r0 := cal(&tmp)
		if r < r0 {
			r = r0
		}
	}
	return r
}

func cal(heights *[]int) int {
	r := 0
	l := len(*heights)
	if l == 0 {
		return 0
	}
	h := 0
	tmpR := 0
	ii := 0

	for i, v := range *heights {
		if r < v {
			r = v
		}
		h = v
		for ii = i - 1; ii >= 0; ii-- {
			if (*heights)[ii] == 0 {
				break
			}
			if h > (*heights)[ii] {
				h = (*heights)[ii]
			}
			tmpR = (i - ii + 1) * h
			if r < tmpR {
				r = tmpR
			}
		}
	}
	return r
}
