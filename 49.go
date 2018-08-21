package main

func main() {
	t := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	for _, t0 := range t {
		for _, t1 := range t0 {
			print(t1, ",")
		}
		println("")
	}
}

/*func groupAnagrams(strs []string) [][]string {
	h := make(map[string][]string, 0)
	for _, str := range strs {
		t := map[byte]int{}
		for _, b := range []byte(str) {
			t[b] += 1
		}
		s := ""
		for i := 97; i <= 122; i++ {
			if v := t[byte(i)]; v != 0 {
				s += string(byte(i)) + string(v)
			}
		}
		if _, ok := h[s];!ok {
			h[s] = []string{str}
		} else {
			h[s] = append(h[s], str)
		}
	}
	r := make([][]string, len(h))
	i := 0
	for _, v := range h{
		r[i] = v
		i++
	}
	return r
}*/

func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string, 0)
	for _, str := range strs {
		t := make([]int, 26)
		for i := len(str)-1; i >= 0; i-- {
			t[str[i]-97] += 1
		}
		s := ""
		for _, v := range t {
			s += string(v) + "#"
		}
		if _, ok := resultMap[s]; !ok {
			resultMap[s] = []string{str}
		} else {
			resultMap[s] = append(resultMap[s], str)
		}
	}
	result := make([][]string, len(resultMap))
	i := 0
	for _, v := range resultMap {
		result[i] = v
		i++
	}
	return result
}

func sorted1(w string) string {
	t := make(map[string]int, 26)
	for i := len(w)-1; i >= 0; i-- {
		t[w[i:(i+1)]] += 1
	}
	s := ""
	for is, v := range t {
		s += is+string(v)
	}
	println("__",s)
	return s
}
