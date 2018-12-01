package main

import "fmt"

func main() {
	println(divide(8, 1))
	println(divide(-2147483648, -1))
	println(divide(8, 4))  //1000,100
	println(divide(10, 3)) //1010,11
	println(divide(10, -3))
}

func divide(dividend int, divisor int) int {
	sign := false
	if divisor == 0 {
		return 0
	}
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}
	if dividend < 0 && divisor < 0 {
		dividend = 0 - dividend
		divisor = 0 - divisor
	}
	if dividend < 0 || divisor < 0 {
		if dividend < 0 {
			dividend = 0 - dividend
		} else {
			divisor = 0 - divisor
		}
		sign = true
	}

	r := 0
	tmp := 1
	tmpDivisor := 1

	for {
		tmpDivisor = divisor
		tmp = 1
		if dividend < divisor {
			break
		}
		println(fmt.Sprintf("before:%b,%b,%d", dividend, divisor, tmp))
		for {
			tmpDivisor = tmpDivisor << 1
			if tmpDivisor > dividend {
				r += tmp
				dividend -= tmpDivisor >> 1
				println(fmt.Sprintf("after:%b-%b=%b,%d", dividend+tmpDivisor>>1, tmpDivisor>>1, dividend, tmp))
				break
			}
			tmp = tmp << 1
		}
	}

	if sign {
		return 0 - r
	} else {
		return r
	}
}
