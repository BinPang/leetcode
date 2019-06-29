package main

func main() {
	println(nthUglyNumber(100))
	return
	println(nthUglyNumber(10))
	println(nthUglyNumber(18))
	println(nthUglyNumber(40))
}

func nthUglyNumber(n int) int {
	num1, num2, num3 := 2, 3, 5
	fact1, fact2, fact3 := 0, 0, 0
	total1, total2, total3 := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	i := 1
	for i < n {
		total1 = num1 * dp[fact1]
		total2 = num2 * dp[fact2]
		total3 = num3 * dp[fact3]

		if total1 < total2 {
			if total1 < total3 {
				fact1++
				if total1 == dp[i-1] {
					continue
				} else {
					dp[i] = total1
				}
			} else {
				fact3++
				if total3 == dp[i-1] {
					continue
				} else {
					dp[i] = total3
				}
			}
		} else { //total2 <= total1
			if total2 < total3 {
				fact2++
				if total2 == dp[i-1] {
					continue
				} else {
					dp[i] = total2
				}
			} else {
				fact3++
				if total3 == dp[i-1] {
					continue
				} else {
					dp[i] = total3
				}
			}
		}
		i++
	}
	return dp[n-1]
}
