package main

import "fmt"

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	r0 := findLadders(beginWord, endWord, wordList)
	println(fmt.Sprintf("result:%+v", r0))

	wordList = []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}
	r0 = findLadders("red", "tax", wordList)
	//[["red","ted","tad","tax"],["red","ted","tex","tax"]]
	//[["red","ted","tad","tax"],["red","ted","tex","tax"],["red","rex","tex","tax"]]
	println(fmt.Sprintf("result:%+v", r0))

	wordList = []string{"hot", "dot", "dog", "lot", "log"}
	r0 = findLadders(beginWord, endWord, wordList)
	println(fmt.Sprintf("result:%+v", r0))

	wordList = []string{"kid", "tag", "pup", "ail", "tun", "woo", "erg", "luz", "brr", "gay", "sip", "kay", "per", "val", "mes", "ohs", "now", "boa", "cet", "pal", "bar", "die", "war", "hay", "eco", "pub", "lob", "rue", "fry", "lit", "rex", "jan", "cot", "bid", "ali", "pay", "col", "gum", "ger", "row", "won", "dan", "rum", "fad", "tut", "sag", "yip", "sui", "ark", "has", "zip", "fez", "own", "ump", "dis", "ads", "max", "jaw", "out", "btu", "ana", "gap", "cry", "led", "abe", "box", "ore", "pig", "fie", "toy", "fat", "cal", "lie", "noh", "sew", "ono", "tam", "flu", "mgm", "ply", "awe", "pry", "tit", "tie", "yet", "too", "tax", "jim", "san", "pan", "map", "ski", "ova", "wed", "non", "wac", "nut", "why", "bye", "lye", "oct", "old", "fin", "feb", "chi", "sap", "owl", "log", "tod", "dot", "bow", "fob", "for", "joe", "ivy", "fan", "age", "fax", "hip", "jib", "mel", "hus", "sob", "ifs", "tab", "ara", "dab", "jag", "jar", "arm", "lot", "tom", "sax", "tex", "yum", "pei", "wen", "wry", "ire", "irk", "far", "mew", "wit", "doe", "gas", "rte", "ian", "pot", "ask", "wag", "hag", "amy", "nag", "ron", "soy", "gin", "don", "tug", "fay", "vic", "boo", "nam", "ave", "buy", "sop", "but", "orb", "fen", "paw", "his", "sub", "bob", "yea", "oft", "inn", "rod", "yam", "pew", "web", "hod", "hun", "gyp", "wei", "wis", "rob", "gad", "pie", "mon", "dog", "bib", "rub", "ere", "dig", "era", "cat", "fox", "bee", "mod", "day", "apr", "vie", "nev", "jam", "pam", "new", "aye", "ani", "and", "ibm", "yap", "can", "pyx", "tar", "kin", "fog", "hum", "pip", "cup", "dye", "lyx", "jog", "nun", "par", "wan", "fey", "bus", "oak", "bad", "ats", "set", "qom", "vat", "eat", "pus", "rev", "axe", "ion", "six", "ila", "lao", "mom", "mas", "pro", "few", "opt", "poe", "art", "ash", "oar", "cap", "lop", "may", "shy", "rid", "bat", "sum", "rim", "fee", "bmw", "sky", "maj", "hue", "thy", "ava", "rap", "den", "fla", "auk", "cox", "ibo", "hey", "saw", "vim", "sec", "ltd", "you", "its", "tat", "dew", "eva", "tog", "ram", "let", "see", "zit", "maw", "nix", "ate", "gig", "rep", "owe", "ind", "hog", "eve", "sam", "zoo", "any", "dow", "cod", "bed", "vet", "ham", "sis", "hex", "via", "fir", "nod", "mao", "aug", "mum", "hoe", "bah", "hal", "keg", "hew", "zed", "tow", "gog", "ass", "dem", "who", "bet", "gos", "son", "ear", "spy", "kit", "boy", "due", "sen", "oaf", "mix", "hep", "fur", "ada", "bin", "nil", "mia", "ewe", "hit", "fix", "sad", "rib", "eye", "hop", "haw", "wax", "mid", "tad", "ken", "wad", "rye", "pap", "bog", "gut", "ito", "woe", "our", "ado", "sin", "mad", "ray", "hon", "roy", "dip", "hen", "iva", "lug", "asp", "hui", "yak", "bay", "poi", "yep", "bun", "try", "lad", "elm", "nat", "wyo", "gym", "dug", "toe", "dee", "wig", "sly", "rip", "geo", "cog", "pas", "zen", "odd", "nan", "lay", "pod", "fit", "hem", "joy", "bum", "rio", "yon", "dec", "leg", "put", "sue", "dim", "pet", "yaw", "nub", "bit", "bur", "sid", "sun", "oil", "red", "doc", "moe", "caw", "eel", "dix", "cub", "end", "gem", "off", "yew", "hug", "pop", "tub", "sgt", "lid", "pun", "ton", "sol", "din", "yup", "jab", "pea", "bug", "gag", "mil", "jig", "hub", "low", "did", "tin", "get", "gte", "sox", "lei", "mig", "fig", "lon", "use", "ban", "flo", "nov", "jut", "bag", "mir", "sty", "lap", "two", "ins", "con", "ant", "net", "tux", "ode", "stu", "mug", "cad", "nap", "gun", "fop", "tot", "sow", "sal", "sic", "ted", "wot", "del", "imp", "cob", "way", "ann", "tan", "mci", "job", "wet", "ism", "err", "him", "all", "pad", "hah", "hie", "aim", "ike", "jed", "ego", "mac", "baa", "min", "com", "ill", "was", "cab", "ago", "ina", "big", "ilk", "gal", "tap", "duh", "ola", "ran", "lab", "top", "gob", "hot", "ora", "tia", "kip", "han", "met", "hut", "she", "sac", "fed", "goo", "tee", "ell", "not", "act", "gil", "rut", "ala", "ape", "rig", "cid", "god", "duo", "lin", "aid", "gel", "awl", "lag", "elf", "liz", "ref", "aha", "fib", "oho", "tho", "her", "nor", "ace", "adz", "fun", "ned", "coo", "win", "tao", "coy", "van", "man", "pit", "guy", "foe", "hid", "mai", "sup", "jay", "hob", "mow", "jot", "are", "pol", "arc", "lax", "aft", "alb", "len", "air", "pug", "pox", "vow", "got", "meg", "zoe", "amp", "ale", "bud", "gee", "pin", "dun", "pat", "ten", "mob"}
	r0 = findLadders("cet", "ism", wordList)
	println(fmt.Sprintf("result:%+v", r0))
}

