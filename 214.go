package main

func main() {
	println(shortestPalindrome(""))
	println(shortestPalindrome("z"))
	println(shortestPalindrome("abcdefg"))
	println(shortestPalindrome("aacecaaa"))
	println(shortestPalindrome("aacecaaaa"))
}

func shortestPalindrome(s string) string {
	l := len(s)
	middle := l / 2
	gap := l%2 == 0
	for {
		if gap {
			if isPalindrome(s[0:middle], s[middle:middle*2]) {
				//find it
				return  reverse(s[middle*2:]) + s
			}
			middle--
		} else {
			if isPalindrome(s[0:middle], s[middle+1:2*middle+1]) {
				//find it
				return  reverse(s[middle*2+1:]) + s
			}
		}
		gap = !gap
	}
}

func isPalindrome(a, b string) bool {
	l := len(a)
	for i := 0; i < l; i++ {
		if a[i] != b[l-i-1] {
			return false
		}
	}
	return true
}

func reverse(a string)string {
	start := 0
	end := len(a)-1
	r := []byte(a)
	for start < end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
	return string(r)
}
