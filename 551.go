package main

func checkRecord(s string) bool {
	l := len(s)
	A := 0
	for k, v := range s {
		if v == 'A' {
			A++
			if A > 1 {
				return false
			}
		} else if v == 'L' {
			if k+1 < l && s[k+1] == 'L' && k+2 < l && s[k+2] == 'L' {
				return false
			}
		}
	}
	return true
}
