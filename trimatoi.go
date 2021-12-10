package piscine

func TrimAtoi(s string) int {
	total := 0
	negative := 1
	for _, i := range s {
		if i >= 48 && i <= 57 {
			total *= 10
			total += int(i - '0')
		}
		if i == 45 && total == 0 {
			negative *= -1
		}
	}
	return total * negative
}
