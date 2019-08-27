package main

import "fmt"

func main() {
	t := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}

	println(fmt.Sprintf("%+v", findWords(t, []string{"oath", "pea", "eat", "rain"})))

	t = [][]byte{
		{'b', 'a', 'c'},
	}
	println(fmt.Sprintf("%+v", findWords(t, []string{"ba", "bac"})))
}

type trie struct {
	data [26]*trie
	ok   bool
}

func findWords(board [][]byte, words []string) []string {
	l1 := len(board)
	if l1 == 0 {
		return nil
	}
	l2 := len(board[0])
	if l2 == 0 {
		return nil
	}
	data := &trie{data: [26]*trie{}}
	for _, v1 := range words {
		tmp := data
		for _, v2 := range v1 {
			if tmp.data[v2-'a'] == nil {
				tmp.data[v2-'a'] = &trie{data: [26]*trie{}}
			}
			tmp = tmp.data[v2-'a']
		}
		tmp.ok = true
	}
	r := make([]string, 0)
	exist := make(map[string]bool)
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			tmpR, _ := _findWords(&board, data, i, j, l1, l2)
			for _, v := range tmpR {
				if exist[v] {
					continue
				}
				exist[v] = true
				r = append(r, v)
			}
		}
	}
	return r
}

func _findWords(board *[][]byte, trie *trie, i, j, l1, l2 int) ([]string, bool) {
	if trie == nil {
		return nil, false
	}
	if i < 0 || j < 0 || i == l1 || j == l2 {
		return nil, false
	}
	if (*board)[i][j] == '0' {
		return nil, false
	}
	if trie.data[(*board)[i][j]-'a'] == nil {
		return nil, false
	}
	//up, right, bottom, left
	r := make([]string, 0)
	tmp := (*board)[i][j]
	(*board)[i][j] = '0'
	data, ok := _findWords(board, trie.data[tmp-'a'], i-1, j, l1, l2)
	if ok {
		for _, v := range data {
			r = append(r, string(tmp)+v)
		}
	}
	data, ok = _findWords(board, trie.data[tmp-'a'], i, j+1, l1, l2)
	if ok {
		for _, v := range data {
			r = append(r, string(tmp)+v)
		}
	}
	data, ok = _findWords(board, trie.data[tmp-'a'], i+1, j, l1, l2)
	if ok {
		for _, v := range data {
			r = append(r, string(tmp)+v)
		}
	}
	data, ok = _findWords(board, trie.data[tmp-'a'], i, j-1, l1, l2)
	if ok {
		for _, v := range data {
			r = append(r, string(tmp)+v)
		}
	}
	if trie.data[tmp-'a'].ok {
		r = append(r, string(tmp))
	}
	(*board)[i][j] = tmp
	return r, true
}
