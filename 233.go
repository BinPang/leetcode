package main

func main() {
	println(countDigitOne(13))
}

//(a + 8) / 10 * m + (a % 10 == 1) * (b + 1);
//copy from internet
func countDigitOne(n int) int {
	r := 0
	m := 0
	for m = 1; m <= n; m *= 10 {
		a := n / m
		b := n % m
		r += (a + 8) / 10 * m
		if a%10 == 1 {
			r += b + 1
		}
	}
	return r
}
