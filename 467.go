package main

func main() {
	println(findSubstringInWraproundString("ttzabcffrst"))
}

func findSubstringInWraproundString(p string) int {
	count := make([]int, 26)
	maxLengthCur := 0

	for i := 0; i < len(p); i++ {
		if i > 0 && (p[i]-p[i-1] == 1 || p[i-1]-p[i] == 25) {
			maxLengthCur++
		} else {
			maxLengthCur = 1
		}

		index := p[i] - 'a'
		if maxLengthCur > count[index] {
			count[index] = maxLengthCur
		}
	}

	sum := 0
	for i := 0; i < 26; i++ {
		sum += count[i]
	}
	return sum
}
