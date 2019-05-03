package main

import "fmt"

func main() {
	println(fmt.Sprintf("%+v", diffWaysToCompute("2*3-4*5")))
	println(fmt.Sprintf("%+v", diffWaysToCompute("2+3")))
}

func diffWaysToCompute(input string) []int {
	isNum := true
	r := make([]int, 0)
	for k, v := range input {
		if v == '+' || v == '-' || v == '*' {
			part1 := diffWaysToCompute(input[0:k])
			part2 := diffWaysToCompute(input[k+1:])
			for _, v1 := range part1 {
				for _, v2 := range part2 {
					r = append(r, calc(v1, v2, v))
				}
			}
			isNum = false
		}
	}
	if isNum {
		num := 0
		for _, v := range input {
			num = num*10 + int(v-'0')
		}
		r = append(r, num)
	}
	return r
}

func calc(a, b int, op int32) int {
	if op == '+' {
		return a + b
	} else if op == '-' {
		return a - b
	} else {
		return a * b
	}
}
