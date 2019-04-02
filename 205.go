package main

func isIsomorphic(s string, t string) bool {
	l := len(s)
	st := make(map[uint8]uint8, 256)
	ts := make(map[uint8]uint8, 256)
	var ok0, ok1 bool
	for i := 0; i < l; i++ {
		_, ok0 = st[s[i]]
		_, ok1 = ts[t[i]]
		if ok0 && ok1 {
			if st[s[i]] == t[i] && ts[t[i]] == s[i] {
			} else {
				return false
			}
		} else if !ok0 && !ok1 {
			st[s[i]] = t[i]
			ts[t[i]] = s[i]
		} else {
			return false
		}
	}

	return true
}
