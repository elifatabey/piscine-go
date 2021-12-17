package piscine

func Join(strs []string, sep string) string {
	finalstr := ""
	for i := 0; i < len(strs); i++ {
		finalstr += strs[i]
		if i < len(strs)-1 {
			finalstr += sep
		}
	}
	return finalstr
}
