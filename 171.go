package main

func titleToNumber(s string) int {
	l := len(s)
	r := 0
	tmp := 1
	for i := l - 1; i >= 0; i-- {
		r = r + int(s[i]-'A'+1)*tmp
		tmp *= 26
	}
	return r
}
