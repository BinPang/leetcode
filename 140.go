package main

import (
	//"fmt"
	"strings"
)

func main() {
	r1 := wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
	println(strings.Join(r1, "\n"))

	r2 := wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
	println(strings.Join(r2, "\n"))

	r3:= wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	println(strings.Join(r3, "\n"))
}

func wordBreak(s string, wordDict []string) []string {
	r := make([]string, 0)
	l := len(s)
	dp := make([]bool, l+1)
	dp[0] = true

	stack1 := make([]int, l)
	stack2 := make([]int, l)
	pstack := 0 //just a stack

	wordMap := make(map[string]bool, len(wordDict))
	maxLen := 0
	for _, v := range wordDict {
		wordMap[v] = true
		if len(v) > maxLen {
			maxLen = len(v) //here to reduce time
		}
	}

	start := 0
	end := 1
	for {
		//println(l, start, end)
		if end == l {
			//maybe find or not
			if wordMap[s[start:end]] {
				//find
				//println(fmt.Sprintf("%+v,%+v", stack1, stack2))
				oneR := ""
				for i := 0; i < pstack; i++ {
					oneR += s[stack1[i]:stack2[i]+1] + " "
				}
				oneR += s[start:end]
				r = append(r, oneR)
			}
			if pstack == 0 {
				break
			}
			pstack--
			start = stack1[pstack]
			end = stack2[pstack] + 2
			continue
		}
		if wordMap[s[start:end]] {
			stack1[pstack] = start
			stack2[pstack] = end - 1
			pstack++
			start = end
			end += 1
			continue
		}
		end += 1
		if end - start > maxLen {
			if pstack == 0 {
				break
			}
			pstack--
			start = stack1[pstack]
			end = stack2[pstack] + 2
		}
	}

	return r
}
