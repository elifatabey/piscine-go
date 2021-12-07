package piscine

func LastRune(s string) rune {
	sliceR := []rune(s)
	sliceL := len(sliceR)
	return sliceR[sliceL-1]
}
