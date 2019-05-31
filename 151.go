package main

func main() {
	println(reverseWords(" 1"))
	println(reverseWords("  "))
	println(reverseWords("t       sky is blue"))
	println(reverseWords("t       sky is b"))
	println(reverseWords("t       sky is b   "))
	println(reverseWords("  the sky is blue   "))
	println(reverseWords("the sky is blue"))
	println(reverseWords("      the sky is blue"))
	println(reverseWords("the       sky is blue"))
	println(reverseWords("the       sky is blue      "))
	println(reverseWords("the"))
	println(reverseWords("   the"))
	println(reverseWords(""))

}

func reverseWords(s string) string {
	l := len(s)
	r := ""
	end := l - 1
	for end >= 0 {
		if s[end] == ' ' {
			end--
			continue
		}
		break
	}
	if end < 0 {
		return ""
	}
	preBlank := false
	start := end
	for {
		//println(start, end, preBlank)
		if start == 0 {
			if s[start] == ' ' {
				start++
			}
			if preBlank {
				r += " " + s[start:end+1]
			} else {
				r += s[start : end+1]
			}
			break
		}
		if s[start] == ' ' {
			if preBlank {
				r += " " + s[start+1:end+1]
			} else {
				r += s[start+1 : end+1]
				preBlank = true
			}
			end = start
			for end >= 0 {
				if s[end] == ' ' {
					end--
					continue
				}
				break
			}
			if end < 0 {
				break
			}
			start = end
		} else {
			start--
		}
	}

	return r
}