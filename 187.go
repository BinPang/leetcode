package main

import "fmt"

func main() {
	println(fmt.Sprintf("%+v", findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")))
	println(fmt.Sprintf("%+v", findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAA")))
}

func findRepeatedDnaSequences(s string) []string {
	r := make([]string, 0)
	dup := make(map[string]int)
	for i := 0; i+9 < len(s); i++ {
		if dup[s[i:i+10]] == 0 {
			dup[s[i:i+10]] = 1
		} else if dup[s[i:i+10]] == 1 {
			r = append(r, s[i:i+10])
			dup[s[i:i+10]] = 2
		}
	}
	return r
}
