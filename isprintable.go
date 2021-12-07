package piscine

func IsPrintable(s string) bool {
	if s == "" {
		return true
	} else {
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
