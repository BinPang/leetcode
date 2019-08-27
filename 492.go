package main

import (
	"fmt"
	"math"
)

func main() {
	println(fmt.Sprintf("%+v", constructRectangle(4)))
	println(fmt.Sprintf("%+v", constructRectangle(5)))
	println(fmt.Sprintf("%+v", constructRectangle(8)))
	println(fmt.Sprintf("%+v", constructRectangle(9)))
}

func constructRectangle(area int) []int {
	l := int(math.Sqrt(float64(area)))
	for {
		if area%l == 0 {
			break
		}
		l--
	}
	return []int{area / l, l}
}