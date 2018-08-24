package main

func main() {
	_printArray(letterCombinations("23"))
	_printArray(letterCombinations("234"))
}

func letterCombinations(digits string) []string {
	l := len(digits)
	if l == 0 {
		return nil
	}

	t := _letterCombinations(digits)
	r := make([]string, len(t))
	for i := range t  {
		r[i] = string(t[i])
	}
	return r
}


func _letterCombinations(digits string) [][]byte  {
	m := map[uint8][]byte{
		50:{'a', 'b', 'c'},
		51:{'d', 'e', 'f'},
		52:{'g', 'h', 'i'},
		53:{'j', 'k', 'l'},
		54:{'m', 'n', 'o'},
		55:{'p', 'q', 'r', 's'},
		56:{'t', 'u', 'v'},
		57:{'w', 'x', 'y', 'z'},
	}

	if len(digits) == 1 {
		tmpT := make([][]byte, len(m[digits[0]]))
		for i, v := range m[digits[0]]{
			tmpT[i] =[]byte{v}
		}
		return tmpT
	}
	t := _letterCombinations(digits[1:])
	r := make([][]byte, len(t)*len(m[digits[0]]))
	for i, v := range m[digits[0]]{
		for i1, v1 := range t {
			r[i*len(t)+i1] = append([]byte{v}, v1...)
		}
	}
	return r
}

func _printArray(t []string)  {
	for _, v := range t{
		print(v, ",")
	}
	println("")
}




