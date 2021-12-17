package piscine

func Unmatch(a []int) int {
	var result []int
	var final int = 0
	for i := range a {
		result = append(result, a[i+1:]...)
		result = append(result, a[:i]...)
		for _, look := range result {
			if look == a[i] {
			} else {
				final = a[i]
			}
		}
	}
	return final
}