/*
Return an empty list if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
 */

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	l0 := len(wordList)
	l1 := len(beginWord)
	mw := make(map[string]bool, l0)
	for _, v := range wordList {
		mw[v] = true
	}
	data := [][]string{{beginWord}}
	reverse := make(map[string]map[string]bool)
	path := 0
	var j byte
	var i int
	find := false
	for {
		if len(data[path]) == 0 || find {
			break
		}
		data = append(data, []string{})
		dup := map[string]bool{}
		for _, v := range data[path] {
			tmpWord := []byte(v)
			for i = 0; i < l1; i++ {
				tmpChar := tmpWord[i]
				for j = 'a'; j <= 'z'; j++ {
					tmpWord[i] = j
					if string(tmpWord) == beginWord {
						continue
					}
					if mw[string(tmpWord)] {
						data[path+1] = append(data[path+1], string(tmpWord))
						if reverse[string(tmpWord)] == nil {
							reverse[string(tmpWord)] = map[string]bool{v: true}
						} else {
							reverse[string(tmpWord)][v] = true
						}
						if string(tmpWord) == endWord {
							find = true
						} else {
							//mw[string(tmpWord)] = false
						}
						dup[string(tmpWord)] = true
					}
				}
				tmpWord[i] = tmpChar
			}
		}
		//mark true --> false
		for v := range dup {
			mw[v] = false
		}
		//println(fmt.Sprintf("%d,list:%+v,reverse:%+v", path, data, reverse))
		path++
	}
	if find {
		return _findLadders(endWord, beginWord, &reverse)
	}
	return nil
}

func _findLadders(end, begin string, mw *map[string]map[string]bool) [][]string {
	r := make([][]string, 0)
	if end == begin {
		return [][]string{{begin}}
	}
	for v := range (*mw)[end] {
		tmp := _findLadders(v, begin, mw)
		for _, v1 := range tmp {
			r = append(r, append(v1, end))
		}
	}
	return r
}
