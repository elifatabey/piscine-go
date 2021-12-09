package piscine

func Capitalize(s string) string {
	L := ToLower(s) // the function I write to make every letter to lower case.
	runeL := []rune(L)
	len := len(runeL) - 1
	Lnew := Concat(ToUpper(string(L[0])), string(runeL[1:len]))
	LnewS := []rune(Lnew)
	newstrings := []rune{}
	for i := range LnewS {
		if IsAlpha(string(LnewS[i])) == true {
			newstrings = append(newstrings, LnewS[i])
		} else {
			newstrings = append(newstrings, LnewS[i])
			if LnewS[i+1] == rune(32) {
				newstrings = append(newstrings)
			}
			if IsLower(string(LnewS[i+1])) {
				LnewS[i+1] = LnewS[i+1] - rune(32)
			}
		}
	}
	return string(newstrings)
}
