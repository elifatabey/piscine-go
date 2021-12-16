package piscine

func Atoi(s string) int {
	runes := []rune(s)
	len := len(runes)
	negative := 1
	numbers := 0
	if checkdigit(s) == true {
		if runes[0] == '-' {
			negative *= -1
			for i := 1; i < len; i++ {
				numbers *= 10
				numbers += int(runes[i] - '0')
			}
		} else if runes[0] == '+' {
			for i := 1; i < len; i++ {
				numbers *= 10
				numbers += int(runes[i] - '0')
			}
		} else {
			for i := 0; i < len; i++ {
				numbers *= 10
				numbers += int(runes[i] - '0')
			}
		}
	}
	if checkdigit(s) == false {
		return 0
	}
	return numbers * negative
}

func checkdigit(s string) bool {
	runes := []rune(s)
	len := len(runes)
	if runes[0] == '-' || runes[0] == '+' {
		for i := 1; i < len; i++ {
			if !(runes[i] >= 48 && runes[i] <= 57) {
				return false
			}
		}
	} else {
		for _, char := range s {
			if !(char >= 48 && char <= 57) {
				return false
			}
		}
	}
	return true
}
