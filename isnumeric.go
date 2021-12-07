package piscine

func IsNumeric(s string) bool {
	if s == "" {
		return false
	} else {
		count := 0
		for i := '0'; i <= '9'; i++ {
			for k := 0; k < len([]rune(s)); k++ {
				if rune(s[k]) == i {
					count = count + 1
				}
			}
		}
		if count == len([]rune(s)) {
			return true
		}
	}
	return false
}
