package main

import (
	//"fmt"
	"strings"
)

func main() {
	r0 := wordBreak("a", []string{"a"})
	println(strings.Join(r0, "\n"))

	r10 := wordBreak("aaaa", []string{"a", "aa"})
	println(strings.Join(r10, "\n"))

	r1 := wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
	println(strings.Join(r1, "\n"))

	r2 := wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
	println(strings.Join(r2, "\n"))

	r3 := wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	println(strings.Join(r3, "\n"))
}

func wordBreak(s string, wordDict []string) []string {
	l := len(s)
	dp := make([][]int, l+1)
	for e := range dp {
		dp[e] = make([]int, 0)
	}
	wordMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordMap[v] = true
	}
	dp[0] = []int{0}

	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			if len(dp[j]) > 0 && wordMap[s[j:i]] {
				dp[i] = append(dp[i], j)
			}
		}
	}
	//println(fmt.Sprintf("%+v", dp))

	return dfs(s, &dp, l)
}

func dfs(s string, dp *[][]int, start int) []string {
	r := make([]string, 0)
	for _, v := range (*dp)[start] {
		if v == 0 {
			if start+1 == len(*dp) {
				r = append(r, s[v:start])
			} else {
				r = append(r, s[v:start]+" ")
			}
		} else {
			tmp := dfs(s, dp, v)
			for _, v2 := range tmp {
				if start+1 == len(*dp) {
					r = append(r, v2+s[v:start])
				} else {
					r = append(r, v2+s[v:start]+" ")
				}
			}
		}
	}
	return r
}
