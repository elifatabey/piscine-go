package piscine

func SplitWhiteSpaces(s string) []string {
	var final []string
	countWords := 1
	lenStr := 0
	result := ""
	for c := range s {
		if isWhiteSpace(rune(s[c])) {
			countWords++
			if s[c] == ' ' && s[c+1] == ' ' {
				countWords = countWords - 1
			}
		}
		lenStr++
	}
	final = make([]string, countWords)
	i := 0
	for j, c := range s {
		if j+1 == lenStr {
			final[i] = result + string(s[j])
		}
		if isWhiteSpace(rune(c)) {
			if result != "" {
				if i <= countWords {
					final[i] = result
					i++
					result = ""
				}
			}
		} else {
			result += string(c)
		}
	}
	return final
}

func isWhiteSpace(r rune) bool {
	return (r == ' ' || r == '\n' || r == '\t' || r == 0)
}
