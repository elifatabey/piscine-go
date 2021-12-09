package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i := range runes {
		if i >= 97 && i <= 122 {
			i = i - 32
		}
	}
	return string(runes)
}
