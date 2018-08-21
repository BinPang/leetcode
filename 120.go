package main

func main() {
	t := [][]int{
		{2},
		{3,4},
		{6,5,7},
		{4,1,8,3},
	}
	println("need 11,result:", minimumTotal(t))
	println("need 0,result:", minimumTotal([][]int{}))
}


func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	if l == 0 {
		return 0
	}
	for i, v := range triangle{
		if i==0 {
			continue
		}
		l1 := len(v)
		for i1 := range v{
			if i1 == 0 {
				v[i1] += triangle[i-1][0]
			} else if i1 == l1 -1 {
				v[i1] += triangle[i-1][l1-2]
			} else {
				if triangle[i-1][i1] > triangle[i-1][i1-1] {
					v[i1] += triangle[i-1][i1-1]
				} else {
					v[i1] += triangle[i-1][i1]
				}
			}
		}
	}

	min := triangle[l-1][0]
	for _, v := range triangle[l-1]{
		if v < min {
			min = v
		}
	}
	return min
}
