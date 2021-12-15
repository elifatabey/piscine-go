package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var count []int
	for i := range a {
		if i != len(a)-1 {
			if f(a[i], a[i+1]) > 0 {
				count = append(count, 0)
			}
			if f(a[i], a[i+1]) < 0 {
				count = append(count, 1)
			} else {
				count = append(count, 2)
			}
		}
		for i := range count {
			if count[i] == count[i+1] {
				return false
			}
		}
	}
	return true
}
