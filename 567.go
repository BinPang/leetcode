package main

func main() {
	println(checkInclusion("aabcc", "abbaacc"))
	println(checkInclusion("aabb", "aabooaabb"))
	println(checkInclusion("aabb", "aaabboo"))
	println(checkInclusion("aabb", "eidbcaaabboo"))
	println(checkInclusion("aabb", "eidbcabbaoo"))
	println(checkInclusion("aabb", "eidbabbaoo"))
	println(checkInclusion("ab", "eidbaooo"))
	println(checkInclusion("ab", "eidbaooo"))
	println(checkInclusion("ab", "eidboaoo"))
	println(checkInclusion("aa", "eidbaaoo"))
}

func checkInclusion(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 > l2 {
		return false
	}
	tmp := make([]int, 26)
	tmp1 := make([]int, 26)
	for _, v := range s1 {
		tmp[v-'a'] += 1
		tmp1[v-'a'] += 1
	}
	index := 0
	start := 0
	l := 0
	for index < l2 {
		if tmp1[s2[index]-'a'] == 0 {
			l = 0
			if index != 0 && tmp1[s2[index-1]-'a'] != 0 {
				for i := 0; i < 26; i++ {
					tmp[i] = tmp1[i]
				}
			}
			index++
			start = index
			continue
		}
		if tmp[s2[index]-'a'] == 0 {
			for {
				if s2[start] == s2[index] {
					start++
					break
				}
				tmp[s2[start]-'a']++
				start++
				l--
			}
		} else {
			tmp[s2[index]-'a']--
			l++
			//println(l, index, l1)
			if l == l1 {
				return true
			}
		}
		index++
	}
	return false
}
