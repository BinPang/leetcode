package main

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	println(ladderLength(beginWord, endWord, wordList))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	l0 := len(wordList)
	mw := make(map[string]int, l0)
	for _, v := range wordList{
		mw[v] = l0+1
	}

	l1 := len(beginWord)
	work := []string{beginWord}
	i:=0
	var j byte
	for {
		if len(work) == 0 {
			break
		}
		tmpWord := []byte(work[0])
		work = work[1:]
		for i = 0; i < l1; i++ {
			tmpChar := tmpWord[i]
			for j = 'a'; j <= 'z'; j++ {
				tmpWord[i] = j
				if mw[string(tmpWord)] != 0 {
					work = append(work, string(tmpWord))
				}
			}
			tmpWord[i] = tmpChar
		}
	}


	return 0
}
