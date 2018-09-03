package main

func main() {
	println(romanToInt("LVIII"))
	println(romanToInt("III"))
	println(romanToInt("IV"))
	println(romanToInt("IX"))
	println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	l := len(s)
	tmp := ""
	r := 0

	for i := l - 1; i >= 0; i-- {
		tmp = string(s[i])
		r += m[tmp]
		if (tmp == "V" || tmp == "X") && i > 0 && string(s[i-1]) == "I"  {
			r -= 1
			i--
		} else if (tmp == "L" || tmp == "C") && i > 0 && string(s[i-1]) == "X"  {
			r -= 10
			i--
		} else if (tmp == "D" || tmp == "M") && i > 0 && string(s[i-1]) == "C"  {
			r -= 100
			i--
		}
	}
	return r
}
