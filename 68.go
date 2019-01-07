package main

import "fmt"

/*
Input:
words = ["What","must","be","acknowledgment","shall","be"]
maxWidth = 16
Output:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]

Input:
words = ["Science","is","what","we","understand","well","enough","to","explain",
         "to","a","computer.","Art","is","everything","else","we","do"]
maxWidth = 20
Output:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]
 */
func main() {
	fullJustify([]string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what",
		"you", "can", "do", "for", "your", "country"}, 16)
	println()
	fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
		"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20)
	println()
	fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	println()
	fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16)
	println()
}

func fullJustify(words []string, maxWidth int) []string {
	//l := len(words)
	r := make([]string, 0)
	item := make([][]int, 1)
	sum := make([]int, 1)
	index := 0
	for i, v := range words {
		if sum[index]+len(item[index])+len(v) > maxWidth {
			item = append(item, []int{i})
			sum = append(sum, len(v))
			index++
		} else {
			item[index] = append(item[index], i)
			sum[index] += len(v)
		}
	}
	lItem := len(item)
	for i := 0; i < lItem-1; i++ {
		lOneItem := len(item[i])
		if lOneItem == 1 {
			r = append(r, fmt.Sprintf(fmt.Sprintf("%%-%ds", maxWidth), words[item[i][0]]))
		} else {
			tmpR := ""
			each := (maxWidth - sum[i]) / (lOneItem - 1)
			first := (maxWidth - sum[i]) % (lOneItem - 1)
			for j := 0; j < lOneItem-1; j++ {
				if first > 0 {
					tmpR += fmt.Sprintf(fmt.Sprintf("%%-%ds", each+1+len(words[item[i][j]])), words[item[i][j]])
					first--
				} else {
					tmpR += fmt.Sprintf(fmt.Sprintf("%%-%ds", each+len(words[item[i][j]])), words[item[i][j]])
				}
			}
			tmpR += words[item[i][lOneItem-1]]
			r = append(r, tmpR)
		}
	}
	//last item
	tmpR := ""
	lOneItem := len(item[lItem-1])
	for i, v := range item[lItem-1] {
		if i == lOneItem-1 {
			tmpR += words[v]
		} else {
			tmpR += words[v] + " "
		}
	}
	r = append(r, fmt.Sprintf(fmt.Sprintf("%%-%ds", maxWidth), tmpR))
	//for _, v := range r {
	//	println(fmt.Sprintf("%s|%d", v, len(v)))
	//}
	return r
}
