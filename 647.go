package main

func main() {
	println(countSubstrings("abc"))
	println(countSubstrings("aaa"))
}

func countSubstrings(s string) int {
	l := len(s)
	r := 0
	for i := 0; i < l; i++ {
		r += 1
		start := i - 1
		end := i + 1
		for {
			if start < 0 || end >= l {
				break
			}
			if s[start] == s[end] {
				r += 1
				start--
				end++
				continue
			}
			break
		}

		start = i
		end = i + 1
		for {
			if start < 0 || end >= l {
				break
			}
			if s[start] == s[end] {
				r += 1
				start--
				end++
				continue
			}
			break
		}
	}
	return r
}
