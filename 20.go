package main

func isValid(s string) bool {
	bytes := []byte(s)
	store := make([]string, len(bytes))
	point := 0
	re := map[string]string{")": "(", "]": "[", "}": "{"}
	e := map[string]bool{"(": true, "[": true, "{": true}
	for _, v := range bytes {
		v1 := string(v)
		if _, ok := e[v1]; ok {
			store[point] = v1
			point++
		} else if _, ok := re[v1]; ok {
			if point > 0 && store[point-1] == re[v1] {
				point--
			} else {
				return false
			}
		}
	}
	if point == 0 {
		return true
	} else {
		return false
	}
}
