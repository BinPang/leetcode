package main

func main() {
	findWords([]string{"a", "b"})
}

func findWords(words []string) []string {
	tmp := map[uint8]uint8{
		'q': 1,
		'w': 1,
		'e': 1,
		'r': 1,
		't': 1,
		'y': 1,
		'u': 1,
		'i': 1,
		'o': 1,
		'p': 1,
		'Q': 1,
		'W': 1,
		'E': 1,
		'R': 1,
		'T': 1,
		'Y': 1,
		'U': 1,
		'I': 1,
		'O': 1,
		'P': 1,
		'a': 2,
		's': 2,
		'd': 2,
		'f': 2,
		'g': 2,
		'h': 2,
		'j': 2,
		'k': 2,
		'l': 2,
		'A': 2,
		'S': 2,
		'D': 2,
		'F': 2,
		'G': 2,
		'H': 2,
		'J': 2,
		'K': 2,
		'L': 2,
		'z': 3,
		'x': 3,
		'c': 3,
		'v': 3,
		'b': 3,
		'n': 3,
		'm': 3,
		'Z': 3,
		'X': 3,
		'C': 3,
		'V': 3,
		'B': 3,
		'N': 3,
		'M': 3,
	}

	r := make([]string, 0)
	var level uint8
	for _, v := range words {
		if len(v) == 0 {
			continue
		}
		level = tmp[v[0]]
		for _, v1 := range v {
			if tmp[uint8(v1)] != level {
				level = 4
				break
			}
		}
		if level != 4 {
			r = append(r, v)
		}
	}
	return r
}
