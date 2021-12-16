package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, str := range *ptr {
		if str != "" {
			count++
		}
	}
	new := make([]string, count)
	i := 0
	for _, str2 := range *ptr {
		if str2 != "" {
			new[i] = str2
			i++
		}
	}
	*ptr = new
	return count
}
