package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	x := 0
	for range a {
		x++
	}
	for i := range a {
		if a[i] != x-1 {
			if f(a[i], a[i+1]) > 0 || f(a[i], a[i+1]) < 0 {
				return true
			}
			if f(a[i], a[i+1]) == 0 {
				return false
			}
		}
	}
	return false
}
