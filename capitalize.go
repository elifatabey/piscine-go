package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	l := len(runes)
	snew := []rune{}
	for i := range s {
		if i >= 65 && i <= 90 {
			runes[i] = runes[i] - rune(32)
		}
	}
	if runes[0] >= 97 && runes[0] <= 122 {
		runes[0] = runes[0] - rune(32)
	}
	for i := range s {
		if runes[i] >= 32 && runes[i] <= 127 {
			snew = append(snew, runes[i])
		} else {
			snew = append(snew, runes[i])
			if i+1 < l-1 {
				for j := i + 1; j < l-1; j++ {
					if runes[j] == rune(32) {
						snew = append(snew)
					}
					if runes[j] >= 97 && runes[j] <= 122 {
						runes[j] = runes[j] - rune(32)
						snew = append(snew, runes[j])
					}
				}
			}
		}
	}
	return string(snew)
}
