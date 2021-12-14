package piscine

func Index(s string, toFind string) int {
	runeS := []rune(s)
	l := len(runeS)
	if toFind == "" {
		return -1
	} else {
		for i := 1; i < l; i++ {
			for j := 0; j < l-i; j++ {
				substrings := string(runeS[j : j+i])
				if len([]rune(substrings)) == len([]rune(toFind)) && substrings == toFind {
					return j
				}
			}
		}
	}
	return -1
}
