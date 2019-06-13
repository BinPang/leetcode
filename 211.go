package main

func main() {
	t := Constructor()
	t.AddWord("bad")
	t.AddWord("dad")
	t.AddWord("mad")
	println(t.Search("pad"))
	println(t.Search("bad"))
	println(t.Search(".ad"))
	println(t.Search("b.."))
}

type WordDictionary struct {
	end  bool
	next [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if word == "" {
		this.end = true
		return
	}
	tmp := this
	for _, v := range word {
		if tmp.next[int(v-'a')] == nil {
			tmp.next[int(v-'a')] = &WordDictionary{}
		}
		tmp = tmp.next[int(v-'a')]
	}
	tmp.end = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return s(this, word)
}

func s(w *WordDictionary, word string) bool {
	if word == "" {
		return w.end == true
	}
	if word[0] == '.' {
		ok := false
		for i := 0; i < 26; i++ {
			if w.next[i] != nil && s(w.next[i], word[1:]) {
				ok = true
				break
			}
		}
		return ok
	} else {
		if w.next[word[0]-'a'] == nil {
			return false
		}
		return s(w.next[word[0]-'a'], word[1:])
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
