package main

import "strconv"

func main() {
	println(evalRPN([]string{"1", "3", "+"}))
	println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	l := len(tokens)
	if l == 0 {
		return 0
	}
	s := make([]int, len(tokens))
	p := 0
	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			p -= 2
			if v == "+" {
				s[p] = s[p] + s[p+1]
			} else if v == "-" {
				s[p] = s[p] - s[p+1]
			} else if v == "*" {
				s[p] = s[p] * s[p+1]
			} else {
				s[p] = s[p] / s[p+1]
			}
			p++
		} else {
			s[p], _ = strconv.Atoi(v)
			p++
		}
	}
	return s[0]
}
