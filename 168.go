package main

func main() {
	println(convertToTitle(26*26*26), 26*26*26)
	println(convertToTitle(1))
	println(convertToTitle(26))
	println(convertToTitle(100))
	println(convertToTitle(10000))
}

func convertToTitle(n int) string {
	data := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	r := ""
	t := 0
	for {
		if n == 0 {
			break
		}

		t = n % 26
		if t == 0 {
			t = 26
		}
		n = (n - t) / 26
		r = data[t-1] + r
	}

	return r
}
