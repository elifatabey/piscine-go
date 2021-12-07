package piscine

func IsLower(s string) bool {
	count := 0
	for i := 'a'; i <= 'z'; i++ {
		for k := 0; k < len([]rune(s)); k++ {
			if rune(s[k]) == i {
				count = count + 1
			}
		}
	}
	if count == len([]rune(s)) {
		return true
	}
	return false
}
