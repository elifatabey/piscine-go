package piscine

func Split(s, sep string) []string {
	resp := []string{}
	new := ""
	for i := 0; i < len(s); i++ {
		if setsep(s, sep, i) {
			if new != "" {
				resp = append(resp, new)
				new = ""
				i = i + len(sep) - 1
			}
		} else {
			new = new + string(s[i])
		}
	}
	if new != "" {
		resp = append(resp, new)
	}
	return resp
}

func setsep(s, sep string, i int) bool {
	j := 0
	for j < len(sep) && i < len(s) {
		if s[i] != sep[j] {
			return false
		}
		i++
		j++
	}
	return true
}
