package piscine

func IsUpper(s string) bool {
	count:= 0
	runes := []rune(s)
	for _,i := range runes {
		if i >= 65 && i<=90 {
			count++
		}
	}
	if count == len(runes){
		return true
	}
return false
}
