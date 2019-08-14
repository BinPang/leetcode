package main

func main() {
	println(addStrings("12", "345"))
	println(addStrings("999", "1"))
	println(addStrings("1", "999"))
}

func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	if l1 < l2 {
		l1, l2 = l2, l1
		num1, num2 = num2, num1
	}
	l := l1 + 1 //l1 is bigger
	r := make([]byte, l)
	var incr uint8 = 0
	var tmp uint8 = 0
	for l != 1 {
		if l2 <= 0 {
			tmp = num1[l1-1] - '0' + incr
		} else {
			tmp = num1[l1-1] - '0' + num2[l2-1] - '0' + incr
		}
		if tmp >= 10 {
			incr = 1
			r[l-1] = tmp - 10 + '0'
		} else {
			incr = 0
			r[l-1] = tmp + '0'
		}
		l--
		l1--
		l2--
	}
	if incr == 1 {
		r[0] = 1 + '0'
	} else {
		r = r[1:]
	}

	return string(r)
}
