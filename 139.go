package main

func main() {
	println(wordBreak("applepenapple", []string{"apple", "pen"}))
	println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	dp := make([]bool, l+1)
	dp[0] = true
	wordMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordMap[v] = true
	}
	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				//println(i, j, s[j:i])
				dp[i] = true
			}
		}
	}
	return dp[l]
}
