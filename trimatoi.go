package piscine

func TrimAtoi(s string) int {
	result := 0
	negative := 1
	for _, c := range s {
		if c >= 48 && c <= 57 {
			result *= 10
			result += int(c - '0')
		}
		if c == 45 && result == 0 {
			negative *= -1
		}
	}
	return result * negative
}
