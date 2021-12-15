package piscine

func Any(f func(string) bool, a []string) bool {
	x := 0
	for range a {
		x = x + 1
	}
	atotal := make([]bool, x)
	for i, value := range a {
		atotal[i] = f(value)
	}
	for _, che := range atotal {
		if che == true {
			return true
		}
	}
	return false
}
