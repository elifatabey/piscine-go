package piscine

func TrimAtoi(s string) int {
	total := 0
	negative := 1
	for i := range s {
		if i >= 48 && i <= 57 {
			total *= 10
			total += int(i - '0')
		}
		if i == '-' && i == 0 {
			negative *= -1
		}
	}
	return total * negative
}
