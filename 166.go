package main

import (
	"strconv"
)

func main() {
	println(fractionToDecimal(-2147483648, -10))
	println(fractionToDecimal(-2147483648, 1))
	println(fractionToDecimal(-50, 8))
	println(fractionToDecimal(4, 333))
	println(fractionToDecimal(1, 50))
	println(fractionToDecimal(1, 6))
	println(fractionToDecimal(4, 2))
	println(fractionToDecimal(2, 3))
	println(fractionToDecimal(8, 7))
	println(fractionToDecimal(508, 17))
}

func fractionToDecimal(numerator int, denominator int) string {
	r := ""
	if numerator < 0 && denominator > 0 {
		r = "-"
		numerator = -numerator
	} else if numerator > 0 && denominator < 0 {
		r = "-"
		denominator = -denominator
	} else if numerator < 0 && denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	}
	if numerator%denominator == 0 {
		return r + strconv.Itoa(numerator/denominator)
	}
	if numerator > denominator {
		r += strconv.Itoa(numerator/denominator) + "."
		numerator = numerator % denominator
	} else {
		r += "0."
	}
	tmp := make(map[int]bool, 0)
	tmp1 := make([]int, 0)
	numerator *= 10

	for {
		if tmp[numerator] {
			for k, v := range tmp1 {
				if tmp1[k] == numerator {
					r += "("
				}
				r += strconv.Itoa(v/denominator)
			}
			r += ")"
			break
		}
		tmp1 = append(tmp1, numerator)

		tmp[numerator] = true
		numerator = (numerator % denominator) * 10
		if numerator == 0 {
			for _, v := range tmp1 {
				r += strconv.Itoa(v/denominator)
			}
			break
		}
	}
	return r
}
