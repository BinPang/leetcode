package main

func main() {
	println(minWindow("bba", "ab"))
	println(minWindow("aa", "aa"))
	println(minWindow("ab", "b"))
	println(minWindow("a", "aa"))
	println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	r := ""
	tmpR := ""
	l := len(t)
	ls := len(s)
	tmap := make(map[byte]int, l)
	store := make(map[byte]int, l)
	for i := 0; i < l; i++ {
		tmap[t[i]] += 1
		store[t[i]] = 0
	}
	start := -1
	end := 0
	loopCount := 0
	for {
		if end == ls {
			break
		}
		if _, ok := store[s[end]]; ok {
			store[s[end]] ++
			if store[s[end]] == tmap[s[end]] {
				loopCount ++
			}
			if start < 0 {
				start = end
			}
		}
		if loopCount == len(tmap) {
			//deal with, bcbbba, ab.first item is gt
			for {
				if store[s[start]] > 0 {
					if store[s[start]] == tmap[s[start]] {
						break
					} else {
						store[s[start]]--
					}
				}
				start++
			}
			r = s[start : end+1]
			break
		}
		end++
	}
	//println(s, t, start, end, loopCount, r)

	end++
	for {
		if end >= ls {
			break
		}
		if _, ok := tmap[s[end]]; ok {
			store[s[end]] ++
			for {
				if _, ok := tmap[s[start]]; ok {
					if tmap[s[start]] == store[s[start]] {
						break
					} else {
						store[s[start]]--
					}
				}
				start++
			}
			tmpR = s[start : end+1]
			//println(s, t, start, end, r, tmpR)
			if len(r) > len(tmpR) {
				r = tmpR
			}
		}
		end++
	}
	return r
}
