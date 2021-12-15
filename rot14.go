package piscine

func Rot14(s string) string {
	nb := rune(14)
	result := ""
	arr := []rune(s)

	for i := 0; i < len(arr); i++ {
		if arr[i] >= 'a' && arr[i] <= 'z' {
			if arr[i] >= 'm' {
				arr[i] = arr[i] - (nb - 2)
			} else {
				arr[i] = arr[i] + nb
			}
		} else if arr[i] >= 'A' && arr[i] <= 'Z' {
			if arr[i] >= 'M' {
				arr[i] = arr[i] - (nb - 2)
			} else {
				arr[i] = arr[i] + nb
			}
		}
		result += string(arr[i])
	}
	return result
}
