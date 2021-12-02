package piscine

func StrRev(s string) string {
	runes := []rune(s)
	n := len(runes) - 1

	runesRev := []rune{}
	for i := n; i >= 0; i-- {
		runesRev = append(runesRev, runes[i])
	}
	return string(runesRev)
}
