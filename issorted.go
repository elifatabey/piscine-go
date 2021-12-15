package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	count := true
	for i := range a {
		if a[i] != len(a)-1 {
			if !(f(a[i], a[i+1]) > 0 || f(a[i], a[i+1]) < 0) {
				count = false
			}
		}
	}
	return count
}
