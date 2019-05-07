package main

func main() {
	println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
}

func maxProduct(words []string) int {
	l := len(words)
	bin := make([]int, 26)
	tmp := make([]int, l)
	r := 0
	tmpR := 0
	for i := 0; i < 26; i++ {
		bin[i] = 1 << uint(i)
	}
	for k, v := range words {
		for _, v1 := range v {
			tmp[k] |= bin[v1-'a']
		}
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if tmp[i]&tmp[j] == 0 {
				tmpR = len(words[i]) * len(words[j])
				if tmpR > r {
					r = tmpR
				}
			}
		}
	}
	return r
}
