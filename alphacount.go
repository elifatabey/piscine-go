package piscine

func AlphaCount(s string) int {
	count := 0
	runeS := []rune(s)
	for _, i := range runeS {
		if i >= 65 && i <= 90 || i >= 97 && i <= 122 {
			count = count + 1
		}
	}
	return count
}
