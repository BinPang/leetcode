package main

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
	hash := make(map[string]int, lw)
	low := len(words[0]) //len one words
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
		if _, ok := hash[sArray[i]]; ok {
			for j = 0; j < lw1; j += low {
				if v, ok := hash[sArray[i+j]]; ok {
					if v > 0 {
						hash[sArray[i+j]] -= 1
					} else {
						break
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
			for tmp := 0; tmp < j; tmp += low {
				hash[sArray[i+tmp]] += 1
			}
		}
	}
	return r
}
