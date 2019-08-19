package main

func main() {
	println(toHex(4234423))
	println(toHex(-4234423))
	println(toHex(26))
	println(toHex(-1))
	println(toHex(0))
}

func toHex(num int) string {
	m := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	r := ""
	if num == 0 {
		r = "0"
	} else if num > 0 {
		for num > 0 {
			r = m[num&15] + r
			num >>= 4
		}
	} else {
		for i := 0; i < 8; i++ {
			r = m[num&15] + r
			num >>= 4
		}
	}
	return r
}
