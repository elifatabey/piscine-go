package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i := range s {
		if runes[i] >= 97 && runes[i] <= 122 {
			runes[i] = runes[i] - rune(32)
		}
	}
	return string(runes)
}
