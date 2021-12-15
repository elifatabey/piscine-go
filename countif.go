package piscine

func CountIf(f func(string) bool, tab []string) int {
	x := 0
	for range tab {
		x = x + 1
	}
	atotal := make([]bool, x)
	for i, value := range tab {
		atotal[i] = f(value)
	}
	count := 0
	for _, che := range atotal {
		if che == true {
			count++
		}
	}
	return count
}
