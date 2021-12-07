package piscine

func IsPrintable(s string) bool {
	count := 0
	for _, char := range s {
		if char >= rune(32) && char <= rune(127) {
			count++
		}
		if len(s) == count {
			return true
		}
	}
	return false
}
