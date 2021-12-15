package piscine

func Map(f func(int) bool, a []int) []bool {
	x := 0
	for range a {
		x = x + 1
	}
	a2 := make([]bool, x)
	for i, value := range a {
		a2[i] = f(value)
	}
	return a2
}
