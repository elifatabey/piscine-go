package piscine

func AlphaCount(s string) int {
	count := 0
	for i := 'a'; i <= 'z'; i++ {
		for k := 0; k < len([]rune(s)); k++ {
			if rune(s[k]) == i {
				count = count + 1
			}
		}
	}
	for i := 'A'; i <= 'Z'; i++ {
		for k := 0; k < len([]rune(s)); k++ {
			if rune(s[k]) == i {
				count = count + 1
			}
		}
	}
	return count
}
