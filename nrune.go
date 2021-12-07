package piscine

func NRune(s string, n int) rune {
	RuneS := []rune(s)
	l := len(RuneS)
	if n > l || n <= 0 {
		return 0
	} else {
		return RuneS[n-1]
	}
}
