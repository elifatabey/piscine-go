package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var count []bool
	for i := range a {
		if i != len(a)-1 {
			if !(f(a[i], a[i+1]) > 0 || f(a[i], a[i+1]) < 0) {
				count = append(count, false)
			} else {
				count = append(count, true)
			}
		}
	}
	for _, any := range count {
		if any == false {
			return false
		}
	}
	return true
}
