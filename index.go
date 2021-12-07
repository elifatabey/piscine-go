package piscine

func Index(s string, toFind string) int {
	RuneS := []rune(s)
	l := len(RuneS)
	for i := 1; i < l; i++ {
		for j := 0; j <= l-i; j++ {
			substrings := string(RuneS[j : j+i])
			if len([]rune(substrings)) == len([]rune(toFind)) && substrings == toFind {
				return j
			}
		}
	}
	return -1
}
