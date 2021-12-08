package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i := range s {
		if IsLower(string(runes[i])) {
			runes[i] = runes[i] - rune(32)
		}
	}
	return string(runes)
}
