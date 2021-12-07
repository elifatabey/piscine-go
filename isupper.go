package piscine

func IsUpper(s string) bool {
	count := 0
	for i := 'A'; i <= 'Z'; i++ {
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
