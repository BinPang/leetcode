package main

import "fmt"

func main() {
	println(multiply("9", "99"), 99*9)
	return

	multiply("123", "456")
}

func multiply(num1 string, num2 string) string {
	if num1 == "" || num2 == "" {
		return ""
	}
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1 := len(num1)
	l2 := len(num2)
	tmp := make([]int, l1+l2)
	t0 := 0
	p := 0
	/**
		1	2	3
		4	5	6
	-------------
	 */
	for i := l2 - 1; i >= 0; i-- {
		for j := l1 - 1; j >= 0; j-- {
			t0 = (int(num2[i]) - 48) * (int(num1[j]) - 48)
			p = l1 - 1 + l2 - 1 - i - j
			tmp[p] += t0 % 10
			tmp[p+1] += t0 / 10
		}
	}
	for i, v := range tmp {
		if v >= 10 {
			tmp[i+1] += v / 10
			tmp[i] = v % 10
		}
	}
	counting := false
	r := ""
	for p = l1 + l2 - 1; p >= 0; p-- {
		if tmp[p] != 0 || counting {
			r+=fmt.Sprintf("%d", tmp[p])
			counting = true
		}
	}
	//for i, v := range tmp {
	//	println(i, v)
	//}
	//println(123 * 456, r)
	return r
}
