package piscine

func Capitalize(s string) string {
	L := ToLower(s)
	runeL := []rune(L)
	len := len(runeL)
	Lnew := Concat(ToUpper(string(L[0])), string(runeL[1:len]))
	LnewS := []rune(Lnew)
	newstrings := []rune{}
	for i := range LnewS {
		if IsAlpha(string(LnewS[i])) == true {
			newstrings = append(newstrings, LnewS[i])
		} else {
			newstrings = append(newstrings, LnewS[i])
			if i+1 > len-1 {
			} else {
				if LnewS[i+1] == rune(32) {
					newstrings = append(newstrings)
				}
				if IsLower(string(LnewS[i+1])) {
					LnewS[i+1] = LnewS[i+1] - rune(32)
				}
			}
		}
	}
	return string(newstrings)
}
