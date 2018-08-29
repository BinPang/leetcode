package main



//we use max's left and max's right
func maxArea(height []int) int {
	if len(height) <2 {
		panic("must gte 2")
	}
	max := 0
	mc := 0
	for i, v := range height {
		if i == 0 {
			max = v
			mc = 1
			continue
		}
		if v > max {
			max = v
			mc = 1
		} else if v == max {
			mc +=1
		}
	}
	if mc >= 2 {

	}
	return 0
}
