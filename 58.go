package main

func main() {
	println("hello world", lengthOfLastWord("hello world"))
	println("hello  ", lengthOfLastWord("hello  "))
	println("", lengthOfLastWord(""))
	println("  ", lengthOfLastWord("  "))
	println("a ", lengthOfLastWord("a "))
}

func lengthOfLastWord(s string) int {
	start := 0
	end := len(s)-1
	lastEmpty := false
	r := 0
	for start <= end {
		if s[end] == ' ' {
			if lastEmpty {
				break
			} else {
				end--
				continue
			}
		} else {
			lastEmpty = true
		}
		r +=1
		end-=1
	}
	return r
}
