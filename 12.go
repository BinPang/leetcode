package main

func main() {
	println("input:1994,need:MCMXCIV,return:", intToRoman(1994))
	println("input:58,need:LVIII,return:", intToRoman(58))
}

func intToRoman(num int) string {
	tmp := 0
	level := 1
	ts := ""
	r := ""
	m := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	for {
		if num == 0 {
			break
		}
		tmp = num % 10
		if tmp == 4 {
			r = m[level] + m[level*5] + r
		} else if tmp == 9 {
			r = m[level] + m[level*10] + r
		} else if tmp == 5 {
			r = m[level*5] + r
		} else {
			ts = ""
			if tmp > 5 {
				ts = m[5*level]
				tmp -= 5
			}
			for i := 0; i < tmp; i++ {
				ts += m[level]
			}
			r = ts + r
		}
		//println(r)
		num /= 10
		level *= 10
	}

	return r
}
