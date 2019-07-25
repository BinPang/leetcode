package main

func main() {
	println(detectCapitalUse("USA"))
	println(detectCapitalUse("FlaG"))
	println(detectCapitalUse("ala"))
	println(detectCapitalUse("alF"))
	println(detectCapitalUse("FFal"))
	println(detectCapitalUse("aFal"))
}

func detectCapitalUse(word string) bool {
	first := false
	second := false
	for k, v := range word {
		if k > 1 {
			if v >= 'A' && v <= 'Z' {
				if first && second {

				} else {
					return false
				}
			} else {
				if first && second {
					return false
				}
			}
		} else if k == 1 {
			if v >= 'A' && v <= 'Z' {
				if !first {
					return false
				}
				second = true
			}
		} else {
			if v >= 'A' && v <= 'Z' {
				first = true
			}
		}
	}
	return true
}
