package main

import "fmt"

func main() {
	r3 := palindromePairs([]string{"bb", "bababab", "baab", "abaabaa", "aaba", "", "bbaa", "aba", "baa", "b"})
	println(fmt.Sprintf("%+v", r3))
	r0 := palindromePairs([]string{"a", ""})
	println(fmt.Sprintf("%+v", r0))
	r1 := palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"})
	println(fmt.Sprintf("%+v", r1))
	r2 := palindromePairs([]string{"bat", "tab", "cat"})
	println(fmt.Sprintf("%+v", r2))
}

type trie struct {
	index int
	ok    bool
	next  map[uint8]*trie
}

func palindromePairs(words []string) [][]int {
	t := &trie{next: make(map[uint8]*trie)}
	tmp := t
	empty := -1
	for k, v := range words {
		tmp = t
		for e := range v {
			if tmp.next[v[e]] == nil {
				tmp.next[v[e]] = &trie{next: make(map[uint8]*trie)}
			}
			tmp = tmp.next[v[e]]
		}
		tmp.index = k
		tmp.ok = true

		if v == "" {
			empty = k
		}
	}
	dup := make(map[int]map[int]bool, 0)
	r := make([][]int, 0)

	for k, v := range words {
		l := len(v)
		//left
		//	1	3	5	7	9	11
		//	a	b	c	d	e	f
		//0   2   4   6   8   10  12
		for i := 0; i <= 2*l; i++ {
			//println(":::",k, v, i, l, palindromeOk(v, i, l))
			if palindromeOk(v, i, l) {
				needTrie := ""
				if i < l {
					needTrie = v[i:]
				} else if i == l {
					//middle is ok
					if empty != -1 && empty != k {
						r = append(r, []int{k, empty}, []int{empty, k})
					}
					continue
				} else {
					needTrie = v[0:(i - l)]
				}
				j, ok := trieOk(t, needTrie)
				//println("_:", k, needTrie, j, ok, i, l)
				if ok {
					if j == k {
						continue
					}
					if i < l {
						if _, ok = dup[j]; ok {
							if _, ok = dup[j][k]; ok {
								continue
							}
						} else {
							dup[j] = map[int]bool{}
						}
						r = append(r, []int{j, k})
						dup[j][k] = true
					} else {
						if _, ok = dup[k]; ok {
							if _, ok = dup[k][j]; ok {
								continue
							}
						} else {
							dup[k] = map[int]bool{}
						}
						r = append(r, []int{k, j})
						dup[k][j] = true
					}
				}
			}
		}
	}

	return r
}

func palindromeOk(s string, p, l int) bool {
	start, end := 0, 0
	if p%2 == 0 {
		start = p/2 - 1
		end = p / 2
	} else {
		start = (p-1)/2 - 1
		end = (p-1)/2 + 1
	}
	for {
		if start < 0 || end == l {
			return true
		}
		if s[start] != s[end] {
			return false
		}
		start -= 1
		end += 1
	}
}

func trieOk(t *trie, s string) (int, bool) {
	//s is reverse
	tmp := t
	for e := len(s) - 1; e >= 0; e-- {
		//for e := range s {
		if tmp == nil {
			return 0, false
		}
		if tmp.next[s[e]] == nil {
			return 0, false
		}
		tmp = tmp.next[s[e]]
	}
	return tmp.index, tmp.ok
}
