package main

func main() {
	println(numDecodings("310"))
}

func numDecodings(s string) int {
	println(s)
	l := len(s)
	if l == 0 {
		return 0
	} else if l == 1 {
		if s[0] == '0' {
			return 0
		} else {
			return 1
		}
	} else if l == 2 {
		if s[0] == '0' {
			return 0
		} else if s[0] == '1' {
			if s[1] == '0' {
				return 1
			} else  {
				return 2
			}
		} else if s[0] == '2' {
			if s[1] == '0' {
				return 1
			} else if s[1] > '6' {
				return 1
			} else {
				return 2
			}
		} else {
			if s[1] == '0' {
				return 0
			} else {
				return 1
			}
		}
	} else {
		return numDecodings(s[1:]) + numDecodings(s[2:])
	}
}
