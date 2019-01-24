package main

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	println(ladderLength(beginWord, endWord, wordList))

	wordList = []string{"hot", "dot", "dog", "lot", "log"}
	println(ladderLength(beginWord, endWord, wordList))
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	l0 := len(wordList)
	mw := make(map[string]int, l0)
	for _, v := range wordList {
		mw[v] = l0 + 1
	}

	l1 := len(beginWord)
	oldWork := []string{beginWord}
	newWork := make([]string, 0)
	i := 0
	path := 0
	var j byte
	for {
		if len(oldWork) == 0 {
			break
		}
		for _, v := range oldWork {
			tmpWord := []byte(v)
			for i = 0; i < l1; i++ {
				tmpChar := tmpWord[i]
				for j = 'a'; j <= 'z'; j++ {
					tmpWord[i] = j
					if mw[string(tmpWord)] != 0 {
						if string(tmpWord) == endWord {
							return path + 2
						}
						if path < mw[string(tmpWord)] {
							mw[string(tmpWord)] = path
							newWork = append(newWork, string(tmpWord))
						}
					}
				}
				tmpWord[i] = tmpChar
			}
		}
		//println(fmt.Sprintf("%d,list:%+v,%+v,%+v", path, wordList, oldWork, newWork))
		oldWork = newWork
		newWork = make([]string, 0)
		path++
	}

	return 0
}

