package main

func main() {
	println("true => ", isScramble("abcd", "cdab"))
	println("true => ", isScramble("abc", "cab"))
	println("false => ",isScramble("aa", "ba"))
	println("false => ",isScramble("aa", "ab"))
	println("true => ", isScramble("great", "rgeat"))
	println("false => ",isScramble("abcde", "caebd"))
}

func isScramble(s1 string, s2 string) bool {
	//println(s1, s2)
	if s1 == s2 {
		return true
	}
	l := len(s1)
	if l != len(s2) {
		return false
	}
	tmp := make([]int, 26)
	for i := range s1 {
		tmp[s1[i]-'a'] ++
		tmp[s2[i]-'a'] --
	}
	for _, v := range tmp {
		if v != 0 {
			return false
		}
	}

	for i := 1; i < l; i++ {
		//println(i, s1, s2, "__:", s1[0:i], s2[0:i], s1[i:], s2[i:])
		if isScramble(s1[0:i], s2[0:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		//println(i, s1, s2, "++:", s1[0:i], s2[l-i:], s1[i:], s2[0:l-i])
		if isScramble(s1[0:i], s2[l-i:]) && isScramble(s1[i:], s2[0:l-i]) {
			return true
		}
	}
	return false
}
