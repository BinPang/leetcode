package main

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	println(ladderLength(beginWord, endWord, wordList))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	l := len(wordList)
	visited := make([]string, l+1)
	visited[0] = beginWord
	pVisited := 1
	lw := len(beginWord)

	hashWordList := make(map[string]bool, l)
	hashVisited := make(map[string]bool, l)
	

	return 0
}
