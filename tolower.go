package piscine

func ToLower(s string) string {
	runeS := []rune(s)
	for i := range s {
		if IsUpper(string(runeS[i])) {
			runeS[i] = runeS[i] + rune(32)
		}
	}
	return string(runeS)
}
