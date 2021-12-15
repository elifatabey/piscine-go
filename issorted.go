package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var count []bool
	for i := range a {
		if i != len(a)-1 {
			if f(a[i], a[i+1]) > 0 || f(a[i], a[i+1]) < 0 {
				count = append(count, true)
			} else {
				count = append(count, false)
			}
		}
	}
	for _, any := range count {
		if any == true {
			return true
		}
	}
	return false
}
