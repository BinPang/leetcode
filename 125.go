package main

func main() {
	println("input:0z;z   ; 0,need true,return:", isPalindrome("0z;z   ; 0"))
	return
	println("input:A man, a plan, a canal: Panama,need true,return:", isPalindrome("A man, a plan, a canal: Panama"))
	println("input:azAZ,need false,return:", isPalindrome("azAZ"))
	println("input:AzZa,need true,return:", isPalindrome("AzZa"))
	println("input:A3zZ4a,need false,return:", isPalindrome("A3zZ4a"))
	println("input:A3zZ3a,need true,return:", isPalindrome("A3zZ3a"))
}

//a-z:97-122,A-Z:65-90,0-9:48~57
func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for {
		if start >= end {
			break
		}
		if !((s[start] >= 97 && s[start] <= 122) || (s[start] >= 65 && s[start] <= 90) || (s[start] >= 48 && s[start] <= 57)) {
			start++
			continue
		}
		if !((s[end] >= 97 && s[end] <= 122) || (s[end] >= 65 && s[end] <= 90) || (s[end] >= 48 && s[end] <= 57)) {
			end--
			continue
		}
		if (s[start] >= 48 && s[start] <= 57) || (s[end] >= 48 && s[end] <= 57) {
			if s[start] == s[end] {
				start++
				end--
				continue
			} else {
				return false
			}
		}
		if s[start] == s[end] || s[start]+32 == s[end] || s[end]+32 == s[start] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}
