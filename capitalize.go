package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	l := len(runes)
	snew := []rune{}
	for i := range runes {
		if runes[i] >= 65 && runes[i] <= 90 {
			runes[i] = runes[i] + rune(32)
		}
	}
	if runes[0] >= 97 && runes[0] <= 122 {
		runes[0] = runes[0] - rune(32)
	}
	for i := range runes {
		if runes[i] >= 48 && runes[i] <= 57 || runes[i] >= 65 && runes[i] <= 90 || runes[i] >= 97 && runes[i] <= 122 {
			snew = append(snew, runes[i])
		} else {
			snew = append(snew, runes[i])
			if i+1 > l-1 {
			} else {
				if runes[i+1] == rune(32) {
					snew = append(snew)
				}
				if runes[i+1] >= 97 && runes[i+1] <= 122 {
					runes[i+1] = runes[i+1] - rune(32)
				}
			}
		}
	}
	return string(snew)
}
