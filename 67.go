package main

func main() {
	println(addBinary("1", "111"))
	return
	println(addBinary("1011", "1011"))
	println(addBinary("10", "1"))
}

func addBinary(a string, b string) string {
	l1 := len(a)
	l2 := len(b)
	l := l1
	if l2 > l1 {
		l = l2
	}
	//0=>48, 1=>49
	by := make([]byte, l)
	l1 -= 1
	l2 -= 1
	l -= 1
	t0 := 0
	t1 := 0
	incr := 0
	tmpSum := 0
	for {
		if l1 < 0 && l2 < 0 {
			break
		}
		t0 = 0
		t1 = 0
		if l1 >= 0 && l2 >= 0 {
			if a[l1] == 49 {
				t0 = 1
			}
			if b[l2] == 49 {
				t1 = 1
			}
		} else if l1 >= 0 && a[l1] == 49 {
			t0 = 1
		} else if l2 >= 0 && b[l2] == 49 {
			t1 = 1
		}
		tmpSum = t0 + t1 + incr
		if tmpSum == 3 {
			incr = 1
			by[l] = 49
		} else if tmpSum == 2 {
			incr = 1
			by[l] = 48
		} else if tmpSum == 1 {
			incr = 0
			by[l] = 49
		} else {
			incr = 0
			by[l] = 48
		}
		l1--
		l2--
		l--
	}
	if incr == 1 {
		return "1" + string(by)
	} else {
		return string(by)
	}
}
