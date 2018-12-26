package main

//Your runtime beats 41.89 % of golang submissions.
func main() {
	ts := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	println(len(ts))

	ts1 := findSubstring("barfoothefoobarfoo", []string{"foo", "bar"})
	println(len(ts1))

	ts2 := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	println(len(ts2))
}

func findSubstring(s string, words []string) []int {
	l := len(s)
	if l == 0 {
		return nil
	}
	lw := len(words)
	if lw == 0 || words[0] == "" {
		return nil
	}
	r := make([]int, 0)
	hash := make(map[string]int)
	low := len(words[0]) //len one word
	lw1 := lw * low
	for _, v := range words {
		if len(v) != low {
			return nil
		}
		hash[v] += 1
	}
	sArray := make([]string, l-low+1)
	for i := range sArray {
		sArray[i] = s[i : i+low]
	}
	ii := l - lw*len(words[0])
	j := 0
	for i := 0; i <= ii; i++ {
		if hash[sArray[i]] > 0 {
			tmp := make(map[string]int)
			for j = 0; j < lw1; j += low {
				if hash[sArray[i+j]] > 0 {
					if tmp[sArray[i+j]] == hash[sArray[i+j]] {
						break
					} else {
						tmp[sArray[i+j]] += 1
					}
				} else {
					break
				}
			}
			if j == lw1 {
				//find
				//println("find,s->", s, ":", i, find, lw)
				r = append(r, i)
			}
			//println("start:", s, i, j, lw1)
		}
	}
	return r
}
