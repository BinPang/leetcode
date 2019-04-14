package main

func main() {
	v := Constructor()
	v.Insert("apple")
	println(v.Search("apple"))
	println(v.StartsWith("apple"))
	println(v.StartsWith("appled"))
	println(v.Search("app"))
	println(v.StartsWith("app"))

	v.Insert("app")
	println(v.Search("app"))
}

/**
Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
*/

type Trie struct {
	data [26]*Trie //a-z
	now  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{data: [26]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	head := this
	for e := range word {
		if head.data[word[e]-'a'] == nil {
			head.data[word[e]-'a'] = &Trie{}
		}
		head = head.data[word[e]-'a']
	}
	head.now = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	head := this
	for e := range word {
		if head.data[word[e]-'a'] == nil {
			return false
		}
		head = head.data[word[e]-'a']
	}
	return head.now
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return false
	}
	head := this
	for e := range prefix {
		if head.data[prefix[e]-'a'] == nil {
			return false
		}
		head = head.data[prefix[e]-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
